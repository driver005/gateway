package repository

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LineItemRepo struct {
	sql.Repository[models.LineItem]
	returnItemRepository *ReturnItemRepo
}

func LineItemRepository(db *gorm.DB) *LineItemRepo {
	return &LineItemRepo{*sql.NewRepository[models.LineItem](db), ReturnItemRepository(db)}
}

func (r *LineItemRepo) FindByReturn(ctx context.Context, returnId uuid.UUID) (*models.LineItem, *models.ReturnItem, *utils.ApplictaionError) {
	var res *models.ReturnItem

	query := sql.BuildQuery[models.ReturnItem](models.ReturnItem{
		ReturnId: uuid.NullUUID{
			UUID: returnId,
		},
	}, &sql.Options{})

	if err := r.returnItemRepository.FindOne(ctx, res, query); err != nil {
		return nil, nil, err
	}

	return res.Item, res, nil
}
