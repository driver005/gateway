package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Invite struct {
	r Registry
}

func NewInvite(r Registry) *Invite {
	m := Invite{r: r}
	return &m
}

func (m *Invite) SetRoutes(router fiber.Router) {
	route := router.Group("/invites")
	route.Get("/", m.List)
	route.Post("/", m.Create)
	route.Delete("/:id", m.Delete)
}

func (m *Invite) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableInvite](context)
	if err != nil {
		return err
	}
	result, err := m.r.InviteService().SetContext(context.Context()).List(model, config)
	if err != nil {
		return err
	}

	// return context.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"data":   result,
	// 	"count":  count,
	// 	"offset": config.Skip,
	// 	"limit":  config.Take,
	// })
	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Invite) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateInviteInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	err = m.r.InviteService().SetContext(context.Context()).Create(model, 0)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).Send(nil)
}

func (m *Invite) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.InviteService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "invite",
		"deleted": true,
	})
}

func (m *Invite) Accept(context fiber.Ctx) error {
	return nil
}

func (m *Invite) Resend(context fiber.Ctx) error {
	return nil
}
