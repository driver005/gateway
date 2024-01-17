package migrations

import "reflect"

type TaxLineConstraints1648641130007 struct {
	r Registry
}

func (m *TaxLineConstraints1648641130007) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *TaxLineConstraints1648641130007) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "line_item_tax_line" ADD CONSTRAINT "UQ_3c2af51043ed7243e7d9775a2ad" UNIQUE ("item_id", "code")`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "shipping_method_tax_line" ADD CONSTRAINT "UQ_cd147fca71e50bc954139fa3104" UNIQUE ("shipping_method_id", "code")`).Error; err != nil {
		return err
	}
	return nil
}
func (m *TaxLineConstraints1648641130007) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "shipping_method_tax_line" DROP CONSTRAINT "UQ_cd147fca71e50bc954139fa3104"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item_tax_line" DROP CONSTRAINT "UQ_3c2af51043ed7243e7d9775a2ad"`).Error; err != nil {
		return err
	}
	return nil
}
