package middlewares

import (
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *Handler) PublishableApiKey(context fiber.Ctx) error {
	pubKey, err := utils.ParseUUID(context.Get("x-publishable-api-key"))
	if err != nil {
		return err
	}

	if pubKey != uuid.Nil {
		publishableApiKeyScopes, err := h.r.PublishableApiKeyService().GetResourceScopes(pubKey)
		if err != nil {
			return err
		}

		context.Locals("publishableApiKeyScopes", publishableApiKeyScopes)
	}

	return context.Next()
}
