package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type UpdatePaymentInput struct {
	OrderId uuid.UUID `json:"order_id,omitempty" validate:"omitempty"`
	SwapId  uuid.UUID `json:"swap_id,omitempty" validate:"omitempty"`
}

type PaymentMethod struct {
	ProviderId uuid.UUID  `json:"provider_id,omitempty" validate:"omitempty"`
	Data       core.JSONB `json:"data,omitempty" validate:"omitempty"`
}

type PaymentSessionInput struct {
	PaymentSessionId   uuid.UUID        `json:"payment_session_id,omitempty" validate:"omitempty"`
	ProviderId         uuid.UUID        `json:"provider_id"`
	Cart               *models.Cart     `json:"cart"`
	Customer           *models.Customer `json:"customer,omitempty" validate:"omitempty"`
	CurrencyCode       string           `json:"currency_code"`
	Amount             float64          `json:"amount"`
	ResourceId         uuid.UUID        `json:"resource_id,omitempty" validate:"omitempty"`
	PaymentSessionData core.JSONB       `json:"paymentSessionData,omitempty" validate:"omitempty"`
	Context            core.JSONB
}

type CreatePaymentInput struct {
	CartId         uuid.UUID              `json:"cart_id,omitempty" validate:"omitempty"`
	Amount         float64                `json:"amount"`
	CurrencyCode   string                 `json:"currency_code"`
	ProviderId     uuid.UUID              `json:"provider_id,omitempty" validate:"omitempty"`
	PaymentSession *models.PaymentSession `json:"payment_session"`
	ResourceId     uuid.UUID              `json:"resource_id,omitempty" validate:"omitempty"`
}
