package models

import (
	"github.com/driver005/gateway/core"
)

// @oas:schema:FulfillmentProvider
// title: "Fulfillment Provider"
// description: "A fulfillment provider represents a fulfillment service installed in the Medusa backend, either through a plugin or backend customizations.
//
//	It holds the fulfillment service's installation status."
//
// type: object
// required:
//   - id
//   - is_installed
//
// properties:
//
//	id:
//	  description: The ID of the fulfillment provider as given by the fulfillment service.
//	  type: string
//	  example: manual
//	is_installed:
//	  description: Whether the fulfillment service is installed in the current version. If a fulfillment service is no longer installed, the `is_installed` attribute is set to `false`.
//	  type: boolean
//	  default: true
type FulfillmentProvider struct {
	core.SoftDeletableModel

	IsInstalled bool `json:"is_installed" gorm:"column:is_installed;default:true"`
}
