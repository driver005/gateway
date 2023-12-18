package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ProductTaxRateRepo struct {
	Repository[models.ProductTaxRate]
}

func ProductTaxRateRepository(db *gorm.DB) ProductTaxRateRepo {
	return ProductTaxRateRepo{*NewRepository[models.ProductTaxRate](db)}
}
