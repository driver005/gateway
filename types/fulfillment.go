package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type FulFillmentItemType struct {
	ItemId   uuid.UUID `json:"item_id"`
	Quantity int       `json:"quantity"`
}

type FulfillmentItemPartition struct {
	ShippingMethod *models.ShippingMethod `json:"shipping_method"`
	Items          []models.LineItem      `json:"items"`
}

type CreateShipmentConfig struct {
	Metadata       core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
	NoNotification bool       `json:"no_notification,omitempty" validate:"omitempty"`
	LocationId     uuid.UUID  `json:"location_id,omitempty" validate:"omitempty"`
}

type CreateFulfillmentOrder struct {
	*models.ClaimOrder
	IsClaim         bool                    `json:"is_claim,omitempty" validate:"omitempty"`
	Email           string                  `json:"email,omitempty" validate:"omitempty"`
	Payments        []models.Payment        `json:"payments"`
	Discounts       []models.Discount       `json:"discounts"`
	CurrencyCode    string                  `json:"currency_code"`
	TaxRate         float64                 `json:"tax_rate,omitempty" validate:"omitempty"`
	RegionId        uuid.UUID               `json:"region_id"`
	Region          *models.Region          `json:"region,omitempty" validate:"omitempty"`
	IsSwap          bool                    `json:"is_swap,omitempty" validate:"omitempty"`
	DisplayId       int                     `json:"display_id"`
	BillingAddress  *models.Address         `json:"billing_address"`
	Items           []models.LineItem       `json:"items"`
	ShippingMethods []models.ShippingMethod `json:"shipping_methods"`
	NoNotification  bool                    `json:"no_notification"`
}
