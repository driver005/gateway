package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type DraftOrderRepo struct {
	sql.Repository[models.DraftOrder]
}

func DraftOrderRepository(db *gorm.DB) *DraftOrderRepo {
	return &DraftOrderRepo{*sql.NewRepository[models.DraftOrder](db)}
}
