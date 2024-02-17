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

// @oas:schema:AdminPostOrdersOrderReq
// type: object
// description: "The details to update of the order."
// properties:
//
//	email:
//	  description: The email associated with the order
//	  type: string
//	billing_address:
//	  description: The order's billing address
//	  $ref: "#/components/schemas/AddressPayload"
//	shipping_address:
//	  description: The order's shipping address
//	  $ref: "#/components/schemas/AddressPayload"
//	items:
//	  description: The line items of the order
//	  type: array
//	  items:
//	    $ref: "#/components/schemas/LineItem"
//	region:
//	  description: ID of the region that the order is associated with.
//	  type: string
//	discounts:
//	  description: The discounts applied to the order
//	  type: array
//	  items:
//	    $ref: "#/components/schemas/Discount"
//	customer_id:
//	  description: The ID of the customer associated with the order.
//	  type: string
//	payment_method:
//	  description: The payment method chosen for the order.
//	  type: object
//	  properties:
//	    provider_id:
//	      type: string
//	      description: The ID of the payment provider.
//	    data:
//	      description: Any data relevant for the given payment method.
//	      type: object
//	shipping_method:
//	  description: The Shipping Method used for shipping the order.
//	  type: object
//	  properties:
//	    provider_id:
//	      type: string
//	      description: The ID of the shipping provider.
//	    profile_id:
//	      type: string
//	      description: The ID of the shipping profile.
//	    price:
//	      type: integer
//	      description: The price of the shipping.
//	    data:
//	      type: object
//	      description: Any data relevant to the specific shipping method.
//	    items:
//	      type: array
//	      items:
//	        $ref: "#/components/schemas/LineItem"
//	      description: Items to ship
//	no_notification:
//	  description: >-
//	    If set to `true`, no notification will be sent to the customer related to this order.
//	  type: boolean
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

// @oas:schema:AdminPostOrdersOrderShippingMethodsReq
// type: object
// description: "The shipping method's details."
// required:
//   - price
//   - option_id
//
// properties:
//
//	price:
//	  type: number
//	  description: The price (excluding VAT) that should be charged for the Shipping Method
//	option_id:
//	  type: string
//	  description: The ID of the Shipping Option to create the Shipping Method from.
//	data:
//	  type: object
//	  description: The data required for the Shipping Option to create a Shipping Method. This depends on the Fulfillment Provider.
type OrderShippingMethod struct {
	OptionId uuid.UUID  `json:"option_id"`
	Data     core.JSONB `json:"data,omitempty" validate:"omitempty"`
	Price    float64    `json:"price"`
}

