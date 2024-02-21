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

func newGiftCardTransaction(db *gorm.DB, opts ...gen.DOOption) giftCardTransaction {
	_giftCardTransaction := giftCardTransaction{}

	_giftCardTransaction.giftCardTransactionDo.UseDB(db, opts...)
	_giftCardTransaction.giftCardTransactionDo.UseModel(&model.GiftCardTransaction{})

	tableName := _giftCardTransaction.giftCardTransactionDo.TableName()
	_giftCardTransaction.ALL = field.NewAsterisk(tableName)
	_giftCardTransaction.ID = field.NewString(tableName, "id")
	_giftCardTransaction.GiftCardID = field.NewString(tableName, "gift_card_id")
	_giftCardTransaction.OrderID = field.NewString(tableName, "order_id")
	_giftCardTransaction.Amount = field.NewInt32(tableName, "amount")
	_giftCardTransaction.CreatedAt = field.NewTime(tableName, "created_at")
	_giftCardTransaction.IsTaxable = field.NewBool(tableName, "is_taxable")
	_giftCardTransaction.TaxRate = field.NewFloat32(tableName, "tax_rate")

	_giftCardTransaction.fillFieldMap()

	return _giftCardTransaction
}

type giftCardTransaction struct {
	giftCardTransactionDo giftCardTransactionDo

	ALL        field.Asterisk
	ID         field.String
	GiftCardID field.String
	OrderID    field.String
	Amount     field.Int32
	CreatedAt  field.Time
	IsTaxable  field.Bool
	TaxRate    field.Float32

	fieldMap map[string]field.Expr
}

func (g giftCardTransaction) Table(newTableName string) *giftCardTransaction {
	g.giftCardTransactionDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g giftCardTransaction) As(alias string) *giftCardTransaction {
	g.giftCardTransactionDo.DO = *(g.giftCardTransactionDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *giftCardTransaction) updateTableName(table string) *giftCardTransaction {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewString(table, "id")
	g.GiftCardID = field.NewString(table, "gift_card_id")
	g.OrderID = field.NewString(table, "order_id")
	g.Amount = field.NewInt32(table, "amount")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.IsTaxable = field.NewBool(table, "is_taxable")
	g.TaxRate = field.NewFloat32(table, "tax_rate")

	g.fillFieldMap()

	return g
}

func (g *giftCardTransaction) WithContext(ctx context.Context) *giftCardTransactionDo {
	return g.giftCardTransactionDo.WithContext(ctx)
}

func (g giftCardTransaction) TableName() string { return g.giftCardTransactionDo.TableName() }

func (g giftCardTransaction) Alias() string { return g.giftCardTransactionDo.Alias() }

func (g giftCardTransaction) Columns(cols ...field.Expr) gen.Columns {
	return g.giftCardTransactionDo.Columns(cols...)
}

func (g *giftCardTransaction) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *giftCardTransaction) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 7)
	g.fieldMap["id"] = g.ID
	g.fieldMap["gift_card_id"] = g.GiftCardID
	g.fieldMap["order_id"] = g.OrderID
	g.fieldMap["amount"] = g.Amount
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["is_taxable"] = g.IsTaxable
	g.fieldMap["tax_rate"] = g.TaxRate
}

func (g giftCardTransaction) clone(db *gorm.DB) giftCardTransaction {
	g.giftCardTransactionDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g giftCardTransaction) replaceDB(db *gorm.DB) giftCardTransaction {
	g.giftCardTransactionDo.ReplaceDB(db)
	return g
}

type giftCardTransactionDo struct{ gen.DO }

func (g giftCardTransactionDo) Debug() *giftCardTransactionDo {
	return g.withDO(g.DO.Debug())
}

func (g giftCardTransactionDo) WithContext(ctx context.Context) *giftCardTransactionDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g giftCardTransactionDo) ReadDB() *giftCardTransactionDo {
	return g.Clauses(dbresolver.Read)
}

