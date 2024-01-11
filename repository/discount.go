package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type DiscountRepo struct {
	sql.Repository[models.Discount]
}

func DiscountRepository(db *gorm.DB) *DiscountRepo {
	return &DiscountRepo{*sql.NewRepository[models.Discount](db)}
}
