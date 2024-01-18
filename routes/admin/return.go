package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Return struct {
	r Registry
}

func NewReturn(r Registry) *Return {
	m := Return{r: r}
	return &m
}

func (m *Return) SetRoutes(router fiber.Router) {
	route := router.Group("/returns")
	route.Get("/", m.List)
}

func (m *Return) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableReturn](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.ReturnService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *Return) Cancel(context fiber.Ctx) error {
	return nil
}

func (m *Return) Receive(context fiber.Ctx) error {
	return nil
}
