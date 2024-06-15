package stocklocation

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"StockLocation",
		db,
		".snapshot-medusa-stock-location.json",
		"plugins",
	)

	module.Migrate()
}
