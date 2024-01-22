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

func (m *Region) SetRoutes(router fiber.Router) {
	route := router.Group("/regions")
	route.Get("/:id", m.Get)
	route.Get("/", m.List)
	route.Post("/", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Get("/:id/fulfillment-options", m.GetFulfillmentOptions)
	route.Post("/:id/countries", m.AddCountry)
	route.Delete("/:id/countries:country_code", m.RemoveCountry)
	route.Post("/:id/payment-providers", m.AddPaymentProvider)
	route.Delete("/:id/payment-providers/:provider_id", m.RemovePaymentProvider)
	route.Post("/:id/fulfillment-providers", m.AddFullfilmentProvider)
	route.Delete("/:id/fulfillment-providers/:provider_id", m.RemoveFullfilmentProvider)
}

func (m *Region) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Region) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableRegion](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.RegionService().SetContext(context.Context()).ListAndCount(model, config)
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
	model, id, err := api.BindUpdate[types.UpdateRegionInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Region) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.RegionService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "region",
		"deleted": true,
	})
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
