package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type AnalyticsConfigRepo struct {
	Repository[models.AnalyticsConfig]
}

func AnalyticsConfigRepository(db *gorm.DB) AnalyticsConfigRepo {
	return AnalyticsConfigRepo{*NewRepository[models.AnalyticsConfig](db)}
}
