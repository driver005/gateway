package interfaces

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type CartCompletionResponse struct {
	ResponseCode int
	ResponseBody map[string]interface{}
}

type ICartCompletionStrategy interface {
	Complete(cartId uuid.UUID, idempotencyKey *models.IdempotencyKey, context types.RequestContext) (*CartCompletionResponse, *utils.ApplictaionError)
}
