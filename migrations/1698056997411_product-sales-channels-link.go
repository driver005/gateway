package migrations

import "reflect"

type ProductSalesChannelsLink1698056997411 struct {
	r Registry
}

func (m *ProductSalesChannelsLink1698056997411) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *ProductSalesChannelsLink1698056997411) Up() error {
	if err := m.r.Context().Exec(`
        ALTER TABLE "product_sales_channel" ADD COLUMN IF NOT EXISTS "id" text;
        UPDATE "product_sales_channel" SET "id" = 'prodsc_' || substr(md5(random()::text), 0, 27) WHERE id is NULL;
        ALTER TABLE "product_sales_channel" ALTER COLUMN "id" SET NOT NULL;

        ALTER TABLE "product_sales_channel" DROP CONSTRAINT IF EXISTS "PK_fd29b6a8bd641052628dee19583";
        ALTER TABLE "product_sales_channel" ADD CONSTRAINT "product_sales_channel_pk" PRIMARY KEY (id);
        ALTER TABLE "product_sales_channel" ADD CONSTRAINT "product_sales_channel_product_id_sales_channel_id_unique" UNIQUE (product_id, sales_channel_id);

        ALTER TABLE "product_sales_channel" ADD COLUMN IF NOT EXISTS "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now();
        ALTER TABLE "product_sales_channel" ADD COLUMN IF NOT EXISTS "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now();
        ALTER TABLE "product_sales_channel" ADD COLUMN IF NOT EXISTS "deleted_at" TIMESTAMP WITH TIME ZONE;
    `).Error; err != nil {
		return err
	}
	return nil
}
func (m *ProductSalesChannelsLink1698056997411) Down() error {
	if err := m.r.Context().Exec(`
        ALTER TABLE product_sales_channel DROP CONSTRAINT IF EXISTS "product_sales_channel_pk";
        ALTER TABLE product_sales_channel DROP CONSTRAINT IF EXISTS "product_sales_channel_product_id_sales_channel_id_unique";
        ALTER TABLE product_sales_channel drop column if exists "id";

        ALTER TABLE "product_sales_channel" DROP COLUMN IF EXISTS "created_at";
        ALTER TABLE "product_sales_channel" DROP COLUMN IF EXISTS "updated_at";
        ALTER TABLE "product_sales_channel" DROP COLUMN IF EXISTS "deleted_at";

        ALTER TABLE product_sales_channel ADD CONSTRAINT "PK_product_sales_channel" PRIMARY KEY (product_id, sales_channel_id);
    `).Error; err != nil {
		return err
	}
	return nil
}
