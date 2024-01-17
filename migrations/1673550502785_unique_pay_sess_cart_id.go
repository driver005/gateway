package migrations

import "reflect"

type UniquePaySessCartId1673550502785 struct {
	r Registry
}

func (m *UniquePaySessCartId1673550502785) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *UniquePaySessCartId1673550502785) Up() error {
	if err := m.r.Context().Exec(`DROP INDEX IF EXISTS "public"."UniqPaymentSessionCartIdProviderId"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX IF EXISTS "UniqPaymentSessionCartIdProviderId"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "UniqPaymentSessionCartIdProviderId" ON "payment_session" ("cart_id", "provider_id") WHERE cart_id IS NOT NULL`).Error; err != nil {
		return err
	}
	return nil
}
func (m *UniquePaySessCartId1673550502785) Down() error {
	if err := m.r.Context().Exec(`DROP INDEX "public"."UniqPaymentSessionCartIdProviderId"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "UniqPaymentSessionCartIdProviderId" ON "payment_session" ("cart_id", "provider_id") `).Error; err != nil {
		return err
	}
	return nil
}
