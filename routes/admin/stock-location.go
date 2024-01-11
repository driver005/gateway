package admin

import "github.com/gofiber/fiber/v3"

type StockLocation struct {
	r Registry
}

func NewStockLocation(r Registry) *StockLocation {
	m := StockLocation{r: r}
	return &m
}

func (m *StockLocation) Get(context fiber.Ctx) error {
	return nil
}

func (m *StockLocation) List(context fiber.Ctx) error {
	return nil
}

func (m *StockLocation) Create(context fiber.Ctx) error {
	return nil
}

func (m *StockLocation) Update(context fiber.Ctx) error {
	return nil
}

func (m *StockLocation) Delete(context fiber.Ctx) error {
	return nil
}
