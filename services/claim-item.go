package services

import (
	"context"
	"reflect"
	"slices"
	"strings"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
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

func (s *ClaimItemService) Retrieve(claimItemId uuid.UUID, config sql.Options) (*models.ClaimItem, *utils.ApplictaionError) {
	if claimItemId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"claimItemId" must be defined`,
			"500",
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

func (s *ClaimItemService) List(selector models.ClaimItem, config sql.Options) ([]models.ClaimItem, *utils.ApplictaionError) {
	var res []models.ClaimItem

	if reflect.DeepEqual(config, sql.Options{}) {
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

func (s *ClaimItemService) Update(id uuid.UUID, model *models.ClaimItem) (*models.ClaimItem, *utils.ApplictaionError) {
	item, err := s.Retrieve(id, sql.Options{
		Relations: []string{"images", "tags"},
	})
	if err != nil {
		return nil, err
	}

	if model.Note != "" {
		item.Note = model.Note
	}

	if model.Reason != "" {
		item.Reason = model.Reason
	}

	if model.Metadata != nil {
		item.Metadata = model.Metadata
	}

	if model.Tags != nil {
		item.Tags = []models.ClaimTag{}
		for _, tag := range model.Tags {
			if tag.Id != uuid.Nil {
				item.Tags = append(item.Tags, tag)
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

	if model.Images != nil {
		var ids uuid.UUIDs
		for i, img := range model.Images {
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
		for _, img := range model.Images {
			if img.Id != uuid.Nil {
				item.Images = append(item.Images, img)
			} else {
				var claimImage *models.ClaimImage

				claimImage.Url = img.Url

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

func (s *ClaimItemService) Create(model *models.ClaimItem) (*models.ClaimItem, *utils.ApplictaionError) {
	if model.Reason != "missing_item" && model.Reason != "wrong_item" && model.Reason != "production_failure" && model.Reason != "other" {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			`claim item reason must be one ,of "missing_item", "wrong_item", "production_failure" or "other".`,
			"500",
			nil,
		)
	}

	lineItem, err := s.r.LineItemService().SetContext(s.ctx).Retrieve(model.ItemId.UUID, sql.Options{})
	if err != nil {
		return nil, err
	}

	if lineItem.VariantId.UUID == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Cannot claim a custom line ite,m",
			"500",
			nil,
		)
	}

	if lineItem.FulfilledQuantity < model.Quantity {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Cannot claim more of an item t,han has been fulfilled.",
			"500",
			nil,
		)
	}

	var tagsToAdd []models.ClaimTag
	if model.Tags != nil {
		for _, tag := range model.Tags {
			var claimTag *models.ClaimTag

			claimTag.Value = strings.ToLower(strings.TrimSpace(tag.Value))
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
	if model.Images != nil {
		for _, img := range model.Images {
			var claimImage *models.ClaimImage

			claimImage.Url = img.Url

			if err := s.r.ClaimImageRepository().Create(s.ctx, claimImage); err != nil {
				return nil, err
			}

			imagesToAdd = append(imagesToAdd, *claimImage)
		}
	}

	model.Tags = tagsToAdd
	model.Images = imagesToAdd

	if err := s.r.ClaimItemRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}
