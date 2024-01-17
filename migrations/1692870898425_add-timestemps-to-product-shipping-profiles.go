package migrations

import "reflect"

type AddTimestempsToProductShippingProfiles1692870898425 struct {
	r Registry
}

func (m *AddTimestempsToProductShippingProfiles1692870898425) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *AddTimestempsToProductShippingProfiles1692870898425) Up() error {
	if err := m.r.Context().Exec(`
		ALTER TABLE "product_shipping_profile" ADD COLUMN IF NOT EXISTS "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now();
		ALTER TABLE "product_shipping_profile" ADD COLUMN IF NOT EXISTS "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now();
    	ALTER TABLE "product_shipping_profile" ADD COLUMN IF NOT EXISTS "deleted_at" TIMESTAMP WITH TIME ZONE;
    `).Error; err != nil {
		return err
	}
	return nil
}
func (m *AddTimestempsToProductShippingProfiles1692870898425) Down() error {
	if err := m.r.Context().Exec(`
        ALTER TABLE "product_shipping_profile" DROP COLUMN IF EXISTS "created_at";
		ALTER TABLE "product_shipping_profile" DROP COLUMN IF EXISTS "updated_at";  
		ALTER TABLE "product_shipping_profile" DROP COLUMN IF EXISTS "deleted_at";

    `).Error; err != nil {
		return err
	}
	return nil
}
