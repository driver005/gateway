package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Swap struct {
	r Registry
}

func NewSwap(r Registry) *Swap {
	m := Swap{r: r}
	return &m
}

func (m *Swap) SetRoutes(router fiber.Router) {
	route := router.Group("/swaps")
	route.Get("/:id", m.Get)
	route.Get("/", m.List)
}

func (m *Swap) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.SwapService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Swap) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableSwap](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.SwapService().SetContext(context.Context()).ListAndCount(model, config)
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
