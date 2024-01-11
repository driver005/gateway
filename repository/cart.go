package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type CartRepo struct {
	sql.Repository[models.Cart]
}

func CartRepository(db *gorm.DB) *CartRepo {
	return &CartRepo{*sql.NewRepository[models.Cart](db)}
}
