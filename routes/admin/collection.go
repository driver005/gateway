package admin

import "github.com/gofiber/fiber/v3"

type Collection struct {
	r Registry
}

func NewCollection(r Registry) *Collection {
	m := Collection{r: r}
	return &m
}

func (m *Collection) Get(context fiber.Ctx) error {
	return nil
}

func (m *Collection) List(context fiber.Ctx) error {
	return nil
}

func (m *Collection) Create(context fiber.Ctx) error {
	return nil
}

func (m *Collection) Update(context fiber.Ctx) error {
	return nil
}

func (m *Collection) Delete(context fiber.Ctx) error {
	return nil
}

func (m *Collection) AddProducts(context fiber.Ctx) error {
	return nil
}

func (m *Collection) RemoveProducts(context fiber.Ctx) error {
	return nil
}
