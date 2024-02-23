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

func newPayment(db *gorm.DB, opts ...gen.DOOption) payment {
	_payment := payment{}

	_payment.paymentDo.UseDB(db, opts...)
	_payment.paymentDo.UseModel(&model.Payment{})

	tableName := _payment.paymentDo.TableName()
	_payment.ALL = field.NewAsterisk(tableName)
	_payment.ID = field.NewString(tableName, "id")
	_payment.Amount = field.NewFloat64(tableName, "amount")
	_payment.AuthorizedAmount = field.NewFloat64(tableName, "authorized_amount")
	_payment.CurrencyCode = field.NewString(tableName, "currency_code")
	_payment.ProviderID = field.NewString(tableName, "provider_id")
	_payment.CartID = field.NewString(tableName, "cart_id")
	_payment.OrderID = field.NewString(tableName, "order_id")
	_payment.OrderEditID = field.NewString(tableName, "order_edit_id")
	_payment.CustomerID = field.NewString(tableName, "customer_id")
	_payment.Data = field.NewString(tableName, "data")
	_payment.CreatedAt = field.NewTime(tableName, "created_at")
	_payment.UpdatedAt = field.NewTime(tableName, "updated_at")
	_payment.DeletedAt = field.NewField(tableName, "deleted_at")
	_payment.CapturedAt = field.NewTime(tableName, "captured_at")
	_payment.CanceledAt = field.NewTime(tableName, "canceled_at")
	_payment.PaymentCollectionID = field.NewString(tableName, "payment_collection_id")
	_payment.SessionID = field.NewString(tableName, "session_id")

	_payment.fillFieldMap()

	return _payment
}

type payment struct {
	paymentDo paymentDo

	ALL                 field.Asterisk
	ID                  field.String
	Amount              field.Float64
	AuthorizedAmount    field.Float64
	CurrencyCode        field.String
	ProviderID          field.String
	CartID              field.String
	OrderID             field.String
	OrderEditID         field.String
	CustomerID          field.String
	Data                field.String
	CreatedAt           field.Time
	UpdatedAt           field.Time
	DeletedAt           field.Field
	CapturedAt          field.Time
	CanceledAt          field.Time
	PaymentCollectionID field.String
	SessionID           field.String

	fieldMap map[string]field.Expr
}

func (p payment) Table(newTableName string) *payment {
	p.paymentDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p payment) As(alias string) *payment {
	p.paymentDo.DO = *(p.paymentDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *payment) updateTableName(table string) *payment {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewString(table, "id")
	p.Amount = field.NewFloat64(table, "amount")
	p.AuthorizedAmount = field.NewFloat64(table, "authorized_amount")
	p.CurrencyCode = field.NewString(table, "currency_code")
	p.ProviderID = field.NewString(table, "provider_id")
	p.CartID = field.NewString(table, "cart_id")
	p.OrderID = field.NewString(table, "order_id")
	p.OrderEditID = field.NewString(table, "order_edit_id")
	p.CustomerID = field.NewString(table, "customer_id")
	p.Data = field.NewString(table, "data")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.CapturedAt = field.NewTime(table, "captured_at")
	p.CanceledAt = field.NewTime(table, "canceled_at")
	p.PaymentCollectionID = field.NewString(table, "payment_collection_id")
	p.SessionID = field.NewString(table, "session_id")

	p.fillFieldMap()

	return p
}

func (p *payment) WithContext(ctx context.Context) *paymentDo { return p.paymentDo.WithContext(ctx) }

func (p payment) TableName() string { return p.paymentDo.TableName() }

func (p payment) Alias() string { return p.paymentDo.Alias() }

func (p payment) Columns(cols ...field.Expr) gen.Columns { return p.paymentDo.Columns(cols...) }

func (p *payment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *payment) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 17)
	p.fieldMap["id"] = p.ID
	p.fieldMap["amount"] = p.Amount
	p.fieldMap["authorized_amount"] = p.AuthorizedAmount
	p.fieldMap["currency_code"] = p.CurrencyCode
	p.fieldMap["provider_id"] = p.ProviderID
	p.fieldMap["cart_id"] = p.CartID
	p.fieldMap["order_id"] = p.OrderID
	p.fieldMap["order_edit_id"] = p.OrderEditID
	p.fieldMap["customer_id"] = p.CustomerID
	p.fieldMap["data"] = p.Data
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["captured_at"] = p.CapturedAt
	p.fieldMap["canceled_at"] = p.CanceledAt
	p.fieldMap["payment_collection_id"] = p.PaymentCollectionID
	p.fieldMap["session_id"] = p.SessionID
}

