package models

import (
	"github.com/driver005/gateway/core"
)

// @oas:schema:PaymentProvider
// title: "Payment Provider"
// description: "A payment provider represents a payment service installed in the Medusa backend, either through a plugin or backend customizations.
//
//	It holds the payment service's installation status."
//
// type: object
// required:
//   - id
//   - is_installed
//
// properties:
//
//	id:
//	  description: The ID of the payment provider as given by the payment service.
//	  type: string
//	  example: manual
//	is_installed:
//	  description: Whether the payment service is installed in the current version. If a payment service is no longer installed, the `is_installed` attribute is set to `false`.
//	  type: boolean
//	  default: true
type PaymentProvider struct {
	core.Model

	IsInstalled bool `json:"is_installed"  gorm:"column:is_installed;default:true"`
}
