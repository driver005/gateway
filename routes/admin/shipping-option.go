package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type ShippingOption struct {
	r Registry
}

func NewShippingOption(r Registry) *ShippingOption {
	m := ShippingOption{r: r}
	return &m
}

func (m *ShippingOption) SetRoutes(router fiber.Router) {
	route := router.Group("/shipping-options")
	route.Get("/:id", m.Get)
	route.Get("/", m.List)
	route.Post("/", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
}

func (m *ShippingOption) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.ShippingOptionService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ShippingOption) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableShippingOption](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.ShippingOptionService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *ShippingOption) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateShippingOptionInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ShippingOptionService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ShippingOption) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateShippingOptionInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ShippingOptionService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ShippingOption) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.ShippingOptionService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "shipping-option",
		"deleted": true,
	})
}
