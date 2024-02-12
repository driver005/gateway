package store

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
	route.Get("", m.List)
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
