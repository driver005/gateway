package registry

import (
	"context"

	db "github.com/driver005/gateway/database"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/sql"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type SQL struct {
	*Base
}

var _ Registry = new(SQL)

func init() {
	RegisterDriver(func() Driver {
		return NewRegistrySQL()
	})
}

func NewRegistrySQL() *SQL {
	r := &SQL{
		Base: new(Base),
	}
	r.Base.with(r)
	return r
}

func (m *SQL) Init(ctx context.Context) error {
	if m.database == nil {
		// new db connection
		c := postgres.New(postgres.Config{
			DriverName: "postgres",
			// DSN:        "postgres://postgres:postgres@127.0.0.1:5432/medusa-docker?sslmode=disable",
			DSN: "postgres://jzyozqtc:Hdh78SBKXkvgs6-Z5fpVQAw_la4Iln__@batyr.db.elephantsql.com/jzyozqtc",
		})

		database, err := gorm.Open(c, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   logger.Default.LogMode(logger.Silent),
		})

		if err != nil {
			return helper.WithStack(err)
		}

		m.database, err = sql.NewManager(database, m.l)

		if err != nil {
			return err
		}
	}

	return nil
}

func (m *SQL) CanHandle(dsn string) bool {
	return m.alwaysCanHandle(dsn)
}

func (m *SQL) alwaysCanHandle(dsn string) bool {
	scheme := helper.Split(dsn, "://")[0]
	s := helper.Canonicalize(scheme)
	return s == helper.DriverMySQL || s == helper.DriverPostgreSQL
}

func (m *SQL) Context() *gorm.DB {
	return m.database.DB(context.Background())
}

func (m *SQL) Manager(ctx context.Context) *gorm.DB {
	return m.database.DB(ctx)
}

func (m *SQL) ClientManager() *db.Handler {
	return m.Db()
}