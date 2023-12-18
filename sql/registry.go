package sql

import (
	"context"

	"gorm.io/gorm"
)

type (
	Database interface {
		DB(ctx context.Context) *gorm.DB
	}
	Provider interface {
		Database() Database
	}
)
