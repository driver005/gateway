package sql

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type queryParams struct {
	Search         string `query:"q"`
	Filter         string `query:"filter"`
	Expand         string `query:"expand"`
	Fields         string `query:"fields"`
	Offset         int    `query:"offset,default=1"`
	Limit          int    `query:"limit,default=10"`
	All            bool   `query:"all,default=false"`
	OrderBy        string `query:"order_by,default=id"`
	OrderDirection string `query:"order_direction,default=desc,oneof=desc asc"`
}

const (
	SEARCH   = 1  // Filter response with LIKE query "search={search_phrase}"
	FILTER   = 2  // Filter response by column name values "filter={column_name}:{value}"
	PAGINATE = 4  // Paginate response with page and page_size
	ORDER_BY = 8  // Order response by column name
	ALL      = 15 // Equivalent to SEARCH|FILTER|PAGINATE|ORDER_BY
	tagKey   = "filter"
)

var (
	columnNameRegexp = regexp.MustCompile(`(?m)column:(\w{1,}).*`)
	paramNameRegexp  = regexp.MustCompile(`(?m)param:(\w{1,}).*`)
)

func orderBy(db *gorm.DB, params queryParams) *gorm.DB {
	return db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: "id"},
		Desc:   params.OrderDirection == "desc",
	})
}

func paginate(db *gorm.DB, params queryParams) *gorm.DB {
	if params.All {
		return db
	}

	if params.Offset == 0 {
		params.Offset = 1
	}

	switch {
	case params.Limit > 100:
		params.Limit = 100
	case params.Limit <= 0:
		params.Limit = 10
	}

	offset := (params.Offset - 1) * params.Limit
	return db.Offset(offset).Limit(params.Limit)
}

func getColumnNameForField(field reflect.StructField) string {
	fieldTag := field.Tag.Get("json")
	gormTag := field.Tag.Get("gorm")
	res := columnNameRegexp.FindStringSubmatch(gormTag)
	if len(res) == 2 {
		return res[1]
	}
	return strings.ReplaceAll(fieldTag, ",omitempty", "")
}

func searchField(field reflect.StructField, phrase string) clause.Expression {
	filterTag := field.Tag.Get(tagKey)
	columnName := getColumnNameForField(field)
	if strings.Contains(filterTag, "searchable") {
		return clause.Like{Column: columnName, Value: "%" + phrase + "%"}
	}
	return nil
}

func filterField(field reflect.StructField, phrase string) clause.Expression {
	var paramName string
	if !strings.Contains(field.Tag.Get(tagKey), "filterable") {
		return nil
	}
	columnName := getColumnNameForField(field)
	paramMatch := paramNameRegexp.FindStringSubmatch(field.Tag.Get(tagKey))
	if len(paramMatch) == 2 {
		paramName = paramMatch[1]
	} else {
		paramName = columnName
	}
	re, err := regexp.Compile(fmt.Sprintf(`(?m)%v:(\w{1,}).*`, paramName))
	if err != nil {
		return nil
	}
	filterSubPhraseMatch := re.FindStringSubmatch(phrase)
	if len(filterSubPhraseMatch) == 2 {
		return clause.Eq{Column: columnName, Value: filterSubPhraseMatch[1]}
	}
	return nil
}

func addExpressions(phrase string, modelType reflect.Type,
	operator func(reflect.StructField, string) clause.Expression,
) []clause.Expression {
	numFields := modelType.NumField()
	expressions := make([]clause.Expression, 0, numFields)
	for i := 0; i < numFields; i++ {
		field := modelType.Field(i)
		if field.Type.Kind() == reflect.Struct {
			expressions = append(expressions, addExpressions(phrase, field.Type, operator)...)
		}
		expression := operator(field, phrase)
		if expression != nil {
			expressions = append(expressions, expression)
		}
	}

	return expressions
}

func expressionByField(
	db *gorm.DB, phrase string, modelType reflect.Type,
	operator func(reflect.StructField, string) clause.Expression,
	predicate func(...clause.Expression) clause.Expression,
) *gorm.DB {
	expressions := addExpressions(phrase, modelType, operator)
	if len(expressions) == 1 {
		db = db.Where(predicate(expressions[0]))
	} else if len(expressions) > 1 {
		db = db.Where(predicate(expressions...))
	}
	return db
}

// Filter DB request with query parameters.
// Note: Don't forget to initialize DB Model first, otherwise filter and search won't work
// Example:
//
//	db.Model(&UserModel).Scope(filter.FilterByQuery(ctx, filter.ALL)).Find(&users)
//
// Or if only pagination and order is needed:
//
//	db.Model(&UserModel).Scope(filter.FilterByQuery(ctx, filter.PAGINATION|filter.ORDER_BY)).Find(&users)
//
// And models should have appropriate`fitler` tags:
//
//	type User struct {
//		gorm.Model
//		Username string `gorm:"uniqueIndex" filter:"param:login;searchable;filterable"`
//		// `param` defines custom column name for the query param
//		FullName string `filter:"searchable"`
//	}
func FilterByQuery(context fiber.Ctx, config int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var params queryParams
		if err := context.Bind().Query(&params); err != nil {
			return db
		}

		model := db.Statement.Model

		modelType := reflect.TypeOf(model)
		if model != nil && modelType.Kind() == reflect.Ptr && modelType.Elem().Kind() == reflect.Struct {
			if config&SEARCH > 0 && params.Search != "" {
				db = expressionByField(db, params.Search, modelType.Elem(), searchField, clause.Or)
			}
			if config&FILTER > 0 && params.Filter != "" {
				db = expressionByField(db, params.Filter, modelType.Elem(), filterField, clause.And)
			}
		}

		if config&PAGINATE > 0 {

		}

		if config&ORDER_BY > 0 {
			db = orderBy(db, params)
		}
		if config&PAGINATE > 0 {
			db = paginate(db, params)
		}
		return db
	}
}
