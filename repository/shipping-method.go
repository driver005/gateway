package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ShippingMethodRepo struct {
	sql.Repository[models.ShippingMethod]
}

func ShippingMethodRepository(db *gorm.DB) *ShippingMethodRepo {
	return &ShippingMethodRepo{*sql.NewRepository[models.ShippingMethod](db)}
}
