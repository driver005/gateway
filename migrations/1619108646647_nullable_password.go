package migrations

import "reflect"

type NullablePassword1619108646647 struct {
	r Registry
}

func (m *NullablePassword1619108646647) GetName() string {
	return reflect.Indirect(reflect.ValueOf(m)).Type().Name()
}

func (m *NullablePassword1619108646647) Up() error {
	if err := m.r.Context().Exec(`ALTER TABLE "user" ALTER COLUMN "password_hash" DROP NOT NULL`).Error; err != nil {
		return err
	}
	return nil
}
func (m *NullablePassword1619108646647) Down() error {
	if err := m.r.Context().Exec(`ALTER TABLE "user" ALTER COLUMN "password_hash" TYPE character varying`).Error; err != nil {
		return err
	}

	// if err := m.r.Context().Exec(`ALTER TABLE "user" ALTER COLUMN if exists "password_hash" NOT NULL`).Error; err != nil {
	// 	return err
	// }
	return nil
}
