package services

import (
	"context"
	"math"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type NewTotalsService struct {
	ctx context.Context
	r   Registry
}

func NewNewTotalsServices(
	r Registry,
) *NewTotalsService {
	return &NewTotalsService{
		context.Background(),
		r,
	}
}

func (s *NewTotalsService) SetContext(context context.Context) *NewTotalsService {
	s.ctx = context
	return s
}

func (s *NewTotalsService) GetLineItemTotals(items []models.LineItem, includeTax bool, calculationContext *interfaces.TaxCalculationContext, taxRate *float64) (map[uuid.UUID]models.LineItem, *utils.ApplictaionError) {
	lineItemsTaxLinesMap := make(map[uuid.UUID][]models.LineItemTaxLine)
	if taxRate == nil && includeTax {
		itemContainsTaxLines := false
		for _, item := range items {
			if len(item.TaxLines) > 0 {
				itemContainsTaxLines = true
				lineItemsTaxLinesMap[item.Id] = item.TaxLines
			}
		}
		if !itemContainsTaxLines {
			lineItemsTaxLines, err := s.r.TaxProviderService().SetContext(s.ctx).GetTaxLinesMap(items, calculationContext)
			if err != nil {
				return nil, err
			}
			for _, item := range items {
				lineItemsTaxLinesMap[item.Id] = lineItemsTaxLines.LineItemsTaxLines[item.Id]
			}
		}
	}

	var calculationMethod func(
		item models.LineItem,
		taxRate *float64,
		includeTax bool,
		lineItemAllocation types.LineAllocations,
		taxLines []models.LineItemTaxLine,
		calculationContext *interfaces.TaxCalculationContext,
	) (*models.LineItem, *utils.ApplictaionError)
	if taxRate != nil {
		calculationMethod = s.getLineItemTotalsLegacy
	} else {
		calculationMethod = s.getLineItemTotals
	}

	itemsTotals := make(map[uuid.UUID]models.LineItem)
	for _, item := range items {
		lineItemAllocation := calculationContext.AllocationMap[item.Id]
		totals, err := calculationMethod(item, taxRate, includeTax, lineItemAllocation, lineItemsTaxLinesMap[item.Id], calculationContext)
		if err != nil {
			return nil, err
		}
		itemsTotals[item.Id] = *totals
	}

	return itemsTotals, nil
}

func (s *NewTotalsService) getLineItemTotals(item models.LineItem, taxRate *float64, includeTax bool, lineItemAllocation types.LineAllocations, taxLines []models.LineItemTaxLine, calculationContext *interfaces.TaxCalculationContext) (*models.LineItem, *utils.ApplictaionError) {
	subtotal := item.UnitPrice * float64(item.Quantity)
	feature := true
	if feature && item.IncludesTax {
		subtotal = 0
	}

	rawDiscountTotal := lineItemAllocation.Discount.Amount
	discountTotal := math.Round(rawDiscountTotal)

	totals := &models.LineItem{
		UnitPrice:        item.UnitPrice,
		Quantity:         item.Quantity,
		Subtotal:         subtotal,
		DiscountTotal:    discountTotal,
		Total:            subtotal - discountTotal,
		OriginalTotal:    subtotal,
		OriginalTaxTotal: 0,
		TaxTotal:         0,
		TaxLines:         item.TaxLines,
		RawDiscountTotal: rawDiscountTotal,
	}

	if includeTax {
		if len(totals.TaxLines) == 0 {
			totals.TaxLines = taxLines
		}
		if totals.TaxLines == nil && item.VariantId.UUID != uuid.Nil {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Tax Lines must be joined to calculate taxes",
				"500",
				nil,
			)
		}
	}

	if item.IsReturn {
		if item.TaxLines == nil && item.VariantId.UUID != uuid.Nil {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Return Line Items must join tax lines",
				"500",
				nil,
			)
		}
	}

	if len(totals.TaxLines) > 0 {
		taxTotal, err := s.r.TaxCalculationStrategy().Calculate([]models.LineItem{item}, totals.TaxLines, calculationContext)
		if err != nil {
			return nil, err
		}

		noDiscountContext := calculationContext
		noDiscountContext.AllocationMap = types.LineAllocationsMap{}
		originalTaxTotal, err := s.r.TaxCalculationStrategy().Calculate([]models.LineItem{item}, totals.TaxLines, noDiscountContext)
		if err != nil {
			return nil, err
		}

		feature = true
		if feature && item.IncludesTax {
			totals.Subtotal += totals.UnitPrice*float64(totals.Quantity) - originalTaxTotal
			totals.Total += totals.Subtotal
			totals.OriginalTotal += totals.Subtotal
		}

		totals.TaxTotal = taxTotal
		totals.OriginalTaxTotal = originalTaxTotal
		totals.Total += totals.TaxTotal
		totals.OriginalTotal += totals.OriginalTaxTotal
	}

	return totals, nil
}

