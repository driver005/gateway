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

func newOrder(db *gorm.DB, opts ...gen.DOOption) order {
	_order := order{}

	_order.orderDo.UseDB(db, opts...)
	_order.orderDo.UseModel(&model.Order{})

	tableName := _order.orderDo.TableName()
	_order.ALL = field.NewAsterisk(tableName)
	_order.ID = field.NewString(tableName, "id")
	_order.Status = field.NewString(tableName, "status")
	_order.FulfillmentStatus = field.NewString(tableName, "fulfillment_status")
	_order.PaymentStatus = field.NewString(tableName, "payment_status")
	_order.DisplayID = field.NewInt32(tableName, "display_id")
	_order.CartID = field.NewString(tableName, "cart_id")
	_order.CustomerID = field.NewString(tableName, "customer_id")
	_order.Email = field.NewString(tableName, "email")
	_order.BillingAddressID = field.NewString(tableName, "billing_address_id")
	_order.ShippingAddressID = field.NewString(tableName, "shipping_address_id")
	_order.RegionID = field.NewString(tableName, "region_id")
	_order.CurrencyCode = field.NewString(tableName, "currency_code")
	_order.TaxRate = field.NewFloat32(tableName, "tax_rate")
	_order.CanceledAt = field.NewTime(tableName, "canceled_at")
	_order.CreatedAt = field.NewTime(tableName, "created_at")
	_order.UpdatedAt = field.NewTime(tableName, "updated_at")
	_order.Metadata = field.NewString(tableName, "metadata")
	_order.IdempotencyKey = field.NewString(tableName, "idempotency_key")
	_order.DraftOrderID = field.NewString(tableName, "draft_order_id")
	_order.NoNotification = field.NewBool(tableName, "no_notification")
	_order.ExternalID = field.NewString(tableName, "external_id")
	_order.SalesChannelID = field.NewString(tableName, "sales_channel_id")

	_order.fillFieldMap()

	return _order
}

type order struct {
	orderDo orderDo

	ALL               field.Asterisk
	ID                field.String
	Status            field.String
	FulfillmentStatus field.String
	PaymentStatus     field.String
	DisplayID         field.Int32
	CartID            field.String
	CustomerID        field.String
	Email             field.String
	BillingAddressID  field.String
	ShippingAddressID field.String
	RegionID          field.String
	CurrencyCode      field.String
	TaxRate           field.Float32
	CanceledAt        field.Time
	CreatedAt         field.Time
	UpdatedAt         field.Time
	Metadata          field.String
	IdempotencyKey    field.String
	DraftOrderID      field.String
	NoNotification    field.Bool
	ExternalID        field.String
	SalesChannelID    field.String

	fieldMap map[string]field.Expr
}

func (o order) Table(newTableName string) *order {
	o.orderDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o order) As(alias string) *order {
	o.orderDo.DO = *(o.orderDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *order) updateTableName(table string) *order {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewString(table, "id")
	o.Status = field.NewString(table, "status")
	o.FulfillmentStatus = field.NewString(table, "fulfillment_status")
	o.PaymentStatus = field.NewString(table, "payment_status")
	o.DisplayID = field.NewInt32(table, "display_id")
	o.CartID = field.NewString(table, "cart_id")
	o.CustomerID = field.NewString(table, "customer_id")
	o.Email = field.NewString(table, "email")
	o.BillingAddressID = field.NewString(table, "billing_address_id")
	o.ShippingAddressID = field.NewString(table, "shipping_address_id")
	o.RegionID = field.NewString(table, "region_id")
	o.CurrencyCode = field.NewString(table, "currency_code")
	o.TaxRate = field.NewFloat32(table, "tax_rate")
	o.CanceledAt = field.NewTime(table, "canceled_at")
	o.CreatedAt = field.NewTime(table, "created_at")
	o.UpdatedAt = field.NewTime(table, "updated_at")
	o.Metadata = field.NewString(table, "metadata")
	o.IdempotencyKey = field.NewString(table, "idempotency_key")
	o.DraftOrderID = field.NewString(table, "draft_order_id")
	o.NoNotification = field.NewBool(table, "no_notification")
	o.ExternalID = field.NewString(table, "external_id")
	o.SalesChannelID = field.NewString(table, "sales_channel_id")

	o.fillFieldMap()

	return o
}

func (o *order) WithContext(ctx context.Context) *orderDo { return o.orderDo.WithContext(ctx) }

func (o order) TableName() string { return o.orderDo.TableName() }

func (o order) Alias() string { return o.orderDo.Alias() }

func (o order) Columns(cols ...field.Expr) gen.Columns { return o.orderDo.Columns(cols...) }

func (o *order) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *order) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 22)
	o.fieldMap["id"] = o.ID
	o.fieldMap["status"] = o.Status
	o.fieldMap["fulfillment_status"] = o.FulfillmentStatus
	o.fieldMap["payment_status"] = o.PaymentStatus
	o.fieldMap["display_id"] = o.DisplayID
	o.fieldMap["cart_id"] = o.CartID
	o.fieldMap["customer_id"] = o.CustomerID
	o.fieldMap["email"] = o.Email
	o.fieldMap["billing_address_id"] = o.BillingAddressID
	o.fieldMap["shipping_address_id"] = o.ShippingAddressID
	o.fieldMap["region_id"] = o.RegionID
	o.fieldMap["currency_code"] = o.CurrencyCode
	o.fieldMap["tax_rate"] = o.TaxRate
	o.fieldMap["canceled_at"] = o.CanceledAt
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
	o.fieldMap["metadata"] = o.Metadata
	o.fieldMap["idempotency_key"] = o.IdempotencyKey
	o.fieldMap["draft_order_id"] = o.DraftOrderID
	o.fieldMap["no_notification"] = o.NoNotification
	o.fieldMap["external_id"] = o.ExternalID
	o.fieldMap["sales_channel_id"] = o.SalesChannelID
}

