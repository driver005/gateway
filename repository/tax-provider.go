package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type TaxProviderRepo struct {
	Repository[models.TaxProvider]
}

func TaxProviderRepository(db *gorm.DB) TaxProviderRepo {
	return TaxProviderRepo{*NewRepository[models.TaxProvider](db)}
}
