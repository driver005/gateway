package admin

import "github.com/gofiber/fiber/v3"

type InventoryItem struct {
	r Registry
}

func NewInventoryItem(r Registry) *InventoryItem {
	m := InventoryItem{r: r}
	return &m
}

func (m *InventoryItem) Get(context fiber.Ctx) error {
	return nil
}

func (m *InventoryItem) List(context fiber.Ctx) error {
	return nil
}

func (m *InventoryItem) Create(context fiber.Ctx) error {
	return nil
}

func (m *InventoryItem) Update(context fiber.Ctx) error {
	return nil
}

func (m *InventoryItem) Delete(context fiber.Ctx) error {
	return nil
}

func (m *InventoryItem) CreateLocationLevel(context fiber.Ctx) error {
	return nil
}

func (m *InventoryItem) DeleteLocationLevels(context fiber.Ctx) error {
	return nil
}

func (m *InventoryItem) ListLocationLevels(context fiber.Ctx) error {
	return nil
}

func (m *InventoryItem) UpdateLocationLevel(context fiber.Ctx) error {
	return nil
}
