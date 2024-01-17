package migrations

import "reflect"

type SwapFulfillmentStatusRequiresAction1661863940645 struct {
	r Registry
}

func (m *SwapFulfillmentStatusRequiresAction1661863940645) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *SwapFulfillmentStatusRequiresAction1661863940645) Up() error {
	if err := m.r.Context().Exec(`ALTER TYPE "swap_fulfillment_status_enum" RENAME TO "swap_fulfillment_status_enum_old"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TYPE "swap_fulfillment_status_enum" AS ENUM('not_fulfilled', 'fulfilled', 'shipped', 'partially_shipped', 'canceled', 'requires_action')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "swap" ALTER COLUMN "fulfillment_status" TYPE "swap_fulfillment_status_enum" USING "fulfillment_status"::"text"::"swap_fulfillment_status_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "swap_fulfillment_status_enum_old"`).Error; err != nil {
		return err
	}
	return nil
}
func (m *SwapFulfillmentStatusRequiresAction1661863940645) Down() error {
	if err := m.r.Context().Exec(`CREATE TYPE "swap_fulfillment_status_enum_old" AS ENUM('not_fulfilled', 'fulfilled', 'shipped', 'canceled', 'requires_action')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "swap" ALTER COLUMN "fulfillment_status" TYPE "swap_fulfillment_status_enum_old" USING "fulfillment_status"::"text"::"swap_fulfillment_status_enum_old"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "swap_fulfillment_status_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TYPE "swap_fulfillment_status_enum_old" RENAME TO "swap_fulfillment_status_enum"`).Error; err != nil {
		return err
	}
	return nil
}
