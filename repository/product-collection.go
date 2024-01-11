package repository

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCollectionRepo struct {
	sql.Repository[models.ProductCollection]
}

func ProductCollectionRepository(db *gorm.DB) *ProductCollectionRepo {
	return &ProductCollectionRepo{*sql.NewRepository[models.ProductCollection](db)}
}

func (r *ProductCollectionRepo) FindAndCountByDiscountConditionId(ctx context.Context, conditionId uuid.UUID, query sql.Query) ([]models.ProductCollection, *int64, *utils.ApplictaionError) {
	var result []models.ProductCollection
	var count int64

	err := r.ParseQuery(ctx, query).Model(&models.ProductCollection{}).
		Joins("JOIN discount_condition_product_collection dc_pc ON dc_pc.product_collection_id = product_collections.id AND dc_pc.condition_id = ?", conditionId).
		Set("gorm:auto_preload", true).
		Find(&result).
		Count(&count).Error

	if err != nil {
		return nil, nil, &utils.ApplictaionError{Message: err.Error()}
	}

	return result, &count, nil
}
