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

func (m *CustomerGroup) SetRoutes(router fiber.Router) {
	route := router.Group("/customer-groups")
	route.Get("/:id", m.Get)
	route.Get("/", m.List)
	route.Post("/", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)
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
	model, config, err := api.BindList[types.FilterableCustomerGroup](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.CustomerGroupService().SetContext(context.Context()).ListAndCount(model, config)
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
	model, id, err := api.BindUpdate[types.UpdateCustomerGroup](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CustomerGroupService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *CustomerGroup) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.CustomerGroupService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "cusomer-group",
		"deleted": true,
	})
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
