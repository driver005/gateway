package types

import "github.com/google/uuid"

// The context to apply on retrieved prices.
type PriceSelectionParams struct {
	FindParams
	// Retrieve prices for a cart Id uuid.UUID.
	CartId uuid.UUID `json:"cart_id,omitempty" validate:"omitempty"`
	// Retrieve prices for a region Id uuid.UUID.
	RegionId uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	// Retrieve prices for a currency code.
	CurrencyCode string `json:"currency_code,omitempty" validate:"omitempty"`
}

// The context to apply on retrieved prices by a user admin.
type AdminPriceSelectionParams struct {
	PriceSelectionParams
	// Retrieve prices for a customer Id uuid.UUID.
	CustomerId uuid.UUID `json:"customer_id,omitempty" validate:"omitempty"`
}
