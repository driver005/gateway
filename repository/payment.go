package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type PaymentRepo struct {
	Repository[models.Payment]
}

func PaymentRepository(db *gorm.DB) PaymentRepo {
	return PaymentRepo{*NewRepository[models.Payment](db)}
}
