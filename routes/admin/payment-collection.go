package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/gofiber/fiber/v3"
)

type PaymentCollection struct {
	r Registry
}

func NewPaymentCollection(r Registry) *PaymentCollection {
	m := PaymentCollection{r: r}
	return &m
}

func (m *PaymentCollection) SetRoutes(router fiber.Router) {
	route := router.Group("/payment-collections")
	route.Get("/:id", m.Get)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
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
	model, id, err := api.BindUpdate[models.PaymentCollection](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.PaymentCollectionService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *PaymentCollection) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.PaymentCollectionService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "payment-collection",
		"deleted": true,
	})
}

func (m *PaymentCollection) MarkAuthorized(context fiber.Ctx) error {
	return nil
}
