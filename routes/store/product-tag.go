package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type ProductTag struct {
	r Registry
}

func NewProductTag(r Registry) *ProductTag {
	m := ProductTag{r: r}
	return &m
}

func (m *ProductTag) SetRoutes(router fiber.Router) {
	route := router.Group("/product-tags")
	route.Get("", m.List)
}

func (m *ProductTag) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableProductTag](context)
	if err != nil {
		return err
	}

	result, count, err := m.r.ProductTagService().SetContext(context.Context()).ListAndCount(model, config)
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
