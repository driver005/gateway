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

func newRegionPaymentProvider(db *gorm.DB, opts ...gen.DOOption) regionPaymentProvider {
	_regionPaymentProvider := regionPaymentProvider{}

	_regionPaymentProvider.regionPaymentProviderDo.UseDB(db, opts...)
	_regionPaymentProvider.regionPaymentProviderDo.UseModel(&model.RegionPaymentProvider{})

	tableName := _regionPaymentProvider.regionPaymentProviderDo.TableName()
	_regionPaymentProvider.ALL = field.NewAsterisk(tableName)
	_regionPaymentProvider.RegionID = field.NewString(tableName, "region_id")
	_regionPaymentProvider.ProviderID = field.NewString(tableName, "provider_id")

	_regionPaymentProvider.fillFieldMap()

	return _regionPaymentProvider
}

type regionPaymentProvider struct {
	regionPaymentProviderDo regionPaymentProviderDo

	ALL        field.Asterisk
	RegionID   field.String
	ProviderID field.String

	fieldMap map[string]field.Expr
}

func (r regionPaymentProvider) Table(newTableName string) *regionPaymentProvider {
	r.regionPaymentProviderDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r regionPaymentProvider) As(alias string) *regionPaymentProvider {
	r.regionPaymentProviderDo.DO = *(r.regionPaymentProviderDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *regionPaymentProvider) updateTableName(table string) *regionPaymentProvider {
	r.ALL = field.NewAsterisk(table)
	r.RegionID = field.NewString(table, "region_id")
	r.ProviderID = field.NewString(table, "provider_id")

	r.fillFieldMap()

	return r
}

func (r *regionPaymentProvider) WithContext(ctx context.Context) *regionPaymentProviderDo {
	return r.regionPaymentProviderDo.WithContext(ctx)
}

func (r regionPaymentProvider) TableName() string { return r.regionPaymentProviderDo.TableName() }

func (r regionPaymentProvider) Alias() string { return r.regionPaymentProviderDo.Alias() }

func (r regionPaymentProvider) Columns(cols ...field.Expr) gen.Columns {
	return r.regionPaymentProviderDo.Columns(cols...)
}

func (r *regionPaymentProvider) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *regionPaymentProvider) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 2)
	r.fieldMap["region_id"] = r.RegionID
	r.fieldMap["provider_id"] = r.ProviderID
}

func (r regionPaymentProvider) clone(db *gorm.DB) regionPaymentProvider {
	r.regionPaymentProviderDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r regionPaymentProvider) replaceDB(db *gorm.DB) regionPaymentProvider {
	r.regionPaymentProviderDo.ReplaceDB(db)
	return r
}

type regionPaymentProviderDo struct{ gen.DO }

func (r regionPaymentProviderDo) Debug() *regionPaymentProviderDo {
	return r.withDO(r.DO.Debug())
}

func (r regionPaymentProviderDo) WithContext(ctx context.Context) *regionPaymentProviderDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r regionPaymentProviderDo) ReadDB() *regionPaymentProviderDo {
	return r.Clauses(dbresolver.Read)
}

func (r regionPaymentProviderDo) WriteDB() *regionPaymentProviderDo {
	return r.Clauses(dbresolver.Write)
}

func (r regionPaymentProviderDo) Session(config *gorm.Session) *regionPaymentProviderDo {
	return r.withDO(r.DO.Session(config))
}

func (r regionPaymentProviderDo) Clauses(conds ...clause.Expression) *regionPaymentProviderDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r regionPaymentProviderDo) Returning(value interface{}, columns ...string) *regionPaymentProviderDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r regionPaymentProviderDo) Not(conds ...gen.Condition) *regionPaymentProviderDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r regionPaymentProviderDo) Or(conds ...gen.Condition) *regionPaymentProviderDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r regionPaymentProviderDo) Select(conds ...field.Expr) *regionPaymentProviderDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r regionPaymentProviderDo) Where(conds ...gen.Condition) *regionPaymentProviderDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r regionPaymentProviderDo) Order(conds ...field.Expr) *regionPaymentProviderDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r regionPaymentProviderDo) Distinct(cols ...field.Expr) *regionPaymentProviderDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r regionPaymentProviderDo) Omit(cols ...field.Expr) *regionPaymentProviderDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r regionPaymentProviderDo) Join(table schema.Tabler, on ...field.Expr) *regionPaymentProviderDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r regionPaymentProviderDo) LeftJoin(table schema.Tabler, on ...field.Expr) *regionPaymentProviderDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r regionPaymentProviderDo) RightJoin(table schema.Tabler, on ...field.Expr) *regionPaymentProviderDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r regionPaymentProviderDo) Group(cols ...field.Expr) *regionPaymentProviderDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r regionPaymentProviderDo) Having(conds ...gen.Condition) *regionPaymentProviderDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r regionPaymentProviderDo) Limit(limit int) *regionPaymentProviderDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r regionPaymentProviderDo) Offset(offset int) *regionPaymentProviderDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r regionPaymentProviderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *regionPaymentProviderDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r regionPaymentProviderDo) Unscoped() *regionPaymentProviderDo {
	return r.withDO(r.DO.Unscoped())
}

func (r regionPaymentProviderDo) Create(values ...*model.RegionPaymentProvider) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r regionPaymentProviderDo) CreateInBatches(values []*model.RegionPaymentProvider, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r regionPaymentProviderDo) Save(values ...*model.RegionPaymentProvider) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r regionPaymentProviderDo) First() (*model.RegionPaymentProvider, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegionPaymentProvider), nil
	}
}

func (r regionPaymentProviderDo) Take() (*model.RegionPaymentProvider, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegionPaymentProvider), nil
	}
}

func (r regionPaymentProviderDo) Last() (*model.RegionPaymentProvider, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegionPaymentProvider), nil
	}
}

func (r regionPaymentProviderDo) Find() ([]*model.RegionPaymentProvider, error) {
	result, err := r.DO.Find()
	return result.([]*model.RegionPaymentProvider), err
}

func (r regionPaymentProviderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RegionPaymentProvider, err error) {
	buf := make([]*model.RegionPaymentProvider, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r regionPaymentProviderDo) FindInBatches(result *[]*model.RegionPaymentProvider, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r regionPaymentProviderDo) Attrs(attrs ...field.AssignExpr) *regionPaymentProviderDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r regionPaymentProviderDo) Assign(attrs ...field.AssignExpr) *regionPaymentProviderDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r regionPaymentProviderDo) Joins(fields ...field.RelationField) *regionPaymentProviderDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r regionPaymentProviderDo) Preload(fields ...field.RelationField) *regionPaymentProviderDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r regionPaymentProviderDo) FirstOrInit() (*model.RegionPaymentProvider, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegionPaymentProvider), nil
	}
}

func (r regionPaymentProviderDo) FirstOrCreate() (*model.RegionPaymentProvider, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegionPaymentProvider), nil
	}
}

func (r regionPaymentProviderDo) FindByPage(offset int, limit int) (result []*model.RegionPaymentProvider, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r regionPaymentProviderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r regionPaymentProviderDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r regionPaymentProviderDo) Delete(models ...*model.RegionPaymentProvider) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *regionPaymentProviderDo) withDO(do gen.Dao) *regionPaymentProviderDo {
	r.DO = *do.(*gen.DO)
	return r
}
