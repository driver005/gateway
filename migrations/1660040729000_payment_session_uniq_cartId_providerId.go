package migrations

import "reflect"

type PaymentSessionUniqCartIdProviderId1660040729000 struct {
	r Registry
}

func (m *PaymentSessionUniqCartIdProviderId1660040729000) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *PaymentSessionUniqCartIdProviderId1660040729000) Up() error {
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "UniqPaymentSessionCartIdProviderId" ON "payment_session" ("cart_id", "provider_id")`).Error; err != nil {
		return err
	}
	return nil
}
func (m *PaymentSessionUniqCartIdProviderId1660040729000) Down() error {
	if err := m.r.Context().Exec(`DROP INDEX "UniqPaymentSessionCartIdProviderId"`).Error; err != nil {
		return err
	}
	return nil
}
