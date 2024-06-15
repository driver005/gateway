package order

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"Order",
		db,
		".snapshot-medusa-order.json",
		"plugins",
	)

	module.Migrate()
}
