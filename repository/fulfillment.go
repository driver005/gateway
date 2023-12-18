package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type FulfillmentRepo struct {
	Repository[models.Fulfillment]
}

func FulfillmentRepository(db *gorm.DB) FulfillmentRepo {
	return FulfillmentRepo{*NewRepository[models.Fulfillment](db)}
}
