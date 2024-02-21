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

func newPaymentProvider(db *gorm.DB, opts ...gen.DOOption) paymentProvider {
	_paymentProvider := paymentProvider{}

	_paymentProvider.paymentProviderDo.UseDB(db, opts...)
	_paymentProvider.paymentProviderDo.UseModel(&model.PaymentProvider{})

	tableName := _paymentProvider.paymentProviderDo.TableName()
	_paymentProvider.ALL = field.NewAsterisk(tableName)
	_paymentProvider.ID = field.NewString(tableName, "id")
	_paymentProvider.IsInstalled = field.NewBool(tableName, "is_installed")

	_paymentProvider.fillFieldMap()

	return _paymentProvider
}

type paymentProvider struct {
	paymentProviderDo paymentProviderDo

	ALL         field.Asterisk
	ID          field.String
	IsInstalled field.Bool

	fieldMap map[string]field.Expr
}

func (p paymentProvider) Table(newTableName string) *paymentProvider {
	p.paymentProviderDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p paymentProvider) As(alias string) *paymentProvider {
	p.paymentProviderDo.DO = *(p.paymentProviderDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *paymentProvider) updateTableName(table string) *paymentProvider {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewString(table, "id")
	p.IsInstalled = field.NewBool(table, "is_installed")

	p.fillFieldMap()

	return p
}

func (p *paymentProvider) WithContext(ctx context.Context) *paymentProviderDo {
	return p.paymentProviderDo.WithContext(ctx)
}

func (p paymentProvider) TableName() string { return p.paymentProviderDo.TableName() }

func (p paymentProvider) Alias() string { return p.paymentProviderDo.Alias() }

func (p paymentProvider) Columns(cols ...field.Expr) gen.Columns {
	return p.paymentProviderDo.Columns(cols...)
}

func (p *paymentProvider) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *paymentProvider) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 2)
	p.fieldMap["id"] = p.ID
	p.fieldMap["is_installed"] = p.IsInstalled
}

func (p paymentProvider) clone(db *gorm.DB) paymentProvider {
	p.paymentProviderDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p paymentProvider) replaceDB(db *gorm.DB) paymentProvider {
	p.paymentProviderDo.ReplaceDB(db)
	return p
}

type paymentProviderDo struct{ gen.DO }

func (p paymentProviderDo) Debug() *paymentProviderDo {
	return p.withDO(p.DO.Debug())
}

func (p paymentProviderDo) WithContext(ctx context.Context) *paymentProviderDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p paymentProviderDo) ReadDB() *paymentProviderDo {
	return p.Clauses(dbresolver.Read)
}

func (p paymentProviderDo) WriteDB() *paymentProviderDo {
	return p.Clauses(dbresolver.Write)
}

func (p paymentProviderDo) Session(config *gorm.Session) *paymentProviderDo {
	return p.withDO(p.DO.Session(config))
}

func (p paymentProviderDo) Clauses(conds ...clause.Expression) *paymentProviderDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p paymentProviderDo) Returning(value interface{}, columns ...string) *paymentProviderDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p paymentProviderDo) Not(conds ...gen.Condition) *paymentProviderDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p paymentProviderDo) Or(conds ...gen.Condition) *paymentProviderDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p paymentProviderDo) Select(conds ...field.Expr) *paymentProviderDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p paymentProviderDo) Where(conds ...gen.Condition) *paymentProviderDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p paymentProviderDo) Order(conds ...field.Expr) *paymentProviderDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p paymentProviderDo) Distinct(cols ...field.Expr) *paymentProviderDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p paymentProviderDo) Omit(cols ...field.Expr) *paymentProviderDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p paymentProviderDo) Join(table schema.Tabler, on ...field.Expr) *paymentProviderDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p paymentProviderDo) LeftJoin(table schema.Tabler, on ...field.Expr) *paymentProviderDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p paymentProviderDo) RightJoin(table schema.Tabler, on ...field.Expr) *paymentProviderDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p paymentProviderDo) Group(cols ...field.Expr) *paymentProviderDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p paymentProviderDo) Having(conds ...gen.Condition) *paymentProviderDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p paymentProviderDo) Limit(limit int) *paymentProviderDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p paymentProviderDo) Offset(offset int) *paymentProviderDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p paymentProviderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *paymentProviderDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p paymentProviderDo) Unscoped() *paymentProviderDo {
	return p.withDO(p.DO.Unscoped())
}

func (p paymentProviderDo) Create(values ...*model.PaymentProvider) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p paymentProviderDo) CreateInBatches(values []*model.PaymentProvider, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p paymentProviderDo) Save(values ...*model.PaymentProvider) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p paymentProviderDo) First() (*model.PaymentProvider, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentProvider), nil
	}
}

func (p paymentProviderDo) Take() (*model.PaymentProvider, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentProvider), nil
	}
}

func (p paymentProviderDo) Last() (*model.PaymentProvider, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentProvider), nil
	}
}

func (p paymentProviderDo) Find() ([]*model.PaymentProvider, error) {
	result, err := p.DO.Find()
	return result.([]*model.PaymentProvider), err
}

func (p paymentProviderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PaymentProvider, err error) {
	buf := make([]*model.PaymentProvider, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p paymentProviderDo) FindInBatches(result *[]*model.PaymentProvider, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p paymentProviderDo) Attrs(attrs ...field.AssignExpr) *paymentProviderDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p paymentProviderDo) Assign(attrs ...field.AssignExpr) *paymentProviderDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p paymentProviderDo) Joins(fields ...field.RelationField) *paymentProviderDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p paymentProviderDo) Preload(fields ...field.RelationField) *paymentProviderDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p paymentProviderDo) FirstOrInit() (*model.PaymentProvider, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentProvider), nil
	}
}

func (p paymentProviderDo) FirstOrCreate() (*model.PaymentProvider, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentProvider), nil
	}
}

func (p paymentProviderDo) FindByPage(offset int, limit int) (result []*model.PaymentProvider, count int64, err error) {
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

func (p paymentProviderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p paymentProviderDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p paymentProviderDo) Delete(models ...*model.PaymentProvider) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *paymentProviderDo) withDO(do gen.Dao) *paymentProviderDo {
	p.DO = *do.(*gen.DO)
	return p
}
