package admin

import "github.com/gofiber/fiber/v3"

type Payment struct {
	r Registry
}

func NewPayment(r Registry) *Payment {
	m := Payment{r: r}
	return &m
}

func (m *Payment) Get(context fiber.Ctx) error {
	return nil
}

func (m *Payment) Capture(context fiber.Ctx) error {
	return nil
}

func (m *Payment) Refund(context fiber.Ctx) error {
	return nil
}
