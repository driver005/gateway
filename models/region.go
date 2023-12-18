package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Region - Regions hold settings for how Customers in a given geographical location shop. The is, for example, where currencies and tax rates are defined. A Region can consist of multiple countries to accomodate common shopping settings across countries.
type Region struct {
	core.Model

	// The name of the region as displayed to the customer. If the Region only has one country it is recommended to write the country name.
	Name string `json:"name"`

	// The 3 character currency code that the Region uses.
	CurrencyCode string `json:"currency_code"`

	Currency *Currency `json:"currency" gorm:"foreignKey:code;references:currency_code"`

	// The tax rate that should be charged on purchases in the Region.
	TaxRate float64 `json:"tax_rate"`

	// The tax rates that are included in the Region. Available if the relation `tax_rates` is expanded.
	TaxRates []TaxRate `json:"tax_rates" gorm:"foreignKey:id"`

	// The tax code used on purchases in the Region. This may be used by other systems for accounting purposes.
	TaxCode string `json:"tax_code" gorm:"default:null"`

	// Whether the gift cards are taxable or not in this region.
	GiftCardsTaxable bool `json:"gift_cards_taxable" gorm:"default:null"`

	// Whether taxes should be automated in this region.
	AutomaticTaxes bool `json:"automatic_taxes" gorm:"default:null"`

	// The countries that are included in the Region. Available if the relation `countries` is expanded.
	Countries []Country `json:"countries" gorm:"foreignKey:id"`

	// The ID of the tax provider used in this region
	TaxProviderId uuid.NullUUID `json:"tax_provider_id" gorm:"default:null"`

	TaxProvider *TaxProvider `json:"tax_provider" gorm:"foreignKey:id;references:tax_provider_id"`

	// The Payment Providers that can be used to process Payments in the Region. Available if the relation `payment_providers` is expanded.
	PaymentProviders []PaymentProvider `json:"payment_providers" gorm:"foreignKey:id"`

	// The Fulfillment Providers that can be used to fulfill orders in the Region. Available if the relation `payment_providers` is expanded.
	FulfillmentProviders []FulfillmentProvider `json:"fulfillment_providers" gorm:"foreignKey:id"`

	// [EXPERIMENTAL] Does the prices for the region include tax
	IncludesTax bool `json:"includes_tax" gorm:"default:null"`
}
