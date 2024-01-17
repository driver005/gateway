package migrations

import "reflect"

type LineitemAdjustmentsAmount1678093365812 struct {
	r Registry
}

func (m *LineitemAdjustmentsAmount1678093365812) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *LineitemAdjustmentsAmount1678093365812) Up() error {
	if err := m.r.Context().Exec(`
        ALTER TABLE line_item_adjustment ALTER COLUMN amount TYPE NUMERIC;
    `).Error; err != nil {
		return err
	}
	return nil
}
func (m *LineitemAdjustmentsAmount1678093365812) Down() error {
	if err := m.r.Context().Exec(`
        ALTER TABLE line_item_adjustment ALTER COLUMN amount TYPE integer;
    `).Error; err != nil {
		return err
	}
	return nil
}
