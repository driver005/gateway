package inventory

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"Inventory",
		db,
		".snapshot-medusa-inventory.json",
		"plugins",
	)

	module.Migrate()
}
