package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type CreateShippingProfile struct {
	Name     string                     `json:"name"`
	Type     models.ShippingProfileType `json:"type"`
	Metadata core.JSONB                 `json:"metadata,omitempty"`
}

type UpdateShippingProfile struct {
	Name            string                     `json:"name,omitempty"`
	Metadata        core.JSONB                 `json:"metadata,omitempty"`
	Type            models.ShippingProfileType `json:"type,omitempty"`
	Products        uuid.UUIDs                 `json:"products,omitempty"`
	ShippingOptions uuid.UUIDs                 `json:"shipping_options,omitempty"`
}
