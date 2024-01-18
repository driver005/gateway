package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type PriceList struct {
	r Registry
}

func NewPriceList(r Registry) *PriceList {
	m := PriceList{r: r}
	return &m
}

func (m *PriceList) Get(context fiber.Ctx) error {
	return nil
}

func (m *PriceList) List(context fiber.Ctx) error {
	return nil
}

func (m *PriceList) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreatePriceListInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.PriceListService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *PriceList) Update(context fiber.Ctx) error {
	return nil
}

func (m *PriceList) Delete(context fiber.Ctx) error {
	return nil
}

func (m *PriceList) AddPricesBatch(context fiber.Ctx) error {
	return nil
}

func (m *PriceList) DeletePricesBatch(context fiber.Ctx) error {
	return nil
}

func (m *PriceList) DeleteProductPrices(context fiber.Ctx) error {
	return nil
}

func (m *PriceList) DeleteProductPricesBatch(context fiber.Ctx) error {
	return nil
}

func (m *PriceList) DeleteVariantPrices(context fiber.Ctx) error {
	return nil
}

func (m *PriceList) ListPriceListProducts(context fiber.Ctx) error {
	return nil
}
