package migrations

import "reflect"

type NewTaxSystem1641636508055 struct {
	r Registry
}

func (m *NewTaxSystem1641636508055) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *NewTaxSystem1641636508055) Up() error {
	if err := m.r.Context().Exec(`CREATE TABLE "tax_rate" ("id" character varying NOT NULL, "rate" real, "code" character varying, "name" character varying NOT NULL, "region_id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "metadata" jsonb, CONSTRAINT "PK_23b71b53f650c0b39e99ccef4fd" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "tax_provider" ("id" character varying NOT NULL, "is_installed" boolean NOT NULL DEFAULT true, CONSTRAINT "PK_b198bf82ba6a317c11763d99b99" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "line_item_tax_line" ("id" character varying NOT NULL, "rate" real NOT NULL, "name" character varying NOT NULL, "code" character varying, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "metadata" jsonb, "item_id" character varying NOT NULL, CONSTRAINT "PK_4a0f4322fcd5ce4af85727f89a8" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_5077fa54b0d037e984385dfe8a" ON "line_item_tax_line" ("item_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "shipping_method_tax_line" ("id" character varying NOT NULL, "rate" real NOT NULL, "name" character varying NOT NULL, "code" character varying, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "metadata" jsonb, "shipping_method_id" character varying NOT NULL, CONSTRAINT "PK_54c94f5908aacbd51cf0a73edb1" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_926ca9f29014af8091722dede0" ON "shipping_method_tax_line" ("shipping_method_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "product_tax_rate" ("product_id" character varying NOT NULL, "rate_id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "metadata" jsonb, CONSTRAINT "PK_326257ce468df46cd5c8c5922e9" PRIMARY KEY ("product_id", "rate_id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "product_type_tax_rate" ("product_type_id" character varying NOT NULL, "rate_id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "metadata" jsonb, CONSTRAINT "PK_ddc9242de1d99bc7674969289f0" PRIMARY KEY ("product_type_id", "rate_id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "shipping_tax_rate" ("shipping_option_id" character varying NOT NULL, "rate_id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "metadata" jsonb, CONSTRAINT "PK_bcd93b14d7e2695365d383f5eae" PRIMARY KEY ("shipping_option_id", "rate_id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "region" ADD "gift_cards_taxable" boolean NOT NULL DEFAULT true`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "region" ADD "automatic_taxes" boolean NOT NULL DEFAULT true`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "region" ADD "tax_provider_id" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item" ADD "is_return" boolean NOT NULL DEFAULT false`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`UPDATE "line_item" SET "is_return" = true WHERE "id" IN (SELECT "id" from "line_item" WHERE "metadata"->>'is_return_line' = 'true')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "order" ALTER COLUMN "tax_rate" DROP NOT NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_2484cf14c437a04586b07e7ddd" ON "product_tax_rate" ("rate_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_1d04aebeabb6a89f87e536a124" ON "product_tax_rate" ("product_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_ece65a774192b34253abc4cd67" ON "product_type_tax_rate" ("rate_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_25a3138bb236f63d9bb6c8ff11" ON "product_type_tax_rate" ("product_type_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_346e0016cf045b998074774764" ON "shipping_tax_rate" ("rate_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_f672727ab020df6c50fb64c1a7" ON "shipping_tax_rate" ("shipping_option_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "tax_rate" ADD CONSTRAINT "FK_b95a1e03b051993d208366cb960" FOREIGN KEY ("region_id") REFERENCES "region"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "region" ADD CONSTRAINT "FK_91f88052197680f9790272aaf5b" FOREIGN KEY ("tax_provider_id") REFERENCES "tax_provider"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item_tax_line" ADD CONSTRAINT "FK_5077fa54b0d037e984385dfe8ad" FOREIGN KEY ("item_id") REFERENCES "line_item"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "shipping_method_tax_line" ADD CONSTRAINT "FK_926ca9f29014af8091722dede08" FOREIGN KEY ("shipping_method_id") REFERENCES "shipping_method"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_tax_rate" ADD CONSTRAINT "FK_1d04aebeabb6a89f87e536a124d" FOREIGN KEY ("product_id") REFERENCES "product"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_tax_rate" ADD CONSTRAINT "FK_2484cf14c437a04586b07e7dddb" FOREIGN KEY ("rate_id") REFERENCES "tax_rate"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_type_tax_rate" ADD CONSTRAINT "FK_25a3138bb236f63d9bb6c8ff111" FOREIGN KEY ("product_type_id") REFERENCES "product_type"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_type_tax_rate" ADD CONSTRAINT "FK_ece65a774192b34253abc4cd672" FOREIGN KEY ("rate_id") REFERENCES "tax_rate"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "shipping_tax_rate" ADD CONSTRAINT "FK_f672727ab020df6c50fb64c1a70" FOREIGN KEY ("shipping_option_id") REFERENCES "shipping_option"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "shipping_tax_rate" ADD CONSTRAINT "FK_346e0016cf045b9980747747645" FOREIGN KEY ("rate_id") REFERENCES "tax_rate"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
func (m *NewTaxSystem1641636508055) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "shipping_tax_rate" DROP CONSTRAINT "FK_346e0016cf045b9980747747645"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "shipping_tax_rate" DROP CONSTRAINT "FK_f672727ab020df6c50fb64c1a70"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_type_tax_rate" DROP CONSTRAINT "FK_ece65a774192b34253abc4cd672"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_type_tax_rate" DROP CONSTRAINT "FK_25a3138bb236f63d9bb6c8ff111"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_tax_rate" DROP CONSTRAINT "FK_2484cf14c437a04586b07e7dddb"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_tax_rate" DROP CONSTRAINT "FK_1d04aebeabb6a89f87e536a124d"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "shipping_method_tax_line" DROP CONSTRAINT "FK_926ca9f29014af8091722dede08"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item_tax_line" DROP CONSTRAINT "FK_5077fa54b0d037e984385dfe8ad"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "region" DROP CONSTRAINT "FK_91f88052197680f9790272aaf5b"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "tax_rate" DROP CONSTRAINT "FK_b95a1e03b051993d208366cb960"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_f672727ab020df6c50fb64c1a7"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_346e0016cf045b998074774764"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_25a3138bb236f63d9bb6c8ff11"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_ece65a774192b34253abc4cd67"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_1d04aebeabb6a89f87e536a124"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_2484cf14c437a04586b07e7ddd"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "order" ALTER COLUMN "tax_rate" SET NOT NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "line_item" DROP COLUMN "is_return"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "region" DROP COLUMN "tax_provider_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "region" DROP COLUMN "automatic_taxes"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "region" DROP COLUMN "gift_cards_taxable"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "shipping_tax_rate"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "product_type_tax_rate"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "product_tax_rate"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_926ca9f29014af8091722dede0"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "shipping_method_tax_line"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_5077fa54b0d037e984385dfe8a"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "line_item_tax_line"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "tax_provider"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "tax_rate"`).Error; err != nil {
		return err
	}
	return nil
}
