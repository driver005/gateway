package database

import (
	"context"

	"gorm.io/gorm"
)

type Handler struct {
	r Registry
}

type Registry interface {
	Manager(ctx context.Context) *gorm.DB
}

func Paginate(offset int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}

func NewHandler(r Registry) *Handler {
	return &Handler{r: r}
}
