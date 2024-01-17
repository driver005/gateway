package migrations

import "reflect"

type NestedReturnReasons1631800727788 struct {
	r Registry
}

func (m *NestedReturnReasons1631800727788) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *NestedReturnReasons1631800727788) Up() error {
	if err := m.r.Context().Exec(
		`ALTER TABLE "return_reason" ADD "parent_return_reason_id" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return_reason" ADD CONSTRAINT "FK_2250c5d9e975987ab212f61a657" FOREIGN KEY ("parent_return_reason_id") REFERENCES "return_reason"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
func (m *NestedReturnReasons1631800727788) Down() error {
	if err := m.r.Context().Exec(
		`ALTER TABLE "return_reason" DROP COLUMN "parent_return_reason_id"`).Error; err != nil {
		return err
	}
	return nil
}
