package repository

import (
	"context"

	"github.com/driver005/gateway/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository[M any] struct {
	db *gorm.DB
}

func NewRepository[M any](db *gorm.DB) *Repository[M] {
	return &Repository[M]{
		db: db,
	}
}

func (r *Repository[M]) Relations(relations []string) *Repository[M] {
	for _, relation := range relations {
		r.db.Association(relation)
	}

	return r
}

func (r *Repository[M]) Specification(specification ...Specification) *Repository[M] {
	for _, s := range specification {
		r.db.Where(s.GetQuery(), s.GetValues()...)
	}
	return r
}

// Create inserts value, returning the inserted data's primary key in value's id
func (r *Repository[M]) Create(ctx context.Context, model *M) *utils.ApplictaionError {
	res := r.db.WithContext(ctx).Create(&model)
	return r.HandleError(res)
}

// CreateInBatches inserts value, returning the inserted data's primary key in value's id
func (r *Repository[M]) CreateBatch(ctx context.Context, model *M, batchSize int) *utils.ApplictaionError {
	res := r.db.WithContext(ctx).CreateInBatches(&model, batchSize)
	return r.HandleError(res)
}

// CreateTx inserts value, returning the inserted data's primary key in value's id
func (r *Repository[M]) CreateTx(ctx context.Context, model *M, tx *gorm.DB) *utils.ApplictaionError {
	res := tx.WithContext(ctx).Create(&model)
	return r.HandleError(res)
}

// Insert inserts value, returning the inserted data's primary key in value's id
func (r *Repository[M]) Insert(ctx context.Context, model *M) *utils.ApplictaionError {
	return r.Create(ctx, model)
}

// InsertTx inserts value, returning the inserted data's primary key in value's id
func (r *Repository[M]) InsertTx(ctx context.Context, model *M, tx *gorm.DB) *utils.ApplictaionError {
	return r.CreateTx(ctx, model, tx)
}

// Save updates value in database. If value doesn't contain a matching primary key, value is inserted.
func (r *Repository[M]) Save(ctx context.Context, model *M) *utils.ApplictaionError {
	res := r.db.WithContext(ctx).Save(&model)
	return r.HandleError(res)
}

// Save updates value in database. If value doesn't contain a matching primary key, value is inserted.
func (r *Repository[M]) SaveSlice(ctx context.Context, model []M) *utils.ApplictaionError {
	res := r.db.WithContext(ctx).Save(&model)
	return r.HandleError(res)
}

// SaveTx updates value in database. If value doesn't contain a matching primary key, value is inserted.
func (r *Repository[M]) SaveTx(ctx context.Context, model *M, tx *gorm.DB) *utils.ApplictaionError {
	res := tx.WithContext(ctx).Save(&model)
	return r.HandleError(res)
}

// Update updates attributes using callbacks. values must be a struct or map. Reference: https://gorm.io/docs/update.html#Update-Changed-Fields
func (r *Repository[M]) Update(ctx context.Context, model *M) *utils.ApplictaionError {
	res := r.db.Model(&model).Updates(&model)

	return r.HandleError(res)
}

// Upsert inserts a given entity into the database, unless a unique constraint conflicts then updates the entity Unlike save method executes a primitive operation without cascades, relations and other operations included. Executes fast and efficient INSERT ... ON CONFLICT DO UPDATE/ON DUPLICATE KEY UPDATE query.
func (r *Repository[M]) Upsert(ctx context.Context, model *M) *utils.ApplictaionError {
	res := r.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&model)
	return r.HandleError(res)
}

// Delete deletes value matching given conditions. If value contains primary key it is included in the conditions. If value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current time if null.
func (r *Repository[M]) Delete(ctx context.Context, model *M) *utils.ApplictaionError {
	res := r.db.WithContext(ctx).Delete(model)
	return r.HandleError(res)
}

// Remove deletes value matching given conditions. If value contains primary key it is included in the conditions. If value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current time if null.
func (r *Repository[M]) DeleteSlice(ctx context.Context, model []M) *utils.ApplictaionError {
	res := r.db.WithContext(ctx).Delete(model)
	return r.HandleError(res)
}

// DeleteTx deletes value matching given conditions. If value contains primary key it is included in the conditions. If value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current time if null.
func (r *Repository[M]) DeleteTx(ctx context.Context, model *M, tx *gorm.DB) *utils.ApplictaionError {
	res := tx.WithContext(ctx).Delete(model)
	return r.HandleError(res)
}

// DeletePermanently deletes value matching given conditions. If value contains primary key it is included in the conditions.
func (r *Repository[M]) DeletePermanently(ctx context.Context, model *M) *utils.ApplictaionError {
	res := r.db.WithContext(ctx).Unscoped().Delete(model)
	return r.HandleError(res)
}

