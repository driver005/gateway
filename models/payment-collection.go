package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// A payment collection allows grouping and managing a list of payments at one. This can be helpful when making additional payment for order edits or integrating installment payments.
type PaymentCollection struct {
	core.Model

	// The type of the payment collection
	Type PaymentCollectionType `json:"type"`

	// The type of the payment collection
	Status PaymentCollectionStatus `json:"status"`

	// Description of the payment collection
	Description string `json:"description"`

	// Amount of the payment collection.
	Amount float64 `json:"amount"`

	// Authorized amount of the payment collection.
	AuthorizedAmount float64 `json:"authorized_amount"`

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
	CreatedBy uuid.UUID `json:"created_by"`
}

type PaymentCollectionStatus string

const (
	PaymentCollectionStatusNotPaid             PaymentCollectionStatus = "not_paid"
	PaymentCollectionStatusAwaiting            PaymentCollectionStatus = "awaiting"
	PaymentCollectionStatusAuthorized          PaymentCollectionStatus = "authorized"
	PaymentCollectionStatusPartiallyAuthorized PaymentCollectionStatus = "partially_authorized"
	PaymentCollectionStatusCanceled            PaymentCollectionStatus = "canceled"
)

func (pl *PaymentCollectionStatus) Scan(value interface{}) error {
	*pl = PaymentCollectionStatus(value.([]byte))
	return nil
}

func (pl PaymentCollectionStatus) Value() (driver.Value, error) {
	return string(pl), nil
}

type PaymentCollectionType string

const (
	PaymentCollectionTypeOrderEdit PaymentCollectionType = "order_edit"
)

func (pl *PaymentCollectionType) Scan(value interface{}) error {
	*pl = PaymentCollectionType(value.([]byte))
	return nil
}

func (pl PaymentCollectionType) Value() (driver.Value, error) {
	return string(pl), nil
}
