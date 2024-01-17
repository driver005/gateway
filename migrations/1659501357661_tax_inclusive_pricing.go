package migrations

import "reflect"

type TaxInclusivePricing1659501357661 struct {
	r Registry
}

func (m *TaxInclusivePricing1659501357661) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *TaxInclusivePricing1659501357661) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "currency" ADD "includes_tax" boolean NOT NULL DEFAULT false`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "region" ADD "includes_tax" boolean NOT NULL DEFAULT false`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "shipping_option" ADD "includes_tax" boolean NOT NULL DEFAULT false`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "price_list" ADD "includes_tax" boolean NOT NULL DEFAULT false`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "shipping_method" ADD "includes_tax" boolean NOT NULL DEFAULT false`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item" ADD "includes_tax" boolean NOT NULL DEFAULT false`).Error; err != nil {
		return err
	}
	return nil
}
func (m *TaxInclusivePricing1659501357661) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "line_item" DROP COLUMN "includes_tax"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "shipping_method" DROP COLUMN "includes_tax"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "price_list" DROP COLUMN "includes_tax"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "shipping_option" DROP COLUMN "includes_tax"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "region" DROP COLUMN "includes_tax"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "currency" DROP COLUMN "includes_tax"`).Error; err != nil {
		return err
	}
	return nil
}
