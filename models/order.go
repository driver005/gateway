package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Order - Represents an order
type Order struct {
	core.Model

	// The order's status
	Status OrderStatus `json:"status" gorm:"default:null"`

	// The order's fulfillment status
	FulfillmentStatus FulfillmentStatus `json:"fulfillment_status" gorm:"default:null"`

	// The order's payment status
	PaymentStatus PaymentStatus `json:"payment_status" gorm:"default:null"`

	// The order's display ID
	DisplayId int `json:"display_id" gorm:"default:null"`

	// The ID of the cart associated with the order
	CartId uuid.NullUUID `json:"cart_id" gorm:"default:null"`

	// A cart object. Available if the relation `cart` is expanded.
	Cart *Cart `json:"cart" gorm:"foreignKey:id;references:cart_id"`

	// The ID of the customer associated with the order
	CustomerId uuid.NullUUID `json:"customer_id"`

	// A customer object. Available if the relation `customer` is expanded.
	Customer *Customer `json:"customer" gorm:"foreignKey:id;references:customer_id"`

	// The email associated with the order
	Email string `json:"email"`

	// The ID of the billing address associated with the order
	BillingAddressId uuid.NullUUID `json:"billing_address_id" gorm:"default:null"`

	BillingAddress *Address `json:"billing_address" gorm:"foreignKey:id;references:billing_address_id"`

	// The ID of the shipping address associated with the order
	ShippingAddressId uuid.NullUUID `json:"shipping_address_id" gorm:"default:null"`

	ShippingAddress *Address `json:"shipping_address" gorm:"foreignKey:id;references:shipping_address_id"`

	// The region's ID
	RegionId uuid.NullUUID `json:"region_id"`

	// A region object. Available if the relation `region` is expanded.
	Region *Region `json:"region" gorm:"foreignKey:id;references:region_id"`

	// The 3 character currency code that is used in the order
	CurrencyCode string `json:"currency_code"`

	Currency *Currency `json:"currency" gorm:"foreignKey:code;references:currency_code"`

	// The order's tax rate
	TaxRate float64 `json:"tax_rate" gorm:"default:null"`

	// The discounts used in the order. Available if the relation `discounts` is expanded.
	Discounts []Discount `json:"discounts" gorm:"foreignKey:id"`

	// The gift cards used in the order. Available if the relation `gift_cards` is expanded.
	GiftCards []GiftCard `json:"gift_cards" gorm:"foreignKey:id"`

	// The shipping methods used in the order. Available if the relation `shipping_methods` is expanded.
	ShippingMethods []ShippingMethod `json:"shipping_methods" gorm:"foreignKey:id"`

	// The payments used in the order. Available if the relation `payments` is expanded.
	Payments []Payment `json:"payments" gorm:"foreignKey:id"`

	// The fulfillments used in the order. Available if the relation `fulfillments` is expanded.
	Fulfillments []Fulfillment `json:"fulfillments" gorm:"foreignKey:id"`

	// The returns associated with the order. Available if the relation `returns` is expanded.
	Returns []Return `json:"returns" gorm:"foreignKey:id"`

	// The claims associated with the order. Available if the relation `claims` is expanded.
	Claims []ClaimOrder `json:"claims" gorm:"foreignKey:id"`

	// The refunds associated with the order. Available if the relation `refunds` is expanded.
	Refunds []Refund `json:"refunds" gorm:"foreignKey:id"`

	// The swaps associated with the order. Available if the relation `swaps` is expanded.
	Swaps []Swap `json:"swaps" gorm:"foreignKey:id"`

	// The ID of the draft order this order is associated with.
	DraftOrderId uuid.NullUUID `json:"draft_order_id" gorm:"default:null"`

	// A draft order object. Available if the relation `draft_order` is expanded.
	DraftOrder *DraftOrder `json:"draft_order" gorm:"foreignKey:id;references:draft_order_id"`

	// The line items that belong to the order. Available if the relation `items` is expanded.
	Items []LineItem `json:"items" gorm:"foreignKey:id"`

	// The returnable items that belong to the order. Available if the relation `returnable_items` is expanded.
	ReturnableItems []LineItem `json:"returnable_items" gorm:"foreignKey:id"`

	// Order edits done on the order. Available if the relation `edits` is expanded.
	Edits []OrderEdit `json:"edits" gorm:"foreignKey:id"`

	// The gift card transactions used in the order. Available if the relation `gift_card_transactions` is expanded.
	GiftCardTransactions []GiftCardTransaction `json:"gift_card_transactions" gorm:"foreignKey:id"`

	// The date the order was canceled on.
	CanceledAt *time.Time `json:"canceled_at" gorm:"default:null"`

	// Flag for describing whether or not notifications related to this should be send.
	NoNotification bool `json:"no_notification" gorm:"default:null"`

	// Randomly generated key used to continue the processing of the order in case of failure.
	IdempotencyKey string `json:"idempotency_key" gorm:"default:null"`

	// The ID of an external order.
	ExternalId uuid.NullUUID `json:"external_id" gorm:"default:null"`

	// The ID of the sales channel this order is associated with.
	SalesChannelId uuid.NullUUID `json:"sales_channel_id" gorm:"default:null"`

	// A sales channel object. Available if the relation `sales_channel` is expanded.
	SalesChannel *SalesChannel `json:"sales_channel" gorm:"foreignKey:id;references:sales_channel_id"`

	// The total of shipping
	ShippingTotal float64 `json:"shipping_total" gorm:"default:null"`

	// The total of gift cards
	ShippingTaxTotal float64 `json:"shipping_tax_total" gorm:"default:null"`

	// The total of discount
	DiscountTotal float64 `json:"discount_total" gorm:"default:null"`

	// The total of the discount
	RawDiscountTotal float64 `json:"raw_discount_total" gorm:"default:null"`

	// The total of gift cards
	ItemTaxTotal float64 `json:"item_tax_total" gorm:"default:null"`

	// The total of tax
	TaxTotal float64 `json:"tax_total" gorm:"default:null"`

	// The total amount refunded if the order is returned.
	RefundedTotal float64 `json:"refunded_total" gorm:"default:null"`

	// The total amount of the order
	Total float64 `json:"total" gorm:"default:null"`

	// The subtotal of the order
	Subtotal float64 `json:"subtotal" gorm:"default:null"`

	// The total amount paid
	PaidTotal float64 `json:"paid_total" gorm:"default:null"`

	// The amount that can be refunded
	RefundableAmount float64 `json:"refundable_amount" gorm:"default:null"`

	// The total of gift cards
	GiftCardTotal float64 `json:"gift_card_total" gorm:"default:null"`

	// The total of gift cards with taxes
	GiftCardTaxTotal float64 `json:"gift_card_tax_total" gorm:"default:null"`
}

type OrderStatus string

const (
	OrderStatusPending        OrderStatus = "pending"
	OrderStatusCompleted      OrderStatus = "completed"
	OrderStatusArchived       OrderStatus = "archived"
	OrderStatusCanceled       OrderStatus = "canceled"
	OrderStatusRefunded       OrderStatus = "refunded"
	OrderStatusRequiresAction OrderStatus = "requires_action"
)

func (pl *OrderStatus) Scan(value interface{}) error {
	*pl = OrderStatus(value.([]byte))
	return nil
}

func (pl OrderStatus) Value() (driver.Value, error) {
	return string(pl), nil
}
