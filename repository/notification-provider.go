package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type NotificationProviderRepo struct {
	sql.Repository[models.NotificationProvider]
}

func NotificationProviderRepository(db *gorm.DB) *NotificationProviderRepo {
	return &NotificationProviderRepo{*sql.NewRepository[models.NotificationProvider](db)}
}
