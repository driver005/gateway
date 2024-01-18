package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type ProductCategory struct {
	r Registry
}

func NewProductCategory(r Registry) *ProductCategory {
	m := ProductCategory{r: r}
	return &m
}

func (m *ProductCategory) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.ProductCategoryService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ProductCategory) List(context fiber.Ctx) error {
	return nil
}

func (m *ProductCategory) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateProductCategoryInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductCategoryService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
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
