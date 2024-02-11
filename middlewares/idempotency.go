package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
)

func (h *Handler) Idempotency() fiber.Handler {
	return idempotency.New(idempotency.Config{
		// LifeTime is the backtrack time window of the presence of an idempotency key. Default is 24h.
		Lifetime: time.Hour * 24,

		// KeyHeader is the header for looking up the idempotency key from. Default is "X-Idempotency-Key".
		KeyHeader: "X-Idempotency-Key",

		// Storage is the storage of idempotency responses. Default is a in-memory fiber.Storage.
		// Storage: fiberstore.NewRedis(), // this is to just illustrate a storage that implements fiber.Storage
	})
}