func (s *NewTotalsService) getLineItemTotalsLegacy(item models.LineItem, taxRate *float64, includeTax bool, lineItemAllocation types.LineAllocations, taxLines []models.LineItemTaxLine, calculationContext *interfaces.TaxCalculationContext) (*models.LineItem, *utils.ApplictaionError) {
	subtotal := item.UnitPrice * float64(item.Quantity)
	feature := true
	if feature && item.IncludesTax {
		subtotal = 0
	}

	rawDiscountTotal := lineItemAllocation.Discount.Amount
	discountTotal := math.Round(rawDiscountTotal)

	totals := &models.LineItem{
		UnitPrice:        item.UnitPrice,
		Quantity:         item.Quantity,
		Subtotal:         subtotal,
		DiscountTotal:    discountTotal,
		Total:            subtotal - discountTotal,
		OriginalTotal:    subtotal,
		OriginalTaxTotal: 0,
		TaxTotal:         0,
		TaxLines:         nil,
		RawDiscountTotal: rawDiscountTotal,
	}

	if taxRate != nil {
		taxRate := *taxRate / 100
		includesTax := feature && item.IncludesTax
		taxIncludedInPrice := 0.0
		if !item.IncludesTax {
			taxIncludedInPrice = math.Round(utils.CalculatePriceTaxAmount(item.UnitPrice, taxRate, includesTax))
		}
		totals.Subtotal = math.Round((item.UnitPrice - taxIncludedInPrice) * float64(item.Quantity))
		totals.Total = totals.Subtotal
		totals.OriginalTaxTotal = math.Round(totals.Subtotal * taxRate)
		totals.TaxTotal = math.Round((totals.Subtotal - discountTotal) * taxRate)
		totals.Total += totals.TaxTotal
		if includesTax {
			totals.OriginalTotal += totals.Subtotal
		}
		totals.OriginalTotal += totals.OriginalTaxTotal
	}

	return totals, nil
}

func (s *NewTotalsService) GetLineItemRefund(lineItem models.LineItem, calculationContext *interfaces.TaxCalculationContext, taxRate *float64) (float64, *utils.ApplictaionError) {
	if taxRate != nil {
		return s.getLineItemRefundLegacy(lineItem, calculationContext, taxRate)
	}

	feature := true
	includesTax := feature && lineItem.IncludesTax
	discountAmount := calculationContext.AllocationMap[lineItem.Id].Discount.UnitAmount * float64(lineItem.Quantity)
	if lineItem.TaxLines == nil {
		return 0, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Cannot compute line item refund amount, tax lines are missing from the line item",
			"500",
			nil,
		)
	}

	totalTaxRate := 0.0
	for _, taxLine := range lineItem.TaxLines {
		totalTaxRate += taxLine.Rate / 100
	}

	taxAmountIncludedInPrice := 0.0
	if includesTax {
		taxAmountIncludedInPrice = math.Round(utils.CalculatePriceTaxAmount(lineItem.UnitPrice, totalTaxRate, includesTax))
	}

	lineSubtotal := (lineItem.UnitPrice-taxAmountIncludedInPrice)*float64(lineItem.Quantity) - discountAmount
	taxTotal := 0.0
	for _, taxLine := range lineItem.TaxLines {
		taxTotal += math.Round(lineSubtotal * (taxLine.Rate / 100))
	}

	return lineSubtotal + taxTotal, nil
}

