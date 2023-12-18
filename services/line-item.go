package services

import (
	"context"
	"errors"
	"reflect"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/driver005/gateway/types"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
)

type LineItemService struct {
	ctx                 context.Context
	repo                *repository.LineItemRepo
	lineItemTaxLineRepo *repository.LineItemTaxLineRepo
	cartRepository      *repository.CartRepo
}

func NewLineItemService(
	ctx context.Context,
	repo *repository.LineItemRepo,
	lineItemTaxLineRepo *repository.LineItemTaxLineRepo,
	cartRepository *repository.CartRepo,
) *LineItemService {
	return &LineItemService{
		ctx,
		repo,
		lineItemTaxLineRepo,
		cartRepository,
	}
}

func (s *LineItemService) Retrieve(lineItemId uuid.UUID, config repository.Options) (*models.LineItem, error) {
	if lineItemId == uuid.Nil {
		return nil, errors.New(`"lineItemId" must be defined`)
	}
	var res *models.LineItem

	query := repository.BuildQuery[models.LineItem](models.LineItem{Model: core.Model{Id: lineItemId}}, config)

	if err := s.repo.FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LineItemService) List(selector models.LineItem, config repository.Options) ([]models.LineItem, error) {
	var res []models.LineItem

	if reflect.DeepEqual(config, repository.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	query := repository.BuildQuery[models.LineItem](selector, config)

	if err := s.repo.Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LineItemService) CreateReturnLines(returnId uuid.UUID, cartId uuid.UUID) (*models.LineItem, error) {
	lineItem, returnItem, err := s.repo.FindByReturn(s.ctx, returnId)
	if err != nil {
		return nil, err
	}

	model := &models.LineItem{
		Model: core.Model{
			Metadata: lineItem.Metadata,
		},
		CartId: uuid.NullUUID{
			UUID:  cartId,
			Valid: true,
		},
		Thumbnail:      lineItem.Thumbnail,
		IsReturn:       true,
		Title:          lineItem.Title,
		VariantId:      lineItem.VariantId,
		UnitPrice:      -1 * lineItem.UnitPrice,
		Quantity:       returnItem.Quantity,
		AllowDiscounts: lineItem.AllowDiscounts,
		IncludesTax:    lineItem.IncludesTax,
		TaxLines:       lineItem.TaxLines,
		Adjustments:    lineItem.Adjustments,
	}

	if err := s.repo.Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *LineItemService) Generate(
	variantId *string,
	variant []types.GenerateInputData,
	regionId *string,
	quantity *int,
	context types.GenerateLineItemContext,
) ([]models.LineItem, error) {
	return nil, nil
}
