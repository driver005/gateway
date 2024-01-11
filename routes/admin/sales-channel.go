package admin

import "github.com/gofiber/fiber/v3"

type SalesChannel struct {
	r Registry
}

func NewSalesChannel(r Registry) *SalesChannel {
	m := SalesChannel{r: r}
	return &m
}

func (m *SalesChannel) Get(context fiber.Ctx) error {
	return nil
}

func (m *SalesChannel) List(context fiber.Ctx) error {
	return nil
}

func (m *SalesChannel) Create(context fiber.Ctx) error {
	return nil
}

func (m *SalesChannel) Update(context fiber.Ctx) error {
	return nil
}

func (m *SalesChannel) Delete(context fiber.Ctx) error {
	return nil
}

func (m *SalesChannel) AddProductsBatch(context fiber.Ctx) error {
	return nil
}

func (m *SalesChannel) DeleteProductsBatch(context fiber.Ctx) error {
	return nil
}

func (m *SalesChannel) AddStockLocation(context fiber.Ctx) error {
	return nil
}

func (m *SalesChannel) RemoveStockLocation(context fiber.Ctx) error {
	return nil
}
