package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ShippingMethodRepo struct {
	Repository[models.ShippingMethod]
}

func ShippingMethodRepository(db *gorm.DB) ShippingMethodRepo {
	return ShippingMethodRepo{*NewRepository[models.ShippingMethod](db)}
}
