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

func (m *Customer) Get(context fiber.Ctx) error {
	return nil
}

func (m *Customer) List(context fiber.Ctx) error {
	return nil
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
	return nil
}
