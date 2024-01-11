package admin

import "github.com/gofiber/fiber/v3"

type ShippingOption struct {
	r Registry
}

func NewShippingOption(r Registry) *ShippingOption {
	m := ShippingOption{r: r}
	return &m
}

func (m *ShippingOption) Get(context fiber.Ctx) error {
	return nil
}

func (m *ShippingOption) List(context fiber.Ctx) error {
	return nil
}

func (m *ShippingOption) Create(context fiber.Ctx) error {
	return nil
}

func (m *ShippingOption) Update(context fiber.Ctx) error {
	return nil
}

func (m *ShippingOption) Delete(context fiber.Ctx) error {
	return nil
}
