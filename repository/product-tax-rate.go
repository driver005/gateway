package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ProductTaxRateRepo struct {
	sql.Repository[models.ProductTaxRate]
}

func ProductTaxRateRepository(db *gorm.DB) *ProductTaxRateRepo {
	return &ProductTaxRateRepo{*sql.NewRepository[models.ProductTaxRate](db)}
}
