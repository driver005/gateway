package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type OrderItemChangeRepo struct {
	sql.Repository[models.OrderItemChange]
}

func OrderItemChangeRepository(db *gorm.DB) *OrderItemChangeRepo {
	return &OrderItemChangeRepo{*sql.NewRepository[models.OrderItemChange](db)}
}
