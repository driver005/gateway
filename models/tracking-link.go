package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// TrackingLink - Tracking Link holds information about tracking numbers for a Fulfillment. Tracking Links can optionally contain a URL that can be visited to see the status of the shipment.
type TrackingLink struct {
	core.Model

	// The URL at which the status of the shipment can be tracked.
	Url string `json:"url" gorm:"default:null"`

	// The tracking number given by the shipping carrier.
	TrackingNumber string `json:"tracking_number"`

	// The id of the Fulfillment that the Tracking Link references.
	FulfillmentId uuid.NullUUID `json:"fulfillment_id"`

	Fulfillment *Fulfillment `json:"fulfillment" gorm:"foreignKey:id;references:fulfillment_id"`

	// Randomly generated key used to continue the completion of a process in case of failure.
	IdempotencyKey string `json:"idempotency_key" gorm:"default:null"`
}
