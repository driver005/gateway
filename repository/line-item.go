package repository

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LineItemRepo struct {
	Repository[models.LineItem]
	returnItemRepository ReturnItemRepo
}

func LineItemRepository(db *gorm.DB) LineItemRepo {
	return LineItemRepo{*NewRepository[models.LineItem](db), ReturnItemRepository(db)}
}

func (r *LineItemRepo) FindByReturn(ctx context.Context, returnId uuid.UUID) (*models.LineItem, *models.ReturnItem, error) {
	var res *models.ReturnItem

	query := BuildQuery[models.ReturnItem](models.ReturnItem{
		ReturnId: uuid.NullUUID{
			UUID: returnId,
		},
	}, Options{})

	if err := r.returnItemRepository.FindOne(ctx, res, query); err != nil {
		return nil, nil, err
	}

	return res.Item, res, nil
}
