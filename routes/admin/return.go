package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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
	route.Get("", m.List)

	route.Post("/:id/receive", m.Receive)
	route.Post("/:id/cancel", m.Cancel)
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
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	model, err := m.r.ReturnService().SetContext(context.Context()).Cancel(id)
	if err != nil {
		return err
	}

	var orderId uuid.UUID

	if model.SwapId.UUID != uuid.Nil {
		data, err := m.r.SwapService().SetContext(context.Context()).Retrieve(model.SwapId.UUID, &sql.Options{})
		if err != nil {
			return err
		}

		orderId = data.OrderId.UUID
	} else if model.ClaimOrderId.UUID != uuid.Nil {
		data, err := m.r.ClaimService().SetContext(context.Context()).Retrieve(model.ClaimOrderId.UUID, &sql.Options{})
		if err != nil {
			return err
		}

		orderId = data.OrderId.UUID
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(orderId, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Return) Receive(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.ReturnReceive](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	refundAmount := model.Refund

	if refundAmount < 0 {
		refundAmount = 0
	}

	receivedReturn, err := m.r.ReturnService().SetContext(context.Context()).Receive(id, model.Items, &refundAmount, true, map[string]interface{}{
		"locationId": model.LocationId,
	})
	if err != nil {
		return err
	}

	if receivedReturn.OrderId.UUID != uuid.Nil {
		if _, err := m.r.OrderService().SetContext(context.Context()).RegisterReturnReceived(receivedReturn.OrderId.UUID, receivedReturn, &refundAmount); err != nil {
			return err
		}
	}

	if receivedReturn.SwapId.UUID != uuid.Nil {
		if _, err := m.r.SwapService().SetContext(context.Context()).RegisterReceived(receivedReturn.SwapId.UUID); err != nil {
			return err
		}
	}

	result, err := m.r.ReturnService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}
