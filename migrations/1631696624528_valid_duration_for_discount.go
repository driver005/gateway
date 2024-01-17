package migrations

import "reflect"

type ValidDurationForDiscount1631696624528 struct {
	r Registry
}

func (m *ValidDurationForDiscount1631696624528) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *ValidDurationForDiscount1631696624528) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "discount" ADD "valid_duration" character varying`).Error; err != nil {
		return err
	}
	return nil
}
func (m *ValidDurationForDiscount1631696624528) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "discount" DROP COLUMN "valid_duration"`).Error; err != nil {
		return err
	}
	return nil
}
