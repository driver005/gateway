package services

import (
	"context"
	"fmt"
	"reflect"
	"slices"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type ShippingOptionService struct {
	ctx context.Context
	r   Registry
}

func NewShippingOptionService(
	r Registry,
) *ShippingOptionService {
	return &ShippingOptionService{
		context.Background(),
		r,
	}
}

func (s *ShippingOptionService) SetContext(context context.Context) *ShippingOptionService {
	s.ctx = context
	return s
}

func (s *ShippingOptionService) ValidateRequirement(data *types.ValidateRequirementTypeInput, optionId uuid.UUID) (*models.ShippingOptionRequirement, *utils.ApplictaionError) {
	if !reflect.ValueOf(data.Type).IsZero() {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"A Shipping Requirement must have a type field",
			nil,
		)
	}
	if data.Type != models.ShippingOptionRequirementMinSubtotal && data.Type != models.ShippingOptionRequirementMaxSubtotal {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Requirement type must be one of min_subtotal, max_subtotal",
			nil,
		)
	}

	var existingReq *models.ShippingOptionRequirement = &models.ShippingOptionRequirement{}
	if data.Id != uuid.Nil {
		query := sql.BuildQuery(models.ShippingOptionRequirement{Model: core.Model{Id: data.Id}}, &sql.Options{})
		err := s.r.ShippingOptionRequirementRepository().FindOne(s.ctx, existingReq, query)
		if err != nil {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				fmt.Sprintf("Shipping option requirement with id %s does not exist", data.Id),
				nil,
			)
		}
	}
	if existingReq == nil && data.Id != uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("Shipping option requirement with id %s does not exist", data.Id),
			nil,
		)
	}

	model := &models.ShippingOptionRequirement{
		Model: core.Model{
			Id: data.Id,
		},
		Type:   data.Type,
		Amount: data.Amount,
	}

	if optionId == uuid.Nil {
		return model, nil
	}
	if reflect.DeepEqual(existingReq, &models.ShippingOptionRequirement{}) {
		model = existingReq
		if !reflect.ValueOf(data.Type).IsZero() {
			model.Type = data.Type
		}
		if !reflect.ValueOf(data.Amount).IsZero() {
			model.Amount = data.Amount
		}
	}
	if err := s.r.ShippingOptionRequirementRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *ShippingOptionService) List(selector *types.FilterableShippingOption, config *sql.Options) ([]models.ShippingOption, *utils.ApplictaionError) {
	var res []models.ShippingOption
	query := sql.BuildQuery(selector, config)
	if err := s.r.ShippingOptionRepository().Find(s.ctx, &res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ShippingOptionService) ListAndCount(selector *types.FilterableShippingOption, config *sql.Options) ([]models.ShippingOption, *int64, *utils.ApplictaionError) {
	var res []models.ShippingOption

	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = 0
		config.Take = 50
	}

	query := sql.BuildQuery(selector, config)
	count, err := s.r.ShippingOptionRepository().FindAndCount(s.ctx, &res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *ShippingOptionService) Retrieve(optionId uuid.UUID, config *sql.Options) (*models.ShippingOption, *utils.ApplictaionError) {
	if optionId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			`"optionId" must be defined`,
			nil,
		)
	}
	var res *models.ShippingOption
	query := sql.BuildQuery(models.ShippingOption{Model: core.Model{Id: optionId}}, config)
	if err := s.r.ShippingOptionRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("Shipping Option with %s was not found", optionId),
			nil,
		)
	}

	return res, nil
}

