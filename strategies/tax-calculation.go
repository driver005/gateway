package strategies

import (
	"context"
	"math"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type TaxCalculationStrategy struct {
	ctx context.Context
	r   Registry
}

func NewTaxCalculationStrategy(
	r Registry,
) *TaxCalculationStrategy {
	return &TaxCalculationStrategy{
		context.Background(),
		r,
	}
}

func (s *TaxCalculationStrategy) Calculate(items []models.LineItem, taxLines []interface{}, calculationContext *interfaces.TaxCalculationContext) float64 {
	lineItemsTaxLines := filterLineItemTaxLines(taxLines)
	shippingMethodsTaxLines := filterShippingMethodTaxLines(taxLines)
	lineItemsTax := s.calculateLineItemsTax(items, lineItemsTaxLines, calculationContext)
	shippingMethodsTax := s.calculateShippingMethodsTax(calculationContext.ShippingMethods, shippingMethodsTaxLines)
	return math.Round(float64(lineItemsTax + shippingMethodsTax))
}

func (s *TaxCalculationStrategy) calculateLineItemsTax(items []models.LineItem, taxLines []models.LineItemTaxLine, context *interfaces.TaxCalculationContext) float64 {
	taxInclusivePricingFeatureFlag := true
	taxTotal := 0.0
	for _, item := range items {
		allocations := context.AllocationMap[item.Id]
		filteredTaxLines := filterTaxLinesForItem(taxLines, item.Id)
		includesTax := taxInclusivePricingFeatureFlag && item.IncludesTax
		var taxableAmount float64
		if includesTax {
			taxRate := 0.0
			for _, taxLine := range filteredTaxLines {
				taxRate += taxLine.Rate / 100
			}
			taxIncludedInPrice := math.Round(calculatePriceTaxAmount(item.UnitPrice, taxRate, includesTax))
			taxableAmount = (item.UnitPrice - taxIncludedInPrice) * float64(item.Quantity)
		} else {
			taxableAmount = item.UnitPrice * float64(item.Quantity)
		}
		taxableAmount -= allocations.Discount.Amount
		for _, filteredTaxLine := range filteredTaxLines {
			taxTotal += math.Round(calculatePriceTaxAmount(taxableAmount, filteredTaxLine.Rate/100, false))
		}
	}
	return taxTotal
}

func (s *TaxCalculationStrategy) calculateShippingMethodsTax(shippingMethods []models.ShippingMethod, taxLines []models.ShippingMethodTaxLine) float64 {
	taxInclusiveEnabled := true
	taxTotal := 0.0
	for _, sm := range shippingMethods {
		lineRates := filterTaxLinesForShippingMethod(taxLines, sm.Id)
		for _, lineRate := range lineRates {
			taxTotal += calculatePriceTaxAmount(sm.Price, lineRate.Rate/100, taxInclusiveEnabled && sm.IncludesTax)
		}
	}
	return taxTotal
}

func filterLineItemTaxLines(taxLines []interface{}) []models.LineItemTaxLine {
	var lineItemTaxLines []models.LineItemTaxLine
	for _, tl := range taxLines {
		if taxLine, ok := tl.(models.LineItemTaxLine); ok {
			lineItemTaxLines = append(lineItemTaxLines, taxLine)
		}
	}
	return lineItemTaxLines
}

func filterShippingMethodTaxLines(taxLines []interface{}) []models.ShippingMethodTaxLine {
	var shippingMethodTaxLines []models.ShippingMethodTaxLine
	for _, tl := range taxLines {
		if taxLine, ok := tl.(models.ShippingMethodTaxLine); ok {
			shippingMethodTaxLines = append(shippingMethodTaxLines, taxLine)
		}
	}
	return shippingMethodTaxLines
}

func filterTaxLinesForItem(taxLines []models.LineItemTaxLine, itemId uuid.UUID) []models.LineItemTaxLine {
	var filteredTaxLines []models.LineItemTaxLine
	for _, taxLine := range taxLines {
		if taxLine.ItemId.UUID == itemId {
			filteredTaxLines = append(filteredTaxLines, taxLine)
		}
	}
	return filteredTaxLines
}

func filterTaxLinesForShippingMethod(taxLines []models.ShippingMethodTaxLine, shippingMethodId uuid.UUID) []models.ShippingMethodTaxLine {
	var filteredTaxLines []models.ShippingMethodTaxLine
	for _, taxLine := range taxLines {
		if taxLine.ShippingMethodId.UUID == shippingMethodId {
			filteredTaxLines = append(filteredTaxLines, taxLine)
		}
	}
	return filteredTaxLines
}

func calculatePriceTaxAmount(price float64, taxRate float64, includesTax bool) float64 {
	if includesTax {
		return price - (price / (1 + taxRate))
	}
	return price * taxRate
}
