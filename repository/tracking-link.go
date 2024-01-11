package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type TrackingLinkRepo struct {
	sql.Repository[models.TrackingLink]
}

func TrackingLinkRepository(db *gorm.DB) *TrackingLinkRepo {
	return &TrackingLinkRepo{*sql.NewRepository[models.TrackingLink](db)}
}
