package interfaces

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
)

type ShippingTaxCalculationLine struct {
	ShippingMethod models.ShippingMethod
	Rates          []types.TaxServiceRate
}

type ItemTaxCalculationLine struct {
	Item  models.LineItem
	Rates []types.TaxServiceRate
}

type TaxCalculationContext struct {
	ShippingAddress models.Address
	Customer        models.Customer
	Region          models.Region
	IsReturn        bool
	ShippingMethods []models.ShippingMethod
	AllocationMap   types.LineAllocationsMap
}

type ITaxService interface {
	GetTaxLines(itemLines []ItemTaxCalculationLine, shippingLines []ShippingTaxCalculationLine, context *TaxCalculationContext) ([]types.ProviderTaxLine, *utils.ApplictaionError)
}
