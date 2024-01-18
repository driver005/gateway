package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/interfaces"
	"github.com/gofiber/fiber/v3"
)

type StockLocation struct {
	r Registry
}

func NewStockLocation(r Registry) *StockLocation {
	m := StockLocation{r: r}
	return &m
}

func (m *StockLocation) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.StockLocationService().Retrieve(context.Context(), id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *StockLocation) List(context fiber.Ctx) error {
	return nil
}

func (m *StockLocation) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[interfaces.CreateStockLocationInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.StockLocationService().Create(context.Context(), *model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *StockLocation) Update(context fiber.Ctx) error {
	return nil
}

func (m *StockLocation) Delete(context fiber.Ctx) error {
	return nil
}
