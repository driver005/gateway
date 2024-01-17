package cmd

import (
	"github.com/driver005/gateway/migrations"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Registry interface {
	Context() *gorm.DB
	Logger() *zap.SugaredLogger
	Migration() *migrations.Handler
}
