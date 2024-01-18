package admin

import (
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
)

type Store struct {
	r Registry
}

func NewStore(r Registry) *Store {
	m := Store{r: r}
	return &m
}

func (m *Store) Get(context fiber.Ctx) error {
	var config *sql.Options
	if err := context.Bind().Query(config); err != nil {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid query parameters",
			nil,
		)
	}

	result, err := m.r.StoreService().SetContext(context.Context()).Retrieve(config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
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
