package migrations

import "reflect"

type PaymentSessionIsInitiated1672906846560 struct {
	r Registry
}

func (m *PaymentSessionIsInitiated1672906846560) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *PaymentSessionIsInitiated1672906846560) Up() error {
	if err := m.r.Context().Exec(`
      ALTER TABLE payment_session ADD COLUMN is_initiated BOOLEAN NOT NULL DEFAULT false
    `).Error; err != nil {
		return err
	}

	// Set is_initiated to true if there is more that 0 key in the data. We assume that if data contains any key
	// A payment has been initiated to the payment provider
	if err := m.r.Context().Exec(`
      UPDATE payment_session SET is_initiated = true WHERE (
          SELECT coalesce(json_array_length(json_agg(keys)), 0)
          FROM jsonb_object_keys(data) AS keys (keys)
) > 0
    `).Error; err != nil {
		return err
	}
	return nil
}
func (m *PaymentSessionIsInitiated1672906846560) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE payment_session DROP COLUMN is_initiated`).Error; err != nil {
		return err
	}
	return nil
}
