package migrations

import "reflect"

type DeleteDateOnShippingOptionRequirements1632828114899 struct {
	r Registry
}

func (m *DeleteDateOnShippingOptionRequirements1632828114899) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *DeleteDateOnShippingOptionRequirements1632828114899) Up() error {
	if err := m.r.Context().Exec(
		`ALTER TABLE "shipping_option_requirement" ADD "deleted_at" TIMESTAMP WITH TIME ZONE`).Error; err != nil {
		return err
	}
	return nil
}
func (m *DeleteDateOnShippingOptionRequirements1632828114899) Down() error {
	if err := m.r.Context().Exec(
		`ALTER TABLE "shipping_option_requirement" DROP COLUMN "deleted_at"`).Error; err != nil {
		return err
	}
	return nil
}
