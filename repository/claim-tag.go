package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ClaimTagRepo struct {
	Repository[models.ClaimTag]
}

func ClaimTagRepository(db *gorm.DB) ClaimTagRepo {
	return ClaimTagRepo{*NewRepository[models.ClaimTag](db)}
}
