package repository

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const productSalesChannelTable = "product_sales_channel"

type SalesChannelRepo struct {
	sql.Repository[models.SalesChannel]
}

func SalesChannelRepository(db *gorm.DB) *SalesChannelRepo {
	return &SalesChannelRepo{*sql.NewRepository[models.SalesChannel](db)}
}

func (r *SalesChannelRepo) GetFreeTextSearchResults(ctx context.Context, q *string, options sql.Options) ([]models.SalesChannel, *int64, error) {
	var selector models.SalesChannel
	var res []models.SalesChannel

	if q != nil {
		v := sql.ILike(*q)
		selector.Name = v
		selector.Description = v
	}

	query := sql.BuildQuery(selector, options)

	count, err := r.FindAndCount(ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil

}

func (r *SalesChannelRepo) RemoveProducts(salesChannelId uuid.UUID, productIds uuid.UUIDs) *utils.ApplictaionError {
	whereOptions := map[string]interface{}{
		"sales_channel_id": salesChannelId,
		"product_id":       productIds,
	}
	if err := r.Db().Table(productSalesChannelTable).Where(whereOptions).Delete(whereOptions).Error; err != nil {
		return r.HandleDBError(err)
	}
	return nil
}

func (r *SalesChannelRepo) AddProducts(salesChannelId uuid.UUID, productIds uuid.UUIDs, isMedusaV2Enabled bool) *utils.ApplictaionError {
	valuesToInsert := make([]map[string]interface{}, len(productIds))
	for i, id := range productIds {
		valuesToInsert[i] = map[string]interface{}{
			"sales_channel_id": salesChannelId,
			"product_id":       id,
		}
	}
	if isMedusaV2Enabled {
		for i, v := range valuesToInsert {
			var err error
			v["id"], err = uuid.NewUUID()
			if err != nil {
				return utils.NewApplictaionError(
					utils.CONFLICT,
					err.Error(),
					"500",
					nil,
				)
			}
			valuesToInsert[i] = v
		}
	}
	if err := r.Db().Table(productSalesChannelTable).Create(valuesToInsert).Error; err != nil {
		return r.HandleDBError(err)
	}

	return nil
}
func (r *SalesChannelRepo) ListProductIdsBySalesChannelIds(salesChannelIds uuid.UUIDs) (map[string][]string, *utils.ApplictaionError) {
	result, err := r.Db().Table(productSalesChannelTable).Where("sales_channel_id IN ?", salesChannelIds.Strings()).Select("sales_channel_id", "product_id").Rows()
	if err != nil {
		return nil, r.HandleDBError(err)
	}
	acc := make(map[string][]string)
	for result.Next() {
		var salesChannelId, productId string
		err := result.Scan(&salesChannelId, &productId)
		if err != nil {
			return nil, r.HandleDBError(err)
		}
		if _, ok := acc[salesChannelId]; !ok {
			acc[salesChannelId] = []string{}
		}
		acc[salesChannelId] = append(acc[salesChannelId], productId)
	}
	return acc, nil
}
