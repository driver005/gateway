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

func newCustomerGroupCustomer(db *gorm.DB, opts ...gen.DOOption) customerGroupCustomer {
	_customerGroupCustomer := customerGroupCustomer{}

	_customerGroupCustomer.customerGroupCustomerDo.UseDB(db, opts...)
	_customerGroupCustomer.customerGroupCustomerDo.UseModel(&model.CustomerGroupCustomer{})

	tableName := _customerGroupCustomer.customerGroupCustomerDo.TableName()
	_customerGroupCustomer.ALL = field.NewAsterisk(tableName)
	_customerGroupCustomer.CustomerGroupID = field.NewString(tableName, "customer_group_id")
	_customerGroupCustomer.CustomerID = field.NewString(tableName, "customer_id")

	_customerGroupCustomer.fillFieldMap()

	return _customerGroupCustomer
}

type customerGroupCustomer struct {
	customerGroupCustomerDo customerGroupCustomerDo

	ALL             field.Asterisk
	CustomerGroupID field.String
	CustomerID      field.String

	fieldMap map[string]field.Expr
}

func (c customerGroupCustomer) Table(newTableName string) *customerGroupCustomer {
	c.customerGroupCustomerDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c customerGroupCustomer) As(alias string) *customerGroupCustomer {
	c.customerGroupCustomerDo.DO = *(c.customerGroupCustomerDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *customerGroupCustomer) updateTableName(table string) *customerGroupCustomer {
	c.ALL = field.NewAsterisk(table)
	c.CustomerGroupID = field.NewString(table, "customer_group_id")
	c.CustomerID = field.NewString(table, "customer_id")

	c.fillFieldMap()

	return c
}

func (c *customerGroupCustomer) WithContext(ctx context.Context) *customerGroupCustomerDo {
	return c.customerGroupCustomerDo.WithContext(ctx)
}

func (c customerGroupCustomer) TableName() string { return c.customerGroupCustomerDo.TableName() }

func (c customerGroupCustomer) Alias() string { return c.customerGroupCustomerDo.Alias() }

func (c customerGroupCustomer) Columns(cols ...field.Expr) gen.Columns {
	return c.customerGroupCustomerDo.Columns(cols...)
}

func (c *customerGroupCustomer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *customerGroupCustomer) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 2)
	c.fieldMap["customer_group_id"] = c.CustomerGroupID
	c.fieldMap["customer_id"] = c.CustomerID
}

func (c customerGroupCustomer) clone(db *gorm.DB) customerGroupCustomer {
	c.customerGroupCustomerDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c customerGroupCustomer) replaceDB(db *gorm.DB) customerGroupCustomer {
	c.customerGroupCustomerDo.ReplaceDB(db)
	return c
}

type customerGroupCustomerDo struct{ gen.DO }

func (c customerGroupCustomerDo) Debug() *customerGroupCustomerDo {
	return c.withDO(c.DO.Debug())
}

func (c customerGroupCustomerDo) WithContext(ctx context.Context) *customerGroupCustomerDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c customerGroupCustomerDo) ReadDB() *customerGroupCustomerDo {
	return c.Clauses(dbresolver.Read)
}

func (c customerGroupCustomerDo) WriteDB() *customerGroupCustomerDo {
	return c.Clauses(dbresolver.Write)
}

func (c customerGroupCustomerDo) Session(config *gorm.Session) *customerGroupCustomerDo {
	return c.withDO(c.DO.Session(config))
}

func (c customerGroupCustomerDo) Clauses(conds ...clause.Expression) *customerGroupCustomerDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c customerGroupCustomerDo) Returning(value interface{}, columns ...string) *customerGroupCustomerDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c customerGroupCustomerDo) Not(conds ...gen.Condition) *customerGroupCustomerDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c customerGroupCustomerDo) Or(conds ...gen.Condition) *customerGroupCustomerDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c customerGroupCustomerDo) Select(conds ...field.Expr) *customerGroupCustomerDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c customerGroupCustomerDo) Where(conds ...gen.Condition) *customerGroupCustomerDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c customerGroupCustomerDo) Order(conds ...field.Expr) *customerGroupCustomerDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c customerGroupCustomerDo) Distinct(cols ...field.Expr) *customerGroupCustomerDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c customerGroupCustomerDo) Omit(cols ...field.Expr) *customerGroupCustomerDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c customerGroupCustomerDo) Join(table schema.Tabler, on ...field.Expr) *customerGroupCustomerDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c customerGroupCustomerDo) LeftJoin(table schema.Tabler, on ...field.Expr) *customerGroupCustomerDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c customerGroupCustomerDo) RightJoin(table schema.Tabler, on ...field.Expr) *customerGroupCustomerDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c customerGroupCustomerDo) Group(cols ...field.Expr) *customerGroupCustomerDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c customerGroupCustomerDo) Having(conds ...gen.Condition) *customerGroupCustomerDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c customerGroupCustomerDo) Limit(limit int) *customerGroupCustomerDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c customerGroupCustomerDo) Offset(offset int) *customerGroupCustomerDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c customerGroupCustomerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *customerGroupCustomerDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c customerGroupCustomerDo) Unscoped() *customerGroupCustomerDo {
	return c.withDO(c.DO.Unscoped())
}

func (c customerGroupCustomerDo) Create(values ...*model.CustomerGroupCustomer) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c customerGroupCustomerDo) CreateInBatches(values []*model.CustomerGroupCustomer, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c customerGroupCustomerDo) Save(values ...*model.CustomerGroupCustomer) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c customerGroupCustomerDo) First() (*model.CustomerGroupCustomer, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CustomerGroupCustomer), nil
	}
}

func (c customerGroupCustomerDo) Take() (*model.CustomerGroupCustomer, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CustomerGroupCustomer), nil
	}
}

func (c customerGroupCustomerDo) Last() (*model.CustomerGroupCustomer, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CustomerGroupCustomer), nil
	}
}

func (c customerGroupCustomerDo) Find() ([]*model.CustomerGroupCustomer, error) {
	result, err := c.DO.Find()
	return result.([]*model.CustomerGroupCustomer), err
}

func (c customerGroupCustomerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CustomerGroupCustomer, err error) {
	buf := make([]*model.CustomerGroupCustomer, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c customerGroupCustomerDo) FindInBatches(result *[]*model.CustomerGroupCustomer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c customerGroupCustomerDo) Attrs(attrs ...field.AssignExpr) *customerGroupCustomerDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c customerGroupCustomerDo) Assign(attrs ...field.AssignExpr) *customerGroupCustomerDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c customerGroupCustomerDo) Joins(fields ...field.RelationField) *customerGroupCustomerDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c customerGroupCustomerDo) Preload(fields ...field.RelationField) *customerGroupCustomerDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c customerGroupCustomerDo) FirstOrInit() (*model.CustomerGroupCustomer, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CustomerGroupCustomer), nil
	}
}

func (c customerGroupCustomerDo) FirstOrCreate() (*model.CustomerGroupCustomer, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CustomerGroupCustomer), nil
	}
}

func (c customerGroupCustomerDo) FindByPage(offset int, limit int) (result []*model.CustomerGroupCustomer, count int64, err error) {
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

func (c customerGroupCustomerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c customerGroupCustomerDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c customerGroupCustomerDo) Delete(models ...*model.CustomerGroupCustomer) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *customerGroupCustomerDo) withDO(do gen.Dao) *customerGroupCustomerDo {
	c.DO = *do.(*gen.DO)
	return c
}
