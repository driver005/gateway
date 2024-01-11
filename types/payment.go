package types

import (
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type PaymentSessionInput struct {
	Cart               models.Cart
	PaymentSessionId   uuid.UUID
	ProviderId         uuid.UUID
	Customer           *models.Customer
	CurrencyCode       string
	Amount             float64
	ResourceId         uuid.UUID
	PaymentSessionData map[string]interface{}
	Context            map[string]interface{}
}
