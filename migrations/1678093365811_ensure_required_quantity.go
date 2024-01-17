package migrations

import "reflect"

type EnsureRequiredQuantity1678093365811 struct {
	r Registry
}

func (m *EnsureRequiredQuantity1678093365811) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *EnsureRequiredQuantity1678093365811) Up() error {
	if err := m.r.Context().Exec(`
        DO
        $$
            BEGIN
                ALTER TABLE product_variant_inventory_item
                    RENAME COLUMN quantity TO required_quantity;
            EXCEPTION
                WHEN undefined_column THEN
            END;
        $$;
    `).Error; err != nil {
		return err
	}
	return nil
}
func (m *EnsureRequiredQuantity1678093365811) Down() error {
	if err := m.r.Context().Exec(`
        DO
        $$
            BEGIN
                ALTER TABLE product_variant_inventory_item
                    RENAME COLUMN required_quantity TO quantity;
            EXCEPTION
                WHEN undefined_column THEN
            END;
        $$;
    `).Error; err != nil {
		return err
	}
	return nil
}
