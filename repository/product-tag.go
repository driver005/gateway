package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ProductTagRepo struct {
	Repository[models.ProductTag]
}

func ProductTagRepository(db *gorm.DB) ProductTagRepo {
	return ProductTagRepo{*NewRepository[models.ProductTag](db)}
}