func (o order) clone(db *gorm.DB) order {
	o.orderDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o order) replaceDB(db *gorm.DB) order {
	o.orderDo.ReplaceDB(db)
	return o
}

type orderDo struct{ gen.DO }

func (o orderDo) Debug() *orderDo {
	return o.withDO(o.DO.Debug())
}

func (o orderDo) WithContext(ctx context.Context) *orderDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o orderDo) ReadDB() *orderDo {
	return o.Clauses(dbresolver.Read)
}

func (o orderDo) WriteDB() *orderDo {
	return o.Clauses(dbresolver.Write)
}

func (o orderDo) Session(config *gorm.Session) *orderDo {
	return o.withDO(o.DO.Session(config))
}

func (o orderDo) Clauses(conds ...clause.Expression) *orderDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o orderDo) Returning(value interface{}, columns ...string) *orderDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o orderDo) Not(conds ...gen.Condition) *orderDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o orderDo) Or(conds ...gen.Condition) *orderDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o orderDo) Select(conds ...field.Expr) *orderDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o orderDo) Where(conds ...gen.Condition) *orderDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o orderDo) Order(conds ...field.Expr) *orderDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o orderDo) Distinct(cols ...field.Expr) *orderDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o orderDo) Omit(cols ...field.Expr) *orderDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o orderDo) Join(table schema.Tabler, on ...field.Expr) *orderDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o orderDo) LeftJoin(table schema.Tabler, on ...field.Expr) *orderDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o orderDo) RightJoin(table schema.Tabler, on ...field.Expr) *orderDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o orderDo) Group(cols ...field.Expr) *orderDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o orderDo) Having(conds ...gen.Condition) *orderDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o orderDo) Limit(limit int) *orderDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o orderDo) Offset(offset int) *orderDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o orderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *orderDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o orderDo) Unscoped() *orderDo {
	return o.withDO(o.DO.Unscoped())
}

func (o orderDo) Create(values ...*model.Order) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o orderDo) CreateInBatches(values []*model.Order, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o orderDo) Save(values ...*model.Order) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o orderDo) First() (*model.Order, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) Take() (*model.Order, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) Last() (*model.Order, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) Find() ([]*model.Order, error) {
	result, err := o.DO.Find()
	return result.([]*model.Order), err
}

func (o orderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Order, err error) {
	buf := make([]*model.Order, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o orderDo) FindInBatches(result *[]*model.Order, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o orderDo) Attrs(attrs ...field.AssignExpr) *orderDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o orderDo) Assign(attrs ...field.AssignExpr) *orderDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o orderDo) Joins(fields ...field.RelationField) *orderDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o orderDo) Preload(fields ...field.RelationField) *orderDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o orderDo) FirstOrInit() (*model.Order, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) FirstOrCreate() (*model.Order, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) FindByPage(offset int, limit int) (result []*model.Order, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o orderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o orderDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o orderDo) Delete(models ...*model.Order) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *orderDo) withDO(do gen.Dao) *orderDo {
	o.DO = *do.(*gen.DO)
	return o
}
