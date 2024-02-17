package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableSwap struct {
	core.FilterModel

	IdempotencyKey string `json:"idempotency_key,omitempty" validate:"omitempty"`
}

// @oas:schema:StorePostSwapsReq
// type: object
// description: "The details of the swap to create."
// required:
//   - order_id
//   - return_items
//   - additional_items
//
// properties:
//
//	order_id:
//	  type: string
//	  description: The ID of the Order to create the Swap for.
//	return_items:
//	  description: "The items to include in the Return."
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - item_id
//	      - quantity
//	    properties:
//	      item_id:
//	        description: The ID of the order's line item to return.
//	        type: string
//	      quantity:
//	        description: The quantity to return.
//	        type: integer
//	      reason_id:
//	        description: The ID of the reason of this return. Return reasons can be retrieved from the List Return Reasons API Route.
//	        type: string
//	      note:
//	        description: The note to add to the item being swapped.
//	        type: string
//	return_shipping_option:
//	  type: string
//	  description: The ID of the Shipping Option to create the Shipping Method from.
//	additional_items:
//	  description: "The items to exchange the returned items with."
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - variant_id
//	      - quantity
//	    properties:
//	      variant_id:
//	        description: The ID of the Product Variant.
//	        type: string
//	      quantity:
//	        description: The quantity of the variant.
//	        type: integer
type CreateSwap struct {
	OrderId              uuid.UUID                            `json:"order_id" validate:"required"`
	ReturnItems          []OrderReturnItem                    `json:"return_items" validate:"dive"`
	AdditionalItems      []CreateClaimItemAdditionalItemInput `json:"additional_items" validate:"dive"`
	ReturnShippingOption uuid.UUID                            `json:"return_shipping_option,omitempty" validate:"omitempty"`
}
