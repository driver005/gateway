package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type ShippingProfile struct {
	r Registry
}

func NewShippingProfile(r Registry) *ShippingProfile {
	m := ShippingProfile{r: r}
	return &m
}

func (m *ShippingProfile) SetRoutes(router fiber.Router) {
	route := router.Group("/shipping-profiles")
	route.Get("/:id", m.Get)
	route.Get("/", m.List)
	route.Post("/", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
}

func (m *ShippingProfile) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.ShippingProfileService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ShippingProfile) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableShippingProfile](context)
	if err != nil {
		return err
	}
	result, err := m.r.ShippingProfileService().SetContext(context.Context()).List(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ShippingProfile) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateShippingProfile](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ShippingProfileService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ShippingProfile) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateShippingProfile](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ShippingProfileService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ShippingProfile) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.ShippingProfileService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "shipping-profile",
		"deleted": true,
	})
}
