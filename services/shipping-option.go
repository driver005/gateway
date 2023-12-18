package services

import (
	"context"
	"fmt"
	"reflect"
	"slices"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
)

type ShippingOptionService struct {
	ctx                   context.Context
	repo                  *repository.ShippingOptionRepo
	methodRepository      *repository.ShippingMethodRepo
	requirementRepository *repository.ShippingOptionRequirementRepo
}

func NewShippingOptionService(
	ctx context.Context,
	repo *repository.ShippingOptionRepo,
	methodRepository *repository.ShippingMethodRepo,
	requirementRepository *repository.ShippingOptionRequirementRepo,
) *ShippingOptionService {
	return &ShippingOptionService{
		ctx,
		repo,
		methodRepository,
		requirementRepository,
	}
}
func (s *ShippingOptionService) ValidateRequirement(requirement *models.ShippingOptionRequirement, optionId uuid.UUID) (*models.ShippingOptionRequirement, *utils.ApplictaionError) {
	if requirement.Type == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"A Shipping Requirement must have a type field",
			"500",
			nil,
		)
	}
	if requirement.Type != "min_subtotal" && requirement.Type != "max_subtotal" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Requirement type must be one of min_subtotal, max_subtotal",
			"500",
			nil,
		)
	}

	existingReq := &models.ShippingOptionRequirement{}
	if requirement.Id != uuid.Nil {
		query := repository.BuildQuery(models.ShippingOptionRequirement{Model: core.Model{Id: requirement.Id}}, repository.Options{})
		err := s.requirementRepository.FindOne(s.ctx, existingReq, query)
		if err != nil {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				fmt.Sprintf("Shipping option requirement with id %s does not exist", requirement.Id),
				"500",
				nil,
			)
		}
	}
	if existingReq == nil && requirement.Id != uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("Shipping option requirement with id %s does not exist", requirement.Id),
			"500",
			nil,
		)
	}
	if optionId == uuid.Nil {
		return requirement, nil
	}
	req := &models.ShippingOptionRequirement{}
	if existingReq != nil {
		req = existingReq
		req.Type = requirement.Type
		req.Amount = requirement.Amount
	} else {
		req = &models.ShippingOptionRequirement{
			Type:             requirement.Type,
			Amount:           requirement.Amount,
			ShippingOptionId: uuid.NullUUID{UUID: optionId},
		}
	}
	if err := s.requirementRepository.Save(s.ctx, req); err != nil {
		return nil, err
	}

	return requirement, nil
}

func (s *ShippingOptionService) List(selector models.ShippingOption, config repository.Options) ([]models.ShippingOption, *utils.ApplictaionError) {
	var res []models.ShippingOption
	query := repository.BuildQuery(selector, config)
	if err := s.repo.Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ShippingOptionService) ListAndCount(selector models.ShippingOption, config repository.Options) ([]models.ShippingOption, *int64, *utils.ApplictaionError) {
	var res []models.ShippingOption

	if reflect.DeepEqual(config, repository.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
	}

	query := repository.BuildQuery(selector, config)
	count, err := s.repo.FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *ShippingOptionService) Retrieve(optionId uuid.UUID, options repository.Options) (*models.ShippingOption, *utils.ApplictaionError) {
	if optionId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			`"optionId" must be defined`,
			"500",
			nil,
		)
	}
	var res *models.ShippingOption
	query := repository.BuildQuery(models.ShippingOption{Model: core.Model{Id: optionId}}, options)
	if err := s.repo.FindOne(s.ctx, res, query); err != nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("Shipping Option with %s was not found", optionId),
			"500",
			nil,
		)
	}

	return res, nil
}

func (s *ShippingOptionService) UpdateShippingMethod(id uuid.UUID, update *models.ShippingMethod) (*models.ShippingMethod, *utils.ApplictaionError) {
	update.Id = id

	if err := s.methodRepository.FindOne(s.ctx, update, repository.Query{}); err != nil {
		return nil, utils.NewApplictaionError(
			utils.DB_ERROR,
			err.Error(),
			"500",
			nil,
		)
	}

	if err := s.methodRepository.Upsert(s.ctx, update); err != nil {
		return nil, utils.NewApplictaionError(
			utils.DB_ERROR,
			err.Error(),
			"500",
			nil,
		)
	}

	return update, nil
}

func (s *ShippingOptionService) DeleteShippingMethods(shippingMethods []models.ShippingMethod) *utils.ApplictaionError {
	err := s.methodRepository.RemoveSlice(s.ctx, shippingMethods)
	if err != nil {
		return err
	}
	return nil
}

