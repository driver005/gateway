package services

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

type ClaimService struct {
	ctx context.Context
	r   Registry
}

func NewClaimService(
	r Registry,
) *ClaimService {
	return &ClaimService{
		context.Background(),
		r,
	}
}

func (s *ClaimService) SetContext(context context.Context) *ClaimService {
	s.ctx = context
	return s
}

func (s *ClaimService) Retrieve(id uuid.UUID, config sql.Options) (*models.ClaimOrder, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			"500",
			nil,
		)
	}
	var res *models.ClaimOrder
	query := sql.BuildQuery(models.ClaimOrder{Model: core.Model{Id: id}}, config)
	if err := s.r.ClaimRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ClaimService) List(selector models.ClaimOrder, config sql.Options) ([]models.ClaimOrder, *utils.ApplictaionError) {
	var res []models.ClaimOrder

	if reflect.DeepEqual(config, sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	query := sql.BuildQuery(selector, config)

	if err := s.r.ClaimRepository().Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ClaimService) Update(id uuid.UUID, data *models.ClaimOrder) (*models.ClaimOrder, *utils.ApplictaionError) {
	claim, err := s.Retrieve(id, sql.Options{Relations: []string{"shipping_methods"}})
	if err != nil {
		return nil, err
	}
	if claim.CanceledAt != nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Canceled claim cannot be updated",
			"500",
			nil,
		)
	}

	if data.ShippingMethods != nil {
		for _, m := range claim.ShippingMethods {
			_, err = s.r.ShippingOptionService().SetContext(s.ctx).UpdateShippingMethod(m.Id, &models.ShippingMethod{
				ClaimOrderId: uuid.NullUUID{},
			})
			if err != nil {
				return nil, err
			}
		}
		for _, method := range data.ShippingMethods {
			if method.Id != uuid.Nil {
				_, err = s.r.ShippingOptionService().SetContext(s.ctx).UpdateShippingMethod(method.Id, &models.ShippingMethod{
					ClaimOrderId: uuid.NullUUID{UUID: claim.Id},
				})
				if err != nil {
					return nil, err
				}
			} else {
				_, err = s.r.ShippingOptionService().SetContext(s.ctx).CreateShippingMethod(method.ShippingOptionId.UUID, method.Data, &models.ShippingMethod{
					ClaimOrderId: uuid.NullUUID{UUID: claim.Id},
					Price:        method.Price,
				})
				if err != nil {
					return nil, err
				}
			}
		}
	}

	if data.ClaimItems != nil {
		for _, i := range data.ClaimItems {
			if i.Id != uuid.Nil {
				_, err := s.r.ClaimItemService().SetContext(s.ctx).Update(i.Id, &i)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	data.Id = claim.Id

	if err := s.r.ClaimRepository().Update(s.ctx, data); err != nil {
		return nil, err
	}

	// err = s.eventBus_.Emit(ClaimServiceEvents.UPDATED, ClaimServiceEventsUpdatedData{
	// 	ID:             claim.ID,
	// 	NoNotification: claim.NoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	return claim, nil
}

func (s *ClaimService) ValidateCreateClaimInput(data *models.ClaimOrder) *utils.ApplictaionError {
	if data.Type != models.ClaimStatusRefund && data.Type != models.ClaimStatusReplace {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			`Claim type must be one of "refund" or "replace".`,
			"500",
			nil,
		)
	}

	if data.Type == models.ClaimStatusReplace && len(data.AdditionalItems) == 0 {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			`Claims with type "replace" must have at least one additional item.`,
			"500",
			nil,
		)
	}

	if len(data.ClaimItems) == 0 {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			`Claims must have at least one claim item.`,
			"500",
			nil,
		)
	}

	if data.RefundAmount != 0 && data.Type != models.ClaimStatusRefund {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf(`Claim has type %s but must be type "refund" to have a refund_amount.`, data.Type),
			"500",
			nil,
		)
	}

	var claimIds uuid.UUIDs
	for _, claimItem := range data.ClaimItems {
		claimIds = append(claimIds, claimItem.ItemId.UUID)
	}

	claimLineItems, err := s.r.LineItemService().SetContext(s.ctx).List(
		models.LineItem{},
		sql.Options{
			Relations:     []string{"order", "swap", "claim_order", "tax_lines"},
			Specification: []sql.Specification{sql.In("id", claimIds)},
		},
	)
	if err != nil {
		return err
	}

	for _, line := range claimLineItems {
		if line.Order != nil && line.Order.CanceledAt != nil ||
			line.Swap != nil && line.Swap.CanceledAt != nil ||
			line.ClaimOrder != nil && line.ClaimOrder.CanceledAt != nil {
			return utils.NewApplictaionError(
				utils.INVALID_DATA,
				`Cannot create a claim on a canceled item.`,
				"500",
				nil,
			)
		}
	}

	return nil
}

