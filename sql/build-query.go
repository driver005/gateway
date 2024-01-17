package sql

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/driver005/gateway/utils"
	"github.com/fatih/structs"
	"github.com/gofiber/fiber/v3"
)

type Options struct {
	Selects       []string `json:"fields,omitempty"`
	Skip          *int     `json:"offset,omitempty"`
	Take          *int     `json:"limit,omitempty"`
	Relations     []string `json:"expand,omitempty"`
	Order         *string  `json:"order,omitempty"`
	Specification []Specification
	Not           []string
	Null          []string
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

func NewOptions(selects []string, skip *int, take *int, relations []string, order *string, secification []Specification) Options {
	return Options{
		Selects:       selects,
		Skip:          skip,
		Take:          take,
		Relations:     relations,
		Order:         order,
		Specification: secification,
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

func FromQuery(context fiber.Ctx) (*Options, *utils.ApplictaionError) {
	var req *Options
	if err := context.Bind().Query(req); err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid query parameters",
			nil,
		)
	}

	return req, nil
}

func BuildQuery[T any](selector T, config *Options) Query {
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
