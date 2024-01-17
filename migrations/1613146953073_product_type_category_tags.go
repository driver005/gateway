package migrations

import "reflect"

type ProductTypeCategoryTags1613146953073 struct {
	r Registry
}

func (m *ProductTypeCategoryTags1613146953073) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *ProductTypeCategoryTags1613146953073) Up() error {
	if err := m.r.Context().Exec(`CREATE TABLE "product_collection" ("id" character varying NOT NULL, "title" character varying NOT NULL, "handle" character varying, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "deleted_at" TIMESTAMP WITH TIME ZONE, "metadata" jsonb, CONSTRAINT "PK_49d419fc77d3aed46c835c558ac" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_6910923cb678fd6e99011a21cc" ON "product_collection" ("handle") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "product_tag" ("id" character varying NOT NULL, "value" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "deleted_at" TIMESTAMP WITH TIME ZONE, "metadata" jsonb, CONSTRAINT "PK_1439455c6528caa94fcc8564fda" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "product_type" ("id" character varying NOT NULL, "value" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "deleted_at" TIMESTAMP WITH TIME ZONE, "metadata" jsonb, CONSTRAINT "PK_e0843930fbb8854fe36ca39dae1" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "product_tags" ("product_id" character varying NOT NULL, "product_tag_id" character varying NOT NULL, CONSTRAINT "PK_1cf5c9537e7198df494b71b993f" PRIMARY KEY ("product_id", "product_tag_id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_5b0c6fc53c574299ecc7f9ee22" ON "product_tags" ("product_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_21683a063fe82dafdf681ecc9c" ON "product_tags" ("product_tag_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product" DROP COLUMN "tags"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product" ADD "collection_id" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product" ADD "type_id" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product" ADD CONSTRAINT "FK_49d419fc77d3aed46c835c558ac" FOREIGN KEY ("collection_id") REFERENCES "product_collection"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_tags" ADD CONSTRAINT "FK_5b0c6fc53c574299ecc7f9ee22e" FOREIGN KEY ("product_id") REFERENCES "product"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_tags" ADD CONSTRAINT "FK_21683a063fe82dafdf681ecc9c4" FOREIGN KEY ("product_tag_id") REFERENCES "product_tag"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product" ADD CONSTRAINT "FK_e0843930fbb8854fe36ca39dae1" FOREIGN KEY ("type_id") REFERENCES "product_type"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
func (m *ProductTypeCategoryTags1613146953073) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "product_tags" DROP CONSTRAINT "FK_21683a063fe82dafdf681ecc9c4"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_tags" DROP CONSTRAINT "FK_5b0c6fc53c574299ecc7f9ee22e"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product" DROP CONSTRAINT "FK_49d419fc77d3aed46c835c558ac"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product" DROP COLUMN "type_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product" DROP COLUMN "collection_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product" DROP CONSTRAINT if exists "FK_e0843930fbb8854fe36ca39dae1"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product" ADD "tags" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_21683a063fe82dafdf681ecc9c"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_5b0c6fc53c574299ecc7f9ee22"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "product_tags"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "product_type"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "product_tag"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_6910923cb678fd6e99011a21cc"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "product_collection"`).Error; err != nil {
		return err
	}
	return nil
}