func (s *ClaimService) GetRefundTotalForClaimLinesOnOrder(order *models.Order, claimItems []models.ClaimItem) (*float64, *utils.ApplictaionError) {
	var claimLines []models.LineItem
	for _, ci := range claimItems {
		predicate := func(item models.LineItem) bool {
			return item.ShippedQuantity != 0 &&
				ci.Quantity <= item.ShippedQuantity &&
				item.Id == ci.ItemId.UUID
		}
		claimLine, ok := lo.Find(order.Items, predicate)
		if ok {
			claimLines = append(claimLines, models.LineItem{
				Model: core.Model{
					Id: claimLine.Id,
				},
				Quantity: ci.Quantity,
			})
			continue
		}
		if len(order.Swaps) > 0 {
			for _, swap := range order.Swaps {
				claimLine, ok := lo.Find(swap.AdditionalItems, predicate)
				if ok {
					claimLines = append(claimLines, models.LineItem{
						Model: core.Model{
							Id: claimLine.Id,
						},
						Quantity: ci.Quantity,
					})
					continue
				}
			}
		}
		if len(order.Claims) > 0 {
			for _, claim := range order.Claims {
				claimLine, ok := lo.Find(claim.AdditionalItems, predicate)
				if ok {
					claimLines = append(claimLines, models.LineItem{
						Model: core.Model{
							Id: claimLine.Id,
						},
						Quantity: ci.Quantity,
					})
					continue
				}
			}
		}
	}

	refundTotal := 0.0
	for _, line := range claimLines {
		refund, err := s.r.TotalsService().GetLineItemRefund(order, line)
		if err != nil {
			return nil, err
		}
		refundTotal += refund
	}

	return &refundTotal, nil
}

