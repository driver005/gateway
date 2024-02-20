package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:ReturnReason
// title: "Return Reason"
// description: "A Return Reason is a value defined by an admin. It can be used on Return Items in order to indicate why a Line Item was returned."
// type: object
// required:
//   - created_at
//   - deleted_at
//   - description
//   - id
//   - label
//   - metadata
//   - parent_return_reason_id
//   - updated_at
//   - value
//
// properties:
//
//	id:
//	  description: The return reason's ID
//	  type: string
//	  example: rr_01G8X82GCCV2KSQHDBHSSAH5TQ
//	value:
//	  description: The value to identify the reason by.
//	  type: string
//	  example: damaged
//	label:
//	  description: A text that can be displayed to the Customer as a reason.
//	  type: string
//	  example: Damaged goods
//	description:
//	  description: A description of the Reason.
//	  nullable: true
//	  type: string
//	  example: Items that are damaged
//	parent_return_reason_id:
//	  description: The ID of the parent reason.
//	  nullable: true
//	  type: string
//	  example: null
//	parent_return_reason:
//	  description: The details of the parent reason.
//	  x-expandable: "parent_return_reason"
//	  nullable: true
//	  $ref: "#/components/schemas/ReturnReason"
//	return_reason_children:
//	  description: The details of the child reasons.
//	  x-expandable: "return_reason_children"
//	  $ref: "#/components/schemas/ReturnReason"
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
type ReturnReason struct {
	core.Model

	Description          string         `json:"description"  gorm:"column:description"`
	Label                string         `json:"label"  gorm:"column:label"`
	Value                string         `json:"value"  gorm:"column:value"`
	ParentReturnReasonId uuid.NullUUID  `json:"parent_return_reason_id"  gorm:"column:parent_return_reason_id"`
	ParentReturnReason   *ReturnReason  `json:"parent_return_reason"  gorm:"column:parent_return_reason;foreignKey:ParentReturnReasonId"`
	ReturnReasonChildren []ReturnReason `json:"return_reason_children"  gorm:"column:return_reason_children;foreignKey:Id"`
}
