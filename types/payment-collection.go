package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type CreatePaymentCollectionInput struct {
	RegionId     uuid.UUID                    `json:"region_id"`
	Type         models.PaymentCollectionType `json:"type"`
	CurrencyCode string                       `json:"currency_code"`
	Amount       float64                      `json:"amount"`
	CreatedBy    uuid.UUID                    `json:"created_by"`
	Metadata     core.JSONB                   `json:"metadata,omitempty" validate:"omitempty"`
	Description  string                       `json:"description,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminUpdatePaymentCollectionsReq
// type: object
// description: "The details to update of the payment collection."
// properties:
//
//	description:
//	  description: A description to create or update the payment collection.
//	  type: string
//	metadata:
//	  description: A set of key-value pairs to hold additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type UpdatePaymentCollectionInput struct {
	Metadata    core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
	Description string     `json:"description,omitempty" validate:"omitempty"`
}

type PaymentCollectionsSessionsBatchInput struct {
	ProviderId uuid.UUID `json:"provider_id"`
	Amount     float64   `json:"amount"`
	SessionId  uuid.UUID `json:"session_id,omitempty" validate:"omitempty"`
}

// @oas:schema:StorePostCartsCartPaymentSessionReq
// type: object
// description: "The details of the payment session to set."
// required:
//   - provider_id
//
// properties:
//
//	provider_id:
//	  type: string
//	  description: The ID of the Payment Provider.
type SessionsInput struct {
	ProviderId uuid.UUID `json:"provider_id"`
}

// @oas:schema:StorePostPaymentCollectionsBatchSessionsAuthorizeReq
// type: object
// description: "The details of the payment sessions to authorize."
// required:
//   - session_ids
//
// properties:
//
//	session_ids:
//	  description: "List of Payment Session IDs to authorize."
//	  type: array
//	  items:
//	    type: string
type PaymentCollectionsAuthorizeBatch struct {
	SessionIds uuid.UUIDs `json:"session_ids"`
}

// @oas:schema:StorePostPaymentCollectionsBatchSessionsReq
// type: object
// description: "The details of the payment sessions to manage."
// required:
//   - sessions
//
// properties:
//
//	sessions:
//	  description: "Payment sessions related to the Payment Collection. Existing sessions that are not added in this array will be deleted."
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - provider_id
//	      - amount
//	    properties:
//	      provider_id:
//	        type: string
//	        description: The ID of the Payment Provider.
//	      amount:
//	        type: integer
//	        description: "The payment amount"
//	      session_id:
//	        type: string
//	        description: "The ID of the Payment Session to be updated. If no ID is provided, a new payment session is created."
type PaymentCollectionsSessionsBatch struct {
	Sessions []PaymentCollectionsSessionsBatchInput `json:"sessions"`
}
