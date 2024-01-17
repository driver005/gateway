package migrations

import "reflect"

type LineItemTaxAdjustmentOnCascadeDelete1680857773272 struct {
	r Registry
}

func (m *LineItemTaxAdjustmentOnCascadeDelete1680857773272) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *LineItemTaxAdjustmentOnCascadeDelete1680857773272) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "line_item_tax_line" DROP CONSTRAINT IF EXISTS "FK_5077fa54b0d037e984385dfe8ad"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item_tax_line" ADD CONSTRAINT "FK_5077fa54b0d037e984385dfe8ad" FOREIGN KEY ("item_id") REFERENCES "line_item"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`ALTER TABLE "line_item_adjustment" DROP CONSTRAINT IF EXISTS "FK_be9aea2ccf3567007b6227da4d2"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item_adjustment" ADD CONSTRAINT "FK_be9aea2ccf3567007b6227da4d2" FOREIGN KEY ("item_id") REFERENCES "line_item"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
func (m *LineItemTaxAdjustmentOnCascadeDelete1680857773272) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "line_item_tax_line" DROP CONSTRAINT "FK_5077fa54b0d037e984385dfe8ad"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item_tax_line" ADD CONSTRAINT "FK_5077fa54b0d037e984385dfe8ad" FOREIGN KEY ("item_id") REFERENCES "line_item"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`ALTER TABLE "line_item_adjustment" DROP CONSTRAINT "FK_be9aea2ccf3567007b6227da4d2"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item_adjustment" ADD CONSTRAINT "FK_be9aea2ccf3567007b6227da4d2" FOREIGN KEY ("item_id") REFERENCES "line_item"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
