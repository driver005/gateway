package admin

import "github.com/gofiber/fiber/v3"

type Return struct {
	r Registry
}

func NewReturn(r Registry) *Return {
	m := Return{r: r}
	return &m
}

func (m *Return) List(context fiber.Ctx) error {
	return nil
}

func (m *Return) Cancel(context fiber.Ctx) error {
	return nil
}

func (m *Return) Receive(context fiber.Ctx) error {
	return nil
}
