package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type RegionRepo struct {
	Repository[models.Region]
}

func RegionRepository(db *gorm.DB) RegionRepo {
	return RegionRepo{*NewRepository[models.Region](db)}
}
