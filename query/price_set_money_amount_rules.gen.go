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

func newPriceSetMoneyAmountRule(db *gorm.DB, opts ...gen.DOOption) priceSetMoneyAmountRule {
	_priceSetMoneyAmountRule := priceSetMoneyAmountRule{}

	_priceSetMoneyAmountRule.priceSetMoneyAmountRuleDo.UseDB(db, opts...)
	_priceSetMoneyAmountRule.priceSetMoneyAmountRuleDo.UseModel(&model.PriceSetMoneyAmountRule{})

	tableName := _priceSetMoneyAmountRule.priceSetMoneyAmountRuleDo.TableName()
	_priceSetMoneyAmountRule.ALL = field.NewAsterisk(tableName)
	_priceSetMoneyAmountRule.ID = field.NewString(tableName, "id")
	_priceSetMoneyAmountRule.PriceSetMoneyAmountID = field.NewString(tableName, "price_set_money_amount_id")
	_priceSetMoneyAmountRule.RuleTypeID = field.NewString(tableName, "rule_type_id")
	_priceSetMoneyAmountRule.Value = field.NewString(tableName, "value")

	_priceSetMoneyAmountRule.fillFieldMap()

	return _priceSetMoneyAmountRule
}

type priceSetMoneyAmountRule struct {
	priceSetMoneyAmountRuleDo priceSetMoneyAmountRuleDo

	ALL                   field.Asterisk
	ID                    field.String
	PriceSetMoneyAmountID field.String
	RuleTypeID            field.String
	Value                 field.String

	fieldMap map[string]field.Expr
}

func (p priceSetMoneyAmountRule) Table(newTableName string) *priceSetMoneyAmountRule {
	p.priceSetMoneyAmountRuleDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p priceSetMoneyAmountRule) As(alias string) *priceSetMoneyAmountRule {
	p.priceSetMoneyAmountRuleDo.DO = *(p.priceSetMoneyAmountRuleDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *priceSetMoneyAmountRule) updateTableName(table string) *priceSetMoneyAmountRule {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewString(table, "id")
	p.PriceSetMoneyAmountID = field.NewString(table, "price_set_money_amount_id")
	p.RuleTypeID = field.NewString(table, "rule_type_id")
	p.Value = field.NewString(table, "value")

	p.fillFieldMap()

	return p
}

func (p *priceSetMoneyAmountRule) WithContext(ctx context.Context) *priceSetMoneyAmountRuleDo {
	return p.priceSetMoneyAmountRuleDo.WithContext(ctx)
}

func (p priceSetMoneyAmountRule) TableName() string { return p.priceSetMoneyAmountRuleDo.TableName() }

func (p priceSetMoneyAmountRule) Alias() string { return p.priceSetMoneyAmountRuleDo.Alias() }

func (p priceSetMoneyAmountRule) Columns(cols ...field.Expr) gen.Columns {
	return p.priceSetMoneyAmountRuleDo.Columns(cols...)
}

func (p *priceSetMoneyAmountRule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *priceSetMoneyAmountRule) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["id"] = p.ID
	p.fieldMap["price_set_money_amount_id"] = p.PriceSetMoneyAmountID
	p.fieldMap["rule_type_id"] = p.RuleTypeID
	p.fieldMap["value"] = p.Value
}

func (p priceSetMoneyAmountRule) clone(db *gorm.DB) priceSetMoneyAmountRule {
	p.priceSetMoneyAmountRuleDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p priceSetMoneyAmountRule) replaceDB(db *gorm.DB) priceSetMoneyAmountRule {
	p.priceSetMoneyAmountRuleDo.ReplaceDB(db)
	return p
}

type priceSetMoneyAmountRuleDo struct{ gen.DO }

func (p priceSetMoneyAmountRuleDo) Debug() *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Debug())
}

func (p priceSetMoneyAmountRuleDo) WithContext(ctx context.Context) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p priceSetMoneyAmountRuleDo) ReadDB() *priceSetMoneyAmountRuleDo {
	return p.Clauses(dbresolver.Read)
}

func (p priceSetMoneyAmountRuleDo) WriteDB() *priceSetMoneyAmountRuleDo {
	return p.Clauses(dbresolver.Write)
}

