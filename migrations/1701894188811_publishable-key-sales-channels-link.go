package migrations

import "reflect"

type PublishableKeySalesChannelsLink1701894188811 struct {
	r Registry
}

func (m *PublishableKeySalesChannelsLink1701894188811) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *PublishableKeySalesChannelsLink1701894188811) Up() error {
	if err := m.r.Context().Exec(`
        ALTER TABLE "publishable_api_key_sales_channel" ADD COLUMN IF NOT EXISTS "id" text;
        UPDATE "publishable_api_key_sales_channel" SET "id" = 'pksc_' || substr(md5(random()::text), 0, 27) WHERE id is NULL;
        ALTER TABLE "publishable_api_key_sales_channel" ALTER COLUMN "id" SET NOT NULL;

        CREATE INDEX IF NOT EXISTS "IDX_id_publishable_api_key_sales_channel" ON "publishable_api_key_sales_channel" ("id");

        ALTER TABLE "publishable_api_key_sales_channel" ADD COLUMN IF NOT EXISTS "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now();
        ALTER TABLE "publishable_api_key_sales_channel" ADD COLUMN IF NOT EXISTS "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now();
        ALTER TABLE "publishable_api_key_sales_channel" ADD COLUMN IF NOT EXISTS "deleted_at" TIMESTAMP WITH TIME ZONE;
    `).Error; err != nil {
		return err
	}
	return nil
}
func (m *PublishableKeySalesChannelsLink1701894188811) Down() error {
	if err := m.r.Context().Exec(`
        DROP INDEX IF EXISTS "IDX_id_publishable_api_key_sales_channel";
        ALTER TABLE "publishable_api_key_sales_channel" DROP COLUMN IF EXISTS "id";
        ALTER TABLE "publishable_api_key_sales_channel" DROP COLUMN IF EXISTS "created_at";
        ALTER TABLE "publishable_api_key_sales_channel" DROP COLUMN IF EXISTS "updated_at";
        ALTER TABLE "publishable_api_key_sales_channel" DROP COLUMN IF EXISTS "deleted_at";
    `).Error; err != nil {
		return err
	}
	return nil
}