func (s *NewTotalsService) getLineItemRefundLegacy(lineItem models.LineItem, calculationContext *interfaces.TaxCalculationContext, taxRate *float64) (float64, *utils.ApplictaionError) {
	feature := true
	includesTax := feature && lineItem.IncludesTax
	taxAmountIncludedInPrice := 0.0
	if includesTax {
		taxAmountIncludedInPrice = math.Round(utils.CalculatePriceTaxAmount(lineItem.UnitPrice, *taxRate/100, includesTax))
	}

	discountAmount := calculationContext.AllocationMap[lineItem.Id].Discount.Amount
	lineSubtotal := (lineItem.UnitPrice-taxAmountIncludedInPrice)*float64(lineItem.Quantity) - discountAmount
	return math.Round(lineSubtotal * (1 + *taxRate/100)), nil
}

func (s *NewTotalsService) GetGiftCardTotals(giftCardableAmount float64, giftCardTransactions []models.GiftCardTransaction, region *models.Region, giftCards []models.GiftCard) (*Total, *utils.ApplictaionError) {
	if len(giftCards) == 0 && len(giftCardTransactions) == 0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Cannot calculate the gift cart totals. Neither the gift cards or gift card transactions have been provided",
			"500",
			nil,
		)
	}

	if len(giftCardTransactions) > 0 {
		return s.GetGiftCardTransactionsTotals(giftCardTransactions, region), nil
	}

	var result *Total
	if len(giftCards) == 0 {
		return result, nil
	}

	totalGiftCardBalance := 0.0
	totalTaxFromGiftCards := 0.0
	giftCardableBalance := giftCardableAmount
	for _, giftCard := range giftCards {
		taxableAmount := 0.0
		totalGiftCardBalance += giftCard.Balance
		taxableAmount = math.Min(giftCardableBalance, giftCard.Balance)
		if taxableAmount <= 0 || giftCard.TaxRate == 0.0 {
			continue
		}
		taxAmountFromGiftCard := math.Round(taxableAmount * (giftCard.TaxRate / 100))
		totalTaxFromGiftCards += taxAmountFromGiftCard
		giftCardableBalance -= taxableAmount
	}

	result.TaxTotal = math.Round(totalTaxFromGiftCards)
	result.Total = math.Min(giftCardableAmount, totalGiftCardBalance)
	return result, nil
}

func (s *NewTotalsService) GetGiftCardTransactionsTotals(giftCardTransactions []models.GiftCardTransaction, region *models.Region) *Total {
	var result *Total
	for _, transaction := range giftCardTransactions {
		taxMultiplier := 0.0
		if transaction.TaxRate != 0.0 {
			taxMultiplier = transaction.TaxRate / 100
		} else if transaction.IsTaxable {
			if region.GiftCardsTaxable {
				taxMultiplier = transaction.GiftCard.TaxRate / 100
			}
		}
		result.Total += transaction.Amount
		result.TaxTotal += math.Round(transaction.Amount * taxMultiplier)
	}
	return result
}

func (s *NewTotalsService) GetShippingMethodTotals(shippingMethods []models.ShippingMethod, includeTax bool, discounts []models.Discount, taxRate *float64, calculationContext *interfaces.TaxCalculationContext) (map[uuid.UUID]models.ShippingMethod, *utils.ApplictaionError) {
	shippingMethodsTaxLinesMap := make(map[uuid.UUID][]models.ShippingMethodTaxLine)
	if taxRate == nil && includeTax {
		shippingMethodContainsTaxLines := false
		for _, method := range shippingMethods {
			if len(method.TaxLines) > 0 {
				shippingMethodContainsTaxLines = true
				shippingMethodsTaxLinesMap[method.Id] = method.TaxLines
			}
		}
		if !shippingMethodContainsTaxLines {
			calculationContext.ShippingMethods = shippingMethods
			shippingMethodsTaxLines, err := s.r.TaxProviderService().SetContext(s.ctx).GetTaxLinesMap([]models.LineItem{}, calculationContext)
			if err != nil {
				return nil, err
			}
			for _, method := range shippingMethods {
				shippingMethodsTaxLinesMap[method.Id] = shippingMethodsTaxLines.ShippingMethodsTaxLines[method.Id]
			}
		}
	}

	var calculationMethod func(shippingMethod models.ShippingMethod, includeTax bool, calculationContext *interfaces.TaxCalculationContext, taxLines []models.ShippingMethodTaxLine, discounts []models.Discount, taxRate float64) (*models.ShippingMethod, *utils.ApplictaionError)
	if taxRate != nil {
		calculationMethod = s.getShippingMethodTotalsLegacy
	} else {
		calculationMethod = s.getShippingMethodTotals
	}

	shippingMethodsTotals := make(map[uuid.UUID]models.ShippingMethod)
	for _, method := range shippingMethods {
		totals, err := calculationMethod(method, includeTax, calculationContext, shippingMethodsTaxLinesMap[method.Id], discounts, *taxRate)
		if err != nil {
			return nil, err
		}
		shippingMethodsTotals[method.Id] = *totals
	}

	return shippingMethodsTotals, nil
}

