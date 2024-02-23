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

func newCartLineItemTaxLine(db *gorm.DB, opts ...gen.DOOption) cartLineItemTaxLine {
	_cartLineItemTaxLine := cartLineItemTaxLine{}

	_cartLineItemTaxLine.cartLineItemTaxLineDo.UseDB(db, opts...)
	_cartLineItemTaxLine.cartLineItemTaxLineDo.UseModel(&model.CartLineItemTaxLine{})

	tableName := _cartLineItemTaxLine.cartLineItemTaxLineDo.TableName()
	_cartLineItemTaxLine.ALL = field.NewAsterisk(tableName)
	_cartLineItemTaxLine.ID = field.NewString(tableName, "id")
	_cartLineItemTaxLine.Description = field.NewString(tableName, "description")
	_cartLineItemTaxLine.TaxRateID = field.NewString(tableName, "tax_rate_id")
	_cartLineItemTaxLine.Code = field.NewString(tableName, "code")
	_cartLineItemTaxLine.Rate = field.NewFloat64(tableName, "rate")
	_cartLineItemTaxLine.ProviderID = field.NewString(tableName, "provider_id")
	_cartLineItemTaxLine.CreatedAt = field.NewTime(tableName, "created_at")
	_cartLineItemTaxLine.UpdatedAt = field.NewTime(tableName, "updated_at")
	_cartLineItemTaxLine.ItemID = field.NewString(tableName, "item_id")

	_cartLineItemTaxLine.fillFieldMap()

	return _cartLineItemTaxLine
}

type cartLineItemTaxLine struct {
	cartLineItemTaxLineDo cartLineItemTaxLineDo

	ALL         field.Asterisk
	ID          field.String
	Description field.String
	TaxRateID   field.String
	Code        field.String
	Rate        field.Float64
	ProviderID  field.String
	CreatedAt   field.Time
	UpdatedAt   field.Time
	ItemID      field.String

	fieldMap map[string]field.Expr
}

func (c cartLineItemTaxLine) Table(newTableName string) *cartLineItemTaxLine {
	c.cartLineItemTaxLineDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cartLineItemTaxLine) As(alias string) *cartLineItemTaxLine {
	c.cartLineItemTaxLineDo.DO = *(c.cartLineItemTaxLineDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cartLineItemTaxLine) updateTableName(table string) *cartLineItemTaxLine {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewString(table, "id")
	c.Description = field.NewString(table, "description")
	c.TaxRateID = field.NewString(table, "tax_rate_id")
	c.Code = field.NewString(table, "code")
	c.Rate = field.NewFloat64(table, "rate")
	c.ProviderID = field.NewString(table, "provider_id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.ItemID = field.NewString(table, "item_id")

	c.fillFieldMap()

	return c
}

func (c *cartLineItemTaxLine) WithContext(ctx context.Context) *cartLineItemTaxLineDo {
	return c.cartLineItemTaxLineDo.WithContext(ctx)
}

func (c cartLineItemTaxLine) TableName() string { return c.cartLineItemTaxLineDo.TableName() }

func (c cartLineItemTaxLine) Alias() string { return c.cartLineItemTaxLineDo.Alias() }

func (c cartLineItemTaxLine) Columns(cols ...field.Expr) gen.Columns {
	return c.cartLineItemTaxLineDo.Columns(cols...)
}

func (c *cartLineItemTaxLine) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cartLineItemTaxLine) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["id"] = c.ID
	c.fieldMap["description"] = c.Description
	c.fieldMap["tax_rate_id"] = c.TaxRateID
	c.fieldMap["code"] = c.Code
	c.fieldMap["rate"] = c.Rate
	c.fieldMap["provider_id"] = c.ProviderID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["item_id"] = c.ItemID
}

func (c cartLineItemTaxLine) clone(db *gorm.DB) cartLineItemTaxLine {
	c.cartLineItemTaxLineDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cartLineItemTaxLine) replaceDB(db *gorm.DB) cartLineItemTaxLine {
	c.cartLineItemTaxLineDo.ReplaceDB(db)
	return c
}

type cartLineItemTaxLineDo struct{ gen.DO }

func (c cartLineItemTaxLineDo) Debug() *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Debug())
}

