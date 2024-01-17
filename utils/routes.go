package utils

import (
	"github.com/driver005/gateway/models"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func GetUser(context fiber.Ctx) uuid.UUID {
	var id uuid.UUID
	user, ok := context.Locals("user").(*models.User)
	if !ok {
		id = context.Locals("user_id").(uuid.UUID)
	} else {
		id = user.Id
	}

	return id
}
