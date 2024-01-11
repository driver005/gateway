package interfaces

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/utils"
)

type ITaxCalculationStrategy interface {
	Calculate(items []models.LineItem, taxLines interface{}, calculationContext *TaxCalculationContext) (float64, *utils.ApplictaionError)
}

func IsTaxCalculationStrategy(object interface{}) bool {
	_, ok := object.(ITaxCalculationStrategy)
	return ok
}
