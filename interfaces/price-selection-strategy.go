package interfaces

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/gofrs/uuid"
)

type PriceSelectionContext struct {
	CartID                string
	CustomerID            string
	RegionID              string
	Quantity              int
	CurrencyCode          string
	IncludeDiscountPrices bool
	TaxRates              []types.TaxServiceRate
	IgnoreCache           bool
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
	OriginalPrice              *float64
	OriginalPriceIncludesTax   *bool
	CalculatedPrice            *float64
	CalculatedPriceIncludesTax *bool
	CalculatedPriceType        *PriceType
	Prices                     []models.MoneyAmount
}

type IPriceSelectionStrategy interface {
	CalculateVariantPrice(data []struct {
		VariantID uuid.UUID
		Quantity  *int
	}, context PriceSelectionContext) (map[string]PriceSelectionResult, error)

	OnVariantsPricesUpdate(variantIDs []string) error
}
