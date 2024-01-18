package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Notification struct {
	r Registry
}

func NewNotification(r Registry) *Notification {
	m := Notification{r: r}
	return &m
}

func (m *Notification) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableNotification](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.NotificationService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *Notification) Resend(context fiber.Ctx) error {
	return nil
}
