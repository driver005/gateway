package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type ProductVariantPricing struct {
	TaxedPricing
	Prices                     []models.MoneyAmount `json:"prices"`
	OriginalPrice              float64              `json:"original_price,omitempty" validate:"omitempty"`
	CalculatedPrice            float64              `json:"calculated_price,omitempty" validate:"omitempty"`
	OriginalPriceIncludesTax   bool                 `json:"original_price_includes_tax,omitempty" validate:"omitempty"`
	CalculatedPriceIncludesTax bool                 `json:"calculated_price_includes_tax,omitempty" validate:"omitempty"`
	CalculatedPriceType        string               `json:"calculated_price_type,omitempty" validate:"omitempty"`
}

type TaxedPricing struct {
	OriginalPriceInclTax   float64          `json:"original_price_incl_tax,omitempty" validate:"omitempty"`
	CalculatedPriceInclTax float64          `json:"calculated_price_incl_tax,omitempty" validate:"omitempty"`
	OriginalTax            float64          `json:"original_tax,omitempty" validate:"omitempty"`
	CalculatedTax          float64          `json:"calculated_tax,omitempty" validate:"omitempty"`
	TaxRates               []TaxServiceRate `json:"tax_rates,omitempty" validate:"omitempty"`
}

type PricingContext struct {
	Context core.JSONB `json:"context"`
}

type PricingFilters struct {
	Id uuid.UUIDs `json:"id"`
}

type PriceSetDTO struct {
	Id           uuid.UUID            `json:"id"`
	MoneyAmounts []models.MoneyAmount `json:"money_amounts,omitempty" validate:"omitempty"`
	RuleTypes    []RuleTypeDTO        `json:"rule_types,omitempty" validate:"omitempty"`
}

type CalculatedPriceSetDTO struct {
	Id            uuid.UUID `json:"id"`
	PriceSetId    uuid.UUID `json:"price_set_id"`
	Amount        string    `json:"amount,omitempty" validate:"omitempty"`
	CurrencyCode  string    `json:"currency_code,omitempty" validate:"omitempty"`
	MinQuantity   string    `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity   string    `json:"max_quantity,omitempty" validate:"omitempty"`
	PriceListType string    `json:"price_list_type,omitempty" validate:"omitempty"`
	PriceListId   uuid.UUID `json:"price_list_id,omitempty" validate:"omitempty"`
}

type CalculatedPriceSet struct {
	Id                         uuid.UUID        `json:"id"`
	IsCalculatedPricePriceList bool             `json:"is_calculated_price_price_list,omitempty" validate:"omitempty"`
	CalculatedAmount           float64          `json:"calculated_amount,omitempty" validate:"omitempty"`
	IsOriginalPricePriceList   bool             `json:"is_original_price_price_list,omitempty" validate:"omitempty"`
	OriginalAmount             float64          `json:"original_amount,omitempty" validate:"omitempty"`
	CurrencyCode               string           `json:"currency_code,omitempty" validate:"omitempty"`
	CalculatedPrice            *CalculatedPrice `json:"calculated_price,omitempty" validate:"omitempty"`
	OriginalPrice              *OriginalPrice   `json:"original_price,omitempty" validate:"omitempty"`
}

type CalculatedPrice struct {
	MoneyAmountId uuid.UUID `json:"money_amount_id,omitempty" validate:"omitempty"`
	PriceListId   uuid.UUID `json:"price_list_id,omitempty" validate:"omitempty"`
	PriceListType string    `json:"price_list_type,omitempty" validate:"omitempty"`
	MinQuantity   int       `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity   int       `json:"max_quantity,omitempty" validate:"omitempty"`
}

type OriginalPrice struct {
	MoneyAmountId uuid.UUID `json:"money_amount_id,omitempty" validate:"omitempty"`
	PriceListId   uuid.UUID `json:"price_list_id,omitempty" validate:"omitempty"`
	PriceListType string    `json:"price_list_type,omitempty" validate:"omitempty"`
	MinQuantity   int       `json:"min_quantity,omitempty" validate:"omitempty"`
	MaxQuantity   int       `json:"max_quantity,omitempty" validate:"omitempty"`
}

type AddRulesDTO struct {
	PriceSetId uuid.UUID `json:"priceSetId uuid.UUID"`
	Rules      []AddRule `json:"rules"`
}

type AddRule struct {
	Attribute string `json:"attribute"`
}

type CreatePricesDTO struct {
	*models.MoneyAmount
	Rules map[string]string `json:"rules"`
}

type AddPricesDTO struct {
	PriceSetId uuid.UUID         `json:"priceSetId uuid.UUID"`
	Prices     []CreatePricesDTO `json:"prices"`
}

type RemovePriceSetRulesDTO struct {
	Id    uuid.UUID `json:"id"`
	Rules []string  `json:"rules"`
}

type CreatePriceSetDTO struct {
	Rules  []CreateRule      `json:"rules,omitempty" validate:"omitempty"`
	Prices []CreatePricesDTO `json:"prices,omitempty" validate:"omitempty"`
}

type CreateRule struct {
	RuleAttribute string `json:"rule_attribute"`
}

type UpdatePriceSetDTO struct {
	Id uuid.UUID `json:"id"`
}

type FilterablePriceSetProps struct {
	Id           uuid.UUIDs          `json:"id,omitempty" validate:"omitempty"`
	MoneyAmounts *models.MoneyAmount `json:"money_amounts,omitempty" validate:"omitempty"`
}

type PricedShippingOption struct {
	*models.ShippingOption
	PriceInclTax float64          `json:"price_incl_tax"`
	TaxRates     []TaxServiceRate `json:"tax_rates"`
	TaxAmount    float64          `json:"tax_amount"`
}
