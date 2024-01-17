package migrations

import "reflect"

type DropVariantIdFkMoneyAmount1680857773273 struct {
	r Registry
}

func (m *DropVariantIdFkMoneyAmount1680857773273) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *DropVariantIdFkMoneyAmount1680857773273) Up() error {
	if err := m.r.Context().Exec(
		`alter table if exists "money_amount" drop constraint if exists "FK_17a06d728e4cfbc5bd2ddb70af0";`).Error; err != nil {
		return err
	}
	return nil
}
func (m *DropVariantIdFkMoneyAmount1680857773273) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" DROP CONSTRAINT "FK_17a06d728e4cfbc5bd2ddb70af0";`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE if exists "money_amount" ADD CONSTRAINT "FK_17a06d728e4cfbc5bd2ddb70af0" FOREIGN KEY ("variant_id") REFERENCES "product_variant"("id") ON DELETE cascade ON UPDATE NO ACTION;`).Error; err != nil {
		return err
	}
	return nil
}
