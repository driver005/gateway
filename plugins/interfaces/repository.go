package interfaces

import (
	"context"

	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"gorm.io/gorm"
)

// RepositoryInterface defines the interface for the generic repository
type RepositoryInterface[M any] interface {
	Relations(relations []string) RepositoryInterface[M]
	Db() *gorm.DB
	Specification(specification ...sql.Specification) RepositoryInterface[M]

	Create(ctx context.Context, model *M) *utils.ApplictaionError
	CreateSlice(ctx context.Context, model *[]M) *utils.ApplictaionError
	CreateBatch(ctx context.Context, model *M, batchSize int) *utils.ApplictaionError
	CreateTx(ctx context.Context, model *M, tx *gorm.DB) *utils.ApplictaionError

	Insert(ctx context.Context, model *M) *utils.ApplictaionError
	InsertSlice(ctx context.Context, model *[]M) *utils.ApplictaionError
	InsertTx(ctx context.Context, model *M, tx *gorm.DB) *utils.ApplictaionError

	Save(ctx context.Context, model *M) *utils.ApplictaionError
	SaveSlice(ctx context.Context, model *[]M) *utils.ApplictaionError
	SaveTx(ctx context.Context, model *M, tx *gorm.DB) *utils.ApplictaionError

	Update(ctx context.Context, model *M) *utils.ApplictaionError
	UpdateSlice(ctx context.Context, model *[]M) *utils.ApplictaionError

	Upsert(ctx context.Context, model *M) *utils.ApplictaionError

	Delete(ctx context.Context, model *M) *utils.ApplictaionError
	DeleteSlice(ctx context.Context, model []M) *utils.ApplictaionError
	DeleteTx(ctx context.Context, model *M, tx *gorm.DB) *utils.ApplictaionError
	DeletePermanently(ctx context.Context, model *M) *utils.ApplictaionError

	Remove(ctx context.Context, model *M) *utils.ApplictaionError
	RemoveSlice(ctx context.Context, model []M) *utils.ApplictaionError
	RemoveTX(ctx context.Context, model *M, tx *gorm.DB) *utils.ApplictaionError

	SoftRemove(ctx context.Context, model *M) *utils.ApplictaionError
	Recover(ctx context.Context, model *M) *utils.ApplictaionError

	Count(ctx context.Context, query sql.Query) (*int64, *utils.ApplictaionError)
	CountBy(ctx context.Context, field []string, query sql.Query) (*int64, *utils.ApplictaionError)

	Find(ctx context.Context, models *[]M, query sql.Query) *utils.ApplictaionError
	FindAndCount(ctx context.Context, models *[]M, query sql.Query) (*int64, *utils.ApplictaionError)
	FindOne(ctx context.Context, model *M, query sql.Query) *utils.ApplictaionError

	ParseQuery(ctx context.Context, query sql.Query) *gorm.DB

	Clear(target interface{}) *utils.ApplictaionError

	HandleOneError(res *gorm.DB) *utils.ApplictaionError
	HandleError(res *gorm.DB) *utils.ApplictaionError
	HandleDBError(err error) *utils.ApplictaionError
}
