package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:DiscountConditionProductType
// title: "Product Type Discount Condition"
// description: "This represents the association between a discount condition and a product type"
// type: object
// required:
//   - condition_id
//   - created_at
//   - metadata
//   - product_type_id
//   - updated_at
//
// properties:
//
//	product_type_id:
//	  description: The ID of the Product Tag
//	  type: string
//	  example: ptyp_01G8X9A7ESKAJXG2H0E6F1MW7A
//	condition_id:
//	  description: The ID of the Discount Condition
//	  type: string
//	  example: discon_01G8X9A7ESKAJXG2H0E6F1MW7A
//	product_type:
//	  description: The details of the product type.
//	  x-expandable: "product_type"
//	  nullable: true
//	  $ref: "#/components/schemas/ProductType"
//	discount_condition:
//	  description: The details of the discount condition.
//	  x-expandable: "discount_condition"
//	  nullable: true
//	  $ref: "#/components/schemas/DiscountCondition"
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
type DiscountConditionProductType struct {
	ProductTypeId     uuid.NullUUID      `json:"product_type_id"  gorm:"column:product_type_id;primaryKey"`
	ConditionId       uuid.NullUUID      `json:"condition_id"  gorm:"column:condition_id;primaryKey"`
	ProductType       *ProductType       `json:"product_type"  gorm:"column:product_type;foreignKey:ProductTypeId"`
	DiscountCondition *DiscountCondition `json:"discount_condition"  gorm:"column:discount_condition;foreignKey:ConditionId"`
	CreatedAt         time.Time          `json:"created_at"  gorm:"column:created_at;created_at"`
	UpdatedAt         time.Time          `json:"updated_at"  gorm:"column:updated_at;updated_at"`
	Metadata          core.JSONB         `json:"metadata"  gorm:"column:metadata"`
}
