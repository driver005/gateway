package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type RegionRepo struct {
	sql.Repository[models.Region]
}

func RegionRepository(db *gorm.DB) *RegionRepo {
	return &RegionRepo{*sql.NewRepository[models.Region](db)}
}
