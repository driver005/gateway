package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// SalesChannel - A Sales Channel
type SalesChannel struct {
	core.Model

	// The name of the sales channel.
	Name string `json:"name"`

	// The description of the sales channel.
	Description string `json:"description" gorm:"default:null"`

	LocationIds uuid.UUIDs `json:"location_id"`

	// Specify if the sales channel is enabled or disabled.
	IsDisabled bool `json:"is_disabled" gorm:"default:null"`
}
