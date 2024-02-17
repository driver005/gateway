package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableReturn struct {
	core.FilterModel

	IdempotencyKey string `json:"idempotency_key,omitempty" validate:"omitempty"`
}

type OrderReturnItem struct {
	ItemId   uuid.UUID `json:"item_id"`
	Quantity int       `json:"quantity"`
	ReasonId uuid.UUID `json:"reason_id,omitempty" validate:"omitempty"`
	Note     string    `json:"note,omitempty" validate:"omitempty"`
}

type CreateReturnInput struct {
	OrderId        uuid.UUID                       `json:"order_id"`
	SwapId         uuid.UUID                       `json:"swap_id,omitempty" validate:"omitempty"`
	ClaimOrderId   uuid.UUID                       `json:"claim_order_id,omitempty" validate:"omitempty"`
	Items          []OrderReturnItem               `json:"items,omitempty" validate:"omitempty"`
	ShippingMethod *CreateClaimReturnShippingInput `json:"shipping_method,omitempty" validate:"omitempty"`
	NoNotification bool                            `json:"no_notification,omitempty" validate:"omitempty"`
	Metadata       core.JSONB                      `json:"metadata,omitempty" validate:"omitempty"`
	RefundAmount   float64                         `json:"refund_amount,omitempty" validate:"omitempty"`
	LocationId     uuid.UUID                       `json:"location_id,omitempty" validate:"omitempty"`
}

type UpdateReturnInput struct {
	Items          []OrderReturnItem               `json:"items,omitempty" validate:"omitempty"`
	ShippingMethod *CreateClaimReturnShippingInput `json:"shipping_method,omitempty" validate:"omitempty"`
	NoNotification bool                            `json:"no_notification,omitempty" validate:"omitempty"`
	Metadata       core.JSONB                      `json:"metadata,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostReturnsReturnReceiveReq
// type: object
// description: "The details of the received return."
// required:
//   - items
//
// properties:
//
//	items:
//	  description: The Line Items that have been received.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - item_id
//	      - quantity
//	    properties:
//	      item_id:
//	        description: The ID of the Line Item.
//	        type: string
//	      quantity:
//	        description: The quantity of the Line Item.
//	        type: integer
//	refund:
//	  description: The amount to refund.
//	  type: number
//	location_id:
//	  description: The ID of the location to return items from.
//	  type: string
type ReturnReceive struct {
	Items      []OrderReturnItem `json:"items"`
	Refund     float64           `json:"refund,omitempty" validate:"omitempty"`
	LocationId uuid.UUID         `json:"location_id,omitempty" validate:"omitempty"`
}

type ReturnShippingMethod struct {
	OptionId uuid.UUID `json:"option_id"`
}

// @oas:schema:StorePostReturnsReq
// type: object
// description: "The details of the return to create."
// required:
//   - order_id
//   - items
//
// properties:
//
//	order_id:
//	  type: string
//	  description: The ID of the Order to create the return for.
//	items:
//	  description: "The items to include in the return."
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - item_id
//	      - quantity
//	    properties:
//	      item_id:
//	        description: The ID of the line item to return.
//	        type: string
//	      quantity:
//	        description: The quantity to return.
//	        type: integer
//	      reason_id:
//	        description: The ID of the return reason. Return reasons can be retrieved from the List Return Reasons API Route.
//	        type: string
//	      note:
//	        description: A note to add to the item returned.
//	        type: string
//	return_shipping:
//	  description: The return shipping method used to return the items. If provided, a fulfillment is automatically created for the return.
//	  type: object
//	  required:
//	    - option_id
//	  properties:
//	    option_id:
//	      type: string
//	      description: The ID of the Shipping Option to create the Shipping Method from.
type CreateReturn struct {
	OrderId        uuid.UUID             `json:"order_id"`
	Items          []OrderReturnItem     `json:"items,omitempty" validate:"omitempty"`
	ReturnShipping *ReturnShippingMethod `json:"return_shipping,omitempty" validate:"omitempty"`
}
