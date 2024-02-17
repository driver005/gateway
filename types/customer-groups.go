package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:AdminPostCustomerGroupsGroupCustomersBatchReq
// type: object
// description: "The customers to add to the customer group."
// required:
//   - customer_ids
//
// properties:
//
//	customer_ids:
//	  description: "The ids of the customers to add"
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - id
//	    properties:
//	      id:
//	        description: ID of the customer
//	        type: string
type CustomersToCustomerGroup struct {
	CustomerIds []uuid.UUID `json:"customer_ids"`
}

type CustomerGroupsBatchCustomer struct {
	Id uuid.UUID `json:"id"`
}

// @oas:schema:AdminPostCustomerGroupsReq
// type: object
// description: "The details of the customer group to create."
// required:
//   - name
//
// properties:
//
//	name:
//	  type: string
//	  description: Name of the customer group
//	metadata:
//	  type: object
//	  description: Metadata of the customer group.
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type CreateCustomerGroup struct {
	Name     string     `json:"name,omitempty" validate:"omitempty"`
	Metadata core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostCustomerGroupsGroupReq
// type: object
// description: "The details to update in the customer group."
// properties:
//
//	name:
//	  description: "Name of the customer group"
//	  type: string
//	metadata:
//	  description: "Metadata of the customer group."
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type UpdateCustomerGroup struct {
	Name     string     `json:"name,omitempty" validate:"omitempty"`
	Metadata core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type FilterableCustomerGroup struct {
	core.FilterModel

	Name                []string  `json:"name,omitempty" validate:"omitempty"`
	DiscountConditionId uuid.UUID `json:"discount_condition_id,omitempty" validate:"omitempty"`
}