func (s *ShippingOptionService) UpdateShippingMethod(id uuid.UUID, data *types.ShippingMethodUpdate) (*models.ShippingMethod, *utils.ApplictaionError) {
	model := &models.ShippingMethod{}

	query := sql.BuildQuery(models.ShippingMethod{Model: core.Model{Id: id}}, &sql.Options{})
	if err := s.r.ShippingMethodRepository().FindOne(s.ctx, model, query); err != nil {
		return nil, err
	}

	if !reflect.ValueOf(data.Data).IsZero() {
		model.Data = data.Data
	}
	if !reflect.ValueOf(data.Price).IsZero() {
		model.Price = data.Price
	}
	if !reflect.ValueOf(data.ReturnId).IsZero() {
		model.ReturnId = uuid.NullUUID{UUID: data.ReturnId}
	}
	if !reflect.ValueOf(data.SwapId).IsZero() {
		model.SwapId = uuid.NullUUID{UUID: data.SwapId}
	}
	if !reflect.ValueOf(data.OrderId).IsZero() {
		model.OrderId = uuid.NullUUID{UUID: data.OrderId}
	}
	if !reflect.ValueOf(data.ClaimOrderId).IsZero() {
		model.ClaimOrderId = uuid.NullUUID{UUID: data.ClaimOrderId}
	}

	if err := s.r.ShippingMethodRepository().Upsert(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *ShippingOptionService) DeleteShippingMethods(shippingMethods []models.ShippingMethod) *utils.ApplictaionError {
	err := s.r.ShippingMethodRepository().RemoveSlice(s.ctx, shippingMethods)
	if err != nil {
		return err
	}
	return nil
}

func (s *ShippingOptionService) CreateShippingMethod(optionId uuid.UUID, data map[string]interface{}, config *types.CreateShippingMethodDto) (*models.ShippingMethod, *utils.ApplictaionError) {
	option, err := s.Retrieve(optionId, &sql.Options{
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
	validatedData, err := s.r.FulfillmentProviderService().SetContext(s.ctx).ValidateFulfillmentData(option, data, config.Cart)
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
	if config.CartId != uuid.Nil {
		toCreate.CartId = uuid.NullUUID{UUID: config.CartId}
	}
	if config.ReturnId != uuid.Nil {
		toCreate.ReturnId = uuid.NullUUID{UUID: config.ReturnId}
	}
	if config.OrderId != uuid.Nil {
		toCreate.OrderId = uuid.NullUUID{UUID: config.OrderId}
	}
	if config.ClaimOrderId != uuid.Nil {
		toCreate.ClaimOrderId = uuid.NullUUID{UUID: config.ClaimOrderId}
	}

	if err = s.r.ShippingMethodRepository().Save(s.ctx, toCreate); err != nil {
		return nil, err
	}

	var method *models.ShippingMethod
	query := sql.BuildQuery(models.ShippingOptionRequirement{Model: core.Model{Id: toCreate.Id}}, &sql.Options{
		Relations: []string{"shipping_option"},
	})
	if err := s.r.ShippingMethodRepository().FindOne(s.ctx, method, query); err != nil {
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

func (s *ShippingOptionService) ValidateAndMutatePrice(option *models.ShippingOption, option2 *types.CreateShippingOptionInput, priceInput types.ValidatePriceTypeAndAmountInput) (*models.ShippingOption, *utils.ApplictaionError) {
	if option2 != nil {
		var requirements []models.ShippingOptionRequirement
		for _, requirement := range option2.Requirements {
			requirements = append(requirements, models.ShippingOptionRequirement{
				Amount: requirement.Amount,
				Type:   requirement.Type,
			})
		}
		option.PriceType = option2.PriceType
		option.Name = option2.Name
		option.RegionId = uuid.NullUUID{UUID: option2.RegionId}
		option.ProfileId = uuid.NullUUID{UUID: option2.ProfileId}
		option.ProviderId = uuid.NullUUID{UUID: option2.ProviderId}
		option.Data = option2.Data
		option.IncludesTax = option2.IncludesTax
		option.Amount = option2.Amount
		option.AdminOnly = option2.AdminOnly
		option.IsReturn = option2.IsReturn
		option.Metadata = option2.Metadata
		option.Requirements = requirements
	}

	option.Amount = priceInput.Amount
	if priceInput.PriceType != "" {
		priceType, err := s.validatePriceType(priceInput.PriceType, option)
		if err != nil {
			return nil, err
		}
		option.PriceType = priceType
		if priceInput.PriceType == models.ShippingOptionPriceCalculated {
			option.Amount = 0.0
		}
	}
	if option.PriceType == models.ShippingOptionPriceFlatRate && option.Amount < 0.0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Shipping options of type `flat_rate` must have an `amount`",
			nil,
		)
	}
	return option, nil
}

// Validates a shipping option price
func (s *ShippingOptionService) validatePriceType(priceType models.ShippingOptionPriceType, option *models.ShippingOption) (models.ShippingOptionPriceType, *utils.ApplictaionError) {
	if priceType == "" || (priceType != models.ShippingOptionPriceFlatRate && priceType != models.ShippingOptionPriceCalculated) {
		return "", utils.NewApplictaionError(
			utils.INVALID_DATA,
			"The price must be of type flat_rate or calculated",
			nil,
		)
	}
	if priceType == models.ShippingOptionPriceCalculated {
		canCalculate, err := s.r.FulfillmentProviderService().SetContext(s.ctx).CanCalculate(types.FulfillmentOptions{
			ProviderId: option.Provider.Id,
			Options:    option.Data,
		})
		if err != nil {
			return "", err
		}
		if !canCalculate {
			return "", utils.NewApplictaionError(
				utils.CONFLICT,
				"The fulfillment provider cannot calculate prices for this option",
				"500",
				nil,
			)
		}
	}
	return priceType, nil
}

func (s *ShippingOptionService) Create(data *types.CreateShippingOptionInput) (*models.ShippingOption, *utils.ApplictaionError) {
	optionWithValidatedPrice, err := s.ValidateAndMutatePrice(nil, data, types.ValidatePriceTypeAndAmountInput{
		PriceType: data.PriceType,
	})
	if err != nil {
		return nil, err
	}
	option := &models.ShippingOption{
		Name:        data.Name,
		ProfileId:   uuid.NullUUID{UUID: data.ProfileId},
		RegionId:    uuid.NullUUID{UUID: data.RegionId},
		ProviderId:  uuid.NullUUID{UUID: data.ProviderId},
		Data:        data.Data,
		PriceType:   optionWithValidatedPrice.PriceType,
		Amount:      optionWithValidatedPrice.Amount,
		IncludesTax: data.IncludesTax,
		IsReturn:    data.IsReturn,
	}

	if err = s.r.ShippingOptionRepository().Save(s.ctx, option); err != nil {
		return nil, err
	}
	return option, nil
}

func (s *ShippingOptionService) Update(optionId uuid.UUID, data *types.UpdateShippingOptionInput) (*models.ShippingOption, *utils.ApplictaionError) {
	option, err := s.Retrieve(optionId, &sql.Options{
		Relations: []string{"requirements"},
	})
	if err != nil {
		return nil, err
	}

	if data.Metadata != nil {
		option.Metadata = utils.MergeMaps(option.Metadata, data.Metadata)
	}
	if data.RegionId != uuid.Nil || data.ProviderId != uuid.Nil || reflect.ValueOf(data.Data).IsZero() {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"Region and Provider cannot be updated after creation",
			nil,
		)
	}
	if !reflect.ValueOf(data.IsReturn).IsZero() {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"is_return cannot be changed after creation",
			nil,
		)
	}
	if data.Requirements != nil {
		acc := []models.ShippingOptionRequirement{}
		for _, r := range data.Requirements {
			validated, err := s.ValidateRequirement(&types.ValidateRequirementTypeInput{
				Amount: r.Amount,
				Type:   r.Type,
			}, optionId)
			if err != nil {
				return nil, err
			}
			if slices.ContainsFunc(option.Requirements, func(m models.ShippingOptionRequirement) bool {
				return m.Type == validated.Type
			}) {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					"Only one requirement of each type is allowed",
					"500",
					nil,
				)
			}
			if slices.ContainsFunc(option.Requirements, func(m models.ShippingOptionRequirement) bool {
				return m.Type == models.ShippingOptionRequirementMaxSubtotal && validated.Amount > m.Amount || m.Type == models.ShippingOptionRequirementMinSubtotal && validated.Amount < m.Amount
			}) {
				return nil, utils.NewApplictaionError(
					utils.INVALID_DATA,
					"Max. subtotal must be greater than Min. subtotal",
					"500",
					nil,
				)
			}
			acc = append(acc, *validated)
		}
		option.Requirements = acc
	}
	optionWithValidatedPrice, err := s.ValidateAndMutatePrice(option, nil, types.ValidatePriceTypeAndAmountInput{
		PriceType: data.PriceType,
		Amount:    data.Amount,
	})
	if err != nil {
		return nil, err
	}

	if !reflect.ValueOf(data.Name).IsZero() {
		optionWithValidatedPrice.Name = data.Name
	}
	if !reflect.ValueOf(data.AdminOnly).IsZero() {
		optionWithValidatedPrice.AdminOnly = data.AdminOnly
	}
	if !reflect.ValueOf(data.ProfileId).IsZero() {
		optionWithValidatedPrice.ProfileId = uuid.NullUUID{UUID: data.ProfileId}
	}

	feature := true
	if feature {
		if !reflect.ValueOf(data.IncludesTax).IsZero() {
			optionWithValidatedPrice.IncludesTax = data.IncludesTax
		}
	}
	if err = s.r.ShippingOptionRepository().Save(s.ctx, optionWithValidatedPrice); err != nil {
		return nil, err
	}
	return optionWithValidatedPrice, nil
}

