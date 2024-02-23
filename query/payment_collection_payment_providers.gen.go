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

func newPaymentCollectionPaymentProvider(db *gorm.DB, opts ...gen.DOOption) paymentCollectionPaymentProvider {
	_paymentCollectionPaymentProvider := paymentCollectionPaymentProvider{}

	_paymentCollectionPaymentProvider.paymentCollectionPaymentProviderDo.UseDB(db, opts...)
	_paymentCollectionPaymentProvider.paymentCollectionPaymentProviderDo.UseModel(&model.PaymentCollectionPaymentProvider{})

	tableName := _paymentCollectionPaymentProvider.paymentCollectionPaymentProviderDo.TableName()
	_paymentCollectionPaymentProvider.ALL = field.NewAsterisk(tableName)
	_paymentCollectionPaymentProvider.PaymentCollectionID = field.NewString(tableName, "payment_collection_id")
	_paymentCollectionPaymentProvider.PaymentProviderID = field.NewString(tableName, "payment_provider_id")

	_paymentCollectionPaymentProvider.fillFieldMap()

	return _paymentCollectionPaymentProvider
}

type paymentCollectionPaymentProvider struct {
	paymentCollectionPaymentProviderDo paymentCollectionPaymentProviderDo

	ALL                 field.Asterisk
	PaymentCollectionID field.String
	PaymentProviderID   field.String

	fieldMap map[string]field.Expr
}

func (p paymentCollectionPaymentProvider) Table(newTableName string) *paymentCollectionPaymentProvider {
	p.paymentCollectionPaymentProviderDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p paymentCollectionPaymentProvider) As(alias string) *paymentCollectionPaymentProvider {
	p.paymentCollectionPaymentProviderDo.DO = *(p.paymentCollectionPaymentProviderDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *paymentCollectionPaymentProvider) updateTableName(table string) *paymentCollectionPaymentProvider {
	p.ALL = field.NewAsterisk(table)
	p.PaymentCollectionID = field.NewString(table, "payment_collection_id")
	p.PaymentProviderID = field.NewString(table, "payment_provider_id")

	p.fillFieldMap()

	return p
}

func (p *paymentCollectionPaymentProvider) WithContext(ctx context.Context) *paymentCollectionPaymentProviderDo {
	return p.paymentCollectionPaymentProviderDo.WithContext(ctx)
}

func (p paymentCollectionPaymentProvider) TableName() string {
	return p.paymentCollectionPaymentProviderDo.TableName()
}

func (p paymentCollectionPaymentProvider) Alias() string {
	return p.paymentCollectionPaymentProviderDo.Alias()
}

func (p paymentCollectionPaymentProvider) Columns(cols ...field.Expr) gen.Columns {
	return p.paymentCollectionPaymentProviderDo.Columns(cols...)
}

func (p *paymentCollectionPaymentProvider) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *paymentCollectionPaymentProvider) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 2)
	p.fieldMap["payment_collection_id"] = p.PaymentCollectionID
	p.fieldMap["payment_provider_id"] = p.PaymentProviderID
}

func (p paymentCollectionPaymentProvider) clone(db *gorm.DB) paymentCollectionPaymentProvider {
	p.paymentCollectionPaymentProviderDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p paymentCollectionPaymentProvider) replaceDB(db *gorm.DB) paymentCollectionPaymentProvider {
	p.paymentCollectionPaymentProviderDo.ReplaceDB(db)
	return p
}

type paymentCollectionPaymentProviderDo struct{ gen.DO }

func (p paymentCollectionPaymentProviderDo) Debug() *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Debug())
}

func (p paymentCollectionPaymentProviderDo) WithContext(ctx context.Context) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p paymentCollectionPaymentProviderDo) ReadDB() *paymentCollectionPaymentProviderDo {
	return p.Clauses(dbresolver.Read)
}

func (p paymentCollectionPaymentProviderDo) WriteDB() *paymentCollectionPaymentProviderDo {
	return p.Clauses(dbresolver.Write)
}

func (p paymentCollectionPaymentProviderDo) Session(config *gorm.Session) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Session(config))
}

