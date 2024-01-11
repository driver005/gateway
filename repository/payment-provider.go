package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type PaymentProviderRepo struct {
	sql.Repository[models.PaymentProvider]
}

func PaymentProviderRepository(db *gorm.DB) *PaymentProviderRepo {
	return &PaymentProviderRepo{*sql.NewRepository[models.PaymentProvider](db)}
}
