package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:LineItemAdjustment
// title: "Line Item Adjustment"
// description: "A Line Item Adjustment includes details on discounts applied on a line item."
// type: object
// required:
//   - amount
//   - description
//   - discount_id
//   - id
//   - item_id
//   - metadata
//
// properties:
//
//	id:
//	  description: The Line Item Adjustment's ID
//	  type: string
//	  example: lia_01G8TKE4XYCTHSCK2GDEP47RE1
//	item_id:
//	  description: The ID of the line item
//	  type: string
//	  example: item_01G8ZC9GWT6B2GP5FSXRXNFNGN
//	item:
//	  description: The details of the line item.
//	  x-expandable: "item"
//	  nullable: true
//	  $ref: "#/components/schemas/LineItem"
//	description:
//	  description: The line item's adjustment description
//	  type: string
//	  example: Adjusted item's price.
//	discount_id:
//	  description: The ID of the discount associated with the adjustment
//	  nullable: true
//	  type: string
//	  example: disc_01F0YESMW10MGHWJKZSDDMN0VN
//	discount:
//	  description: The details of the discount associated with the adjustment.
//	  x-expandable: "discount"
//	  nullable: true
//	  $ref: "#/components/schemas/Discount"
//	amount:
//	  description: The adjustment amount
//	  type: number
//	  example: 1000
//	metadata:
//	  description: An optional key-value map with additional details
//	  nullable: true
//	  type: object
//	  example: {car: "white"}
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type LineItemAdjustment struct {
	core.Model

	ItemId      uuid.NullUUID `json:"item_id"  gorm:"column:item_id"`
	Item        *LineItem     `json:"item"  gorm:"column:item;foreignKey:ItemId"`
	Description string        `json:"description"  gorm:"column:description"`
	DiscountId  uuid.NullUUID `json:"discount_id"  gorm:"column:discount_id"`
	Discount    *Discount     `json:"discount"  gorm:"column:discount;foreignKey:DiscountId"`
	Amount      float64       `json:"amount"  gorm:"column:amount"`
}
