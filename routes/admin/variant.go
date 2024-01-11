package admin

import "github.com/gofiber/fiber/v3"

type Variant struct {
	r Registry
}

func NewVariant(r Registry) *Variant {
	m := Variant{r: r}
	return &m
}

func (m *Variant) Get(context fiber.Ctx) error {
	return nil
}

func (m *Variant) List(context fiber.Ctx) error {
	return nil
}

func (m *Variant) GetInventory(context fiber.Ctx) error {
	return nil
}
