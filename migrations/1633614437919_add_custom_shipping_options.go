package migrations

import "reflect"

type AddCustomShippingOptions1633614437919 struct {
	r Registry
}

func (m *AddCustomShippingOptions1633614437919) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *AddCustomShippingOptions1633614437919) Up() error {
	if err := m.r.Context().Exec(`CREATE TABLE "custom_shipping_option" ("id" character varying NOT NULL, "price" integer NOT NULL, "shipping_option_id" character varying NOT NULL, "cart_id" character varying, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "deleted_at" TIMESTAMP WITH TIME ZONE, "metadata" jsonb, CONSTRAINT "UQ_0f838b122a9a01d921aa1cdb669" UNIQUE ("shipping_option_id", "cart_id"), CONSTRAINT "PK_8dfcb5c1172c29eec4a728420cc" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_44090cb11b06174cbcc667e91c" ON "custom_shipping_option" ("shipping_option_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_93caeb1bb70d37c1d36d6701a7" ON "custom_shipping_option" ("cart_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TYPE "cart_type_enum" RENAME TO "cart_type_enum_old"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TYPE "cart_type_enum" AS ENUM('default', 'swap', 'draft_order', 'payment_link', 'claim')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "cart" ALTER COLUMN "type" DROP DEFAULT`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "cart" ALTER COLUMN "type" TYPE "cart_type_enum" USING "type"::"text"::"cart_type_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "cart" ALTER COLUMN "type" SET DEFAULT 'default'`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "cart_type_enum_old"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "custom_shipping_option" ADD CONSTRAINT "FK_44090cb11b06174cbcc667e91ca" FOREIGN KEY ("shipping_option_id") REFERENCES "shipping_option"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "custom_shipping_option" ADD CONSTRAINT "FK_93caeb1bb70d37c1d36d6701a7a" FOREIGN KEY ("cart_id") REFERENCES "cart"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
func (m *AddCustomShippingOptions1633614437919) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "custom_shipping_option" DROP CONSTRAINT "FK_93caeb1bb70d37c1d36d6701a7a"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "custom_shipping_option" DROP CONSTRAINT "FK_44090cb11b06174cbcc667e91ca"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TYPE "cart_type_enum_old" AS ENUM('default', 'swap', 'draft_order', 'payment_link')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "cart" ALTER COLUMN "type" DROP DEFAULT`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "cart" ALTER COLUMN "type" TYPE "cart_type_enum_old" USING "type"::"text"::"cart_type_enum_old"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "cart" ALTER COLUMN "type" SET DEFAULT 'default'`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "cart_type_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TYPE "cart_type_enum_old" RENAME TO "cart_type_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_93caeb1bb70d37c1d36d6701a7"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_44090cb11b06174cbcc667e91c"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "custom_shipping_option"`).Error; err != nil {
		return err
	}
	return nil
}
