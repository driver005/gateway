package migrations

import "reflect"

type ExtendedUserApi1633512755401 struct {
	r Registry
}

func (m *ExtendedUserApi1633512755401) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *ExtendedUserApi1633512755401) Up() error {
	if err := m.r.Context().Exec(`CREATE TYPE "invite_role_enum" AS ENUM('admin', 'member', 'developer')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE TABLE "invite" ("id" character varying NOT NULL, "user_email" character varying NOT NULL, "role" "invite_role_enum" DEFAULT 'member', "accepted" boolean NOT NULL DEFAULT false, "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "deleted_at" TIMESTAMP WITH TIME ZONE, "metadata" jsonb, CONSTRAINT "PK_fc9fa190e5a3c5d80604a4f63e1" PRIMARY KEY ("id"))`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "invite" ADD "token" character varying NOT NULL`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "invite" ADD "expires_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`CREATE TYPE "user_role_enum" AS ENUM('admin', 'member', 'developer')`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "user" ADD "role" "user_role_enum" DEFAULT 'member'`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`ALTER TABLE "store" ADD "invite_link_template" character varying`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_6b0ce4b4bcfd24491510bf19d1" ON "invite" ("user_email")`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_e12875dfb3b1d92d7d7c5377e2"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_ba8de19442d86957a3aa3b5006" ON "user" ("email") WHERE deleted_at IS NULL`).Error; err != nil {
		return err
	}
	return nil
}
func (m *ExtendedUserApi1633512755401) Down() error {
	if err := m.r.Context().Exec(`DROP INDEX "IDX_ba8de19442d86957a3aa3b5006"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`CREATE UNIQUE INDEX "IDX_e12875dfb3b1d92d7d7c5377e2" ON "user" ("email") `).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP INDEX "IDX_6b0ce4b4bcfd24491510bf19d1"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "user" DROP COLUMN "role"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "user_role_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "invite" DROP COLUMN "expires_at"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "invite" DROP COLUMN "token"`).Error; err != nil {
		return err
	}

	if err := m.r.Context().Exec(`DROP TABLE "invite"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`DROP TYPE "invite_role_enum"`).Error; err != nil {
		return err
	}
	if err := m.r.Context().Exec(`ALTER TABLE "store" DROP COLUMN "invite_link_template"`).Error; err != nil {
		return err
	}
	return nil
}
