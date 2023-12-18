package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ProductOptionRepo struct {
	Repository[models.ProductOption]
}

func ProductOptionRepository(db *gorm.DB) ProductOptionRepo {
	return ProductOptionRepo{*NewRepository[models.ProductOption](db)}
}
