package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type GiftCardTransactionRepo struct {
	Repository[models.GiftCardTransaction]
}

func GiftCardTransactionRepository(db *gorm.DB) GiftCardTransactionRepo {
	return GiftCardTransactionRepo{*NewRepository[models.GiftCardTransaction](db)}
}
