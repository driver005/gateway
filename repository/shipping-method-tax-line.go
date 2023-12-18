package repository

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShippingMethodTaxLineRepo struct {
	Repository[models.ShippingMethodTaxLine]
	cartRepository CartRepo
}

func ShippingMethodTaxLineRepository(db *gorm.DB) ShippingMethodTaxLineRepo {
	return ShippingMethodTaxLineRepo{*NewRepository[models.ShippingMethodTaxLine](db), CartRepository(db)}
}

func (r *ShippingMethodTaxLineRepo) DeleteForCart(ctx context.Context, cartId uuid.UUID) error {
	var cart models.Cart
	cart.Id = cartId

	if err := r.cartRepository.FindOne(ctx, &cart, Query{}); err != nil {
		return err
	}

	for _, shipping_method := range cart.ShippingMethods {
		r.Delete(ctx, &models.ShippingMethodTaxLine{
			ShippingMethodId: uuid.NullUUID{
				UUID: shipping_method.Id,
			},
		})
	}

	return nil
}
