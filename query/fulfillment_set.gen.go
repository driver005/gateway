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

func newFulfillmentSet(db *gorm.DB, opts ...gen.DOOption) fulfillmentSet {
	_fulfillmentSet := fulfillmentSet{}

	_fulfillmentSet.fulfillmentSetDo.UseDB(db, opts...)
	_fulfillmentSet.fulfillmentSetDo.UseModel(&model.FulfillmentSet{})

	tableName := _fulfillmentSet.fulfillmentSetDo.TableName()
	_fulfillmentSet.ALL = field.NewAsterisk(tableName)
	_fulfillmentSet.ID = field.NewString(tableName, "id")
	_fulfillmentSet.Name = field.NewString(tableName, "name")
	_fulfillmentSet.Type = field.NewString(tableName, "type")
	_fulfillmentSet.Metadata = field.NewString(tableName, "metadata")
	_fulfillmentSet.CreatedAt = field.NewTime(tableName, "created_at")
	_fulfillmentSet.UpdatedAt = field.NewTime(tableName, "updated_at")
	_fulfillmentSet.DeletedAt = field.NewField(tableName, "deleted_at")

	_fulfillmentSet.fillFieldMap()

	return _fulfillmentSet
}

type fulfillmentSet struct {
	fulfillmentSetDo fulfillmentSetDo

	ALL       field.Asterisk
	ID        field.String
	Name      field.String
	Type      field.String
	Metadata  field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (f fulfillmentSet) Table(newTableName string) *fulfillmentSet {
	f.fulfillmentSetDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fulfillmentSet) As(alias string) *fulfillmentSet {
	f.fulfillmentSetDo.DO = *(f.fulfillmentSetDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fulfillmentSet) updateTableName(table string) *fulfillmentSet {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewString(table, "id")
	f.Name = field.NewString(table, "name")
	f.Type = field.NewString(table, "type")
	f.Metadata = field.NewString(table, "metadata")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.DeletedAt = field.NewField(table, "deleted_at")

	f.fillFieldMap()

	return f
}

func (f *fulfillmentSet) WithContext(ctx context.Context) *fulfillmentSetDo {
	return f.fulfillmentSetDo.WithContext(ctx)
}

func (f fulfillmentSet) TableName() string { return f.fulfillmentSetDo.TableName() }

func (f fulfillmentSet) Alias() string { return f.fulfillmentSetDo.Alias() }

func (f fulfillmentSet) Columns(cols ...field.Expr) gen.Columns {
	return f.fulfillmentSetDo.Columns(cols...)
}

func (f *fulfillmentSet) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fulfillmentSet) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 7)
	f.fieldMap["id"] = f.ID
	f.fieldMap["name"] = f.Name
	f.fieldMap["type"] = f.Type
	f.fieldMap["metadata"] = f.Metadata
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["deleted_at"] = f.DeletedAt
}

func (f fulfillmentSet) clone(db *gorm.DB) fulfillmentSet {
	f.fulfillmentSetDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fulfillmentSet) replaceDB(db *gorm.DB) fulfillmentSet {
	f.fulfillmentSetDo.ReplaceDB(db)
	return f
}

type fulfillmentSetDo struct{ gen.DO }

func (f fulfillmentSetDo) Debug() *fulfillmentSetDo {
	return f.withDO(f.DO.Debug())
}

func (f fulfillmentSetDo) WithContext(ctx context.Context) *fulfillmentSetDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fulfillmentSetDo) ReadDB() *fulfillmentSetDo {
	return f.Clauses(dbresolver.Read)
}

func (f fulfillmentSetDo) WriteDB() *fulfillmentSetDo {
	return f.Clauses(dbresolver.Write)
}

func (f fulfillmentSetDo) Session(config *gorm.Session) *fulfillmentSetDo {
	return f.withDO(f.DO.Session(config))
}

func (f fulfillmentSetDo) Clauses(conds ...clause.Expression) *fulfillmentSetDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fulfillmentSetDo) Returning(value interface{}, columns ...string) *fulfillmentSetDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fulfillmentSetDo) Not(conds ...gen.Condition) *fulfillmentSetDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fulfillmentSetDo) Or(conds ...gen.Condition) *fulfillmentSetDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fulfillmentSetDo) Select(conds ...field.Expr) *fulfillmentSetDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fulfillmentSetDo) Where(conds ...gen.Condition) *fulfillmentSetDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fulfillmentSetDo) Order(conds ...field.Expr) *fulfillmentSetDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fulfillmentSetDo) Distinct(cols ...field.Expr) *fulfillmentSetDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fulfillmentSetDo) Omit(cols ...field.Expr) *fulfillmentSetDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fulfillmentSetDo) Join(table schema.Tabler, on ...field.Expr) *fulfillmentSetDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fulfillmentSetDo) LeftJoin(table schema.Tabler, on ...field.Expr) *fulfillmentSetDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fulfillmentSetDo) RightJoin(table schema.Tabler, on ...field.Expr) *fulfillmentSetDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fulfillmentSetDo) Group(cols ...field.Expr) *fulfillmentSetDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fulfillmentSetDo) Having(conds ...gen.Condition) *fulfillmentSetDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fulfillmentSetDo) Limit(limit int) *fulfillmentSetDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fulfillmentSetDo) Offset(offset int) *fulfillmentSetDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fulfillmentSetDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *fulfillmentSetDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fulfillmentSetDo) Unscoped() *fulfillmentSetDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fulfillmentSetDo) Create(values ...*model.FulfillmentSet) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fulfillmentSetDo) CreateInBatches(values []*model.FulfillmentSet, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fulfillmentSetDo) Save(values ...*model.FulfillmentSet) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fulfillmentSetDo) First() (*model.FulfillmentSet, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentSet), nil
	}
}

func (f fulfillmentSetDo) Take() (*model.FulfillmentSet, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentSet), nil
	}
}

func (f fulfillmentSetDo) Last() (*model.FulfillmentSet, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentSet), nil
	}
}

func (f fulfillmentSetDo) Find() ([]*model.FulfillmentSet, error) {
	result, err := f.DO.Find()
	return result.([]*model.FulfillmentSet), err
}

func (f fulfillmentSetDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FulfillmentSet, err error) {
	buf := make([]*model.FulfillmentSet, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fulfillmentSetDo) FindInBatches(result *[]*model.FulfillmentSet, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fulfillmentSetDo) Attrs(attrs ...field.AssignExpr) *fulfillmentSetDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fulfillmentSetDo) Assign(attrs ...field.AssignExpr) *fulfillmentSetDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fulfillmentSetDo) Joins(fields ...field.RelationField) *fulfillmentSetDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fulfillmentSetDo) Preload(fields ...field.RelationField) *fulfillmentSetDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fulfillmentSetDo) FirstOrInit() (*model.FulfillmentSet, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentSet), nil
	}
}

func (f fulfillmentSetDo) FirstOrCreate() (*model.FulfillmentSet, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentSet), nil
	}
}

func (f fulfillmentSetDo) FindByPage(offset int, limit int) (result []*model.FulfillmentSet, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fulfillmentSetDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fulfillmentSetDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fulfillmentSetDo) Delete(models ...*model.FulfillmentSet) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fulfillmentSetDo) withDO(do gen.Dao) *fulfillmentSetDo {
	f.DO = *do.(*gen.DO)
	return f
}
