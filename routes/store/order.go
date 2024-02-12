package store

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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
	route.Get("", m.Lookup)
	route.Get("/:id", m.Get)

	route.Get("/cart/:cart_id", m.GetByCart)
	route.Post("/customer/confirm", m.ConfirmRequest)
	route.Post("/batch/customer/token", m.Request)
}

func (m *Order) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByIdWithTotals(id, config, types.TotalsContext{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) GetByCart(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "cart_id")
	if err != nil {
		return err
	}

	result, err := m.r.OrderService().SetContext(context.Context()).RetrieveByCartIdWithTotals(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Order) Lookup(context fiber.Ctx) error {
	model, err := api.Bind[types.OrderLookup](context, m.r.Validator())
	if err != nil {
		return err
	}

	orders, err := m.r.OrderService().SetContext(context.Context()).List(&types.FilterableOrder{
		DisplayId: model.DisplayId,
		Email:     model.Email,
	}, &sql.Options{})
	if err != nil {
		return err
	}

	if len(orders) != 1 {
		return context.SendStatus(fiber.StatusNotFound)
	}

	return context.Status(fiber.StatusOK).JSON(orders[0])
}

func (m *Order) ConfirmRequest(context fiber.Ctx) error {
	model, err := api.Bind[types.CustomerAcceptClaim](context, m.r.Validator())
	if err != nil {
		return err
	}

	_, claims, errObj := m.r.TockenService().SetContext(context.Context()).VerifyToken(model.Token)
	if errObj != nil {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Invalid token",
		)
	}

	customerId := claims["claimingCustomerId"].(uuid.UUID)

	customer, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(customerId, &sql.Options{})
	if err != nil {
		return err
	}

	orders, err := m.r.OrderService().SetContext(context.Context()).List(&types.FilterableOrder{FilterModel: core.FilterModel{Id: claims["orders"].(uuid.UUIDs)}}, &sql.Options{})
	if err != nil {
		return err
	}

	for _, order := range orders {
		if _, err := m.r.OrderService().SetContext(context.Context()).Update(order.Id, &types.UpdateOrderInput{
			CustomerId: customerId,
			Email:      customer.Email,
		}); err != nil {
			return err
		}
	}

	return context.SendStatus(fiber.StatusOK)
}

func (m *Order) Request(context fiber.Ctx) error {
	model, err := api.Bind[types.CustomerOrderClaim](context, m.r.Validator())
	if err != nil {
		return err
	}

	customerId := context.Locals("customer_id").(uuid.UUID)

	customer, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(customerId, &sql.Options{})
	if err != nil {
		return err
	}

	if !customer.HasAccount {
		utils.NewApplictaionError(
			utils.UNAUTHORIZED,
			"Customer does not have an account",
		)
	}

	orders, err := m.r.OrderService().SetContext(context.Context()).List(&types.FilterableOrder{FilterModel: core.FilterModel{Id: model.OrderIds}}, &sql.Options{Selects: []string{"id", "email"}})
	if err != nil {
		return err
	}

	emailOrderMapping := make(map[string]uuid.UUIDs)
	for _, order := range orders {
		emailOrderMapping[order.Email] = append(emailOrderMapping[order.Email], order.Id)
	}

	// 1. email
	for _, ids := range emailOrderMapping {
		// 1. token
		_, errObj := m.r.TockenService().SetContext(context.Context()).SignToken(map[string]interface{}{
			"claimingCustomerId": customerId,
			"orders":             ids,
		})
		if errObj != nil {
			return utils.NewApplictaionError(
				utils.NOT_ALLOWED,
				"Invalid token",
			)
		}

		// err = eventBusService.Emit(TokenEventsOrderUpdateTokenCreated, TokenEventPayload{
		// 	OldEmail:      email,
		// 	NewCustomerID: customer.ID,
		// 	Orders:        ids,
		// 	Token:         token,
		// })
		// if err != nil {
		// 	return err
		// }
	}

	return context.SendStatus(fiber.StatusOK)
}