func (s *ClaimService) Create(data *models.ClaimOrder, locationId string, returnShipping *models.ShippingMethod) (*models.ClaimOrder, *utils.ApplictaionError) {
	err := s.ValidateCreateClaimInput(data)
	if err != nil {
		return nil, err
	}

	if data.ShippingAddress != nil {

		if err := s.r.AddressRepository().Save(s.ctx, data.ShippingAddress); err != nil {
			return nil, err
		}
		data.ShippingAddressId = uuid.NullUUID{UUID: data.ShippingAddress.Id}
	}

	if data.Type == models.ClaimStatusRefund && data.RefundAmount == 0 {
		toRefund, err := s.GetRefundTotalForClaimLinesOnOrder(data.Order, data.ClaimItems)
		if err != nil {
			return nil, err
		}
		data.RefundAmount = *toRefund
	}

	var newItems []models.LineItem
	if data.AdditionalItems != nil {
		for _, i := range data.AdditionalItems {
			newItem, err := s.r.LineItemService().SetContext(s.ctx).Generate(i.VariantId.UUID, nil, data.Order.RegionId.UUID, i.Quantity, types.GenerateLineItemContext{})
			if err != nil {
				return nil, err
			}
			newItems = append(newItems, newItem...)
		}

		for _, newItem := range data.AdditionalItems {
			if newItem.VariantId.UUID != uuid.Nil {
				_, err = s.r.ProductVariantInventoryService().SetContext(s.ctx).ReserveQuantity(newItem.VariantId.UUID, newItem.Quantity, ReserveQuantityContext{
					LineItemId:     newItem.Id,
					SalesChannelId: data.Order.SalesChannelId.UUID,
				})
				if err != nil {
					return nil, err
				}
			}
		}
	}

	evaluatedNoNotification := data.NoNotification
	if !evaluatedNoNotification {
		evaluatedNoNotification = data.Order.NoNotification
	}

	data.AdditionalItems = newItems
	data.PaymentStatus = models.ClaimPaymentStatusNotRefunded
	data.NoNotification = evaluatedNoNotification

	if err := s.r.ClaimRepository().Save(s.ctx, data); err != nil {
		return nil, err
	}

	var lineItemIds uuid.UUIDs
	for _, lineItem := range data.AdditionalItems {
		lineItemIds = append(lineItemIds, lineItem.Id)
	}

	if len(data.AdditionalItems) > 0 {
		calcContext, err := s.r.TotalsService().GetCalculationContext(nil, data.Order, CalculationContextOptions{})
		if err != nil {
			return nil, err
		}
		lineItems, err := s.r.LineItemService().SetContext(s.ctx).List(models.LineItem{}, sql.Options{
			Relations:     []string{"variant.product.profiles"},
			Specification: []sql.Specification{sql.In("id", lineItemIds)},
		})
		if err != nil {
			return nil, err
		}
		_, _, err = s.r.TaxProviderService().SetContext(s.ctx).CreateTaxLines(nil, lineItems, calcContext)
		if err != nil {
			return nil, err
		}
	}

	if data.ShippingMethods != nil {
		for _, method := range data.ShippingMethods {
			if method.Id != uuid.Nil {
				_, err = s.r.ShippingOptionService().SetContext(s.ctx).UpdateShippingMethod(method.Id, &models.ShippingMethod{
					ClaimOrderId: uuid.NullUUID{UUID: data.Id},
				})
				if err != nil {
					return nil, err
				}
			} else {
				_, err = s.r.ShippingOptionService().SetContext(s.ctx).CreateShippingMethod(method.ShippingOptionId.UUID, method.Data, &models.ShippingMethod{
					ClaimOrderId: uuid.NullUUID{UUID: data.Id},
					Price:        method.Price,
				})
				if err != nil {
					return nil, err
				}
			}
		}
	}

	var items []models.ReturnItem
	for _, ci := range data.ClaimItems {
		ci.ClaimOrderId = uuid.NullUUID{UUID: data.Id}
		_, err := s.r.ClaimItemService().SetContext(s.ctx).Create(&ci)
		if err != nil {
			return nil, err
		}

		items = append(items, models.ReturnItem{
			ItemId:   ci.ItemId,
			Quantity: ci.Quantity,
			Metadata: ci.Metadata,
		})
	}

	if returnShipping != nil {
		_, err = s.r.ReturnService().Create(&models.Return{
			RefundAmount:   data.RefundAmount,
			OrderId:        data.OrderId,
			ClaimOrderId:   uuid.NullUUID{UUID: data.Id},
			Items:          items,
			ShippingMethod: returnShipping,
			NoNotification: evaluatedNoNotification,
			LocationId:     locationId,
		})
		if err != nil {
			return nil, err
		}
	}

	// err = s.eventBus_.Emit(ClaimServiceEvents.CREATED, ClaimServiceEventsCreatedData{
	// 	ID:             result.ID,
	// 	NoNotification: result.NoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	return data, nil
}