func (p priceSetMoneyAmountRuleDo) Session(config *gorm.Session) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Session(config))
}

func (p priceSetMoneyAmountRuleDo) Clauses(conds ...clause.Expression) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p priceSetMoneyAmountRuleDo) Returning(value interface{}, columns ...string) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p priceSetMoneyAmountRuleDo) Not(conds ...gen.Condition) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p priceSetMoneyAmountRuleDo) Or(conds ...gen.Condition) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p priceSetMoneyAmountRuleDo) Select(conds ...field.Expr) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p priceSetMoneyAmountRuleDo) Where(conds ...gen.Condition) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p priceSetMoneyAmountRuleDo) Order(conds ...field.Expr) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p priceSetMoneyAmountRuleDo) Distinct(cols ...field.Expr) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p priceSetMoneyAmountRuleDo) Omit(cols ...field.Expr) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p priceSetMoneyAmountRuleDo) Join(table schema.Tabler, on ...field.Expr) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p priceSetMoneyAmountRuleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p priceSetMoneyAmountRuleDo) RightJoin(table schema.Tabler, on ...field.Expr) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p priceSetMoneyAmountRuleDo) Group(cols ...field.Expr) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p priceSetMoneyAmountRuleDo) Having(conds ...gen.Condition) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p priceSetMoneyAmountRuleDo) Limit(limit int) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p priceSetMoneyAmountRuleDo) Offset(offset int) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p priceSetMoneyAmountRuleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p priceSetMoneyAmountRuleDo) Unscoped() *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Unscoped())
}

func (p priceSetMoneyAmountRuleDo) Create(values ...*model.PriceSetMoneyAmountRule) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p priceSetMoneyAmountRuleDo) CreateInBatches(values []*model.PriceSetMoneyAmountRule, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p priceSetMoneyAmountRuleDo) Save(values ...*model.PriceSetMoneyAmountRule) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p priceSetMoneyAmountRuleDo) First() (*model.PriceSetMoneyAmountRule, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceSetMoneyAmountRule), nil
	}
}

func (p priceSetMoneyAmountRuleDo) Take() (*model.PriceSetMoneyAmountRule, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceSetMoneyAmountRule), nil
	}
}

func (p priceSetMoneyAmountRuleDo) Last() (*model.PriceSetMoneyAmountRule, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceSetMoneyAmountRule), nil
	}
}

func (p priceSetMoneyAmountRuleDo) Find() ([]*model.PriceSetMoneyAmountRule, error) {
	result, err := p.DO.Find()
	return result.([]*model.PriceSetMoneyAmountRule), err
}

func (p priceSetMoneyAmountRuleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PriceSetMoneyAmountRule, err error) {
	buf := make([]*model.PriceSetMoneyAmountRule, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p priceSetMoneyAmountRuleDo) FindInBatches(result *[]*model.PriceSetMoneyAmountRule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p priceSetMoneyAmountRuleDo) Attrs(attrs ...field.AssignExpr) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p priceSetMoneyAmountRuleDo) Assign(attrs ...field.AssignExpr) *priceSetMoneyAmountRuleDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p priceSetMoneyAmountRuleDo) Joins(fields ...field.RelationField) *priceSetMoneyAmountRuleDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p priceSetMoneyAmountRuleDo) Preload(fields ...field.RelationField) *priceSetMoneyAmountRuleDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p priceSetMoneyAmountRuleDo) FirstOrInit() (*model.PriceSetMoneyAmountRule, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceSetMoneyAmountRule), nil
	}
}

func (p priceSetMoneyAmountRuleDo) FirstOrCreate() (*model.PriceSetMoneyAmountRule, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceSetMoneyAmountRule), nil
	}
}

func (p priceSetMoneyAmountRuleDo) FindByPage(offset int, limit int) (result []*model.PriceSetMoneyAmountRule, count int64, err error) {
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

func (p priceSetMoneyAmountRuleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p priceSetMoneyAmountRuleDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p priceSetMoneyAmountRuleDo) Delete(models ...*model.PriceSetMoneyAmountRule) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *priceSetMoneyAmountRuleDo) withDO(do gen.Dao) *priceSetMoneyAmountRuleDo {
	p.DO = *do.(*gen.DO)
	return p
}
