package promotion

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"Promotion",
		db,
		".snapshot-medusa-promotion.json",
		"plugins",
	)

	module.Migrate()
}
