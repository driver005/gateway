package services

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

// TODO:OrdersReturnItem
type Transformer func(item *models.LineItem, quantity int, additional *types.OrderReturnItem) (*models.ReturnItem, *utils.ApplictaionError)

type ReturnService struct {
	ctx context.Context
	r   Registry
}

func NewReturnService(
	r Registry,
) *ReturnService {
	return &ReturnService{
		context.Background(),
		r,
	}
}

func (s *ReturnService) SetContext(context context.Context) *ReturnService {
	s.ctx = context
	return s
}

func (s *ReturnService) GetFulfillmentItems(order *models.Order, items []types.OrderReturnItem, transformer Transformer) ([]models.ReturnItem, *utils.ApplictaionError) {
	merged := append([]models.LineItem{}, order.Items...)
	if order.Swaps != nil && len(order.Swaps) > 0 {
		for _, s := range order.Swaps {
			merged = append(merged, s.AdditionalItems...)
		}
	}
	if order.Claims != nil && len(order.Claims) > 0 {
		for _, c := range order.Claims {
			merged = append(merged, c.AdditionalItems...)
		}
	}
	toReturn := make([]models.ReturnItem, len(items))
	for i, data := range items {
		var item *models.LineItem = &models.LineItem{}
		for _, it := range merged {
			if it.Id == data.ItemId {
				item = &it
				break
			}
		}
		if reflect.DeepEqual(item, &models.LineItem{}) {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Item not found",
				"500",
				nil,
			)
		}
		line, err := transformer(item, data.Quantity, &data)
		if err != nil {
			return nil, err
		}
		toReturn[i] = *line
	}
	filtered := make([]models.ReturnItem, 0)
	for _, i := range toReturn {
		if !reflect.DeepEqual(i, models.ReturnItem{}) {
			filtered = append(filtered, i)
		}
	}
	return filtered, nil
}

func (s *ReturnService) List(selector *types.FilterableReturn, config *sql.Options) ([]models.Return, *utils.ApplictaionError) {
	returns, _, err := s.ListAndCount(selector, config)
	if err != nil {
		return nil, err
	}
	return returns, nil
}

func (s *ReturnService) ListAndCount(selector *types.FilterableReturn, config *sql.Options) ([]models.Return, *int64, *utils.ApplictaionError) {
	var res []models.Return

	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = 0
		config.Take = 20
	}

	query := sql.BuildQuery(selector, config)
	count, err := s.r.ReturnRepository().FindAndCount(s.ctx, &res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *ReturnService) Cancel(returnId uuid.UUID) (*models.Return, *utils.ApplictaionError) {
	res, err := s.Retrieve(returnId, &sql.Options{})
	if err != nil {
		return nil, err
	}
	if res.Status == models.ReturnReceived {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Can't cancel a return which has been returned",
			nil,
		)
	}

	res.Status = models.ReturnCanceled

	if err := s.r.ReturnRepository().Save(s.ctx, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ReturnService) ValidateReturnStatuses(order *models.Order) *utils.ApplictaionError {
	if order.FulfillmentStatus == models.FulfillmentStatusNotFulfilled || order.FulfillmentStatus == models.FulfillmentStatusReturned {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Can't return an unfulfilled or already returned order",
			nil,
		)
	}
	if order.PaymentStatus != models.PaymentStatusCaptured {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Can't return an order with payment unprocessed",
			nil,
		)
	}
	return nil
}

func (s *ReturnService) ValidateReturnLineItem(item *models.LineItem, quantity int, additional *types.OrderReturnItem) (*models.ReturnItem, *utils.ApplictaionError) {
	if item == nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Return contains invalid line item",
			nil,
		)
	}
	returnable := item.Quantity - item.ReturnedQuantity
	if quantity > returnable {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Cannot return more items than have been purchased",
			nil,
		)
	}
	toReturn := &models.ReturnItem{
		Item:     item,
		Quantity: quantity,
	}
	if additional.ReasonId != uuid.Nil {
		toReturn.ReasonId = uuid.NullUUID{UUID: additional.ReasonId}
	}
	if additional.Note != "" {
		toReturn.Note = additional.Note
	}
	return toReturn, nil
}

