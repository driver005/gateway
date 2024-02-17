package types

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableGiftCard struct {
	core.FilterModel
}

// @oas:schema:AdminPostGiftCardsReq
// type: object
// description: "The details of the gift card to create."
// required:
//   - region_id
//
// properties:
//
//	value:
//	  type: integer
//	  description: The value (excluding VAT) that the Gift Card should represent.
//	is_disabled:
//	  type: boolean
//	  description: >-
//	    Whether the Gift Card is disabled on creation. If set to `true`, the gift card will not be available for customers.
//	ends_at:
//	  type: string
//	  format: date-time
//	  description: The date and time at which the Gift Card should no longer be available.
//	region_id:
//	  description: The ID of the Region in which the Gift Card can be used.
//	  type: string
//	metadata:
//	  description: An optional set of key-value pairs to hold additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type CreateGiftCardInput struct {
	OrderId    uuid.UUID  `json:"order_id,omitempty" validate:"omitempty"`
	Value      float64    `json:"value,omitempty" validate:"omitempty"`
	Balance    float64    `json:"balance,omitempty" validate:"omitempty"`
	EndsAt     *time.Time `json:"ends_at,omitempty" validate:"omitempty"`
	IsDisabled bool       `json:"is_disabled,omitempty" validate:"omitempty"`
	RegionId   uuid.UUID  `json:"region_id"`
	Metadata   core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
	TaxRate    float64    `json:"tax_rate,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostGiftCardsGiftCardReq
// type: object
// description: "The details to update of the gift card."
// properties:
//
//	balance:
//	  type: integer
//	  description: The value (excluding VAT) that the Gift Card should represent.
//	is_disabled:
//	  type: boolean
//	  description: >-
//	    Whether the Gift Card is disabled on creation. If set to `true`, the gift card will not be available for customers.
//	ends_at:
//	  type: string
//	  format: date-time
//	  description: The date and time at which the Gift Card should no longer be available.
//	region_id:
//	  description: The ID of the Region in which the Gift Card can be used.
//	  type: string
//	metadata:
//	  description: An optional set of key-value pairs to hold additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type UpdateGiftCardInput struct {
	Balance    float64    `json:"balance,omitempty" validate:"omitempty"`
	EndsAt     *time.Time `json:"ends_at,omitempty" validate:"omitempty"`
	IsDisabled bool       `json:"is_disabled,omitempty" validate:"omitempty"`
	RegionId   uuid.UUID  `json:"region_id,omitempty" validate:"omitempty"`
	Metadata   core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type CreateGiftCardTransactionInput struct {
	GiftCardId uuid.UUID `json:"gift_card_id"`
	OrderId    uuid.UUID `json:"order_id"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at,omitempty" validate:"omitempty"`
	IsTaxable  bool      `json:"is_taxable,omitempty" validate:"omitempty"`
	TaxRate    float64   `json:"tax_rate,omitempty" validate:"omitempty"`
}
