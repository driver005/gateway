package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type GiftCardRepo struct {
	sql.Repository[models.GiftCard]
}

func GiftCardRepository(db *gorm.DB) *GiftCardRepo {
	return &GiftCardRepo{*sql.NewRepository[models.GiftCard](db)}
}
