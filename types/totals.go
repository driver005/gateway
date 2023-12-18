package types

import (
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

// CalculationContextData represents the data required for performing calculations
type CalculationContextData struct {
	Discounts       []models.Discount       `json:"discounts,omitempty"`
	Items           []models.LineItem       `json:"items,omitempty"`
	Customer        models.Customer         `json:"customer"`
	Region          models.Region           `json:"region"`
	ShippingAddress *models.Address         `json:"shipping_address,omitempty"`
	Swaps           []models.Swap           `json:"swaps,omitempty"`
	Claims          []models.ClaimOrder     `json:"claims,omitempty"`
	ShippingMethods []models.ShippingMethod `json:"shipping_methods,omitempty"`
}

// GiftCardAllocation represents the amount of a gift card allocated to a line item
type GiftCardAllocation struct {
	Amount     float64 `json:"amount,omitempty"`
	UnitAmount float64 `json:"unit_amount,omitempty"`
}

// DiscountAllocation represents the amount of a discount allocated to a line item
type DiscountAllocation struct {
	Amount     float64 `json:"amount,omitempty"`
	UnitAmount float64 `json:"unit_amount,omitempty"`
}

// LineAllocationsMap represents a map of line item ids and its corresponding gift card and discount allocations
type LineAllocations struct {
	GiftCard *GiftCardAllocation `json:"gift_card,omitempty"`
	Discount *DiscountAllocation `json:"discount,omitempty"`
}

// LineAllocationsMap represents a map of line item ids and its corresponding gift card and discount allocations
type LineAllocationsMap map[uuid.UUID]struct {
	GiftCard *GiftCardAllocation `json:"gift_card,omitempty"`
	Discount *DiscountAllocation `json:"discount,omitempty"`
}

// SubtotalOptions represents options to use for subtotal calculations
type SubtotalOptions struct {
	ExcludeNonDiscounts bool `json:"excludeNonDiscounts,omitempty"`
}

// LineDiscount associates a line item and discount allocation
type LineDiscount struct {
	LineItem models.LineItem `json:"lineItem,omitempty"`
	Variant  uuid.UUID       `json:"variant,omitempty"`
	Amount   float64         `json:"amount,omitempty"`
}

// LineDiscountAmount associates a line item and discount allocation
type LineDiscountAmount struct {
	Item                    models.LineItem `json:"item,omitempty"`
	Amount                  float64         `json:"amount,omitempty"`
	CustomAdjustmentsAmount float64         `json:"customAdjustmentsAmount,omitempty"`
}
