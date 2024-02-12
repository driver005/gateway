package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type ReturnReason struct {
	r Registry
}

func NewReturnReason(r Registry) *ReturnReason {
	m := ReturnReason{r: r}
	return &m
}

func (m *ReturnReason) SetRoutes(router fiber.Router) {
	route := router.Group("/return-reasons")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
}

func (m *ReturnReason) Get(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	result, err := m.r.ReturnReasonService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *ReturnReason) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableReturnReason](context)
	if err != nil {
		return err
	}

	model.ParentReturnReasonId = uuid.Nil

	result, err := m.r.ReturnReasonService().SetContext(context.Context()).List(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}
