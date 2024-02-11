package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Payment struct {
	r Registry
}

func NewPayment(r Registry) *Payment {
	m := Payment{r: r}
	return &m
}

func (m *Payment) SetRoutes(router fiber.Router) {
	route := router.Group("/payments")
	route.Get("/:id", m.Get)

	route.Post("/:id/capture", m.Capture)
	route.Post("/:id/refund", m.Refund)
}

func (m *Payment) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.PaymentService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Payment) Capture(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	result, err := m.r.PaymentService().SetContext(context.Context()).Capture(id, nil)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Payment) Refund(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.PaymentRefund](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.PaymentService().SetContext(context.Context()).Refund(id, nil, model.Amount, model.Reason, &model.Note)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}
