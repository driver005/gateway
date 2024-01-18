package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Order struct {
	r Registry
}

func NewOrder(r Registry) *Order {
	m := Order{r: r}
	return &m
}

func (m *Order) SetRoutes(router fiber.Router) {
	route := router.Group("/orders")
	route.Get("/:id", m.Get)
	route.Get("/", m.List)
	route.Post("/:id", m.Update)
}

func (m *Order) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableOrder](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.OrderService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *Order) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateOrderInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) AddShippingMethod(context fiber.Ctx) error {
	return nil
}

func (m *Order) Archive(context fiber.Ctx) error {
	return nil
}

func (m *Order) Cancel(context fiber.Ctx) error {
	return nil
}

func (m *Order) CancelSwap(context fiber.Ctx) error {
	return nil
}

func (m *Order) CancelClaim(context fiber.Ctx) error {
	return nil
}

func (m *Order) CancelFullfillmentClaim(context fiber.Ctx) error {
	return nil
}

func (m *Order) CancelFullfillmentSwap(context fiber.Ctx) error {
	return nil
}

func (m *Order) CancelFullfillment(context fiber.Ctx) error {
	return nil
}

func (m *Order) CapturePayment(context fiber.Ctx) error {
	return nil
}

func (m *Order) Complete(context fiber.Ctx) error {
	return nil
}

func (m *Order) CreateClaimShippment(context fiber.Ctx) error {
	return nil
}

func (m *Order) CreateClaim(context fiber.Ctx) error {
	return nil
}

func (m *Order) CreateFulfillment(context fiber.Ctx) error {
	return nil
}

func (m *Order) CreateReservationForLineItem(context fiber.Ctx) error {
	return nil
}

func (m *Order) CreateShipment(context fiber.Ctx) error {
	return nil
}

func (m *Order) CreateSwapShipment(context fiber.Ctx) error {
	return nil
}

func (m *Order) CreateSwap(context fiber.Ctx) error {
	return nil
}

func (m *Order) FulfillClaim(context fiber.Ctx) error {
	return nil
}

func (m *Order) FulfillSwap(context fiber.Ctx) error {
	return nil
}

func (m *Order) GetReservations(context fiber.Ctx) error {
	return nil
}

func (m *Order) ProcessSwapPayment(context fiber.Ctx) error {
	return nil
}

func (m *Order) RefundPayment(context fiber.Ctx) error {
	return nil
}

func (m *Order) RequestReturn(context fiber.Ctx) error {
	return nil
}

func (m *Order) UpdateClaim(context fiber.Ctx) error {
	return nil
}
