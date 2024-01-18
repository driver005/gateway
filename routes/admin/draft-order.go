package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type DraftOrder struct {
	r Registry
}

func NewDraftOrder(r Registry) *DraftOrder {
	m := DraftOrder{r: r}
	return &m
}

func (m *DraftOrder) Get(context fiber.Ctx) error {
	return nil
}

func (m *DraftOrder) List(context fiber.Ctx) error {
	return nil
}

func (m *DraftOrder) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.DraftOrderCreate](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.DraftOrderService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *DraftOrder) Update(context fiber.Ctx) error {
	return nil
}

func (m *DraftOrder) Delete(context fiber.Ctx) error {
	return nil
}

func (m *DraftOrder) CreateLineItem(context fiber.Ctx) error {
	return nil
}

func (m *DraftOrder) RegisterPayment(context fiber.Ctx) error {
	return nil
}

func (m *DraftOrder) UpdateLineItem(context fiber.Ctx) error {
	return nil
}
