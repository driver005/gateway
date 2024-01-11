package repository

import (
	"context"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductCategoryRepo struct {
	sql.Repository[models.ProductCategory]
}

func ProductCategoryRepository(db *gorm.DB) *ProductCategoryRepo {
	return &ProductCategoryRepo{*sql.NewRepository[models.ProductCategory](db)}
}

func (r *ProductCategoryRepo) FindOneWithDescendants(ctx context.Context, query sql.Query, q *string) (*models.ProductCategory, *utils.ApplictaionError) {
	var productCategory models.ProductCategory
	err := r.Db().Preload(clause.Associations).First(&productCategory, query).Error
	if err != nil {
		return nil, r.HandleDBError(err)
	}
	return &productCategory, nil
}

func (r *ProductCategoryRepo) GetFreeTextSearchResultsAndCount(ctx context.Context, query sql.Query, q *string, includeTree bool) ([]models.ProductCategory, *int64, *utils.ApplictaionError) {
	var categories []models.ProductCategory
	var count int64

	db := r.ParseQuery(ctx, query)

	if q != nil && *q != "" {
		db = db.Where("name LIKE ?", "%"+*q+"%")
		db = db.Or("handle LIKE ?", "%"+*q+"%")
	}

	if err := db.
		Preload(clause.Associations).
		Order("rank ASC").
		Order("handle ASC").
		Find(&categories).Error; err != nil {
		return nil, nil, r.HandleDBError(err)
	}

	if includeTree {
		for i := range categories {
			if err := db.Preload("Children").First(&categories[i]).Error; err != nil {
				return nil, nil, r.HandleDBError(err)
			}
		}
	}

	if err := db.Model(&models.ProductCategory{}).Count(&count).Error; err != nil {
		return nil, nil, r.HandleDBError(err)
	}

	return categories, &count, nil
}

func (r *ProductCategoryRepo) AddProducts(productCategoryId uuid.UUID, productIds uuid.UUIDs) *utils.ApplictaionError {
	for _, productId := range productIds {
		err := r.Db().Model(&models.ProductCategory{}).Association("Products").Append(&models.Product{Model: core.Model{Id: productId}})
		if err != nil {
			return r.HandleDBError(err)
		}
	}
	return nil
}

func (r *ProductCategoryRepo) RemoveProducts(productCategoryId uuid.UUID, productIds uuid.UUIDs) *utils.ApplictaionError {
	for _, productId := range productIds {
		err := r.Db().Model(&models.ProductCategory{}).Association("Products").Delete(&models.Product{Model: core.Model{Id: productId}})
		if err != nil {
			return r.HandleDBError(err)
		}
	}
	return nil
}
