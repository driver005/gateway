package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type PriceListRepo struct {
	Repository[models.PriceList]
}

func PriceListRepository(db *gorm.DB) PriceListRepo {
	return PriceListRepo{*NewRepository[models.PriceList](db)}
}
