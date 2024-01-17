package migrations

import "reflect"

type CategoryCreateIndexes1679950645253 struct {
	r Registry
}

func (m *CategoryCreateIndexes1679950645253) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *CategoryCreateIndexes1679950645253) Up() error {
	if err := m.r.Context().Exec(`
      CREATE INDEX "IDX_product_category_active_public" ON "product_category" (
        "parent_category_id",
        "is_active",
        "is_internal"
) WHERE (
        ("is_active" IS TRUE) AND
        ("is_internal" IS FALSE)
);
    `).Error; err != nil {
		return err
	}
	return nil
}
func (m *CategoryCreateIndexes1679950645253) Down() error {
	if err := m.r.Context().Exec(`
      DROP INDEX "IDX_product_category_active_public";
    `).Error; err != nil {
		return err
	}
	return nil
}
