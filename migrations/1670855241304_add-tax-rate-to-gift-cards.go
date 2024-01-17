package migrations

import "reflect"

type AddTaxRateToGiftCards1670855241304 struct {
	r Registry
}

func (m *AddTaxRateToGiftCards1670855241304) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *AddTaxRateToGiftCards1670855241304) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "gift_card" ADD COLUMN IF NOT EXISTS tax_rate REAL`).Error; err != nil {
		return err
	}
	return nil
}
func (m *AddTaxRateToGiftCards1670855241304) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "gift_card" DROP COLUMN IF EXISTS "tax_rate"`).Error; err != nil {
		return err
	}
	return nil
}
