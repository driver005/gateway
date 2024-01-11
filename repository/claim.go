package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ClaimRepo struct {
	sql.Repository[models.ClaimOrder]
}

func ClaimRepository(db *gorm.DB) *ClaimRepo {
	return &ClaimRepo{*sql.NewRepository[models.ClaimOrder](db)}
}
