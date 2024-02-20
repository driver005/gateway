package models

import (
	"database/sql/driver"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// @oas:schema:PaymentCollection
// title: "Payment Collection"
// description: "A payment collection allows grouping and managing a list of payments at one. This can be helpful when making additional payment for order edits or integrating installment payments."
// type: object
// required:
//   - amount
//   - authorized_amount
//   - created_at
//   - created_by
//   - currency_code
//   - deleted_at
//   - description
//   - id
//   - metadata
//   - region_id
//   - status
//   - type
//   - updated_at
//
// properties:
//
//	id:
//	  description: The payment collection's ID
//	  type: string
//	  example: paycol_01G8TJSYT9M6AVS5N4EMNFS1EK
//	type:
//	  description: The type of the payment collection
//	  type: string
//	  enum:
//	    - order_edit
//	status:
//	  description: The type of the payment collection
//	  type: string
//	  enum:
//	    - not_paid
//	    - awaiting
//	    - authorized
//	    - partially_authorized
//	    - canceled
//	description:
//	  description: Description of the payment collection
//	  nullable: true
//	  type: string
//	amount:
//	  description: Amount of the payment collection.
//	  type: integer
//	authorized_amount:
//	  description: Authorized amount of the payment collection.
//	  nullable: true
//	  type: integer
//	region_id:
//	  description: The ID of the region this payment collection is associated with.
//	  type: string
//	  example: reg_01G1G5V26T9H8Y0M4JNE3YGA4G
//	region:
//	  description: The details of the region this payment collection is associated with.
//	  x-expandable: "region"
//	  nullable: true
//	  $ref: "#/components/schemas/Region"
//	currency_code:
//	  description: The three character ISO code for the currency this payment collection is associated with.
//	  type: string
//	  example: usd
//	  externalDocs:
//	    url: https://en.wikipedia.org/wiki/ISO_4217#Active_codes
//	    description: See a list of codes.
//	currency:
//	  description: The details of the currency this payment collection is associated with.
//	  x-expandable: "currency"
//	  nullable: true
//	  $ref: "#/components/schemas/Currency"
//	payment_sessions:
//	  description: The details of the payment sessions created as part of the payment collection.
//	  type: array
//	  x-expandable: "payment_sessions"
//	  items:
//	    $ref: "#/components/schemas/PaymentSession"
//	payments:
//	  description: The details of the payments created as part of the payment collection.
//	  type: array
//	  x-expandable: "payments"
//	  items:
//	    $ref: "#/components/schemas/Payment"
//	created_by:
//	  description: The ID of the user that created the payment collection.
//	  type: string
//	created_at:
//	  description: The date with timezone at which the resource was created.
//	  type: string
//	  format: date-time
//	updated_at:
//	  description: The date with timezone at which the resource was updated.
//	  type: string
//	  format: date-time
//	deleted_at:
//	  description: The date with timezone at which the resource was deleted.
//	  nullable: true
//	  type: string
//	  format: date-time
//	metadata:
//	  description: An optional key-value map with additional details
//	  nullable: true
//	  type: object
//	  example: {car: "white"}
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type PaymentCollection struct {
	core.Model

	Type             PaymentCollectionType   `json:"type"  gorm:"column:type"`
	Status           PaymentCollectionStatus `json:"status"  gorm:"column:status"`
	Description      string                  `json:"description"  gorm:"column:description"`
	Amount           float64                 `json:"amount"  gorm:"column:amount"`
	AuthorizedAmount float64                 `json:"authorized_amount"  gorm:"column:authorized_amount"`
	RegionId         uuid.NullUUID           `json:"region_id"  gorm:"column:region_id"`
	Region           *Region                 `json:"region"  gorm:"column:region;foreignKey:RegionId"`
	CurrencyCode     string                  `json:"currency_code"  gorm:"column:currency_code"`
	Currency         *Currency               `json:"currency"  gorm:"column:currency;foreignKey:CurrencyCode;foreignKey:Code"`
	PaymentSessions  []PaymentSession        `json:"payment_sessions"  gorm:"column:payment_sessions;many2many:payment_collection_sessions"`
	Payments         []Payment               `json:"payments"  gorm:"column:payments;many2many:payment_collection_payments"`
	CreatedBy        uuid.UUID               `json:"created_by"  gorm:"column:created_by"`
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
