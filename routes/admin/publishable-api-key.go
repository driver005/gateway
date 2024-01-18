package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type PublishableApiKey struct {
	r Registry
}

func NewPublishableApiKey(r Registry) *PublishableApiKey {
	m := PublishableApiKey{r: r}
	return &m
}

func (m *PublishableApiKey) Get(context fiber.Ctx) error {
	return nil
}

func (m *PublishableApiKey) List(context fiber.Ctx) error {
	return nil
}

func (m *PublishableApiKey) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreatePublishableApiKeyInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.PublishableApiKeyService().SetContext(context.Context()).Create(model, uuid.Nil)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *PublishableApiKey) Update(context fiber.Ctx) error {
	return nil
}

func (m *PublishableApiKey) Delete(context fiber.Ctx) error {
	return nil
}

func (m *PublishableApiKey) AddChannelsBatch(context fiber.Ctx) error {
	return nil
}

func (m *PublishableApiKey) DeleteChannelsBatch(context fiber.Ctx) error {
	return nil
}

func (m *PublishableApiKey) ListChannels(context fiber.Ctx) error {
	return nil
}

func (m *PublishableApiKey) Revoke(context fiber.Ctx) error {
	return nil
}
