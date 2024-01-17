package migrations

import "reflect"

type AddAnalyticsConfig1666173221888 struct {
	r Registry
}

func (m *AddAnalyticsConfig1666173221888) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *AddAnalyticsConfig1666173221888) Up() error {
	if err := m.r.Context().Exec(`CREATE TABLE "analytics_config" ("id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "deleted_at" TIMESTAMP WITH TIME ZONE, "user_id" character varying NOT NULL, "opt_out" boolean NOT NULL DEFAULT false, "anonymize" boolean NOT NULL DEFAULT false, CONSTRAINT "PK_93505647c5d7cb479becb810b0f" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_379ca70338ce9991f3affdeedf" ON "analytics_config" ("id", "user_id") WHERE deleted_at IS NULL`).Error; err != nil {
		return err
	}
	return nil
}
func (m *AddAnalyticsConfig1666173221888) Down() error {
	if err := m.r.Context().Exec(`DROP INDEX "IDX_379ca70338ce9991f3affdeedf"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "analytics_config"`).Error; err != nil {
		return err
	}
	return nil
}
