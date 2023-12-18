package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type PaymentProviderRepo struct {
	Repository[models.PaymentProvider]
}

func PaymentProviderRepository(db *gorm.DB) PaymentProviderRepo {
	return PaymentProviderRepo{*NewRepository[models.PaymentProvider](db)}
}
