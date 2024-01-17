package migrations

import "reflect"

type StatusOnProduct1631864388026 struct {
	r Registry
}

func (m *StatusOnProduct1631864388026) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *StatusOnProduct1631864388026) Up() error {
	if err := m.r.Context().Exec(
		`CREATE TYPE "product_status_enum" AS ENUM('draft', 'proposed', 'published', 'rejected')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(
		`ALTER TABLE "product" ADD "status" "product_status_enum" `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(
		`UPDATE "product" SET "status" = 'published' WHERE "status" IS NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(
		`ALTER TABLE "product" ALTER COLUMN "status" SET NOT NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(
		`ALTER TABLE "product" ALTER COLUMN "status" SET DEFAULT 'draft'`).Error; err != nil {
		return err
	}
	return nil
}
func (m *StatusOnProduct1631864388026) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "product" DROP COLUMN "status"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "product_status_enum"`).Error; err != nil {
		return err
	}
	return nil
}