// @oas:schema:AdminPostOrdersOrderClaimsClaimShipmentsReq
// type: object
// required:
//   - fulfillment_id
//
// properties:
//
//	fulfillment_id:
//	  description: The ID of the Fulfillment.
//	  type: string
//	tracking_numbers:
//	  description: An array of tracking numbers for the shipment.
//	  type: array
//	  items:
//	    type: string
type OrderClaimShipments struct {
	FulfillmentId   uuid.UUID `json:"fulfillment_id"`
	TrackingNumbers []string  `json:"tracking_numbers,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostOrdersOrderFulfillmentsReq
// type: object
// description: "The details of the fulfillment to be created."
// required:
//   - items
//
// properties:
//
//	items:
//	  description: The Line Items to include in the Fulfillment.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - item_id
//	      - quantity
//	    properties:
//	      item_id:
//	        description: The ID of the Line Item to fulfill.
//	        type: string
//	      quantity:
//	        description: The quantity of the Line Item to fulfill.
//	        type: integer
//	location_id:
//	  type: string
//	  description: "The ID of the location where the items will be fulfilled from."
//	no_notification:
//	  description: >-
//	    If set to `true`, no notification will be sent to the customer related to this fulfillment.
//	  type: boolean
//	metadata:
//	  description: An optional set of key-value pairs to hold additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type OrderFulfillments struct {
	Items          []FulFillmentItemType `json:"items"`
	LocationId     uuid.UUID             `json:"location_id,omitempty" validate:"omitempty"`
	NoNotification bool                  `json:"no_notification,omitempty" validate:"omitempty"`
	Metadata       core.JSONB            `json:"metadata,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminOrdersOrderLineItemReservationReq
// type: object
// required:
// - location_id
// properties:
//
//	location_id:
//	  description: "The ID of the location of the reservation"
//	  type: string
//	quantity:
//	  description: "The quantity to reserve"
//	  type: number
type OrderLineItemReservation struct {
	LocationId uuid.UUID `json:"location_id"`
	Quantity   int       `json:"quantity,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostOrdersOrderShipmentReq
// type: object
// description: "The details of the shipment to create."
// required:
//   - fulfillment_id
//
// properties:
//
//	fulfillment_id:
//	  description: The ID of the Fulfillment.
//	  type: string
//	tracking_numbers:
//	  description: The tracking numbers for the shipment.
//	  type: array
//	  items:
//	    type: string
//	no_notification:
//	  description: If set to true no notification will be send related to this Shipment.
//	  type: boolean
type CreateOrderShipment struct {
	FulfillmentId   uuid.UUID `json:"fulfillment_id"`
	TrackingNumbers []string  `json:"tracking_numbers,omitempty" validate:"omitempty"`
	NoNotification  bool      `json:"no_notification,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostOrdersOrderSwapsReq
// type: object
// description: "The details of the swap to create."
// required:
//   - return_items
//
// properties:
//
//	return_items:
//	  description: The Line Items to associate with the swap's return.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - item_id
//	      - quantity
//	    properties:
//	      item_id:
//	        description: The ID of the Line Item that will be returned.
//	        type: string
//	      quantity:
//	        description: The number of items that will be returned
//	        type: integer
//	      reason_id:
//	        description: The ID of the Return Reason to use.
//	        type: string
//	      note:
//	        description: An optional note with information about the Return.
//	        type: string
//	return_shipping:
//	  description: The shipping method associated with the swap's return.
//	  type: object
//	  required:
//	    - option_id
//	  properties:
//	    option_id:
//	      type: string
//	      description: The ID of the Shipping Option to create the Shipping Method from.
//	    price:
//	      type: integer
//	      description: The price to charge for the Shipping Method.
//	additional_items:
//	  description: The new items to send to the Customer.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - variant_id
//	      - quantity
//	    properties:
//	      variant_id:
//	        description: The ID of the Product Variant.
//	        type: string
//	      quantity:
//	        description: The quantity of the Product Variant.
//	        type: integer
//	sales_channel_id:
//	  type: string
//	  description: "The ID of the sales channel associated with the swap."
//	custom_shipping_options:
//	  description: An array of custom shipping options to potentially create a Shipping Method from to send the additional items.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - option_id
//	      - price
//	    properties:
//	      option_id:
//	        description: The ID of the Shipping Option.
//	        type: string
//	      price:
//	        description: The custom price of the Shipping Option.
//	        type: integer
//	no_notification:
//	  description: >-
//	    If set to `true`, no notification will be sent to the customer related to this Swap.
//	  type: boolean
//	return_location_id:
//	  type: string
//	  description: "The ID of the location used for the associated return."
//	allow_backorder:
//	  description: >-
//	    If set to `true`, swaps can be completed with items out of stock
//	  type: boolean
//	  default: true
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

// @oas:schema:AdminPostOrdersOrderClaimsClaimFulfillmentsReq
// type: object
// properties:
//
//	metadata:
//	  description: An optional set of key-value pairs to hold additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
//	no_notification:
//	  description: >-
//	    If set to `true`, no notification will be sent to the customer related to this Claim.
//	  type: boolean
//	location_id:
//	  description: "The ID of the fulfillment's location."
//	  type: string
type OrderClaimFulfillments struct {
	LocationId     uuid.UUID  `json:"location_id,omitempty" validate:"omitempty"`
	NoNotification bool       `json:"no_notification,omitempty" validate:"omitempty"`
	Metadata       core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostOrdersOrderRefundsReq
// type: object
// description: "The details of the order refund."
// required:
//   - amount
//   - reason
//
// properties:
//
//	amount:
//	  description: The amount to refund. It should be less than or equal the `refundable_amount` of the order.
//	  type: integer
//	reason:
//	  description: The reason for the Refund.
//	  type: string
//	note:
//	  description: A note with additional details about the Refund.
//	  type: string
//	no_notification:
//	  description: >-
//	    If set to `true`, no notification will be sent to the customer related to this Refund.
//	  type: boolean
type OrderRefunds struct {
	Amount         float64             `json:"amount"`
	Reason         models.RefundReason `json:"reason"`
	Note           string              `json:"note,omitempty" validate:"omitempty"`
	NoNotification bool                `json:"no_notification,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminPostOrdersOrderReturnsReq
// type: object
// description: "The details of the requested return."
// required:
//   - items
//
// properties:
//
//	items:
//	  description: The line items that will be returned.
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - item_id
//	      - quantity
//	    properties:
//	      item_id:
//	        description: The ID of the Line Item.
//	        type: string
//	      reason_id:
//	        description: The ID of the Return Reason to use.
//	        type: string
//	      note:
//	        description: An optional note with information about the Return.
//	        type: string
//	      quantity:
//	        description: The quantity of the Line Item.
//	        type: integer
//	return_shipping:
//	  description: The Shipping Method to be used to handle the return shipment.
//	  type: object
//	  properties:
//	    option_id:
//	      type: string
//	      description: The ID of the Shipping Option to create the Shipping Method from.
//	    price:
//	      type: integer
//	      description: The price to charge for the Shipping Method.
//	note:
//	  description: An optional note with information about the Return.
//	  type: string
//	receive_now:
//	  description: A flag to indicate if the Return should be registerd as received immediately.
//	  type: boolean
//	  default: false
//	no_notification:
//	  description: >-
//	    If set to `true`, no notification will be sent to the customer related to this Return.
//	  type: boolean
//	refund:
//	  description: The amount to refund.
//	  type: integer
//	location_id:
//	  description: "The ID of the location used for the return."
//	  type: string
type OrderReturns struct {
	Items          []OrderReturnItem              `json:"items" validate:"dive"`
	ReturnShipping CreateClaimReturnShippingInput `json:"return_shipping,omitempty" validate:"omitempty,nested"`
	Note           string                         `json:"note,omitempty" validate:"omitempty,alphanum"`
	ReceiveNow     bool                           `json:"receive_now,omitempty" validate:"omitempty,boolean"`
	NoNotification bool                           `json:"no_notification,omitempty" validate:"omitempty,boolean"`
	Refund         float64                        `json:"refund,omitempty" validate:"omitempty,numeric"`
	LocationId     uuid.UUID                      `json:"location_id,omitempty" validate:"omitempty,alphanum"`
}

// @oas:schema:StorePostCustomersCustomerAcceptClaimReq
// type: object
// description: "The details necessary to grant order access."
// required:
//   - token
//
// properties:
//
//	token:
//	  description: "The claim token generated by previous request to the Claim Order API Route."
//	  type: string
type CustomerAcceptClaim struct {
	Token string `json:"token"`
}

type OrderLookup struct {
	DisplayId       string          `json:"display_id"`
	Email           string          `json:"email" validate:"email"`
	ShippingAddress *AddressPayload `json:"shipping_address,omitempty" validate:"omitempty"`
}

// @oas:schema:StorePostCustomersCustomerOrderClaimReq
// type: object
// description: "The details of the orders to claim."
// required:
//   - order_ids
//
// properties:
//
//	order_ids:
//	  description: "The ID of the orders to claim"
//	  type: array
//	  items:
//	   type: string
type CustomerOrderClaim struct {
	OrderIds uuid.UUIDs `json:"order_ids"`
}
