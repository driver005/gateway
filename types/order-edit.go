package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type FilterableOrderEdit struct {
	core.FilterModel
	OrderId      uuid.UUID `json:"order_id,omitempty" validate:"omitempty"`
	InternalNote string    `json:"internal_note,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostOrderEditsReq
// type: object
// description: "The details of the order edit to create."
// required:
//   - order_id
//
// properties:
//
//	order_id:
//	  description: The ID of the order to create the edit for.
//	  type: string
//	internal_note:
//	  description: An optional note to associate with the order edit.
//	  type: string
type CreateOrderEditInput struct {
	OrderId      uuid.UUID `json:"order_id"`
	InternalNote string    `json:"internal_note,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostOrderEditsOrderEditReq
// type: object
// description: "The details to update of the order edit."
// properties:
//
//	internal_note:
//	  description: An optional note to create or update in the order edit.
//	  type: string
type UpdateOrderEditInput struct {
	InternalNote string `json:"internal_note"`
}

// @oas:schema:AdminPostOrderEditsEditLineItemsReq
// type: object
// description: "The details of the line item change to create."
// required:
//   - variant_id
//   - quantity
//
// properties:
//
//	variant_id:
//	  description: The ID of the product variant associated with the item.
//	  type: string
//	quantity:
//	  description: The quantity of the item.
//	  type: number
//	metadata:
//	  description: An optional set of key-value pairs to hold additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type AddOrderEditLineItemInput struct {
	Quantity  int        `json:"quantity"`
	VariantId uuid.UUID  `json:"variant_id"`
	Metadata  core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type CreateOrderEditItemChangeInput struct {
	Type               models.OrderEditItemChangeType `json:"type"`
	OrderEditId        uuid.UUID                      `json:"order_edit_id"`
	OriginalLineItemId uuid.UUID                      `json:"original_line_item_id,omitempty" validate:"omitempty"`
	LineItemId         uuid.UUID                      `json:"line_item_id,omitempty" validate:"omitempty"`
}

type OrderEditsRequestConfirmation struct {
	PaymentCollectionDescription string `json:"payment_collection_description,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostOrderEditsEditLineItemsLineItemReq
// type: object
// description: "The details to create or update of the line item change."
// required:
//   - quantity
//
// properties:
//
//	quantity:
//	  description: The quantity to update
//	  type: number
type OrderEditsEditLineItem struct {
	Quantity int `json:"quantity"`
}

// @oas:schema:StorePostOrderEditsOrderEditDecline
// type: object
// description: "The details of the order edit's decline."
// properties:
//
//	declined_reason:
//	  type: string
//	  description: The reason for declining the Order Edit.
type OrderEditsDecline struct {
	DeclinedReason string `json:"declined_reason,omitempty" validate:"omitempty"`
}

// @oas:schema:StorePostCartsCartLineItemsItemReq
// type: object
// description: "The details to update of the line item."
// required:
//   - quantity
//
// properties:
//
//	quantity:
//	  type: number
//	  description: The quantity of the line item in the cart.
//	metadata:
//	  type: object
//	  description: An optional key-value map with additional details about the Line Item. If omitted, the metadata will remain unchanged."
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type UpdateLineItem struct {
	Quantity int        `json:"quantity"`
	Metadata core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}
