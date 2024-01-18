package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type ShippingProfile struct {
	r Registry
}

func NewShippingProfile(r Registry) *ShippingProfile {
	m := ShippingProfile{r: r}
	return &m
}

func (m *ShippingProfile) Get(context fiber.Ctx) error {
	return nil
}

func (m *ShippingProfile) List(context fiber.Ctx) error {
	return nil
}

func (m *ShippingProfile) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateShippingProfile](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ShippingProfileService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ShippingProfile) Update(context fiber.Ctx) error {
	return nil
}

func (m *ShippingProfile) Delete(context fiber.Ctx) error {
	return nil
}
