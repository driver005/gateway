package migrations

import "reflect"

type ProductDomainImpovedIndexes1679950645254 struct {
	r Registry
}

func (m *ProductDomainImpovedIndexes1679950645254) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *ProductDomainImpovedIndexes1679950645254) Up() error {
	if err := m.r.Context().Exec(`
      /* You can turn of this settings if you are in a context with lots of variants) set enable_nestloop to off; */
      
      DROP INDEX IF EXISTS "IDX_17a06d728e4cfbc5bd2ddb70af";
      CREATE INDEX IF NOT EXISTS idx_money_amount_variant_id ON money_amount (variant_id);

      DROP INDEX IF EXISTS "IDX_b433e27b7a83e6d12ab26b15b0";
      CREATE INDEX IF NOT EXISTS idx_money_amount_region_id ON money_amount (region_id);

      DROP INDEX IF EXISTS "IDX_7234ed737ff4eb1b6ae6e6d7b0";
      CREATE INDEX IF NOT EXISTS idx_product_option_value_variant_id ON product_option_value (variant_id);

      DROP INDEX IF EXISTS "IDX_cdf4388f294b30a25c627d69fe";
      CREATE INDEX IF NOT EXISTS idx_product_option_value_option_id ON product_option_value (option_id);
    `).Error; err != nil {
		return err
	}
	return nil
}
func (m *ProductDomainImpovedIndexes1679950645254) Down() error {
	if err := m.r.Context().Exec(`
      DROP INDEX IF EXISTS idx_money_amount_variant_id;
      DROP INDEX IF EXISTS idx_money_amount_region_id;
      DROP INDEX IF EXISTS idx_product_option_value_variant_id;
      DROP INDEX IF EXISTS idx_product_option_value_option_id;
      
      CREATE INDEX IF NOT EXISTS "IDX_17a06d728e4cfbc5bd2ddb70af" ON "money_amount" ("variant_id");
      CREATE INDEX IF NOT EXISTS "IDX_b433e27b7a83e6d12ab26b15b0" ON "money_amount" ("region_id");
      CREATE INDEX IF NOT EXISTS "IDX_7234ed737ff4eb1b6ae6e6d7b0" ON "product_option_value" ("variant_id");
      CREATE INDEX IF NOT EXISTS "IDX_cdf4388f294b30a25c627d69fe" ON "product_option_value" ("option_id");
    `).Error; err != nil {
		return err
	}
	return nil
}
