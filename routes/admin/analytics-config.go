package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type AnalyticsConfig struct {
	r Registry
}

func NewAnalyticsConfig(r Registry) *AnalyticsConfig {
	m := AnalyticsConfig{r: r}
	return &m
}

func (m *AnalyticsConfig) SetRoutes(router fiber.Router) {
	route := router.Group("/analytics-configs")
	route.Get("", m.Get)
	route.Post("", m.Create)
	route.Post("/update", m.Update)
	route.Delete("", m.Delete)
}

func (m *AnalyticsConfig) Get(context fiber.Ctx) error {
	userId := api.GetUser(context)

	user, err := m.r.AnalyticsConfigService().SetContext(context.Context()).Retrive(userId)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(user)
}

func (m *AnalyticsConfig) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateAnalyticsConfig](context, m.r.Validator())
	if err != nil {
		return err
	}

	userId := api.GetUser(context)

	result, err := m.r.AnalyticsConfigService().SetContext(context.Context()).Create(userId, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *AnalyticsConfig) Update(context fiber.Ctx) error {
	model, err := api.BindCreate[types.UpdateAnalyticsConfig](context, m.r.Validator())
	if err != nil {
		return err
	}

	userId := api.GetUser(context)

	result, err := m.r.AnalyticsConfigService().SetContext(context.Context()).Update(userId, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *AnalyticsConfig) Delete(context fiber.Ctx) error {
	userId := api.GetUser(context)

	if err := m.r.AnalyticsConfigService().SetContext(context.Context()).Delete(userId); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_id": userId,
		"object":  "analytics_config",
		"deleted": true,
	})
}
