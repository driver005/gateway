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

func newCartGiftCard(db *gorm.DB, opts ...gen.DOOption) cartGiftCard {
	_cartGiftCard := cartGiftCard{}

	_cartGiftCard.cartGiftCardDo.UseDB(db, opts...)
	_cartGiftCard.cartGiftCardDo.UseModel(&model.CartGiftCard{})

	tableName := _cartGiftCard.cartGiftCardDo.TableName()
	_cartGiftCard.ALL = field.NewAsterisk(tableName)
	_cartGiftCard.CartID = field.NewString(tableName, "cart_id")
	_cartGiftCard.GiftCardID = field.NewString(tableName, "gift_card_id")

	_cartGiftCard.fillFieldMap()

	return _cartGiftCard
}

type cartGiftCard struct {
	cartGiftCardDo cartGiftCardDo

	ALL        field.Asterisk
	CartID     field.String
	GiftCardID field.String

	fieldMap map[string]field.Expr
}

func (c cartGiftCard) Table(newTableName string) *cartGiftCard {
	c.cartGiftCardDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cartGiftCard) As(alias string) *cartGiftCard {
	c.cartGiftCardDo.DO = *(c.cartGiftCardDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cartGiftCard) updateTableName(table string) *cartGiftCard {
	c.ALL = field.NewAsterisk(table)
	c.CartID = field.NewString(table, "cart_id")
	c.GiftCardID = field.NewString(table, "gift_card_id")

	c.fillFieldMap()

	return c
}

func (c *cartGiftCard) WithContext(ctx context.Context) *cartGiftCardDo {
	return c.cartGiftCardDo.WithContext(ctx)
}

func (c cartGiftCard) TableName() string { return c.cartGiftCardDo.TableName() }

func (c cartGiftCard) Alias() string { return c.cartGiftCardDo.Alias() }

func (c cartGiftCard) Columns(cols ...field.Expr) gen.Columns {
	return c.cartGiftCardDo.Columns(cols...)
}

func (c *cartGiftCard) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cartGiftCard) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 2)
	c.fieldMap["cart_id"] = c.CartID
	c.fieldMap["gift_card_id"] = c.GiftCardID
}

func (c cartGiftCard) clone(db *gorm.DB) cartGiftCard {
	c.cartGiftCardDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cartGiftCard) replaceDB(db *gorm.DB) cartGiftCard {
	c.cartGiftCardDo.ReplaceDB(db)
	return c
}

type cartGiftCardDo struct{ gen.DO }

func (c cartGiftCardDo) Debug() *cartGiftCardDo {
	return c.withDO(c.DO.Debug())
}

func (c cartGiftCardDo) WithContext(ctx context.Context) *cartGiftCardDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cartGiftCardDo) ReadDB() *cartGiftCardDo {
	return c.Clauses(dbresolver.Read)
}

func (c cartGiftCardDo) WriteDB() *cartGiftCardDo {
	return c.Clauses(dbresolver.Write)
}

func (c cartGiftCardDo) Session(config *gorm.Session) *cartGiftCardDo {
	return c.withDO(c.DO.Session(config))
}

func (c cartGiftCardDo) Clauses(conds ...clause.Expression) *cartGiftCardDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cartGiftCardDo) Returning(value interface{}, columns ...string) *cartGiftCardDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cartGiftCardDo) Not(conds ...gen.Condition) *cartGiftCardDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cartGiftCardDo) Or(conds ...gen.Condition) *cartGiftCardDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cartGiftCardDo) Select(conds ...field.Expr) *cartGiftCardDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cartGiftCardDo) Where(conds ...gen.Condition) *cartGiftCardDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cartGiftCardDo) Order(conds ...field.Expr) *cartGiftCardDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cartGiftCardDo) Distinct(cols ...field.Expr) *cartGiftCardDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cartGiftCardDo) Omit(cols ...field.Expr) *cartGiftCardDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cartGiftCardDo) Join(table schema.Tabler, on ...field.Expr) *cartGiftCardDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cartGiftCardDo) LeftJoin(table schema.Tabler, on ...field.Expr) *cartGiftCardDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cartGiftCardDo) RightJoin(table schema.Tabler, on ...field.Expr) *cartGiftCardDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cartGiftCardDo) Group(cols ...field.Expr) *cartGiftCardDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cartGiftCardDo) Having(conds ...gen.Condition) *cartGiftCardDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cartGiftCardDo) Limit(limit int) *cartGiftCardDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cartGiftCardDo) Offset(offset int) *cartGiftCardDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cartGiftCardDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *cartGiftCardDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cartGiftCardDo) Unscoped() *cartGiftCardDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cartGiftCardDo) Create(values ...*model.CartGiftCard) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cartGiftCardDo) CreateInBatches(values []*model.CartGiftCard, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cartGiftCardDo) Save(values ...*model.CartGiftCard) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cartGiftCardDo) First() (*model.CartGiftCard, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CartGiftCard), nil
	}
}

func (c cartGiftCardDo) Take() (*model.CartGiftCard, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CartGiftCard), nil
	}
}

func (c cartGiftCardDo) Last() (*model.CartGiftCard, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CartGiftCard), nil
	}
}

func (c cartGiftCardDo) Find() ([]*model.CartGiftCard, error) {
	result, err := c.DO.Find()
	return result.([]*model.CartGiftCard), err
}

func (c cartGiftCardDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CartGiftCard, err error) {
	buf := make([]*model.CartGiftCard, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cartGiftCardDo) FindInBatches(result *[]*model.CartGiftCard, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cartGiftCardDo) Attrs(attrs ...field.AssignExpr) *cartGiftCardDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cartGiftCardDo) Assign(attrs ...field.AssignExpr) *cartGiftCardDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cartGiftCardDo) Joins(fields ...field.RelationField) *cartGiftCardDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cartGiftCardDo) Preload(fields ...field.RelationField) *cartGiftCardDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cartGiftCardDo) FirstOrInit() (*model.CartGiftCard, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CartGiftCard), nil
	}
}

func (c cartGiftCardDo) FirstOrCreate() (*model.CartGiftCard, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CartGiftCard), nil
	}
}

func (c cartGiftCardDo) FindByPage(offset int, limit int) (result []*model.CartGiftCard, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cartGiftCardDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cartGiftCardDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cartGiftCardDo) Delete(models ...*model.CartGiftCard) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cartGiftCardDo) withDO(do gen.Dao) *cartGiftCardDo {
	c.DO = *do.(*gen.DO)
	return c
}
