package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type NotificationProviderRepo struct {
	Repository[models.NotificationProvider]
}

func NotificationProviderRepository(db *gorm.DB) NotificationProviderRepo {
	return NotificationProviderRepo{*NewRepository[models.NotificationProvider](db)}
}
