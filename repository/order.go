package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type OrderRepo struct {
	sql.Repository[models.Order]
}

func OrderRepository(db *gorm.DB) *OrderRepo {
	return &OrderRepo{*sql.NewRepository[models.Order](db)}
}
