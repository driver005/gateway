package region

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"Region",
		db,
		".snapshot-medusa-region.json",
		"plugins",
	)

	module.Migrate()
}
