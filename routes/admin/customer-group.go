package admin

import "github.com/gofiber/fiber/v3"

type CustomerGroup struct {
	r Registry
}

func NewCustomerGroup(r Registry) *CustomerGroup {
	m := CustomerGroup{r: r}
	return &m
}

func (m *CustomerGroup) Get(context fiber.Ctx) error {
	return nil
}

func (m *CustomerGroup) List(context fiber.Ctx) error {
	return nil
}

func (m *CustomerGroup) Create(context fiber.Ctx) error {
	return nil
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
