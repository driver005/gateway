package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type OrderItemChangeRepo struct {
	Repository[models.OrderItemChange]
}

func OrderItemChangeRepository(db *gorm.DB) OrderItemChangeRepo {
	return OrderItemChangeRepo{*NewRepository[models.OrderItemChange](db)}
}
