package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type PaymentSessionRepo struct {
	Repository[models.PaymentSession]
}

func PaymentSessionRepository(db *gorm.DB) PaymentSessionRepo {
	return PaymentSessionRepo{*NewRepository[models.PaymentSession](db)}
}
