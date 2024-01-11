package admin

import "github.com/gofiber/fiber/v3"

type Notification struct {
	r Registry
}

func NewNotification(r Registry) *Notification {
	m := Notification{r: r}
	return &m
}

func (m *Notification) List(context fiber.Ctx) error {
	return nil
}

func (m *Notification) Resend(context fiber.Ctx) error {
	return nil
}