func (p paymentCollectionPaymentProviderDo) Clauses(conds ...clause.Expression) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p paymentCollectionPaymentProviderDo) Returning(value interface{}, columns ...string) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p paymentCollectionPaymentProviderDo) Not(conds ...gen.Condition) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p paymentCollectionPaymentProviderDo) Or(conds ...gen.Condition) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p paymentCollectionPaymentProviderDo) Select(conds ...field.Expr) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p paymentCollectionPaymentProviderDo) Where(conds ...gen.Condition) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p paymentCollectionPaymentProviderDo) Order(conds ...field.Expr) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p paymentCollectionPaymentProviderDo) Distinct(cols ...field.Expr) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p paymentCollectionPaymentProviderDo) Omit(cols ...field.Expr) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p paymentCollectionPaymentProviderDo) Join(table schema.Tabler, on ...field.Expr) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p paymentCollectionPaymentProviderDo) LeftJoin(table schema.Tabler, on ...field.Expr) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p paymentCollectionPaymentProviderDo) RightJoin(table schema.Tabler, on ...field.Expr) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p paymentCollectionPaymentProviderDo) Group(cols ...field.Expr) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p paymentCollectionPaymentProviderDo) Having(conds ...gen.Condition) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p paymentCollectionPaymentProviderDo) Limit(limit int) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p paymentCollectionPaymentProviderDo) Offset(offset int) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p paymentCollectionPaymentProviderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p paymentCollectionPaymentProviderDo) Unscoped() *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Unscoped())
}

func (p paymentCollectionPaymentProviderDo) Create(values ...*model.PaymentCollectionPaymentProvider) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p paymentCollectionPaymentProviderDo) CreateInBatches(values []*model.PaymentCollectionPaymentProvider, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p paymentCollectionPaymentProviderDo) Save(values ...*model.PaymentCollectionPaymentProvider) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p paymentCollectionPaymentProviderDo) First() (*model.PaymentCollectionPaymentProvider, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentCollectionPaymentProvider), nil
	}
}

func (p paymentCollectionPaymentProviderDo) Take() (*model.PaymentCollectionPaymentProvider, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentCollectionPaymentProvider), nil
	}
}

func (p paymentCollectionPaymentProviderDo) Last() (*model.PaymentCollectionPaymentProvider, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentCollectionPaymentProvider), nil
	}
}

func (p paymentCollectionPaymentProviderDo) Find() ([]*model.PaymentCollectionPaymentProvider, error) {
	result, err := p.DO.Find()
	return result.([]*model.PaymentCollectionPaymentProvider), err
}

func (p paymentCollectionPaymentProviderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PaymentCollectionPaymentProvider, err error) {
	buf := make([]*model.PaymentCollectionPaymentProvider, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p paymentCollectionPaymentProviderDo) FindInBatches(result *[]*model.PaymentCollectionPaymentProvider, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p paymentCollectionPaymentProviderDo) Attrs(attrs ...field.AssignExpr) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p paymentCollectionPaymentProviderDo) Assign(attrs ...field.AssignExpr) *paymentCollectionPaymentProviderDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p paymentCollectionPaymentProviderDo) Joins(fields ...field.RelationField) *paymentCollectionPaymentProviderDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p paymentCollectionPaymentProviderDo) Preload(fields ...field.RelationField) *paymentCollectionPaymentProviderDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p paymentCollectionPaymentProviderDo) FirstOrInit() (*model.PaymentCollectionPaymentProvider, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentCollectionPaymentProvider), nil
	}
}

func (p paymentCollectionPaymentProviderDo) FirstOrCreate() (*model.PaymentCollectionPaymentProvider, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentCollectionPaymentProvider), nil
	}
}

func (p paymentCollectionPaymentProviderDo) FindByPage(offset int, limit int) (result []*model.PaymentCollectionPaymentProvider, count int64, err error) {
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

func (p paymentCollectionPaymentProviderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p paymentCollectionPaymentProviderDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p paymentCollectionPaymentProviderDo) Delete(models ...*model.PaymentCollectionPaymentProvider) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *paymentCollectionPaymentProviderDo) withDO(do gen.Dao) *paymentCollectionPaymentProviderDo {
	p.DO = *do.(*gen.DO)
	return p
}
