package migrations

import "reflect"

type ProductCategoryRank1677234878504 struct {
	r Registry
}

func (m *ProductCategoryRank1677234878504) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *ProductCategoryRank1677234878504) Up() error {
	if err := m.r.Context().Exec(`
      ALTER TABLE "product_category"
      ADD COLUMN "rank" integer DEFAULT '0' NOT NULL CHECK ("rank" >= 0);
    `).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`
      CREATE UNIQUE INDEX "UniqProductCategoryParentIdRank"
      ON "product_category" ("parent_category_id", "rank");
    `).Error; err != nil {
		return err
	}
	return nil
}
func (m *ProductCategoryRank1677234878504) Down() error {
	if err := m.r.Context().Exec(`
      DROP INDEX "UniqProductCategoryParentIdRank";
    `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`
      ALTER TABLE "product_category" DROP COLUMN "rank";
    `).Error; err != nil {
		return err
	}
	return nil
}
