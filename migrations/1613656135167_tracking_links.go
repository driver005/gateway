package migrations

import "reflect"

type TrackingLinks1613656135167 struct {
	r Registry
}

func (m *TrackingLinks1613656135167) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *TrackingLinks1613656135167) Up() error {
	if err := m.r.Context().Exec(`CREATE TABLE "tracking_link" ("id" character varying NOT NULL, "url" character varying, "tracking_number" character varying NOT NULL, "fulfillment_id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "deleted_at" TIMESTAMP WITH TIME ZONE, "metadata" jsonb, "idempotency_key" character varying, CONSTRAINT "PK_fcfd77feb9012ec2126d7c0bfb6" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "tracking_link" ADD CONSTRAINT "FK_471e9e4c96e02ba209a307db32b" FOREIGN KEY ("fulfillment_id") REFERENCES "fulfillment"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`).Error; err != nil {
		return err
	}
	return nil
}
func (m *TrackingLinks1613656135167) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "tracking_link" DROP CONSTRAINT "FK_471e9e4c96e02ba209a307db32b"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "tracking_link"`).Error; err != nil {
		return err
	}
	return nil
}
