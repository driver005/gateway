package apikey

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"ApiKey",
		db,
		".snapshot-medusa-api-key.json",
		"plugins",
	)

	module.Migrate()
}
