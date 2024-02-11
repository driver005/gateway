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

type CreateOrderEditInput struct {
	OrderId      uuid.UUID `json:"order_id"`
	InternalNote string    `json:"internal_note,omitempty" validate:"omitempty"`
}

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

type OrderEditsEditLineItem struct {
	Quantity int `json:"quantity"`
}
