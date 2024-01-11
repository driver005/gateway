package admin

import "github.com/gofiber/fiber/v3"

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
	return nil
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
