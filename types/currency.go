package types

type UpdateCurrencyInput struct {
	IncludesTax bool `json:"includes_tax,omitempty" validate:"omitempty"`
}
