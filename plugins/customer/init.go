package customer

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"Customer",
		db,
		".snapshot-medusa-customer.json",
		"plugins",
	)

	module.Migrate()
}
