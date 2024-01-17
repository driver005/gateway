package migrations

import "reflect"

type AddDescriptionToProductCategory1680857773272 struct {
	r Registry
}

func (m *AddDescriptionToProductCategory1680857773272) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *AddDescriptionToProductCategory1680857773272) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "product_category" ADD COLUMN IF NOT EXISTS "description" TEXT NOT NULL DEFAULT ''`).Error; err != nil {
		return err
	}
	return nil
}
func (m *AddDescriptionToProductCategory1680857773272) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "product_category" DROP COLUMN IF EXISTS "description"`).Error; err != nil {
		return err
	}
	return nil
}
