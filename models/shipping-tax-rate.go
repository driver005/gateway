package models

import "github.com/google/uuid"

// ShippingTaxRate - Associates a tax rate with a shipping option to indicate that the shipping option is taxed in a certain way
type ShippingTaxRate struct {

	// The ID of the Shipping Option
	ShippingOptionId uuid.NullUUID `json:"shipping_option_id"`

	ShippingOption *ShippingOption `json:"shipping_option" gorm:"foreignKey:id;references:shipping_option_id"`

	// The ID of the Tax Rate
	RateId uuid.NullUUID `json:"rate_id"`

	TaxRate *TaxRate `json:"tax_rate" gorm:"foreignKey:id;references:rate_id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" gorm:"default:null"`
}
