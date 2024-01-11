package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type FulfillmentProviderRepo struct {
	sql.Repository[models.FulfillmentProvider]
}

func FulfillmentProviderRepository(db *gorm.DB) *FulfillmentProviderRepo {
	return &FulfillmentProviderRepo{*sql.NewRepository[models.FulfillmentProvider](db)}
}
