package admin

import "github.com/gofiber/fiber/v3"

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
	return nil
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
