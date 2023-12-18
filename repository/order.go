package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type OrderRepo struct {
	Repository[models.Order]
}

func OrderRepository(db *gorm.DB) OrderRepo {
	return OrderRepo{*NewRepository[models.Order](db)}
}
