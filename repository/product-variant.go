package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ProductVariantRepo struct {
	sql.Repository[models.ProductVariant]
}

func ProductVariantRepository(db *gorm.DB) *ProductVariantRepo {
	return &ProductVariantRepo{*sql.NewRepository[models.ProductVariant](db)}
}