func (s *ClaimService) CreateFulfillment(id uuid.UUID, noNotification bool, locationId string, metadata map[string]interface{}) (*models.ClaimOrder, *utils.ApplictaionError) {
	claim, err := s.Retrieve(id, sql.Options{
		Relations: []string{
			"additional_items.tax_lines",
			"additional_items.variant.product.profiles",
			"shipping_methods",
			"shipping_methods.shipping_option",
			"shipping_methods.tax_lines",
			"shipping_address",
			"order",
			"order.billing_address",
			"order.discounts",
			"order.discounts.rule",
			"order.payments",
		},
	})
	if err != nil {
		return nil, err
	}
	if claim.CanceledAt != nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Canceled claim cannot be fulfilled",
			"500",
			nil,
		)
	}

	if claim.FulfillmentStatus != models.ClaimFulfillmentStatusNotFulfilled && claim.FulfillmentStatus != models.ClaimFulfillmentStatusCanceled {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"The claim has already been fulfilled.",
			"500",
			nil,
		)
	}
	if claim.Type != models.ClaimStatusReplace {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			fmt.Sprintf(`Claims with the type %s can not be fulfilled.`, claim.Type),
			"500",
			nil,
		)
	}
	if len(claim.ShippingMethods) == 0 {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Cannot fulfill a claim without a shipping method.",
			"500",
			nil,
		)
	}
	evaluatedNoNotification := noNotification
	if !evaluatedNoNotification {
		evaluatedNoNotification = claim.NoNotification
	}

	claim.NoNotification = evaluatedNoNotification

	var lineItems []types.FulFillmentItemType
	for _, item := range claim.AdditionalItems {
		lineItems = append(lineItems, types.FulFillmentItemType{
			ItemId:   item.Id,
			Quantity: item.Quantity,
		})
	}

	fulfillmentOrder := &types.CreateFulfillmentOrder{}

	copier.CopyWithOption(&fulfillmentOrder, claim, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	fulfillmentOrder.Email = claim.Order.Email
	fulfillmentOrder.Discounts = claim.Order.Discounts
	fulfillmentOrder.CurrencyCode = claim.Order.CurrencyCode
	fulfillmentOrder.TaxRate = claim.Order.TaxRate
	fulfillmentOrder.RegionId = claim.Order.RegionId.UUID
	fulfillmentOrder.Region = claim.Order.Region
	fulfillmentOrder.DisplayId = claim.Order.DisplayId
	fulfillmentOrder.BillingAddress = claim.Order.BillingAddress
	fulfillmentOrder.Items = claim.AdditionalItems
	fulfillmentOrder.ShippingMethods = claim.ShippingMethods
	fulfillmentOrder.IsClaim = true
	fulfillmentOrder.NoNotification = evaluatedNoNotification

	fulfillments, err := s.r.FulfillmentService().SetContext(s.ctx).CreateFulfillment(fulfillmentOrder, lineItems, models.Fulfillment{
		Model: core.Model{
			Metadata: metadata,
		},
		ClaimOrderId: uuid.NullUUID{UUID: id},
		LocationId:   locationId,
	})
	if err != nil {
		return nil, err
	}

	var successfullyFulfilledItems []models.FulfillmentItem
	for _, fulfillment := range fulfillments {
		successfullyFulfilledItems = append(successfullyFulfilledItems, fulfillment.Items...)
	}

	claim.FulfillmentStatus = models.ClaimFulfillmentStatusFulfilled

	for _, item := range claim.AdditionalItems {
		fulfillmentItem, ok := lo.Find(successfullyFulfilledItems, func(i models.FulfillmentItem) bool {
			return i.ItemId.UUID == item.Id
		})
		if ok {
			fulfilledQuantity := item.FulfilledQuantity + fulfillmentItem.Quantity
			_, err = s.r.LineItemService().SetContext(s.ctx).Update(item.Id, nil, &models.LineItem{
				FulfilledQuantity: fulfilledQuantity,
			}, sql.Options{})
			if err != nil {
				return nil, err
			}
			if item.Quantity != fulfilledQuantity {
				claim.FulfillmentStatus = models.ClaimFulfillmentStatusRequiresAction
			}
		} else if item.Quantity != item.FulfilledQuantity {
			claim.FulfillmentStatus = models.ClaimFulfillmentStatusRequiresAction
		}
	}

	if err := s.r.ClaimRepository().Save(s.ctx, claim); err != nil {
		return nil, err
	}

	// eventsToEmit := fulfillments.Map(func(fulfillment Fulfillment) EventBusEvent {
	// 	return EventBusEvent{
	// 		EventName: ClaimServiceEventsFULFILLMENT_CREATED,
	// 		Data: ClaimServiceEventsFulfillmentCreatedData{
	// 			ID:             id,
	// 			FulfillmentID:  fulfillment.ID,
	// 			NoNotification: claim.NoNotification,
	// 		},
	// 	}
	// })
	// err = s.eventBus_.Emit(eventsToEmit)
	// if err != nil {
	// 	return nil, err
	// }

	return claim, nil
}

func (s *ClaimService) CancelFulfillment(fulfillmentId uuid.UUID) (*models.ClaimOrder, *utils.ApplictaionError) {
	canceled, err := s.r.FulfillmentService().SetContext(s.ctx).CancelFulfillment(fulfillmentId, nil)
	if err != nil {
		return nil, err
	}
	if canceled.ClaimOrderId.UUID == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			`Fufillment not related to a claim`,
			"500",
			nil,
		)
	}
	claim, err := s.Retrieve(canceled.ClaimOrderId.UUID, sql.Options{})
	if err != nil {
		return nil, err
	}
	claim.FulfillmentStatus = models.ClaimFulfillmentStatusCanceled

	if err := s.r.ClaimRepository().Save(s.ctx, claim); err != nil {
		return nil, err
	}

	return claim, nil
}

