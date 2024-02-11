package admin

import (
	"fmt"
	"reflect"

	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

const includes bool = true

type Order struct {
	r Registry
}

func NewOrder(r Registry) *Order {
	m := Order{r: r}
	return &m
}

func (m *Order) SetRoutes(router fiber.Router) {
	route := router.Group("/orders")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("/:id", m.Update)

	route.Post("/:id/complete", m.Complete)
	route.Post("/:id/cancel", m.Cancel)
	route.Post("/:id/archive", m.Archive)
	route.Post("/:id/refund", m.RefundPayment)
	route.Post("/:id/capture", m.CapturePayment)
	route.Post("/:id/shipment", m.CreateShipment)
	route.Post("/:id/return", m.RequestReturn)
	route.Post("/:id/shipping-methods", m.AddShippingMethod)

	route.Post("/:id/swaps", m.CreateSwap)
	route.Post("/:id/swaps/:swap_id/cancel", m.CancelSwap)
	route.Post("/:id/swaps/:swap_id/fulfillments", m.FulfillSwap)
	route.Post("/:id/swaps/:swap_id/shipments", m.CreateSwapShipment)
	route.Post("/:id/swaps/:swap_id/process-payment", m.ProcessSwapPayment)

	route.Post("/:id/claims", m.CreateClaim)
	route.Post("/:id/claims/:claim_id/cancel", m.CancelClaim)
	route.Post("/:id/claims/:claim_id", m.UpdateClaim)
	route.Post("/:id/claims/:claim_id/fulfillments", m.FulfillClaim)
	route.Post("/:id/claims/:claim_id/shipments", m.CreateClaimShippment)

	route.Post("/:id/reservations", m.GetReservations)
	route.Post("/:id/line-items/:line_item_id/reserve", m.CreateReservationForLineItem)

	route.Post("/:id/fulfillment", m.CreateFulfillment)
	route.Post("/:id/fulfillments/:fulfillment_id/cancel", m.CancelFullfillment)
	route.Post("/:id/swaps/:swap_id/fulfillments/:fulfillment_id/cancel", m.CancelFullfillmentSwap)
	route.Post("/:id/claims/:claim_id/fulfillments/:fulfillment_id/cancel", m.CancelFullfillmentClaim)
}

func (m *Order) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableOrder](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.OrderService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
}

func (m *Order) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateOrderInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) AddShippingMethod(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderShippingMethod](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).AddShippingMethod(id, model.OptionId, model.Data, &types.CreateShippingMethodDto{
		CreateShippingMethod: types.CreateShippingMethod{
			Price: model.Price,
		},
	}); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) Archive(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).Archive(id); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) Cancel(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).Cancel(id); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)

}