func (s *ReturnService) Retrieve(returnId uuid.UUID, config *sql.Options) (*models.Return, *utils.ApplictaionError) {
	if returnId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"returnId" must be defined`,
			nil,
		)
	}

	var res *models.Return = &models.Return{}

	query := sql.BuildQuery(models.Return{BaseModel: core.BaseModel{Id: returnId}}, config)
	if err := s.r.ReturnRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ReturnService) RetrieveBySwap(swapId uuid.UUID, relations []string) (*models.Return, *utils.ApplictaionError) {
	if swapId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"swapId" must be defined`,
			nil,
		)
	}

	var res *models.Return = &models.Return{}

	query := sql.BuildQuery(models.Return{SwapId: uuid.NullUUID{UUID: swapId}}, &sql.Options{
		Relations: relations,
	})
	if err := s.r.ReturnRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ReturnService) Update(returnId uuid.UUID, data *types.UpdateReturnInput) (*models.Return, *utils.ApplictaionError) {
	ret, err := s.Retrieve(returnId, &sql.Options{})
	if err != nil {
		return nil, err
	}
	if ret.Status == models.ReturnCanceled {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Cannot Update a canceled return",
			nil,
		)
	}

	if !reflect.ValueOf(data.Items).IsZero() {
		var returns []models.ReturnItem
		for _, r := range data.Items {
			returns = append(returns, models.ReturnItem{
				ItemId:   uuid.NullUUID{UUID: r.ItemId},
				Quantity: r.Quantity,
				ReasonId: uuid.NullUUID{UUID: r.ReasonId},
				Note:     r.Note,
			})
		}

		ret.Items = returns
	}
	if data.ShippingMethod != nil {
		ret.ShippingMethod = &models.ShippingMethod{
			ShippingOptionId: uuid.NullUUID{UUID: data.ShippingMethod.OptionId},
			Price:            data.ShippingMethod.Price,
		}
	}
	if !reflect.ValueOf(data.NoNotification).IsZero() {
		ret.NoNotification = data.NoNotification
	}
	if data.Metadata != nil {
		ret.Metadata = utils.MergeMaps(ret.Metadata, data.Metadata)
	}

	if err := s.r.ReturnRepository().Update(s.ctx, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *ReturnService) Create(data *types.CreateReturnInput) (*models.Return, *utils.ApplictaionError) {
	orderId := data.OrderId
	if data.SwapId == uuid.Nil {
		data.OrderId = uuid.Nil
	}

	if data.Items != nil {
		for _, item := range data.Items {
			line, err := s.r.LineItemService().SetContext(s.ctx).Retrieve(item.ItemId, &sql.Options{Relations: []string{"order", "swap", "claim_order"}})
			if err != nil {
				return nil, err
			}
			if line.Order.CanceledAt != nil || line.Swap.CanceledAt != nil || line.ClaimOrder.CanceledAt != nil {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					"Cannot Create a return for a canceled item.",
					"500",
					nil,
				)
			}
		}
	}
	order, err := s.r.OrderService().SetContext(s.ctx).RetrieveById(orderId, &sql.Options{Selects: []string{"refunded_total", "total", "refundable_amount"}, Relations: []string{"swaps", "swaps.additional_items", "swaps.additional_items.tax_lines", "claims", "claims.additional_items", "claims.additional_items.tax_lines", "items", "items.tax_lines", "region", "region.tax_rates"}})
	if err != nil {
		return nil, err
	}
	returnLines, err := s.GetFulfillmentItems(order, data.Items, s.ValidateReturnLineItem)
	if err != nil {
		return nil, err
	}

	var lineItem []models.LineItem
	var ids uuid.UUIDs
	for _, item := range returnLines {
		lineItem = append(lineItem, *item.Item)
		ids = append(ids, item.ReturnId.UUID)
	}

	toRefund := 0.0
	if !reflect.ValueOf(data.RefundAmount).IsZero() {
		toRefund = data.RefundAmount
		if toRefund == 0.0 {
			refund, err := s.r.TotalsService().GetRefundTotal(order, lineItem)
			if err != nil {
				return nil, err
			}
			toRefund = refund
		}
	}

	returnObject := &models.Return{
		OrderId:      uuid.NullUUID{UUID: orderId},
		Status:       models.ReturnRequested,
		RefundAmount: math.Floor(toRefund),
		Items:        make([]models.ReturnItem, len(returnLines)),
	}

	returnReasons, err := s.r.ReturnReasonService().SetContext(s.ctx).List(&types.FilterableReturnReason{}, &sql.Options{
		Relations:     []string{"return_reason_children"},
		Specification: []sql.Specification{sql.In("id", ids)},
	})
	if err != nil {
		return nil, err
	}
	for _, rr := range returnReasons {
		if len(rr.ReturnReasonChildren) > 0 {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Cannot apply return reason category",
				"500",
				nil,
			)
		}
	}

	for _, l := range returnLines {
		returnObject.Items = append(returnObject.Items, models.ReturnItem{
			ItemId:            uuid.NullUUID{UUID: l.Item.Id},
			Quantity:          l.Quantity,
			RequestedQuantity: l.Quantity,
			ReasonId:          l.ReasonId,
			Note:              l.Note,
			Metadata:          l.Metadata,
		})
	}

	if err := s.r.ReturnRepository().Save(s.ctx, returnObject); err != nil {
		return nil, err
	}

	if data.ShippingMethod != nil && data.ShippingMethod.OptionId != uuid.Nil {
		shippingMethod, err := s.r.ShippingOptionService().SetContext(s.ctx).CreateShippingMethod(data.ShippingMethod.OptionId, map[string]interface{}{}, &types.CreateShippingMethodDto{CreateShippingMethod: types.CreateShippingMethod{Price: data.ShippingMethod.Price, ReturnId: returnObject.Id}})
		if err != nil {
			return nil, err
		}
		calculationContext, err := s.r.TotalsService().GetCalculationContext(types.CalculationContextData{
			Discounts:       order.Discounts,
			Items:           order.Items,
			Customer:        order.Customer,
			Region:          order.Region,
			ShippingAddress: order.ShippingAddress,
			Swaps:           order.Swaps,
			Claims:          order.Claims,
			ShippingMethods: order.ShippingMethods,
		}, CalculationContextOptions{})
		if err != nil {
			return nil, err
		}
		taxLines, err := s.r.TaxProviderService().SetContext(s.ctx).CreateShippingTaxLines(shippingMethod, calculationContext)
		if err != nil {
			return nil, err
		}
		includesTax := true && shippingMethod.IncludesTax
		taxRate := 0.0
		for _, tl := range taxLines {
			taxRate += tl.Rate / 100
		}
		taxAmountIncludedInPrice := 0.0
		if !includesTax {
			taxAmountIncludedInPrice = 0.0
		} else {
			taxAmountIncludedInPrice = math.Round(utils.CalculatePriceTaxAmount(
				shippingMethod.Price,
				taxRate,
				includesTax,
			))
		}
		shippingPriceWithoutTax := shippingMethod.Price - taxAmountIncludedInPrice
		shippingTotal := shippingPriceWithoutTax
		for _, tl := range taxLines {
			shippingTotal += math.Round(float64(shippingPriceWithoutTax) * (tl.Rate / 100))
		}
		if !reflect.ValueOf(data.RefundAmount).IsZero() {
			returnObject.RefundAmount = math.Floor(toRefund) - shippingTotal
			if err := s.r.ReturnRepository().Save(s.ctx, returnObject); err != nil {
				return nil, err
			}
		}
	}
	return returnObject, nil
}

