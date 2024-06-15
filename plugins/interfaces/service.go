package interfaces

import (
	"context"

	"github.com/driver005/gateway/sql"
)

type InternalModuleService[TEntity any, TContainer any] interface {
	GetContainer() TContainer
	Retrieve(idOrObject interface{}, config *sql.Options, sharedContext *context.Context) (*TEntity, error)
	List(filters interface{}, config *sql.Options, sharedContext *context.Context) ([]TEntity, error)
	ListAndCount(filters interface{}, config *sql.Options, sharedContext *context.Context) ([]TEntity, int, error)
	Create(data interface{}, sharedContext *context.Context) ([]TEntity, error)
	Update(data interface{}, sharedContext *context.Context) ([]TEntity, error)
	Delete(idOrSelector interface{}, sharedContext *context.Context) error
	SoftDelete(idsOrFilter interface{}, sharedContext *context.Context) ([]TEntity, map[string][]interface{}, error)
	Restore(idsOrFilter interface{}, sharedContext *context.Context) ([]TEntity, map[string][]interface{}, error)
	Upsert(data interface{}, sharedContext *context.Context) ([]TEntity, error)
}
