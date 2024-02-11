package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type TaxRate struct {
	r Registry
}

func NewTaxRate(r Registry) *TaxRate {
	m := TaxRate{r: r}
	return &m
}

func (m *TaxRate) SetRoutes(router fiber.Router) {
	route := router.Group("/store")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Post("/:id/products/batch", m.AddToProducts)
	route.Post("/:id/product-types/batch", m.AddProductTypes)
	route.Post("/:id/shipping-options/batch", m.AddToShippingOptions)
	route.Delete("/:id/products/batch", m.RemoveFromProducts)
	route.Delete("/:id/product-types/batch", m.RemoveFromProductTypes)
	route.Delete("/:id/shipping-options/batch", m.RemoveFromShippingOptions)
}

func (m *TaxRate) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.TaxRateService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *TaxRate) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableTaxRate](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.TaxRateService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *TaxRate) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateTaxRateInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *TaxRate) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateTaxRateInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *TaxRate) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.TaxRateService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "tax-rate",
		"deleted": true,
	})
}

func (m *TaxRate) AddProductTypes(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.TaxRateProductTypes](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.TaxRateService().SetContext(context.Context()).AddToProductType(id, model.ProductTypes, false); err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *TaxRate) AddToProducts(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.TaxRateProducts](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.TaxRateService().SetContext(context.Context()).AddToProduct(id, model.Products, false); err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *TaxRate) AddToShippingOptions(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.TaxRateShippingOptions](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.TaxRateService().SetContext(context.Context()).AddToShippingOption(id, model.ShippingOptions, false); err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *TaxRate) RemoveFromProductTypes(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.TaxRateProductTypes](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.TaxRateService().SetContext(context.Context()).RemoveFromProductType(id, model.ProductTypes); err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *TaxRate) RemoveFromProducts(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.TaxRateProducts](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.TaxRateService().SetContext(context.Context()).RemoveFromProduct(id, model.Products); err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *TaxRate) RemoveFromShippingOptions(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.TaxRateShippingOptions](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if err := m.r.TaxRateService().SetContext(context.Context()).RemoveFromShippingOption(id, model.ShippingOptions); err != nil {
		return err
	}

	result, err := m.r.TaxRateService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}
