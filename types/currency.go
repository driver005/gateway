package types

import "github.com/driver005/gateway/core"

type UpdateCurrencyInput struct {
	IncludesTax bool `json:"includes_tax,omitempty" validate:"omitempty"`
}

type FilterableCurrencyProps struct {
	core.FilterModel
	Code        string `json:"code" validate:"omitempty,string"`
	IncludesTax bool   `json:"includes_tax" validate:"omitempty,bool"`
}
