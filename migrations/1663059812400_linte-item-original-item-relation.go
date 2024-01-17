package migrations

import "reflect"

type LinteItemOriginalItemRelation1663059812400 struct {
	r Registry
}

func (m *LinteItemOriginalItemRelation1663059812400) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *LinteItemOriginalItemRelation1663059812400) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "line_item"
          ADD COLUMN IF NOT EXISTS original_item_id character varying,
          ADD COLUMN IF NOT EXISTS order_edit_id character varying`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`ALTER TABLE "line_item"
          ADD CONSTRAINT "line_item_original_item_fk" FOREIGN KEY ("original_item_id") REFERENCES "line_item" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item"
          ADD CONSTRAINT "line_item_order_edit_fk" FOREIGN KEY ("order_edit_id") REFERENCES "order_edit" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "unique_li_original_item_id_order_edit_id" ON "line_item" ("order_edit_id", "original_item_id") WHERE original_item_id IS NOT NULL AND order_edit_id IS NOT NULL`).Error; err != nil {
		return err
	}
	return nil
}
func (m *LinteItemOriginalItemRelation1663059812400) Down() error {
	if err := m.r.Context().Exec(`DROP INDEX IF EXISTS "unique_li_original_item_id_order_edit_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item" DROP CONSTRAINT "line_item_original_item_fk"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item" DROP CONSTRAINT "line_item_order_edit_fk"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item" DROP COLUMN "original_item_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item" DROP COLUMN "order_edit_id"`).Error; err != nil {
		return err
	}
	return nil
}
