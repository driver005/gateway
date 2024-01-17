package migrations

import "reflect"

type ProductCategory1672906846559 struct {
	r Registry
}

func (m *ProductCategory1672906846559) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *ProductCategory1672906846559) Up() error {
	if err := m.r.Context().Exec(`
      CREATE TABLE "product_category"
        (
          "id" character varying NOT NULL,
          "name" text NOT NULL,
          "handle" text NOT NULL,
          "parent_category_id" character varying,
          "mpath" text,
          "is_active" boolean DEFAULT false,
          "is_internal" boolean DEFAULT false,
          "deleted_at" TIMESTAMP WITH TIME ZONE,
          "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
          "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
          CONSTRAINT "PK_qgguwbn1cwstxk93efl0px9oqwt" PRIMARY KEY ("id")
)
    `).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_product_category_handle" ON "product_category" ("handle") WHERE deleted_at IS NULL`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`CREATE INDEX "IDX_product_category_path" ON "product_category" ("mpath")`).Error; err != nil {
		return err
	}
	return nil
}
func (m *ProductCategory1672906846559) Down() error {
	if err := m.r.Context().Exec(`DROP INDEX "IDX_product_category_path"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_product_category_handle"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "product_category"`).Error; err != nil {
		return err
	}
	return nil
}
