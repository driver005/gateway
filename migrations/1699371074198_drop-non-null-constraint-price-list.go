package migrations

import "reflect"

type DropNonNullConstraintPriceList1699371074198 struct {
	r Registry
}

func (m *DropNonNullConstraintPriceList1699371074198) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *DropNonNullConstraintPriceList1699371074198) Up() error {
	if err := m.r.Context().Exec(`
      ALTER TABLE IF EXISTS price_list ALTER COLUMN name DROP NOT NULL;
    `).Error; err != nil {
		return err
	}
	return nil
}
func (m *DropNonNullConstraintPriceList1699371074198) Down() error {
	if err := m.r.Context().Exec(`
      ALTER TABLE IF EXISTS price_list ALTER COLUMN name SET NOT NULL;
    `).Error; err != nil {
		return err
	}
	return nil
}
