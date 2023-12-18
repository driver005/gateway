package models

import "github.com/driver005/gateway/core"

// OAuth - Represent an OAuth app
type OAuth struct {
	core.Model

	// The app's display name
	DisplayName string `json:"display_name"`

	// The app's name
	ApplicationName string `json:"application_name"`

	// The URL to install the app
	InstallUrl string `json:"install_url" gorm:"default:null"`

	// The URL to uninstall the app
	UninstallUrl string `json:"uninstall_url" gorm:"default:null"`

	// Any data necessary to the app.
	Data JSONB `json:"data" gorm:"default:null"`
}
