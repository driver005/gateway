package repository

import (
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"gorm.io/gorm"
)

type NotificationRepo struct {
	sql.Repository[models.Notification]
}

func NotificationRepository(db *gorm.DB) *NotificationRepo {
	return &NotificationRepo{*sql.NewRepository[models.Notification](db)}
}
