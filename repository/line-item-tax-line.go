package repository

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LineItemTaxLineRepo struct {
	sql.Repository[models.LineItemTaxLine]
	cartRepository *CartRepo
}

func LineItemTaxLineRepository(db *gorm.DB) *LineItemTaxLineRepo {
	return &LineItemTaxLineRepo{*sql.NewRepository[models.LineItemTaxLine](db), CartRepository(db)}
}

func (r *LineItemTaxLineRepo) DeleteForCart(ctx context.Context, cartId uuid.UUID) *utils.ApplictaionError {
	var cart models.Cart
	cart.Id = cartId

	if err := r.cartRepository.FindOne(ctx, &cart, sql.Query{}); err != nil {
		return err
	}

	for _, item := range cart.Items {
		r.Delete(ctx, &models.LineItemTaxLine{
			ItemId: uuid.NullUUID{
				UUID: item.Id,
			},
		})
	}

	return nil
}
