package types

import "github.com/driver005/gateway/core"

type AddressPayload struct {
	FirstName   string     `json:"first_name,omitempty" validate:"omitempty"`
	LastName    string     `json:"last_name,omitempty" validate:"omitempty"`
	Phone       string     `json:"phone,omitempty" validate:"omitempty"`
	Metadata    core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
	Company     string     `json:"company,omitempty" validate:"omitempty"`
	Address1    string     `json:"address_1,omitempty" validate:"omitempty"`
	Address2    string     `json:"address_2,omitempty" validate:"omitempty"`
	City        string     `json:"city,omitempty" validate:"omitempty"`
	CountryCode string     `json:"country_code,omitempty" validate:"omitempty"`
	Province    string     `json:"province,omitempty" validate:"omitempty"`
	PostalCode  string     `json:"postal_code,omitempty" validate:"omitempty"`
}

type AddressCreatePayload struct {
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Phone       string     `json:"phone,omitempty" validate:"omitempty"`
	Metadata    core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
	Company     string     `json:"company,omitempty" validate:"omitempty"`
	Address1    string     `json:"address_1"`
	Address2    string     `json:"address_2,omitempty" validate:"omitempty"`
	City        string     `json:"city"`
	CountryCode string     `json:"country_code"`
	Province    string     `json:"province,omitempty" validate:"omitempty"`
	PostalCode  string     `json:"postal_code"`
}
