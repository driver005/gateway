package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// A payment collection allows grouping and managing a list of payments at one. This can be helpful when making additional payment for order edits or integrating installment payments.
type PaymentCollection struct {
	core.Model

	// The type of the payment collection
	Type string `json:"type"`

	// The type of the payment collection
	Status string `json:"status"`

	// Description of the payment collection
	Description string `json:"description"`

	// Amount of the payment collection.
	Amount int32 `json:"amount"`

	// Authorized amount of the payment collection.
	AuthorizedAmount int32 `json:"authorized_amount"`

	// The ID of the region this payment collection is associated with.
	RegionId uuid.NullUUID `json:"region_id"`

	Region *Region `json:"region,omitempty" gorm:"foreignKey:id;references:region_id"`

	// The three character ISO code for the currency this payment collection is associated with.
	CurrencyCode string `json:"currency_code"`

	Currency *Currency `json:"currency,omitempty"`

	// The details of the payment sessions created as part of the payment collection.
	PaymentSessions []PaymentSession `json:"payment_sessions,omitempty"`

	// The details of the payments created as part of the payment collection.
	Payments []Payment `json:"payments,omitempty"`

	// The ID of the user that created the payment collection.
	CreatedBy string `json:"created_by"`
}
