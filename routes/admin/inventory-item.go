package admin

import "github.com/gofiber/fiber/v3"

type InventoryItem struct {
	r Registry
}

func NewInventoryItem(r Registry) *InventoryItem {
	m := InventoryItem{r: r}
	return &m
}

func (m *InventoryItem) SetRoutes(router fiber.Router) {
	route := router.Group("/inventory-items")
	route.Get("/:id", m.Get)
	route.Get("/", m.List)
	route.Post("/", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Get("/:id/location-levels", m.ListLocationLevels)
	route.Post("/:id/location-levels", m.CreateLocationLevel)
	route.Post("/:id/location-levels/:location_id", m.UpdateLocationLevel)
	route.Delete("/:id/location-levels/:location_id", m.DeleteLocationLevels)
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
