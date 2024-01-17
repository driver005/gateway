package migrations

import "reflect"

type AddSalesChannelMetadata1680714052628 struct {
	r Registry
}

func (m *AddSalesChannelMetadata1680714052628) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *AddSalesChannelMetadata1680714052628) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "sales_channel" ADD COLUMN "metadata" jsonb NULL;`).Error; err != nil {
		return err
	}
	return nil
}
func (m *AddSalesChannelMetadata1680714052628) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "sales_channel" DROP COLUMN "metadata"`).Error; err != nil {
		return err
	}
	return nil
}
