package migrations

import "reflect"

type AllowBackorderSwaps1630505790603 struct {
	r Registry
}

func (m *AllowBackorderSwaps1630505790603) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *AllowBackorderSwaps1630505790603) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "swap" ADD "allow_backorder" boolean NOT NULL DEFAULT false`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "cart" ADD "payment_authorized_at" TIMESTAMP WITH TIME ZONE`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TYPE "swap_payment_status_enum" RENAME TO "swap_payment_status_enum_old"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TYPE "swap_payment_status_enum" AS ENUM('not_paid', 'awaiting', 'captured', 'confirmed', 'canceled', 'difference_refunded', 'partially_refunded', 'refunded', 'requires_action')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "swap" ALTER COLUMN "payment_status" TYPE "swap_payment_status_enum" USING "payment_status"::"text"::"swap_payment_status_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "swap_payment_status_enum_old"`).Error; err != nil {
		return err
	}
	return nil
}
func (m *AllowBackorderSwaps1630505790603) Down() error {
	if err := m.r.Context().Exec(`CREATE TYPE "swap_payment_status_enum_old" AS ENUM('not_paid', 'awaiting', 'captured', 'canceled', 'difference_refunded', 'partially_refunded', 'refunded', 'requires_action')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "swap" ALTER COLUMN "payment_status" TYPE "swap_payment_status_enum_old" USING "payment_status"::"text"::"swap_payment_status_enum_old"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "swap_payment_status_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TYPE "swap_payment_status_enum_old" RENAME TO  "swap_payment_status_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "cart" DROP COLUMN "payment_authorized_at"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "swap" DROP COLUMN "allow_backorder"`).Error; err != nil {
		return err
	}
	return nil
}
