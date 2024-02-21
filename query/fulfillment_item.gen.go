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

func newFulfillmentItem(db *gorm.DB, opts ...gen.DOOption) fulfillmentItem {
	_fulfillmentItem := fulfillmentItem{}

	_fulfillmentItem.fulfillmentItemDo.UseDB(db, opts...)
	_fulfillmentItem.fulfillmentItemDo.UseModel(&model.FulfillmentItem{})

	tableName := _fulfillmentItem.fulfillmentItemDo.TableName()
	_fulfillmentItem.ALL = field.NewAsterisk(tableName)
	_fulfillmentItem.FulfillmentID = field.NewString(tableName, "fulfillment_id")
	_fulfillmentItem.ItemID = field.NewString(tableName, "item_id")
	_fulfillmentItem.Quantity = field.NewInt32(tableName, "quantity")

	_fulfillmentItem.fillFieldMap()

	return _fulfillmentItem
}

type fulfillmentItem struct {
	fulfillmentItemDo fulfillmentItemDo

	ALL           field.Asterisk
	FulfillmentID field.String
	ItemID        field.String
	Quantity      field.Int32

	fieldMap map[string]field.Expr
}

func (f fulfillmentItem) Table(newTableName string) *fulfillmentItem {
	f.fulfillmentItemDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fulfillmentItem) As(alias string) *fulfillmentItem {
	f.fulfillmentItemDo.DO = *(f.fulfillmentItemDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fulfillmentItem) updateTableName(table string) *fulfillmentItem {
	f.ALL = field.NewAsterisk(table)
	f.FulfillmentID = field.NewString(table, "fulfillment_id")
	f.ItemID = field.NewString(table, "item_id")
	f.Quantity = field.NewInt32(table, "quantity")

	f.fillFieldMap()

	return f
}

func (f *fulfillmentItem) WithContext(ctx context.Context) *fulfillmentItemDo {
	return f.fulfillmentItemDo.WithContext(ctx)
}

func (f fulfillmentItem) TableName() string { return f.fulfillmentItemDo.TableName() }

func (f fulfillmentItem) Alias() string { return f.fulfillmentItemDo.Alias() }

func (f fulfillmentItem) Columns(cols ...field.Expr) gen.Columns {
	return f.fulfillmentItemDo.Columns(cols...)
}

func (f *fulfillmentItem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fulfillmentItem) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 3)
	f.fieldMap["fulfillment_id"] = f.FulfillmentID
	f.fieldMap["item_id"] = f.ItemID
	f.fieldMap["quantity"] = f.Quantity
}

func (f fulfillmentItem) clone(db *gorm.DB) fulfillmentItem {
	f.fulfillmentItemDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fulfillmentItem) replaceDB(db *gorm.DB) fulfillmentItem {
	f.fulfillmentItemDo.ReplaceDB(db)
	return f
}

type fulfillmentItemDo struct{ gen.DO }

func (f fulfillmentItemDo) Debug() *fulfillmentItemDo {
	return f.withDO(f.DO.Debug())
}

func (f fulfillmentItemDo) WithContext(ctx context.Context) *fulfillmentItemDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fulfillmentItemDo) ReadDB() *fulfillmentItemDo {
	return f.Clauses(dbresolver.Read)
}

func (f fulfillmentItemDo) WriteDB() *fulfillmentItemDo {
	return f.Clauses(dbresolver.Write)
}

func (f fulfillmentItemDo) Session(config *gorm.Session) *fulfillmentItemDo {
	return f.withDO(f.DO.Session(config))
}

func (f fulfillmentItemDo) Clauses(conds ...clause.Expression) *fulfillmentItemDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fulfillmentItemDo) Returning(value interface{}, columns ...string) *fulfillmentItemDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fulfillmentItemDo) Not(conds ...gen.Condition) *fulfillmentItemDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fulfillmentItemDo) Or(conds ...gen.Condition) *fulfillmentItemDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fulfillmentItemDo) Select(conds ...field.Expr) *fulfillmentItemDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fulfillmentItemDo) Where(conds ...gen.Condition) *fulfillmentItemDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fulfillmentItemDo) Order(conds ...field.Expr) *fulfillmentItemDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fulfillmentItemDo) Distinct(cols ...field.Expr) *fulfillmentItemDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fulfillmentItemDo) Omit(cols ...field.Expr) *fulfillmentItemDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fulfillmentItemDo) Join(table schema.Tabler, on ...field.Expr) *fulfillmentItemDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fulfillmentItemDo) LeftJoin(table schema.Tabler, on ...field.Expr) *fulfillmentItemDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fulfillmentItemDo) RightJoin(table schema.Tabler, on ...field.Expr) *fulfillmentItemDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fulfillmentItemDo) Group(cols ...field.Expr) *fulfillmentItemDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fulfillmentItemDo) Having(conds ...gen.Condition) *fulfillmentItemDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fulfillmentItemDo) Limit(limit int) *fulfillmentItemDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fulfillmentItemDo) Offset(offset int) *fulfillmentItemDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fulfillmentItemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *fulfillmentItemDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fulfillmentItemDo) Unscoped() *fulfillmentItemDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fulfillmentItemDo) Create(values ...*model.FulfillmentItem) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fulfillmentItemDo) CreateInBatches(values []*model.FulfillmentItem, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fulfillmentItemDo) Save(values ...*model.FulfillmentItem) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fulfillmentItemDo) First() (*model.FulfillmentItem, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentItem), nil
	}
}

func (f fulfillmentItemDo) Take() (*model.FulfillmentItem, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentItem), nil
	}
}

func (f fulfillmentItemDo) Last() (*model.FulfillmentItem, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentItem), nil
	}
}

func (f fulfillmentItemDo) Find() ([]*model.FulfillmentItem, error) {
	result, err := f.DO.Find()
	return result.([]*model.FulfillmentItem), err
}

func (f fulfillmentItemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FulfillmentItem, err error) {
	buf := make([]*model.FulfillmentItem, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fulfillmentItemDo) FindInBatches(result *[]*model.FulfillmentItem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fulfillmentItemDo) Attrs(attrs ...field.AssignExpr) *fulfillmentItemDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fulfillmentItemDo) Assign(attrs ...field.AssignExpr) *fulfillmentItemDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fulfillmentItemDo) Joins(fields ...field.RelationField) *fulfillmentItemDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fulfillmentItemDo) Preload(fields ...field.RelationField) *fulfillmentItemDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fulfillmentItemDo) FirstOrInit() (*model.FulfillmentItem, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentItem), nil
	}
}

func (f fulfillmentItemDo) FirstOrCreate() (*model.FulfillmentItem, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentItem), nil
	}
}

func (f fulfillmentItemDo) FindByPage(offset int, limit int) (result []*model.FulfillmentItem, count int64, err error) {
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

func (f fulfillmentItemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fulfillmentItemDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fulfillmentItemDo) Delete(models ...*model.FulfillmentItem) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fulfillmentItemDo) withDO(do gen.Dao) *fulfillmentItemDo {
	f.DO = *do.(*gen.DO)
	return f
}
