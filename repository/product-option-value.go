package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ProductOptionValueRepo struct {
	sql.Repository[models.ProductOptionValue]
}

func ProductOptionValueRepository(db *gorm.DB) *ProductOptionValueRepo {
	return &ProductOptionValueRepo{*sql.NewRepository[models.ProductOptionValue](db)}
}
