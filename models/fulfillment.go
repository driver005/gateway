package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Fulfillment - Fulfillments are created once store operators can prepare the purchased goods. Fulfillments will eventually be shipped and hold information about how to track shipments. Fulfillments are created through a provider, which is typically an external shipping aggregator, shipping partner og 3PL, most plugins will have asynchronous communications with these providers through webhooks in order to automatically update and synchronize the state of Fulfillments.
type Fulfillment struct {
	core.Model

	// The id of the Claim that the Fulfillment belongs to.
	ClaimOrderId uuid.NullUUID `json:"claim_order_id" gorm:"default:null"`

	// A claim order object. Available if the relation `claim_order` is expanded.
	ClaimOrder *ClaimOrder `json:"claim_order" gorm:"foreignKey:id;references:claim_order_id"`

	// The id of the Swap that the Fulfillment belongs to.
	SwapId uuid.NullUUID `json:"swap_id" gorm:"default:null"`

	// A swap object. Available if the relation `swap` is expanded.
	Swap *Swap `json:"swap" gorm:"foreignKey:id;references:swap_id"`

	// The id of the Order that the Fulfillment belongs to.
	OrderId uuid.NullUUID `json:"order_id" gorm:"default:null"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" gorm:"foreignKey:id;references:order_id"`

	// The id of the Fulfillment Provider responsible for handling the fulfillment
	ProviderId uuid.NullUUID `json:"provider_id"`

	Provider *FulfillmentProvider `json:"provider" gorm:"foreignKey:id;references:provider_id"`

	// The Fulfillment Items in the Fulfillment - these hold information about how many of each Line Item has been fulfilled. Available if the relation `items` is expanded.
	Items []FulfillmentItem `json:"items" gorm:"foreignKey:fulfillment_id"`

	// The Tracking Links that can be used to track the status of the Fulfillment, these will usually be provided by the Fulfillment Provider. Available if the relation `tracking_links` is expanded.
	TrackingLinks []TrackingLink `json:"tracking_links" gorm:"foreignKey:id"`

	// The tracking numbers that can be used to track the status of the fulfillment.
	// Deprecated
	// TODO
	TrackingNumbers string `json:"tracking_numbers" gorm:"default:null"`

	// This contains all the data necessary for the Fulfillment provider to handle the fulfillment.
	Data JSONB `json:"data" gorm:"default:null"`

	// The date with timezone at which the Fulfillment was shipped.
	ShippedAt time.Time `json:"shipped_at" gorm:"default:null"`

	// Flag for describing whether or not notifications related to this should be send.
	NoNotification bool `json:"no_notification" gorm:"default:null"`

	// The date with timezone at which the Fulfillment was canceled.
	CanceledAt time.Time `json:"canceled_at" gorm:"default:null"`

	// Randomly generated key used to continue the completion of the fulfillment in case of failure.
	IdempotencyKey string `json:"idempotency_key" gorm:"default:null"`
}
