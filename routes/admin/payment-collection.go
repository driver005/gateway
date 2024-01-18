package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/gofiber/fiber/v3"
)

type PaymentCollection struct {
	r Registry
}

func NewPaymentCollection(r Registry) *PaymentCollection {
	m := PaymentCollection{r: r}
	return &m
}

func (m *PaymentCollection) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
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
