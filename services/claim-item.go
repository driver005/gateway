package services

import (
	"context"
	"reflect"
	"slices"
	"strings"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
)

type ClaimItemService struct {
	ctx context.Context
	r   Registry
}

func NewClaimItemService(
	r Registry,
) *ClaimItemService {
	return &ClaimItemService{
		context.Background(),
		r,
	}
}

func (s *ClaimItemService) SetContext(context context.Context) *ClaimItemService {
	s.ctx = context
	return s
}

func (s *ClaimItemService) Retrieve(claimItemId uuid.UUID, config *sql.Options) (*models.ClaimItem, *utils.ApplictaionError) {
	if claimItemId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"claimItemId" must be defined`,
			nil,
		)
	}
	var res *models.ClaimItem

	query := sql.BuildQuery[models.ClaimItem](models.ClaimItem{Model: core.Model{Id: claimItemId}}, config)

	if err := s.r.ClaimItemRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ClaimItemService) List(selector models.ClaimItem, config *sql.Options) ([]models.ClaimItem, *utils.ApplictaionError) {
	var res []models.ClaimItem

	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	query := sql.BuildQuery[models.ClaimItem](selector, config)

	if err := s.r.ClaimItemRepository().Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ClaimItemService) Update(id uuid.UUID, data *types.UpdateClaimItemInput) (*models.ClaimItem, *utils.ApplictaionError) {
	item, err := s.Retrieve(id, &sql.Options{
		Relations: []string{"images", "tags"},
	})
	if err != nil {
		return nil, err
	}

	if data.Note != "" {
		item.Note = data.Note
	}

	if data.Reason != "" {
		item.Reason = data.Reason
	}

	if data.Metadata != nil {
		item.Metadata = data.Metadata
	}

	if data.Tags != nil {
		item.Tags = []models.ClaimTag{}
		for _, tag := range data.Tags {
			if tag.Id != uuid.Nil {
				item.Tags = append(item.Tags, models.ClaimTag{
					Id:    tag.Id,
					Value: tag.Value,
				})
			} else {
				var claimTag *models.ClaimTag

				claimTag.Value = strings.ToLower(strings.TrimSpace(tag.Value))
				if err := s.r.ClaimTagRepository().FindOne(s.ctx, claimTag, sql.Query{}); err != nil {
					if err := s.r.ClaimTagRepository().Create(s.ctx, claimTag); err != nil {
						return nil, err
					}

					item.Tags = append(item.Tags, *claimTag)
				} else {
					item.Tags = append(item.Tags, *claimTag)
				}
			}
		}
	}

	if data.Images != nil {
		var ids uuid.UUIDs
		for i, img := range data.Images {
			ids[i] = img.Id
		}
		for _, img := range item.Images {
			if !slices.Contains(ids, img.Id) {
				if err := s.r.ClaimImageRepository().Remove(s.ctx, &img); err != nil {
					return nil, err
				}
			}
		}
		item.Images = []models.ClaimImage{}
		for _, img := range data.Images {
			if img.Id != uuid.Nil {
				item.Images = append(item.Images, models.ClaimImage{
					Model: core.Model{
						Id: img.Id,
					},
					Url: img.URL,
				})
			} else {
				var claimImage *models.ClaimImage

				claimImage.Url = img.URL

				if err := s.r.ClaimImageRepository().Create(s.ctx, claimImage); err != nil {
					return nil, err
				}

				item.Images = append(item.Images, *claimImage)
			}
		}
	}

	if err := s.r.ClaimItemRepository().Save(s.ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ClaimItemService) Create(data *types.CreateClaimItemInput) (*models.ClaimItem, *utils.ApplictaionError) {
	if data.Reason != models.ClaimReasonTypeMissingItem && data.Reason != models.ClaimReasonTypeWrongItem && data.Reason != models.ClaimReasonTypeProductionFailure && data.Reason != models.ClaimReasonTypeOther {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			`claim item reason must be one ,of "missing_item", "wrong_item", "production_failure" or "other".`,
			nil,
		)
	}

	lineItem, err := s.r.LineItemService().SetContext(s.ctx).Retrieve(data.ItemId, &sql.Options{})
	if err != nil {
		return nil, err
	}

	if lineItem.VariantId.UUID == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Cannot claim a custom line ite,m",
			nil,
		)
	}

	if lineItem.FulfilledQuantity < data.Quantity {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Cannot claim more of an item t,han has been fulfilled.",
			nil,
		)
	}

	var tagsToAdd []models.ClaimTag
	if data.Tags != nil {
		for _, tag := range data.Tags {
			var claimTag *models.ClaimTag

			claimTag.Value = strings.ToLower(strings.TrimSpace(tag))
			if err := s.r.ClaimTagRepository().FindOne(s.ctx, claimTag, sql.Query{}); err != nil {
				if err := s.r.ClaimTagRepository().Create(s.ctx, claimTag); err != nil {
					return nil, err
				}
				tagsToAdd = append(tagsToAdd, *claimTag)
			} else {
				tagsToAdd = append(tagsToAdd, *claimTag)
			}
		}
	}

	var imagesToAdd []models.ClaimImage
	if data.Images != nil {
		for _, url := range data.Images {
			var claimImage *models.ClaimImage

			claimImage.Url = url

			if err := s.r.ClaimImageRepository().Create(s.ctx, claimImage); err != nil {
				return nil, err
			}

			imagesToAdd = append(imagesToAdd, *claimImage)
		}
	}

	model := &models.ClaimItem{
		ItemId:       uuid.NullUUID{UUID: data.ItemId},
		Quantity:     data.Quantity,
		ClaimOrderId: uuid.NullUUID{UUID: data.ClaimOrderId},
		Reason:       data.Reason,
		Note:         data.Note,
		VariantId:    lineItem.VariantId,
		Tags:         tagsToAdd,
		Images:       imagesToAdd,
	}

	if err := s.r.ClaimItemRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}
