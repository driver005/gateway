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

func newReturnItem(db *gorm.DB, opts ...gen.DOOption) returnItem {
	_returnItem := returnItem{}

	_returnItem.returnItemDo.UseDB(db, opts...)
	_returnItem.returnItemDo.UseModel(&model.ReturnItem{})

	tableName := _returnItem.returnItemDo.TableName()
	_returnItem.ALL = field.NewAsterisk(tableName)
	_returnItem.ReturnID = field.NewString(tableName, "return_id")
	_returnItem.ItemID = field.NewString(tableName, "item_id")
	_returnItem.Quantity = field.NewInt32(tableName, "quantity")
	_returnItem.IsRequested = field.NewBool(tableName, "is_requested")
	_returnItem.RequestedQuantity = field.NewInt32(tableName, "requested_quantity")
	_returnItem.ReceivedQuantity = field.NewInt32(tableName, "received_quantity")
	_returnItem.Metadata = field.NewString(tableName, "metadata")
	_returnItem.ReasonID = field.NewString(tableName, "reason_id")
	_returnItem.Note = field.NewString(tableName, "note")

	_returnItem.fillFieldMap()

	return _returnItem
}

type returnItem struct {
	returnItemDo returnItemDo

	ALL               field.Asterisk
	ReturnID          field.String
	ItemID            field.String
	Quantity          field.Int32
	IsRequested       field.Bool
	RequestedQuantity field.Int32
	ReceivedQuantity  field.Int32
	Metadata          field.String
	ReasonID          field.String
	Note              field.String

	fieldMap map[string]field.Expr
}

func (r returnItem) Table(newTableName string) *returnItem {
	r.returnItemDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r returnItem) As(alias string) *returnItem {
	r.returnItemDo.DO = *(r.returnItemDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *returnItem) updateTableName(table string) *returnItem {
	r.ALL = field.NewAsterisk(table)
	r.ReturnID = field.NewString(table, "return_id")
	r.ItemID = field.NewString(table, "item_id")
	r.Quantity = field.NewInt32(table, "quantity")
	r.IsRequested = field.NewBool(table, "is_requested")
	r.RequestedQuantity = field.NewInt32(table, "requested_quantity")
	r.ReceivedQuantity = field.NewInt32(table, "received_quantity")
	r.Metadata = field.NewString(table, "metadata")
	r.ReasonID = field.NewString(table, "reason_id")
	r.Note = field.NewString(table, "note")

	r.fillFieldMap()

	return r
}

func (r *returnItem) WithContext(ctx context.Context) *returnItemDo {
	return r.returnItemDo.WithContext(ctx)
}

func (r returnItem) TableName() string { return r.returnItemDo.TableName() }

func (r returnItem) Alias() string { return r.returnItemDo.Alias() }

func (r returnItem) Columns(cols ...field.Expr) gen.Columns { return r.returnItemDo.Columns(cols...) }

func (r *returnItem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *returnItem) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 9)
	r.fieldMap["return_id"] = r.ReturnID
	r.fieldMap["item_id"] = r.ItemID
	r.fieldMap["quantity"] = r.Quantity
	r.fieldMap["is_requested"] = r.IsRequested
	r.fieldMap["requested_quantity"] = r.RequestedQuantity
	r.fieldMap["received_quantity"] = r.ReceivedQuantity
	r.fieldMap["metadata"] = r.Metadata
	r.fieldMap["reason_id"] = r.ReasonID
	r.fieldMap["note"] = r.Note
}

func (r returnItem) clone(db *gorm.DB) returnItem {
	r.returnItemDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r returnItem) replaceDB(db *gorm.DB) returnItem {
	r.returnItemDo.ReplaceDB(db)
	return r
}

type returnItemDo struct{ gen.DO }

func (r returnItemDo) Debug() *returnItemDo {
	return r.withDO(r.DO.Debug())
}

func (r returnItemDo) WithContext(ctx context.Context) *returnItemDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r returnItemDo) ReadDB() *returnItemDo {
	return r.Clauses(dbresolver.Read)
}

func (r returnItemDo) WriteDB() *returnItemDo {
	return r.Clauses(dbresolver.Write)
}

func (r returnItemDo) Session(config *gorm.Session) *returnItemDo {
	return r.withDO(r.DO.Session(config))
}

func (r returnItemDo) Clauses(conds ...clause.Expression) *returnItemDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r returnItemDo) Returning(value interface{}, columns ...string) *returnItemDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r returnItemDo) Not(conds ...gen.Condition) *returnItemDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r returnItemDo) Or(conds ...gen.Condition) *returnItemDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r returnItemDo) Select(conds ...field.Expr) *returnItemDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r returnItemDo) Where(conds ...gen.Condition) *returnItemDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r returnItemDo) Order(conds ...field.Expr) *returnItemDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r returnItemDo) Distinct(cols ...field.Expr) *returnItemDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r returnItemDo) Omit(cols ...field.Expr) *returnItemDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r returnItemDo) Join(table schema.Tabler, on ...field.Expr) *returnItemDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r returnItemDo) LeftJoin(table schema.Tabler, on ...field.Expr) *returnItemDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r returnItemDo) RightJoin(table schema.Tabler, on ...field.Expr) *returnItemDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r returnItemDo) Group(cols ...field.Expr) *returnItemDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r returnItemDo) Having(conds ...gen.Condition) *returnItemDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r returnItemDo) Limit(limit int) *returnItemDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r returnItemDo) Offset(offset int) *returnItemDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r returnItemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *returnItemDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r returnItemDo) Unscoped() *returnItemDo {
	return r.withDO(r.DO.Unscoped())
}

func (r returnItemDo) Create(values ...*model.ReturnItem) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r returnItemDo) CreateInBatches(values []*model.ReturnItem, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r returnItemDo) Save(values ...*model.ReturnItem) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r returnItemDo) First() (*model.ReturnItem, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReturnItem), nil
	}
}

func (r returnItemDo) Take() (*model.ReturnItem, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReturnItem), nil
	}
}

func (r returnItemDo) Last() (*model.ReturnItem, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReturnItem), nil
	}
}

func (r returnItemDo) Find() ([]*model.ReturnItem, error) {
	result, err := r.DO.Find()
	return result.([]*model.ReturnItem), err
}

func (r returnItemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ReturnItem, err error) {
	buf := make([]*model.ReturnItem, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r returnItemDo) FindInBatches(result *[]*model.ReturnItem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r returnItemDo) Attrs(attrs ...field.AssignExpr) *returnItemDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r returnItemDo) Assign(attrs ...field.AssignExpr) *returnItemDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r returnItemDo) Joins(fields ...field.RelationField) *returnItemDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r returnItemDo) Preload(fields ...field.RelationField) *returnItemDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r returnItemDo) FirstOrInit() (*model.ReturnItem, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReturnItem), nil
	}
}

func (r returnItemDo) FirstOrCreate() (*model.ReturnItem, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReturnItem), nil
	}
}

func (r returnItemDo) FindByPage(offset int, limit int) (result []*model.ReturnItem, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r returnItemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r returnItemDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r returnItemDo) Delete(models ...*model.ReturnItem) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *returnItemDo) withDO(do gen.Dao) *returnItemDo {
	r.DO = *do.(*gen.DO)
	return r
}
