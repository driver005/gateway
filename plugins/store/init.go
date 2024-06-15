package store

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"Store",
		db,
		".snapshot-medusa-store.json",
		"plugins",
	)

	module.Migrate()
}
