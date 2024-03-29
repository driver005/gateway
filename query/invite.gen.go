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

func newInvite(db *gorm.DB, opts ...gen.DOOption) invite {
	_invite := invite{}

	_invite.inviteDo.UseDB(db, opts...)
	_invite.inviteDo.UseModel(&model.Invite{})

	tableName := _invite.inviteDo.TableName()
	_invite.ALL = field.NewAsterisk(tableName)
	_invite.ID = field.NewString(tableName, "id")
	_invite.Email = field.NewString(tableName, "email")
	_invite.Accepted = field.NewBool(tableName, "accepted")
	_invite.Token = field.NewString(tableName, "token")
	_invite.ExpiresAt = field.NewTime(tableName, "expires_at")
	_invite.Metadata = field.NewString(tableName, "metadata")
	_invite.CreatedAt = field.NewTime(tableName, "created_at")
	_invite.UpdatedAt = field.NewTime(tableName, "updated_at")
	_invite.DeletedAt = field.NewField(tableName, "deleted_at")

	_invite.fillFieldMap()

	return _invite
}

type invite struct {
	inviteDo inviteDo

	ALL       field.Asterisk
	ID        field.String
	Email     field.String
	Accepted  field.Bool
	Token     field.String
	ExpiresAt field.Time
	Metadata  field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (i invite) Table(newTableName string) *invite {
	i.inviteDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i invite) As(alias string) *invite {
	i.inviteDo.DO = *(i.inviteDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *invite) updateTableName(table string) *invite {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewString(table, "id")
	i.Email = field.NewString(table, "email")
	i.Accepted = field.NewBool(table, "accepted")
	i.Token = field.NewString(table, "token")
	i.ExpiresAt = field.NewTime(table, "expires_at")
	i.Metadata = field.NewString(table, "metadata")
	i.CreatedAt = field.NewTime(table, "created_at")
	i.UpdatedAt = field.NewTime(table, "updated_at")
	i.DeletedAt = field.NewField(table, "deleted_at")

	i.fillFieldMap()

	return i
}

func (i *invite) WithContext(ctx context.Context) *inviteDo { return i.inviteDo.WithContext(ctx) }

func (i invite) TableName() string { return i.inviteDo.TableName() }

func (i invite) Alias() string { return i.inviteDo.Alias() }

func (i invite) Columns(cols ...field.Expr) gen.Columns { return i.inviteDo.Columns(cols...) }

func (i *invite) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *invite) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 9)
	i.fieldMap["id"] = i.ID
	i.fieldMap["email"] = i.Email
	i.fieldMap["accepted"] = i.Accepted
	i.fieldMap["token"] = i.Token
	i.fieldMap["expires_at"] = i.ExpiresAt
	i.fieldMap["metadata"] = i.Metadata
	i.fieldMap["created_at"] = i.CreatedAt
	i.fieldMap["updated_at"] = i.UpdatedAt
	i.fieldMap["deleted_at"] = i.DeletedAt
}

func (i invite) clone(db *gorm.DB) invite {
	i.inviteDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i invite) replaceDB(db *gorm.DB) invite {
	i.inviteDo.ReplaceDB(db)
	return i
}

type inviteDo struct{ gen.DO }

func (i inviteDo) Debug() *inviteDo {
	return i.withDO(i.DO.Debug())
}

func (i inviteDo) WithContext(ctx context.Context) *inviteDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i inviteDo) ReadDB() *inviteDo {
	return i.Clauses(dbresolver.Read)
}

func (i inviteDo) WriteDB() *inviteDo {
	return i.Clauses(dbresolver.Write)
}

func (i inviteDo) Session(config *gorm.Session) *inviteDo {
	return i.withDO(i.DO.Session(config))
}

func (i inviteDo) Clauses(conds ...clause.Expression) *inviteDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i inviteDo) Returning(value interface{}, columns ...string) *inviteDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i inviteDo) Not(conds ...gen.Condition) *inviteDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i inviteDo) Or(conds ...gen.Condition) *inviteDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i inviteDo) Select(conds ...field.Expr) *inviteDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i inviteDo) Where(conds ...gen.Condition) *inviteDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i inviteDo) Order(conds ...field.Expr) *inviteDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i inviteDo) Distinct(cols ...field.Expr) *inviteDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i inviteDo) Omit(cols ...field.Expr) *inviteDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i inviteDo) Join(table schema.Tabler, on ...field.Expr) *inviteDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i inviteDo) LeftJoin(table schema.Tabler, on ...field.Expr) *inviteDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i inviteDo) RightJoin(table schema.Tabler, on ...field.Expr) *inviteDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i inviteDo) Group(cols ...field.Expr) *inviteDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i inviteDo) Having(conds ...gen.Condition) *inviteDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i inviteDo) Limit(limit int) *inviteDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i inviteDo) Offset(offset int) *inviteDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i inviteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *inviteDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i inviteDo) Unscoped() *inviteDo {
	return i.withDO(i.DO.Unscoped())
}

func (i inviteDo) Create(values ...*model.Invite) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i inviteDo) CreateInBatches(values []*model.Invite, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i inviteDo) Save(values ...*model.Invite) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i inviteDo) First() (*model.Invite, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Invite), nil
	}
}

func (i inviteDo) Take() (*model.Invite, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Invite), nil
	}
}

func (i inviteDo) Last() (*model.Invite, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Invite), nil
	}
}

func (i inviteDo) Find() ([]*model.Invite, error) {
	result, err := i.DO.Find()
	return result.([]*model.Invite), err
}

func (i inviteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Invite, err error) {
	buf := make([]*model.Invite, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i inviteDo) FindInBatches(result *[]*model.Invite, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i inviteDo) Attrs(attrs ...field.AssignExpr) *inviteDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i inviteDo) Assign(attrs ...field.AssignExpr) *inviteDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i inviteDo) Joins(fields ...field.RelationField) *inviteDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i inviteDo) Preload(fields ...field.RelationField) *inviteDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i inviteDo) FirstOrInit() (*model.Invite, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Invite), nil
	}
}

func (i inviteDo) FirstOrCreate() (*model.Invite, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Invite), nil
	}
}

func (i inviteDo) FindByPage(offset int, limit int) (result []*model.Invite, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i inviteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i inviteDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i inviteDo) Delete(models ...*model.Invite) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *inviteDo) withDO(do gen.Dao) *inviteDo {
	i.DO = *do.(*gen.DO)
	return i
}
