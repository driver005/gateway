package migrations

import "reflect"

type OrderTaxRateToRealType1638543550000 struct {
	r Registry
}

func (m *OrderTaxRateToRealType1638543550000) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *OrderTaxRateToRealType1638543550000) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "order" ALTER COLUMN "tax_rate" TYPE REAL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "region" ALTER COLUMN "tax_rate" TYPE REAL`).Error; err != nil {
		return err
	}
	return nil
}
func (m *OrderTaxRateToRealType1638543550000) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "order" ALTER COLUMN "tax_rate" TYPE INTEGER`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "region" ALTER COLUMN "tax_rate" TYPE NUMERIC`).Error; err != nil {
		return err
	}
	return nil
}
