package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ClaimImageRepo struct {
	Repository[models.ClaimImage]
}

func ClaimImageRepository(db *gorm.DB) ClaimImageRepo {
	return ClaimImageRepo{*NewRepository[models.ClaimImage](db)}
}
