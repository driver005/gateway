package registry

import (
	"context"
	"sync"

	db "github.com/driver005/gateway/database"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/logger"
	"github.com/driver005/gateway/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var (
	drivers = make([]func() Driver, 0)
	dmtx    sync.Mutex

	// ErrNoResponsibleDriverFound is returned when no driver was found for the provided DSN.
	ErrNoResponsibleDriverFound = errors.New("dsn value requested an unknown driver")
)

type Driver interface {
	// CanHandle returns true if the driver is capable of handling the given DSN or false otherwise.
	CanHandle(dsn string) bool

	Manager(ctx context.Context) *gorm.DB

	Context() *gorm.DB
}

type Registry interface {
	Driver

	Init(ctx context.Context) error
	WithLogger(l *logger.Logger) Registry
	RegisterRoutes(router *fiber.App)
	Setup()
	ClientManager() *db.Handler
	// Service() *service.Handler

	sql.Provider
}

// RegisterDriver registers a driver
func RegisterDriver(d func() Driver) {
	dmtx.Lock()
	drivers = append(drivers, d)
	dmtx.Unlock()
}

// GetDriverFor returns a driver for the given DSN or ErrNoResponsibleDriverFound if no driver was found.
func GetDriverFor(dsn string) (Driver, error) {
	for _, f := range drivers {
		driver := f()
		if driver.CanHandle(dsn) {
			return driver, nil
		}
	}
	return nil, ErrNoResponsibleDriverFound
}

func NewRegistryFromDSN(ctx context.Context, l *logger.Logger) (Registry, error) {
	driver, err := GetDriverFor("postgres://jzyozqtc:Hdh78SBKXkvgs6-Z5fpVQAw_la4Iln__@batyr.db.elephantsql.com/jzyozqt?=columnsWithAlias=true")
	if err != nil {
		return nil, helper.WithStack(err)
	}

	registry, ok := driver.(Registry)
	if !ok {
		return nil, errors.Errorf("driver of type %T does not implement interface Registry", driver)
	}

	registry = registry.WithLogger(l)

	if err := registry.Init(ctx); err != nil {
		return nil, err
	}

	return registry, nil
}

func CallRegistry(ctx context.Context, r Registry) {
	r.Context()
}