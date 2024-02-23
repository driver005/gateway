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

func newFulfillmentLabel(db *gorm.DB, opts ...gen.DOOption) fulfillmentLabel {
	_fulfillmentLabel := fulfillmentLabel{}

	_fulfillmentLabel.fulfillmentLabelDo.UseDB(db, opts...)
	_fulfillmentLabel.fulfillmentLabelDo.UseModel(&model.FulfillmentLabel{})

	tableName := _fulfillmentLabel.fulfillmentLabelDo.TableName()
	_fulfillmentLabel.ALL = field.NewAsterisk(tableName)
	_fulfillmentLabel.ID = field.NewString(tableName, "id")
	_fulfillmentLabel.TrackingNumber = field.NewString(tableName, "tracking_number")
	_fulfillmentLabel.TrackingURL = field.NewString(tableName, "tracking_url")
	_fulfillmentLabel.LabelURL = field.NewString(tableName, "label_url")
	_fulfillmentLabel.FulfillmentID = field.NewString(tableName, "fulfillment_id")
	_fulfillmentLabel.CreatedAt = field.NewTime(tableName, "created_at")
	_fulfillmentLabel.UpdatedAt = field.NewTime(tableName, "updated_at")
	_fulfillmentLabel.DeletedAt = field.NewField(tableName, "deleted_at")

	_fulfillmentLabel.fillFieldMap()

	return _fulfillmentLabel
}

type fulfillmentLabel struct {
	fulfillmentLabelDo fulfillmentLabelDo

	ALL            field.Asterisk
	ID             field.String
	TrackingNumber field.String
	TrackingURL    field.String
	LabelURL       field.String
	FulfillmentID  field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field

	fieldMap map[string]field.Expr
}

func (f fulfillmentLabel) Table(newTableName string) *fulfillmentLabel {
	f.fulfillmentLabelDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fulfillmentLabel) As(alias string) *fulfillmentLabel {
	f.fulfillmentLabelDo.DO = *(f.fulfillmentLabelDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fulfillmentLabel) updateTableName(table string) *fulfillmentLabel {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewString(table, "id")
	f.TrackingNumber = field.NewString(table, "tracking_number")
	f.TrackingURL = field.NewString(table, "tracking_url")
	f.LabelURL = field.NewString(table, "label_url")
	f.FulfillmentID = field.NewString(table, "fulfillment_id")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.DeletedAt = field.NewField(table, "deleted_at")

	f.fillFieldMap()

	return f
}

func (f *fulfillmentLabel) WithContext(ctx context.Context) *fulfillmentLabelDo {
	return f.fulfillmentLabelDo.WithContext(ctx)
}

func (f fulfillmentLabel) TableName() string { return f.fulfillmentLabelDo.TableName() }

func (f fulfillmentLabel) Alias() string { return f.fulfillmentLabelDo.Alias() }

func (f fulfillmentLabel) Columns(cols ...field.Expr) gen.Columns {
	return f.fulfillmentLabelDo.Columns(cols...)
}

func (f *fulfillmentLabel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fulfillmentLabel) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 8)
	f.fieldMap["id"] = f.ID
	f.fieldMap["tracking_number"] = f.TrackingNumber
	f.fieldMap["tracking_url"] = f.TrackingURL
	f.fieldMap["label_url"] = f.LabelURL
	f.fieldMap["fulfillment_id"] = f.FulfillmentID
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["deleted_at"] = f.DeletedAt
}

func (f fulfillmentLabel) clone(db *gorm.DB) fulfillmentLabel {
	f.fulfillmentLabelDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fulfillmentLabel) replaceDB(db *gorm.DB) fulfillmentLabel {
	f.fulfillmentLabelDo.ReplaceDB(db)
	return f
}

type fulfillmentLabelDo struct{ gen.DO }

func (f fulfillmentLabelDo) Debug() *fulfillmentLabelDo {
	return f.withDO(f.DO.Debug())
}

func (f fulfillmentLabelDo) WithContext(ctx context.Context) *fulfillmentLabelDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fulfillmentLabelDo) ReadDB() *fulfillmentLabelDo {
	return f.Clauses(dbresolver.Read)
}

