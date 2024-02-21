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

func newPaymentCollection(db *gorm.DB, opts ...gen.DOOption) paymentCollection {
	_paymentCollection := paymentCollection{}

	_paymentCollection.paymentCollectionDo.UseDB(db, opts...)
	_paymentCollection.paymentCollectionDo.UseModel(&model.PaymentCollection{})

	tableName := _paymentCollection.paymentCollectionDo.TableName()
	_paymentCollection.ALL = field.NewAsterisk(tableName)
	_paymentCollection.ID = field.NewString(tableName, "id")
	_paymentCollection.CreatedAt = field.NewTime(tableName, "created_at")
	_paymentCollection.UpdatedAt = field.NewTime(tableName, "updated_at")
	_paymentCollection.DeletedAt = field.NewField(tableName, "deleted_at")
	_paymentCollection.Type = field.NewString(tableName, "type")
	_paymentCollection.Status = field.NewString(tableName, "status")
	_paymentCollection.Description = field.NewString(tableName, "description")
	_paymentCollection.Amount = field.NewInt32(tableName, "amount")
	_paymentCollection.AuthorizedAmount = field.NewInt32(tableName, "authorized_amount")
	_paymentCollection.RegionID = field.NewString(tableName, "region_id")
	_paymentCollection.CurrencyCode = field.NewString(tableName, "currency_code")
	_paymentCollection.Metadata = field.NewString(tableName, "metadata")
	_paymentCollection.CreatedBy = field.NewString(tableName, "created_by")

	_paymentCollection.fillFieldMap()

	return _paymentCollection
}

type paymentCollection struct {
	paymentCollectionDo paymentCollectionDo

	ALL              field.Asterisk
	ID               field.String
	CreatedAt        field.Time
	UpdatedAt        field.Time
	DeletedAt        field.Field
	Type             field.String
	Status           field.String
	Description      field.String
	Amount           field.Int32
	AuthorizedAmount field.Int32
	RegionID         field.String
	CurrencyCode     field.String
	Metadata         field.String
	CreatedBy        field.String

	fieldMap map[string]field.Expr
}

func (p paymentCollection) Table(newTableName string) *paymentCollection {
	p.paymentCollectionDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p paymentCollection) As(alias string) *paymentCollection {
	p.paymentCollectionDo.DO = *(p.paymentCollectionDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *paymentCollection) updateTableName(table string) *paymentCollection {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewString(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Type = field.NewString(table, "type")
	p.Status = field.NewString(table, "status")
	p.Description = field.NewString(table, "description")
	p.Amount = field.NewInt32(table, "amount")
	p.AuthorizedAmount = field.NewInt32(table, "authorized_amount")
	p.RegionID = field.NewString(table, "region_id")
	p.CurrencyCode = field.NewString(table, "currency_code")
	p.Metadata = field.NewString(table, "metadata")
	p.CreatedBy = field.NewString(table, "created_by")

	p.fillFieldMap()

	return p
}

func (p *paymentCollection) WithContext(ctx context.Context) *paymentCollectionDo {
	return p.paymentCollectionDo.WithContext(ctx)
}

func (p paymentCollection) TableName() string { return p.paymentCollectionDo.TableName() }

func (p paymentCollection) Alias() string { return p.paymentCollectionDo.Alias() }

func (p paymentCollection) Columns(cols ...field.Expr) gen.Columns {
	return p.paymentCollectionDo.Columns(cols...)
}

func (p *paymentCollection) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *paymentCollection) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 13)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["type"] = p.Type
	p.fieldMap["status"] = p.Status
	p.fieldMap["description"] = p.Description
	p.fieldMap["amount"] = p.Amount
	p.fieldMap["authorized_amount"] = p.AuthorizedAmount
	p.fieldMap["region_id"] = p.RegionID
	p.fieldMap["currency_code"] = p.CurrencyCode
	p.fieldMap["metadata"] = p.Metadata
	p.fieldMap["created_by"] = p.CreatedBy
}

func (p paymentCollection) clone(db *gorm.DB) paymentCollection {
	p.paymentCollectionDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p paymentCollection) replaceDB(db *gorm.DB) paymentCollection {
	p.paymentCollectionDo.ReplaceDB(db)
	return p
}

