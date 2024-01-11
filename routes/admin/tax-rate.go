package admin

import "github.com/gofiber/fiber/v3"

type TaxRate struct {
	r Registry
}

func NewTaxRate(r Registry) *TaxRate {
	m := TaxRate{r: r}
	return &m
}

func (m *TaxRate) Get(context fiber.Ctx) error {
	return nil
}

func (m *TaxRate) List(context fiber.Ctx) error {
	return nil
}

func (m *TaxRate) Create(context fiber.Ctx) error {
	return nil
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
