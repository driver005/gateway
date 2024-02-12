package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Region struct {
	r Registry
}

func NewRegion(r Registry) *Region {
	m := Region{r: r}
	return &m
}

func (m *Region) SetRoutes(router fiber.Router) {
	route := router.Group("/regions")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
}

func (m *Region) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	result, err := m.r.RegionService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Region) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableRegion](context)
	if err != nil {
		return err
	}

	result, count, err := m.r.RegionService().SetContext(context.Context()).ListAndCount(model, config)
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
