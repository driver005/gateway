package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableDraftOrder struct {
	core.FilterModel
}

type DraftOrderListSelector struct {
	Q string `json:"q,omitempty" validate:"omitempty"`
}

type DraftOrderCreate struct {
	Status              string           `json:"status,omitempty" validate:"omitempty"`
	Email               string           `json:"email"`
	BillingAddressId    uuid.UUID        `json:"billing_address_id,omitempty" validate:"omitempty"`
	BillingAddress      *AddressPayload  `json:"billing_address,omitempty" validate:"omitempty"`
	ShippingAddressId   uuid.UUID        `json:"shipping_address_id,omitempty" validate:"omitempty"`
	ShippingAddress     *AddressPayload  `json:"shipping_address,omitempty" validate:"omitempty"`
	Items               []Item           `json:"items,omitempty" validate:"omitempty"`
	RegionId            uuid.UUID        `json:"region_id"`
	Discounts           []Discount       `json:"discounts,omitempty" validate:"omitempty"`
	CustomerId          uuid.UUID        `json:"customer_id,omitempty" validate:"omitempty"`
	NoNotificationOrder bool             `json:"no_notification_order,omitempty" validate:"omitempty"`
	ShippingMethods     []ShippingMethod `json:"shipping_methods"`
	Metadata            core.JSONB       `json:"metadata,omitempty" validate:"omitempty"`
	IdempotencyKey      string           `json:"idempotency_key,omitempty" validate:"omitempty"`
}

type ShippingMethod struct {
	OptionId uuid.UUID  `json:"option_id"`
	Data     core.JSONB `json:"data,omitempty" validate:"omitempty"`
	Price    float64    `json:"price,omitempty" validate:"omitempty"`
}

type Item struct {
	Title     string     `json:"title,omitempty" validate:"omitempty"`
	UnitPrice float64    `json:"unit_price,omitempty" validate:"omitempty"`
	VariantId uuid.UUID  `json:"variant_id,omitempty" validate:"omitempty"`
	Quantity  int        `json:"quantity"`
	Metadata  core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}
