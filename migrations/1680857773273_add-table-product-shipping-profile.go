package migrations

import "reflect"

type AddTableProductShippingProfile1680857773273 struct {
	r Registry
}

func (m *AddTableProductShippingProfile1680857773273) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *AddTableProductShippingProfile1680857773273) Up() error {
	if err := m.r.Context().Exec(
		`
        CREATE TABLE IF NOT EXISTS "product_shipping_profile"
        (
            "profile_id" text NOT NULL,
            "product_id" text NOT NULL
);

        INSERT INTO "product_shipping_profile" ("profile_id", "product_id")
        SELECT "profile_id", "id" FROM "product";

        ALTER TABLE "product" DROP COLUMN IF EXISTS "profile_id";
        CREATE UNIQUE INDEX IF NOT EXISTS "idx_product_shipping_profile_profile_id_product_id_unique" ON "product_shipping_profile" ("profile_id", "product_id");
        CREATE INDEX IF NOT EXISTS "idx_product_shipping_profile_product_id" ON "product_shipping_profile" ("product_id");
        CREATE INDEX IF NOT EXISTS "idx_product_shipping_profile_profile_id" ON "product_shipping_profile" ("profile_id");
        DROP INDEX IF EXISTS "IDX_80823b7ae866dc5acae2dac6d2";
      `).Error; err != nil {
		return err
	}
	return nil
}
func (m *AddTableProductShippingProfile1680857773273) Down() error {
	if err := m.r.Context().Exec(`
        DROP INDEX IF EXISTS "idx_product_shipping_profile_profile_id_product_id_unique";
        DROP INDEX IF EXISTS "idx_product_shipping_profile_product_id";
        DROP INDEX IF EXISTS "idx_product_shipping_profile_profile_id";

        ALTER TABLE "product" ADD COLUMN IF NOT EXISTS "profile_id" CHARACTER VARYING;

        UPDATE "product" SET "profile_id" = "product_shipping_profile"."profile_id"
            FROM "product_shipping_profile"
        WHERE "product"."id" = "product_shipping_profile"."product_id";

        DROP TABLE IF EXISTS "product_shipping_profile";

        ALTER TABLE "product" ALTER COLUMN profile_id SET NOT NULL;

        CREATE INDEX IF NOT EXISTS "IDX_80823b7ae866dc5acae2dac6d2" ON "product" ("profile_id");
        ALTER TABLE "product" ADD CONSTRAINT "FK_80823b7ae866dc5acae2dac6d2c" FOREIGN KEY ("profile_id") REFERENCES "shipping_profile"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
    `).Error; err != nil {
		return err
	}
	return nil
}
