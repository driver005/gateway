package admin

import (
	"github.com/driver005/gateway/api"
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

func (m *OrderEdit) Get(context fiber.Ctx) error {
	return nil
}

func (m *OrderEdit) List(context fiber.Ctx) error {
	return nil
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
	return nil
}

func (m *OrderEdit) Delete(context fiber.Ctx) error {
	return nil
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
