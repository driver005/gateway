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

func newNotification(db *gorm.DB, opts ...gen.DOOption) notification {
	_notification := notification{}

	_notification.notificationDo.UseDB(db, opts...)
	_notification.notificationDo.UseModel(&model.Notification{})

	tableName := _notification.notificationDo.TableName()
	_notification.ALL = field.NewAsterisk(tableName)
	_notification.ID = field.NewString(tableName, "id")
	_notification.EventName = field.NewString(tableName, "event_name")
	_notification.ResourceType = field.NewString(tableName, "resource_type")
	_notification.ResourceID = field.NewString(tableName, "resource_id")
	_notification.CustomerID = field.NewString(tableName, "customer_id")
	_notification.To = field.NewString(tableName, "to")
	_notification.Data = field.NewString(tableName, "data")
	_notification.ParentID = field.NewString(tableName, "parent_id")
	_notification.ProviderID = field.NewString(tableName, "provider_id")
	_notification.CreatedAt = field.NewTime(tableName, "created_at")
	_notification.UpdatedAt = field.NewTime(tableName, "updated_at")

	_notification.fillFieldMap()

	return _notification
}

type notification struct {
	notificationDo notificationDo

	ALL          field.Asterisk
	ID           field.String
	EventName    field.String
	ResourceType field.String
	ResourceID   field.String
	CustomerID   field.String
	To           field.String
	Data         field.String
	ParentID     field.String
	ProviderID   field.String
	CreatedAt    field.Time
	UpdatedAt    field.Time

	fieldMap map[string]field.Expr
}

func (n notification) Table(newTableName string) *notification {
	n.notificationDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n notification) As(alias string) *notification {
	n.notificationDo.DO = *(n.notificationDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *notification) updateTableName(table string) *notification {
	n.ALL = field.NewAsterisk(table)
	n.ID = field.NewString(table, "id")
	n.EventName = field.NewString(table, "event_name")
	n.ResourceType = field.NewString(table, "resource_type")
	n.ResourceID = field.NewString(table, "resource_id")
	n.CustomerID = field.NewString(table, "customer_id")
	n.To = field.NewString(table, "to")
	n.Data = field.NewString(table, "data")
	n.ParentID = field.NewString(table, "parent_id")
	n.ProviderID = field.NewString(table, "provider_id")
	n.CreatedAt = field.NewTime(table, "created_at")
	n.UpdatedAt = field.NewTime(table, "updated_at")

	n.fillFieldMap()

	return n
}

func (n *notification) WithContext(ctx context.Context) *notificationDo {
	return n.notificationDo.WithContext(ctx)
}

func (n notification) TableName() string { return n.notificationDo.TableName() }

func (n notification) Alias() string { return n.notificationDo.Alias() }

func (n notification) Columns(cols ...field.Expr) gen.Columns {
	return n.notificationDo.Columns(cols...)
}

func (n *notification) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *notification) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 11)
	n.fieldMap["id"] = n.ID
	n.fieldMap["event_name"] = n.EventName
	n.fieldMap["resource_type"] = n.ResourceType
	n.fieldMap["resource_id"] = n.ResourceID
	n.fieldMap["customer_id"] = n.CustomerID
	n.fieldMap["to"] = n.To
	n.fieldMap["data"] = n.Data
	n.fieldMap["parent_id"] = n.ParentID
	n.fieldMap["provider_id"] = n.ProviderID
	n.fieldMap["created_at"] = n.CreatedAt
	n.fieldMap["updated_at"] = n.UpdatedAt
}

func (n notification) clone(db *gorm.DB) notification {
	n.notificationDo.ReplaceConnPool(db.Statement.ConnPool)
	return n
}

func (n notification) replaceDB(db *gorm.DB) notification {
	n.notificationDo.ReplaceDB(db)
	return n
}

type notificationDo struct{ gen.DO }

func (n notificationDo) Debug() *notificationDo {
	return n.withDO(n.DO.Debug())
}

func (n notificationDo) WithContext(ctx context.Context) *notificationDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n notificationDo) ReadDB() *notificationDo {
	return n.Clauses(dbresolver.Read)
}

func (n notificationDo) WriteDB() *notificationDo {
	return n.Clauses(dbresolver.Write)
}

func (n notificationDo) Session(config *gorm.Session) *notificationDo {
	return n.withDO(n.DO.Session(config))
}

func (n notificationDo) Clauses(conds ...clause.Expression) *notificationDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n notificationDo) Returning(value interface{}, columns ...string) *notificationDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n notificationDo) Not(conds ...gen.Condition) *notificationDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n notificationDo) Or(conds ...gen.Condition) *notificationDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n notificationDo) Select(conds ...field.Expr) *notificationDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n notificationDo) Where(conds ...gen.Condition) *notificationDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n notificationDo) Order(conds ...field.Expr) *notificationDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n notificationDo) Distinct(cols ...field.Expr) *notificationDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n notificationDo) Omit(cols ...field.Expr) *notificationDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n notificationDo) Join(table schema.Tabler, on ...field.Expr) *notificationDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n notificationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *notificationDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n notificationDo) RightJoin(table schema.Tabler, on ...field.Expr) *notificationDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n notificationDo) Group(cols ...field.Expr) *notificationDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n notificationDo) Having(conds ...gen.Condition) *notificationDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n notificationDo) Limit(limit int) *notificationDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n notificationDo) Offset(offset int) *notificationDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n notificationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *notificationDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n notificationDo) Unscoped() *notificationDo {
	return n.withDO(n.DO.Unscoped())
}

func (n notificationDo) Create(values ...*model.Notification) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n notificationDo) CreateInBatches(values []*model.Notification, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n notificationDo) Save(values ...*model.Notification) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n notificationDo) First() (*model.Notification, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Notification), nil
	}
}

func (n notificationDo) Take() (*model.Notification, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Notification), nil
	}
}

func (n notificationDo) Last() (*model.Notification, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Notification), nil
	}
}

func (n notificationDo) Find() ([]*model.Notification, error) {
	result, err := n.DO.Find()
	return result.([]*model.Notification), err
}

func (n notificationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Notification, err error) {
	buf := make([]*model.Notification, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n notificationDo) FindInBatches(result *[]*model.Notification, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n notificationDo) Attrs(attrs ...field.AssignExpr) *notificationDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n notificationDo) Assign(attrs ...field.AssignExpr) *notificationDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n notificationDo) Joins(fields ...field.RelationField) *notificationDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n notificationDo) Preload(fields ...field.RelationField) *notificationDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n notificationDo) FirstOrInit() (*model.Notification, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Notification), nil
	}
}

func (n notificationDo) FirstOrCreate() (*model.Notification, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Notification), nil
	}
}

func (n notificationDo) FindByPage(offset int, limit int) (result []*model.Notification, count int64, err error) {
	result, err = n.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = n.Offset(-1).Limit(-1).Count()
	return
}

func (n notificationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n notificationDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n notificationDo) Delete(models ...*model.Notification) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *notificationDo) withDO(do gen.Dao) *notificationDo {
	n.DO = *do.(*gen.DO)
	return n
}
