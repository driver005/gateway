package saleschannel

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"SalesChannel",
		db,
		".snapshot-medusa-sales-channel-tst.json",
		"plugins",
	)

	module.Migrate()
}
