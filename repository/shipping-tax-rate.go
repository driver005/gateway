package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ShippingTaxRateRepo struct {
	Repository[models.ShippingTaxRate]
}

func ShippingTaxRateRepository(db *gorm.DB) ShippingTaxRateRepo {
	return ShippingTaxRateRepo{*NewRepository[models.ShippingTaxRate](db)}
}