func (s *ReturnService) Fulfill(returnId uuid.UUID) (*models.Return, *utils.ApplictaionError) {
	returnData, err := s.Retrieve(returnId, &sql.Options{
		Relations: []string{"items", "shipping_method", "shipping_method.tax_lines", "shipping_method.shipping_option", "swap", "claim_order"},
	})
	if err != nil {
		return nil, err
	}
	if returnData.Status == models.ReturnCanceled {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Cannot fulfill a canceled return",
			nil,
		)
	}

	items, err := s.r.LineItemService().SetContext(s.ctx).List(models.LineItem{}, &sql.Options{Relations: []string{"tax_lines", "variant.product.profiles"}})
	if err != nil {
		return nil, err
	}
	for i, item := range returnData.Items {
		var found *models.LineItem = &models.LineItem{}
		for _, it := range items {
			if it.Id == item.ItemId.UUID {
				found = &it
				break
			}
		}
		if reflect.DeepEqual(found, &models.LineItem{}) {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Item not found",
				"500",
				nil,
			)
		}
		returnData.Items[i] = models.ReturnItem{
			Item: found,
		}
	}
	if returnData.ShippingData != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Return has already been fulfilled",
			nil,
		)
	}
	if returnData.ShippingMethod == nil {
		return returnData, nil
	}
	returnData.ShippingData, err = s.r.FulfillmentProviderService().SetContext(s.ctx).CreateReturn(returnData)
	if err != nil {
		return nil, err
	}
	if err := s.r.ReturnRepository().Save(s.ctx, returnData); err != nil {
		return nil, err
	}
	return returnData, nil
}

