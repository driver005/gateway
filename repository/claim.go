package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type ClaimRepo struct {
	Repository[models.ClaimOrder]
}

func ClaimRepository(db *gorm.DB) ClaimRepo {
	return ClaimRepo{*NewRepository[models.ClaimOrder](db)}
}
