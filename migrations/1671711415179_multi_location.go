package migrations

import "reflect"

type MultiLocation1671711415179 struct {
	r Registry
}

func (m *MultiLocation1671711415179) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *MultiLocation1671711415179) Up() error {
	if err := m.r.Context().Exec(`CREATE TABLE "sales_channel_location" ("id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "sales_channel_id" text NOT NULL, "location_id" text NOT NULL, CONSTRAINT "PK_afd2c2c52634bc8280a9c9ee533" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_6caaa358f12ed0b846f00e2dcd" ON "sales_channel_location" ("sales_channel_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_c2203162ca946a71aeb98390b0" ON "sales_channel_location" ("location_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "product_variant_inventory_item" ("id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "inventory_item_id" text NOT NULL, "variant_id" text NOT NULL, "quantity" integer NOT NULL DEFAULT '1', CONSTRAINT "UQ_c9be7c1b11a1a729eb51d1b6bca" UNIQUE ("variant_id", "inventory_item_id"), CONSTRAINT "PK_9a1188b8d36f4d198303b4f7efa" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_c74e8c2835094a37dead376a3b" ON "product_variant_inventory_item" ("inventory_item_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_bf5386e7f2acc460adbf96d6f3" ON "product_variant_inventory_item" ("variant_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return" ADD "location_id" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "fulfillment" ADD "location_id" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "store" ADD "default_location_id" character varying`).Error; err != nil {
		return err
	}
	return nil
}
func (m *MultiLocation1671711415179) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "store" DROP COLUMN "default_location_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "fulfillment" DROP COLUMN "location_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return" DROP COLUMN "location_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_bf5386e7f2acc460adbf96d6f3"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_c74e8c2835094a37dead376a3b"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "product_variant_inventory_item"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_c2203162ca946a71aeb98390b0"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_6caaa358f12ed0b846f00e2dcd"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "sales_channel_location"`).Error; err != nil {
		return err
	}
	return nil
}
