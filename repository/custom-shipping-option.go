package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type CustomShippingOptionRepo struct {
	Repository[models.CustomShippingOption]
}

func CustomShippingOptionRepository(db *gorm.DB) CustomShippingOptionRepo {
	return CustomShippingOptionRepo{*NewRepository[models.CustomShippingOption](db)}
}
