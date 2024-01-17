package migrations

import "reflect"

type AddLineItemAdjustments1648600574750 struct {
	r Registry
}

func (m *AddLineItemAdjustments1648600574750) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *AddLineItemAdjustments1648600574750) Up() error {
	if err := m.r.Context().Exec(`CREATE TABLE "line_item_adjustment" ("id" character varying NOT NULL, "item_id" character varying NOT NULL, "description" character varying NOT NULL, "discount_id" character varying, "amount" integer NOT NULL, "metadata" jsonb, CONSTRAINT "PK_2b1360103753df2dc8257c2c8c3" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_be9aea2ccf3567007b6227da4d" ON "line_item_adjustment" ("item_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_2f41b20a71f30e60471d7e3769" ON "line_item_adjustment" ("discount_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_bf701b88d2041392a288785ada" ON "line_item_adjustment" ("discount_id", "item_id") WHERE "discount_id" IS NOT NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item_adjustment" ADD CONSTRAINT "FK_be9aea2ccf3567007b6227da4d2" FOREIGN KEY ("item_id") REFERENCES "line_item"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item_adjustment" ADD CONSTRAINT "FK_2f41b20a71f30e60471d7e3769c" FOREIGN KEY ("discount_id") REFERENCES "discount"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
func (m *AddLineItemAdjustments1648600574750) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "line_item_adjustment" DROP CONSTRAINT "FK_2f41b20a71f30e60471d7e3769c"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item_adjustment" DROP CONSTRAINT "FK_be9aea2ccf3567007b6227da4d2"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_bf701b88d2041392a288785ada"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_2f41b20a71f30e60471d7e3769"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_be9aea2ccf3567007b6227da4d"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "line_item_adjustment"`).Error; err != nil {
		return err
	}
	return nil
}
