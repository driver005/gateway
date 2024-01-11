package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type AnalyticsConfigRepo struct {
	sql.Repository[models.AnalyticsConfig]
}

func AnalyticsConfigRepository(db *gorm.DB) *AnalyticsConfigRepo {
	return &AnalyticsConfigRepo{*sql.NewRepository[models.AnalyticsConfig](db)}
}
