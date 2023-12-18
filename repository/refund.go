package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type RefundRepo struct {
	Repository[models.Refund]
}

func RefundRepository(db *gorm.DB) RefundRepo {
	return RefundRepo{*NewRepository[models.Refund](db)}
}
