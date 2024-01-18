package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Region struct {
	r Registry
}

func NewRegion(r Registry) *Region {
	m := Region{r: r}
	return &m
}

func (m *Region) Get(context fiber.Ctx) error {
	return nil
}

func (m *Region) List(context fiber.Ctx) error {
	return nil
}

func (m *Region) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateRegionInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Region) Update(context fiber.Ctx) error {
	return nil
}

func (m *Region) Delete(context fiber.Ctx) error {
	return nil
}

func (m *Region) AddCountry(context fiber.Ctx) error {
	return nil
}

func (m *Region) AddFullfilmentProvider(context fiber.Ctx) error {
	return nil
}

func (m *Region) AddPaymentProvider(context fiber.Ctx) error {
	return nil
}

func (m *Region) GetFulfillmentOptions(context fiber.Ctx) error {
	return nil
}

func (m *Region) RemoveCountry(context fiber.Ctx) error {
	return nil
}

func (m *Region) RemoveFullfilmentProvider(context fiber.Ctx) error {
	return nil
}

func (m *Region) RemovePaymentProvider(context fiber.Ctx) error {
	return nil
}
