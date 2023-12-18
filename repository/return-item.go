package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ReturnItemRepo struct {
	Repository[models.ReturnItem]
}

func ReturnItemRepository(db *gorm.DB) ReturnItemRepo {
	return ReturnItemRepo{*NewRepository[models.ReturnItem](db)}
}
