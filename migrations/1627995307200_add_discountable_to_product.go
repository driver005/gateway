package migrations

import "reflect"

type AddDiscountableToProduct1627995307200 struct {
	r Registry
}

func (m *AddDiscountableToProduct1627995307200) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *AddDiscountableToProduct1627995307200) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "product" ADD "discountable" boolean NOT NULL DEFAULT true`).Error; err != nil {
		return err
	}
	return nil
}
func (m *AddDiscountableToProduct1627995307200) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "product" DROP COLUMN "discountable"`).Error; err != nil {
		return err
	}
	return nil
}
