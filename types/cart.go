package types

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type GiftCard struct {
	Code string `json:"code"`
}

type Discount struct {
	Code string `json:"code"`
}

type CartCreateProps struct {
	RegionId          uuid.UUID        `json:"region_id,omitempty" validate:"omitempty"`
	Region            *models.Region   `json:"region,omitempty" validate:"omitempty"`
	Email             string           `json:"email,omitempty" validate:"omitempty"`
	BillingAddressId  uuid.UUID        `json:"billing_address_id,omitempty" validate:"omitempty"`
	BillingAddress    *AddressPayload  `json:"billing_address,omitempty" validate:"omitempty"`
	ShippingAddressId uuid.UUID        `json:"shipping_address_id,omitempty" validate:"omitempty"`
	ShippingAddress   *AddressPayload  `json:"shipping_address,omitempty" validate:"omitempty"`
	GiftCards         []GiftCard       `json:"gift_cards,omitempty" validate:"omitempty"`
	Discounts         []Discount       `json:"discounts,omitempty" validate:"omitempty"`
	Customer          *models.Customer `json:"customer,omitempty" validate:"omitempty"`
	CustomerId        uuid.UUID        `json:"customer_id,omitempty" validate:"omitempty"`
	Type              models.CartType  `json:"type,omitempty" validate:"omitempty"`
	Context           core.JSONB       `json:"context,omitempty" validate:"omitempty"`
	Metadata          core.JSONB       `json:"metadata,omitempty" validate:"omitempty"`
	SalesChannelId    uuid.UUID        `json:"sales_channel_id,omitempty" validate:"omitempty"`
	CountryCode       string           `json:"country_code,omitempty" validate:"omitempty"`
}

type CartUpdateProps struct {
	RegionId            uuid.UUID       `json:"region_id,omitempty" validate:"omitempty"`
	CountryCode         string          `json:"country_code,omitempty" validate:"omitempty"`
	Email               string          `json:"email,omitempty" validate:"omitempty"`
	ShippingAddressId   uuid.UUID       `json:"shipping_address_id,omitempty" validate:"omitempty"`
	BillingAddressId    uuid.UUID       `json:"billing_address_id,omitempty" validate:"omitempty"`
	BillingAddress      *AddressPayload `json:"billing_address,omitempty" validate:"omitempty"`
	ShippingAddress     *AddressPayload `json:"shipping_address,omitempty" validate:"omitempty"`
	CompletedAt         *time.Time      `json:"completed_at,omitempty" validate:"omitempty"`
	PaymentAuthorizedAt *time.Time      `json:"payment_authorized_at,omitempty" validate:"omitempty"`
	GiftCards           []GiftCard      `json:"gift_cards,omitempty" validate:"omitempty"`
	Discounts           []Discount      `json:"discounts,omitempty" validate:"omitempty"`
	CustomerId          uuid.UUID       `json:"customer_id,omitempty" validate:"omitempty"`
	Context             core.JSONB      `json:"context,omitempty" validate:"omitempty"`
	Metadata            core.JSONB      `json:"metadata,omitempty" validate:"omitempty"`
	SalesChannelId      uuid.UUID       `json:"sales_channel_id,omitempty" validate:"omitempty"`
}

type FilterableCartProps struct {
	core.FilterModel
}

type CreateCart struct {
	RegionId       uuid.UUID                            `json:"region_id,omitempty" validate:"omitempty,string"`
	CountryCode    string                               `json:"country_code,omitempty" validate:"omitempty,string"`
	Items          []CreateClaimItemAdditionalItemInput `json:"items,omitempty" validate:"omitempty,dive"`
	Context        core.JSONB                           `json:"context,omitempty" validate:"omitempty"`
	SalesChannelId uuid.UUID                            `json:"sales_channel_id,omitempty" validate:"omitempty,string"`
}

type UpdatePaymentSession struct {
	Data core.JSONB `json:"data"`
}

type AddShippingMethod struct {
	OptionId uuid.UUID  `json:"option_id"`
	Data     core.JSONB `json:"data,omitempty" validate:"omitempty"`
}
