package migrations

import "reflect"

type AddMetadataToProductCategory1699564794649 struct {
	r Registry
}

func (m *AddMetadataToProductCategory1699564794649) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *AddMetadataToProductCategory1699564794649) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "product_category" ADD COLUMN "metadata" jsonb NULL;`).Error; err != nil {
		return err
	}
	return nil
}
func (m *AddMetadataToProductCategory1699564794649) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "product_category" DROP COLUMN "metadata"`).Error; err != nil {
		return err
	}
	return nil
}