func (m *Order) CancelSwap(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	swapId, err := api.BindDelete(context, "swap_id")
	if err != nil {
		return err
	}

	swap, err := m.r.SwapService().SetContext(context.Context()).Retrieve(swapId, &sql.Options{})
	if err != nil {
		return err
	}

	if swap.OrderId.UUID != id {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("no swap was found with the id: %s related to order: %s", swapId, id),
		)
	}

	if _, err := m.r.SwapService().SetContext(context.Context()).Cancel(swapId); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) CancelClaim(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	claimId, err := api.BindDelete(context, "claim_id")
	if err != nil {
		return err
	}

	claim, err := m.r.ClaimService().SetContext(context.Context()).Retrieve(claimId, &sql.Options{})
	if err != nil {
		return err
	}

	if claim.OrderId.UUID != id {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("no claim was found with the id: %s related to order: %s", claimId, id),
		)
	}

	if _, err := m.r.ClaimService().SetContext(context.Context()).Cancel(claimId); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) CancelFullfillmentClaim(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	claimId, err := api.BindDelete(context, "claim_id")
	if err != nil {
		return err
	}

	fulfillmentId, err := api.BindDelete(context, "fulfillment_id")
	if err != nil {
		return err
	}

	fulfillment, err := m.r.FulfillmentService().SetContext(context.Context()).Retrieve(fulfillmentId, &sql.Options{})
	if err != nil {
		return err
	}

	if fulfillment.ClaimOrderId.UUID != id {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("no fulfillment was found with the id: %s related to claim: %s", fulfillmentId, claimId),
		)
	}

	claim, err := m.r.ClaimService().SetContext(context.Context()).Retrieve(claimId, &sql.Options{})
	if err != nil {
		return err
	}

	if claim.OrderId.UUID != id {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("no claim was found with the id: %s related to order: %s", claimId, id),
		)
	}

	if _, err := m.r.ClaimService().SetContext(context.Context()).CancelFulfillment(fulfillmentId); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) CancelFullfillmentSwap(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	swapId, err := api.BindDelete(context, "swap_id")
	if err != nil {
		return err
	}

	fulfillmentId, err := api.BindDelete(context, "fulfillment_id")
	if err != nil {
		return err
	}

	fulfillment, err := m.r.FulfillmentService().SetContext(context.Context()).Retrieve(fulfillmentId, &sql.Options{})
	if err != nil {
		return err
	}

	if fulfillment.SwapId.UUID != swapId {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("no fulfillment was found with the id: %s related to swap: %s", fulfillmentId, swapId),
		)
	}

	swap, err := m.r.SwapService().SetContext(context.Context()).Retrieve(swapId, &sql.Options{})
	if err != nil {
		return err
	}

	if swap.OrderId.UUID != id {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("no swap was found with the id: %s related to order: %s", swapId, id),
		)
	}

	if _, err := m.r.SwapService().SetContext(context.Context()).CancelFulfillment(fulfillmentId); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) CancelFullfillment(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	fulfillmentId, err := api.BindDelete(context, "fulfillment_id")
	if err != nil {
		return err
	}

	fulfillment, err := m.r.FulfillmentService().SetContext(context.Context()).Retrieve(fulfillmentId, &sql.Options{})
	if err != nil {
		return err
	}

	if fulfillment.OrderId.UUID != id {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("no fulfillment was found with the id: %s related to order: %s", fulfillmentId, id),
		)
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).CancelFulfillment(fulfillmentId); err != nil {
		return err
	}

	data, err := m.r.FulfillmentService().SetContext(context.Context()).Retrieve(fulfillmentId, &sql.Options{Relations: []string{"items", "items.item"}})
	if err != nil {
		return err
	}

	if data.LocationId != uuid.Nil && m.r.InventoryService() != nil {
		for _, item := range data.Items {
			if item.Item.VariantId.UUID != uuid.Nil {
				if err := m.r.ProductVariantInventoryService().SetContext(context.Context()).AdjustInventory(item.Item.VariantId.UUID, item.Fulfillment.LocationId, item.Quantity); err != nil {
					return err
				}
			}
		}
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) CapturePayment(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).CapturePayment(id); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) Complete(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).CompleteOrder(id); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) CreateClaimShippment(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderClaimShipments](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	claimId, err := api.BindDelete(context, "claim_id")
	if err != nil {
		return err
	}

	trakingLinks := []models.TrackingLink{}
	for _, link := range model.TrackingNumbers {
		trakingLinks = append(trakingLinks, models.TrackingLink{TrackingNumber: link})
	}

	if _, err := m.r.ClaimService().SetContext(context.Context()).CreateShipment(claimId, model.FulfillmentId, trakingLinks, false, nil); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) CreateClaim(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.CreateClaimInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	headerKey := context.Get("Idempotency-Key")

	idempotencyKey := &models.IdempotencyKey{}
	idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).InitializeRequest(headerKey, context.Method(), nil, context.Path())
	if err != nil {
		return err
	}

	context.Set("Access-Control-Expose-Headers", "Idempotency-Key")
	context.Set("Idempotency-Key", idempotencyKey.IdempotencyKey)

	if idempotencyKey.RecoveryPoint == "started" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			order, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(id, &sql.Options{
				Relations: []string{
					"customer",
					"shipping_address",
					"region",
					"items",
					"items.tax_lines",
					"discounts",
					"discounts.rule",
					"claims",
					"claims.additional_items",
					"claims.additional_items.tax_lines",
					"swaps",
					"swaps.additional_items",
					"swaps.additional_items.tax_lines",
				},
			})
			if err != nil {
				return nil, err
			}

			model.Order = order
			// model.IdempotencyKey = idempotencyKey.IdempotencyKey

			if _, err := m.r.ClaimService().SetContext(context.Context()).Create(model); err != nil {
				return nil, err
			}

			return &types.IdempotencyCallbackResult{RecoveryPoint: "claim_created"}, nil
		})
		if err != nil {
			return err
		}
	} else if idempotencyKey.RecoveryPoint == "claim_created" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			claims, err := m.r.ClaimService().SetContext(context.Context()).List(&models.ClaimOrder{IdempotencyKey: idempotencyKey.IdempotencyKey}, &sql.Options{})
			if err != nil {
				return nil, err
			}

			claim := claims[0]

			if claim.Type == "refund" {
				if _, err := m.r.ClaimService().SetContext(context.Context()).ProcessRefund(claim.Id); err != nil {
					return nil, err
				}
			}

			return &types.IdempotencyCallbackResult{RecoveryPoint: "refund_handled"}, nil
		})
		if err != nil {
			return err
		}
	} else if idempotencyKey.RecoveryPoint == "refund_handled" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			claims, err := m.r.ClaimService().SetContext(context.Context()).List(&models.ClaimOrder{IdempotencyKey: idempotencyKey.IdempotencyKey}, &sql.Options{Relations: []string{"return_order"}})
			if err != nil {
				return nil, err
			}

			claim := claims[0]

			if !reflect.ValueOf(claim.ReturnOrder).IsZero() {
				if _, err := m.r.ReturnService().SetContext(context.Context()).Fulfill(claim.ReturnOrder.Id); err != nil {
					return nil, err
				}
			}

			result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
				ReturnableItems: includes,
			})
			if err != nil {
				return nil, err
			}

			return &types.IdempotencyCallbackResult{
				ResponseCode: 200,
				ResponseBody: core.JSONB{
					"data": result,
				},
			}, nil
		})
		if err != nil {
			return err
		}
	} else {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).Update(idempotencyKey.IdempotencyKey, &models.IdempotencyKey{
			RecoveryPoint: "finished",
			ResponseCode:  500,
			ResponseBody:  core.JSONB{"message": "Unknown recovery point"},
		})
		if err != nil {
			return err
		}
	}

	return context.Status(idempotencyKey.ResponseCode).JSON(idempotencyKey.ResponseBody)
}

