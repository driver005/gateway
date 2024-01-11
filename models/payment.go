package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Payment - Payments represent an amount authorized with a given payment method, Payments can be captured, canceled or refunded.
type Payment struct {
	core.Model

	// The ID of the Swap that the Payment is used for.
	SwapId uuid.NullUUID `json:"swap_id" gorm:"default:null"`

	// A swap object. Available if the relation `swap` is expanded.
	Swap *Swap `json:"swap" gorm:"foreignKey:id;references:swap_id"`

	// The id of the Cart that the Payment Session is created for.
	CartId uuid.NullUUID `json:"cart_id" gorm:"default:null"`

	// A cart object. Available if the relation `cart` is expanded.
	Cart *Cart `json:"cart" gorm:"foreignKey:id;references:cart_id"`

	// The ID of the Order that the Payment is used for.
	OrderId uuid.NullUUID `json:"order_id" gorm:"default:null"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" gorm:"foreignKey:id;references:order_id"`

	// The amount that the Payment has been authorized for.
	Amount float64 `json:"amount"`

	// The 3 character ISO currency code that the Payment is completed in.
	CurrencyCode string `json:"currency_code"`

	Currency *Currency `json:"currency" gorm:"foreignKey:code;references:currency_code"`

	// The amount of the original Payment amount that has been refunded back to the Customer.
	AmountRefunded float64 `json:"amount_refunded" gorm:"default:null"`

	// The id of the Payment Provider that is responsible for the Payment
	ProviderId uuid.NullUUID `json:"provider_id"`

	// The data required for the Payment Provider to identify, modify and process the Payment. Typically this will be an object that holds an id to the external payment session, but can be an empty object if the Payment Provider doesn't hold any state.
	Data JSONB `json:"data" gorm:"default:null"`

	// The date with timezone at which the Payment was captured.
	CapturedAt *time.Time `json:"captured_at" gorm:"default:null"`

	// The date with timezone at which the Payment was canceled.
	CanceledAt *time.Time `json:"canceled_at" gorm:"default:null"`

	// Randomly generated key used to continue the completion of a payment in case of failure.
	IdempotencyKey string `json:"idempotency_key" gorm:"default:null"`
}

type PaymentStatus string

const (
	PaymentStatusNotPaid           PaymentStatus = "not_paid"
	PaymentStatusAwaiting          PaymentStatus = "awaiting"
	PaymentStatusCaptured          PaymentStatus = "captured"
	PaymentStatusPartiallyRefunded PaymentStatus = "partially_refunded"
	PaymentStatusRefunded          PaymentStatus = "refunded"
	PaymentStatusCanceled          PaymentStatus = "canceled"
	PaymentStatusRequiresAction    PaymentStatus = "requires_action"
)

func (pl *PaymentStatus) Scan(value interface{}) error {
	*pl = PaymentStatus(value.([]byte))
	return nil
}

func (pl PaymentStatus) Value() (driver.Value, error) {
	return string(pl), nil
}
