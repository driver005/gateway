package repository

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PriceListRepo struct {
	sql.Repository[models.PriceList]
}

func PriceListRepository(db *gorm.DB) *PriceListRepo {
	return &PriceListRepo{*sql.NewRepository[models.PriceList](db)}
}

func (r *PriceListRepo) ListAndCount(ctx context.Context, selector *types.FilterablePriceList, config *sql.Options, q *string) ([]models.PriceList, *int64, *utils.ApplictaionError) {
	var res []models.PriceList

	if q != nil {
		v := sql.ILike(*q)
		selector.Name = v
		selector.Description = v
	}

	query := sql.BuildQuery(selector, config)

	count, err := r.FindAndCount(ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (r *PriceListRepo) ListPriceListsVariantIdsMap(priceListIds uuid.UUIDs) (map[string][]string, *utils.ApplictaionError) {
	var data []struct {
		PLID          string
		PVMAVariantId string
	}
	if err := r.Db().Table("pl").
		Joins("inner join prices on pl.id = prices.price_list_id").
		Joins("inner join product_variant_money_amount pvma on pvma.money_amount_id = prices.id").
		Where("pl.id in ?", priceListIds.Strings()).
		Select("pl.id as plid, pvma.variant_id as pvma_variant_id").
		Scan(&data).
		Error; err != nil {
		return nil, r.HandleDBError(err)
	}

	result := make(map[string][]string)
	for _, curr := range data {
		if _, ok := result[curr.PLID]; !ok {
			result[curr.PLID] = []string{}
		}
		result[curr.PLID] = append(result[curr.PLID], curr.PVMAVariantId)
	}
	return result, nil
}
