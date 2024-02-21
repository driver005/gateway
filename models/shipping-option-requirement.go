package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:ShippingOptionRequirement
// title: "Shipping Option Requirement"
// description: "A shipping option requirement defines conditions that a Cart must satisfy for the Shipping Option to be available for usage in the Cart."
// type: object
// required:
//   - amount
//   - deleted_at
//   - id
//   - shipping_option_id
//   - type
//
// properties:
//
//	id:
//	  description: The shipping option requirement's ID
//	  type: string
//	  example: sor_01G1G5V29AB4CTNDRFSRWSRKWD
//	shipping_option_id:
//	  description: The ID of the shipping option that the requirements belong to.
//	  type: string
//	  example: so_01G1G5V27GYX4QXNARRQCW1N8T
//	shipping_option:
//	  description: The details of the shipping option that the requirements belong to.
//	  x-expandable: "shipping_option"
//	  nullable: true
//	  $ref: "#/components/schemas/ShippingOption"
//	type:
//	  description: The type of the requirement, this defines how the value will be compared to the Cart's total. `min_subtotal` requirements define the minimum subtotal that is needed for the Shipping Option to be available, while the `max_subtotal` defines the maximum subtotal that the Cart can have for the Shipping Option to be available.
//	  type: string
//	  enum:
//	    - min_subtotal
//	    - max_subtotal
//	  example: min_subtotal
//	amount:
//	  description: The amount to compare the Cart subtotal to.
//	  type: integer
//	  example: 100
//	deleted_at:
//	  description: The date with timezone at which the resource was deleted.
//	  nullable: true
//	  type: string
//	  format: date-time
type ShippingOptionRequirement struct {
	core.Model

	Amount           float64                       `json:"amount" gorm:"column:amount"`
	ShippingOption   *ShippingOption               `json:"shipping_option" gorm:"foreignKey:ShippingOptionId"`
	ShippingOptionId uuid.NullUUID                 `json:"shipping_option_id" gorm:"column:shipping_option_id"`
	Type             ShippingOptionRequirementType `json:"type" gorm:"column:type"`
}

// The type of the requirement, this defines how the value will be compared to the Cart's total. `min_subtotal` requirements define the minimum subtotal that is needed for the Shipping Option to be available, while the `max_subtotal` defines the maximum subtotal that the Cart can have for the Shipping Option to be available.
type ShippingOptionRequirementType string

// Defines values for ShippingOptionRequirementType.
const (
	ShippingOptionRequirementMaxSubtotal ShippingOptionRequirementType = "max_subtotal"
	ShippingOptionRequirementMinSubtotal ShippingOptionRequirementType = "min_subtotal"
)

func (so *ShippingOptionRequirementType) Scan(value interface{}) error {
	*so = ShippingOptionRequirementType(value.([]byte))
	return nil
}

func (so ShippingOptionRequirementType) Value() (driver.Value, error) {
	return string(so), nil
}
