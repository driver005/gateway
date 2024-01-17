package migrations

import "reflect"

type SoftDeletingUniqueConstraints1624610325746 struct {
	r Registry
}

func (m *SoftDeletingUniqueConstraints1624610325746) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *SoftDeletingUniqueConstraints1624610325746) Up() error {
	if err := m.r.Context().Exec(`DROP INDEX "IDX_6910923cb678fd6e99011a21cc"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_db7355f7bd36c547c8a4f539e5"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_087926f6fec32903be3c8eedfa"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_f4dc2c0888b66d547c175f090e"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_9db95c4b71f632fc93ecbc3d8b"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_7124082c8846a06a857cca386c"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_a0a3f124dc5b167622217fee02"`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_e08af711f3493df1e921c4c9ef" ON "product_collection" ("handle") WHERE deleted_at IS NOT NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_77c4073c30ea7793f484750529" ON "product" ("handle") WHERE deleted_at IS NOT NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_ae3e22c67d7c7a969a363533c0" ON "discount" ("code") WHERE deleted_at IS NOT NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_0683952543d7d3f4fffc427034" ON "product_variant" ("sku") WHERE deleted_at IS NOT NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_410649600ce31c10c4b667ca10" ON "product_variant" ("barcode") WHERE deleted_at IS NOT NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_5248fda27b9f16ef818604bb6f" ON "product_variant" ("ean") WHERE deleted_at IS NOT NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_832f86daf8103491d634a967da" ON "product_variant" ("upc") WHERE deleted_at IS NOT NULL`).Error; err != nil {
		return err
	}
	return nil
}
func (m *SoftDeletingUniqueConstraints1624610325746) Down() error {
	if err := m.r.Context().Exec(`DROP INDEX "IDX_ae3e22c67d7c7a969a363533c0"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_77c4073c30ea7793f484750529"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_e08af711f3493df1e921c4c9ef"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_832f86daf8103491d634a967da"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_5248fda27b9f16ef818604bb6f"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_410649600ce31c10c4b667ca10"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_0683952543d7d3f4fffc427034"`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_087926f6fec32903be3c8eedfa" ON "discount" ("code") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_db7355f7bd36c547c8a4f539e5" ON "product" ("handle") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_6910923cb678fd6e99011a21cc" ON "product_collection" ("handle") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_a0a3f124dc5b167622217fee02" ON "product_variant" ("upc") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_7124082c8846a06a857cca386c" ON "product_variant" ("ean") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_9db95c4b71f632fc93ecbc3d8b" ON "product_variant" ("barcode") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_f4dc2c0888b66d547c175f090e" ON "product_variant" ("sku") `).Error; err != nil {
		return err
	}
	return nil
}
