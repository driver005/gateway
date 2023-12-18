package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ShippingOptionRepo struct {
	Repository[models.ShippingOption]
}

func ShippingOptionRepository(db *gorm.DB) ShippingOptionRepo {
	return ShippingOptionRepo{*NewRepository[models.ShippingOption](db)}
}