func (s *ReturnService) Receive(returnId uuid.UUID, receivedItems []types.OrderReturnItem, refundAmount *float64, allowMismatch bool, context map[string]interface{}) (*models.Return, *utils.ApplictaionError) {
	returnObj, err := s.Retrieve(returnId, &sql.Options{Relations: []string{"items", "swap", "swap.additional_items"}})
	if err != nil {
		return nil, err
	}
	if returnObj.Status == models.ReturnCanceled {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Cannot receive a canceled return",
			nil,
		)
	}
	orderId := returnObj.OrderId
	if returnObj.Swap != nil {
		orderId = returnObj.Swap.OrderId
	}
	order, err := s.r.OrderService().SetContext(s.ctx).RetrieveById(orderId.UUID, &sql.Options{Relations: []string{"items", "returns", "payments", "discounts", "discounts.rule", "refunds", "shipping_methods", "shipping_methods.shipping_option", "region", "swaps", "swaps.additional_items", "claims", "claims.additional_items"}})
	if err != nil {
		return nil, err
	}
	if returnObj.Status == models.ReturnReceived {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("Return with id %s has already been received", returnId),
			nil,
		)
	}
	returnLines, err := s.GetFulfillmentItems(order, receivedItems, s.ValidateReturnLineItem)
	if err != nil {
		return nil, err
	}
	var newLines []models.ReturnItem
	returnStatus := models.ReturnReceived
	isMatching := true
	for i, l := range returnLines {
		existing, ok := lo.Find(returnObj.Items, func(m models.ReturnItem) bool {
			return m.ItemId.UUID == l.ItemId.UUID
		})
		if ok {
			existing.IsRequested = l.Quantity == existing.Quantity
			existing.Quantity = l.Quantity
			existing.RequestedQuantity = l.Quantity
			existing.RecievedQuantity = l.Quantity
			newLines = append(newLines, existing)
		} else {
			newLines = append(newLines, models.ReturnItem{
				Metadata:         l.Metadata,
				ReturnId:         uuid.NullUUID{UUID: returnObj.Id},
				ItemId:           l.ItemId,
				Quantity:         l.Quantity,
				RecievedQuantity: l.Quantity,
				IsRequested:      false,
			})
		}
		if !newLines[i].IsRequested {
			isMatching = false
		}
	}
	if !isMatching && !allowMismatch {
		returnStatus = models.ReturnRequiresAction
	}
	totalRefundableAmount := refundAmount
	if totalRefundableAmount == nil {
		totalRefundableAmount = &returnObj.RefundAmount
	}
	now := time.Now()
	updateObj := returnObj
	updateObj.LocationId = uuid.NullUUID{UUID: context["locationId"].(uuid.UUID)}
	updateObj.Status = returnStatus
	updateObj.Items = newLines
	updateObj.RefundAmount = math.Floor(*totalRefundableAmount)
	updateObj.ReceivedAt = &now

	if err := s.r.ReturnRepository().Save(s.ctx, updateObj); err != nil {
		return nil, err
	}

	for _, i := range returnObj.Items {
		lineItem, err := s.r.LineItemService().SetContext(s.ctx).Retrieve(i.ItemId.UUID, &sql.Options{})
		if err != nil {
			return nil, err
		}
		returnedQuantity := lineItem.ReturnedQuantity + i.Quantity
		_, err = s.r.LineItemService().SetContext(s.ctx).Update(i.ItemId.UUID, nil, &models.LineItem{ReturnedQuantity: returnedQuantity}, &sql.Options{})
		if err != nil {
			return nil, err
		}
	}

	for _, line := range newLines {
		var orderItem *models.LineItem = &models.LineItem{}
		for _, item := range order.Items {
			if item.Id == line.ItemId.UUID {
				orderItem = &item
				break
			}
		}
		if orderItem != nil && orderItem.VariantId.UUID != uuid.Nil {
			err = s.r.ProductVariantInventoryService().SetContext(s.ctx).AdjustInventory(orderItem.VariantId.UUID, updateObj.LocationId.UUID, line.RecievedQuantity)
			if err != nil {
				return nil, err
			}
		}
	}
	return updateObj, nil
}
