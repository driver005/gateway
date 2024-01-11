package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type TaxProviderRepo struct {
	sql.Repository[models.TaxProvider]
}

func TaxProviderRepository(db *gorm.DB) *TaxProviderRepo {
	return &TaxProviderRepo{*sql.NewRepository[models.TaxProvider](db)}
}
