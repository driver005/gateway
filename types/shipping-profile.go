package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type CreateShippingProfile struct {
	Name     string                     `json:"name"`
	Type     models.ShippingProfileType `json:"type"`
	Metadata core.JSONB                 `json:"metadata,omitempty" validate:"omitempty"`
}

type UpdateShippingProfile struct {
	Name            string                     `json:"name,omitempty" validate:"omitempty"`
	Metadata        core.JSONB                 `json:"metadata,omitempty" validate:"omitempty"`
	Type            models.ShippingProfileType `json:"type,omitempty" validate:"omitempty"`
	Products        uuid.UUIDs                 `json:"products,omitempty" validate:"omitempty"`
	ShippingOptions uuid.UUIDs                 `json:"shipping_options,omitempty" validate:"omitempty"`
}
