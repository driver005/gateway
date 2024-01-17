package services

import (
	"context"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type LineItemAdjustmentService struct {
	ctx context.Context
	r   Registry
}

func NewLineItemAdjustmentService(
	r Registry,
) *LineItemAdjustmentService {
	return &LineItemAdjustmentService{
		context.Background(),
		r,
	}
}

func (s *LineItemAdjustmentService) SetContext(context context.Context) *LineItemAdjustmentService {
	s.ctx = context
	return s
}

func (s *LineItemAdjustmentService) Retrieve(id uuid.UUID, config *sql.Options) (*models.LineItemAdjustment, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			`"id" must be defined`,
			nil,
		)
	}
	var res *models.LineItemAdjustment
	query := sql.BuildQuery(models.LineItemAdjustment{Model: core.Model{Id: id}}, config)

	if err := s.r.LineItemAdjustmentRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *LineItemAdjustmentService) Create(data *models.LineItemAdjustment) (*models.LineItemAdjustment, *utils.ApplictaionError) {
	if err := s.r.LineItemAdjustmentRepository().Save(s.ctx, data); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *LineItemAdjustmentService) Update(id uuid.UUID, Update *models.LineItemAdjustment) (*models.LineItemAdjustment, *utils.ApplictaionError) {
	lineItemAdjustment, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	Update.Id = lineItemAdjustment.Id

	if err := s.r.LineItemAdjustmentRepository().Update(s.ctx, Update); err != nil {
		return nil, err
	}

	return Update, nil
}

func (s *LineItemAdjustmentService) List(selector types.FilterableLineItemAdjustmentProps, config *sql.Options) ([]models.LineItemAdjustment, *utils.ApplictaionError) {
	var res []models.LineItemAdjustment
	query := sql.BuildQuery(selector, config)

	if err := s.r.LineItemAdjustmentRepository().Find(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *LineItemAdjustmentService) Delete(id uuid.UUID, selector *models.LineItemAdjustment, config *sql.Options) *utils.ApplictaionError {
	var data *models.LineItemAdjustment
	if id != uuid.Nil {
		var err *utils.ApplictaionError
		data, err = s.Retrieve(id, config)
		if err != nil {
			return err
		}
	} else if selector != nil {
		data = selector
	}

	if err := s.r.LineItemAdjustmentRepository().SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}

func (s *LineItemAdjustmentService) DeleteSlice(ids uuid.UUIDs, selector []models.LineItemAdjustment) *utils.ApplictaionError {
	var data []models.LineItemAdjustment
	if ids != nil {
		for _, id := range ids {
			var err *utils.ApplictaionError
			lineItem, err := s.Retrieve(id, &sql.Options{})
			if err != nil {
				return err
			}
			data = append(data, *lineItem)
		}

	} else if selector != nil {
		data = selector
	}

	if err := s.r.LineItemAdjustmentRepository().RemoveSlice(s.ctx, data); err != nil {
		return err
	}

	return nil
}

func (s *LineItemAdjustmentService) GenerateAdjustments(calculationContextData types.CalculationContextData, generatedLineItem *models.LineItem, context *models.ProductVariant) ([]models.LineItemAdjustment, *utils.ApplictaionError) {
	lineItem := generatedLineItem

	if !lineItem.AllowDiscounts || lineItem.IsReturn || len(calculationContextData.Discounts) == 0 {
		return nil, nil
	}

	var discount *models.Discount
	for _, d := range calculationContextData.Discounts {
		if d.Rule.Type != models.DiscountRuleFreeShipping {
			discount = &d
			break
		}
	}

	if discount == nil {
		return nil, nil
	}

	isValid, err := s.r.DiscountService().SetContext(s.ctx).ValidateDiscountForProduct(discount.RuleId.UUID, context.ProductId.UUID)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, nil
	}

	lineItem.Id = generatedLineItem.Id

	amount, err := s.r.DiscountService().SetContext(s.ctx).CalculateDiscountForLineItem(discount.Id, lineItem, calculationContextData)
	if err != nil {
		return nil, err
	}
	if amount == 0 {
		return nil, nil
	}

	adjustment := models.LineItemAdjustment{
		Amount:      amount,
		DiscountId:  uuid.NullUUID{UUID: discount.Id},
		Description: "discount",
	}

	return []models.LineItemAdjustment{adjustment}, nil
}

func (s *LineItemAdjustmentService) CreateAdjustmentForLineItem(cart *models.Cart, lineItem *models.LineItem) ([]models.LineItemAdjustment, *utils.ApplictaionError) {
	adjustments, err := s.GenerateAdjustments(types.CalculationContextData{
		Discounts:       cart.Discounts,
		Items:           cart.Items,
		Customer:        cart.Customer,
		Region:          cart.Region,
		ShippingAddress: cart.ShippingAddress,
		ShippingMethods: cart.ShippingMethods,
	}, lineItem, lineItem.Variant)
	if err != nil {
		return nil, err
	}

	var createdAdjustments []models.LineItemAdjustment
	for _, adjustment := range adjustments {
		created, err := s.Create(&adjustment)
		if err != nil {
			return nil, err
		}
		createdAdjustments = append(createdAdjustments, *created)
	}

	return createdAdjustments, nil
}

func (s *LineItemAdjustmentService) CreateAdjustments(cart *models.Cart, lineItem *models.LineItem) ([]models.LineItemAdjustment, [][]models.LineItemAdjustment, *utils.ApplictaionError) {
	if lineItem != nil {
		adjustments, err := s.CreateAdjustmentForLineItem(cart, lineItem)
		if err != nil {
			return nil, nil, err
		}
		return adjustments, nil, nil
	}

	if len(cart.Items) == 0 {
		return nil, nil, nil
	}

	var allAdjustments [][]models.LineItemAdjustment
	for _, li := range cart.Items {
		adjustments, err := s.CreateAdjustmentForLineItem(cart, &li)
		if err != nil {
			return nil, nil, err
		}
		allAdjustments = append(allAdjustments, adjustments)
	}

	return nil, allAdjustments, nil
}
