package admin

import "github.com/gofiber/fiber/v3"

type Invite struct {
	r Registry
}

func NewInvite(r Registry) *Invite {
	m := Invite{r: r}
	return &m
}

func (m *Invite) List(context fiber.Ctx) error {
	return nil
}

func (m *Invite) Create(context fiber.Ctx) error {
	return nil
}

func (m *Invite) Delete(context fiber.Ctx) error {
	return nil
}

func (m *Invite) Accept(context fiber.Ctx) error {
	return nil
}

func (m *Invite) Resend(context fiber.Ctx) error {
	return nil
}