func (s *ShippingOptionService) Delete(optionId uuid.UUID) *utils.ApplictaionError {
	data, err := s.Retrieve(optionId, &sql.Options{})
	if err != nil {
		return err
	}

	if err := s.r.ShippingOptionRepository().SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}

func (s *ShippingOptionService) AddRequirement(optionId uuid.UUID, requirement *models.ShippingOptionRequirement) (*models.ShippingOption, *utils.ApplictaionError) {
	option, err := s.Retrieve(optionId, &sql.Options{
		Relations: []string{"requirements"},
	})
	if err != nil {
		return nil, err
	}
	validatedReq, err := s.ValidateRequirement(&types.ValidateRequirementTypeInput{
		Id:     requirement.Id,
		Type:   requirement.Type,
		Amount: requirement.Amount,
	}, optionId)
	if err != nil {
		return nil, err
	}
	if slices.ContainsFunc(option.Requirements, func(m models.ShippingOptionRequirement) bool {
		return m.Type == validatedReq.Type
	}) {
		return nil, utils.NewApplictaionError(
			utils.DUPLICATE_ERROR,
			fmt.Sprintf("A requirement with type: %s already exists", validatedReq.Type),
			nil,
		)
	}
	option.Requirements = append(option.Requirements, *validatedReq)
	if err = s.r.ShippingOptionRepository().Save(s.ctx, option); err != nil {
		return nil, err
	}
	return option, nil
}

