package models

import "github.com/google/uuid"

// PaymentProvider - Represents a Payment Provider plugin and holds its installation status.
type PaymentProvider struct {

	// The id of the payment provider as given by the plugin.
	Id uuid.UUID `json:"id gorm:primarykey"`

	// Whether the plugin is installed in the current version. Plugins that are no longer installed are not deleted by will have this field set to `false`.
	IsInstalled bool `json:"is_installed" gorm:"default:null"`
}
