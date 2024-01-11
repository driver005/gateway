package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Swap - Swaps can be created when a Customer wishes to exchange Products that they have purchased to different Products. Swaps consist of a Return of previously purchased Products and a Fulfillment of new Products, the amount paid for the Products being returned will be used towards payment for the new Products. In the case where the amount paid for the the Products being returned exceed the amount to be paid for the new Products, a Refund will be issued for the difference.
type Swap struct {
	core.Model

	// The status of the Fulfillment of the Swap.
	FulfillmentStatus SwapFulfillmentStatus `json:"fulfillment_status"`

	// The status of the Payment of the Swap. The payment may either refer to the refund of an amount or the authorization of a new amount.
	PaymentStatus SwapPaymentStatus `json:"payment_status"`

	// The ID of the Order where the Line Items to be returned where purchased.
	OrderId uuid.NullUUID `json:"order_id"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" gorm:"foreignKey:id;references:order_id"`

	// The new Line Items to ship to the Customer. Available if the relation `additional_items` is expanded.
	AdditionalItems []LineItem `json:"additional_items" gorm:"foreignKey:id"`

	// A return order object. The Return that is issued for the return part of the Swap. Available if the relation `return_order` is expanded.
	ReturnOrder *Return `json:"return_order" gorm:"foreignKey:id"`

	// The Fulfillments used to send the new Line Items. Available if the relation `fulfillments` is expanded.
	Fulfillments []Fulfillment `json:"fulfillments" gorm:"foreignKey:id"`

	Payment *Payment `json:"payment" gorm:"foreignKey:id"`

	// The difference that is paid or refunded as a result of the Swap. May be negative when the amount paid for the returned items exceed the total of the new Products.
	DifferenceDue float64 `json:"difference_due" gorm:"default:null"`

	// The Address to send the new Line Items to - in most cases this will be the same as the shipping address on the Order.
	ShippingAddressId uuid.NullUUID `json:"shipping_address_id" gorm:"default:null"`

	ShippingAddress *Address `json:"shipping_address" gorm:"foreignKey:id;references:shipping_address_id"`

	// The Shipping Methods used to fulfill the additional items purchased. Available if the relation `shipping_methods` is expanded.
	ShippingMethods []ShippingMethod `json:"shipping_methods" gorm:"foreignKey:id"`

	// The id of the Cart that the Customer will use to confirm the Swap.
	CartId uuid.NullUUID `json:"cart_id" gorm:"default:null"`

	// A cart object. Available if the relation `cart` is expanded.
	Cart *Cart `json:"cart" gorm:"foreignKey:id;references:cart_id"`

	// If true, swaps can be completed with items out of stock
	AllowBackorder bool `json:"allow_backorder" gorm:"default:null"`

	// Randomly generated key used to continue the completion of the swap in case of failure.
	IdempotencyKey string `json:"idempotency_key" gorm:"default:null"`

	// The date with timezone at which the Swap was confirmed by the Customer.
	ConfirmedAt *time.Time `json:"confirmed_at" gorm:"default:null"`

	// The date with timezone at which the Swap was canceled.
	CanceledAt *time.Time `json:"canceled_at" gorm:"default:null"`

	// If set to true, no notification will be sent related to this swap
	NoNotification bool `json:"no_notification" gorm:"default:null"`
}

type SwapFulfillmentStatus string

const (
	SwapFulfillmentNotFulfilled     SwapFulfillmentStatus = "not_fulfilled"
	SwapFulfillmentFulfilled        SwapFulfillmentStatus = "fulfilled"
	SwapFulfillmentShipped          SwapFulfillmentStatus = "shipped"
	SwapFulfillmentPartiallyShipped SwapFulfillmentStatus = "partially_shipped"
	SwapFulfillmentCanceled         SwapFulfillmentStatus = "canceled"
	SwapFulfillmentRequiresAction   SwapFulfillmentStatus = "requires_action"
)

func (pl *SwapFulfillmentStatus) Scan(value interface{}) error {
	*pl = SwapFulfillmentStatus(value.([]byte))
	return nil
}

func (pl SwapFulfillmentStatus) Value() (driver.Value, error) {
	return string(pl), nil
}

type SwapPaymentStatus string

const (
	SwapPaymentNotPaid            SwapPaymentStatus = "not_paid"
	SwapPaymentAwaiting           SwapPaymentStatus = "awaiting"
	SwapPaymentCaptured           SwapPaymentStatus = "captured"
	SwapPaymentConfirmed          SwapPaymentStatus = "confirmed"
	SwapPaymentCanceled           SwapPaymentStatus = "canceled"
	SwapPaymentDifferenceRefunded SwapPaymentStatus = "difference_refunded"
	SwapPaymentPartiallyRefunded  SwapPaymentStatus = "partially_refunded"
	SwapPaymentRefunded           SwapPaymentStatus = "refunded"
	SwapPaymentRequiresAction     SwapPaymentStatus = "requires_action"
)

func (pl *SwapPaymentStatus) Scan(value interface{}) error {
	*pl = SwapPaymentStatus(value.([]byte))
	return nil
}

func (pl SwapPaymentStatus) Value() (driver.Value, error) {
	return string(pl), nil
}
