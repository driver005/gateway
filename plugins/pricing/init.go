package pricing

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"Pricing",
		db,
		".snapshot-medusa-pricing.json",
		"plugins",
	)

	module.Migrate()
}
