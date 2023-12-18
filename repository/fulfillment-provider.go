package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type FulfillmentProviderRepo struct {
	Repository[models.FulfillmentProvider]
}

func FulfillmentProviderRepository(db *gorm.DB) FulfillmentProviderRepo {
	return FulfillmentProviderRepo{*NewRepository[models.FulfillmentProvider](db)}
}
