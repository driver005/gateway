package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:AdminPostProductsToCollectionReq
// type: object
// description: "The details of the products to add to the collection."
// required:
//   - product_ids
//
// properties:
//
//	product_ids:
//	  description: "An array of Product IDs to add to the Product Collection."
//	  type: array
//	  items:
//	    description: "The ID of a Product to add to the Product Collection."
//	    type: string
type AddProductsToCollectionInput struct {
	ProductIds []uuid.UUID `json:"product_ids"`
}

type FilterableCollection struct {
	core.FilterModel
	Title               string    `json:"title,omitempty" validate:"omitempty"`
	Handle              string    `json:"handle,omitempty" validate:"omitempty"`
	DiscountConditionId uuid.UUID `json:"discount_condition_id,omitempty" validate:"omitempty"`
}
