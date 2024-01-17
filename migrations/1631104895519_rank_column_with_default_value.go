package migrations

import "reflect"

type RankColumnWithDefaultValue1631104895519 struct {
	r Registry
}

func (m *RankColumnWithDefaultValue1631104895519) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *RankColumnWithDefaultValue1631104895519) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "product_variant" ADD "variant_rank" integer DEFAULT '0'`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_option_value" DROP CONSTRAINT "FK_7234ed737ff4eb1b6ae6e6d7b01"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_option_value" ADD CONSTRAINT "FK_7234ed737ff4eb1b6ae6e6d7b01" FOREIGN KEY ("variant_id") REFERENCES "product_variant"("id") ON DELETE cascade ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" DROP CONSTRAINT "FK_17a06d728e4cfbc5bd2ddb70af0"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" ADD CONSTRAINT "FK_17a06d728e4cfbc5bd2ddb70af0" FOREIGN KEY ("variant_id") REFERENCES "product_variant"("id") ON DELETE cascade ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
func (m *RankColumnWithDefaultValue1631104895519) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "product_variant" DROP COLUMN "variant_rank"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_option_value" DROP CONSTRAINT "FK_7234ed737ff4eb1b6ae6e6d7b01"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_option_value" ADD CONSTRAINT "FK_7234ed737ff4eb1b6ae6e6d7b01" FOREIGN KEY ("variant_id") REFERENCES "product_variant"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" DROP CONSTRAINT "FK_17a06d728e4cfbc5bd2ddb70af0"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" ADD CONSTRAINT "FK_17a06d728e4cfbc5bd2ddb70af0" FOREIGN KEY ("variant_id") REFERENCES "product_variant"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
