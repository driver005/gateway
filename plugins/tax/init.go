package tax

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"Tax",
		db,
		".snapshot-medusa-tax.json",
		"plugins",
	)

	module.Migrate()
}
