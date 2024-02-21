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

func newSwap(db *gorm.DB, opts ...gen.DOOption) swap {
	_swap := swap{}

	_swap.swapDo.UseDB(db, opts...)
	_swap.swapDo.UseModel(&model.Swap{})

	tableName := _swap.swapDo.TableName()
	_swap.ALL = field.NewAsterisk(tableName)
	_swap.ID = field.NewString(tableName, "id")
	_swap.FulfillmentStatus = field.NewString(tableName, "fulfillment_status")
	_swap.PaymentStatus = field.NewString(tableName, "payment_status")
	_swap.OrderID = field.NewString(tableName, "order_id")
	_swap.DifferenceDue = field.NewInt32(tableName, "difference_due")
	_swap.ShippingAddressID = field.NewString(tableName, "shipping_address_id")
	_swap.CartID = field.NewString(tableName, "cart_id")
	_swap.ConfirmedAt = field.NewTime(tableName, "confirmed_at")
	_swap.CreatedAt = field.NewTime(tableName, "created_at")
	_swap.UpdatedAt = field.NewTime(tableName, "updated_at")
	_swap.DeletedAt = field.NewField(tableName, "deleted_at")
	_swap.Metadata = field.NewString(tableName, "metadata")
	_swap.IdempotencyKey = field.NewString(tableName, "idempotency_key")
	_swap.NoNotification = field.NewBool(tableName, "no_notification")
	_swap.CanceledAt = field.NewTime(tableName, "canceled_at")
	_swap.AllowBackorder = field.NewBool(tableName, "allow_backorder")

	_swap.fillFieldMap()

	return _swap
}

type swap struct {
	swapDo swapDo

	ALL               field.Asterisk
	ID                field.String
	FulfillmentStatus field.String
	PaymentStatus     field.String
	OrderID           field.String
	DifferenceDue     field.Int32
	ShippingAddressID field.String
	CartID            field.String
	ConfirmedAt       field.Time
	CreatedAt         field.Time
	UpdatedAt         field.Time
	DeletedAt         field.Field
	Metadata          field.String
	IdempotencyKey    field.String
	NoNotification    field.Bool
	CanceledAt        field.Time
	AllowBackorder    field.Bool

	fieldMap map[string]field.Expr
}

func (s swap) Table(newTableName string) *swap {
	s.swapDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s swap) As(alias string) *swap {
	s.swapDo.DO = *(s.swapDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *swap) updateTableName(table string) *swap {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewString(table, "id")
	s.FulfillmentStatus = field.NewString(table, "fulfillment_status")
	s.PaymentStatus = field.NewString(table, "payment_status")
	s.OrderID = field.NewString(table, "order_id")
	s.DifferenceDue = field.NewInt32(table, "difference_due")
	s.ShippingAddressID = field.NewString(table, "shipping_address_id")
	s.CartID = field.NewString(table, "cart_id")
	s.ConfirmedAt = field.NewTime(table, "confirmed_at")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.Metadata = field.NewString(table, "metadata")
	s.IdempotencyKey = field.NewString(table, "idempotency_key")
	s.NoNotification = field.NewBool(table, "no_notification")
	s.CanceledAt = field.NewTime(table, "canceled_at")
	s.AllowBackorder = field.NewBool(table, "allow_backorder")

	s.fillFieldMap()

	return s
}

func (s *swap) WithContext(ctx context.Context) *swapDo { return s.swapDo.WithContext(ctx) }

func (s swap) TableName() string { return s.swapDo.TableName() }

func (s swap) Alias() string { return s.swapDo.Alias() }

func (s swap) Columns(cols ...field.Expr) gen.Columns { return s.swapDo.Columns(cols...) }

func (s *swap) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *swap) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 16)
	s.fieldMap["id"] = s.ID
	s.fieldMap["fulfillment_status"] = s.FulfillmentStatus
	s.fieldMap["payment_status"] = s.PaymentStatus
	s.fieldMap["order_id"] = s.OrderID
	s.fieldMap["difference_due"] = s.DifferenceDue
	s.fieldMap["shipping_address_id"] = s.ShippingAddressID
	s.fieldMap["cart_id"] = s.CartID
	s.fieldMap["confirmed_at"] = s.ConfirmedAt
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["metadata"] = s.Metadata
	s.fieldMap["idempotency_key"] = s.IdempotencyKey
	s.fieldMap["no_notification"] = s.NoNotification
	s.fieldMap["canceled_at"] = s.CanceledAt
	s.fieldMap["allow_backorder"] = s.AllowBackorder
}

