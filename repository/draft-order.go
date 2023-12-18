package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type DraftOrderRepo struct {
	Repository[models.DraftOrder]
}

func DraftOrderRepository(db *gorm.DB) DraftOrderRepo {
	return DraftOrderRepo{*NewRepository[models.DraftOrder](db)}
}
