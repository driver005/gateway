package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type CustomShippingOptionRepo struct {
	sql.Repository[models.CustomShippingOption]
}

func CustomShippingOptionRepository(db *gorm.DB) *CustomShippingOptionRepo {
	return &CustomShippingOptionRepo{*sql.NewRepository[models.CustomShippingOption](db)}
}
