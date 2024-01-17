package migrations

import "reflect"

type StagedJobOptions1673003729870 struct {
	r Registry
}

func (m *StagedJobOptions1673003729870) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *StagedJobOptions1673003729870) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "staged_job" ADD "options" jsonb NOT NULL DEFAULT '{}'::JSONB`).Error; err != nil {
		return err
	}
	return nil
}
func (m *StagedJobOptions1673003729870) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "staged_job" DROP COLUMN "options"`).Error; err != nil {
		return err
	}
	return nil
}
