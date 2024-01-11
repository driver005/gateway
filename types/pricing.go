package types

import (
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type ProductVariantPricing struct {
	Prices                     []models.MoneyAmount
	OriginalPrice              float64
	CalculatedPrice            float64
	OriginalPriceIncludesTax   bool
	CalculatedPriceIncludesTax bool
	CalculatedPriceType        string
	TaxedPricing
}

type TaxedPricing struct {
	OriginalPriceInclTax   float64
	CalculatedPriceInclTax float64
	OriginalTax            float64
	CalculatedTax          float64
	TaxRates               []TaxServiceRate
}

type PricingContext struct {
	Context map[string]interface{} `json:"context"`
}

type PricingFilters struct {
	Id []string `json:"id"`
}

type PriceSetDTO struct {
	Id           uuid.UUID            `json:"id"`
	MoneyAmounts []models.MoneyAmount `json:"money_amounts,omitempty"`
	RuleTypes    []RuleTypeDTO        `json:"rule_types,omitempty"`
}

type CalculatedPriceSetDTO struct {
	Id            uuid.UUID `json:"id"`
	PriceSetID    string    `json:"price_set_id"`
	Amount        string    `json:"amount,omitempty"`
	CurrencyCode  string    `json:"currency_code,omitempty"`
	MinQuantity   string    `json:"min_quantity,omitempty"`
	MaxQuantity   string    `json:"max_quantity,omitempty"`
	PriceListType string    `json:"price_list_type,omitempty"`
	PriceListID   string    `json:"price_list_id,omitempty"`
}

type CalculatedPriceSet struct {
	Id                         uuid.UUID        `json:"id"`
	IsCalculatedPricePriceList bool             `json:"is_calculated_price_price_list,omitempty"`
	CalculatedAmount           float64          `json:"calculated_amount,omitempty"`
	IsOriginalPricePriceList   bool             `json:"is_original_price_price_list,omitempty"`
	OriginalAmount             float64          `json:"original_amount,omitempty"`
	CurrencyCode               string           `json:"currency_code,omitempty"`
	CalculatedPrice            *CalculatedPrice `json:"calculated_price,omitempty"`
	OriginalPrice              *OriginalPrice   `json:"original_price,omitempty"`
}

type CalculatedPrice struct {
	MoneyAmountId uuid.UUID `json:"money_amount_id,omitempty"`
	PriceListId   uuid.UUID `json:"price_list_id,omitempty"`
	PriceListType string    `json:"price_list_type,omitempty"`
	MinQuantity   int       `json:"min_quantity,omitempty"`
	MaxQuantity   int       `json:"max_quantity,omitempty"`
}

type OriginalPrice struct {
	MoneyAmountId uuid.UUID `json:"money_amount_id,omitempty"`
	PriceListId   uuid.UUID `json:"price_list_id,omitempty"`
	PriceListType string    `json:"price_list_type,omitempty"`
	MinQuantity   int       `json:"min_quantity,omitempty"`
	MaxQuantity   int       `json:"max_quantity,omitempty"`
}

type AddRulesDTO struct {
	PriceSetId uuid.UUID `json:"priceSetId"`
	Rules      []AddRule `json:"rules"`
}

type AddRule struct {
	Attribute string `json:"attribute"`
}

type CreatePricesDTO struct {
	models.MoneyAmount
	Rules map[string]string `json:"rules"`
}

type AddPricesDTO struct {
	PriceSetId uuid.UUID         `json:"priceSetId"`
	Prices     []CreatePricesDTO `json:"prices"`
}

type RemovePriceSetRulesDTO struct {
	Id    uuid.UUID `json:"id"`
	Rules []string  `json:"rules"`
}

type CreatePriceSetDTO struct {
	Rules  []CreateRule      `json:"rules,omitempty"`
	Prices []CreatePricesDTO `json:"prices,omitempty"`
}

type CreateRule struct {
	RuleAttribute string `json:"rule_attribute"`
}

type UpdatePriceSetDTO struct {
	Id string `json:"id"`
}

type FilterablePriceSetProps struct {
	Id           []string            `json:"id,omitempty"`
	MoneyAmounts *models.MoneyAmount `json:"money_amounts,omitempty"`
}

type PricedShippingOption struct {
	models.ShippingOption
	PriceInclTax float64          `json:"price_incl_tax"`
	TaxRates     []TaxServiceRate `json:"tax_rates"`
	TaxAmount    float64          `json:"tax_amount"`
}
