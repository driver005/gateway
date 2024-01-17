package migrations

import "reflect"

type Notifications1613146953072 struct {
	r Registry
}

func (m *Notifications1613146953072) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *Notifications1613146953072) Up() error {
	if err := m.r.Context().Exec(`CREATE TABLE "notification_provider" ("id" character varying NOT NULL, "is_installed" boolean NOT NULL DEFAULT true, CONSTRAINT "PK_0425c2423e2ce9fdfd5c23761d9" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "notification" ("id" character varying NOT NULL, "event_name" character varying, "resource_type" character varying NOT NULL, "resource_id" character varying NOT NULL, "customer_id" character varying, "to" character varying NOT NULL, "data" jsonb NOT NULL, "parent_id" character varying, "provider_id" character varying, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), CONSTRAINT "PK_705b6c7cdf9b2c2ff7ac7872cb7" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_df1494d263740fcfb1d09a98fc" ON "notification" ("resource_type") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_ea6a358d9ce41c16499aae55f9" ON "notification" ("resource_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE INDEX "IDX_b5df0f53a74b9d0c0a2b652c88" ON "notification" ("customer_id") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "notification" ADD CONSTRAINT "FK_b5df0f53a74b9d0c0a2b652c88d" FOREIGN KEY ("customer_id") REFERENCES "customer"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "notification" ADD CONSTRAINT "FK_371db513192c083f48ba63c33be" FOREIGN KEY ("parent_id") REFERENCES "notification"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "notification" ADD CONSTRAINT "FK_0425c2423e2ce9fdfd5c23761d9" FOREIGN KEY ("provider_id") REFERENCES "notification_provider"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
func (m *Notifications1613146953072) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "notification" DROP CONSTRAINT "FK_0425c2423e2ce9fdfd5c23761d9"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "notification" DROP CONSTRAINT "FK_371db513192c083f48ba63c33be"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "notification" DROP CONSTRAINT "FK_b5df0f53a74b9d0c0a2b652c88d"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_b5df0f53a74b9d0c0a2b652c88"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_ea6a358d9ce41c16499aae55f9"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_df1494d263740fcfb1d09a98fc"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "notification"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "notification_provider"`).Error; err != nil {
		return err
	}
	return nil
}
