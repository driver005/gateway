package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
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
	route.Get("/", m.List)
	route.Post("/", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
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

	result, err := m.r.OrderEditService().SetContext(context.Context()).Create(model, "")
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
	return nil
}

func (m *OrderEdit) Cancel(context fiber.Ctx) error {
	return nil
}

func (m *OrderEdit) Confirm(context fiber.Ctx) error {
	return nil
}

func (m *OrderEdit) RequestConfirmation(context fiber.Ctx) error {
	return nil
}

func (m *OrderEdit) DeleteLineItem(context fiber.Ctx) error {
	return nil
}

func (m *OrderEdit) DeleteItemChange(context fiber.Ctx) error {
	return nil
}

func (m *OrderEdit) UpdateLineItem(context fiber.Ctx) error {
	return nil
}
