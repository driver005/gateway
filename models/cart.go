package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Cart - Represents a user cart
type Cart struct {
	core.Model

	// The email associated with the cart
	Email string `json:"email" gorm:"default:null"`

	// The billing address's ID
	BillingAddressId uuid.NullUUID `json:"billing_address_id" gorm:"default:null"`

	BillingAddress *Address `json:"billing_address" gorm:"foreignKey:id;references:billing_address_id"`

	// The shipping address's ID
	ShippingAddressId uuid.NullUUID `json:"shipping_address_id" gorm:"default:null"`

	ShippingAddress *Address `json:"shipping_address" gorm:"foreignKey:id;references:shipping_address_id"`

	// Available if the relation `items` is expanded.
	Items []LineItem `json:"items" gorm:"foreignKey:id"`

	// The region's ID
	RegionId uuid.NullUUID `json:"region_id" gorm:"default:null"`

	// A region object. Available if the relation `region` is expanded.
	Region *Region `json:"region" gorm:"foreignKey:id;references:region_id"`

	// Available if the relation `discounts` is expanded.
	Discounts []Discount `json:"discounts" gorm:"foreignKey:id"`

	// Available if the relation `gift_cards` is expanded.
	GiftCards []GiftCard `json:"gift_cards" gorm:"foreignKey:id"`

	// The customer's ID
	CustomerId uuid.NullUUID `json:"customer_id" gorm:"default:null"`

	// A customer object. Available if the relation `customer` is expanded.
	Customer *Customer `json:"customer" gorm:"foreignKey:id;references:customer_id"`

	PaymentSession *PaymentSession `json:"payment_session" gorm:"foreignKey:id"`

	// The payment sessions created on the cart.
	PaymentSessions []PaymentSession `json:"payment_sessions" gorm:"foreignKey:id"`

	// The payment's ID if available
	PaymentId uuid.NullUUID `json:"payment_id" gorm:"default:null"`

	Payment *Payment `json:"payment" gorm:"foreignKey:id;references:payment_id"`

	// The shipping methods added to the cart.
	ShippingMethods []ShippingMethod `json:"shipping_methods" gorm:"foreignKey:id"`

	// The cart's type.
	Type string `json:"type" gorm:"default:null"`

	// The date with timezone at which the cart was completed.
	CompletedAt time.Time `json:"completed_at" gorm:"default:null"`

	// The date with timezone at which the payment was authorized.
	PaymentAuthorizedAt time.Time `json:"payment_authorized_at" gorm:"default:null"`

	// Randomly generated key used to continue the completion of a cart in case of failure.
	IdempotencyKey string `json:"idempotency_key" gorm:"default:null"`

	// The context of the cart which can include info like IP or user agent.
	Context JSONB `json:"context" gorm:"default:null"`

	// The sales channel ID the cart is associated with.
	SalesChannelId uuid.NullUUID `json:"sales_channel_id" gorm:"default:null"`

	// A sales channel object. Available if the relation `sales_channel` is expanded.
	SalesChannel *SalesChannel `json:"sales_channel" gorm:"foreignKey:id;references:sales_channel_id"`

	// The total of shipping
	ShippingTotal float64 `json:"shipping_total" gorm:"default:null"`

	// The total of discount
	DiscountTotal float64 `json:"discount_total" gorm:"default:null"`

	// The total of tax
	TaxTotal float64 `json:"tax_total" gorm:"default:null"`

	// The total amount refunded if the order associated with this cart is returned.
	RefundedTotal float64 `json:"refunded_total" gorm:"default:null"`

	// The total amount of the cart
	Total float64 `json:"total" gorm:"default:null"`

	// The subtotal of the cart
	Subtotal float64 `json:"subtotal" gorm:"default:null"`

	// The amount that can be refunded
	RefundableAmount float64 `json:"refundable_amount" gorm:"default:null"`

	// The total of gift cards
	GiftCardTotal float64 `json:"gift_card_total" gorm:"default:null"`

	// The total of gift cards with taxes
	GiftCardTaxTotal float64 `json:"gift_card_tax_total" gorm:"default:null"`
}
