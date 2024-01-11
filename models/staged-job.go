package models

import "github.com/driver005/gateway/core"

// StagedJob - A staged job resource
type StagedJob struct {
	core.Model

	// The name of the event
	EventName string `json:"event_name"`

	// Data necessary for the job
	Data core.JSONB `json:"data" gorm:"default:null"`

	// Options necessary for the job
	Options core.JSONB `json:"options" gorm:"default:null"`
}
