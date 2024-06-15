package notification

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"Notification",
		db,
		".snapshot-medusa-notification.json",
		"plugins",
	)

	module.Migrate()
}