func (c cartLineItemTaxLineDo) WithContext(ctx context.Context) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cartLineItemTaxLineDo) ReadDB() *cartLineItemTaxLineDo {
	return c.Clauses(dbresolver.Read)
}

func (c cartLineItemTaxLineDo) WriteDB() *cartLineItemTaxLineDo {
	return c.Clauses(dbresolver.Write)
}

func (c cartLineItemTaxLineDo) Session(config *gorm.Session) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Session(config))
}

func (c cartLineItemTaxLineDo) Clauses(conds ...clause.Expression) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cartLineItemTaxLineDo) Returning(value interface{}, columns ...string) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cartLineItemTaxLineDo) Not(conds ...gen.Condition) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cartLineItemTaxLineDo) Or(conds ...gen.Condition) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cartLineItemTaxLineDo) Select(conds ...field.Expr) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cartLineItemTaxLineDo) Where(conds ...gen.Condition) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cartLineItemTaxLineDo) Order(conds ...field.Expr) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cartLineItemTaxLineDo) Distinct(cols ...field.Expr) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cartLineItemTaxLineDo) Omit(cols ...field.Expr) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cartLineItemTaxLineDo) Join(table schema.Tabler, on ...field.Expr) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cartLineItemTaxLineDo) LeftJoin(table schema.Tabler, on ...field.Expr) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cartLineItemTaxLineDo) RightJoin(table schema.Tabler, on ...field.Expr) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cartLineItemTaxLineDo) Group(cols ...field.Expr) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cartLineItemTaxLineDo) Having(conds ...gen.Condition) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cartLineItemTaxLineDo) Limit(limit int) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cartLineItemTaxLineDo) Offset(offset int) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cartLineItemTaxLineDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cartLineItemTaxLineDo) Unscoped() *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cartLineItemTaxLineDo) Create(values ...*model.CartLineItemTaxLine) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cartLineItemTaxLineDo) CreateInBatches(values []*model.CartLineItemTaxLine, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cartLineItemTaxLineDo) Save(values ...*model.CartLineItemTaxLine) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cartLineItemTaxLineDo) First() (*model.CartLineItemTaxLine, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CartLineItemTaxLine), nil
	}
}

func (c cartLineItemTaxLineDo) Take() (*model.CartLineItemTaxLine, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CartLineItemTaxLine), nil
	}
}

func (c cartLineItemTaxLineDo) Last() (*model.CartLineItemTaxLine, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CartLineItemTaxLine), nil
	}
}

func (c cartLineItemTaxLineDo) Find() ([]*model.CartLineItemTaxLine, error) {
	result, err := c.DO.Find()
	return result.([]*model.CartLineItemTaxLine), err
}

func (c cartLineItemTaxLineDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CartLineItemTaxLine, err error) {
	buf := make([]*model.CartLineItemTaxLine, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cartLineItemTaxLineDo) FindInBatches(result *[]*model.CartLineItemTaxLine, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cartLineItemTaxLineDo) Attrs(attrs ...field.AssignExpr) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cartLineItemTaxLineDo) Assign(attrs ...field.AssignExpr) *cartLineItemTaxLineDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cartLineItemTaxLineDo) Joins(fields ...field.RelationField) *cartLineItemTaxLineDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cartLineItemTaxLineDo) Preload(fields ...field.RelationField) *cartLineItemTaxLineDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cartLineItemTaxLineDo) FirstOrInit() (*model.CartLineItemTaxLine, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CartLineItemTaxLine), nil
	}
}

func (c cartLineItemTaxLineDo) FirstOrCreate() (*model.CartLineItemTaxLine, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CartLineItemTaxLine), nil
	}
}

func (c cartLineItemTaxLineDo) FindByPage(offset int, limit int) (result []*model.CartLineItemTaxLine, count int64, err error) {
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

func (c cartLineItemTaxLineDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cartLineItemTaxLineDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cartLineItemTaxLineDo) Delete(models ...*model.CartLineItemTaxLine) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cartLineItemTaxLineDo) withDO(do gen.Dao) *cartLineItemTaxLineDo {
	c.DO = *do.(*gen.DO)
	return c
}
