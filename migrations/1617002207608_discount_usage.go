package migrations

import "reflect"

type DiscountUsage1617002207608 struct {
	r Registry
}

func (m *DiscountUsage1617002207608) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *DiscountUsage1617002207608) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "discount_rule" DROP COLUMN "usage_limit"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_rule" DROP COLUMN "usage_count"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount" ADD "usage_limit" integer`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount" ADD "usage_count" integer NOT NULL DEFAULT '0'`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_rule" ALTER COLUMN "description" DROP NOT NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`COMMENT ON COLUMN "discount_rule"."description" IS NULL`).Error; err != nil {
		return err
	}
	return nil
}
func (m *DiscountUsage1617002207608) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "discount" DROP COLUMN "usage_count"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount" DROP COLUMN "usage_limit"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_rule" ADD "usage_count" integer NOT NULL DEFAULT '0'`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_rule" ADD "usage_limit" integer`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`COMMENT ON COLUMN "discount_rule"."description" IS NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_rule" ALTER COLUMN "description" SET NOT NULL`).Error; err != nil {
		return err
	}
	return nil
}
