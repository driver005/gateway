package auth

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"Auth",
		db,
		".snapshot-medusa-auth.json",
		"plugins",
	)

	module.Migrate()
}
