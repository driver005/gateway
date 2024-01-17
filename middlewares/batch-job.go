package middlewares

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *Handler) CanAccessBatchJob(context fiber.Ctx) error {
	Id, err := utils.ParseUUID(context.Params("user_id"))
	if err != nil {
		return err
	}

	batch, err := h.r.BatchJobService().Retrive(Id)
	if err != nil {
		return err
	}

	var userId uuid.UUID
	user, ok := context.Locals("user").(*models.User)
	if !ok {
		userId = context.Locals("user_id").(uuid.UUID)
	} else {
		userId = user.Id
	}

	if batch.CreatedBy != userId {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Cannot access a batch job that does not belong to the logged in user",
			nil,
		)
	}

	context.Locals("batch-job", batch)

	return context.Next()
}
