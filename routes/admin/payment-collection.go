package admin

import "github.com/gofiber/fiber/v3"

type PaymentCollection struct {
	r Registry
}

func NewPaymentCollection(r Registry) *PaymentCollection {
	m := PaymentCollection{r: r}
	return &m
}

func (m *PaymentCollection) Get(context fiber.Ctx) error {
	return nil
}

func (m *PaymentCollection) Update(context fiber.Ctx) error {
	return nil
}

func (m *PaymentCollection) Delete(context fiber.Ctx) error {
	return nil
}

func (m *PaymentCollection) MarkAuthorized(context fiber.Ctx) error {
	return nil
}
