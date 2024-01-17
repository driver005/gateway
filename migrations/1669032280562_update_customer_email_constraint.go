package migrations

import "reflect"

type UpdateCustomerEmailConstraint1669032280562 struct {
	r Registry
}

func (m *UpdateCustomerEmailConstraint1669032280562) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *UpdateCustomerEmailConstraint1669032280562) Up() error {
	if err := m.r.Context().Exec(`DROP INDEX "IDX_fdb2f3ad8115da4c7718109a6e"`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`ALTER TABLE "customer" ADD CONSTRAINT "UQ_unique_email_for_guests_and_customer_accounts" UNIQUE ("email", "has_account")`).Error; err != nil {
		return err
	}
	return nil
}
func (m *UpdateCustomerEmailConstraint1669032280562) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "customer" DROP CONSTRAINT "UQ_unique_email_for_guests_and_customer_accounts"`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_fdb2f3ad8115da4c7718109a6e" ON "customer" ("email") `).Error; err != nil {
		return err
	}
	return nil
}
