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

func (m *StockLocation) SetRoutes(router fiber.Router) {
	route := router.Group("/stock-locations")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
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
	model, config, err := api.BindList[interfaces.FilterableStockLocation](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.StockLocationService().ListAndCount(context.Context(), *model, config)
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
	model, id, err := api.BindUpdate[interfaces.UpdateStockLocationInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.StockLocationService().Update(context.Context(), id, *model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *StockLocation) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.StockLocationService().Delete(context.Context(), id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "stock-location",
		"deleted": true,
	})
}
