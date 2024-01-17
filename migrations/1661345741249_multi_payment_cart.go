package migrations

import "reflect"

type MultiPaymentCart1661345741249 struct {
	r Registry
}

func (m *MultiPaymentCart1661345741249) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *MultiPaymentCart1661345741249) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "payment" DROP CONSTRAINT "REL_4665f17abc1e81dd58330e5854"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "UniquePaymentActive" ON "payment" ("cart_id") WHERE canceled_at IS NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_aac4855eadda71aa1e4b6d7684" ON "payment" ("cart_id") WHERE canceled_at IS NOT NULL`).Error; err != nil {
		return err
	}
	return nil
}
func (m *MultiPaymentCart1661345741249) Down() error {
	if err := m.r.Context().Exec(`DROP INDEX "IDX_aac4855eadda71aa1e4b6d7684"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "UniquePaymentActive"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "payment" ADD CONSTRAINT "REL_4665f17abc1e81dd58330e5854" UNIQUE ("cart_id")`).Error; err != nil {
		return err
	}
	return nil
}
