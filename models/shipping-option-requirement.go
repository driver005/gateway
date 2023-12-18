package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// A requirement that a Cart must satisfy for the Shipping Option to be available to the Cart.
type ShippingOptionRequirement struct {
	core.Model

	// The amount to compare the Cart subtotal to.
	Amount float64 `json:"amount"`

	// Shipping Options represent a way in which an Order or Return can be shipped. Shipping Options have an associated Fulfillment Provider that will be used when the fulfillment of an Order is initiated. Shipping Options themselves cannot be added to Carts, but serve as a template for Shipping Methods. This distinction makes it possible to customize individual Shipping Methods with additional information.
	ShippingOption *ShippingOption `json:"shipping_option" gorm:"foreignKey:id;references:shipping_option_id"`

	// The id of the Shipping Option that the hipping option requirement belongs to
	ShippingOptionId uuid.NullUUID `json:"shipping_option_id"`

	// The type of the requirement, this defines how the value will be compared to the Cart's total. `min_subtotal` requirements define the minimum subtotal that is needed for the Shipping Option to be available, while the `max_subtotal` defines the maximum subtotal that the Cart can have for the Shipping Option to be available.
	Type ShippingOptionRequirementType `json:"type"`
}

// The type of the requirement, this defines how the value will be compared to the Cart's total. `min_subtotal` requirements define the minimum subtotal that is needed for the Shipping Option to be available, while the `max_subtotal` defines the maximum subtotal that the Cart can have for the Shipping Option to be available.
type ShippingOptionRequirementType string

// Defines values for ShippingOptionRequirementType.
const (
	MaxSubtotal ShippingOptionRequirementType = "max_subtotal"
	MinSubtotal ShippingOptionRequirementType = "min_subtotal"
)

func (so *ShippingOptionRequirementType) Scan(value interface{}) error {
	*so = ShippingOptionRequirementType(value.([]byte))
	return nil
}

func (so ShippingOptionRequirementType) Value() (driver.Value, error) {
	return string(so), nil
}
