package payment

import (
	"github.com/driver005/gateway/internall/modules"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	module := modules.NewModule(
		"payment",
		db,
		".snapshot-medusa-payment.json",
		"plugins",
	)

	module.Migrate()
}
