package migrations

import "reflect"

type PublishableApiKey1667815005070 struct {
	r Registry
}

func (m *PublishableApiKey1667815005070) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *PublishableApiKey1667815005070) Up() error {
	if err := m.r.Context().Exec(`CREATE TABLE "publishable_api_key_sales_channel" ("sales_channel_id" character varying NOT NULL, "publishable_key_id" character varying NOT NULL, CONSTRAINT "PK_68eaeb14bdac8954460054c567c" PRIMARY KEY ("sales_channel_id", "publishable_key_id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "publishable_api_key" ("id" character varying NOT NULL, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "created_by" character varying, "revoked_by" character varying, "revoked_at" TIMESTAMP WITH TIME ZONE, "title" character varying NOT NULL, CONSTRAINT "PK_9e613278673a87de92c606b4494" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	return nil
}
func (m *PublishableApiKey1667815005070) Down() error {
	if err := m.r.Context().Exec(`DROP TABLE "publishable_api_key"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TABLE "publishable_api_key_sales_channel"`).Error; err != nil {
		return err
	}
	return nil
}
