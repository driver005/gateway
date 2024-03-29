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

func newPriceList(db *gorm.DB, opts ...gen.DOOption) priceList {
	_priceList := priceList{}

	_priceList.priceListDo.UseDB(db, opts...)
	_priceList.priceListDo.UseModel(&model.PriceList{})

	tableName := _priceList.priceListDo.TableName()
	_priceList.ALL = field.NewAsterisk(tableName)
	_priceList.ID = field.NewString(tableName, "id")
	_priceList.Status = field.NewString(tableName, "status")
	_priceList.StartsAt = field.NewTime(tableName, "starts_at")
	_priceList.EndsAt = field.NewTime(tableName, "ends_at")
	_priceList.RulesCount = field.NewInt32(tableName, "rules_count")
	_priceList.Title = field.NewString(tableName, "title")
	_priceList.Name = field.NewString(tableName, "name")
	_priceList.Description = field.NewString(tableName, "description")
	_priceList.Type = field.NewString(tableName, "type")
	_priceList.CreatedAt = field.NewTime(tableName, "created_at")
	_priceList.UpdatedAt = field.NewTime(tableName, "updated_at")
	_priceList.DeletedAt = field.NewField(tableName, "deleted_at")

	_priceList.fillFieldMap()

	return _priceList
}

type priceList struct {
	priceListDo priceListDo

	ALL         field.Asterisk
	ID          field.String
	Status      field.String
	StartsAt    field.Time
	EndsAt      field.Time
	RulesCount  field.Int32
	Title       field.String
	Name        field.String
	Description field.String
	Type        field.String
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field

	fieldMap map[string]field.Expr
}

func (p priceList) Table(newTableName string) *priceList {
	p.priceListDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p priceList) As(alias string) *priceList {
	p.priceListDo.DO = *(p.priceListDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *priceList) updateTableName(table string) *priceList {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewString(table, "id")
	p.Status = field.NewString(table, "status")
	p.StartsAt = field.NewTime(table, "starts_at")
	p.EndsAt = field.NewTime(table, "ends_at")
	p.RulesCount = field.NewInt32(table, "rules_count")
	p.Title = field.NewString(table, "title")
	p.Name = field.NewString(table, "name")
	p.Description = field.NewString(table, "description")
	p.Type = field.NewString(table, "type")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *priceList) WithContext(ctx context.Context) *priceListDo {
	return p.priceListDo.WithContext(ctx)
}

func (p priceList) TableName() string { return p.priceListDo.TableName() }

func (p priceList) Alias() string { return p.priceListDo.Alias() }

func (p priceList) Columns(cols ...field.Expr) gen.Columns { return p.priceListDo.Columns(cols...) }

func (p *priceList) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *priceList) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["id"] = p.ID
	p.fieldMap["status"] = p.Status
	p.fieldMap["starts_at"] = p.StartsAt
	p.fieldMap["ends_at"] = p.EndsAt
	p.fieldMap["rules_count"] = p.RulesCount
	p.fieldMap["title"] = p.Title
	p.fieldMap["name"] = p.Name
	p.fieldMap["description"] = p.Description
	p.fieldMap["type"] = p.Type
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p priceList) clone(db *gorm.DB) priceList {
	p.priceListDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p priceList) replaceDB(db *gorm.DB) priceList {
	p.priceListDo.ReplaceDB(db)
	return p
}

type priceListDo struct{ gen.DO }

func (p priceListDo) Debug() *priceListDo {
	return p.withDO(p.DO.Debug())
}

func (p priceListDo) WithContext(ctx context.Context) *priceListDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p priceListDo) ReadDB() *priceListDo {
	return p.Clauses(dbresolver.Read)
}

func (p priceListDo) WriteDB() *priceListDo {
	return p.Clauses(dbresolver.Write)
}

func (p priceListDo) Session(config *gorm.Session) *priceListDo {
	return p.withDO(p.DO.Session(config))
}

func (p priceListDo) Clauses(conds ...clause.Expression) *priceListDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p priceListDo) Returning(value interface{}, columns ...string) *priceListDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p priceListDo) Not(conds ...gen.Condition) *priceListDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p priceListDo) Or(conds ...gen.Condition) *priceListDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p priceListDo) Select(conds ...field.Expr) *priceListDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p priceListDo) Where(conds ...gen.Condition) *priceListDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p priceListDo) Order(conds ...field.Expr) *priceListDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p priceListDo) Distinct(cols ...field.Expr) *priceListDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p priceListDo) Omit(cols ...field.Expr) *priceListDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p priceListDo) Join(table schema.Tabler, on ...field.Expr) *priceListDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p priceListDo) LeftJoin(table schema.Tabler, on ...field.Expr) *priceListDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p priceListDo) RightJoin(table schema.Tabler, on ...field.Expr) *priceListDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p priceListDo) Group(cols ...field.Expr) *priceListDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p priceListDo) Having(conds ...gen.Condition) *priceListDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p priceListDo) Limit(limit int) *priceListDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p priceListDo) Offset(offset int) *priceListDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p priceListDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *priceListDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p priceListDo) Unscoped() *priceListDo {
	return p.withDO(p.DO.Unscoped())
}

func (p priceListDo) Create(values ...*model.PriceList) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p priceListDo) CreateInBatches(values []*model.PriceList, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p priceListDo) Save(values ...*model.PriceList) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p priceListDo) First() (*model.PriceList, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceList), nil
	}
}

func (p priceListDo) Take() (*model.PriceList, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceList), nil
	}
}

func (p priceListDo) Last() (*model.PriceList, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceList), nil
	}
}

func (p priceListDo) Find() ([]*model.PriceList, error) {
	result, err := p.DO.Find()
	return result.([]*model.PriceList), err
}

func (p priceListDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PriceList, err error) {
	buf := make([]*model.PriceList, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p priceListDo) FindInBatches(result *[]*model.PriceList, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p priceListDo) Attrs(attrs ...field.AssignExpr) *priceListDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p priceListDo) Assign(attrs ...field.AssignExpr) *priceListDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p priceListDo) Joins(fields ...field.RelationField) *priceListDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p priceListDo) Preload(fields ...field.RelationField) *priceListDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p priceListDo) FirstOrInit() (*model.PriceList, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceList), nil
	}
}

func (p priceListDo) FirstOrCreate() (*model.PriceList, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PriceList), nil
	}
}

func (p priceListDo) FindByPage(offset int, limit int) (result []*model.PriceList, count int64, err error) {
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

func (p priceListDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p priceListDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p priceListDo) Delete(models ...*model.PriceList) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *priceListDo) withDO(do gen.Dao) *priceListDo {
	p.DO = *do.(*gen.DO)
	return p
}
