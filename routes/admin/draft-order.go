package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type DraftOrder struct {
	r Registry
}

func NewDraftOrder(r Registry) *DraftOrder {
	m := DraftOrder{r: r}
	return &m
}

func (m *DraftOrder) SetRoutes(router fiber.Router) {
	route := router.Group("/draft-orders")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Delete("/:id/line-items/:line_id", m.DeleteLineItem)
	route.Post("/:id/line-items", m.CreateLineItem)
	route.Post("/:id/line-items/:line_id", m.UpdateLineItem)
	route.Post("/:id/pay", m.RegisterPayment)
}

func (m *DraftOrder) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.DraftOrderService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *DraftOrder) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableDraftOrder](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.DraftOrderService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *DraftOrder) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.DraftOrderCreate](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.DraftOrderService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *DraftOrder) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[models.DraftOrder](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.DraftOrderService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *DraftOrder) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.DraftOrderService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "draft-order",
		"deleted": true,
	})
}

func (m *DraftOrder) CreateLineItem(context fiber.Ctx) error {
	model, id, err := api.BindWithUUID[types.Item](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	draftOrder, err := m.r.DraftOrderService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	if draftOrder.Status != models.DraftOrderStatusCompleted {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"You are only allowed to update open draft orders",
		)
	}

	if model.VariantId == uuid.Nil {
		line, err := m.r.LineItemService().SetContext(context.Context()).Generate(
			model.VariantId,
			nil,
			draftOrder.Cart.RegionId.UUID,
			model.Quantity,
			types.GenerateLineItemContext{
				Metadata:  model.Metadata,
				UnitPrice: model.UnitPrice,
			},
		)
		if err != nil {
			return err
		}

		if err := m.r.CartService().SetContext(context.Context()).AddOrUpdateLineItems(draftOrder.CartId.UUID, line, false); err != nil {
			return err
		}
	} else {
		_, err := m.r.LineItemService().SetContext(context.Context()).Create(
			[]models.LineItem{
				{
					Model: core.Model{
						Metadata: model.Metadata,
					},
					CartId:         draftOrder.CartId,
					HasShipping:    true,
					Title:          model.Title,
					AllowDiscounts: false,
					UnitPrice:      model.UnitPrice,
					Quantity:       model.Quantity,
				},
			},
		)
		if err != nil {
			return err
		}
	}

	cart, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(draftOrder.CartId.UUID, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	draftOrder.Cart = *cart

	return context.Status(fiber.StatusOK).JSON(draftOrder)
}

func (m *DraftOrder) RegisterPayment(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	draftOrder, err := m.r.DraftOrderService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	cart, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(draftOrder.CartId.UUID, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	_, err = m.r.PaymentProviderService().SetContext(context.Context()).CreateSession(uuid.Nil, &types.PaymentSessionInput{
		PaymentSessionId:   cart.PaymentSession.Id,
		ProviderId:         cart.PaymentSession.ProviderId.UUID,
		Cart:               cart,
		Customer:           cart.Customer,
		CurrencyCode:       cart.Payment.CurrencyCode,
		Amount:             cart.Payment.Amount,
		ResourceId:         cart.Id,
		PaymentSessionData: cart.PaymentSession.Data,
		Context:            cart.Context,
	})
	if err != nil {
		return err
	}

	if err := m.r.CartService().SetContext(context.Context()).SetPaymentSession(cart.Id, uuid.Nil); err != nil {
		return err
	}

	if err := m.r.CartService().SetContext(context.Context()).CreateTaxLines(cart.Id, nil); err != nil {
		return err
	}

	_, err = m.r.CartService().SetContext(context.Context()).AuthorizePayment(cart.Id, nil, map[string]interface{}{})
	if err != nil {
		return err
	}

	order, err := m.r.OrderService().SetContext(context.Context()).CreateFromCart(id, nil)
	if err != nil {
		return err
	}

	_, err = m.r.DraftOrderService().SetContext(context.Context()).RegisterCartCompletion(id, order.Id)
	if err != nil {
		return err
	}

	_, err = m.r.OrderService().SetContext(context.Context()).CapturePayment(order.Id)
	if err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, &sql.Options{}, types.TotalsContext{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *DraftOrder) UpdateLineItem(context fiber.Ctx) error {
	model, id, err := api.BindWithUUID[types.Item](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	lineId, err := utils.ParseUUID(context.Params("line_id"))
	if err != nil {
		return err
	}

	draftOrder, err := m.r.DraftOrderService().SetContext(context.Context()).Retrieve(id, &sql.Options{Relations: []string{"cart", "cart.items"}})
	if err != nil {
		return err
	}

	if draftOrder.Status == models.DraftOrderStatusCompleted {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"You are only allowed to update open draft orders",
		)
	}

	if model.Quantity == 0 {
		if err := m.r.CartService().SetContext(context.Context()).RemoveLineItem(draftOrder.CartId.UUID, uuid.UUIDs{lineId}); err != nil {
			return err
		}
	} else {
		_, ok := lo.Find(draftOrder.Cart.Items, func(v models.LineItem) bool {
			return v.Id == lineId
		})

		if !ok {
			return utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Could not find the line item",
			)
		}

		item := &types.LineItemUpdate{
			RegionId: draftOrder.Cart.RegionId.UUID,
		}

		item.Title = model.Title
		item.UnitPrice = model.UnitPrice
		item.VariantId = model.VariantId
		item.Quantity = model.Quantity
		item.Metadata = model.Metadata

		_, err := m.r.CartService().SetContext(context.Context()).UpdateLineItem(draftOrder.CartId.UUID, lineId, item)
		if err != nil {
			return err
		}
	}

	cart, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(draftOrder.CartId.UUID, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	draftOrder.Cart = *cart

	return context.Status(fiber.StatusOK).JSON(draftOrder)
}

func (m *DraftOrder) DeleteLineItem(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	lineId, err := utils.ParseUUID(context.Params("line_id"))
	if err != nil {
		return err
	}

	draftOrder, err := m.r.DraftOrderService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	if draftOrder.Status == models.DraftOrderStatusCompleted {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"You are only allowed to update open draft orders",
		)
	}

	if err := m.r.CartService().SetContext(context.Context()).RemoveLineItem(draftOrder.CartId.UUID, uuid.UUIDs{lineId}); err != nil {
		return err
	}

	cart, err := m.r.CartService().SetContext(context.Context()).RetrieveWithTotals(draftOrder.CartId.UUID, &sql.Options{}, services.TotalsConfig{})
	if err != nil {
		return err
	}

	draftOrder.Cart = *cart

	return context.Status(fiber.StatusOK).JSON(draftOrder)
}