func (m *Order) CreateFulfillment(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderFulfillments](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	fulfillments, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(id, &sql.Options{Relations: []string{"fulfillments"}})
	if err != nil {
		return err
	}

	var fulfillmentIds uuid.UUIDs
	for _, fulfillment := range fulfillments.Fulfillments {
		fulfillmentIds = append(fulfillmentIds, fulfillment.Id)
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).CreateFulfillment(id, model.Items, map[string]interface{}{
		"metadata":        model.Metadata,
		"no_notification": model.NoNotification,
		"location_id":     model.LocationId,
	}); err != nil {
		return err
	}

	if model.LocationId != uuid.Nil {
		order, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(id, &sql.Options{Relations: []string{"fulfillments", "fulfillments.items", "fulfillments.items.item"}})
		if err != nil {
			return err
		}

		data := lo.Filter[models.Fulfillment](order.Fulfillments, func(item models.Fulfillment, index int) bool {
			for _, fulfillmentId := range fulfillmentIds {
				return item.Id == fulfillmentId
			}

			return false
		})

		for _, fulfillment := range data {
			items := []models.LineItem{}
			for _, item := range fulfillment.Items {
				lineItem := item.Item
				lineItem.Quantity = item.Quantity
				items = append(items, *lineItem)
			}
			if err := m.r.ProductVariantInventoryService().ValidateInventoryAtLocation(items, model.LocationId); err != nil {
				return err
			}

			for _, item := range fulfillment.Items {
				if item.Item.VariantId.UUID != uuid.Nil {
					break
				}

				if err := m.r.ProductVariantInventoryService().AdjustReservationsQuantityByLineItem(item.Item.Id, item.Item.VariantId.UUID, model.LocationId, -item.Quantity); err != nil {
					return err
				}

				if err := m.r.ProductVariantInventoryService().AdjustInventory(item.Item.VariantId.UUID, model.LocationId, -item.Quantity); err != nil {
					return err
				}
			}
		}

	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) CreateReservationForLineItem(context fiber.Ctx) error {
	model, err := api.Bind[types.OrderLineItemReservation](context, m.r.Validator())
	if err != nil {
		return err
	}

	lineItemId, err := api.BindDelete(context, "line_item_id")
	if err != nil {
		return err
	}

	lineItem, err := m.r.LineItemService().SetContext(context.Context()).Retrieve(lineItemId, &sql.Options{})
	if err != nil {
		return err
	}

	if lineItem.VariantId.UUID == uuid.Nil {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			`Can't create a reservation for a Line Item wihtout a variant`,
		)
	}

	quantity := 0
	if !reflect.ValueOf(model.Quantity).IsZero() {
		quantity = model.Quantity
	} else {
		quantity = lineItem.Quantity
	}

	reservations, err := m.r.ProductVariantInventoryService().SetContext(context.Context()).ReserveQuantity(lineItem.VariantId.UUID, quantity, services.ReserveQuantityContext{LocationId: model.LocationId})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(reservations[0])
}

