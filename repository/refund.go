package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type RefundRepo struct {
	sql.Repository[models.Refund]
}

func RefundRepository(db *gorm.DB) *RefundRepo {
	return &RefundRepo{*sql.NewRepository[models.Refund](db)}
}
