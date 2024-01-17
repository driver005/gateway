package migrations

import "reflect"

type UpdateReturnReasonIndex1692870898423 struct {
	r Registry
}

func (m *UpdateReturnReasonIndex1692870898423) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *UpdateReturnReasonIndex1692870898423) Up() error {
	if err := m.r.Context().Exec(`DROP INDEX IF EXISTS "IDX_00605f9d662c06b81c1b60ce24";`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_return_reason_value" ON "return_reason" ("value") WHERE deleted_at IS NULL;`).Error; err != nil {
		return err
	}
	return nil
}
func (m *UpdateReturnReasonIndex1692870898423) Down() error {
	if err := m.r.Context().Exec(`DROP INDEX IF EXISTS "IDX_return_reason_value";`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_00605f9d662c06b81c1b60ce24" ON "return_reason" ("value") `).Error; err != nil {
		return err
	}
	return nil
}
