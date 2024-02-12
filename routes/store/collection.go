package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Collection struct {
	r Registry
}

func NewCollection(r Registry) *Collection {
	m := Collection{r: r}
	return &m
}

func (m *Collection) SetRoutes(router fiber.Router) {
	route := router.Group("/collections")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
}

func (m *Collection) Get(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	result, err := m.r.ProductCollectionService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Collection) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableCollection](context)
	if err != nil {
		return err
	}

	result, count, err := m.r.ProductCollectionService().SetContext(context.Context()).ListAndCount(model, config)
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
