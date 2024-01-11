package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ProductOptionRepo struct {
	sql.Repository[models.ProductOption]
}

func ProductOptionRepository(db *gorm.DB) *ProductOptionRepo {
	return &ProductOptionRepo{*sql.NewRepository[models.ProductOption](db)}
}
