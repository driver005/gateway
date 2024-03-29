package migrations

import "reflect"

type CartSalesChannelsLink1698160215000 struct {
	r Registry
}

func (m *CartSalesChannelsLink1698160215000) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *CartSalesChannelsLink1698160215000) Up() error {
	if err := m.r.Context().Exec(`
        CREATE TABLE IF NOT EXISTS "cart_sales_channel"
        (
            "id"                character varying        NOT NULL,
            "cart_id"           character varying        NOT NULL,
            "sales_channel_id"  character varying        NOT NULL,
            "created_at"        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
            "updated_at"        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
            "deleted_at"        TIMESTAMP WITH TIME ZONE,
                                CONSTRAINT "cart_sales_channel_pk" PRIMARY KEY ("cart_id", "sales_channel_id"),
                                CONSTRAINT "cart_sales_channel_cart_id_unique"  UNIQUE ("cart_id")
    );

        CREATE INDEX IF NOT EXISTS "IDX_id_cart_sales_channel" ON "cart_sales_channel" ("id");

        insert into "cart_sales_channel" (id, cart_id, sales_channel_id)
            (select 'cartsc_' || substr(md5(random()::text), 0, 27), id, sales_channel_id from "cart" WHERE sales_channel_id IS NOT NULL);

        ALTER TABLE IF EXISTS "cart" DROP CONSTRAINT IF EXISTS "FK_a2bd3c26f42e754b9249ba78fd6";

        ALTER TABLE IF EXISTS "store" DROP CONSTRAINT IF EXISTS "FK_61b0f48cccbb5f41c750bac7286";
    `).Error; err != nil {
		return err
	}
	return nil
}
func (m *CartSalesChannelsLink1698160215000) Down() error {
	if err := m.r.Context().Exec(`
        UPDATE "cart" SET "sales_channel_id" = "cart_sales_channel"."sales_channel_id"
            FROM "cart_sales_channel"
        WHERE "cart"."id" = "cart_sales_channel"."cart_id";

        DROP TABLE IF EXISTS "cart_sales_channel";

        ALTER TABLE IF EXISTS "cart" ADD CONSTRAINT "FK_a2bd3c26f42e754b9249ba78fd6" FOREIGN KEY ("sales_channel_id") REFERENCES "sales_channel"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

        ALTER TABLE IF EXISTS "store" ADD CONSTRAINT "FK_61b0f48cccbb5f41c750bac7286" FOREIGN KEY ("default_sales_channel_id") REFERENCES "sales_channel"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
    `).Error; err != nil {
		return err
	}
	return nil
}
