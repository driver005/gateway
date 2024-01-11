package admin

import "github.com/gofiber/fiber/v3"

type Batch struct {
	r Registry
}

func NewBatch(r Registry) *Batch {
	m := Batch{r: r}
	return &m
}

func (m *Auth) Get(context fiber.Ctx) error {
	return nil
}

func (m *Auth) List(context fiber.Ctx) error {
	return nil
}

func (m *Auth) Create(context fiber.Ctx) error {
	return nil
}

func (m *Auth) Cancel(context fiber.Ctx) error {
	return nil
}

func (m *Auth) Confirm(context fiber.Ctx) error {
	return nil
}
