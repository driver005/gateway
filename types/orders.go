package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type FilterableOrder struct {
	core.FilterModel

	DisplayId       string           `json:"display_id,omitempty" validate:"omitempty"`
	Email           string           `json:"email,omitempty" validate:"omitempty"`
	BillingAddress  *AddressPayload  `json:"billing_address,omitempty" validate:"omitempty"`
	ShippingAddress *AddressPayload  `json:"shipping_address,omitempty" validate:"omitempty"`
	Customer        *models.Customer `json:"customer,omitempty" validate:"omitempty"`

	Status            []models.OrderStatus       `json:"status,omitempty" validate:"omitempty"`
	FulfillmentStatus []models.FulfillmentStatus `json:"fulfillment_status,omitempty" validate:"omitempty"`
	PaymentStatus     []models.PaymentStatus     `json:"payment_status,omitempty" validate:"omitempty"`
	CartId            uuid.UUID                  `json:"cart_id,omitempty" validate:"omitempty"`
	RegionId          uuid.UUID                  `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode      string                     `json:"currency_code,omitempty" validate:"omitempty"`
	TaxRate           string                     `json:"tax_rate,omitempty" validate:"omitempty"`

	CustomerId uuid.UUID `json:"customer_id,omitempty" validate:"omitempty"`
}
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

type OrderShippingMethod struct {
	OptionId uuid.UUID  `json:"option_id"`
	Data     core.JSONB `json:"data,omitempty" validate:"omitempty"`
	Price    float64    `json:"price"`
}

type OrderClaimShipments struct {
	FulfillmentId   uuid.UUID `json:"fulfillment_id"`
	TrackingNumbers []string  `json:"tracking_numbers,omitempty" validate:"omitempty"`
}

type OrderFulfillments struct {
	Items          []FulFillmentItemType `json:"items"`
	LocationId     uuid.UUID             `json:"location_id,omitempty" validate:"omitempty"`
	NoNotification bool                  `json:"no_notification,omitempty" validate:"omitempty"`
	Metadata       core.JSONB            `json:"metadata,omitempty" validate:"omitempty"`
}

type OrderLineItemReservation struct {
	LocationId uuid.UUID `json:"location_id"`
	Quantity   int       `json:"quantity,omitempty" validate:"omitempty"`
}

type CreateOrderShipment struct {
	FulfillmentId   uuid.UUID `json:"fulfillment_id"`
	TrackingNumbers []string  `json:"tracking_numbers,omitempty" validate:"omitempty"`
	NoNotification  bool      `json:"no_notification,omitempty" validate:"omitempty"`
}

type OrderSwap struct {
	ReturnItems           []OrderReturnItem                    `json:"return_items" validate:"required,dive"`
	ReturnShipping        CreateClaimReturnShippingInput       `json:"return_shipping,omitempty" validate:"omitempty,dive"`
	SalesChannelId        string                               `json:"sales_channel_id,omitempty" validate:"omitempty,uuid"`
	AdditionalItems       []CreateClaimItemAdditionalItemInput `json:"additional_items,omitempty" validate:"omitempty,dive"`
	CustomShippingOptions []CreateCustomShippingOptionInput    `json:"custom_shipping_options,omitempty" validate:"omitempty,dive"`
	NoNotification        bool                                 `json:"no_notification,omitempty" validate:"omitempty"`
	ReturnLocationId      string                               `json:"return_location_id,omitempty" validate:"omitempty,uuid"`
	AllowBackorder        bool                                 `json:"allow_backorder,omitempty" validate:"omitempty"`
}

type OrderClaimFulfillments struct {
	LocationId     uuid.UUID  `json:"location_id,omitempty" validate:"omitempty"`
	NoNotification bool       `json:"no_notification,omitempty" validate:"omitempty"`
	Metadata       core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type OrderRefunds struct {
	Amount         float64             `json:"amount"`
	Reason         models.RefundReason `json:"reason"`
	Note           string              `json:"note,omitempty" validate:"omitempty"`
	NoNotification bool                `json:"no_notification,omitempty" validate:"omitempty"`
}

type OrderReturns struct {
	Items          []OrderReturnItem              `json:"items" validate:"dive"`
	ReturnShipping CreateClaimReturnShippingInput `json:"return_shipping,omitempty" validate:"omitempty,nested"`
	Note           string                         `json:"note,omitempty" validate:"omitempty,alphanum"`
	ReceiveNow     bool                           `json:"receive_now,omitempty" validate:"omitempty,boolean"`
	NoNotification bool                           `json:"no_notification,omitempty" validate:"omitempty,boolean"`
	Refund         float64                        `json:"refund,omitempty" validate:"omitempty,numeric"`
	LocationId     uuid.UUID                      `json:"location_id,omitempty" validate:"omitempty,alphanum"`
}

type CustomerAcceptClaim struct {
	Token string `json:"token"`
}

type OrderLookup struct {
	DisplayId       string          `json:"display_id"`
	Email           string          `json:"email" validate:"email"`
	ShippingAddress *AddressPayload `json:"shipping_address,omitempty" validate:"omitempty"`
}

type CustomerOrderClaim struct {
	OrderIds uuid.UUIDs `json:"order_ids"`
}
