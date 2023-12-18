package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
)

// Shipping Profiles have a set of defined Shipping Options that can be used to fulfill a given set of Products.
type ShippingProfile struct {
	core.Model

	// The name given to the Shipping profile - this may be displayed to the Customer.
	Name string `json:"name"`

	// The Products that the Shipping Profile defines Shipping Options for. Available if the relation `products` is expanded.
	Products []Product `json:"products" gorm:"foreignKey:id"`

	// The Shipping Options that can be used to fulfill the Products in the Shipping Profile. Available if the relation `shipping_options` is expanded.
	ShippingOptions []ShippingOption `json:"shipping_options" gorm:"foreignKey:id"`

	// The type of the Shipping Profile, may be `default`, `gift_card` or `custom`.
	Type ShippingProfileType `json:"type"`
}

// The type of the Shipping Profile, may be `default`, `gift_card` or `custom`.
type ShippingProfileType string

// Defines values for ShippingProfileType.
const (
	ShippingProfileTypeCustom   ShippingProfileType = "custom"
	ShippingProfileTypeDefault  ShippingProfileType = "default"
	ShippingProfileTypeGiftCard ShippingProfileType = "gift_card"
)

func (sp *ShippingProfileType) Scan(value interface{}) error {
	*sp = ShippingProfileType(value.([]byte))
	return nil
}

func (sp ShippingProfileType) Value() (driver.Value, error) {
	return string(sp), nil
}
