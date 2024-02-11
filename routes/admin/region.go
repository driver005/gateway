package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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
	route.Get("", m.List)
	route.Post("", m.Create)
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
	model, id, err := api.BindUpdate[types.RegionCountries](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.RegionService().SetContext(context.Context()).AddCountry(id, model.CountryCode); err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Region) AddFullfilmentProvider(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.RegionFulfillmentProvider](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.RegionService().SetContext(context.Context()).AddFulfillmentProvider(id, model.ProviderId); err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Region) AddPaymentProvider(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.RegionPaymentProvider](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.RegionService().SetContext(context.Context()).AddPaymentProvider(id, model.ProviderId); err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Region) GetFulfillmentOptions(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	region, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	var fpsIds uuid.UUIDs
	for _, f := range region.FulfillmentProviders {
		fpsIds = append(fpsIds, f.Id)
	}

	result, err := m.r.FulfillmentProviderService().SetContext(context.Context()).ListFulfillmentOptions(fpsIds)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Region) RemoveCountry(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	countryCode := context.Params("country_code")

	if _, err := m.r.RegionService().SetContext(context.Context()).RemoveCountry(id, countryCode); err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Region) RemoveFullfilmentProvider(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	providerId, err := api.BindDelete(context, "provider_id")
	if err != nil {
		return err
	}

	if _, err := m.r.RegionService().SetContext(context.Context()).RemoveFulfillmentProvider(id, providerId); err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Region) RemovePaymentProvider(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	providerId, err := api.BindDelete(context, "provider_id")
	if err != nil {
		return err
	}

	if _, err := m.r.RegionService().SetContext(context.Context()).RemovePaymentProvider(id, providerId); err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}
