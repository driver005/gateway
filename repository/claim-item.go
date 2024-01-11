package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ClaimItemRepo struct {
	sql.Repository[models.ClaimItem]
}

func ClaimItemRepository(db *gorm.DB) *ClaimItemRepo {
	return &ClaimItemRepo{*sql.NewRepository[models.ClaimItem](db)}
}
