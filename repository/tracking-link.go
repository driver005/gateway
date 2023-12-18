package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type TrackingLinkRepo struct {
	Repository[models.TrackingLink]
}

func TrackingLinkRepository(db *gorm.DB) TrackingLinkRepo {
	return TrackingLinkRepo{*NewRepository[models.TrackingLink](db)}
}
