package migrations

import "reflect"

type SalesChannel1656949291839 struct {
	r Registry
}

func (m *SalesChannel1656949291839) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *SalesChannel1656949291839) Up() error {
	if err := m.r.Context().Exec(`CREATE TABLE "sales_channel" ("id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "deleted_at" TIMESTAMP WITH TIME ZONE, "name" character varying NOT NULL, "description" character varying, "is_disabled" boolean NOT NULL DEFAULT false, CONSTRAINT "PK_d1eb0b923ea5a0eb1e0916191f1" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "product_sales_channel" ("product_id" character varying NOT NULL, "sales_channel_id" character varying NOT NULL, CONSTRAINT "PK_fd29b6a8bd641052628dee19583" PRIMARY KEY ("product_id", "sales_channel_id"))`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`CREATE INDEX "IDX_5a4d5e1e60f97633547821ec8d" ON "product_sales_channel" ("product_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_37341bad297fe5cca91f921032" ON "product_sales_channel" ("sales_channel_id") `).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`ALTER TABLE "cart" ADD "sales_channel_id" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "order" ADD "sales_channel_id" character varying`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`ALTER TABLE "store" ADD "default_sales_channel_id" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "store" ADD CONSTRAINT "UQ_61b0f48cccbb5f41c750bac7286" UNIQUE ("default_sales_channel_id")`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`ALTER TABLE "cart" ADD CONSTRAINT "FK_a2bd3c26f42e754b9249ba78fd6" FOREIGN KEY ("sales_channel_id") REFERENCES "sales_channel"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "order" ADD CONSTRAINT "FK_6ff7e874f01b478c115fdd462eb" FOREIGN KEY ("sales_channel_id") REFERENCES "sales_channel"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "store" ADD CONSTRAINT "FK_61b0f48cccbb5f41c750bac7286" FOREIGN KEY ("default_sales_channel_id") REFERENCES "sales_channel"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`ALTER TABLE "product_sales_channel" ADD CONSTRAINT "FK_5a4d5e1e60f97633547821ec8d6" FOREIGN KEY ("product_id") REFERENCES "product"("id") ON DELETE CASCADE ON UPDATE CASCADE`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_sales_channel" ADD CONSTRAINT "FK_37341bad297fe5cca91f921032b" FOREIGN KEY ("sales_channel_id") REFERENCES "sales_channel"("id") ON DELETE CASCADE ON UPDATE CASCADE`).Error; err != nil {
		return err
	}
	return nil
}
func (m *SalesChannel1656949291839) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "product_sales_channel" DROP CONSTRAINT "FK_37341bad297fe5cca91f921032b"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_sales_channel" DROP CONSTRAINT "FK_5a4d5e1e60f97633547821ec8d6"`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`ALTER TABLE "store" DROP COLUMN "default_sales_channel_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "order" DROP COLUMN "sales_channel_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "cart" DROP COLUMN "sales_channel_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "product_sales_channel"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "sales_channel"`).Error; err != nil {
		return err
	}
	return nil
}