func (f fulfillmentLabelDo) WriteDB() *fulfillmentLabelDo {
	return f.Clauses(dbresolver.Write)
}

func (f fulfillmentLabelDo) Session(config *gorm.Session) *fulfillmentLabelDo {
	return f.withDO(f.DO.Session(config))
}

func (f fulfillmentLabelDo) Clauses(conds ...clause.Expression) *fulfillmentLabelDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fulfillmentLabelDo) Returning(value interface{}, columns ...string) *fulfillmentLabelDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fulfillmentLabelDo) Not(conds ...gen.Condition) *fulfillmentLabelDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fulfillmentLabelDo) Or(conds ...gen.Condition) *fulfillmentLabelDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fulfillmentLabelDo) Select(conds ...field.Expr) *fulfillmentLabelDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fulfillmentLabelDo) Where(conds ...gen.Condition) *fulfillmentLabelDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fulfillmentLabelDo) Order(conds ...field.Expr) *fulfillmentLabelDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fulfillmentLabelDo) Distinct(cols ...field.Expr) *fulfillmentLabelDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fulfillmentLabelDo) Omit(cols ...field.Expr) *fulfillmentLabelDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fulfillmentLabelDo) Join(table schema.Tabler, on ...field.Expr) *fulfillmentLabelDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fulfillmentLabelDo) LeftJoin(table schema.Tabler, on ...field.Expr) *fulfillmentLabelDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fulfillmentLabelDo) RightJoin(table schema.Tabler, on ...field.Expr) *fulfillmentLabelDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fulfillmentLabelDo) Group(cols ...field.Expr) *fulfillmentLabelDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fulfillmentLabelDo) Having(conds ...gen.Condition) *fulfillmentLabelDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fulfillmentLabelDo) Limit(limit int) *fulfillmentLabelDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fulfillmentLabelDo) Offset(offset int) *fulfillmentLabelDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fulfillmentLabelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *fulfillmentLabelDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fulfillmentLabelDo) Unscoped() *fulfillmentLabelDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fulfillmentLabelDo) Create(values ...*model.FulfillmentLabel) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fulfillmentLabelDo) CreateInBatches(values []*model.FulfillmentLabel, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fulfillmentLabelDo) Save(values ...*model.FulfillmentLabel) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fulfillmentLabelDo) First() (*model.FulfillmentLabel, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentLabel), nil
	}
}

func (f fulfillmentLabelDo) Take() (*model.FulfillmentLabel, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentLabel), nil
	}
}

func (f fulfillmentLabelDo) Last() (*model.FulfillmentLabel, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentLabel), nil
	}
}

func (f fulfillmentLabelDo) Find() ([]*model.FulfillmentLabel, error) {
	result, err := f.DO.Find()
	return result.([]*model.FulfillmentLabel), err
}

func (f fulfillmentLabelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FulfillmentLabel, err error) {
	buf := make([]*model.FulfillmentLabel, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fulfillmentLabelDo) FindInBatches(result *[]*model.FulfillmentLabel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fulfillmentLabelDo) Attrs(attrs ...field.AssignExpr) *fulfillmentLabelDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fulfillmentLabelDo) Assign(attrs ...field.AssignExpr) *fulfillmentLabelDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fulfillmentLabelDo) Joins(fields ...field.RelationField) *fulfillmentLabelDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fulfillmentLabelDo) Preload(fields ...field.RelationField) *fulfillmentLabelDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fulfillmentLabelDo) FirstOrInit() (*model.FulfillmentLabel, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentLabel), nil
	}
}

func (f fulfillmentLabelDo) FirstOrCreate() (*model.FulfillmentLabel, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FulfillmentLabel), nil
	}
}

func (f fulfillmentLabelDo) FindByPage(offset int, limit int) (result []*model.FulfillmentLabel, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fulfillmentLabelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fulfillmentLabelDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fulfillmentLabelDo) Delete(models ...*model.FulfillmentLabel) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fulfillmentLabelDo) withDO(do gen.Dao) *fulfillmentLabelDo {
	f.DO = *do.(*gen.DO)
	return f
}