func (s *NewTotalsService) GetGiftCardableAmount(giftCardsTaxable bool, subtotal, shippingTotal, discountTotal, taxTotal float64) float64 {
	if giftCardsTaxable {
		return subtotal + shippingTotal - discountTotal
	}
	return subtotal + shippingTotal + taxTotal - discountTotal
}

func (s *NewTotalsService) getShippingMethodTotals(
	shippingMethod models.ShippingMethod,
	includeTax bool,
	calculationContext *interfaces.TaxCalculationContext,
	taxLines []models.ShippingMethodTaxLine,
	discounts []models.Discount,
	taxRate float64,
) (*models.ShippingMethod, *utils.ApplictaionError) {
	totals := &models.ShippingMethod{
		Price:            shippingMethod.Price,
		OriginalTotal:    shippingMethod.Price,
		Total:            shippingMethod.Price,
		Subtotal:         shippingMethod.Price,
		OriginalTaxTotal: 0,
		TaxTotal:         0,
		TaxLines:         shippingMethod.TaxLines,
	}

	if includeTax {
		if len(totals.TaxLines) == 0 {
			totals.TaxLines = taxLines
		}

		if len(totals.TaxLines) == 0 {
			return totals, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Tax Lines must be joined to calculate shipping taxes",
				"500",
				nil,
			)
		}
	}

	calculationContext.ShippingMethods = []models.ShippingMethod{shippingMethod}

	feature := false

	if len(totals.TaxLines) > 0 {
		var err *utils.ApplictaionError
		includesTax := feature && shippingMethod.IncludesTax
		totals.OriginalTaxTotal, err = s.r.TaxCalculationStrategy().Calculate([]models.LineItem{}, totals.TaxLines, calculationContext)
		if err != nil {
			return totals, err
		}
		totals.TaxTotal = totals.OriginalTaxTotal

		if includesTax {
			totals.Subtotal -= totals.TaxTotal
		} else {
			totals.OriginalTotal += totals.OriginalTaxTotal
			totals.Total += totals.TaxTotal
		}
	}

	hasFreeShipping := false
	for _, d := range discounts {
		if d.Rule.Type == models.DiscountRuleFreeShipping {
			hasFreeShipping = true
			break
		}
	}

	if hasFreeShipping {
		totals.Total = 0
		totals.Subtotal = 0
		totals.TaxTotal = 0
	}

	return totals, nil
}

func (s *NewTotalsService) getShippingMethodTotalsLegacy(
	shippingMethod models.ShippingMethod,
	includeTax bool,
	calculationContext *interfaces.TaxCalculationContext,
	taxLines []models.ShippingMethodTaxLine,
	discounts []models.Discount,
	taxRate float64,
) (*models.ShippingMethod, *utils.ApplictaionError) {
	totals := &models.ShippingMethod{
		Price:            shippingMethod.Price,
		OriginalTotal:    shippingMethod.Price,
		Total:            shippingMethod.Price,
		Subtotal:         shippingMethod.Price,
		OriginalTaxTotal: 0,
		TaxTotal:         0,
		TaxLines:         []models.ShippingMethodTaxLine{},
	}

	totals.OriginalTaxTotal = math.Round(float64(totals.Price) * (taxRate / 100))
	totals.TaxTotal = math.Round(float64(totals.Price) * (taxRate / 100))

	hasFreeShipping := false
	for _, d := range discounts {
		if d.Rule.Type == models.DiscountRuleFreeShipping {
			hasFreeShipping = true
			break
		}
	}

	if hasFreeShipping {
		totals.Total = 0
		totals.Subtotal = 0
		totals.TaxTotal = 0
	}

	return totals, nil
}
