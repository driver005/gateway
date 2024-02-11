package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type CreatePaymentCollectionInput struct {
	RegionId     uuid.UUID                    `json:"region_id"`
	Type         models.PaymentCollectionType `json:"type"`
	CurrencyCode string                       `json:"currency_code"`
	Amount       float64                      `json:"amount"`
	CreatedBy    uuid.UUID                    `json:"created_by"`
	Metadata     core.JSONB                   `json:"metadata,omitempty" validate:"omitempty"`
	Description  string                       `json:"description,omitempty" validate:"omitempty"`
}

type PaymentCollectionsSessionsBatchInput struct {
	ProviderId uuid.UUID `json:"provider_id"`
	Amount     float64   `json:"amount"`
	SessionId  uuid.UUID `json:"session_id,omitempty" validate:"omitempty"`
}

type PaymentCollectionsSessionsInput struct {
	ProviderId uuid.UUID `json:"provider_id"`
}
