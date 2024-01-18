package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Discount struct {
	r Registry
}

func NewDiscount(r Registry) *Discount {
	m := Discount{r: r}
	return &m
}

func (m *Discount) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableDiscount](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.DiscountService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *Discount) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateDiscountInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateDiscountInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.DiscountService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "disocunt",
		"deleted": true,
	})
}

func (m *Discount) GetConditon(context fiber.Ctx) error {
	return nil
}

func (m *Discount) GetDiscountByCode(context fiber.Ctx) error {
	return nil
}

func (m *Discount) AddRegion(context fiber.Ctx) error {
	return nil
}

func (m *Discount) AddResourcesToConditionBatch(context fiber.Ctx) error {
	return nil
}

func (m *Discount) CreateConditon(context fiber.Ctx) error {
	return nil
}

func (m *Discount) CreateDynamicCode(context fiber.Ctx) error {
	return nil
}

func (m *Discount) DeleteConditon(context fiber.Ctx) error {
	return nil
}

func (m *Discount) DeleteDynamicCode(context fiber.Ctx) error {
	return nil
}

func (m *Discount) DeleteResourcesToConditionBatch(context fiber.Ctx) error {
	return nil
}

func (m *Discount) RemoveRegion(context fiber.Ctx) error {
	return nil
}
