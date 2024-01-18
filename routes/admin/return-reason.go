package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type ReturnReason struct {
	r Registry
}

func NewReturnReason(r Registry) *ReturnReason {
	m := ReturnReason{r: r}
	return &m
}

func (m *ReturnReason) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.ReturnReasonService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ReturnReason) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableReturnReason](context)
	if err != nil {
		return err
	}
	result, err := m.r.ReturnReasonService().SetContext(context.Context()).List(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ReturnReason) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateReturnReason](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ReturnReasonService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ReturnReason) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateReturnReason](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ReturnReasonService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ReturnReason) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.ReturnReasonService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "return-reason",
		"deleted": true,
	})
}
