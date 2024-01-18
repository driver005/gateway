package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type GiftCard struct {
	r Registry
}

func NewGiftCard(r Registry) *GiftCard {
	m := GiftCard{r: r}
	return &m
}

func (m *GiftCard) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.GiftCardService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *GiftCard) List(context fiber.Ctx) error {
	return nil
}

func (m *GiftCard) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateGiftCardInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.GiftCardService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *GiftCard) Update(context fiber.Ctx) error {
	return nil
}

func (m *GiftCard) Delete(context fiber.Ctx) error {
	return nil
}
