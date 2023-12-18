package repository

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/fatih/structs"
)

type Options struct {
	Selects   []string
	Skip      *int
	Take      *int
	Relations []string
	Order     *string
}

type Query struct {
	Selects     []string
	Skip        *int
	Take        *int
	Relations   []string
	Where       *string
	WithDeleted bool
	Order       *string
}

func NewOptions(selects []string, skip *int, take *int, relations []string, order *string) Options {
	return Options{
		Selects:   selects,
		Skip:      skip,
		Take:      take,
		Relations: relations,
		Order:     order,
	}
}

func NewQuery(where *string, selects []string, skip *int, take *int, relations []string, order *string, withDeleted bool) Query {
	return Query{
		Selects:     selects,
		Skip:        skip,
		Take:        take,
		Relations:   relations,
		Where:       where,
		WithDeleted: withDeleted,
		Order:       order,
	}
}

func BuildQuery[T any](selector T, config Options) Query {
	s := structs.New(selector)
	whereString := Build(s.Fields())

	return NewQuery(
		&whereString,
		config.Selects,
		config.Skip,
		config.Take,
		config.Relations,
		config.Order,
		!s.Field("DeletedAt").IsZero(),
	)
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

					if field.Value().(string)[0:2] == "IN" {
						result += fmt.Sprintf(`%+v IN '%+v'`, name, field.Value())
					} else {
						result += fmt.Sprintf(`%+v = '%+v'`, name, field.Value())
					}

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
				result += fmt.Sprintf(`%+v IN %+v`, name, field.Value())
			}
		}
	}
	return result
}

func ILike(value string) string {
	return fmt.Sprintf(`Like '%+v'`, value)
}
