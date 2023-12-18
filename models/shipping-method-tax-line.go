package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// ShippingMethodTaxLine - Shipping Method Tax Line
type ShippingMethodTaxLine struct {
	core.Model

	// The ID of the line item
	ShippingMethodId uuid.NullUUID `json:"shipping_method_id"`

	ShippingMethod *ShippingMethod `json:"shipping_method" gorm:"foreignKey:id;references:shipping_method_id"`

	// A code to identify the tax type by
	Code string `json:"code" gorm:"default:null"`

	// A human friendly name for the tax
	Name string `json:"name"`

	// The numeric rate to charge tax by
	Rate float64 `json:"rate"`
}
