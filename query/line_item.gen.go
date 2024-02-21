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

func newLineItem(db *gorm.DB, opts ...gen.DOOption) lineItem {
	_lineItem := lineItem{}

	_lineItem.lineItemDo.UseDB(db, opts...)
	_lineItem.lineItemDo.UseModel(&model.LineItem{})

	tableName := _lineItem.lineItemDo.TableName()
	_lineItem.ALL = field.NewAsterisk(tableName)
	_lineItem.ID = field.NewString(tableName, "id")
	_lineItem.CartID = field.NewString(tableName, "cart_id")
	_lineItem.OrderID = field.NewString(tableName, "order_id")
	_lineItem.SwapID = field.NewString(tableName, "swap_id")
	_lineItem.Title = field.NewString(tableName, "title")
	_lineItem.Description = field.NewString(tableName, "description")
	_lineItem.Thumbnail = field.NewString(tableName, "thumbnail")
	_lineItem.IsGiftcard = field.NewBool(tableName, "is_giftcard")
	_lineItem.ShouldMerge = field.NewBool(tableName, "should_merge")
	_lineItem.AllowDiscounts = field.NewBool(tableName, "allow_discounts")
	_lineItem.HasShipping = field.NewBool(tableName, "has_shipping")
	_lineItem.UnitPrice = field.NewInt32(tableName, "unit_price")
	_lineItem.VariantID = field.NewString(tableName, "variant_id")
	_lineItem.Quantity = field.NewInt32(tableName, "quantity")
	_lineItem.FulfilledQuantity = field.NewInt32(tableName, "fulfilled_quantity")
	_lineItem.ReturnedQuantity = field.NewInt32(tableName, "returned_quantity")
	_lineItem.ShippedQuantity = field.NewInt32(tableName, "shipped_quantity")
	_lineItem.CreatedAt = field.NewTime(tableName, "created_at")
	_lineItem.UpdatedAt = field.NewTime(tableName, "updated_at")
	_lineItem.Metadata = field.NewString(tableName, "metadata")
	_lineItem.ClaimOrderID = field.NewString(tableName, "claim_order_id")
	_lineItem.IsReturn = field.NewBool(tableName, "is_return")
	_lineItem.IncludesTax = field.NewBool(tableName, "includes_tax")
	_lineItem.OriginalItemID = field.NewString(tableName, "original_item_id")
	_lineItem.OrderEditID = field.NewString(tableName, "order_edit_id")
	_lineItem.ProductID = field.NewString(tableName, "product_id")

	_lineItem.fillFieldMap()

	return _lineItem
}

type lineItem struct {
	lineItemDo lineItemDo

	ALL               field.Asterisk
	ID                field.String
	CartID            field.String
	OrderID           field.String
	SwapID            field.String
	Title             field.String
	Description       field.String
	Thumbnail         field.String
	IsGiftcard        field.Bool
	ShouldMerge       field.Bool
	AllowDiscounts    field.Bool
	HasShipping       field.Bool
	UnitPrice         field.Int32
	VariantID         field.String
	Quantity          field.Int32
	FulfilledQuantity field.Int32
	ReturnedQuantity  field.Int32
	ShippedQuantity   field.Int32
	CreatedAt         field.Time
	UpdatedAt         field.Time
	Metadata          field.String
	ClaimOrderID      field.String
	IsReturn          field.Bool
	IncludesTax       field.Bool
	OriginalItemID    field.String
	OrderEditID       field.String
	ProductID         field.String

	fieldMap map[string]field.Expr
}

func (l lineItem) Table(newTableName string) *lineItem {
	l.lineItemDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l lineItem) As(alias string) *lineItem {
	l.lineItemDo.DO = *(l.lineItemDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *lineItem) updateTableName(table string) *lineItem {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewString(table, "id")
	l.CartID = field.NewString(table, "cart_id")
	l.OrderID = field.NewString(table, "order_id")
	l.SwapID = field.NewString(table, "swap_id")
	l.Title = field.NewString(table, "title")
	l.Description = field.NewString(table, "description")
	l.Thumbnail = field.NewString(table, "thumbnail")
	l.IsGiftcard = field.NewBool(table, "is_giftcard")
	l.ShouldMerge = field.NewBool(table, "should_merge")
	l.AllowDiscounts = field.NewBool(table, "allow_discounts")
	l.HasShipping = field.NewBool(table, "has_shipping")
	l.UnitPrice = field.NewInt32(table, "unit_price")
	l.VariantID = field.NewString(table, "variant_id")
	l.Quantity = field.NewInt32(table, "quantity")
	l.FulfilledQuantity = field.NewInt32(table, "fulfilled_quantity")
	l.ReturnedQuantity = field.NewInt32(table, "returned_quantity")
	l.ShippedQuantity = field.NewInt32(table, "shipped_quantity")
	l.CreatedAt = field.NewTime(table, "created_at")
	l.UpdatedAt = field.NewTime(table, "updated_at")
	l.Metadata = field.NewString(table, "metadata")
	l.ClaimOrderID = field.NewString(table, "claim_order_id")
	l.IsReturn = field.NewBool(table, "is_return")
	l.IncludesTax = field.NewBool(table, "includes_tax")
	l.OriginalItemID = field.NewString(table, "original_item_id")
	l.OrderEditID = field.NewString(table, "order_edit_id")
	l.ProductID = field.NewString(table, "product_id")

	l.fillFieldMap()

	return l
}

func (l *lineItem) WithContext(ctx context.Context) *lineItemDo { return l.lineItemDo.WithContext(ctx) }

func (l lineItem) TableName() string { return l.lineItemDo.TableName() }

func (l lineItem) Alias() string { return l.lineItemDo.Alias() }

func (l lineItem) Columns(cols ...field.Expr) gen.Columns { return l.lineItemDo.Columns(cols...) }

func (l *lineItem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *lineItem) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 26)
	l.fieldMap["id"] = l.ID
	l.fieldMap["cart_id"] = l.CartID
	l.fieldMap["order_id"] = l.OrderID
	l.fieldMap["swap_id"] = l.SwapID
	l.fieldMap["title"] = l.Title
	l.fieldMap["description"] = l.Description
	l.fieldMap["thumbnail"] = l.Thumbnail
	l.fieldMap["is_giftcard"] = l.IsGiftcard
	l.fieldMap["should_merge"] = l.ShouldMerge
	l.fieldMap["allow_discounts"] = l.AllowDiscounts
	l.fieldMap["has_shipping"] = l.HasShipping
	l.fieldMap["unit_price"] = l.UnitPrice
	l.fieldMap["variant_id"] = l.VariantID
	l.fieldMap["quantity"] = l.Quantity
	l.fieldMap["fulfilled_quantity"] = l.FulfilledQuantity
	l.fieldMap["returned_quantity"] = l.ReturnedQuantity
	l.fieldMap["shipped_quantity"] = l.ShippedQuantity
	l.fieldMap["created_at"] = l.CreatedAt
	l.fieldMap["updated_at"] = l.UpdatedAt
	l.fieldMap["metadata"] = l.Metadata
	l.fieldMap["claim_order_id"] = l.ClaimOrderID
	l.fieldMap["is_return"] = l.IsReturn
	l.fieldMap["includes_tax"] = l.IncludesTax
	l.fieldMap["original_item_id"] = l.OriginalItemID
	l.fieldMap["order_edit_id"] = l.OrderEditID
	l.fieldMap["product_id"] = l.ProductID
}

