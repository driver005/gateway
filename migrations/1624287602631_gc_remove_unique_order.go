package migrations

import "reflect"

type GcRemoveUniqueOrder1624287602631 struct {
	r Registry
}

func (m *GcRemoveUniqueOrder1624287602631) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *GcRemoveUniqueOrder1624287602631) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "gift_card" DROP CONSTRAINT "FK_dfc1f02bb0552e79076aa58dbb0"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`COMMENT ON COLUMN "gift_card"."order_id" IS NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "gift_card" DROP CONSTRAINT "REL_dfc1f02bb0552e79076aa58dbb"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "gift_card" ADD CONSTRAINT "FK_dfc1f02bb0552e79076aa58dbb0" FOREIGN KEY ("order_id") REFERENCES "order"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}

	return nil
}
func (m *GcRemoveUniqueOrder1624287602631) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "gift_card" DROP CONSTRAINT "FK_dfc1f02bb0552e79076aa58dbb0"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "gift_card" ADD CONSTRAINT "REL_dfc1f02bb0552e79076aa58dbb" UNIQUE ("order_id")`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`COMMENT ON COLUMN "gift_card"."order_id" IS NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "gift_card" ADD CONSTRAINT "FK_dfc1f02bb0552e79076aa58dbb0" FOREIGN KEY ("order_id") REFERENCES "order"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
