package interfaces

import "github.com/driver005/gateway/models"

type ITaxCalculationStrategy interface {
	Calculate(items []models.LineItem, taxLines interface{}, calculationContext *TaxCalculationContext) (float64, error)
}

func IsTaxCalculationStrategy(object interface{}) bool {
	_, ok := object.(ITaxCalculationStrategy)
	return ok
}
