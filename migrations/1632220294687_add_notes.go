package migrations

import "reflect"

type AddNotes1632220294687 struct {
	r Registry
}

func (m *AddNotes1632220294687) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *AddNotes1632220294687) Up() error {
	if err := m.r.Context().Exec(
		`CREATE TABLE "note" ("id" character varying NOT NULL, "value" character varying NOT NULL, "resource_type" character varying NOT NULL, "resource_id" character varying NOT NULL, "author_id" character varying, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "deleted_at" TIMESTAMP WITH TIME ZONE, "metadata" jsonb, CONSTRAINT "PK_96d0c172a4fba276b1bbed43058" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(
		`CREATE INDEX "IDX_f74980b411cf94af523a72af7d" ON "note" ("resource_type") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(
		`CREATE INDEX "IDX_3287f98befad26c3a7dab088cf" ON "note" ("resource_id") `).Error; err != nil {
		return err
	}
	return nil
}
func (m *AddNotes1632220294687) Down() error {
	if err := m.r.Context().Exec(`DROP INDEX "IDX_3287f98befad26c3a7dab088cf"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_f74980b411cf94af523a72af7d"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "note"`).Error; err != nil {
		return err
	}
	return nil
}