func (p payment) clone(db *gorm.DB) payment {
	p.paymentDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p payment) replaceDB(db *gorm.DB) payment {
	p.paymentDo.ReplaceDB(db)
	return p
}

type paymentDo struct{ gen.DO }

func (p paymentDo) Debug() *paymentDo {
	return p.withDO(p.DO.Debug())
}

func (p paymentDo) WithContext(ctx context.Context) *paymentDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p paymentDo) ReadDB() *paymentDo {
	return p.Clauses(dbresolver.Read)
}

func (p paymentDo) WriteDB() *paymentDo {
	return p.Clauses(dbresolver.Write)
}

func (p paymentDo) Session(config *gorm.Session) *paymentDo {
	return p.withDO(p.DO.Session(config))
}

func (p paymentDo) Clauses(conds ...clause.Expression) *paymentDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p paymentDo) Returning(value interface{}, columns ...string) *paymentDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p paymentDo) Not(conds ...gen.Condition) *paymentDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p paymentDo) Or(conds ...gen.Condition) *paymentDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p paymentDo) Select(conds ...field.Expr) *paymentDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p paymentDo) Where(conds ...gen.Condition) *paymentDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p paymentDo) Order(conds ...field.Expr) *paymentDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p paymentDo) Distinct(cols ...field.Expr) *paymentDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p paymentDo) Omit(cols ...field.Expr) *paymentDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p paymentDo) Join(table schema.Tabler, on ...field.Expr) *paymentDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p paymentDo) LeftJoin(table schema.Tabler, on ...field.Expr) *paymentDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p paymentDo) RightJoin(table schema.Tabler, on ...field.Expr) *paymentDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p paymentDo) Group(cols ...field.Expr) *paymentDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p paymentDo) Having(conds ...gen.Condition) *paymentDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p paymentDo) Limit(limit int) *paymentDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p paymentDo) Offset(offset int) *paymentDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p paymentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *paymentDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p paymentDo) Unscoped() *paymentDo {
	return p.withDO(p.DO.Unscoped())
}

func (p paymentDo) Create(values ...*model.Payment) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p paymentDo) CreateInBatches(values []*model.Payment, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p paymentDo) Save(values ...*model.Payment) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p paymentDo) First() (*model.Payment, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Payment), nil
	}
}

func (p paymentDo) Take() (*model.Payment, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Payment), nil
	}
}

func (p paymentDo) Last() (*model.Payment, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Payment), nil
	}
}

func (p paymentDo) Find() ([]*model.Payment, error) {
	result, err := p.DO.Find()
	return result.([]*model.Payment), err
}

func (p paymentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Payment, err error) {
	buf := make([]*model.Payment, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p paymentDo) FindInBatches(result *[]*model.Payment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p paymentDo) Attrs(attrs ...field.AssignExpr) *paymentDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p paymentDo) Assign(attrs ...field.AssignExpr) *paymentDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p paymentDo) Joins(fields ...field.RelationField) *paymentDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p paymentDo) Preload(fields ...field.RelationField) *paymentDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p paymentDo) FirstOrInit() (*model.Payment, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Payment), nil
	}
}

func (p paymentDo) FirstOrCreate() (*model.Payment, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Payment), nil
	}
}

func (p paymentDo) FindByPage(offset int, limit int) (result []*model.Payment, count int64, err error) {
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

func (p paymentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p paymentDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p paymentDo) Delete(models ...*model.Payment) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *paymentDo) withDO(do gen.Dao) *paymentDo {
	p.DO = *do.(*gen.DO)
	return p
}
