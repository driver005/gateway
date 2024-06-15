package utils

import (
	"reflect"

	"gorm.io/gorm"
)

type Query struct {
	Selects     []string
	Skip        int
	Take        int
	Relations   []string
	Where       string
	WithDeleted bool
	Order       string
}

func (q *Query) Parse(db *gorm.DB) *gorm.DB {
	if !reflect.ValueOf(q.Where).IsZero() {
		db = db.Where(q.Where)
	}
	if !reflect.ValueOf(q.WithDeleted).IsZero() {
		db = db.Unscoped()
	}
	if !reflect.ValueOf(q.Skip).IsZero() {
		db = db.Offset(q.Skip)
	}
	if !reflect.ValueOf(q.Take).IsZero() {
		db = db.Limit(q.Take)
	}
	if !reflect.ValueOf(q.Relations).IsZero() {
		for _, relation := range q.Relations {
			db = db.Association(relation).DB
		}
	}
	if !reflect.ValueOf(q.Selects).IsZero() {
		db = db.Select(q.Selects)
	}
	if !reflect.ValueOf(q.Order).IsZero() {
		db = db.Order(q.Order)
	}

	return db
}
