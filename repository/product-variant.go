package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ProductVariantRepo struct {
	Repository[models.ProductVariant]
}

func ProductVariantRepository(db *gorm.DB) ProductVariantRepo {
	return ProductVariantRepo{*NewRepository[models.ProductVariant](db)}
}
