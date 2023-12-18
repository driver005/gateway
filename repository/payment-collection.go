package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

// TODO: Add
type PaymentCollectionRepo struct {
	Repository[models.PaymentCollection]
}

func PaymentCollectionRepository(db *gorm.DB) PaymentCollectionRepo {
	return PaymentCollectionRepo{*NewRepository[models.PaymentCollection](db)}
}
