package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductTypeRepo struct {
	sql.Repository[models.ProductType]
}

func ProductTypeRepository(db *gorm.DB) *ProductTypeRepo {
	return &ProductTypeRepo{*sql.NewRepository[models.ProductType](db)}
}

func (r *ProductTypeRepo) UpsertType(t *models.ProductType) (*models.ProductType, *utils.ApplictaionError) {
	if t == nil {
		return nil, nil
	}

	var existing models.ProductType
	err := r.Db().Where("value = ?", t.Value).First(&existing).Error
	if err == nil {
		return &existing, nil
	}

	created := &models.ProductType{
		Value: t.Value,
	}
	err = r.Db().Create(created).Error
	if err != nil {
		return nil, r.HandleDBError(err)
	}

	return created, nil
}

func (r *ProductTypeRepo) FindAndCountByDiscountConditionId(conditionId uuid.UUID, query sql.Query) ([]models.ProductType, int64, *utils.ApplictaionError) {
	var pt []models.ProductType
	err := r.Db().Model(&models.ProductType{}).
		Where(query.Where).
		Joins("INNER JOIN discount_condition_product_type dc_pt ON dc_pt.product_type_id = product_types.id AND dc_pt.condition_id = ?", conditionId).
		Find(&pt).Error
	if err != nil {
		return nil, 0, r.HandleDBError(err)
	}

	var count int64
	err = r.Db().Model(&models.ProductType{}).
		Where(query.Where).
		Joins("INNER JOIN discount_condition_product_type dc_pt ON dc_pt.product_type_id = product_types.id AND dc_pt.condition_id = ?", conditionId).
		Count(&count).Error
	if err != nil {
		return nil, 0, r.HandleDBError(err)
	}

	return pt, count, nil
}
