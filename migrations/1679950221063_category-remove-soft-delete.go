package migrations

import "reflect"

type CategoryRemoveSoftDelete1679950221063 struct {
	r Registry
}

func (m *CategoryRemoveSoftDelete1679950221063) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *CategoryRemoveSoftDelete1679950221063) Up() error {
	if err := m.r.Context().Exec(`DELETE FROM "product_category" WHERE "deleted_at" IS NOT NULL;`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_category" DROP COLUMN "deleted_at";`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX IF EXISTS "IDX_product_category_handle";`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_product_category_handle" ON "product_category" ("handle");`).Error; err != nil {
		return err
	}
	return nil
}
func (m *CategoryRemoveSoftDelete1679950221063) Down() error {
	if err := m.r.Context().Exec(`DROP INDEX IF EXISTS "IDX_product_category_handle";`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "product_category" ADD COLUMN "deleted_at" timestamp with time zone;`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_product_category_handle" ON "product_category" ("handle") WHERE deleted_at IS NULL;`).Error; err != nil {
		return err
	}
	return nil
}
