package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type StoreRepo struct {
	sql.Repository[models.Store]
}

func StoreRepository(db *gorm.DB) *StoreRepo {
	return &StoreRepo{*sql.NewRepository[models.Store](db)}
}
