package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type CustomerGroup struct {
	r Registry
}

func NewCustomerGroup(r Registry) *CustomerGroup {
	m := CustomerGroup{r: r}
	return &m
}

func (m *CustomerGroup) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.CustomerGroupService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *CustomerGroup) List(context fiber.Ctx) error {
	return nil
}

func (m *CustomerGroup) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateCustomerGroup](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CustomerGroupService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *CustomerGroup) Update(context fiber.Ctx) error {
	return nil
}

func (m *CustomerGroup) Delete(context fiber.Ctx) error {
	return nil
}

func (m *CustomerGroup) AddCustomers(context fiber.Ctx) error {
	return nil
}

func (m *CustomerGroup) GetBatch(context fiber.Ctx) error {
	return nil
}

func (m *CustomerGroup) DeleteBatch(context fiber.Ctx) error {
	return nil
}
