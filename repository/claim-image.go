package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type ClaimImageRepo struct {
	sql.Repository[models.ClaimImage]
}

func ClaimImageRepository(db *gorm.DB) *ClaimImageRepo {
	return &ClaimImageRepo{*sql.NewRepository[models.ClaimImage](db)}
}
