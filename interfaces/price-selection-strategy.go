package interfaces

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type PricingContext struct {
	CartId                uuid.UUID
	CustomerId            uuid.UUID
	CustomerGroupId       uuid.UUIDs
	RegionId              uuid.UUID
	Quantity              int
	CurrencyCode          string
	IncludeDiscountPrices bool
	TaxRates              []types.TaxServiceRate
	IgnoreCache           bool
	AutomaticTaxes        bool
	TaxRate               float64
}

type DefaultPriceType string

const (
	DEFAULT DefaultPriceType = "default"
)

type PriceType string

const (
	DEFAULT_PRICE PriceType = "default"
)

type PriceSelectionResult struct {
	OriginalPrice              float64
	OriginalPriceIncludesTax   bool
	CalculatedPrice            float64
	CalculatedPriceIncludesTax bool
	CalculatedPriceType        string
	Prices                     []models.MoneyAmount
}

type Pricing struct {
	VariantId uuid.UUID
	Quantity  int
}

type ProductPricing struct {
	ProductId uuid.UUID
	Variants  []models.ProductVariant
}
type IPriceSelectionStrategy interface {
	CalculateVariantPrice(data []Pricing, context *PricingContext) (map[uuid.UUID]PriceSelectionResult, *utils.ApplictaionError)

	OnVariantsPricesUpdate(variantIds uuid.UUIDs) error
}