func (s swap) clone(db *gorm.DB) swap {
	s.swapDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s swap) replaceDB(db *gorm.DB) swap {
	s.swapDo.ReplaceDB(db)
	return s
}

type swapDo struct{ gen.DO }

func (s swapDo) Debug() *swapDo {
	return s.withDO(s.DO.Debug())
}

func (s swapDo) WithContext(ctx context.Context) *swapDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s swapDo) ReadDB() *swapDo {
	return s.Clauses(dbresolver.Read)
}

func (s swapDo) WriteDB() *swapDo {
	return s.Clauses(dbresolver.Write)
}

func (s swapDo) Session(config *gorm.Session) *swapDo {
	return s.withDO(s.DO.Session(config))
}

func (s swapDo) Clauses(conds ...clause.Expression) *swapDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s swapDo) Returning(value interface{}, columns ...string) *swapDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s swapDo) Not(conds ...gen.Condition) *swapDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s swapDo) Or(conds ...gen.Condition) *swapDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s swapDo) Select(conds ...field.Expr) *swapDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s swapDo) Where(conds ...gen.Condition) *swapDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s swapDo) Order(conds ...field.Expr) *swapDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s swapDo) Distinct(cols ...field.Expr) *swapDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s swapDo) Omit(cols ...field.Expr) *swapDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s swapDo) Join(table schema.Tabler, on ...field.Expr) *swapDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s swapDo) LeftJoin(table schema.Tabler, on ...field.Expr) *swapDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s swapDo) RightJoin(table schema.Tabler, on ...field.Expr) *swapDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s swapDo) Group(cols ...field.Expr) *swapDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s swapDo) Having(conds ...gen.Condition) *swapDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s swapDo) Limit(limit int) *swapDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s swapDo) Offset(offset int) *swapDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s swapDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *swapDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s swapDo) Unscoped() *swapDo {
	return s.withDO(s.DO.Unscoped())
}

func (s swapDo) Create(values ...*model.Swap) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s swapDo) CreateInBatches(values []*model.Swap, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s swapDo) Save(values ...*model.Swap) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s swapDo) First() (*model.Swap, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Swap), nil
	}
}

func (s swapDo) Take() (*model.Swap, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Swap), nil
	}
}

func (s swapDo) Last() (*model.Swap, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Swap), nil
	}
}

func (s swapDo) Find() ([]*model.Swap, error) {
	result, err := s.DO.Find()
	return result.([]*model.Swap), err
}

func (s swapDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Swap, err error) {
	buf := make([]*model.Swap, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s swapDo) FindInBatches(result *[]*model.Swap, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s swapDo) Attrs(attrs ...field.AssignExpr) *swapDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s swapDo) Assign(attrs ...field.AssignExpr) *swapDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s swapDo) Joins(fields ...field.RelationField) *swapDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s swapDo) Preload(fields ...field.RelationField) *swapDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s swapDo) FirstOrInit() (*model.Swap, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Swap), nil
	}
}

func (s swapDo) FirstOrCreate() (*model.Swap, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Swap), nil
	}
}

func (s swapDo) FindByPage(offset int, limit int) (result []*model.Swap, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s swapDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s swapDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s swapDo) Delete(models ...*model.Swap) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *swapDo) withDO(do gen.Dao) *swapDo {
	s.DO = *do.(*gen.DO)
	return s
}
