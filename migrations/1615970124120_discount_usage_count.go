package migrations

import "reflect"

type DiscountUsageCount1615970124120 struct {
	r Registry
}

func (m *DiscountUsageCount1615970124120) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *DiscountUsageCount1615970124120) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "discount_rule" ADD "usage_count" integer NOT NULL DEFAULT '0'`).Error; err != nil {
		return err
	}
	return nil
}
func (m *DiscountUsageCount1615970124120) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "discount_rule" DROP COLUMN "usage_count"`).Error; err != nil {
		return err
	}
	return nil
}
