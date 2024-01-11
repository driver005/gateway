package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type SalesChannelLocationRepo struct {
	sql.Repository[models.SalesChannelLocation]
}

func SalesChannelLocationRepository(db *gorm.DB) *SalesChannelLocationRepo {
	return &SalesChannelLocationRepo{*sql.NewRepository[models.SalesChannelLocation](db)}
}
