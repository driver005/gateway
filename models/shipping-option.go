package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Shipping Options represent a way in which an Order or Return can be shipped. Shipping Options have an associated Fulfillment Provider that will be used when the fulfillment of an Order is initiated. Shipping Options themselves cannot be added to Carts, but serve as a template for Shipping Methods. This distinction makes it possible to customize individual Shipping Methods with additional information.
type ShippingOption struct {
	core.Model

	// The amount to charge for shipping when the Shipping Option price type is `flat_rate`.
	Amount float64 `json:"amount" gorm:"default:null"`

	// The data needed for the Fulfillment Provider to identify the Shipping Option.
	Data core.JSONB `json:"data" gorm:"default:null"`

	// [EXPERIMENTAL] Does the shipping option price include tax
	IncludesTax bool `json:"includes_tax" gorm:"default:null"`

	// Flag to indicate if the Shipping Option can be used for Return shipments.
	IsReturn bool `json:"is_return" gorm:"default:null"`

	AdminOnly bool `json:"admin_only" gorm:"default:null"`

	// The name given to the Shipping Option - this may be displayed to the Customer.
	Name string `json:"name"`

	// The type of pricing calculation that is used when creatin Shipping Methods from the Shipping Option. Can be `flat_rate` for fixed prices or `calculated` if the Fulfillment Provider can provide price calulations.
	PriceType ShippingOptionPriceType `json:"price_type"`

	// Shipping Profiles have a set of defined Shipping Options that can be used to fulfill a given set of Products.
	Profile *ShippingProfile `json:"profile" gorm:"foreignKey:id;references:profile_id"`

	// The ID of the Shipping Profile that the shipping option belongs to. Shipping Profiles have a set of defined Shipping Options that can be used to Fulfill a given set of Products.
	ProfileId uuid.NullUUID `json:"profile_id"`

	// Represents a fulfillment provider plugin and holds its installation status.
	Provider *FulfillmentProvider `json:"provider" gorm:"foreignKey:id;references:provider_id"`

	// The id of the Fulfillment Provider, that will be used to process Fulfillments from the Shipping Option.
	ProviderId uuid.NullUUID `json:"provider_id"`

	// A region object. Available if the relation `region` is expanded.
	Region *Region `json:"region" gorm:"foreignKey:id;references:region_id"`

	// The region's ID
	RegionId uuid.NullUUID `json:"region_id"`

	// The requirements that must be satisfied for the Shipping Option to be available for a Cart. Available if the relation `requirements` is expanded.
	Requirements []ShippingOptionRequirement `json:"requirements" gorm:"foreignKey:id"`
}

// The type of pricing calculation that is used when creatin Shipping Methods from the Shipping Option. Can be `flat_rate` for fixed prices or `calculated` if the Fulfillment Provider can provide price calulations.
type ShippingOptionPriceType string

// Defines values for ShippingOptionPriceType.
const (
	ShippingOptionPriceCalculated ShippingOptionPriceType = "calculated"
	ShippingOptionPriceFlatRate   ShippingOptionPriceType = "flat_rate"
)

func (so *ShippingOptionPriceType) Scan(value interface{}) error {
	*so = ShippingOptionPriceType(value.([]byte))
	return nil
}

func (so ShippingOptionPriceType) Value() (driver.Value, error) {
	return string(so), nil
}
