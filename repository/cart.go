package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type CartRepo struct {
	Repository[models.Cart]
}

func CartRepository(db *gorm.DB) CartRepo {
	return CartRepo{*NewRepository[models.Cart](db)}
}
