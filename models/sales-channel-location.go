package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// This represents the association between a sales channel and a stock locations.
type SalesChannelLocation struct {
	core.Model

	// The ID of the Sales Channel
	SalesChannelId uuid.NullUUID `json:"sales_channel_id"`

	// The ID of the Location Stock.
	LocationId uuid.NullUUID `json:"location_id"`

	// The details of the sales channel the location is associated with.
	SalesChannel *interface{} `json:"sales_channel,omitempty" gorm:"foreignKey:id;references:location_id"`
}
