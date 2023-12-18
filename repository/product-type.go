package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ProductTypeRepo struct {
	Repository[models.ProductType]
}

func ProductTypeRepository(db *gorm.DB) ProductTypeRepo {
	return ProductTypeRepo{*NewRepository[models.ProductType](db)}
}