func (s *ShippingOptionService) RemoveRequirement(requirementId uuid.UUID) (*models.ShippingOptionRequirement, *utils.ApplictaionError) {
	var requirement *models.ShippingOptionRequirement = &models.ShippingOptionRequirement{}
	query := sql.BuildQuery(models.ShippingOptionRequirement{Model: core.Model{Id: requirementId}}, &sql.Options{})
	if err := s.r.ShippingOptionRequirementRepository().FindOne(s.ctx, requirement, query); err != nil {
		return nil, err
	}

	if err := s.r.ShippingOptionRequirementRepository().SoftRemove(s.ctx, requirement); err != nil {
		return nil, err
	}
	return requirement, nil
}

func (s *ShippingOptionService) UpdateShippingProfile(optionIds uuid.UUIDs, profileId uuid.UUID) ([]models.ShippingOption, *utils.ApplictaionError) {
	var res *models.ShippingOption
	var model []models.ShippingOption

	res.ProfileId = uuid.NullUUID{UUID: profileId}

	if err := s.r.ShippingOptionRepository().Specification(sql.In[uuid.UUID]("id", optionIds)).Upsert(s.ctx, res); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *ShippingOptionService) GetPrice(option *models.ShippingOption, data core.JSONB, cart *models.Cart) (float64, *utils.ApplictaionError) {
	if option.PriceType == "calculated" {
		return s.r.FulfillmentProviderService().SetContext(s.ctx).CalculatePrice(option, data, cart)
	}
	return option.Amount, nil
}
