package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ReturnItemRepo struct {
	sql.Repository[models.ReturnItem]
}

func ReturnItemRepository(db *gorm.DB) *ReturnItemRepo {
	return &ReturnItemRepo{*sql.NewRepository[models.ReturnItem](db)}
}
