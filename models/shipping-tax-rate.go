package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:ShippingTaxRate
// title: "Shipping Tax Rate"
// description: "This represents the tax rates applied on a shipping option."
// type: object
// required:
//   - created_at
//   - metadata
//   - rate_id
//   - shipping_option_id
//   - updated_at
//
// properties:
//
//	shipping_option_id:
//	  description: The ID of the shipping option.
//	  type: string
//	  example: so_01G1G5V27GYX4QXNARRQCW1N8T
//	shipping_option:
//	  description: The details of the shipping option.
//	  x-expandable: "shipping_option"
//	  nullable: true
//	  $ref: "#/components/schemas/ShippingOption"
//	rate_id:
//	  description: The ID of the associated tax rate.
//	  type: string
//	  example: txr_01G8XDBAWKBHHJRKH0AV02KXBR
//	tax_rate:
//	  description: The details of the associated tax rate.
//	  x-expandable: "tax_rate"
//	  nullable: true
//	  $ref: "#/components/schemas/TaxRate"
//	created_at:
//	  description: The date with timezone at which the resource was created.
//	  type: string
//	  format: date-time
//	updated_at:
//	  description: The date with timezone at which the resource was updated.
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
type ShippingTaxRate struct {
	ShippingOptionId uuid.NullUUID   `json:"shipping_option_id" gorm:"column:shipping_option_id;primaryKey"`
	RateId           uuid.NullUUID   `json:"rate_id" gorm:"column:rate_id;primaryKey"`
	ShippingOption   *ShippingOption `json:"shipping_option" gorm:"foreignKey:ShippingOptionId"`
	TaxRate          *TaxRate        `json:"tax_rate" gorm:"foreignKey:RateId"`
	CreatedAt        time.Time       `json:"created_at" gorm:"column:created_at;created_at"`
	UpdatedAt        time.Time       `json:"updated_at" gorm:"column:updated_at;updated_at"`
	Metadata         core.JSONB      `json:"metadata" gorm:"column:metadata"`
}
