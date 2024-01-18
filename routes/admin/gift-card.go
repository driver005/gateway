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
	model, config, err := api.BindList[types.FilterableGiftCard](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.GiftCardService().SetContext(context.Context()).ListAndCount(model, config)
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
	model, id, err := api.BindUpdate[types.UpdateGiftCardInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.GiftCardService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *GiftCard) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.GiftCardService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "gift-card",
		"deleted": true,
	})
}
