package migrations

import "reflect"

type DropProductIdFkSalesChannels1680857773273 struct {
	r Registry
}

func (m *DropProductIdFkSalesChannels1680857773273) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *DropProductIdFkSalesChannels1680857773273) Up() error {
	if err := m.r.Context().Exec(`
        alter table if exists "product_sales_channel" drop constraint if exists "FK_5a4d5e1e60f97633547821ec8d6";
    `).Error; err != nil {
		return err
	}
	return nil
}
func (m *DropProductIdFkSalesChannels1680857773273) Down() error {
	if err := m.r.Context().Exec(`
	    ALTER TABLE if exists "product_sales_channel" ADD CONSTRAINT "FK_5a4d5e1e60f97633547821ec8d6" FOREIGN KEY ("product_id") REFERENCES "product"("id") ON DELETE cascade ON UPDATE NO ACTION;
	`).Error; err != nil {
		return err
	}
	return nil
}
