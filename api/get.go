package api

import (
	"github.com/driver005/gateway/models"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func GetUser(context fiber.Ctx) uuid.UUID {
	var id uuid.UUID
	user, ok := context.Locals("user").(*models.User)
	if !ok {
		if context.Locals("user_id").(uuid.UUID) != uuid.Nil {
			id = context.Locals("user_id").(uuid.UUID)
		} else {
			id = uuid.Nil
		}
	} else {
		id = user.Id
	}

	return id
}

func GetUserStore(context fiber.Ctx) uuid.UUID {
	userId := GetUser(context)
	if userId == uuid.Nil {
		userId = context.Locals("customer_id").(uuid.UUID)
	}

	return userId
}
