package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type GiftCardTransactionRepo struct {
	sql.Repository[models.GiftCardTransaction]
}

func GiftCardTransactionRepository(db *gorm.DB) *GiftCardTransactionRepo {
	return &GiftCardTransactionRepo{*sql.NewRepository[models.GiftCardTransaction](db)}
}
