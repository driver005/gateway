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

func newOrderGiftCard(db *gorm.DB, opts ...gen.DOOption) orderGiftCard {
	_orderGiftCard := orderGiftCard{}

	_orderGiftCard.orderGiftCardDo.UseDB(db, opts...)
	_orderGiftCard.orderGiftCardDo.UseModel(&model.OrderGiftCard{})

	tableName := _orderGiftCard.orderGiftCardDo.TableName()
	_orderGiftCard.ALL = field.NewAsterisk(tableName)
	_orderGiftCard.OrderID = field.NewString(tableName, "order_id")
	_orderGiftCard.GiftCardID = field.NewString(tableName, "gift_card_id")

	_orderGiftCard.fillFieldMap()

	return _orderGiftCard
}

type orderGiftCard struct {
	orderGiftCardDo orderGiftCardDo

	ALL        field.Asterisk
	OrderID    field.String
	GiftCardID field.String

	fieldMap map[string]field.Expr
}

func (o orderGiftCard) Table(newTableName string) *orderGiftCard {
	o.orderGiftCardDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o orderGiftCard) As(alias string) *orderGiftCard {
	o.orderGiftCardDo.DO = *(o.orderGiftCardDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *orderGiftCard) updateTableName(table string) *orderGiftCard {
	o.ALL = field.NewAsterisk(table)
	o.OrderID = field.NewString(table, "order_id")
	o.GiftCardID = field.NewString(table, "gift_card_id")

	o.fillFieldMap()

	return o
}

func (o *orderGiftCard) WithContext(ctx context.Context) *orderGiftCardDo {
	return o.orderGiftCardDo.WithContext(ctx)
}

func (o orderGiftCard) TableName() string { return o.orderGiftCardDo.TableName() }

func (o orderGiftCard) Alias() string { return o.orderGiftCardDo.Alias() }

func (o orderGiftCard) Columns(cols ...field.Expr) gen.Columns {
	return o.orderGiftCardDo.Columns(cols...)
}

func (o *orderGiftCard) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *orderGiftCard) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 2)
	o.fieldMap["order_id"] = o.OrderID
	o.fieldMap["gift_card_id"] = o.GiftCardID
}

func (o orderGiftCard) clone(db *gorm.DB) orderGiftCard {
	o.orderGiftCardDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o orderGiftCard) replaceDB(db *gorm.DB) orderGiftCard {
	o.orderGiftCardDo.ReplaceDB(db)
	return o
}

type orderGiftCardDo struct{ gen.DO }

func (o orderGiftCardDo) Debug() *orderGiftCardDo {
	return o.withDO(o.DO.Debug())
}

func (o orderGiftCardDo) WithContext(ctx context.Context) *orderGiftCardDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o orderGiftCardDo) ReadDB() *orderGiftCardDo {
	return o.Clauses(dbresolver.Read)
}

func (o orderGiftCardDo) WriteDB() *orderGiftCardDo {
	return o.Clauses(dbresolver.Write)
}

func (o orderGiftCardDo) Session(config *gorm.Session) *orderGiftCardDo {
	return o.withDO(o.DO.Session(config))
}

func (o orderGiftCardDo) Clauses(conds ...clause.Expression) *orderGiftCardDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o orderGiftCardDo) Returning(value interface{}, columns ...string) *orderGiftCardDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o orderGiftCardDo) Not(conds ...gen.Condition) *orderGiftCardDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o orderGiftCardDo) Or(conds ...gen.Condition) *orderGiftCardDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o orderGiftCardDo) Select(conds ...field.Expr) *orderGiftCardDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o orderGiftCardDo) Where(conds ...gen.Condition) *orderGiftCardDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o orderGiftCardDo) Order(conds ...field.Expr) *orderGiftCardDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o orderGiftCardDo) Distinct(cols ...field.Expr) *orderGiftCardDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o orderGiftCardDo) Omit(cols ...field.Expr) *orderGiftCardDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o orderGiftCardDo) Join(table schema.Tabler, on ...field.Expr) *orderGiftCardDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o orderGiftCardDo) LeftJoin(table schema.Tabler, on ...field.Expr) *orderGiftCardDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o orderGiftCardDo) RightJoin(table schema.Tabler, on ...field.Expr) *orderGiftCardDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o orderGiftCardDo) Group(cols ...field.Expr) *orderGiftCardDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o orderGiftCardDo) Having(conds ...gen.Condition) *orderGiftCardDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o orderGiftCardDo) Limit(limit int) *orderGiftCardDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o orderGiftCardDo) Offset(offset int) *orderGiftCardDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o orderGiftCardDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *orderGiftCardDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o orderGiftCardDo) Unscoped() *orderGiftCardDo {
	return o.withDO(o.DO.Unscoped())
}

func (o orderGiftCardDo) Create(values ...*model.OrderGiftCard) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o orderGiftCardDo) CreateInBatches(values []*model.OrderGiftCard, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o orderGiftCardDo) Save(values ...*model.OrderGiftCard) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o orderGiftCardDo) First() (*model.OrderGiftCard, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderGiftCard), nil
	}
}

func (o orderGiftCardDo) Take() (*model.OrderGiftCard, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderGiftCard), nil
	}
}

func (o orderGiftCardDo) Last() (*model.OrderGiftCard, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderGiftCard), nil
	}
}

func (o orderGiftCardDo) Find() ([]*model.OrderGiftCard, error) {
	result, err := o.DO.Find()
	return result.([]*model.OrderGiftCard), err
}

func (o orderGiftCardDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OrderGiftCard, err error) {
	buf := make([]*model.OrderGiftCard, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o orderGiftCardDo) FindInBatches(result *[]*model.OrderGiftCard, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o orderGiftCardDo) Attrs(attrs ...field.AssignExpr) *orderGiftCardDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o orderGiftCardDo) Assign(attrs ...field.AssignExpr) *orderGiftCardDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o orderGiftCardDo) Joins(fields ...field.RelationField) *orderGiftCardDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o orderGiftCardDo) Preload(fields ...field.RelationField) *orderGiftCardDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o orderGiftCardDo) FirstOrInit() (*model.OrderGiftCard, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderGiftCard), nil
	}
}

func (o orderGiftCardDo) FirstOrCreate() (*model.OrderGiftCard, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderGiftCard), nil
	}
}

func (o orderGiftCardDo) FindByPage(offset int, limit int) (result []*model.OrderGiftCard, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o orderGiftCardDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o orderGiftCardDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o orderGiftCardDo) Delete(models ...*model.OrderGiftCard) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *orderGiftCardDo) withDO(do gen.Dao) *orderGiftCardDo {
	o.DO = *do.(*gen.DO)
	return o
}
