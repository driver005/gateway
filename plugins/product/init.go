package product

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"Product",
		db,
		".snapshot-medusa-products.json",
		"plugins",
	)

	module.Migrate()
}
