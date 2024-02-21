package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
)

// @oas:schema:ShippingProfile
// title: "Shipping Profile"
// description: "A Shipping Profile has a set of defined Shipping Options that can be used to fulfill a given set of Products. For example, gift cards are shipped differently than physical products,
//
//	so a shipping profile with the type `gift_card` groups together the shipping options that can only be used for gift cards."
//
// type: object
// required:
//   - created_at
//   - deleted_at
//   - id
//   - metadata
//   - name
//   - type
//   - updated_at
//
// properties:
//
//	id:
//	  description: The shipping profile's ID
//	  type: string
//	  example: sp_01G1G5V239ENSZ5MV4JAR737BM
//	name:
//	  description: The name given to the Shipping profile - this may be displayed to the Customer.
//	  type: string
//	  example: Default Shipping Profile
//	type:
//	  description: The type of the Shipping Profile, may be `default`, `gift_card` or `custom`.
//	  type: string
//	  enum:
//	    - default
//	    - gift_card
//	    - custom
//	  example: default
//	products:
//	  description: The details of the products that the Shipping Profile defines Shipping Options for. Available if the relation `products` is expanded.
//	  type: array
//	  x-expandable: "products"
//	  items:
//	    $ref: "#/components/schemas/Product"
//	shipping_options:
//	  description: The details of the shipping options that can be used to create shipping methods for the Products in the Shipping Profile.
//	  type: array
//	  x-expandable: "shipping_options"
//	  items:
//	    $ref: "#/components/schemas/ShippingOption"
//	created_at:
//	  description: The date with timezone at which the resource was created.
//	  type: string
//	  format: date-time
//	updated_at:
//	  description: The date with timezone at which the resource was updated.
//	  type: string
//	  format: date-time
//	deleted_at:
//	  description: The date with timezone at which the resource was deleted.
//	  nullable: true
//	  type: string
//	  format: date-time
//	metadata:
//	  description: An optional key-value map with additional details
//	  nullable: true
//	  type: object
//	  example: {car: "white"}
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type ShippingProfile struct {
	core.SoftDeletableModel

	Name            string              `json:"name" gorm:"column:name"`
	Products        []Product           `json:"products" gorm:"many2many:product_shipping_profile"`
	ShippingOptions []ShippingOption    `json:"shipping_options" gorm:"foreignKey:Id"`
	Type            ShippingProfileType `json:"type" gorm:"column:type"`
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
