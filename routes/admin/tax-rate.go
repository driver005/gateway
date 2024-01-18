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
	return nil
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
	return nil
}

func (m *TaxRate) AddProductTypes(context fiber.Ctx) error {
	return nil
}

func (m *TaxRate) AddToProducts(context fiber.Ctx) error {
	return nil
}

func (m *TaxRate) AddToShippingOptions(context fiber.Ctx) error {
	return nil
}

func (m *TaxRate) RemoveFromProductTypes(context fiber.Ctx) error {
	return nil
}

func (m *TaxRate) RemoveFromProducts(context fiber.Ctx) error {
	return nil
}

func (m *TaxRate) RemoveFromShippingOptions(context fiber.Ctx) error {
	return nil
}
