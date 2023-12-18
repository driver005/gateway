package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ProductOptionValueRepo struct {
	Repository[models.ProductOptionValue]
}

func ProductOptionValueRepository(db *gorm.DB) ProductOptionValueRepo {
	return ProductOptionValueRepo{*NewRepository[models.ProductOptionValue](db)}
}
