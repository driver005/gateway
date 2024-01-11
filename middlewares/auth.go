package middlewares

import (
	"encoding/base64"
	"strings"
	"unsafe"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/keyauth"
)

func (h *Handler) Authenticate() []fiber.Handler {
	return []fiber.Handler{
		h.AdminBearer(),
		h.AdminApiTocken(),
	}
}

func (h *Handler) AdminBasic() fiber.Handler {
	return keyauth.New(keyauth.Config{
		KeyLookup:  "header:Authorization",
		AuthScheme: "Basic",
		Validator: func(c fiber.Ctx, key string) (bool, error) {
			// Decode the header contents
			raw, err := base64.StdEncoding.DecodeString(key)
			if err != nil {
				return false, keyauth.ErrMissingOrMalformedAPIKey
			}

			// Get the credentials
			creds := unsafe.String(unsafe.SliceData(raw), len(raw))

			// Check if the credentials are in the correct form
			// which is "email:password".
			index := strings.Index(creds, ":")
			if index == -1 {
				return false, keyauth.ErrMissingOrMalformedAPIKey
			}

			result := h.r.AuthService().SetContext(c.Context()).Authenticate(creds[:index], creds[index+1:])
			if result.Success {
				c.Locals("user", result.User)
				return true, nil
			}

			return false, keyauth.ErrMissingOrMalformedAPIKey
		},
	})
}

func (h *Handler) AdminApiTocken() fiber.Handler {
	return keyauth.New(keyauth.Config{
		KeyLookup: "header:x-medusa-access-token",
		Validator: func(c fiber.Ctx, key string) (bool, error) {
			result := h.r.AuthService().SetContext(c.Context()).AuthenticateAPIToken(key)
			if result.Success {
				c.Locals("user", result.User)
				return true, nil
			}

			return false, keyauth.ErrMissingOrMalformedAPIKey
		},
	})
}

func (h *Handler) AdminBearer() fiber.Handler {
	return keyauth.New(keyauth.Config{
		KeyLookup:  "header:Authorization",
		AuthScheme: "Bearer",
		Validator: func(c fiber.Ctx, key string) (bool, error) {
			_, data, err := h.r.TockenService().SetContext(c.Context()).VerifyToken(key)
			if err != nil {
				return false, keyauth.ErrMissingOrMalformedAPIKey
			}

			if data["domain"] != "admin" {
				return false, keyauth.ErrMissingOrMalformedAPIKey
			}

			userId, ok := data["user_id"]
			if !ok {
				return false, keyauth.ErrMissingOrMalformedAPIKey
			}

			c.Locals("user_id", userId)

			return true, nil
		},
	})
}

func (h *Handler) StoreBearer() fiber.Handler {
	return keyauth.New(keyauth.Config{
		KeyLookup:  "header:Authorization",
		AuthScheme: "Bearer",
		Validator: func(c fiber.Ctx, key string) (bool, error) {
			_, data, err := h.r.TockenService().SetContext(c.Context()).VerifyToken(key)
			if err != nil {
				return false, keyauth.ErrMissingOrMalformedAPIKey
			}

			if data["domain"] != "store" {
				return false, keyauth.ErrMissingOrMalformedAPIKey
			}

			customerId, ok := data["customer_id"]
			if !ok {
				return false, keyauth.ErrMissingOrMalformedAPIKey
			}

			c.Locals("customer_id", customerId)

			return true, nil
		},
	})
}
