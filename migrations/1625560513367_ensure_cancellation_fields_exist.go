package migrations

import "reflect"

type EnsureCancellationFieldsExist1625560513367 struct {
	r Registry
}

func (m *EnsureCancellationFieldsExist1625560513367) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *EnsureCancellationFieldsExist1625560513367) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "swap" ADD "canceled_at" TIMESTAMP WITH TIME ZONE`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TYPE "return_status_enum" RENAME TO "return_status_enum_old"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TYPE "return_status_enum" AS ENUM('requested', 'received', 'requires_action', 'canceled')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return" ALTER COLUMN "status" DROP DEFAULT`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return" ALTER COLUMN "status" TYPE "return_status_enum" USING "status"::"text"::"return_status_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return" ALTER COLUMN "status" SET DEFAULT 'requested'`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "return_status_enum_old"`).Error; err != nil {
		return err
	}
	return nil
}
func (m *EnsureCancellationFieldsExist1625560513367) Down() error {
	if err := m.r.Context().Exec(`CREATE TYPE "return_status_enum_old" AS ENUM('requested', 'received', 'requires_action')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return" ALTER COLUMN "status" DROP DEFAULT`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return" ALTER COLUMN "status" TYPE "return_status_enum_old" USING "status"::"text"::"return_status_enum_old"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return" ALTER COLUMN "status" SET DEFAULT 'requested'`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "return_status_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TYPE "return_status_enum_old" RENAME TO "return_status_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "swap" DROP COLUMN "canceled_at"`).Error; err != nil {
		return err
	}
	return nil
}
