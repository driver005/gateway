package admin

import "github.com/gofiber/fiber/v3"

type Store struct {
	r Registry
}

func NewStore(r Registry) *Store {
	m := Store{r: r}
	return &m
}

func (m *Store) Get(context fiber.Ctx) error {
	return nil
}

func (m *Store) Update(context fiber.Ctx) error {
	return nil
}

func (m *Store) AddCurrency(context fiber.Ctx) error {
	return nil
}

func (m *Store) RemoveCurrency(context fiber.Ctx) error {
	return nil
}

func (m *Store) ListPaymentProviders(context fiber.Ctx) error {
	return nil
}

func (m *Store) ListTaxProviders(context fiber.Ctx) error {
	return nil
}
