package migrations

import "reflect"

type ExtendedBatchJob1657267320181 struct {
	r Registry
}

func (m *ExtendedBatchJob1657267320181) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *ExtendedBatchJob1657267320181) Up() error {
	var err = m.r.Context().Raw(`
      	SELECT exists (
			SELECT FROM information_schema.columns
			WHERE  table_name = 'batch_job'
			AND    column_name   = 'status'
	)`).Error

	// if the table exists, we alter the table to add the new columns
	if err != nil {
		if err := m.r.Context().Exec(`
        	ALTER TABLE "batch_job" DROP COLUMN "status";
        	DROP TYPE "batch_job_status_enum";
        	ALTER TABLE "batch_job" ADD "dry_run" boolean NOT NULL DEFAULT false;
        	ALTER TABLE "batch_job" ADD "pre_processed_at" TIMESTAMP WITH TIME ZONE;
        	ALTER TABLE "batch_job" ADD "processing_at" TIMESTAMP WITH TIME ZONE;
        	ALTER TABLE "batch_job" ADD "confirmed_at" TIMESTAMP WITH TIME ZONE;
        	ALTER TABLE "batch_job" ADD "completed_at" TIMESTAMP WITH TIME ZONE;
        	ALTER TABLE "batch_job" ADD "canceled_at" TIMESTAMP WITH TIME ZONE;
        	ALTER TABLE "batch_job" ADD "failed_at" TIMESTAMP WITH TIME ZONE;
        	ALTER TABLE "batch_job" DROP COLUMN "created_by";
        	ALTER TABLE "batch_job" ADD "created_by" character varying;
        	ALTER TABLE "batch_job" ADD CONSTRAINT "FK_cdf30493ba1c9ef207e1e80c10a" FOREIGN KEY ("created_by") REFERENCES "user"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
        `).Error; err != nil {
			return err
		}
	}
	return nil
}
func (m *ExtendedBatchJob1657267320181) Down() error {
	return nil
}
