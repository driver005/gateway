package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type Currencie struct {
	r Registry
}

func NewCurrencie(r Registry) *Currencie {
	m := Currencie{r: r}
	return &m
}

func (m *Currencie) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableCurrencyProps](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.CurrencyService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *Currencie) Update(context fiber.Ctx) error {
	model, code, err := api.BindWithString[types.UpdateCurrencyInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.CurrencyService().SetContext(context.Context()).Update(code, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}