func (s *ClaimService) ProcessRefund(id uuid.UUID) (*models.ClaimOrder, *utils.ApplictaionError) {
	claim, err := s.Retrieve(id, sql.Options{
		Relations: []string{"order", "order.payments"},
	})
	if err != nil {
		return nil, err
	}
	if claim.CanceledAt != nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Canceled claim cannot be processed",
			"500",
			nil,
		)
	}
	if claim.Type != models.ClaimStatusRefund {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			`Claim must have type "refund" to create a refund.`,
			"500",
			nil,
		)
	}
	if claim.RefundAmount != 0 {
		_, err = s.r.PaymentProviderService().SetContext(s.ctx).RefundPayments(claim.Order.Payments, claim.RefundAmount, "claim", nil)
		if err != nil {
			return nil, err
		}
	}
	claim.PaymentStatus = models.ClaimPaymentStatusRefunded

	if err := s.r.ClaimRepository().Save(s.ctx, claim); err != nil {
		return nil, err
	}

	// err = s.eventBus_.Emit(ClaimServiceEventsREFUND_PROCESSED, ClaimServiceEventsRefundProcessedData{
	// 	ID:             id,
	// 	NoNotification: claimOrder.NoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	return claim, nil
}

func (s *ClaimService) CreateShipment(id uuid.UUID, fulfillmentId uuid.UUID, trackingLinks []models.TrackingLink, noNotification bool, metadata map[string]interface{}) (*models.ClaimOrder, *utils.ApplictaionError) {
	claim, err := s.Retrieve(id, sql.Options{
		Relations: []string{"additional_items"},
	})
	if err != nil {
		return nil, err
	}

	if claim.CanceledAt != nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Canceled claim cannot be fulfilled as shipped",
			"500",
			nil,
		)
	}

	evaluatedNoNotification := noNotification
	if !noNotification {
		evaluatedNoNotification = claim.NoNotification
	}

	shipment, err := s.r.FulfillmentService().SetContext(s.ctx).CreateShipment(fulfillmentId, trackingLinks, &models.Fulfillment{
		Model:          core.Model{Metadata: metadata},
		NoNotification: evaluatedNoNotification,
	})
	if err != nil {
		return nil, err
	}

	claim.FulfillmentStatus = models.ClaimFulfillmentStatusShipped

	for _, additionalItem := range claim.AdditionalItems {
		shipped, ok := lo.Find(shipment.Items, func(item models.FulfillmentItem) bool {
			return item.ItemId.UUID == additionalItem.Id
		})
		if ok {
			shippedQty := additionalItem.ShippedQuantity + shipped.Quantity
			_, err := s.r.LineItemService().SetContext(s.ctx).Update(additionalItem.Id, nil, &models.LineItem{
				ShippedQuantity: shippedQty,
			}, sql.Options{})
			if err != nil {
				return nil, err
			}
			if shippedQty != additionalItem.Quantity {
				claim.FulfillmentStatus = models.ClaimFulfillmentStatusPartiallyShipped
			}
		} else if additionalItem.ShippedQuantity != additionalItem.Quantity {
			claim.FulfillmentStatus = models.ClaimFulfillmentStatusPartiallyShipped
		}
	}

	if err := s.r.ClaimRepository().Save(s.ctx, claim); err != nil {
		return nil, err
	}

	// err = eventBus.emit(ClaimServiceEvents.SHIPMENT_CREATED, map[string]interface{}{
	// 	"id":              id,
	// 	"fulfillment_id":  shipment.ID,
	// 	"no_notification": evaluatedNoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	return claim, nil
}

func (s *ClaimService) Cancel(id uuid.UUID) (*models.ClaimOrder, *utils.ApplictaionError) {
	claim, err := s.Retrieve(id, sql.Options{
		Relations: []string{"return_order", "fulfillments", "order", "order.refunds"},
	})
	if err != nil {
		return nil, err
	}

	if claim.RefundAmount != 0 {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Claim with a refund cannot be canceled",
			"500",
			nil,
		)
	}

	if claim.Fulfillments != nil {
		for _, f := range claim.Fulfillments {
			if f.CanceledAt == nil {
				return nil, utils.NewApplictaionError(
					utils.NOT_ALLOWED,
					"All fulfillments must be canceled before the claim can be canceled",
					"500",
					nil,
				)
			}
		}
	}

	if claim.ReturnOrder != nil && claim.ReturnOrder.Status != models.ReturnCanceled {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Return must be canceled before the claim can be canceled",
			"500",
			nil,
		)
	}

	now := time.Now()
	claim.CanceledAt = &now
	claim.FulfillmentStatus = models.ClaimFulfillmentStatusCanceled

	if err := s.r.ClaimRepository().Save(s.ctx, claim); err != nil {
		return nil, err
	}

	// err = eventBus.emit(ClaimServiceEvents.CANCELED, map[string]interface{}{
	// 	"id":              claimOrder.ID,
	// 	"no_notification": claimOrder.NoNotification,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	return claim, nil
}
