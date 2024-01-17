package migrations

import (
	"fmt"
	"reflect"
)

type ProductSearchGinIndexes1679950645254 struct {
	r Registry
}

func (m *ProductSearchGinIndexes1679950645254) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *ProductSearchGinIndexes1679950645254) Up() error {
	if err := m.r.Context().Exec(`
        CREATE EXTENSION IF NOT EXISTS pg_trgm;
        
        CREATE INDEX IF NOT EXISTS idx_gin_product_title ON product USING gin (title gin_trgm_ops) WHERE deleted_at is null;
        CREATE INDEX IF NOT EXISTS idx_gin_product_description ON product USING gin (description gin_trgm_ops) WHERE deleted_at is null;
        
        CREATE INDEX IF NOT EXISTS idx_gin_product_variant_title ON product_variant USING gin (title gin_trgm_ops) WHERE deleted_at is null;
        CREATE INDEX IF NOT EXISTS idx_gin_product_variant_sku ON product_variant USING gin (sku gin_trgm_ops) WHERE deleted_at is null;
        
        CREATE INDEX IF NOT EXISTS idx_gin_product_collection ON product_collection USING gin (title gin_trgm_ops) WHERE deleted_at is null;
      `).Error; err != nil {
		fmt.Println("Could not create pg_trgm extension or indexes, skipping. If you want to use the pg_trgm extension, please install it manually and then run the migration productSearchGinIndexes1679950645254.")
	}
	return nil
}
func (m *ProductSearchGinIndexes1679950645254) Down() error {
	if err := m.r.Context().Exec(`
		DROP INDEX IF EXISTS idx_gin_product_title;
		DROP INDEX IF EXISTS idx_gin_product_description;
		DROP INDEX IF EXISTS idx_gin_product_variant_title;
		DROP INDEX IF EXISTS idx_gin_product_variant_sku;
		DROP INDEX IF EXISTS idx_gin_product_collection;
    `).Error; err != nil {
		return err
	}
	return nil
}
