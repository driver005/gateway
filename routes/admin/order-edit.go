package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type OrderEdit struct {
	r Registry
}

func NewOrderEdit(r Registry) *OrderEdit {
	m := OrderEdit{r: r}
	return &m
}

func (m *OrderEdit) SetRoutes(router fiber.Router) {
	route := router.Group("/order-edits")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Post("/:id/cancel", m.Cancel)
	route.Post("/:id/confirm", m.Confirm)
	route.Post("/:id/request", m.RequestConfirmation)
	route.Post("/:id/items", m.AddLineItem)
	route.Post("/:id/items/:item_id", m.UpdateLineItem)
	route.Delete("/:id/items/:item_id", m.DeleteLineItem)
	route.Delete("/:id/changes/:change_id", m.DeleteItemChange)
}

func (m *OrderEdit) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *OrderEdit) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableOrderEdit](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.OrderEditService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *OrderEdit) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateOrderEditInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	userId := api.GetUser(context)

	result, err := m.r.OrderEditService().SetContext(context.Context()).Create(model, userId)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *OrderEdit) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[models.OrderEdit](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *OrderEdit) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.OrderEditService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "order-edit",
		"deleted": true,
	})
}

func (m *OrderEdit) AddLineItem(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.AddOrderEditLineItemInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.OrderEditService().SetContext(context.Context()).AddLineItem(id, model); err != nil {
		return err
	}

	order, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(order)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *OrderEdit) Cancel(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	userId := api.GetUser(context)

	if _, err := m.r.OrderEditService().SetContext(context.Context()).Cancel(id, userId); err != nil {
		return err
	}

	order, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(order)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *OrderEdit) Confirm(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	userId := api.GetUser(context)

	if _, err := m.r.OrderEditService().SetContext(context.Context()).Confirm(id, userId); err != nil {
		return err
	}

	order, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(order)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *OrderEdit) RequestConfirmation(context fiber.Ctx) error {
	model, id, err := api.BindWithUUID[types.OrderEditsRequestConfirmation](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	userId := api.GetUser(context)

	orderEdit, err := m.r.OrderEditService().SetContext(context.Context()).RequestConfirmation(id, userId)
	if err != nil {
		return err
	}

	total, err := m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(orderEdit)
	if err != nil {
		return err
	}

	if total.DifferenceDue > 0 {
		order, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(orderEdit.OrderId.UUID, &sql.Options{
			Selects: []string{"currency_code", "region_id"},
		})
		if err != nil {
			return err
		}

		paymentCollection, err := m.r.PaymentCollectionService().SetContext(context.Context()).Create(&types.CreatePaymentCollectionInput{
			Type:         models.PaymentCollectionTypeOrderEdit,
			Amount:       total.DifferenceDue,
			CurrencyCode: order.CurrencyCode,
			RegionId:     order.RegionId.UUID,
			Description:  model.PaymentCollectionDescription,
			CreatedBy:    userId,
		})
		if err != nil {
			return err
		}

		orderEdit.PaymentCollectionId = uuid.NullUUID{UUID: paymentCollection.Id}

		_, err = m.r.OrderEditService().SetContext(context.Context()).Update(orderEdit.Id, orderEdit)
		if err != nil {
			return err
		}
	}

	order, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(order)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *OrderEdit) DeleteLineItem(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	itemId, err := utils.ParseUUID(context.Params("item_id"))
	if err != nil {
		return err
	}

	if err := m.r.OrderEditService().SetContext(context.Context()).RemoveLineItem(id, itemId); err != nil {
		return err
	}

	order, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(order)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *OrderEdit) DeleteItemChange(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	changeId, err := utils.ParseUUID(context.Params("change_id"))
	if err != nil {
		return err
	}

	if err := m.r.OrderEditService().SetContext(context.Context()).DeleteItemChange(id, changeId); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      changeId,
		"object":  "item-change",
		"deleted": true,
	})
}

func (m *OrderEdit) UpdateLineItem(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.OrderEditsEditLineItem](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	itemId, err := utils.ParseUUID(context.Params("item_id"))
	if err != nil {
		return err
	}

	if err := m.r.OrderEditService().SetContext(context.Context()).UpdateLineItem(id, itemId, model.Quantity); err != nil {
		return err
	}

	order, err := m.r.OrderEditService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.OrderEditService().SetContext(context.Context()).DecorateTotals(order)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}
