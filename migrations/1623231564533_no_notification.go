package migrations

import "reflect"

type NoNotification1623231564533 struct {
	r Registry
}

func (m *NoNotification1623231564533) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *NoNotification1623231564533) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "return" ADD "no_notification" boolean`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "claim_order" ADD "no_notification" boolean`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "swap" ADD "no_notification" boolean`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "order" ADD "no_notification" boolean`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "draft_order" ADD "no_notification_order" boolean`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "fulfillment" ADD "no_notification" boolean`).Error; err != nil {
		return err
	}
	return nil
}
func (m *NoNotification1623231564533) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "fulfillment" DROP COLUMN "no_notification"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "draft_order" DROP COLUMN "no_notification_order"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "order" DROP COLUMN "no_notification"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "swap" DROP COLUMN "no_notification"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "claim_order" DROP COLUMN "no_notification"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return" DROP COLUMN "no_notification"`).Error; err != nil {
		return err
	}
	return nil
}
