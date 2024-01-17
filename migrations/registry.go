package migrations

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Registry interface {
	Context() *gorm.DB
	Logger() *zap.SugaredLogger
}
