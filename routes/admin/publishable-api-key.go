package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
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

func (m *PublishableApiKey) SetRoutes(router fiber.Router) {
	route := router.Group("/publishable-api-keys")
	route.Get("/:id", m.Get)
	route.Get("/", m.List)
	route.Post("/", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Post("/:id/revoke", m.Revoke)
	route.Get("/:id/sales-channels", m.ListChannels)
	route.Post("/:id/sales-channels/batch", m.AddChannelsBatch)
	route.Delete("/:id/sales-channels/batch", m.DeleteChannelsBatch)
}

func (m *PublishableApiKey) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.PublishableApiKeyService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *PublishableApiKey) List(context fiber.Ctx) error {
	model, config, err := api.BindList[models.PublishableApiKey](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.PublishableApiKeyService().SetContext(context.Context()).ListAndCount(model, config)
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
	model, id, err := api.BindUpdate[types.UpdatePublishableApiKeyInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.PublishableApiKeyService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *PublishableApiKey) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.PublishableApiKeyService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "publishable-api-key",
		"deleted": true,
	})
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