func (s *ShippingOptionService) CreateShippingMethod(optionId uuid.UUID, data map[string]interface{}, config *models.ShippingMethod) (*models.ShippingMethod, *utils.ApplictaionError) {
	option, err := s.Retrieve(optionId, repository.Options{
		Relations: []string{"requirements"},
	})
	if err != nil {
		return nil, err
	}

	if config.Cart != nil {
		var err *utils.ApplictaionError
		option, err = s.ValidateCartOption(option, config.Cart)
		if err != nil {
			return nil, err
		}
	}
	validatedData, err := s.providerService_.validateFulfillmentData(option, data, config.Cart)
	if err != nil {
		return nil, err
	}
	var methodPrice float64
	if config.Price != 0.0 {
		methodPrice = config.Price
	} else {
		methodPrice, err = s.GetPrice(option, validatedData, config.Cart)
		if err != nil {
			return nil, err
		}
	}
	toCreate := &models.ShippingMethod{
		ShippingOptionId: uuid.NullUUID{UUID: option.Id},
		Data:             validatedData,
		Price:            methodPrice,
	}

	feature := true
	if feature {
		toCreate.IncludesTax = option.IncludesTax
	}
	if config.Order != nil {
		toCreate.OrderId = uuid.NullUUID{UUID: config.Order.Id}
	}
	if config.Cart != nil {
		toCreate.CartId = uuid.NullUUID{UUID: config.Cart.Id}
	}
	if config.CartId.UUID != uuid.Nil {
		toCreate.CartId = config.CartId
	}
	if config.ReturnId.UUID != uuid.Nil {
		toCreate.ReturnId = config.ReturnId
	}
	if config.OrderId.UUID != uuid.Nil {
		toCreate.OrderId = config.OrderId
	}
	if config.ClaimOrderId.UUID != uuid.Nil {
		toCreate.ClaimOrderId = config.ClaimOrderId
	}

	if err = s.methodRepository.Save(s.ctx, toCreate); err != nil {
		return nil, err
	}

	var method *models.ShippingMethod
	query := repository.BuildQuery(models.ShippingOptionRequirement{Model: core.Model{Id: toCreate.Id}}, repository.Options{
		Relations: []string{"shipping_option"},
	})
	if err := s.methodRepository.FindOne(s.ctx, method, query); err != nil {
		return nil, err
	}

	return method, nil
}

func (s *ShippingOptionService) ValidateCartOption(option *models.ShippingOption, cart *models.Cart) (*models.ShippingOption, *utils.ApplictaionError) {
	if option.IsReturn {
		return nil, nil
	}
	if cart.RegionId != option.RegionId {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"The shipping option is not available in the cart's region",
			"500",
			nil,
		)
	}
	var amount float64
	if option.IncludesTax {
		amount = cart.Total
	} else {
		amount = cart.Subtotal
	}
	requirementResults := []bool{}
	for _, requirement := range option.Requirements {
		switch requirement.Type {
		case "max_subtotal":
			requirementResults = append(requirementResults, requirement.Amount > amount)
		case "min_subtotal":
			requirementResults = append(requirementResults, requirement.Amount <= amount)
		default:
			requirementResults = append(requirementResults, true)
		}
	}
	if slices.Contains(requirementResults, true) {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"The Cart does not satisfy the shipping option's requirements",
			"500",
			nil,
		)
	}
	var err *utils.ApplictaionError
	option.Amount, err = s.GetPrice(option, option.Data, cart)
	if err != nil {
		return nil, err
	}
	return option, nil
}

func (s *ShippingOptionService) ValidateAndMutatePrice(option *models.ShippingOption, priceInput types.ValidatePriceTypeAndAmountInput) (*models.ShippingOption, *utils.ApplictaionError) {
	option.Amount = priceInput.Amount
	if priceInput.PriceType != nil {
		priceType, err := s.validatePriceType_(priceInput.PriceType, option)
		if err != nil {
			return nil, err
		}
		option.PriceType = priceType
		if *priceInput.PriceType == models.ShippingOptionPriceCalculated {
			option.Amount = 0.0
		}
	}
	if option.PriceType == models.ShippingOptionPriceFlatRate && option.Amount < 0.0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Shipping options of type `flat_rate` must have an `amount`",
			"500",
			nil,
		)
	}
	return option, nil
}

func (s *ShippingOptionService) Create(data *models.ShippingOption) (*models.ShippingOption, *utils.ApplictaionError) {
	optionWithValidatedPrice, err := s.ValidateAndMutatePrice(data, types.ValidatePriceTypeAndAmountInput{
		PriceType: &data.PriceType,
	})
	if err != nil {
		return nil, err
	}
	option := &models.ShippingOption{
		Name:        data.Name,
		ProfileId:   data.ProfileId,
		RegionId:    data.RegionId,
		ProviderId:  data.ProviderId,
		Data:        data.Data,
		PriceType:   optionWithValidatedPrice.PriceType,
		Amount:      optionWithValidatedPrice.Amount,
		IncludesTax: data.IncludesTax,
		IsReturn:    data.IsReturn,
	}

	if err = s.repo.Save(s.ctx, option); err != nil {
		return nil, err
	}
	return option, nil
}

