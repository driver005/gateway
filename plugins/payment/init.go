package payment

import (
	"github.com/driver005/gateway/plugins/payment/migrations"
	"github.com/pressly/goose/v3"
)

func Init() {
	goose.AddNamedMigrationContext("20240502183800_payment.go", migrations.Up, migrations.Down)
}
