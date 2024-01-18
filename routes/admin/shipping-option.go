package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type ShippingOption struct {
	r Registry
}

func NewShippingOption(r Registry) *ShippingOption {
	m := ShippingOption{r: r}
	return &m
}

func (m *ShippingOption) Get(context fiber.Ctx) error {
	return nil
}

func (m *ShippingOption) List(context fiber.Ctx) error {
	return nil
}

func (m *ShippingOption) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateShippingOptionInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ShippingOptionService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ShippingOption) Update(context fiber.Ctx) error {
	return nil
}

func (m *ShippingOption) Delete(context fiber.Ctx) error {
	return nil
}
