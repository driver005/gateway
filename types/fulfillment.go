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
	Metadata       core.JSONB `json:"metadata,omitempty"`
	NoNotification bool       `json:"no_notification,omitempty"`
	LocationID     string     `json:"location_id,omitempty"`
}

type CreateFulfillmentOrder struct {
	models.ClaimOrder
	IsClaim         bool                    `json:"is_claim,omitempty"`
	Email           string                  `json:"email,omitempty"`
	Payments        []models.Payment        `json:"payments"`
	Discounts       []models.Discount       `json:"discounts"`
	CurrencyCode    string                  `json:"currency_code"`
	TaxRate         float64                 `json:"tax_rate,omitempty"`
	RegionId        uuid.UUID               `json:"region_id"`
	Region          *models.Region          `json:"region,omitempty"`
	IsSwap          bool                    `json:"is_swap,omitempty"`
	DisplayId       string                  `json:"display_id"`
	BillingAddress  *models.Address         `json:"billing_address"`
	Items           []models.LineItem       `json:"items"`
	ShippingMethods []models.ShippingMethod `json:"shipping_methods"`
	NoNotification  bool                    `json:"no_notification"`
}
