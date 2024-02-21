// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/driver005/gateway/model"
)

func newAnalyticsConfig(db *gorm.DB, opts ...gen.DOOption) analyticsConfig {
	_analyticsConfig := analyticsConfig{}

	_analyticsConfig.analyticsConfigDo.UseDB(db, opts...)
	_analyticsConfig.analyticsConfigDo.UseModel(&model.AnalyticsConfig{})

	tableName := _analyticsConfig.analyticsConfigDo.TableName()
	_analyticsConfig.ALL = field.NewAsterisk(tableName)
	_analyticsConfig.ID = field.NewString(tableName, "id")
	_analyticsConfig.CreatedAt = field.NewTime(tableName, "created_at")
	_analyticsConfig.UpdatedAt = field.NewTime(tableName, "updated_at")
	_analyticsConfig.DeletedAt = field.NewField(tableName, "deleted_at")
	_analyticsConfig.UserID = field.NewString(tableName, "user_id")
	_analyticsConfig.OptOut = field.NewBool(tableName, "opt_out")
	_analyticsConfig.Anonymize = field.NewBool(tableName, "anonymize")

	_analyticsConfig.fillFieldMap()

	return _analyticsConfig
}

type analyticsConfig struct {
	analyticsConfigDo analyticsConfigDo

	ALL       field.Asterisk
	ID        field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	UserID    field.String
	OptOut    field.Bool
	Anonymize field.Bool

	fieldMap map[string]field.Expr
}

func (a analyticsConfig) Table(newTableName string) *analyticsConfig {
	a.analyticsConfigDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a analyticsConfig) As(alias string) *analyticsConfig {
	a.analyticsConfigDo.DO = *(a.analyticsConfigDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *analyticsConfig) updateTableName(table string) *analyticsConfig {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewString(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.UserID = field.NewString(table, "user_id")
	a.OptOut = field.NewBool(table, "opt_out")
	a.Anonymize = field.NewBool(table, "anonymize")

	a.fillFieldMap()

	return a
}

func (a *analyticsConfig) WithContext(ctx context.Context) *analyticsConfigDo {
	return a.analyticsConfigDo.WithContext(ctx)
}

func (a analyticsConfig) TableName() string { return a.analyticsConfigDo.TableName() }

func (a analyticsConfig) Alias() string { return a.analyticsConfigDo.Alias() }

func (a analyticsConfig) Columns(cols ...field.Expr) gen.Columns {
	return a.analyticsConfigDo.Columns(cols...)
}

func (a *analyticsConfig) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *analyticsConfig) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 7)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["opt_out"] = a.OptOut
	a.fieldMap["anonymize"] = a.Anonymize
}

func (a analyticsConfig) clone(db *gorm.DB) analyticsConfig {
	a.analyticsConfigDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a analyticsConfig) replaceDB(db *gorm.DB) analyticsConfig {
	a.analyticsConfigDo.ReplaceDB(db)
	return a
}

type analyticsConfigDo struct{ gen.DO }

func (a analyticsConfigDo) Debug() *analyticsConfigDo {
	return a.withDO(a.DO.Debug())
}

func (a analyticsConfigDo) WithContext(ctx context.Context) *analyticsConfigDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a analyticsConfigDo) ReadDB() *analyticsConfigDo {
	return a.Clauses(dbresolver.Read)
}

func (a analyticsConfigDo) WriteDB() *analyticsConfigDo {
	return a.Clauses(dbresolver.Write)
}

func (a analyticsConfigDo) Session(config *gorm.Session) *analyticsConfigDo {
	return a.withDO(a.DO.Session(config))
}

func (a analyticsConfigDo) Clauses(conds ...clause.Expression) *analyticsConfigDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a analyticsConfigDo) Returning(value interface{}, columns ...string) *analyticsConfigDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a analyticsConfigDo) Not(conds ...gen.Condition) *analyticsConfigDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a analyticsConfigDo) Or(conds ...gen.Condition) *analyticsConfigDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a analyticsConfigDo) Select(conds ...field.Expr) *analyticsConfigDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a analyticsConfigDo) Where(conds ...gen.Condition) *analyticsConfigDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a analyticsConfigDo) Order(conds ...field.Expr) *analyticsConfigDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a analyticsConfigDo) Distinct(cols ...field.Expr) *analyticsConfigDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a analyticsConfigDo) Omit(cols ...field.Expr) *analyticsConfigDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a analyticsConfigDo) Join(table schema.Tabler, on ...field.Expr) *analyticsConfigDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a analyticsConfigDo) LeftJoin(table schema.Tabler, on ...field.Expr) *analyticsConfigDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a analyticsConfigDo) RightJoin(table schema.Tabler, on ...field.Expr) *analyticsConfigDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a analyticsConfigDo) Group(cols ...field.Expr) *analyticsConfigDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a analyticsConfigDo) Having(conds ...gen.Condition) *analyticsConfigDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a analyticsConfigDo) Limit(limit int) *analyticsConfigDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a analyticsConfigDo) Offset(offset int) *analyticsConfigDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a analyticsConfigDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *analyticsConfigDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a analyticsConfigDo) Unscoped() *analyticsConfigDo {
	return a.withDO(a.DO.Unscoped())
}

func (a analyticsConfigDo) Create(values ...*model.AnalyticsConfig) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a analyticsConfigDo) CreateInBatches(values []*model.AnalyticsConfig, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a analyticsConfigDo) Save(values ...*model.AnalyticsConfig) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a analyticsConfigDo) First() (*model.AnalyticsConfig, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnalyticsConfig), nil
	}
}

func (a analyticsConfigDo) Take() (*model.AnalyticsConfig, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnalyticsConfig), nil
	}
}

func (a analyticsConfigDo) Last() (*model.AnalyticsConfig, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnalyticsConfig), nil
	}
}

func (a analyticsConfigDo) Find() ([]*model.AnalyticsConfig, error) {
	result, err := a.DO.Find()
	return result.([]*model.AnalyticsConfig), err
}

func (a analyticsConfigDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AnalyticsConfig, err error) {
	buf := make([]*model.AnalyticsConfig, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a analyticsConfigDo) FindInBatches(result *[]*model.AnalyticsConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a analyticsConfigDo) Attrs(attrs ...field.AssignExpr) *analyticsConfigDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a analyticsConfigDo) Assign(attrs ...field.AssignExpr) *analyticsConfigDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a analyticsConfigDo) Joins(fields ...field.RelationField) *analyticsConfigDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a analyticsConfigDo) Preload(fields ...field.RelationField) *analyticsConfigDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a analyticsConfigDo) FirstOrInit() (*model.AnalyticsConfig, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnalyticsConfig), nil
	}
}

func (a analyticsConfigDo) FirstOrCreate() (*model.AnalyticsConfig, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnalyticsConfig), nil
	}
}

func (a analyticsConfigDo) FindByPage(offset int, limit int) (result []*model.AnalyticsConfig, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a analyticsConfigDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a analyticsConfigDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a analyticsConfigDo) Delete(models ...*model.AnalyticsConfig) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *analyticsConfigDo) withDO(do gen.Dao) *analyticsConfigDo {
	a.DO = *do.(*gen.DO)
	return a
}
