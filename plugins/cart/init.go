package cart

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"Cart",
		db,
		".snapshot-medusa-cart.json",
		"plugins",
	)

	module.Migrate()
}
