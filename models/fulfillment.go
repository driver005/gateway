package models

import (
	"database/sql/driver"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// @oas:schema:Fulfillment
// title: "Fulfillment"
// description: "A Fulfillment is created once an admin can prepare the purchased goods. Fulfillments will eventually be shipped and hold information about how to track shipments. Fulfillments are created through a fulfillment provider, which typically integrates a third-party shipping service. Fulfillments can be associated with orders, claims, swaps, and returns."
// type: object
// required:
//   - canceled_at
//   - claim_order_id
//   - created_at
//   - data
//   - id
//   - idempotency_key
//   - location_id
//   - metadata
//   - no_notification
//   - order_id
//   - provider_id
//   - shipped_at
//   - swap_id
//   - tracking_numbers
//   - updated_at
//
// properties:
//
//	id:
//	  description: The fulfillment's ID
//	  type: string
//	  example: ful_01G8ZRTMQCA76TXNAT81KPJZRF
//	claim_order_id:
//	  description: The ID of the Claim that the Fulfillment belongs to.
//	  nullable: true
//	  type: string
//	  example: null
//	claim_order:
//	  description: The details of the claim that the fulfillment may belong to.
//	  x-expandable: "claim_order"
//	  nullable: true
//	  $ref: "#/components/schemas/ClaimOrder"
//	swap_id:
//	  description: The ID of the Swap that the Fulfillment belongs to.
//	  nullable: true
//	  type: string
//	  example: null
//	swap:
//	  description: The details of the swap that the fulfillment may belong to.
//	  x-expandable: "swap"
//	  nullable: true
//	  $ref: "#/components/schemas/Swap"
//	order_id:
//	  description: The ID of the Order that the Fulfillment belongs to.
//	  nullable: true
//	  type: string
//	  example: order_01G8TJSYT9M6AVS5N4EMNFS1EK
//	order:
//	  description: The details of the order that the fulfillment may belong to.
//	  x-expandable: "order"
//	  nullable: true
//	  $ref: "#/components/schemas/Order"
//	provider_id:
//	  description: The ID of the Fulfillment Provider responsible for handling the fulfillment.
//	  type: string
//	  example: manual
//	provider:
//	  description: The details of the fulfillment provider responsible for handling the fulfillment.
//	  x-expandable: "provider"
//	  nullable: true
//	  $ref: "#/components/schemas/FulfillmentProvider"
//	location_id:
//	  description: The ID of the stock location the fulfillment will be shipped from
//	  nullable: true
//	  type: string
//	  example: sloc_01G8TJSYT9M6AVS5N4EMNFS1EK
//	items:
//	  description: The Fulfillment Items in the Fulfillment. These hold information about how many of each Line Item has been fulfilled.
//	  type: array
//	  x-expandable: "items"
//	  items:
//	    $ref: "#/components/schemas/FulfillmentItem"
//	tracking_links:
//	  description: The Tracking Links that can be used to track the status of the Fulfillment. These will usually be provided by the Fulfillment Provider.
//	  type: array
//	  x-expandable: "tracking_links"
//	  items:
//	    $ref: "#/components/schemas/TrackingLink"
//	tracking_numbers:
//	  description: The tracking numbers that can be used to track the status of the fulfillment.
//	  deprecated: true
//	  type: array
//	  items:
//	    type: string
//	data:
//	  description: This contains all the data necessary for the Fulfillment provider to handle the fulfillment.
//	  type: object
//	  example: {}
//	shipped_at:
//	  description: The date with timezone at which the Fulfillment was shipped.
//	  nullable: true
//	  type: string
//	  format: date-time
//	no_notification:
//	  description: Flag for describing whether or not notifications related to this should be sent.
//	  nullable: true
//	  type: boolean
//	  example: false
//	canceled_at:
//	  description: The date with timezone at which the Fulfillment was canceled.
//	  nullable: true
//	  type: string
//	  format: date-time
//	idempotency_key:
//	  description: Randomly generated key used to continue the completion of the fulfillment in case of failure.
//	  nullable: true
//	  type: string
//	  externalDocs:
//	    url: https://docs.medusajs.com/development/idempotency-key/overview.md
//	    description: Learn more how to use the idempotency key.
//	created_at:
//	  description: The date with timezone at which the resource was created.
//	  type: string
//	  format: date-time
//	updated_at:
//	  description: The date with timezone at which the resource was updated.
//	  type: string
//	  format: date-time
//	metadata:
//	  description: An optional key-value map with additional details
//	  nullable: true
//	  type: object
//	  example: {car: "white"}
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type Fulfillment struct {
	core.Model

	ClaimOrderId uuid.NullUUID        `json:"claim_order_id" gorm:"column:claim_order_id"`
	ClaimOrder   *ClaimOrder          `json:"claim_order" gorm:"foreignKey:ClaimOrderId"`
	SwapId       uuid.NullUUID        `json:"swap_id" gorm:"column:swap_id"`
	Swap         *Swap                `json:"swap" gorm:"foreignKey:SwapId"`
	OrderId      uuid.NullUUID        `json:"order_id" gorm:"column:order_id"`
	Order        *Order               `json:"order" gorm:"foreignKey:OrderId"`
	ProviderId   uuid.NullUUID        `json:"provider_id" gorm:"column:provider_id"`
	Provider     *FulfillmentProvider `json:"provider" gorm:"foreignKey:ProviderId"`
	LocationId   uuid.NullUUID        `json:"location_id" gorm:"column:location_id"`
	//TODO:ADD
	Items         []FulfillmentItem `json:"items" gorm:"foreignKey:FulfillmentId"`
	TrackingLinks []TrackingLink    `json:"tracking_links" gorm:"foreignKey:Id"`

	//TODO: add ;default:[]
	TrackingNumbers pq.StringArray `json:"tracking_numbers" gorm:"column:tracking_numbers;type:text[];default:[]"`
	Data            core.JSONB     `json:"data" gorm:"column:data"`
	ShippedAt       *time.Time     `json:"shipped_at" gorm:"column:shipped_at"`
	NoNotification  bool           `json:"no_notification" gorm:"column:no_notification"`
	CanceledAt      *time.Time     `json:"canceled_at" gorm:"column:canceled_at"`
	IdempotencyKey  string         `json:"idempotency_key" gorm:"column:idempotency_key"`
}

type FulfillmentStatus string

const (
	FulfillmentStatusNotFulfilled       FulfillmentStatus = "not_fulfilled"
	FulfillmentStatusPartiallyFulfilled FulfillmentStatus = "partially_fulfilled"
	FulfillmentStatusFulfilled          FulfillmentStatus = "fulfilled"
	FulfillmentStatusPartiallyShipped   FulfillmentStatus = "partially_shipped"
	FulfillmentStatusShipped            FulfillmentStatus = "shipped"
	FulfillmentStatusPartiallyReturned  FulfillmentStatus = "partially_returned"
	FulfillmentStatusReturned           FulfillmentStatus = "returned"
	FulfillmentStatusCanceled           FulfillmentStatus = "canceled"
	FulfillmentStatusRequiresAction     FulfillmentStatus = "requires_action"
)

func (pl *FulfillmentStatus) Scan(value interface{}) error {
	*pl = FulfillmentStatus(value.([]byte))
	return nil
}

func (pl FulfillmentStatus) Value() (driver.Value, error) {
	return string(pl), nil
}
