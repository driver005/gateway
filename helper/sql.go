package helper

import (
	"errors"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func ParseError(err error) error {
	pgError := err.(*pq.Error)

	if !errors.Is(err, pgError) {
		return WithStack(err)
	}

	return errors.New(pgError.Detail)
}

func Paginate(offset int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}
