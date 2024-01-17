package migrations

import "reflect"

type LineitemProductId1692870898424 struct {
	r Registry
}

func (m *LineitemProductId1692870898424) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *LineitemProductId1692870898424) Up() error {
	if err := m.r.Context().Exec(`
      ALTER TABLE "line_item" DROP CONSTRAINT IF EXISTS "FK_5371cbaa3be5200f373d24e3d5b";
      ALTER TABLE "line_item" ADD COLUMN IF NOT EXISTS "product_id" text;

      UPDATE "line_item" SET "product_id" = pv."product_id"
      FROM "product_variant" pv
      WHERE "line_item"."variant_id" = "pv"."id";
    `).Error; err != nil {
		return err
	}
	return nil
}
func (m *LineitemProductId1692870898424) Down() error {
	if err := m.r.Context().Exec(`
      	ALTER TABLE "line_item" DROP COLUMN IF EXISTS "product_id";
      	ALTER TABLE "line_item" ADD CONSTRAINT "FK_5371cbaa3be5200f373d24e3d5b" FOREIGN KEY ("variant_id") REFERENCES "product_variant" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
    `).Error; err != nil {
		return err
	}
	return nil
}
