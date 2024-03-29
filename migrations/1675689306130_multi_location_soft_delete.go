package migrations

import "reflect"

type MultiLocationSoftDelete1675689306130 struct {
	r Registry
}

func (m *MultiLocationSoftDelete1675689306130) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *MultiLocationSoftDelete1675689306130) Up() error {
	if err := m.r.Context().Exec(`
      ALTER TABLE sales_channel_location
      ADD COLUMN "deleted_at" TIMESTAMP WITH TIME ZONE;

      DROP INDEX "IDX_6caaa358f12ed0b846f00e2dcd";
      DROP INDEX "IDX_c2203162ca946a71aeb98390b0";

      CREATE INDEX "IDX_sales_channel_location_sales_channel_id" ON "sales_channel_location" ("sales_channel_id") WHERE deleted_at IS NULL;
      CREATE INDEX "IDX_sales_channel_location_location_id" ON "sales_channel_location" ("location_id") WHERE deleted_at IS NULL;

      ALTER TABLE product_variant_inventory_item
      ADD COLUMN "deleted_at" TIMESTAMP WITH TIME ZONE;

      DROP INDEX "IDX_c74e8c2835094a37dead376a3b";
      DROP INDEX "IDX_bf5386e7f2acc460adbf96d6f3";

      CREATE INDEX "IDX_product_variant_inventory_item_inventory_item_id" ON "product_variant_inventory_item" ("inventory_item_id") WHERE deleted_at IS NULL;
      CREATE INDEX "IDX_product_variant_inventory_item_variant_id" ON "product_variant_inventory_item" ("variant_id") WHERE deleted_at IS NULL;
    `).Error; err != nil {
		return err
	}
	return nil
}
func (m *MultiLocationSoftDelete1675689306130) Down() error {
	if err := m.r.Context().Exec(`
      DROP INDEX "IDX_sales_channel_location_sales_channel_id";
      DROP INDEX "IDX_sales_channel_location_location_id";

      DROP INDEX "IDX_product_variant_inventory_item_inventory_item_id";
      DROP INDEX "IDX_product_variant_inventory_item_variant_id";

      CREATE INDEX "IDX_6caaa358f12ed0b846f00e2dcd" ON "sales_channel_location" ("sales_channel_id");
      CREATE INDEX "IDX_c2203162ca946a71aeb98390b0" ON "sales_channel_location" ("location_id");

      CREATE INDEX "IDX_c74e8c2835094a37dead376a3b" ON "product_variant_inventory_item" ("inventory_item_id");
      CREATE INDEX "IDX_bf5386e7f2acc460adbf96d6f3" ON "product_variant_inventory_item" ("variant_id");

      ALTER TABLE sales_channel_location
      DROP COLUMN "deleted_at";

      ALTER TABLE product_variant_inventory_item
      DROP COLUMN "deleted_at";
   `).Error; err != nil {
		return err
	}
	return nil
}
