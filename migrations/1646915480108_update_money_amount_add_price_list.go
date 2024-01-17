package migrations

import "reflect"

type UpdateMoneyAmountAddPriceList1646915480108 struct {
	r Registry
}

func (m *UpdateMoneyAmountAddPriceList1646915480108) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *UpdateMoneyAmountAddPriceList1646915480108) Up() error {
	if err := m.r.Context().Exec(`CREATE TYPE "price_list_type_enum" AS ENUM('sale', 'override')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TYPE "price_list_status_enum" AS ENUM('active', 'draft')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "price_list" ("id" character varying NOT NULL, "name" character varying NOT NULL, "description" character varying NOT NULL, "type" "price_list_type_enum" NOT NULL DEFAULT 'sale', "status" "price_list_status_enum" NOT NULL DEFAULT 'draft', "starts_at" TIMESTAMP WITH TIME ZONE, "ends_at" TIMESTAMP WITH TIME ZONE, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "deleted_at" TIMESTAMP WITH TIME ZONE, CONSTRAINT "PK_52ea7826468b1c889cb2c28df03" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "price_list_customer_groups" ("price_list_id" character varying NOT NULL, "customer_group_id" character varying NOT NULL, CONSTRAINT "PK_1afcbe15cc8782dc80c05707df9" PRIMARY KEY ("price_list_id", "customer_group_id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_52875734e9dd69064f0041f4d9" ON "price_list_customer_groups" ("price_list_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_c5516f550433c9b1c2630d787a" ON "price_list_customer_groups" ("customer_group_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" DROP COLUMN "sale_amount"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" ADD "min_quantity" integer`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" ADD "max_quantity" integer`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" ADD "price_list_id" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" ADD CONSTRAINT "FK_f249976b079375499662eb80c40" FOREIGN KEY ("price_list_id") REFERENCES "price_list"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "price_list_customer_groups" ADD CONSTRAINT "FK_52875734e9dd69064f0041f4d92" FOREIGN KEY ("price_list_id") REFERENCES "price_list"("id") ON DELETE CASCADE ON UPDATE CASCADE`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "price_list_customer_groups" ADD CONSTRAINT "FK_c5516f550433c9b1c2630d787a7" FOREIGN KEY ("customer_group_id") REFERENCES "customer_group"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
func (m *UpdateMoneyAmountAddPriceList1646915480108) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "price_list_customer_groups" DROP CONSTRAINT "FK_c5516f550433c9b1c2630d787a7"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "price_list_customer_groups" DROP CONSTRAINT "FK_52875734e9dd69064f0041f4d92"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" DROP CONSTRAINT "FK_f249976b079375499662eb80c40"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" DROP COLUMN "price_list_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" DROP COLUMN "max_quantity"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" DROP COLUMN "min_quantity"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "money_amount" ADD "sale_amount" integer`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_c5516f550433c9b1c2630d787a"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_52875734e9dd69064f0041f4d9"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "price_list_customer_groups"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "price_list"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "price_list_status_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "price_list_type_enum"`).Error; err != nil {
		return err
	}
	return nil
}
