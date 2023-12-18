package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ProductCollectionRepo struct {
	Repository[models.ProductCollection]
}

func ProductCollectionRepository(db *gorm.DB) ProductCollectionRepo {
	return ProductCollectionRepo{*NewRepository[models.ProductCollection](db)}
}
