package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ShippingTaxRateRepo struct {
	sql.Repository[models.ShippingTaxRate]
}

func ShippingTaxRateRepository(db *gorm.DB) *ShippingTaxRateRepo {
	return &ShippingTaxRateRepo{*sql.NewRepository[models.ShippingTaxRate](db)}
}
