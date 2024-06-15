package fulfillment

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"Fulfillment",
		db,
		".snapshot-medusa-fulfillment.json",
		"plugins",
	)

	module.Migrate()
}
