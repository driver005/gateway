package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ProductRepo struct {
	Repository[models.Product]
}

func ProductRepository(db *gorm.DB) ProductRepo {
	return ProductRepo{*NewRepository[models.Product](db)}
}