func (m *Order) CreateShipment(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.CreateOrderShipment](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	trakingLinks := []models.TrackingLink{}
	for _, link := range model.TrackingNumbers {
		trakingLinks = append(trakingLinks, models.TrackingLink{TrackingNumber: link})
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).CreateShipment(id, model.FulfillmentId, trakingLinks, struct {
		NoNotification bool
		Metadata       map[string]interface{}
	}{NoNotification: model.NoNotification}); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) CreateSwapShipment(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.CreateOrderShipment](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	swapId, err := api.BindDelete(context, "swap_id")
	if err != nil {
		return err
	}

	trakingLinks := []models.TrackingLink{}
	for _, link := range model.TrackingNumbers {
		trakingLinks = append(trakingLinks, models.TrackingLink{TrackingNumber: link})
	}

	if _, err := m.r.SwapService().SetContext(context.Context()).CreateShipment(swapId, model.FulfillmentId, trakingLinks, &types.CreateShipmentConfig{NoNotification: model.NoNotification}); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) CreateSwap(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderSwap](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	headerKey := context.Get("Idempotency-Key")

	idempotencyKey := &models.IdempotencyKey{}
	idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).InitializeRequest(headerKey, context.Method(), nil, context.Path())
	if err != nil {
		return err
	}

	context.Set("Access-Control-Expose-Headers", "Idempotency-Key")
	context.Set("Idempotency-Key", idempotencyKey.IdempotencyKey)

	if idempotencyKey.RecoveryPoint == "started" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			order, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, &sql.Options{
				Relations: []string{
					"cart",
					"items",
					"items.variant",
					"items.tax_lines",
					"swaps",
					"swaps.additional_items",
					"swaps.additional_items.variant",
					"swaps.additional_items.tax_lines",
				},
			}, types.TotalsContext{})
			if err != nil {
				return nil, err
			}

			swap, err := m.r.SwapService().SetContext(context.Context()).Create(order, model.ReturnItems, model.AdditionalItems, &model.ReturnShipping, map[string]interface{}{
				// "idempotency_key": idempotencyKey.idempotency_key,
				"no_notification": model.NoNotification,
				"allow_backorder": model.AllowBackorder,
				"location_id":     model.ReturnLocationId,
			})
			if err != nil {
				return nil, err
			}

			_, err = m.r.SwapService().SetContext(context.Context()).CreateCart(swap.Id, model.CustomShippingOptions, map[string]interface{}{
				"sales_channel_id": model.SalesChannelId,
			})
			if err != nil {
				return nil, err
			}

			returnOrder, err := m.r.ReturnService().SetContext(context.Context()).RetrieveBySwap(swap.Id, []string{})
			if err != nil {
				return nil, err
			}

			if _, err := m.r.ReturnService().SetContext(context.Context()).Fulfill(returnOrder.Id); err != nil {
				return nil, err
			}

			return &types.IdempotencyCallbackResult{RecoveryPoint: "swap_created"}, nil
		})
		if err != nil {
			return err
		}
	} else if idempotencyKey.RecoveryPoint == "swap_created" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			_, err := m.r.SwapService().SetContext(context.Context()).List(&types.FilterableSwap{IdempotencyKey: idempotencyKey.IdempotencyKey}, &sql.Options{})
			if err != nil {
				return nil, err
			}

			result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
				ReturnableItems: includes,
			})
			if err != nil {
				return nil, err
			}

			return &types.IdempotencyCallbackResult{
				ResponseCode: 200,
				ResponseBody: core.JSONB{
					"data": result,
				},
			}, nil
		})
		if err != nil {
			return err
		}
	} else {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).Update(idempotencyKey.IdempotencyKey, &models.IdempotencyKey{
			RecoveryPoint: "finished",
			ResponseCode:  500,
			ResponseBody:  core.JSONB{"message": "Unknown recovery point"},
		})
		if err != nil {
			return err
		}
	}

	return context.Status(idempotencyKey.ResponseCode).JSON(idempotencyKey.ResponseBody)
}

