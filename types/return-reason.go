package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableReturnReason struct {
	core.FilterModel
	ParentReturnReasonId uuid.UUID `json:"parent_return_reason_id,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostReturnReasonsReq
// type: object
// description: "The details of the return reason to create."
// required:
//   - label
//   - value
//
// properties:
//
//	label:
//	  description: "The label to display to the Customer."
//	  type: string
//	value:
//	  description: "A unique value of the return reason."
//	  type: string
//	parent_return_reason_id:
//	  description: "The ID of the parent return reason."
//	  type: string
//	description:
//	  description: "The description of the Reason."
//	  type: string
//	metadata:
//	  description: An optional set of key-value pairs with additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type CreateReturnReason struct {
	Value                string     `json:"value"`
	Label                string     `json:"label"`
	ParentReturnReasonId uuid.UUID  `json:"parent_return_reason_id,omitempty" validate:"omitempty"`
	Description          string     `json:"description,omitempty" validate:"omitempty"`
	Metadata             core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostReturnReasonsReasonReq
// type: object
// description: "The details to update of the return reason."
// properties:
//
//	label:
//	  description: "The label to display to the Customer."
//	  type: string
//	value:
//	  description: "A unique value of the return reason."
//	  type: string
//	description:
//	  description: "The description of the Reason."
//	  type: string
//	metadata:
//	  description: An optional set of key-value pairs with additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type UpdateReturnReason struct {
	Description          string     `json:"description,omitempty" validate:"omitempty"`
	Label                string     `json:"label,omitempty" validate:"omitempty"`
	ParentReturnReasonId uuid.UUID  `json:"parent_return_reason_id,omitempty" validate:"omitempty"`
	Metadata             core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}
