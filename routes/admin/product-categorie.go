package admin

import "github.com/gofiber/fiber/v3"

type ProductCategory struct {
	r Registry
}

func NewProductCategory(r Registry) *ProductCategory {
	m := ProductCategory{r: r}
	return &m
}

func (m *ProductCategory) Get(context fiber.Ctx) error {
	return nil
}

func (m *ProductCategory) List(context fiber.Ctx) error {
	return nil
}

func (m *ProductCategory) Create(context fiber.Ctx) error {
	return nil
}

func (m *ProductCategory) Update(context fiber.Ctx) error {
	return nil
}

func (m *ProductCategory) AddProductsBatch(context fiber.Ctx) error {
	return nil
}

func (m *ProductCategory) DeleteProductsBatch(context fiber.Ctx) error {
	return nil
}
