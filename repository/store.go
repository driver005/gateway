package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type StoreRepo struct {
	Repository[models.Store]
}

func StoreRepository(db *gorm.DB) StoreRepo {
	return StoreRepo{*NewRepository[models.Store](db)}
}
