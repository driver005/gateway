package migrations

import "reflect"

type AddBatchJobModel1649775522087 struct {
	r Registry
}

func (m *AddBatchJobModel1649775522087) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *AddBatchJobModel1649775522087) Up() error {
	if err := m.r.Context().Exec(
		`CREATE TABLE "batch_job"
         (
             "id"                       character varying                NOT NULL,
             "type"                     text                             NOT NULL,
             "created_by"               character varying,
             "context"                  jsonb,
             "result"                   jsonb,
             "dry_run"                  boolean                          NOT NULL DEFAULT FALSE,
             "created_at"               TIMESTAMP WITH TIME ZONE         NOT NULL DEFAULT now(),
             "pre_processed_at"            TIMESTAMP WITH TIME ZONE,
             "confirmed_at"             TIMESTAMP WITH TIME ZONE,
             "processing_at"            TIMESTAMP WITH TIME ZONE,
             "completed_at"             TIMESTAMP WITH TIME ZONE,
             "failed_at"                TIMESTAMP WITH TIME ZONE,
             "canceled_at"              TIMESTAMP WITH TIME ZONE,
             "updated_at"               TIMESTAMP WITH TIME ZONE         NOT NULL DEFAULT now(),
             "deleted_at"               TIMESTAMP WITH TIME ZONE,
             CONSTRAINT "PK_e57f84d485145d5be96bc6d871e" PRIMARY KEY ("id")
    )`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`ALTER TABLE "batch_job" ADD CONSTRAINT "FK_fa53ca4f5fd90605b532802a626" FOREIGN KEY ("created_by") REFERENCES "user" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
func (m *AddBatchJobModel1649775522087) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "batch_job" DROP CONSTRAINT "FK_fa53ca4f5fd90605b532802a626"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "batch_job"`).Error; err != nil {
		return err
	}
	return nil
}
