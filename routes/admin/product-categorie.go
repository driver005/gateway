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

func (m *ProductCategory) SetRoutes(router fiber.Router) {
	route := router.Group("/product-categories")
	route.Get("/:id", m.Get)
	route.Get("/", m.List)
	route.Post("/", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
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
	model, config, err := api.BindList[types.FilterableProductCategory](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.ProductCategoryService().SetContext(context.Context()).ListAndCount(model, config)
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
	model, id, err := api.BindUpdate[types.UpdateProductCategoryInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductCategoryService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ProductCategory) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.ProductCategoryService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "product-category",
		"deleted": true,
	})
}

func (m *ProductCategory) AddProductsBatch(context fiber.Ctx) error {
	return nil
}

func (m *ProductCategory) DeleteProductsBatch(context fiber.Ctx) error {
	return nil
}
