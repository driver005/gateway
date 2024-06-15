package user

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"User",
		db,
		".snapshot-medusa-user.json",
		"plugins",
	)

	module.Migrate()
}
