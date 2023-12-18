package database

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/fatih/structs"
)

type Options[M any] struct {
	selector  *M
	skip      *int
	take      *int
	relations []string
	order     *string
}

type Query[T any] struct {
	selector    *T
	skip        *int
	take        *int
	relations   []string
	where       *string
	withDeleted bool
	order       *string
}

func BuildQuery[T any, M any](selector T, config Options[M]) Query[M] {
	var query Query[M]
	s := structs.New(selector)
	whereString := Build(s.Fields())

	query.where = &whereString
	if !s.Field("DeletedAt").IsZero() {
		query.withDeleted = true
	}

	if config.skip != nil {
		query.skip = config.skip
	}

	if config.take != nil {
		query.take = config.take
	}

	if config.relations != nil {
		query.relations = config.relations
	}

	if config.selector != nil {
		query.selector = config.selector
	}

	if config.order != nil {
		query.order = config.order
	}

	return query
}

func Build(structFields []*structs.Field) string {
	var result string
	for _, field := range structFields {
		name := strings.Replace(field.Tag("json"), ",omitempty", "", -1)
		if field.Kind() == reflect.Struct {
			if !field.IsZero() {
				if field.Name() == "N" {
					result += fmt.Sprintf(`= '%+v'`, field.Value().(time.Time).Format(time.RFC3339))
				} else if field.Name() == "Lt" {
					result += fmt.Sprintf(`< '%+v'`, field.Value().(time.Time).Format(time.RFC3339))
				} else if field.Name() == "Gt" {
					result += fmt.Sprintf(`> '%+v'`, field.Value().(time.Time).Format(time.RFC3339))
				} else if field.Name() == "Lte" {
					result += fmt.Sprintf(`<= '%+v'`, field.Value().(time.Time).Format(time.RFC3339))
				} else if field.Name() == "Gte" {
					result += fmt.Sprintf(`>= '%+v'`, field.Value().(time.Time).Format(time.RFC3339))
				} else {
					if field.IsEmbedded() {
						result += Build(field.Fields())
					} else {
						if len(result) > 0 {
							result += " AND "
						}
						result += fmt.Sprintf("%+v = %+v", name, Build(field.Fields()))
					}
				}
			}
		} else if field.Kind() == reflect.String {
			if !field.IsZero() {
				if field.Name() == "N" {
					result += fmt.Sprintf(`= '%+v'`, field.Value().(string))
				} else if field.Name() == "Lt" {
					result += fmt.Sprintf(`< '%+v'`, field.Value().(string))
				} else if field.Name() == "Gt" {
					result += fmt.Sprintf(`> '%+v'`, field.Value().(string))
				} else if field.Name() == "Lte" {
					result += fmt.Sprintf(`<= '%+v'`, field.Value().(string))
				} else if field.Name() == "Gte" {
					result += fmt.Sprintf(`>= '%+v'`, field.Value().(string))
				} else {
					if len(result) > 0 {
						result += " AND "
					}
					result += fmt.Sprintf(`%+v = '%+v'`, name, field.Value())
				}
			}
		} else if field.Kind() == reflect.Int {
			if !field.IsZero() {
				if field.Name() == "N" {
					result += fmt.Sprintf(`= '%+v'`, field.Value().(int))
				} else if field.Name() == "Lt" {
					result += fmt.Sprintf(`< '%+v'`, field.Value().(int))
				} else if field.Name() == "Gt" {
					result += fmt.Sprintf(`> '%+v'`, field.Value().(int))
				} else if field.Name() == "Lte" {
					result += fmt.Sprintf(`<= '%+v'`, field.Value().(int))
				} else if field.Name() == "Gte" {
					result += fmt.Sprintf(`>= '%+v'`, field.Value().(int))
				} else {
					if len(result) > 0 {
						result += " AND "
					}
					result += fmt.Sprintf(`%+v = %+v`, name, field.Value())

				}
			}
		} else if field.Kind() == reflect.Bool {
			if !field.IsZero() {
				if field.Name() == "N" {
					result += fmt.Sprintf(`= '%+v'`, field.Value().(bool))
				} else if field.Name() == "Lt" {
					result += fmt.Sprintf(`< '%+v'`, field.Value().(bool))
				} else if field.Name() == "Gt" {
					result += fmt.Sprintf(`> '%+v'`, field.Value().(bool))
				} else if field.Name() == "Lte" {
					result += fmt.Sprintf(`<= '%+v'`, field.Value().(bool))
				} else if field.Name() == "Gte" {
					result += fmt.Sprintf(`>= '%+v'`, field.Value().(bool))
				} else {
					if len(result) > 0 {
						result += " AND "
					}
					result += fmt.Sprintf(`%+v = %+v`, name, field.Value())
				}
			}
		} else if field.Kind() == reflect.Slice {
			if !field.IsZero() {
				if len(result) > 0 {
					result += " AND "
				}
				result += fmt.Sprintf(`%+v = %+v`, name, field.Value())
			}
		} else if field.Kind() == reflect.Array {
			if !field.IsZero() {
				if len(result) > 0 {
					result += " AND "
				}
				result += fmt.Sprintf(`%+v = %+v`, name, field.Value())
			}
		}
	}
	return result
}
