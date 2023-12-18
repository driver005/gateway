package sql

import (
	"context"
	"embed"

	"github.com/driver005/gateway/logger"
	"gorm.io/gorm"

	"github.com/pkg/errors"
)

var Migrations embed.FS

const transactionContextKey transactionContextType = "transactionConnection"

var (
	ErrTransactionOpen   = errors.New("There is already a transaction in this context.")
	ErrNoTransactionOpen = errors.New("There is no transaction in this context.")
)

type (
	Manager struct {
		db *gorm.DB
		l  *logger.Logger
	}
	// Dependencies interface {
	// 	Tracer(ctx context.Context) trace.Tracer
	// }
	transactionContextType string
)

func NewManager(db *gorm.DB, l *logger.Logger) (*Manager, error) {
	return &Manager{
		db: db,
		l:  l,
	}, nil
}

func (p *Manager) DB(ctx context.Context) *gorm.DB {
	if c, ok := ctx.Value(transactionContextKey).(*gorm.DB); ok {
		return c.WithContext(ctx)
	}
	return p.db.WithContext(ctx)
}
