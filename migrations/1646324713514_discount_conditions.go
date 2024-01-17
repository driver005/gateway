package migrations

import "reflect"

type DiscountConditions1646324713514 struct {
	r Registry
}

func (m *DiscountConditions1646324713514) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *DiscountConditions1646324713514) Up() error {
	if err := m.r.Context().Exec(`CREATE TYPE "discount_condition_type_enum" AS ENUM('products', 'product_types', 'product_collections', 'product_tags', 'customer_groups')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TYPE "discount_condition_operator_enum" AS ENUM('in', 'not_in')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "discount_condition" ("id" character varying NOT NULL, "type" "discount_condition_type_enum" NOT NULL, "operator" "discount_condition_operator_enum" NOT NULL, "discount_rule_id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "deleted_at" TIMESTAMP WITH TIME ZONE, "metadata" jsonb, CONSTRAINT "PK_e6b81d83133ddc21a2baf2e2204" PRIMARY KEY ("id"), CONSTRAINT "dctypeuniq" UNIQUE ("type", "operator", "discount_rule_id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_efff700651718e452ca9580a62" ON "discount_condition" ("discount_rule_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "discount_condition_customer_group" ("customer_group_id" character varying NOT NULL, "condition_id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "metadata" jsonb, CONSTRAINT "PK_cdc8b2277169a16b8b7d4c73e0e" PRIMARY KEY ("customer_group_id", "condition_id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "discount_condition_product_collection" ("product_collection_id" character varying NOT NULL, "condition_id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "metadata" jsonb, CONSTRAINT "PK_b3508fc787aa4a38705866cbb6d" PRIMARY KEY ("product_collection_id", "condition_id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "discount_condition_product_tag" ("product_tag_id" character varying NOT NULL, "condition_id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "metadata" jsonb, CONSTRAINT "PK_a95382c1e62205b121aa058682b" PRIMARY KEY ("product_tag_id", "condition_id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "discount_condition_product_type" ("product_type_id" character varying NOT NULL, "condition_id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "metadata" jsonb, CONSTRAINT "PK_35d538a5a24399d0df978df12ed" PRIMARY KEY ("product_type_id", "condition_id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "discount_condition_product" ("product_id" character varying NOT NULL, "condition_id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "metadata" jsonb, CONSTRAINT "PK_994eb4529fdbf14450d64ec17e8" PRIMARY KEY ("product_id", "condition_id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_f05132301e95bdab4ba1cf29a2" ON "discount_condition_product" ("condition_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_c759f53b2e48e8cfb50638fe4e" ON "discount_condition_product" ("product_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_6ef23ce0b1d9cf9b5b833e52b9" ON "discount_condition_product_type" ("condition_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_e706deb68f52ab2756119b9e70" ON "discount_condition_product_type" ("product_type_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_fbb2499551ed074526f3ee3624" ON "discount_condition_product_tag" ("condition_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_01486cc9dc6b36bf658685535f" ON "discount_condition_product_tag" ("product_tag_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_a1c4f9cfb599ad1f0db39cadd5" ON "discount_condition_product_collection" ("condition_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_a0b05dc4257abe639cb75f8eae" ON "discount_condition_product_collection" ("product_collection_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_8486ee16e69013c645d0b8716b" ON "discount_condition_customer_group" ("condition_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_4d5f98645a67545d8dea42e2eb" ON "discount_condition_customer_group" ("customer_group_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition" ADD CONSTRAINT "FK_efff700651718e452ca9580a624" FOREIGN KEY ("discount_rule_id") REFERENCES "discount_rule"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_customer_group" ADD CONSTRAINT "FK_4d5f98645a67545d8dea42e2eb8" FOREIGN KEY ("customer_group_id") REFERENCES "customer_group"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_customer_group" ADD CONSTRAINT "FK_8486ee16e69013c645d0b8716b6" FOREIGN KEY ("condition_id") REFERENCES "discount_condition"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_collection" ADD CONSTRAINT "FK_a0b05dc4257abe639cb75f8eae2" FOREIGN KEY ("product_collection_id") REFERENCES "product_collection"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_collection" ADD CONSTRAINT "FK_a1c4f9cfb599ad1f0db39cadd5f" FOREIGN KEY ("condition_id") REFERENCES "discount_condition"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_tag" ADD CONSTRAINT "FK_01486cc9dc6b36bf658685535f6" FOREIGN KEY ("product_tag_id") REFERENCES "product_tag"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_tag" ADD CONSTRAINT "FK_fbb2499551ed074526f3ee36241" FOREIGN KEY ("condition_id") REFERENCES "discount_condition"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_type" ADD CONSTRAINT "FK_e706deb68f52ab2756119b9e704" FOREIGN KEY ("product_type_id") REFERENCES "product_type"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_type" ADD CONSTRAINT "FK_6ef23ce0b1d9cf9b5b833e52b9d" FOREIGN KEY ("condition_id") REFERENCES "discount_condition"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product" ADD CONSTRAINT "FK_c759f53b2e48e8cfb50638fe4e0" FOREIGN KEY ("product_id") REFERENCES "product"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product" ADD CONSTRAINT "FK_f05132301e95bdab4ba1cf29a24" FOREIGN KEY ("condition_id") REFERENCES "discount_condition"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
func (m *DiscountConditions1646324713514) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product" DROP CONSTRAINT "FK_f05132301e95bdab4ba1cf29a24"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product" DROP CONSTRAINT "FK_c759f53b2e48e8cfb50638fe4e0"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_type" DROP CONSTRAINT "FK_6ef23ce0b1d9cf9b5b833e52b9d"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_type" DROP CONSTRAINT "FK_e706deb68f52ab2756119b9e704"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_tag" DROP CONSTRAINT "FK_fbb2499551ed074526f3ee36241"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_tag" DROP CONSTRAINT "FK_01486cc9dc6b36bf658685535f6"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_collection" DROP CONSTRAINT "FK_a1c4f9cfb599ad1f0db39cadd5f"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_collection" DROP CONSTRAINT "FK_a0b05dc4257abe639cb75f8eae2"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_customer_group" DROP CONSTRAINT "FK_8486ee16e69013c645d0b8716b6"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_customer_group" DROP CONSTRAINT "FK_4d5f98645a67545d8dea42e2eb8"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition" DROP CONSTRAINT "FK_efff700651718e452ca9580a624"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_4d5f98645a67545d8dea42e2eb"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_8486ee16e69013c645d0b8716b"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_a0b05dc4257abe639cb75f8eae"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_a1c4f9cfb599ad1f0db39cadd5"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_01486cc9dc6b36bf658685535f"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_fbb2499551ed074526f3ee3624"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_e706deb68f52ab2756119b9e70"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_6ef23ce0b1d9cf9b5b833e52b9"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_c759f53b2e48e8cfb50638fe4e"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_f05132301e95bdab4ba1cf29a2"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product" DROP COLUMN "metadata"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product" DROP COLUMN "updated_at"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product" DROP COLUMN "created_at"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_type" DROP COLUMN "metadata"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_type" DROP COLUMN "updated_at"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_type" DROP COLUMN "created_at"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_tag" DROP COLUMN "metadata"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_tag" DROP COLUMN "updated_at"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_tag" DROP COLUMN "created_at"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_collection" DROP COLUMN "metadata"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_collection" DROP COLUMN "updated_at"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_collection" DROP COLUMN "created_at"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_customer_group" DROP COLUMN "metadata"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_customer_group" DROP COLUMN "updated_at"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_customer_group" DROP COLUMN "created_at"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_customer_group" ADD "metadata" jsonb`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_customer_group" ADD "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_customer_group" ADD "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_collection" ADD "metadata" jsonb`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_collection" ADD "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_collection" ADD "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_tag" ADD "metadata" jsonb`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_tag" ADD "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_tag" ADD "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_type" ADD "metadata" jsonb`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_type" ADD "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product_type" ADD "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product" ADD "metadata" jsonb`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product" ADD "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "discount_condition_product" ADD "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "discount_condition_product"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "discount_condition_product_type"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "discount_condition_product_tag"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "discount_condition_product_collection"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "discount_condition_customer_group"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_efff700651718e452ca9580a62"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "discount_condition"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "discount_condition_operator_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "discount_condition_type_enum"`).Error; err != nil {
		return err
	}
	return nil
}
