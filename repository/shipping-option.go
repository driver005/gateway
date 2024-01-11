package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ShippingOptionRepo struct {
	sql.Repository[models.ShippingOption]
}

func ShippingOptionRepository(db *gorm.DB) *ShippingOptionRepo {
	return &ShippingOptionRepo{*sql.NewRepository[models.ShippingOption](db)}
}
