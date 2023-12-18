package repository

import (
	"github.com/driver005/gateway/models"
	"gorm.io/gorm"
)

type NotificationRepo struct {
	Repository[models.Notification]
}

func NotificationRepository(db *gorm.DB) NotificationRepo {
	return NotificationRepo{*NewRepository[models.Notification](db)}
}
