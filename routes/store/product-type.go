package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type ProductType struct {
	r Registry
}

func NewProductType(r Registry) *ProductType {
	m := ProductType{r: r}
	return &m
}

func (m *ProductType) SetRoutes(router fiber.Router) {
	route := router.Group("/product-types")
	route.Get("", m.List)
}

func (m *ProductType) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableProductType](context)
	if err != nil {
		return err
	}

	result, count, err := m.r.ProductTypeService().SetContext(context.Context()).ListAndCount(model, config)
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
