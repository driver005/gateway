package models

import "github.com/driver005/gateway/core"

// TaxLine - Line item that specifies an amount of tax to add to a line item.
type TaxLine struct {
	core.Model

	// A code to identify the tax type by
	Code string `json:"code" gorm:"default:null"`

	// A human friendly name for the tax
	Name string `json:"name"`

	// The numeric rate to charge tax by
	Rate float64 `json:"rate"`
}
