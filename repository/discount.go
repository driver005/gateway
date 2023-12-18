package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type DiscountRepo struct {
	Repository[models.Discount]
}

func DiscountRepository(db *gorm.DB) DiscountRepo {
	return DiscountRepo{*NewRepository[models.Discount](db)}
}
