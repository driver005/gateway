package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type LineItemAdjustmentRepo struct {
	Repository[models.LineItemAdjustment]
}

func LineItemAdjustmentRepository(db *gorm.DB) LineItemAdjustmentRepo {
	return LineItemAdjustmentRepo{*NewRepository[models.LineItemAdjustment](db)}
}
