package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type FulfillmentRepo struct {
	sql.Repository[models.Fulfillment]
}

func FulfillmentRepository(db *gorm.DB) *FulfillmentRepo {
	return &FulfillmentRepo{*sql.NewRepository[models.Fulfillment](db)}
}
