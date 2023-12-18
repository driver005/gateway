package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// ClaimOrder - Claim Orders represent a group of faulty or missing items. Each claim order consists of a subset of items associated with an original order, and can contain additional information about fulfillments and returns.
type ClaimOrder struct {
	core.Model

	Type ClaimStatus `json:"type"`

	// The status of the claim's payment
	PaymentStatus string `json:"payment_status" gorm:"default:null"`

	FulfillmentStatus string `json:"fulfillment_status" gorm:"default:null"`

	// The items that have been claimed
	ClaimItems []ClaimItem `json:"claim_items" gorm:"foreignKey:id"`

	// Refers to the new items to be shipped when the claim order has the type `replace`
	AdditionalItems []LineItem `json:"additional_items" gorm:"foreignKey:id"`

	// The ID of the order that the claim comes from.
	OrderId uuid.NullUUID `json:"order_id"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" gorm:"foreignKey:id;references:order_id"`

	// A return object. Holds information about the return if the claim is to be returned. Available if the relation 'return_order' is expanded
	ReturnOrder *Return `json:"return_order" gorm:"foreignKey:id"`

	// The ID of the address that the new items should be shipped to
	ShippingAddressId uuid.NullUUID `json:"shipping_address_id" gorm:"default:null"`

	ShippingAddress *Address `json:"shipping_address" gorm:"foreignKey:id;references:shipping_address_id"`

	// The shipping methods that the claim order will be shipped with.
	ShippingMethods []ShippingMethod `json:"shipping_methods" gorm:"foreignKey:id"`

	// The fulfillments of the new items to be shipped
	Fulfillments []Fulfillment `json:"fulfillments" gorm:"foreignKey:id"`

	// The amount that will be refunded in conjunction with the claim
	RefundAmount int32 `json:"refund_amount" gorm:"default:null"`

	// The date with timezone at which the claim was canceled.
	CanceledAt time.Time `json:"canceled_at" gorm:"default:null"`

	// Flag for describing whether or not notifications related to this should be send.
	NoNotification bool `json:"no_notification" gorm:"default:null"`

	// Randomly generated key used to continue the completion of the cart associated with the claim in case of failure.
	IdempotencyKey string `json:"idempotency_key" gorm:"default:null"`
}