type paymentCollectionDo struct{ gen.DO }

func (p paymentCollectionDo) Debug() *paymentCollectionDo {
	return p.withDO(p.DO.Debug())
}

func (p paymentCollectionDo) WithContext(ctx context.Context) *paymentCollectionDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p paymentCollectionDo) ReadDB() *paymentCollectionDo {
	return p.Clauses(dbresolver.Read)
}

func (p paymentCollectionDo) WriteDB() *paymentCollectionDo {
	return p.Clauses(dbresolver.Write)
}

func (p paymentCollectionDo) Session(config *gorm.Session) *paymentCollectionDo {
	return p.withDO(p.DO.Session(config))
}

func (p paymentCollectionDo) Clauses(conds ...clause.Expression) *paymentCollectionDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p paymentCollectionDo) Returning(value interface{}, columns ...string) *paymentCollectionDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p paymentCollectionDo) Not(conds ...gen.Condition) *paymentCollectionDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p paymentCollectionDo) Or(conds ...gen.Condition) *paymentCollectionDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p paymentCollectionDo) Select(conds ...field.Expr) *paymentCollectionDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p paymentCollectionDo) Where(conds ...gen.Condition) *paymentCollectionDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p paymentCollectionDo) Order(conds ...field.Expr) *paymentCollectionDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p paymentCollectionDo) Distinct(cols ...field.Expr) *paymentCollectionDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p paymentCollectionDo) Omit(cols ...field.Expr) *paymentCollectionDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p paymentCollectionDo) Join(table schema.Tabler, on ...field.Expr) *paymentCollectionDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p paymentCollectionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *paymentCollectionDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p paymentCollectionDo) RightJoin(table schema.Tabler, on ...field.Expr) *paymentCollectionDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p paymentCollectionDo) Group(cols ...field.Expr) *paymentCollectionDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p paymentCollectionDo) Having(conds ...gen.Condition) *paymentCollectionDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p paymentCollectionDo) Limit(limit int) *paymentCollectionDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p paymentCollectionDo) Offset(offset int) *paymentCollectionDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p paymentCollectionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *paymentCollectionDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p paymentCollectionDo) Unscoped() *paymentCollectionDo {
	return p.withDO(p.DO.Unscoped())
}

func (p paymentCollectionDo) Create(values ...*model.PaymentCollection) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p paymentCollectionDo) CreateInBatches(values []*model.PaymentCollection, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p paymentCollectionDo) Save(values ...*model.PaymentCollection) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p paymentCollectionDo) First() (*model.PaymentCollection, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentCollection), nil
	}
}

func (p paymentCollectionDo) Take() (*model.PaymentCollection, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentCollection), nil
	}
}

func (p paymentCollectionDo) Last() (*model.PaymentCollection, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentCollection), nil
	}
}

func (p paymentCollectionDo) Find() ([]*model.PaymentCollection, error) {
	result, err := p.DO.Find()
	return result.([]*model.PaymentCollection), err
}

func (p paymentCollectionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PaymentCollection, err error) {
	buf := make([]*model.PaymentCollection, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p paymentCollectionDo) FindInBatches(result *[]*model.PaymentCollection, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p paymentCollectionDo) Attrs(attrs ...field.AssignExpr) *paymentCollectionDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p paymentCollectionDo) Assign(attrs ...field.AssignExpr) *paymentCollectionDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p paymentCollectionDo) Joins(fields ...field.RelationField) *paymentCollectionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p paymentCollectionDo) Preload(fields ...field.RelationField) *paymentCollectionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p paymentCollectionDo) FirstOrInit() (*model.PaymentCollection, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentCollection), nil
	}
}

func (p paymentCollectionDo) FirstOrCreate() (*model.PaymentCollection, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentCollection), nil
	}
}

func (p paymentCollectionDo) FindByPage(offset int, limit int) (result []*model.PaymentCollection, count int64, err error) {
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

func (p paymentCollectionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p paymentCollectionDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p paymentCollectionDo) Delete(models ...*model.PaymentCollection) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *paymentCollectionDo) withDO(do gen.Dao) *paymentCollectionDo {
	p.DO = *do.(*gen.DO)
	return p
}
