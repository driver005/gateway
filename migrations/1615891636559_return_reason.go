package migrations

import "reflect"

type ReturnReason1615891636559 struct {
	r Registry
}

func (m *ReturnReason1615891636559) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *ReturnReason1615891636559) Up() error {
	if err := m.r.Context().Exec(`CREATE TABLE "return_reason" ("id" character varying NOT NULL, "value" character varying NOT NULL, "label" character varying NOT NULL, "description" character varying, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "deleted_at" TIMESTAMP WITH TIME ZONE, "metadata" jsonb, CONSTRAINT "PK_95fd1172973165790903e65660a" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_00605f9d662c06b81c1b60ce24" ON "return_reason" ("value") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return_item" ADD "reason_id" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return_item" ADD "note" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return_item" ADD CONSTRAINT "FK_d742532378a65022e7ceb328828" FOREIGN KEY ("reason_id") REFERENCES "return_reason"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
func (m *ReturnReason1615891636559) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "return_item" DROP CONSTRAINT "FK_d742532378a65022e7ceb328828"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return_item" DROP COLUMN "note"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "return_item" DROP COLUMN "reason_id"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_00605f9d662c06b81c1b60ce24"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "return_reason"`).Error; err != nil {
		return err
	}
	return nil
}
