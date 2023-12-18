package services

import (
	"context"
	"errors"
	"reflect"
	"slices"
	"strings"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
)

type ClaimItemService struct {
	ctx                context.Context
	repo               *repository.ClaimItemRepo
	claimTagRepository *repository.ClaimTagRepo
	claimImageRepo     *repository.ClaimImageRepo
	lineItemService    *LineItemService
}

func NewClaimItemService(
	ctx context.Context,
	repo *repository.ClaimItemRepo,
	claimTagRepository *repository.ClaimTagRepo,
	claimImageRepo *repository.ClaimImageRepo,
	lineItemService *LineItemService,
) *ClaimItemService {
	return &ClaimItemService{
		ctx,
		repo,
		claimTagRepository,
		claimImageRepo,
		lineItemService,
	}
}

func (s *ClaimItemService) Retrieve(claimItemId uuid.UUID, config repository.Options) (*models.ClaimItem, error) {
	if claimItemId == uuid.Nil {
		return nil, errors.New(`"claimItemId" must be defined`)
	}
	var res *models.ClaimItem

	query := repository.BuildQuery[models.ClaimItem](models.ClaimItem{Model: core.Model{Id: claimItemId}}, config)

	if err := s.repo.FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ClaimItemService) List(selector models.ClaimItem, config repository.Options) ([]models.ClaimItem, error) {
	var res []models.ClaimItem

	if reflect.DeepEqual(config, repository.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	query := repository.BuildQuery[models.ClaimItem](selector, config)

	if err := s.repo.Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ClaimItemService) Update(id uuid.UUID, model *models.ClaimItem) (*models.ClaimItem, error) {
	item, err := s.Retrieve(id, repository.Options{
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
				if err := s.claimTagRepository.FindOne(s.ctx, claimTag, repository.Query{}); err != nil {
					if err := s.claimTagRepository.Create(s.ctx, claimTag); err != nil {
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
				if err := s.claimImageRepo.Remove(s.ctx, &img); err != nil {
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

				if err := s.claimImageRepo.Create(s.ctx, claimImage); err != nil {
					return nil, err
				}

				item.Images = append(item.Images, *claimImage)
			}
		}
	}

	if err := s.repo.Save(s.ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ClaimItemService) Create(model *models.ClaimItem) (*models.ClaimItem, error) {
	if model.Reason != "missing_item" && model.Reason != "wrong_item" && model.Reason != "production_failure" && model.Reason != "other" {
		return nil, errors.New(`claim item reason must be one of "missing_item", "wrong_item", "production_failure" or "other".`)
	}

	lineItem, err := s.lineItemService.Retrieve(model.ItemId.UUID, repository.Options{})
	if err != nil {
		return nil, err
	}

	if lineItem.VariantId.UUID == uuid.Nil {
		return nil, errors.New("Cannot claim a custom line item")
	}

	if lineItem.FulfilledQuantity < model.Quantity {
		return nil, errors.New("Cannot claim more of an item than has been fulfilled.")
	}

	var tagsToAdd []models.ClaimTag
	if model.Tags != nil {
		for _, tag := range model.Tags {
			var claimTag *models.ClaimTag

			claimTag.Value = strings.ToLower(strings.TrimSpace(tag.Value))
			if err := s.claimTagRepository.FindOne(s.ctx, claimTag, repository.Query{}); err != nil {
				if err := s.claimTagRepository.Create(s.ctx, claimTag); err != nil {
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

			if err := s.claimImageRepo.Create(s.ctx, claimImage); err != nil {
				return nil, err
			}

			imagesToAdd = append(imagesToAdd, *claimImage)
		}
	}

	model.Tags = tagsToAdd
	model.Images = imagesToAdd

	if err := s.repo.Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}
