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

func newPriceListRuleValue(db *gorm.DB, opts ...gen.DOOption) priceListRuleValue {
	_priceListRuleValue := priceListRuleValue{}

	_priceListRuleValue.priceListRuleValueDo.UseDB(db, opts...)
	_priceListRuleValue.priceListRuleValueDo.UseModel(&model.PriceListRuleValue{})

	tableName := _priceListRuleValue.priceListRuleValueDo.TableName()
	_priceListRuleValue.ALL = field.NewAsterisk(tableName)
	_priceListRuleValue.ID = field.NewString(tableName, "id")
	_priceListRuleValue.Value = field.NewString(tableName, "value")
	_priceListRuleValue.PriceListRuleID = field.NewString(tableName, "price_list_rule_id")

	_priceListRuleValue.fillFieldMap()

	return _priceListRuleValue
}

type priceListRuleValue struct {
	priceListRuleValueDo priceListRuleValueDo

	ALL             field.Asterisk
	ID              field.String
	Value           field.String
	PriceListRuleID field.String

	fieldMap map[string]field.Expr
}

func (p priceListRuleValue) Table(newTableName string) *priceListRuleValue {
	p.priceListRuleValueDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p priceListRuleValue) As(alias string) *priceListRuleValue {
	p.priceListRuleValueDo.DO = *(p.priceListRuleValueDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *priceListRuleValue) updateTableName(table string) *priceListRuleValue {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewString(table, "id")
	p.Value = field.NewString(table, "value")
	p.PriceListRuleID = field.NewString(table, "price_list_rule_id")

	p.fillFieldMap()

	return p
}

func (p *priceListRuleValue) WithContext(ctx context.Context) *priceListRuleValueDo {
	return p.priceListRuleValueDo.WithContext(ctx)
}

func (p priceListRuleValue) TableName() string { return p.priceListRuleValueDo.TableName() }

func (p priceListRuleValue) Alias() string { return p.priceListRuleValueDo.Alias() }

func (p priceListRuleValue) Columns(cols ...field.Expr) gen.Columns {
	return p.priceListRuleValueDo.Columns(cols...)
}

func (p *priceListRuleValue) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *priceListRuleValue) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 3)
	p.fieldMap["id"] = p.ID
	p.fieldMap["value"] = p.Value
	p.fieldMap["price_list_rule_id"] = p.PriceListRuleID
}

func (p priceListRuleValue) clone(db *gorm.DB) priceListRuleValue {
	p.priceListRuleValueDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p priceListRuleValue) replaceDB(db *gorm.DB) priceListRuleValue {
	p.priceListRuleValueDo.ReplaceDB(db)
	return p
}

type priceListRuleValueDo struct{ gen.DO }

func (p priceListRuleValueDo) Debug() *priceListRuleValueDo {
	return p.withDO(p.DO.Debug())
}

func (p priceListRuleValueDo) WithContext(ctx context.Context) *priceListRuleValueDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p priceListRuleValueDo) ReadDB() *priceListRuleValueDo {
	return p.Clauses(dbresolver.Read)
}

func (p priceListRuleValueDo) WriteDB() *priceListRuleValueDo {
	return p.Clauses(dbresolver.Write)
}

func (p priceListRuleValueDo) Session(config *gorm.Session) *priceListRuleValueDo {
	return p.withDO(p.DO.Session(config))
}

func (p priceListRuleValueDo) Clauses(conds ...clause.Expression) *priceListRuleValueDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p priceListRuleValueDo) Returning(value interface{}, columns ...string) *priceListRuleValueDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p priceListRuleValueDo) Not(conds ...gen.Condition) *priceListRuleValueDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p priceListRuleValueDo) Or(conds ...gen.Condition) *priceListRuleValueDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p priceListRuleValueDo) Select(conds ...field.Expr) *priceListRuleValueDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p priceListRuleValueDo) Where(conds ...gen.Condition) *priceListRuleValueDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p priceListRuleValueDo) Order(conds ...field.Expr) *priceListRuleValueDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p priceListRuleValueDo) Distinct(cols ...field.Expr) *priceListRuleValueDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p priceListRuleValueDo) Omit(cols ...field.Expr) *priceListRuleValueDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p priceListRuleValueDo) Join(table schema.Tabler, on ...field.Expr) *priceListRuleValueDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p priceListRuleValueDo) LeftJoin(table schema.Tabler, on ...field.Expr) *priceListRuleValueDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p priceListRuleValueDo) RightJoin(table schema.Tabler, on ...field.Expr) *priceListRuleValueDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p priceListRuleValueDo) Group(cols ...field.Expr) *priceListRuleValueDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p priceListRuleValueDo) Having(conds ...gen.Condition) *priceListRuleValueDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p priceListRuleValueDo) Limit(limit int) *priceListRuleValueDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p priceListRuleValueDo) Offset(offset int) *priceListRuleValueDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p priceListRuleValueDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *priceListRuleValueDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p priceListRuleValueDo) Unscoped() *priceListRuleValueDo {
	return p.withDO(p.DO.Unscoped())
}

func (p priceListRuleValueDo) Create(values ...*model.PriceListRuleValue) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p priceListRuleValueDo) CreateInBatches(values []*model.PriceListRuleValue, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p priceListRuleValueDo) Save(values ...*model.PriceListRuleValue) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p priceListRuleValueDo) First() (*model.PriceListRuleValue, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceListRuleValue), nil
	}
}

func (p priceListRuleValueDo) Take() (*model.PriceListRuleValue, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceListRuleValue), nil
	}
}

func (p priceListRuleValueDo) Last() (*model.PriceListRuleValue, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceListRuleValue), nil
	}
}

func (p priceListRuleValueDo) Find() ([]*model.PriceListRuleValue, error) {
	result, err := p.DO.Find()
	return result.([]*model.PriceListRuleValue), err
}

func (p priceListRuleValueDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PriceListRuleValue, err error) {
	buf := make([]*model.PriceListRuleValue, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p priceListRuleValueDo) FindInBatches(result *[]*model.PriceListRuleValue, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p priceListRuleValueDo) Attrs(attrs ...field.AssignExpr) *priceListRuleValueDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p priceListRuleValueDo) Assign(attrs ...field.AssignExpr) *priceListRuleValueDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p priceListRuleValueDo) Joins(fields ...field.RelationField) *priceListRuleValueDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p priceListRuleValueDo) Preload(fields ...field.RelationField) *priceListRuleValueDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p priceListRuleValueDo) FirstOrInit() (*model.PriceListRuleValue, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceListRuleValue), nil
	}
}

func (p priceListRuleValueDo) FirstOrCreate() (*model.PriceListRuleValue, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceListRuleValue), nil
	}
}

func (p priceListRuleValueDo) FindByPage(offset int, limit int) (result []*model.PriceListRuleValue, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p priceListRuleValueDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p priceListRuleValueDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p priceListRuleValueDo) Delete(models ...*model.PriceListRuleValue) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *priceListRuleValueDo) withDO(do gen.Dao) *priceListRuleValueDo {
	p.DO = *do.(*gen.DO)
	return p
}
