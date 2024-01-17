package migrations

import "reflect"

type ExternalIdOrder1638952072999 struct {
	r Registry
}

func (m *ExternalIdOrder1638952072999) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *ExternalIdOrder1638952072999) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "order" ADD "external_id" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product" ADD "external_id" character varying`).Error; err != nil {
		return err
	}
	return nil
}
func (m *ExternalIdOrder1638952072999) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "product" DROP COLUMN "external_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "order" DROP COLUMN "external_id"`).Error; err != nil {
		return err
	}
	return nil
}
