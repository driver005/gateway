package admin

import "github.com/gofiber/fiber/v3"

type ShippingProfile struct {
	r Registry
}

func NewShippingProfile(r Registry) *ShippingProfile {
	m := ShippingProfile{r: r}
	return &m
}

func (m *ShippingProfile) Get(context fiber.Ctx) error {
	return nil
}

func (m *ShippingProfile) List(context fiber.Ctx) error {
	return nil
}

func (m *ShippingProfile) Create(context fiber.Ctx) error {
	return nil
}

func (m *ShippingProfile) Update(context fiber.Ctx) error {
	return nil
}

func (m *ShippingProfile) Delete(context fiber.Ctx) error {
	return nil
}
