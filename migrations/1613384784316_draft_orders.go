package migrations

import "reflect"

type DraftOrders1613384784316 struct {
	r Registry
}

func (m *DraftOrders1613384784316) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *DraftOrders1613384784316) Up() error {
	if err := m.r.Context().Exec(`CREATE TYPE "draft_order_status_enum" AS ENUM('open', 'completed')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "draft_order" ("id" character varying NOT NULL, "status" "draft_order_status_enum" NOT NULL DEFAULT 'open', "display_id" SERIAL NOT NULL, "cart_id" character varying, "order_id" character varying, "canceled_at" TIMESTAMP WITH TIME ZONE, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "completed_at" TIMESTAMP WITH TIME ZONE, "metadata" jsonb, "idempotency_key" character varying, CONSTRAINT "REL_5bd11d0e2a9628128e2c26fd0a" UNIQUE ("cart_id"), CONSTRAINT "REL_8f6dd6c49202f1466ebf21e77d" UNIQUE ("order_id"), CONSTRAINT "PK_f478946c183d98f8d88a94cfcd7" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_e87cc617a22ef4edce5601edab" ON "draft_order" ("display_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_5bd11d0e2a9628128e2c26fd0a" ON "draft_order" ("cart_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_8f6dd6c49202f1466ebf21e77d" ON "draft_order" ("order_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "order" ADD "draft_order_id" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "order" ADD CONSTRAINT "UQ_727b872f86c7378474a8fa46147" UNIQUE ("draft_order_id")`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TYPE "cart_type_enum" RENAME TO "cart_type_enum_old"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TYPE "cart_type_enum" AS ENUM('default', 'swap', 'draft_order', 'payment_link')`).Error; err != nil {
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
	if err := m.r.Context().Exec(`COMMENT ON COLUMN "cart"."type" IS NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "draft_order" ADD CONSTRAINT "FK_5bd11d0e2a9628128e2c26fd0a6" FOREIGN KEY ("cart_id") REFERENCES "cart"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "draft_order" ADD CONSTRAINT "FK_8f6dd6c49202f1466ebf21e77da" FOREIGN KEY ("order_id") REFERENCES "order"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "order" ADD CONSTRAINT "FK_727b872f86c7378474a8fa46147" FOREIGN KEY ("draft_order_id") REFERENCES "draft_order"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "store" ADD "payment_link_template" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "shipping_option" ADD "admin_only" boolean NOT NULL DEFAULT false`).Error; err != nil {
		return err
	}
	return nil
}
func (m *DraftOrders1613384784316) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "order" DROP CONSTRAINT "FK_727b872f86c7378474a8fa46147"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "draft_order" DROP CONSTRAINT "FK_8f6dd6c49202f1466ebf21e77da"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "draft_order" DROP CONSTRAINT "FK_5bd11d0e2a9628128e2c26fd0a6"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`COMMENT ON COLUMN "cart"."type" IS NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TYPE "cart_type_enum_old" AS ENUM('default', 'swap', 'payment_link')`).Error; err != nil {
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
	if err := m.r.Context().Exec(`ALTER TYPE "cart_type_enum_old" RENAME TO  "cart_type_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "order" DROP CONSTRAINT "UQ_727b872f86c7378474a8fa46147"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "order" DROP COLUMN "draft_order_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_8f6dd6c49202f1466ebf21e77d"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_5bd11d0e2a9628128e2c26fd0a"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_e87cc617a22ef4edce5601edab"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "draft_order"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "draft_order_status_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "store" DROP COLUMN "payment_link_template"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "shipping_option" DROP COLUMN "admin_only"`).Error; err != nil {
		return err
	}
	return nil
}
