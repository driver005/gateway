package models

import (
	"github.com/driver005/gateway/core"
)

// Represents a fulfillment provider plugin and holds its installation status.
type FulfillmentProvider struct {
	core.Model

	// Whether the plugin is installed in the current version. Plugins that are no longer installed are not deleted by will have this field set to `false`.
	IsInstalled bool `json:"is_installed" gorm:"default:null"`
}
