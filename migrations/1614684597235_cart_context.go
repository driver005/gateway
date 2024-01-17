package migrations

import "reflect"

type CartContext1614684597235 struct {
	r Registry
}

func (m *CartContext1614684597235) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *CartContext1614684597235) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "cart" ADD "context" jsonb`).Error; err != nil {
		return err
	}
	return nil
}
func (m *CartContext1614684597235) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "cart" DROP COLUMN "context"`).Error; err != nil {
		return err
	}
	return nil
}