func (g giftCardTransactionDo) WriteDB() *giftCardTransactionDo {
	return g.Clauses(dbresolver.Write)
}

func (g giftCardTransactionDo) Session(config *gorm.Session) *giftCardTransactionDo {
	return g.withDO(g.DO.Session(config))
}

func (g giftCardTransactionDo) Clauses(conds ...clause.Expression) *giftCardTransactionDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g giftCardTransactionDo) Returning(value interface{}, columns ...string) *giftCardTransactionDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g giftCardTransactionDo) Not(conds ...gen.Condition) *giftCardTransactionDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g giftCardTransactionDo) Or(conds ...gen.Condition) *giftCardTransactionDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g giftCardTransactionDo) Select(conds ...field.Expr) *giftCardTransactionDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g giftCardTransactionDo) Where(conds ...gen.Condition) *giftCardTransactionDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g giftCardTransactionDo) Order(conds ...field.Expr) *giftCardTransactionDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g giftCardTransactionDo) Distinct(cols ...field.Expr) *giftCardTransactionDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g giftCardTransactionDo) Omit(cols ...field.Expr) *giftCardTransactionDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g giftCardTransactionDo) Join(table schema.Tabler, on ...field.Expr) *giftCardTransactionDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g giftCardTransactionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *giftCardTransactionDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g giftCardTransactionDo) RightJoin(table schema.Tabler, on ...field.Expr) *giftCardTransactionDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g giftCardTransactionDo) Group(cols ...field.Expr) *giftCardTransactionDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g giftCardTransactionDo) Having(conds ...gen.Condition) *giftCardTransactionDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g giftCardTransactionDo) Limit(limit int) *giftCardTransactionDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g giftCardTransactionDo) Offset(offset int) *giftCardTransactionDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g giftCardTransactionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *giftCardTransactionDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g giftCardTransactionDo) Unscoped() *giftCardTransactionDo {
	return g.withDO(g.DO.Unscoped())
}

func (g giftCardTransactionDo) Create(values ...*model.GiftCardTransaction) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g giftCardTransactionDo) CreateInBatches(values []*model.GiftCardTransaction, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g giftCardTransactionDo) Save(values ...*model.GiftCardTransaction) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g giftCardTransactionDo) First() (*model.GiftCardTransaction, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiftCardTransaction), nil
	}
}

func (g giftCardTransactionDo) Take() (*model.GiftCardTransaction, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiftCardTransaction), nil
	}
}

func (g giftCardTransactionDo) Last() (*model.GiftCardTransaction, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiftCardTransaction), nil
	}
}

func (g giftCardTransactionDo) Find() ([]*model.GiftCardTransaction, error) {
	result, err := g.DO.Find()
	return result.([]*model.GiftCardTransaction), err
}

func (g giftCardTransactionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GiftCardTransaction, err error) {
	buf := make([]*model.GiftCardTransaction, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g giftCardTransactionDo) FindInBatches(result *[]*model.GiftCardTransaction, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g giftCardTransactionDo) Attrs(attrs ...field.AssignExpr) *giftCardTransactionDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g giftCardTransactionDo) Assign(attrs ...field.AssignExpr) *giftCardTransactionDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g giftCardTransactionDo) Joins(fields ...field.RelationField) *giftCardTransactionDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g giftCardTransactionDo) Preload(fields ...field.RelationField) *giftCardTransactionDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g giftCardTransactionDo) FirstOrInit() (*model.GiftCardTransaction, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiftCardTransaction), nil
	}
}

func (g giftCardTransactionDo) FirstOrCreate() (*model.GiftCardTransaction, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiftCardTransaction), nil
	}
}

func (g giftCardTransactionDo) FindByPage(offset int, limit int) (result []*model.GiftCardTransaction, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g giftCardTransactionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g giftCardTransactionDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g giftCardTransactionDo) Delete(models ...*model.GiftCardTransaction) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *giftCardTransactionDo) withDO(do gen.Dao) *giftCardTransactionDo {
	g.DO = *do.(*gen.DO)
	return g
}
