package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ShippingProfileRepo struct {
	Repository[models.ShippingProfile]
}

func ShippingProfileRepository(db *gorm.DB) ShippingProfileRepo {
	return ShippingProfileRepo{*NewRepository[models.ShippingProfile](db)}
}