// Remove deletes value matching given conditions. If value contains primary key it is included in the conditions. If value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current time if null.
func (r *Repository[M]) Remove(ctx context.Context, model *M) *utils.ApplictaionError {
	return r.Delete(ctx, model)
}

// Remove deletes value matching given conditions. If value contains primary key it is included in the conditions. If value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current time if null.
func (r *Repository[M]) RemoveSlice(ctx context.Context, model []M) *utils.ApplictaionError {
	return r.DeleteSlice(ctx, model)
}

// RemoveTX deletes value matching given conditions. If value contains primary key it is included in the conditions. If value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current time if null.
func (r *Repository[M]) RemoveTX(ctx context.Context, model *M, tx *gorm.DB) *utils.ApplictaionError {
	return r.DeleteTx(ctx, model, tx)
}

// SoftRemove deletes value matching given conditions. If value contains primary key it is included in the conditions. If value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current time if null.
func (r *Repository[M]) SoftRemove(ctx context.Context, model *M) *utils.ApplictaionError {
	return r.Remove(ctx, model)
}

// Recover recovers given entitie in the database.
func (r *Repository[M]) Recover(ctx context.Context, model *M) *utils.ApplictaionError {
	res := r.db.Model(&model).Update("deleted_at", nil)

	return r.HandleError(res)
}

// Count counts all records matching given conditions conds
func (r *Repository[M]) Count(ctx context.Context, query Query) (*int64, *utils.ApplictaionError) {
	model := new(M)
	var count int64

	r.db.WithContext(ctx)
	res := r.ParseQuery(ctx, query).Find(model).Count(&count)

	return &count, r.HandleError(res)
}

// Find finds all records matching given conditions conds
func (r *Repository[M]) Find(ctx context.Context, models []M, query Query) *utils.ApplictaionError {
	res := r.ParseQuery(ctx, query).Find(&models)

	return r.HandleError(res)
}

// FindAndCount finds all records matching given conditions conds and count them
func (r *Repository[M]) FindAndCount(ctx context.Context, models []M, query Query) (*int64, *utils.ApplictaionError) {
	res := r.ParseQuery(ctx, query).Find(&models)

	return &res.RowsAffected, r.HandleError(res)
}

// FindOne finds the first record ordered by primary key, matching given conditions conds
func (r *Repository[M]) FindOne(ctx context.Context, model *M, query Query) *utils.ApplictaionError {
	res := r.ParseQuery(ctx, query).First(&model)

	return r.HandleError(res)
}

func (r *Repository[M]) ParseQuery(ctx context.Context, query Query) *gorm.DB {
	model := new(M)
	db := r.db.WithContext(ctx).Model(model)

	if query.Where != nil {
		db = db.Where(query.Where)
	}
	if query.WithDeleted {
		db = db.Unscoped()
	}
	if query.Skip != nil {
		db = db.Offset(*query.Skip)
	}
	if query.Take != nil {
		db = db.Limit(*query.Take)
	}
	if query.Relations != nil {
		for _, relation := range query.Relations {
			db = db.Association(relation).DB
		}
	}
	if query.Selects != nil {
		db = db.Select(query.Selects)
	}
	if query.Order != nil {
		db = db.Order(query.Order)
	}

	return db
}

func (r *Repository[M]) Clear(target interface{}) *utils.ApplictaionError {
	res := r.db.Migrator().DropTable(target)

	return utils.NewApplictaionError(
		utils.DB_ERROR,
		res.Error(),
		"500",
		nil,
	)
}

func (r *Repository[M]) HandleOneError(res *gorm.DB) *utils.ApplictaionError {
	if err := r.HandleError(res); err != nil {
		return err
	}

	if res.RowsAffected != 1 {
		return utils.NewApplictaionError(
			utils.DB_ERROR,
			gorm.ErrRecordNotFound.Error(),
			"500",
			nil,
		)
	}

	return nil
}

func (r *Repository[M]) HandleError(res *gorm.DB) *utils.ApplictaionError {
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return utils.NewApplictaionError(
			utils.DB_ERROR,
			res.Error.Error(),
			"500",
			nil,
		)
	}

	return nil
}

// func (r *Repository[M]) FindWithLimit(ctx context.Context, limit int, offset int, specifications ...Specification) ([]M, *utils.ApplictaionError) {
// 	var models []M

// 	dbPrewarm := r.getPreWarmDbForSelect(ctx, specifications...)
// 	err := dbPrewarm.Limit(limit).Offset(offset).Find(&models).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	result := make([]M, 0, len(models))
// 	result = append(result, models...)

// 	return result, nil
// }

// func (r *Repository[M]) FindAll(ctx context.Context) ([]M, *utils.ApplictaionError) {
// 	return r.FindWithLimit(ctx, -1, -1)
// }
