package sql

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func InitMigrate(up func(tx *sql.Tx) error, down func(tx *sql.Tx) error) {
	goose.AddMigration(up, down)
}

// func (m *Manager) InitMigrate() {
// 	var err error

// 	m.db, err = gorm.Open(postgres.Open("postgres://jzyozqtc:Hdh78SBKXkvgs6-Z5fpVQAw_la4Iln__@batyr.db.elephantsql.com/jzyozqtc"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect gorm")
// 	}
// }

// func (m *Manager) MigrationStatus(ctx context.Context) (migrate.MigrationStatuses, error) {
// 	status, err := m.m.Status(ctx)
// 	fmt.Println(status)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return status, nil
// }

// func (p *Manager) MigrateDown(ctx context.Context, steps int) error {
// 	return helper.WithStack(p.m.Down(ctx, steps))
// }

// func (p *Manager) MigrateUp(ctx context.Context) error {
// 	if err := p.migrateOldMigrationTables(); err != nil {
// 		return err
// 	}
// 	return helper.WithStack(p.m.Up(ctx))
// }

// func (p *Manager) MigrateUpTo(ctx context.Context, steps int) (int, error) {
// 	if err := p.migrateOldMigrationTables(); err != nil {
// 		return 0, err
// 	}
// 	n, err := p.m.UpTo(ctx, steps)
// 	return n, helper.WithStack(err)
// }

// func (p *Manager) PrepareMigration(_ context.Context) error {
// 	return p.migrateOldMigrationTables()
// }

// type oldTableName string

// const (
// 	clientMigrationTableName  oldTableName = "cards_migration"
// 	// jwkMigrationTableName     oldTableName = "hydra_jwk_migration"
// 	// consentMigrationTableName oldTableName = "hydra_oauth2_authentication_consent_migration"
// 	// oauth2MigrationTableName  oldTableName = "hydra_oauth2_migration"
// )

// // this type is copied from sql-migrate to remove the dependency
// type OldMigrationRecord struct {
// 	ID        string    `db:"id"`
// 	AppliedAt time.Time `db:"applied_at"`
// }

// // this function is idempotent
// func (p *Manager) migrateOldMigrationTables() error {
// 	if err := p.conn.RawQuery(fmt.Sprintf("SELECT * FROM %s", clientMigrationTableName)).Exec(); err != nil {
// 		// assume there are no old migration tables => done
// 		return nil
// 	}

// 	if err := pop.CreateSchemaMigrations(p.conn); err != nil {
// 		return helper.WithStack(err)
// 	}

// 	// in this order the migrations only depend on already done ones
// 	for i, table := range []oldTableName{clientMigrationTableName} {
// 		// If table does not exist, we will skip it. Previously, we created a stub table here which
// 		// caused the cached statements to fail, see:
// 		//
// 		// https://github.com/flynn/flynn/pull/2306/files
// 		// https://github.com/jackc/pgx/issues/110
// 		// https://github.com/flynn/flynn/issues/2235
// 		// get old migrations
// 		var migrations []OldMigrationRecord

// 		q := p.conn.RawQuery(fmt.Sprintf("SELECT * FROM %s", table))
// 		fmt.Println(q)
// 		/* #nosec G201 table is static */
// 		if err := p.conn.RawQuery(fmt.Sprintf("SELECT * FROM %s", table)).All(&migrations); err != nil {
// 			if strings.Contains(err.Error(), string(table)) {
// 				continue
// 			}
// 			return err
// 		}

// 		// translate migrations
// 		for _, m := range migrations {
// 			// mark the migration as run for fizz
// 			// fizz standard version pattern: YYYYMMDDhhmmss
// 			migrationNumber, err := strconv.ParseInt(m.ID, 10, 0)
// 			if err != nil {
// 				return helper.WithStack(err)
// 			}

// 			/* #nosec G201 - i is static (0..3) and migrationNumber is from the gorm */
// 			if err := p.conn.RawQuery(
// 				fmt.Sprintf("INSERT INTO schema_migration (version) VALUES ('2019%02d%08d')", i+1, migrationNumber)).
// 				Exec(); err != nil {
// 				return helper.WithStack(err)
// 			}
// 		}

// 		// delete old migration table
// 		if err := p.conn.RawQuery(fmt.Sprintf("DROP TABLE %s", table)).Exec(); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (m *Manager) AutoMigrate(dst ...interface{}) {
// 	m.db.AutoMigrate(dst...)
// }

// func (m *Manager) CreateConstraint(dst interface{}, field string) bool {
// 	m.db.Migrator().CreateConstraint(&dst, field)
// 	constraint := m.db.Migrator().HasConstraint(&dst, field)
// 	return constraint
// }
