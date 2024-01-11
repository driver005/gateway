package admin

import "github.com/gofiber/fiber/v3"

type ReturnReason struct {
	r Registry
}

func NewReturnReason(r Registry) *ReturnReason {
	m := ReturnReason{r: r}
	return &m
}

func (m *ReturnReason) Get(context fiber.Ctx) error {
	return nil
}

func (m *ReturnReason) List(context fiber.Ctx) error {
	return nil
}

func (m *ReturnReason) Create(context fiber.Ctx) error {
	return nil
}

func (m *ReturnReason) Update(context fiber.Ctx) error {
	return nil
}

func (m *ReturnReason) Delete(context fiber.Ctx) error {
	return nil
}
