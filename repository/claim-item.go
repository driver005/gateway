package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ClaimItemRepo struct {
	Repository[models.ClaimItem]
}

func ClaimItemRepository(db *gorm.DB) ClaimItemRepo {
	return ClaimItemRepo{*NewRepository[models.ClaimItem](db)}
}
