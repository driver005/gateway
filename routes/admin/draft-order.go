package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
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
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.DraftOrderService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *DraftOrder) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableDraftOrder](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.DraftOrderService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
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
	model, id, err := api.BindUpdate[models.DraftOrder](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.DraftOrderService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *DraftOrder) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.DraftOrderService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "draft-order",
		"deleted": true,
	})
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
