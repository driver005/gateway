package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ProductCategoryRepo struct {
	Repository[models.ProductCategory]
}

func ProductCategoryRepository(db *gorm.DB) ProductCategoryRepo {
	return ProductCategoryRepo{*NewRepository[models.ProductCategory](db)}
}
