package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type ShippingMethodOrder struct {
	ProviderId uuid.UUID    `json:"provider_id,omitempty" validate:"omitempty"`
	ProfileId  uuid.UUID    `json:"profile_id,omitempty" validate:"omitempty"`
	Price      float64      `json:"price,omitempty" validate:"omitempty"`
	Data       core.JSONB   `json:"data,omitempty" validate:"omitempty"`
	Items      []core.JSONB `json:"items,omitempty" validate:"omitempty"`
}

type CreateOrderInput struct {
	Status          *models.OrderStatus   `json:"status,omitempty" validate:"omitempty"`
	Email           string                `json:"email"`
	BillingAddress  *AddressPayload       `json:"billing_address"`
	ShippingAddress *AddressPayload       `json:"shipping_address"`
	Items           []models.LineItem     `json:"items"`
	Region          string                `json:"region"`
	Discounts       []models.Discount     `json:"discounts,omitempty" validate:"omitempty"`
	CustomerId      uuid.UUID             `json:"customer_id"`
	PaymentMethod   *PaymentMethod        `json:"payment_method"`
	ShippingMethod  []ShippingMethodOrder `json:"shipping_method,omitempty" validate:"omitempty"`
	NoNotification  bool                  `json:"no_notification,omitempty" validate:"omitempty"`
}

type UpdateOrderInput struct {
	Email             string                   `json:"email,omitempty" validate:"omitempty"`
	BillingAddress    *AddressPayload          `json:"billing_address,omitempty" validate:"omitempty"`
	ShippingAddress   *AddressPayload          `json:"shipping_address,omitempty" validate:"omitempty"`
	Items             []models.LineItem        `json:"items,omitempty" validate:"omitempty"`
	Region            string                   `json:"region,omitempty" validate:"omitempty"`
	Discounts         []models.Discount        `json:"discounts,omitempty" validate:"omitempty"`
	CustomerId        uuid.UUID                `json:"customer_id,omitempty" validate:"omitempty"`
	PaymentMethod     *PaymentMethod           `json:"payment_method,omitempty" validate:"omitempty"`
	ShippingMethod    []ShippingMethodOrder    `json:"shipping_method,omitempty" validate:"omitempty"`
	NoNotification    bool                     `json:"no_notification,omitempty" validate:"omitempty"`
	Payment           *models.Payment          `json:"payment,omitempty" validate:"omitempty"`
	Status            models.OrderStatus       `json:"status,omitempty" validate:"omitempty"`
	FulfillmentStatus models.FulfillmentStatus `json:"fulfillment_status,omitempty" validate:"omitempty"`
	PaymentStatus     models.PaymentStatus     `json:"payment_status,omitempty" validate:"omitempty"`
	Metadata          core.JSONB               `json:"metadata,omitempty" validate:"omitempty"`
}

type AdminListOrdersSelector struct {
	Q                 string          `json:"q,omitempty" validate:"omitempty"`
	Id                uuid.UUID       `json:"id,omitempty" validate:"omitempty"`
	Status            []string        `json:"status,omitempty" validate:"omitempty,dive,oneof=OrderStatus"`
	FulfillmentStatus []string        `json:"fulfillment_status,omitempty" validate:"omitempty,dive,oneof=FulfillmentStatus"`
	PaymentStatus     []string        `json:"payment_status,omitempty" validate:"omitempty,dive,oneof=PaymentStatus"`
	DisplayId         uuid.UUID       `json:"display_id,omitempty" validate:"omitempty"`
	CartId            uuid.UUID       `json:"cart_id,omitempty" validate:"omitempty"`
	CustomerId        uuid.UUID       `json:"customer_id,omitempty" validate:"omitempty"`
	Email             string          `json:"email,omitempty" validate:"omitempty"`
	RegionId          uuid.UUIDs      `json:"region_id,omitempty" validate:"omitempty,dive,oneof=string"`
	CurrencyCode      string          `json:"currency_code,omitempty" validate:"omitempty"`
	TaxRate           string          `json:"tax_rate,omitempty" validate:"omitempty"`
	SalesChannelId    uuid.UUIDs      `json:"sales_channel_id,omitempty" validate:"omitempty"`
	CanceledAt        *core.TimeModel `json:"canceled_at,omitempty" validate:"omitempty"`
	CreatedAt         *core.TimeModel `json:"created_at,omitempty" validate:"omitempty"`
	UpdatedAt         *core.TimeModel `json:"updated_at,omitempty" validate:"omitempty"`
}

type OrdersReturnItem struct {
	ItemId   uuid.UUID `json:"item_id"`
	Quantity int       `json:"quantity"`
	ReasonId uuid.UUID `json:"reason_id,omitempty" validate:"omitempty"`
	Note     string    `json:"note,omitempty" validate:"omitempty"`
}

type TotalsContext struct {
	ForceTaxes      bool `json:"force_taxes,omitempty" validate:"omitempty"`
	ReturnableItems bool `json:"returnable_items,omitempty" validate:"omitempty"`
}
