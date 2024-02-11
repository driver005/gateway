package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableRegion struct {
	core.FilterModel
	Name string `json:"name,omitempty" validate:"omitempty"`
}

type UpdateRegionInput struct {
	Name                 string     `json:"name,omitempty" validate:"omitempty"`
	CurrencyCode         string     `json:"currency_code,omitempty" validate:"omitempty"`
	TaxCode              string     `json:"tax_code,omitempty" validate:"omitempty"`
	TaxRate              float64    `json:"tax_rate,omitempty" validate:"omitempty"`
	GiftCardsTaxable     bool       `json:"gift_cards_taxable,omitempty" validate:"omitempty"`
	AutomaticTaxes       bool       `json:"automatic_taxes,omitempty" validate:"omitempty"`
	TaxProviderId        uuid.UUID  `json:"tax_provider_id,omitempty" validate:"omitempty"`
	PaymentProviders     uuid.UUIDs `json:"payment_providers,omitempty" validate:"omitempty"`
	FulfillmentProviders uuid.UUIDs `json:"fulfillment_providers,omitempty" validate:"omitempty"`
	Countries            []string   `json:"countries,omitempty" validate:"omitempty"`
	IncludesTax          bool       `json:"includes_tax,omitempty" validate:"omitempty"`
	Metadata             core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type CreateRegionInput struct {
	Name                 string     `json:"name"`
	CurrencyCode         string     `json:"currency_code"`
	TaxCode              string     `json:"tax_code,omitempty" validate:"omitempty"`
	TaxRate              float64    `json:"tax_rate"`
	PaymentProviders     uuid.UUIDs `json:"payment_providers"`
	FulfillmentProviders uuid.UUIDs `json:"fulfillment_providers"`
	Countries            []string   `json:"countries"`
	IncludesTax          bool       `json:"includes_tax,omitempty" validate:"omitempty"`
	Metadata             core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type RegionCountries struct {
	CountryCode string `json:"country_code"`
}

type RegionFulfillmentProvider struct {
	ProviderId uuid.UUID `json:"provider_id"`
}

type RegionPaymentProvider struct {
	ProviderId uuid.UUID `json:"provider_id"`
}