func (m *Order) FulfillClaim(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderClaimFulfillments](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	claimId, err := api.BindDelete(context, "claim_id")
	if err != nil {
		return err
	}

	fulfillments, err := m.r.ClaimService().SetContext(context.Context()).Retrieve(claimId, &sql.Options{Relations: []string{"fulfillments"}})
	if err != nil {
		return err
	}

	var fulfillmentIds uuid.UUIDs
	for _, fulfillment := range fulfillments.Fulfillments {
		fulfillmentIds = append(fulfillmentIds, fulfillment.Id)
	}

	if _, err := m.r.ClaimService().SetContext(context.Context()).CreateFulfillment(claimId, model.NoNotification, model.LocationId, model.Metadata); err != nil {
		return err
	}

	if model.LocationId != uuid.Nil {
		order, err := m.r.ClaimService().SetContext(context.Context()).Retrieve(claimId, &sql.Options{Relations: []string{"fulfillments", "fulfillments.items", "fulfillments.items.item"}})
		if err != nil {
			return err
		}

		data := lo.Filter[models.Fulfillment](order.Fulfillments, func(item models.Fulfillment, index int) bool {
			for _, fulfillmentId := range fulfillmentIds {
				return item.Id == fulfillmentId
			}

			return false
		})

		for _, fulfillment := range data {
			items := []models.LineItem{}
			for _, item := range fulfillment.Items {
				lineItem := item.Item
				lineItem.Quantity = item.Quantity
				items = append(items, *lineItem)
			}
			if err := m.r.ProductVariantInventoryService().ValidateInventoryAtLocation(items, model.LocationId); err != nil {
				return err
			}

			for _, item := range fulfillment.Items {
				if item.Item.VariantId.UUID != uuid.Nil {
					break
				}

				if err := m.r.ProductVariantInventoryService().AdjustReservationsQuantityByLineItem(item.Item.Id, item.Item.VariantId.UUID, model.LocationId, -item.Quantity); err != nil {
					return err
				}

				if err := m.r.ProductVariantInventoryService().AdjustInventory(item.Item.VariantId.UUID, model.LocationId, -item.Quantity); err != nil {
					return err
				}
			}
		}

	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) FulfillSwap(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderClaimFulfillments](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	swapId, err := api.BindDelete(context, "swap_id")
	if err != nil {
		return err
	}

	fulfillments, err := m.r.SwapService().SetContext(context.Context()).Retrieve(swapId, &sql.Options{Relations: []string{"fulfillments"}})
	if err != nil {
		return err
	}

	var fulfillmentIds uuid.UUIDs
	for _, fulfillment := range fulfillments.Fulfillments {
		fulfillmentIds = append(fulfillmentIds, fulfillment.Id)
	}

	if _, err := m.r.SwapService().SetContext(context.Context()).CreateFulfillment(swapId, &types.CreateShipmentConfig{
		NoNotification: model.NoNotification,
		LocationId:     model.LocationId,
		Metadata:       model.Metadata,
	}); err != nil {
		return err
	}

	if model.LocationId != uuid.Nil {
		order, err := m.r.SwapService().SetContext(context.Context()).Retrieve(swapId, &sql.Options{Relations: []string{"fulfillments", "fulfillments.items", "fulfillments.items.item"}})
		if err != nil {
			return err
		}

		data := lo.Filter[models.Fulfillment](order.Fulfillments, func(item models.Fulfillment, index int) bool {
			for _, fulfillmentId := range fulfillmentIds {
				return item.Id == fulfillmentId
			}

			return false
		})

		for _, fulfillment := range data {
			items := []models.LineItem{}
			for _, item := range fulfillment.Items {
				lineItem := item.Item
				lineItem.Quantity = item.Quantity
				items = append(items, *lineItem)
			}
			if err := m.r.ProductVariantInventoryService().ValidateInventoryAtLocation(items, model.LocationId); err != nil {
				return err
			}

			for _, item := range fulfillment.Items {
				if item.Item.VariantId.UUID != uuid.Nil {
					break
				}

				if err := m.r.ProductVariantInventoryService().AdjustReservationsQuantityByLineItem(item.Item.Id, item.Item.VariantId.UUID, model.LocationId, -item.Quantity); err != nil {
					return err
				}

				if err := m.r.ProductVariantInventoryService().AdjustInventory(item.Item.VariantId.UUID, model.LocationId, -item.Quantity); err != nil {
					return err
				}
			}
		}

	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) GetReservations(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	order, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(id, &sql.Options{Relations: []string{"items"}})
	if err != nil {
		return err
	}

	var itemIds uuid.UUIDs
	for _, item := range order.Items {
		itemIds = append(itemIds, item.Id)
	}

	result, count, err := m.r.InventoryService().ListReservationItems(context.Context(), interfaces.FilterableReservationItemProps{LineItemId: itemIds}, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
}

func (m *Order) ProcessSwapPayment(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	swapId, err := api.BindDelete(context, "swap_id")
	if err != nil {
		return err
	}

	if _, err := m.r.SwapService().ProcessDifference(swapId); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) RefundPayment(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderRefunds](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.OrderService().SetContext(context.Context()).CreateRefund(id, model.Amount, model.Reason, &model.Note, &model.NoNotification); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) RequestReturn(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.OrderReturns](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	headerKey := context.Get("Idempotency-Key")

	idempotencyKey := &models.IdempotencyKey{}
	idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).InitializeRequest(headerKey, context.Method(), nil, context.Path())
	if err != nil {
		return err
	}

	context.Set("Access-Control-Expose-Headers", "Idempotency-Key")
	context.Set("Idempotency-Key", idempotencyKey.IdempotencyKey)

	if idempotencyKey.RecoveryPoint == "started" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			data := &types.CreateReturnInput{
				OrderId: id,
				Items:   model.Items,
			}

			if !reflect.ValueOf(model.LocationId).IsZero() && (m.r.InventoryService() != nil || m.r.StockLocationService() != nil) {
				data.LocationId = model.LocationId
			}

			if !reflect.ValueOf(model.ReturnShipping).IsZero() {
				data.ShippingMethod = &model.ReturnShipping
			}

			if !reflect.ValueOf(model.Refund).IsZero() && model.Refund < 0 {
				data.RefundAmount = 0
			} else {
				if !reflect.ValueOf(model.Refund).IsZero() && model.Refund >= 0 {
					data.RefundAmount = model.Refund
				}
			}

			evaluatedNoNotification := model.NoNotification

			if !reflect.ValueOf(model.NoNotification).IsZero() {
				order, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
				if err != nil {
					return nil, err
				}

				evaluatedNoNotification = order.NoNotification
			}

			data.NoNotification = evaluatedNoNotification

			createdReturn, err := m.r.ReturnService().SetContext(context.Context()).Create(data)
			if err != nil {
				return nil, err
			}

			if !reflect.ValueOf(model.ReturnShipping).IsZero() {
				if _, err := m.r.ReturnService().SetContext(context.Context()).Fulfill(createdReturn.Id); err != nil {
					return nil, err
				}
			}

			// eventBus
			//             .withTransaction(manager)
			//             .emit(OrderService.Events.RETURN_REQUESTED, {
			//               id,
			//               return_id: createdReturn.id,
			//               no_notification: evaluatedNoNotification,
			//             })

			return &types.IdempotencyCallbackResult{RecoveryPoint: "return_requested"}, nil
		})
		if err != nil {
			return err
		}
	} else if idempotencyKey.RecoveryPoint == "return_requested" {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).WorkStage(idempotencyKey.IdempotencyKey, func() (*types.IdempotencyCallbackResult, *utils.ApplictaionError) {
			if model.ReceiveNow {
				returns, err := m.r.ReturnService().SetContext(context.Context()).List(&types.FilterableReturn{IdempotencyKey: idempotencyKey.IdempotencyKey}, &sql.Options{})
				if err != nil {
					return nil, err
				}

				returnOrder := returns[0]

				if _, err := m.r.ReturnService().SetContext(context.Context()).Receive(returnOrder.Id, model.Items, &model.Refund, false, map[string]interface{}{}); err != nil {
					return nil, err
				}
			}

			result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
				ReturnableItems: includes,
			})
			if err != nil {
				return nil, err
			}

			return &types.IdempotencyCallbackResult{
				ResponseCode: 200,
				ResponseBody: core.JSONB{
					"data": result,
				},
			}, nil
		})
		if err != nil {
			return err
		}
	} else {
		idempotencyKey, err = m.r.IdempotencyKeyService().SetContext(context.Context()).Update(idempotencyKey.IdempotencyKey, &models.IdempotencyKey{
			RecoveryPoint: "finished",
			ResponseCode:  500,
			ResponseBody:  core.JSONB{"message": "Unknown recovery point"},
		})
		if err != nil {
			return err
		}
	}

	return context.Status(idempotencyKey.ResponseCode).JSON(idempotencyKey.ResponseBody)
}

func (m *Order) UpdateClaim(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.UpdateClaimInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	claimId, err := api.BindDelete(context, "claim_id")
	if err != nil {
		return err
	}

	if _, err := m.r.ClaimService().SetContext(context.Context()).Update(claimId, model); err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{
		ReturnableItems: includes,
	})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}
