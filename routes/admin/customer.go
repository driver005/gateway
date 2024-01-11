package admin

import "github.com/gofiber/fiber/v3"

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
	return nil
}

func (m *Customer) Update(context fiber.Ctx) error {
	return nil
}
