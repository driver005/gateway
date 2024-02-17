package types

import "github.com/driver005/gateway/core"

// @oas:schema:AdminPostCurrenciesCurrencyReq
// type: object
// description: "The details to update in the currency"
// properties:
//
//	includes_tax:
//	  type: boolean
//	  x-featureFlag: "tax_inclusive_pricing"
//	  description: "Tax included in prices of currency."
type UpdateCurrencyInput struct {
	IncludesTax bool `json:"includes_tax,omitempty" validate:"omitempty"`
}

type FilterableCurrencyProps struct {
	core.FilterModel
	Code        string `json:"code" validate:"omitempty,string"`
	IncludesTax bool   `json:"includes_tax" validate:"omitempty,bool"`
}
