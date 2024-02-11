package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Customer struct {
	r Registry
}

func NewCustomer(r Registry) *Customer {
	m := Customer{r: r}
	return &m
}

func (m *Customer) SetRoutes(router fiber.Router) {
	route := router.Group("/customers")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
}

func (m *Customer) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.CustomerService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Customer) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableCustomer](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.CustomerService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *Customer) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateCustomerInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Customer) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateCustomerInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CustomerService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}
