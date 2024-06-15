package currency

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"Currency",
		db,
		".snapshot-medusa-currency.json",
		"plugins",
	)

	module.Migrate()
}