func (l lineItem) clone(db *gorm.DB) lineItem {
	l.lineItemDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l lineItem) replaceDB(db *gorm.DB) lineItem {
	l.lineItemDo.ReplaceDB(db)
	return l
}

type lineItemDo struct{ gen.DO }

func (l lineItemDo) Debug() *lineItemDo {
	return l.withDO(l.DO.Debug())
}

func (l lineItemDo) WithContext(ctx context.Context) *lineItemDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l lineItemDo) ReadDB() *lineItemDo {
	return l.Clauses(dbresolver.Read)
}

func (l lineItemDo) WriteDB() *lineItemDo {
	return l.Clauses(dbresolver.Write)
}

func (l lineItemDo) Session(config *gorm.Session) *lineItemDo {
	return l.withDO(l.DO.Session(config))
}

func (l lineItemDo) Clauses(conds ...clause.Expression) *lineItemDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l lineItemDo) Returning(value interface{}, columns ...string) *lineItemDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l lineItemDo) Not(conds ...gen.Condition) *lineItemDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l lineItemDo) Or(conds ...gen.Condition) *lineItemDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l lineItemDo) Select(conds ...field.Expr) *lineItemDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l lineItemDo) Where(conds ...gen.Condition) *lineItemDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l lineItemDo) Order(conds ...field.Expr) *lineItemDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l lineItemDo) Distinct(cols ...field.Expr) *lineItemDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l lineItemDo) Omit(cols ...field.Expr) *lineItemDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l lineItemDo) Join(table schema.Tabler, on ...field.Expr) *lineItemDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l lineItemDo) LeftJoin(table schema.Tabler, on ...field.Expr) *lineItemDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l lineItemDo) RightJoin(table schema.Tabler, on ...field.Expr) *lineItemDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l lineItemDo) Group(cols ...field.Expr) *lineItemDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l lineItemDo) Having(conds ...gen.Condition) *lineItemDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l lineItemDo) Limit(limit int) *lineItemDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l lineItemDo) Offset(offset int) *lineItemDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l lineItemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *lineItemDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l lineItemDo) Unscoped() *lineItemDo {
	return l.withDO(l.DO.Unscoped())
}

func (l lineItemDo) Create(values ...*model.LineItem) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l lineItemDo) CreateInBatches(values []*model.LineItem, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l lineItemDo) Save(values ...*model.LineItem) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l lineItemDo) First() (*model.LineItem, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LineItem), nil
	}
}

func (l lineItemDo) Take() (*model.LineItem, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LineItem), nil
	}
}

func (l lineItemDo) Last() (*model.LineItem, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LineItem), nil
	}
}

func (l lineItemDo) Find() ([]*model.LineItem, error) {
	result, err := l.DO.Find()
	return result.([]*model.LineItem), err
}

func (l lineItemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LineItem, err error) {
	buf := make([]*model.LineItem, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l lineItemDo) FindInBatches(result *[]*model.LineItem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l lineItemDo) Attrs(attrs ...field.AssignExpr) *lineItemDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l lineItemDo) Assign(attrs ...field.AssignExpr) *lineItemDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l lineItemDo) Joins(fields ...field.RelationField) *lineItemDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l lineItemDo) Preload(fields ...field.RelationField) *lineItemDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l lineItemDo) FirstOrInit() (*model.LineItem, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LineItem), nil
	}
}

func (l lineItemDo) FirstOrCreate() (*model.LineItem, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LineItem), nil
	}
}

func (l lineItemDo) FindByPage(offset int, limit int) (result []*model.LineItem, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l lineItemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l lineItemDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l lineItemDo) Delete(models ...*model.LineItem) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *lineItemDo) withDO(do gen.Dao) *lineItemDo {
	l.DO = *do.(*gen.DO)
	return l
}
