package migrations

import "reflect"

type CustomerGroups1644943746861 struct {
	r Registry
}

func (m *CustomerGroups1644943746861) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *CustomerGroups1644943746861) Up() error {
	if err := m.r.Context().Exec(`CREATE TABLE "customer_group" ("id" character varying NOT NULL, "name" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "deleted_at" TIMESTAMP WITH TIME ZONE, "metadata" jsonb, CONSTRAINT "PK_88e7da3ff7262d9e0a35aa3664e" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_c4c3a5225a7a1f0af782c40abc" ON "customer_group" ("name") WHERE deleted_at IS NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "customer_group_customers" ("customer_group_id" character varying NOT NULL, "customer_id" character varying NOT NULL, CONSTRAINT "PK_e28a55e34ad1e2d3df9a0ac86d3" PRIMARY KEY ("customer_group_id", "customer_id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_620330964db8d2999e67b0dbe3" ON "customer_group_customers" ("customer_group_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_3c6412d076292f439269abe1a2" ON "customer_group_customers" ("customer_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "customer_group_customers" ADD CONSTRAINT "FK_620330964db8d2999e67b0dbe3e" FOREIGN KEY ("customer_group_id") REFERENCES "customer_group"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "customer_group_customers" ADD CONSTRAINT "FK_3c6412d076292f439269abe1a23" FOREIGN KEY ("customer_id") REFERENCES "customer"("id") ON DELETE CASCADE ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
func (m *CustomerGroups1644943746861) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "customer_group_customers" DROP CONSTRAINT "FK_3c6412d076292f439269abe1a23"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "customer_group_customers" DROP CONSTRAINT "FK_620330964db8d2999e67b0dbe3e"`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`DROP INDEX "IDX_3c6412d076292f439269abe1a2"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_620330964db8d2999e67b0dbe3"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "customer_group_customers"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_c4c3a5225a7a1f0af782c40abc"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "customer_group"`).Error; err != nil {
		return err
	}
	return nil
}
