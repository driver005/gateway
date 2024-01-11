package services

import (
	"context"
	"math"
	"slices"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type Total struct {
	Total    float64
	TaxTotal float64
}

type ShippingMethodTotals struct {
	Price            float64
	TaxTotal         float64
	Total            float64
	Subtotal         float64
	OriginalTotal    float64
	OriginalTaxTotal float64
	TaxLines         []models.ShippingMethodTaxLine
}

type GetShippingMethodTotalsOptions struct {
	IncludeTax         bool
	UseTaxLines        bool
	CalculationContext interfaces.TaxCalculationContext
}

type LineItemTotalsOptions struct {
	IncludeTax         bool
	UseTaxLines        bool
	ExcludeGiftCards   bool
	CalculationContext *interfaces.TaxCalculationContext
}

type GetLineItemTotalOptions struct {
	IncludeTax       bool
	ExcludeDiscounts bool
}

type TotalsServiceProps struct {
	TaxProviderService     TaxProviderService
	NewTotalsService       NewTotalsService
	TaxCalculationStrategy interfaces.ITaxCalculationStrategy
}

type GetTotalsOptions struct {
	ExcludeGiftCards bool
	ForceTaxes       bool
}

type AllocationMapOptions struct {
	ExcludeGiftCards bool
	ExcludeDiscounts bool
}

type CalculationContextOptions struct {
	IsReturn         bool
	ExcludeShipping  bool
	ExcludeGiftCards bool
	ExcludeDiscounts bool
}

type TotalsService struct {
	ctx context.Context
	r   Registry
}

func NewTotalService(
	r Registry,
) *TotalsService {
	return &TotalsService{
		context.Background(),
		r,
	}
}

func (s *TotalsService) SetContext(context context.Context) *TotalsService {
	s.ctx = context
	return s
}

func (s *TotalsService) GetTotal(cart *models.Cart, order *models.Order, options GetTotalsOptions) (float64, *utils.ApplictaionError) {
	subtotal, err := s.GetSubtotal(cart, order, types.SubtotalOptions{})
	if err != nil {
		return 0, err
	}
	taxTotal, err := s.GetTaxTotal(cart, order, options.ForceTaxes)
	if err != nil {
		return 0, err
	}
	discountTotal, err := s.GetDiscountTotal(cart, order)
	if err != nil {
		return 0, err
	}
	giftCardTotal, err := s.GetGiftCardTotal(cart, order, map[string]interface{}{})
	if err != nil {
		return 0, err
	}
	shippingTotal, err := s.GetShippingTotal(cart, order)
	if err != nil {
		return 0, err
	}
	return subtotal + taxTotal + shippingTotal - discountTotal - giftCardTotal.Total, nil
}

func (s *TotalsService) GetPaidTotal(order *models.Order) float64 {
	total := 0.0
	for _, payment := range order.Payments {
		total += payment.Amount
	}
	return total
}

func (s *TotalsService) GetSwapTotal(order *models.Order) float64 {
	swapTotal := 0.0
	for _, swap := range order.Swaps {
		swapTotal += swap.DifferenceDue
	}
	return swapTotal
}

func (s *TotalsService) GetShippingMethodTotals(shippingMethod models.ShippingMethod, cart *models.Cart, order *models.Order, opts GetShippingMethodTotalsOptions) (*ShippingMethodTotals, *utils.ApplictaionError) {
	calculationContext, err := s.GetCalculationContext(cart, order, CalculationContextOptions{ExcludeShipping: true})
	if err != nil {
		return nil, err
	}
	calculationContext.ShippingMethods = []models.ShippingMethod{shippingMethod}
	totals := &ShippingMethodTotals{
		Price:            shippingMethod.Price,
		OriginalTotal:    shippingMethod.Price,
		Total:            shippingMethod.Price,
		Subtotal:         shippingMethod.Price,
		OriginalTaxTotal: 0,
		TaxTotal:         0,
		TaxLines:         shippingMethod.TaxLines,
	}
	if opts.IncludeTax {
		if order != nil {
			totals.OriginalTaxTotal = math.Round(float64(totals.Price) * (order.TaxRate / 100))
			totals.TaxTotal = math.Round(float64(totals.Price) * (order.TaxRate / 100))
		} else if len(totals.TaxLines) == 0 {
			var lineItem []models.LineItem
			if cart != nil {
				lineItem = cart.Items
			}
			if order != nil {
				lineItem = order.Items
			}
			shipping, _, err := s.r.TaxProviderService().SetContext(s.ctx).GetTaxLines(lineItem, calculationContext)
			if err != nil {
				return nil, err
			}
			for _, smtl := range shipping {
				if smtl.ShippingMethodId.UUID == shippingMethod.Id {
					totals.TaxLines = append(totals.TaxLines, smtl)
				}
			}
			if len(totals.TaxLines) == 0 && order != nil {
				return nil, utils.NewApplictaionError(
					utils.CONFLICT,
					"Tax Lines must be joined on shipping method to calculate taxes",
					"500",
					nil,
				)
			}
		}
		if len(totals.TaxLines) > 0 {
			includesTax := true
			totals.OriginalTaxTotal, err = s.r.TaxCalculationStrategy().Calculate([]models.LineItem{}, totals.TaxLines, calculationContext)
			if err != nil {
				return nil, err
			}
			totals.TaxTotal = totals.OriginalTaxTotal
			if includesTax {
				totals.Subtotal -= totals.TaxTotal
			} else {
				totals.OriginalTotal += totals.OriginalTaxTotal
				totals.Total += totals.TaxTotal
			}
		}
	}
	hasFreeShipping := false
	if cart != nil {
		for _, discount := range cart.Discounts {
			if discount.Rule.Type == models.DiscountRuleFreeShipping {
				hasFreeShipping = true
				break
			}
		}
	}
	if order != nil {
		for _, discount := range order.Discounts {
			if discount.Rule.Type == models.DiscountRuleFreeShipping {
				hasFreeShipping = true
				break
			}
		}
	}

	if hasFreeShipping {
		totals.Total = 0
		totals.Subtotal = 0
		totals.TaxTotal = 0
	}
	return totals, nil
}

func (s *TotalsService) GetSubtotal(cart *models.Cart, order *models.Order, opts types.SubtotalOptions) (float64, *utils.ApplictaionError) {
	subtotal := 0.0
	if cart.Items == nil || order.Items == nil {
		return subtotal, nil
	}
	getLineItemSubtotal := func(item models.LineItem) (float64, *utils.ApplictaionError) {
		totals, err := s.GetLineItemTotals(item, cart, order, &LineItemTotalsOptions{IncludeTax: true, ExcludeGiftCards: true})
		if err != nil {
			return 0, err
		}
		return totals.Subtotal, nil
	}

	if cart != nil {
		for _, item := range cart.Items {
			if opts.ExcludeNonDiscounts {
				if item.AllowDiscounts {
					lineItemSubtotal, err := getLineItemSubtotal(item)
					if err != nil {
						return 0, err
					}
					subtotal += lineItemSubtotal
				}
				continue
			}
			lineItemSubtotal, err := getLineItemSubtotal(item)
			if err != nil {
				return 0, err
			}
			subtotal += lineItemSubtotal
		}
	}
	if order != nil {
		for _, item := range order.Items {
			if opts.ExcludeNonDiscounts {
				if item.AllowDiscounts {
					lineItemSubtotal, err := getLineItemSubtotal(item)
					if err != nil {
						return 0, err
					}
					subtotal += lineItemSubtotal
				}
				continue
			}
			lineItemSubtotal, err := getLineItemSubtotal(item)
			if err != nil {
				return 0, err
			}
			subtotal += lineItemSubtotal
		}
	}

	return s.Rounded(subtotal), nil
}

func (s *TotalsService) GetShippingTotal(cart *models.Cart, order *models.Order) (float64, *utils.ApplictaionError) {
	var shippingMethods []models.ShippingMethod
	if cart != nil {
		shippingMethods = cart.ShippingMethods
	}
	if order != nil {
		shippingMethods = order.ShippingMethods
	}
	total := 0.0
	for _, shippingMethod := range shippingMethods {
		totals, err := s.GetShippingMethodTotals(shippingMethod, cart, order, GetShippingMethodTotalsOptions{IncludeTax: true})
		if err != nil {
			return 0, err
		}
		total += totals.Subtotal
	}
	return total, nil
}

func (s *TotalsService) GetTaxTotal(cart *models.Cart, order *models.Order, forceTaxes bool) (float64, *utils.ApplictaionError) {
	if cart != nil && !forceTaxes && !cart.Region.AutomaticTaxes {
		return 0, nil
	}
	calculationContext, err := s.GetCalculationContext(cart, order, CalculationContextOptions{})
	if err != nil {
		return 0, err
	}
	giftCardTotal, err := s.GetGiftCardTotal(cart, order, map[string]interface{}{})
	if err != nil {
		return 0, err
	}
	var taxLines []interface{}
	if order != nil {
		taxLinesJoined := true
		for _, item := range order.Items {
			if item.TaxLines == nil {
				taxLinesJoined = false
				break
			}
		}
		if !taxLinesJoined {
			subtotal, err := s.GetSubtotal(cart, order, types.SubtotalOptions{})
			if err != nil {
				return 0, err
			}
			shippingTotal, err := s.GetShippingTotal(cart, order)
			if err != nil {
				return 0, err
			}
			discountTotal, err := s.GetDiscountTotal(cart, order)
			if err != nil {
				return 0, err
			}
			return s.Rounded((subtotal - discountTotal - giftCardTotal.Total + shippingTotal) * (order.TaxRate / 100)), nil
		}
		if order.TaxRate == 0 {
			for _, item := range order.Items {
				taxLines = append(taxLines, item.TaxLines)
			}
			for _, shippingMethod := range order.ShippingMethods {
				taxLines = append(taxLines, shippingMethod.TaxLines)
			}
		} else {
			subtotal, err := s.GetSubtotal(cart, order, types.SubtotalOptions{})
			if err != nil {
				return 0, err
			}
			shippingTotal, err := s.GetShippingTotal(cart, order)
			if err != nil {
				return 0, err
			}
			discountTotal, err := s.GetDiscountTotal(cart, order)
			if err != nil {
				return 0, err
			}
			return s.Rounded((subtotal - discountTotal - giftCardTotal.Total + shippingTotal) * (order.TaxRate / 100)), nil
		}
	} else {
		shipping, lineItem, err := s.r.TaxProviderService().SetContext(s.ctx).GetTaxLines(cart.Items, calculationContext)
		if err != nil {
			return 0, err
		}

		taxLines = append(taxLines, lineItem)
		taxLines = append(taxLines, shipping)

		if cart.Type == "swap" {
			for _, item := range cart.Items {
				if item.IsReturn {
					if item.TaxLines == nil {
						return 0, utils.NewApplictaionError(
							utils.CONFLICT,
							"Return Line Items must join tax lines",
							"500",
							nil,
						)
					}
				}
				taxLines = append(taxLines, item.TaxLines)
			}

		}
	}
	var lineItem []models.LineItem
	if cart != nil {
		lineItem = cart.Items
	}
	if order != nil {
		lineItem = order.Items
	}
	toReturn, err := s.r.TaxCalculationStrategy().Calculate(lineItem, taxLines, calculationContext)
	if err != nil {
		return 0, err
	}
	if cart != nil && cart.Region.GiftCardsTaxable {
		return s.Rounded(toReturn - giftCardTotal.TaxTotal), nil
	}
	return s.Rounded(toReturn), nil
}

func (s *TotalsService) GetAllocationMap(cart *models.Cart, order *models.Order, options AllocationMapOptions) (types.LineAllocationsMap, *utils.ApplictaionError) {
	allocationMap := make(types.LineAllocationsMap)
	if !options.ExcludeDiscounts {
		var discount models.Discount
		if cart != nil {
			for _, d := range cart.Discounts {
				if d.Rule.Type != models.DiscountRuleFreeShipping {
					discount = d
					break
				}
			}

		}
		if order != nil {
			for _, d := range order.Discounts {
				if d.Rule.Type != models.DiscountRuleFreeShipping {
					discount = d
					break
				}
			}
		}

		lineDiscounts := s.GetLineDiscounts(cart, order, discount)
		for _, ld := range lineDiscounts {
			adjustmentAmount := ld.Amount + ld.CustomAdjustmentsAmount
			if entity, ok := allocationMap[ld.Item.Id]; ok {
				entity.Discount = &types.DiscountAllocation{
					Amount:     adjustmentAmount,
					UnitAmount: adjustmentAmount / float64(ld.Item.Quantity),
				}

				allocationMap[ld.Item.Id] = entity
			} else {
				allocationMap[ld.Item.Id] = struct {
					GiftCard *types.GiftCardAllocation `json:"gift_card,omitempty"`
					Discount *types.DiscountAllocation `json:"discount,omitempty"`
				}{
					Discount: &types.DiscountAllocation{
						Amount:     adjustmentAmount,
						UnitAmount: adjustmentAmount / float64(ld.Item.Quantity),
					},
				}
			}
		}
	}

	return allocationMap, nil
}

func (s *TotalsService) GetRefundedTotal(order *models.Order) float64 {
	if order.Refunds == nil {
		return 0
	}
	total := 0.0
	for _, refund := range order.Refunds {
		total += refund.Amount
	}
	return s.Rounded(total)
}

func (s *TotalsService) GetLineItemRefund(order *models.Order, lineItem models.LineItem) (float64, *utils.ApplictaionError) {
	allocationMap, err := s.GetAllocationMap(nil, order, AllocationMapOptions{})
	if err != nil {
		return 0, err
	}
	includesTax := true
	discountAmount := allocationMap[lineItem.Id].GiftCard.UnitAmount * float64(lineItem.Quantity)
	var taxAmountIncludedInPrice float64

	if order.TaxRate != 0.0 {
		if includesTax {
			taxAmountIncludedInPrice = s.Rounded(utils.CalculatePriceTaxAmount(lineItem.UnitPrice, order.TaxRate/100, includesTax))
		}
		lineSubtotal := (lineItem.UnitPrice-taxAmountIncludedInPrice)*float64(lineItem.Quantity) - discountAmount
		taxRate := order.TaxRate / 100
		return s.Rounded(lineSubtotal * (1 + taxRate)), nil
	}
	if lineItem.TaxLines == nil {
		return 0, utils.NewApplictaionError(
			utils.CONFLICT,
			"Tax calculation did not receive tax_lines",
			"500",
			nil,
		)
	}
	taxRate := 0.0
	for _, tl := range lineItem.TaxLines {
		taxRate += tl.Rate / 100
	}

	if includesTax {
		taxAmountIncludedInPrice = s.Rounded(utils.CalculatePriceTaxAmount(lineItem.UnitPrice, taxRate, includesTax))
	}
	lineSubtotal := (lineItem.UnitPrice-taxAmountIncludedInPrice)*float64(lineItem.Quantity) - discountAmount
	taxTotal := 0.0
	for _, tl := range lineItem.TaxLines {
		taxRate := tl.Rate / 100
		taxTotal += s.Rounded(lineSubtotal * taxRate)
	}
	return s.Rounded(lineSubtotal + taxTotal), nil
}

func (s *TotalsService) GetRefundTotal(order *models.Order, lineItems []models.LineItem) (float64, *utils.ApplictaionError) {
	var itemIds uuid.UUIDs
	for _, item := range order.Items {
		itemIds = append(itemIds, item.Id)
	}
	if order.Swaps != nil {
		for _, swap := range order.Swaps {
			for _, el := range swap.AdditionalItems {
				itemIds = append(itemIds, el.Id)
			}
		}
	}
	if order.Claims != nil {
		for _, claim := range order.Claims {
			for _, el := range claim.AdditionalItems {
				itemIds = append(itemIds, el.Id)
			}
		}
	}
	refunds := []float64{}
	for _, item := range lineItems {
		if !slices.Contains(itemIds, item.Id) {
			return 0, utils.NewApplictaionError(
				utils.CONFLICT,
				"Line item does not exist on order",
				"500",
				nil,
			)
		}
		refund, err := s.GetLineItemRefund(order, item)
		if err != nil {
			return 0, err
		}
		refunds = append(refunds, refund)
	}
	refund := 0.0
	for _, r := range refunds {
		refund += r
	}
	return s.Rounded(refund), nil
}

func (s *TotalsService) CalculateDiscount(lineItem models.LineItem, variant uuid.UUID, variantPrice float64, value float64, discountType models.DiscountRuleType) types.LineDiscount {
	if !lineItem.AllowDiscounts {
		return types.LineDiscount{
			LineItem: lineItem,
			Variant:  variant,
			Amount:   0,
		}
	}
	if discountType == models.DiscountRulePersentage {
		return types.LineDiscount{
			LineItem: lineItem,
			Variant:  variant,
			Amount:   ((variantPrice * float64(lineItem.Quantity)) / 100) * value,
		}
	} else {
		return types.LineDiscount{
			LineItem: lineItem,
			Variant:  variant,
			Amount:   math.Min(value*float64(lineItem.Quantity), variantPrice*float64(lineItem.Quantity)),
		}
	}
}

func (s *TotalsService) GetAllocationItemDiscounts(discount models.Discount, cart models.Cart) []types.LineDiscount {
	discounts := []types.LineDiscount{}
	for _, item := range cart.Items {
		discounts = append(discounts, types.LineDiscount{
			LineItem: item,
			Variant:  item.Variant.Id,
			Amount:   s.GetLineItemDiscountAdjustment(item, discount),
		})
	}
	return discounts
}

func (s *TotalsService) GetLineItemDiscountAdjustment(lineItem models.LineItem, discount models.Discount) float64 {
	for _, adjustment := range lineItem.Adjustments {
		if adjustment.DiscountId.UUID == discount.Id {
			return adjustment.Amount
		}
	}
	return 0
}

func (s *TotalsService) GetLineItemAdjustmentsTotal(cart *models.Cart, order *models.Order) float64 {
	if cart == nil {
		return 0
	}
	if order == nil {
		return 0
	}

	total := 0.0

	if cart != nil {
		for _, item := range cart.Items {
			if item.Adjustments != nil {
				for _, adjustment := range item.Adjustments {
					total += adjustment.Amount
				}
			}
		}
	}
	if order != nil {
		for _, item := range order.Items {
			if item.Adjustments != nil {
				for _, adjustment := range item.Adjustments {
					total += adjustment.Amount
				}
			}
		}
	}

	return total
}

func (s *TotalsService) GetLineDiscounts(cart *models.Cart, order *models.Order, discount models.Discount) []types.LineDiscountAmount {
	merged := append([]models.LineItem{}, cart.Items...)
	if order.Items != nil {
		merged = append(merged, order.Items...)
	}

	if order != nil {
		if order.Swaps != nil {
			for _, swap := range order.Swaps {
				merged = append(merged, swap.AdditionalItems...)
			}
			if order.Claims != nil {
				for _, claim := range order.Claims {
					merged = append(merged, claim.AdditionalItems...)
				}

			}
		}
	}

	lineDiscounts := []types.LineDiscountAmount{}
	for _, item := range merged {
		adjustments := item.Adjustments
		discountAdjustments := []models.LineItemAdjustment{}
		for _, adjustment := range adjustments {
			if adjustment.DiscountId.UUID == discount.Id {
				discountAdjustments = append(discountAdjustments, adjustment)
			}
		}
		customAdjustments := []models.LineItemAdjustment{}
		for _, adjustment := range adjustments {
			if adjustment.DiscountId.UUID == uuid.Nil {
				customAdjustments = append(customAdjustments, adjustment)
			}
		}
		amount := 0.0
		for _, adjustment := range discountAdjustments {
			amount += adjustment.Amount
		}
		customAdjustmentsAmount := 0.0
		for _, adjustment := range customAdjustments {
			customAdjustmentsAmount += adjustment.Amount
		}
		lineDiscounts = append(lineDiscounts, types.LineDiscountAmount{
			Item:                    item,
			Amount:                  amount,
			CustomAdjustmentsAmount: customAdjustmentsAmount,
		})
	}
	return lineDiscounts
}

func (s *TotalsService) GetLineItemTotals(lineItem models.LineItem, cart *models.Cart, order *models.Order, options *LineItemTotalsOptions) (*models.LineItem, *utils.ApplictaionError) {
	var err *utils.ApplictaionError
	calculationContext := options.CalculationContext
	if calculationContext == nil {
		calculationContext, err = s.GetCalculationContext(cart, order, CalculationContextOptions{
			ExcludeShipping:  true,
			ExcludeGiftCards: options.ExcludeGiftCards,
		})
		if err != nil {
			return nil, err
		}
	}
	lineItemAllocation := calculationContext.AllocationMap[lineItem.Id]
	subtotal := lineItem.UnitPrice * float64(lineItem.Quantity)
	feature := true
	if feature && lineItem.IncludesTax && options.IncludeTax {
		subtotal = 0
	}
	rawDiscountTotal := lineItemAllocation.Discount.Amount
	discountTotal := math.Round(rawDiscountTotal)
	lineItemTotals := &models.LineItem{
		UnitPrice:        lineItem.UnitPrice,
		Quantity:         lineItem.Quantity,
		Subtotal:         subtotal,
		DiscountTotal:    discountTotal,
		Total:            subtotal - discountTotal,
		OriginalTotal:    subtotal,
		OriginalTaxTotal: 0,
		TaxTotal:         0,
		TaxLines:         lineItem.TaxLines,
		RawDiscountTotal: rawDiscountTotal,
	}
	if options.IncludeTax {
		if order != nil && order.TaxRate != 0.0 {
			taxRate := order.TaxRate / 100
			includesTax := true
			taxIncludedInPrice := 0
			if !lineItem.IncludesTax {
				taxIncludedInPrice = int(math.Round(utils.CalculatePriceTaxAmount(lineItem.UnitPrice, taxRate, includesTax)))
			}
			lineItemTotals.Subtotal = (lineItem.UnitPrice - float64(taxIncludedInPrice)) * float64(lineItem.Quantity)
			lineItemTotals.Total = lineItemTotals.Subtotal
			lineItemTotals.OriginalTaxTotal = lineItemTotals.Subtotal * taxRate
			lineItemTotals.TaxTotal = (lineItemTotals.Subtotal - discountTotal) * taxRate
			lineItemTotals.Total += lineItemTotals.TaxTotal
			lineItemTotals.OriginalTotal += lineItemTotals.OriginalTaxTotal
		} else {
			var taxLines []models.LineItemTaxLine
			if options.UseTaxLines || order != nil {
				if lineItem.TaxLines == nil && lineItem.VariantId.UUID != uuid.Nil {
					return nil, utils.NewApplictaionError(
						utils.CONFLICT,
						"Tax Lines must be joined on items to calculate taxes",
						"500",
						nil,
					)
				}
				taxLines = lineItem.TaxLines
			} else {
				if lineItem.IsReturn {
					if lineItem.TaxLines == nil && lineItem.VariantId.UUID != uuid.Nil {
						return nil, utils.NewApplictaionError(
							utils.CONFLICT,
							"Return Line Items must join tax lines",
							"500",
							nil,
						)
					}
					taxLines = lineItem.TaxLines
				} else {
					_, lineitem, err := s.r.TaxProviderService().SetContext(s.ctx).GetTaxLines([]models.LineItem{lineItem}, calculationContext)
					if err != nil {
						return nil, err
					}
					taxLines = append(taxLines, lineitem...)
				}
			}
			lineItemTotals.TaxLines = taxLines
		}
	}
	if len(lineItemTotals.TaxLines) > 0 {
		var err *utils.ApplictaionError
		lineItemTotals.TaxTotal, err = s.r.TaxCalculationStrategy().Calculate([]models.LineItem{lineItem}, lineItemTotals.TaxLines, calculationContext)
		if err != nil {
			return nil, err
		}
		noDiscountContext := calculationContext
		lineItemTotals.OriginalTaxTotal, err = s.r.TaxCalculationStrategy().Calculate([]models.LineItem{lineItem}, lineItemTotals.TaxLines, noDiscountContext)
		if err != nil {
			return nil, err
		}
		if feature && lineItem.IncludesTax {
			lineItemTotals.Subtotal += lineItem.UnitPrice*float64(lineItem.Quantity) - lineItemTotals.OriginalTaxTotal
			lineItemTotals.Total += lineItemTotals.Subtotal
			lineItemTotals.OriginalTotal += lineItemTotals.Subtotal
		}
		lineItemTotals.Total += lineItemTotals.TaxTotal
		lineItemTotals.OriginalTotal += lineItemTotals.OriginalTaxTotal
	}
	return lineItemTotals, nil
}

func (s *TotalsService) GetLineItemTotal(lineItem models.LineItem, cart *models.Cart, order *models.Order, options GetLineItemTotalOptions) (float64, *utils.ApplictaionError) {
	lineItemTotals, err := s.GetLineItemTotals(lineItem, cart, order, &LineItemTotalsOptions{
		IncludeTax: options.IncludeTax,
	})
	if err != nil {
		return 0, err
	}
	toReturn := lineItemTotals.Subtotal
	if !options.ExcludeDiscounts {
		toReturn += lineItemTotals.DiscountTotal
	}
	if options.IncludeTax {
		toReturn += lineItemTotals.TaxTotal
	}
	return toReturn, nil
}

func (s *TotalsService) GetGiftCardableAmount(cart *models.Cart, order *models.Order) (float64, *utils.ApplictaionError) {
	if cart.Region.GiftCardsTaxable {
		subtotal, err := s.GetSubtotal(cart, order, types.SubtotalOptions{})
		if err != nil {
			return 0, err
		}
		discountTotal, err := s.GetDiscountTotal(cart, order)
		if err != nil {
			return 0, err
		}
		return subtotal - discountTotal, nil
	}

	if order.Region.GiftCardsTaxable {
		subtotal, err := s.GetSubtotal(cart, order, types.SubtotalOptions{})
		if err != nil {
			return 0, err
		}
		discountTotal, err := s.GetDiscountTotal(cart, order)
		if err != nil {
			return 0, err
		}
		return subtotal - discountTotal, nil
	}

	return s.GetTotal(cart, order, GetTotalsOptions{
		ExcludeGiftCards: true,
	})
}

func (s *TotalsService) GetGiftCardTotal(cart *models.Cart, order *models.Order, opts map[string]interface{}) (*Total, *utils.ApplictaionError) {
	var giftCardable float64
	if val, ok := opts["gift_cardable"].(float64); ok {
		giftCardable = val
	} else {
		subtotal, err := s.GetSubtotal(cart, order, types.SubtotalOptions{})
		if err != nil {
			return nil, err
		}
		discountTotal, err := s.GetDiscountTotal(cart, order)
		if err != nil {
			return nil, err
		}
		giftCardable = subtotal - discountTotal
	}

	var giftCardTotals *Total
	if cart != nil {
		var err *utils.ApplictaionError
		giftCardTotals, err = s.r.NewTotalsService().SetContext(s.ctx).GetGiftCardTotals(giftCardable, []models.GiftCardTransaction{}, cart.Region, cart.GiftCards)
		if err != nil {
			return nil, err
		}
	}
	if order != nil {
		var err *utils.ApplictaionError
		giftCardTotals, err = s.r.NewTotalsService().SetContext(s.ctx).GetGiftCardTotals(giftCardable, order.GiftCardTransactions, order.Region, order.GiftCards)
		if err != nil {
			return nil, err
		}
	}

	return giftCardTotals, nil
}

func (s *TotalsService) GetDiscountTotal(cart *models.Cart, order *models.Order) (float64, *utils.ApplictaionError) {
	subtotal, err := s.GetSubtotal(cart, order, types.SubtotalOptions{})
	if err != nil {
		return 0, err
	}
	discountTotal := int(math.Round(s.GetLineItemAdjustmentsTotal(cart, order)))
	if subtotal < 0 {
		return s.Rounded(math.Max(float64(subtotal), float64(discountTotal))), nil
	}
	return s.Rounded(math.Min(float64(subtotal), float64(discountTotal))), nil
}

func (s *TotalsService) GetCalculationContext(cart *models.Cart, order *models.Order, options CalculationContextOptions) (*interfaces.TaxCalculationContext, *utils.ApplictaionError) {
	allocationMap, err := s.GetAllocationMap(cart, order, AllocationMapOptions{
		ExcludeGiftCards: options.ExcludeGiftCards,
		ExcludeDiscounts: options.ExcludeDiscounts,
	})
	if err != nil {
		return nil, err
	}
	var shippingMethods []models.ShippingMethod
	var res *interfaces.TaxCalculationContext
	if !options.ExcludeShipping {
		if cart != nil {
			shippingMethods = cart.ShippingMethods
		}
		if order != nil {
			shippingMethods = order.ShippingMethods
		}
	}
	if cart != nil {
		res = &interfaces.TaxCalculationContext{
			ShippingAddress: *cart.ShippingAddress,
			ShippingMethods: shippingMethods,
			Customer:        *cart.Customer,
			Region:          *cart.Region,
			IsReturn:        options.IsReturn,
			AllocationMap:   allocationMap,
		}
	}
	if order != nil {
		res = &interfaces.TaxCalculationContext{
			ShippingAddress: *order.ShippingAddress,
			ShippingMethods: shippingMethods,
			Customer:        *order.Customer,
			Region:          *order.Region,
			IsReturn:        options.IsReturn,
			AllocationMap:   allocationMap,
		}
	}
	return res, nil
}

func (s *TotalsService) Rounded(value float64) float64 {
	return math.Round(value)
}
