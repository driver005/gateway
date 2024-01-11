package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type LineItemAdjustmentRepo struct {
	sql.Repository[models.LineItemAdjustment]
}

func LineItemAdjustmentRepository(db *gorm.DB) *LineItemAdjustmentRepo {
	return &LineItemAdjustmentRepo{*sql.NewRepository[models.LineItemAdjustment](db)}
}