func (s *ShippingOptionService) Update(optionId uuid.UUID, update *models.ShippingOption) (*models.ShippingOption, *utils.ApplictaionError) {
	option, err := s.Retrieve(optionId, repository.Options{
		Relations: []string{"requirements"},
	})

	if update.Metadata != nil {
		option.Metadata = update.Metadata
	}
	if update.RegionId.UUID != uuid.Nil || update.ProviderId.UUID != uuid.Nil || update.Data != nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Region and Provider cannot be updated after creation",
			"500",
			nil,
		)
	}
	if update.IsReturn {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"is_return cannot be changed after creation",
			"500",
			nil,
		)
	}
	if update.Requirements != nil {
		acc := []models.ShippingOptionRequirement{}
		for _, r := range update.Requirements {
			validated, err := s.ValidateRequirement(&r, optionId)
			if err != nil {
				return nil, err
			}
			if containsRequirement(acc, validated.Type) {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					fmt.Sprintf("Only one requirement of each type is allowed"),
					"500",
					nil,
				)
			}
			if containsRequirementWithAmount(acc, validated.Type, validated.Amount) {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					fmt.Sprintf("Max. subtotal must be greater than Min. subtotal"),
					"500",
					nil,
				)
			}
			acc = append(acc, *validated)
		}
		option.Requirements = acc
	}
	optionWithValidatedPrice, err := s.ValidateAndMutatePrice(option, types.ValidatePriceTypeAndAmountInput{
		PriceType: &update.PriceType,
		Amount:    update.Amount,
	})
	if err != nil {
		return nil, err
	}

	optionWithValidatedPrice.Name = update.Name
	optionWithValidatedPrice.AdminOnly = update.AdminOnly
	optionWithValidatedPrice.ProfileId = update.ProfileId

	feature := true
	if feature {
		optionWithValidatedPrice.IncludesTax = update.IncludesTax
	}
	if err = s.repo.Save(s.ctx, optionWithValidatedPrice); err != nil {
		return nil, err
	}
	return optionWithValidatedPrice, nil
}

func (s *ShippingOptionService) Delete(optionId uuid.UUID) (*models.ShippingOption, *utils.ApplictaionError) {
	data, err := s.Retrieve(optionId, repository.Options{})
	if err != nil {
		return nil, err
	}

	if err := s.repo.SoftRemove(s.ctx, data); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ShippingOptionService) AddRequirement(optionId uuid.UUID, requirement *models.ShippingOptionRequirement) (*models.ShippingOption, *utils.ApplictaionError) {
	option, err := s.Retrieve(optionId, repository.Options{
		Relations: []string{"requirements"},
	})
	if err != nil {
		return nil, err
	}
	validatedReq, err := s.ValidateRequirement(requirement, optionId)
	if err != nil {
		return nil, err
	}
	if slices.Contains(option.Requirements, validatedReq.Type) {
		return nil, utils.NewApplictaionError(
			utils.DUPLICATE_ERROR,
			fmt.Sprintf("A requirement with type: %s already exists", validatedReq.Type),
			"500",
			nil,
		)
	}
	option.Requirements = append(option.Requirements, *validatedReq)
	if err = s.repo.Save(s.ctx, option); err != nil {
		return nil, err
	}
	return option, nil
}

func (s *ShippingOptionService) RemoveRequirement(requirementId uuid.UUID) (*models.ShippingOptionRequirement, *utils.ApplictaionError) {
	var requirement *models.ShippingOptionRequirement
	query := repository.BuildQuery(models.ShippingOptionRequirement{Model: core.Model{Id: requirementId}}, repository.Options{})
	if err := s.requirementRepository.FindOne(s.ctx, requirement, query); err != nil {
		return nil, err
	}

	if err := s.requirementRepository.SoftRemove(s.ctx, requirement); err != nil {
		return nil, err
	}
	return requirement, nil
}

func (s *ShippingOptionService) UpdateShippingProfile(optionIds uuid.UUIDs, profileId uuid.UUID) ([]models.ShippingOption, *utils.ApplictaionError) {
	var res *models.ShippingOption
	var model []models.ShippingOption

	res.ProfileId = res.ProfileId

	if err := s.repo.Specification(repository.In[uuid.UUID]("id", optionIds)).Upsert(s.ctx, res); err != nil {
		return nil, err
	}

	if err := s.repo.Specification(repository.In[uuid.UUID]("id", optionIds)).Find(s.ctx, model, repository.Query{}); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *ShippingOptionService) GetPrice(option *models.ShippingOption, data core.JSONB, cart *models.Cart) (float64, *utils.ApplictaionError) {
	if option.PriceType == "calculated" {
		return s.providerService_.CalculatePrice(option, data, cart)
	}
	return option.Amount, nil
}
