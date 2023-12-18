package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type GiftCardRepo struct {
	Repository[models.GiftCard]
}

func GiftCardRepository(db *gorm.DB) GiftCardRepo {
	return GiftCardRepo{*NewRepository[models.GiftCard](db)}
}
