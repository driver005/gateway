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

func newRegionFulfillmentProvider(db *gorm.DB, opts ...gen.DOOption) regionFulfillmentProvider {
	_regionFulfillmentProvider := regionFulfillmentProvider{}

	_regionFulfillmentProvider.regionFulfillmentProviderDo.UseDB(db, opts...)
	_regionFulfillmentProvider.regionFulfillmentProviderDo.UseModel(&model.RegionFulfillmentProvider{})

	tableName := _regionFulfillmentProvider.regionFulfillmentProviderDo.TableName()
	_regionFulfillmentProvider.ALL = field.NewAsterisk(tableName)
	_regionFulfillmentProvider.RegionID = field.NewString(tableName, "region_id")
	_regionFulfillmentProvider.ProviderID = field.NewString(tableName, "provider_id")

	_regionFulfillmentProvider.fillFieldMap()

	return _regionFulfillmentProvider
}

type regionFulfillmentProvider struct {
	regionFulfillmentProviderDo regionFulfillmentProviderDo

	ALL        field.Asterisk
	RegionID   field.String
	ProviderID field.String

	fieldMap map[string]field.Expr
}

func (r regionFulfillmentProvider) Table(newTableName string) *regionFulfillmentProvider {
	r.regionFulfillmentProviderDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r regionFulfillmentProvider) As(alias string) *regionFulfillmentProvider {
	r.regionFulfillmentProviderDo.DO = *(r.regionFulfillmentProviderDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *regionFulfillmentProvider) updateTableName(table string) *regionFulfillmentProvider {
	r.ALL = field.NewAsterisk(table)
	r.RegionID = field.NewString(table, "region_id")
	r.ProviderID = field.NewString(table, "provider_id")

	r.fillFieldMap()

	return r
}

func (r *regionFulfillmentProvider) WithContext(ctx context.Context) *regionFulfillmentProviderDo {
	return r.regionFulfillmentProviderDo.WithContext(ctx)
}

func (r regionFulfillmentProvider) TableName() string {
	return r.regionFulfillmentProviderDo.TableName()
}

func (r regionFulfillmentProvider) Alias() string { return r.regionFulfillmentProviderDo.Alias() }

func (r regionFulfillmentProvider) Columns(cols ...field.Expr) gen.Columns {
	return r.regionFulfillmentProviderDo.Columns(cols...)
}

func (r *regionFulfillmentProvider) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *regionFulfillmentProvider) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 2)
	r.fieldMap["region_id"] = r.RegionID
	r.fieldMap["provider_id"] = r.ProviderID
}

func (r regionFulfillmentProvider) clone(db *gorm.DB) regionFulfillmentProvider {
	r.regionFulfillmentProviderDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r regionFulfillmentProvider) replaceDB(db *gorm.DB) regionFulfillmentProvider {
	r.regionFulfillmentProviderDo.ReplaceDB(db)
	return r
}

type regionFulfillmentProviderDo struct{ gen.DO }

func (r regionFulfillmentProviderDo) Debug() *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Debug())
}

func (r regionFulfillmentProviderDo) WithContext(ctx context.Context) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r regionFulfillmentProviderDo) ReadDB() *regionFulfillmentProviderDo {
	return r.Clauses(dbresolver.Read)
}

func (r regionFulfillmentProviderDo) WriteDB() *regionFulfillmentProviderDo {
	return r.Clauses(dbresolver.Write)
}

func (r regionFulfillmentProviderDo) Session(config *gorm.Session) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Session(config))
}

func (r regionFulfillmentProviderDo) Clauses(conds ...clause.Expression) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r regionFulfillmentProviderDo) Returning(value interface{}, columns ...string) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r regionFulfillmentProviderDo) Not(conds ...gen.Condition) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r regionFulfillmentProviderDo) Or(conds ...gen.Condition) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r regionFulfillmentProviderDo) Select(conds ...field.Expr) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r regionFulfillmentProviderDo) Where(conds ...gen.Condition) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r regionFulfillmentProviderDo) Order(conds ...field.Expr) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r regionFulfillmentProviderDo) Distinct(cols ...field.Expr) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r regionFulfillmentProviderDo) Omit(cols ...field.Expr) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r regionFulfillmentProviderDo) Join(table schema.Tabler, on ...field.Expr) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r regionFulfillmentProviderDo) LeftJoin(table schema.Tabler, on ...field.Expr) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r regionFulfillmentProviderDo) RightJoin(table schema.Tabler, on ...field.Expr) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r regionFulfillmentProviderDo) Group(cols ...field.Expr) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r regionFulfillmentProviderDo) Having(conds ...gen.Condition) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r regionFulfillmentProviderDo) Limit(limit int) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r regionFulfillmentProviderDo) Offset(offset int) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r regionFulfillmentProviderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r regionFulfillmentProviderDo) Unscoped() *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Unscoped())
}

func (r regionFulfillmentProviderDo) Create(values ...*model.RegionFulfillmentProvider) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r regionFulfillmentProviderDo) CreateInBatches(values []*model.RegionFulfillmentProvider, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r regionFulfillmentProviderDo) Save(values ...*model.RegionFulfillmentProvider) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r regionFulfillmentProviderDo) First() (*model.RegionFulfillmentProvider, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegionFulfillmentProvider), nil
	}
}

func (r regionFulfillmentProviderDo) Take() (*model.RegionFulfillmentProvider, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegionFulfillmentProvider), nil
	}
}

func (r regionFulfillmentProviderDo) Last() (*model.RegionFulfillmentProvider, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegionFulfillmentProvider), nil
	}
}

func (r regionFulfillmentProviderDo) Find() ([]*model.RegionFulfillmentProvider, error) {
	result, err := r.DO.Find()
	return result.([]*model.RegionFulfillmentProvider), err
}

func (r regionFulfillmentProviderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RegionFulfillmentProvider, err error) {
	buf := make([]*model.RegionFulfillmentProvider, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r regionFulfillmentProviderDo) FindInBatches(result *[]*model.RegionFulfillmentProvider, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r regionFulfillmentProviderDo) Attrs(attrs ...field.AssignExpr) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r regionFulfillmentProviderDo) Assign(attrs ...field.AssignExpr) *regionFulfillmentProviderDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r regionFulfillmentProviderDo) Joins(fields ...field.RelationField) *regionFulfillmentProviderDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r regionFulfillmentProviderDo) Preload(fields ...field.RelationField) *regionFulfillmentProviderDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r regionFulfillmentProviderDo) FirstOrInit() (*model.RegionFulfillmentProvider, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegionFulfillmentProvider), nil
	}
}

func (r regionFulfillmentProviderDo) FirstOrCreate() (*model.RegionFulfillmentProvider, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegionFulfillmentProvider), nil
	}
}

func (r regionFulfillmentProviderDo) FindByPage(offset int, limit int) (result []*model.RegionFulfillmentProvider, count int64, err error) {
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

func (r regionFulfillmentProviderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r regionFulfillmentProviderDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r regionFulfillmentProviderDo) Delete(models ...*model.RegionFulfillmentProvider) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *regionFulfillmentProviderDo) withDO(do gen.Dao) *regionFulfillmentProviderDo {
	r.DO = *do.(*gen.DO)
	return r
}
