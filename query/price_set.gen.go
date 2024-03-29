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

func newPriceSet(db *gorm.DB, opts ...gen.DOOption) priceSet {
	_priceSet := priceSet{}

	_priceSet.priceSetDo.UseDB(db, opts...)
	_priceSet.priceSetDo.UseModel(&model.PriceSet{})

	tableName := _priceSet.priceSetDo.TableName()
	_priceSet.ALL = field.NewAsterisk(tableName)
	_priceSet.ID = field.NewString(tableName, "id")

	_priceSet.fillFieldMap()

	return _priceSet
}

type priceSet struct {
	priceSetDo priceSetDo

	ALL field.Asterisk
	ID  field.String

	fieldMap map[string]field.Expr
}

func (p priceSet) Table(newTableName string) *priceSet {
	p.priceSetDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p priceSet) As(alias string) *priceSet {
	p.priceSetDo.DO = *(p.priceSetDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *priceSet) updateTableName(table string) *priceSet {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewString(table, "id")

	p.fillFieldMap()

	return p
}

func (p *priceSet) WithContext(ctx context.Context) *priceSetDo { return p.priceSetDo.WithContext(ctx) }

func (p priceSet) TableName() string { return p.priceSetDo.TableName() }

func (p priceSet) Alias() string { return p.priceSetDo.Alias() }

func (p priceSet) Columns(cols ...field.Expr) gen.Columns { return p.priceSetDo.Columns(cols...) }

func (p *priceSet) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *priceSet) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 1)
	p.fieldMap["id"] = p.ID
}

func (p priceSet) clone(db *gorm.DB) priceSet {
	p.priceSetDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p priceSet) replaceDB(db *gorm.DB) priceSet {
	p.priceSetDo.ReplaceDB(db)
	return p
}

type priceSetDo struct{ gen.DO }

func (p priceSetDo) Debug() *priceSetDo {
	return p.withDO(p.DO.Debug())
}

func (p priceSetDo) WithContext(ctx context.Context) *priceSetDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p priceSetDo) ReadDB() *priceSetDo {
	return p.Clauses(dbresolver.Read)
}

func (p priceSetDo) WriteDB() *priceSetDo {
	return p.Clauses(dbresolver.Write)
}

func (p priceSetDo) Session(config *gorm.Session) *priceSetDo {
	return p.withDO(p.DO.Session(config))
}

func (p priceSetDo) Clauses(conds ...clause.Expression) *priceSetDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p priceSetDo) Returning(value interface{}, columns ...string) *priceSetDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p priceSetDo) Not(conds ...gen.Condition) *priceSetDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p priceSetDo) Or(conds ...gen.Condition) *priceSetDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p priceSetDo) Select(conds ...field.Expr) *priceSetDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p priceSetDo) Where(conds ...gen.Condition) *priceSetDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p priceSetDo) Order(conds ...field.Expr) *priceSetDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p priceSetDo) Distinct(cols ...field.Expr) *priceSetDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p priceSetDo) Omit(cols ...field.Expr) *priceSetDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p priceSetDo) Join(table schema.Tabler, on ...field.Expr) *priceSetDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p priceSetDo) LeftJoin(table schema.Tabler, on ...field.Expr) *priceSetDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p priceSetDo) RightJoin(table schema.Tabler, on ...field.Expr) *priceSetDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p priceSetDo) Group(cols ...field.Expr) *priceSetDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p priceSetDo) Having(conds ...gen.Condition) *priceSetDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p priceSetDo) Limit(limit int) *priceSetDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p priceSetDo) Offset(offset int) *priceSetDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p priceSetDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *priceSetDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p priceSetDo) Unscoped() *priceSetDo {
	return p.withDO(p.DO.Unscoped())
}

func (p priceSetDo) Create(values ...*model.PriceSet) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p priceSetDo) CreateInBatches(values []*model.PriceSet, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p priceSetDo) Save(values ...*model.PriceSet) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p priceSetDo) First() (*model.PriceSet, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceSet), nil
	}
}

func (p priceSetDo) Take() (*model.PriceSet, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceSet), nil
	}
}

func (p priceSetDo) Last() (*model.PriceSet, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceSet), nil
	}
}

func (p priceSetDo) Find() ([]*model.PriceSet, error) {
	result, err := p.DO.Find()
	return result.([]*model.PriceSet), err
}

func (p priceSetDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PriceSet, err error) {
	buf := make([]*model.PriceSet, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p priceSetDo) FindInBatches(result *[]*model.PriceSet, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p priceSetDo) Attrs(attrs ...field.AssignExpr) *priceSetDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p priceSetDo) Assign(attrs ...field.AssignExpr) *priceSetDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p priceSetDo) Joins(fields ...field.RelationField) *priceSetDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p priceSetDo) Preload(fields ...field.RelationField) *priceSetDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p priceSetDo) FirstOrInit() (*model.PriceSet, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceSet), nil
	}
}

func (p priceSetDo) FirstOrCreate() (*model.PriceSet, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceSet), nil
	}
}

func (p priceSetDo) FindByPage(offset int, limit int) (result []*model.PriceSet, count int64, err error) {
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

func (p priceSetDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p priceSetDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p priceSetDo) Delete(models ...*model.PriceSet) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *priceSetDo) withDO(do gen.Dao) *priceSetDo {
	p.DO = *do.(*gen.DO)
	return p
}
