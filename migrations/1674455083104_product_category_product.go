package migrations

import "reflect"

type ProductCategoryProduct1674455083104 struct {
	r Registry
}

func (m *ProductCategoryProduct1674455083104) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *ProductCategoryProduct1674455083104) Up() error {
	if err := m.r.Context().Exec(`
        CREATE TABLE "product_category_product" (
          "product_category_id" character varying NOT NULL,
          "product_id" character varying NOT NULL,
          CONSTRAINT "FK_product_category_id" FOREIGN KEY ("product_category_id") REFERENCES product_category("id") ON DELETE CASCADE ON UPDATE NO ACTION,
          CONSTRAINT "FK_product_id" FOREIGN KEY ("product_id") REFERENCES product("id") ON DELETE CASCADE ON UPDATE NO ACTION
)
      `).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`
        CREATE UNIQUE INDEX "IDX_upcp_product_id_product_category_id"
        ON "product_category_product" ("product_category_id", "product_id")
      `).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`
        CREATE INDEX "IDX_pcp_product_category_id"
        ON "product_category_product" ("product_category_id")
      `).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`
        CREATE INDEX "IDX_pcp_product_id"
        ON "product_category_product" ("product_id")
      `).Error; err != nil {
		return err
	}
	return nil
}
func (m *ProductCategoryProduct1674455083104) Down() error {
	if err := m.r.Context().Exec(`DROP INDEX "IDX_upcp_product_id_product_category_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_pcp_product_category_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_pcp_product_id"`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`DROP TABLE "product_category_product"`).Error; err != nil {
		return err
	}
	return nil
}
