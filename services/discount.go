package services

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/repository"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
)

type DiscountService struct {
	ctx                         context.Context
	repo                        *repository.DiscountRepo
	discountRuleRepository      *repository.DiscountRuleRepo
	giftCardRepository          *repository.GiftCardRepo
	discountConditionRepository *repository.DiscountConditionRepo
	regionService               *RegionService
}

func NewDiscountService(
	ctx context.Context,
	repo *repository.DiscountRepo,
	discountRuleRepository *repository.DiscountRuleRepo,
	giftCardRepository *repository.GiftCardRepo,
	discountConditionRepository *repository.DiscountConditionRepo,
	regionService *RegionService,
) *DiscountService {
	return &DiscountService{
		ctx,
		repo,
		discountRuleRepository,
		giftCardRepository,
		discountConditionRepository,
		regionService,
	}
}

func (s *DiscountService) validateDiscountRule(discountRule *models.DiscountRule) (*models.DiscountRule, *utils.ApplictaionError) {
	if discountRule.Type == "percentage" && discountRule.Value > 100 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"models.Discount value above 100 is not allowed when type is percentage",
			"500",
			nil,
		)
	}
	return discountRule, nil
}

func (s *DiscountService) list(selector types.FilterableDiscount, config repository.Options) ([]models.Discount, error) {
	var discounts []models.Discount
	query := repository.BuildQuery(selector, config)
	if err := s.repo.Find(s.ctx, discounts, query); err != nil {
		return nil, err
	}
	return discounts, nil
}

func (s *DiscountService) listAndCount(selector types.FilterableDiscount, config repository.Options) ([]models.Discount, *int64, error) {
	var discounts []models.Discount
	query := repository.BuildQuery(selector, config)
	count, err := s.repo.FindAndCount(s.ctx, discounts, query)
	if err != nil {
		return nil, nil, err
	}
	return discounts, count, nil
}

func (s *DiscountService) create(discount *models.Discount) (*models.Discount, *utils.ApplictaionError) {
	conditions := discount.Rule.Conditions
	ruleToCreate := discount.Rule
	ruleToCreate.Conditions = nil
	validatedRule, err := s.validateDiscountRule(ruleToCreate)
	if err != nil {
		return nil, err
	}
	if len(discount.Regions) > 1 && discount.Rule.Type == "fixed" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Fixed discounts can have one region",
			"500",
			nil,
		)
	}
	if discount.Regions != nil {
		discount.Regions = make([]models.Region, len(discount.Regions))
		for i, regionId := range discount.Regions {
			discount.Regions[i] = s.regionService.Retrieve(regionId)
		}
	}
	if len(discount.Regions) == 0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"models.Discount must have atleast 1 region",
			"500",
			nil,
		)
	}
	if err := s.discountRuleRepository.Save(s.ctx, validatedRule); err != nil {
		return nil, err
	}
	discount.Rule = validatedRule
	if err := s.repo.Save(s.ctx, discount); err != nil {
		return nil, err
	}
	if len(conditions) > 0 {
		for _, cond := range conditions {
			err := s.discountConditionService_.upsertCondition(map[string]interface{}{
				"rule_id": result.RuleId,
				"cond":    cond,
			})
			if err != nil {
				return nil, err
			}
		}
	}
	err = s.eventBus_.emit(DiscountService.Events.CREATED, map[string]interface{}{
		"id": result.Id,
	})
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (s *DiscountService) retrieve(discountId string, config repository.Options) (*models.Discount, error) {
	if discountId == "" {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			`"discountId" must be defined`,
			"500",
			nil,
		)
	}
	query := buildQuery(map[string]interface{}{"id": discountId}, config)
	discount, err := s.repo.findOne(query)
	if err != nil {
		return nil, err
	}
	if discount == nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("models.Discount with id %s was not found", discountId),
			"500",
			nil,
		)
	}
	return discount, nil
}

func (s *DiscountService) retrieveByCode(discountCode string, config repository.Options) (*models.Discount, error) {
	normalizedCode := strings.ToUpper(strings.TrimSpace(discountCode))
	query := buildQuery(map[string]interface{}{"code": normalizedCode}, config)
	discount, err := s.repo.findOne(query)
	if err != nil {
		return nil, err
	}
	if discount == nil {
		return models.Discount{}, utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("Discounts with code %s was not found", discountCode),
			"500",
			nil,
		)
	}
	return discount, nil
}

func (s *DiscountService) listByCodes(discountCodes []string, config repository.Options) ([]models.Discount, error) {
	normalizedCodes := make([]string, len(discountCodes))
	for i, code := range discountCodes {
		normalizedCodes[i] = strings.ToUpper(strings.TrimSpace(code))
	}
	query := buildQuery(map[string]interface{}{"code": In(normalizedCodes)}, config)
	discounts, err := s.repo.Find(query)
	if err != nil {
		return nil, err
	}
	if len(discounts) != len(discountCodes) {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("Discounts with code [%s] was not found", strings.Join(normalizedCodes, ", ")),
			"500",
			nil,
		)
	}
	return discounts, nil
}

func (s *DiscountService) update(discountId string, update UpdateDiscountInput) (*models.Discount, error) {
	s.discountRuleRepository := manager.withRepository(s.discountRuleRepository_)
	discount, err := s.retrieve(discountId, repository.Options{Relations: []string{"rule"}})
	if err != nil {
		return nil, err
	}
	conditions := update.Rule.Conditions
	ruleToUpdate := omit(update.Rule, "conditions")
	if !isEmpty(ruleToUpdate) {
		update.Rule = ruleToUpdate.(UpdateDiscountRuleInput)
	}
	rule := update.Rule
	metadata := update.Metadata
	regions := update.Regions
	rest := omit(update, "rule", "metadata", "regions")
	if rest.EndsAt != nil {
		if discount.StartsAt >= *rest.EndsAt {
			return models.Discount{}, utils.NewApplictaionError(
				utils.INVALID_DATA,
				`"ends_at" must be greater than "starts_at"`,
				"500",
				nil,
			)
		}
	}
	if len(regions) > 1 && discount.Rule.Type == "fixed" {
		return models.Discount{}, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Fixed discounts can have one region",
			"500",
			nil,
		)
	}
	if len(conditions) > 0 {
		for _, cond := range conditions {
			err := s.discountConditionService_.upsertCondition(map[string]interface{}{
				"rule_id": discount.RuleId,
				"cond":    cond,
			})
			if err != nil {
				return nil, err
			}
		}
	}
	if regions != nil {
		discount.Regions = make([]Region, len(regions))
		for i, regionId := range regions {
			discount.Regions[i] = s.regionService_.retrieve(regionId)
		}
	}
	if metadata != nil {
		discount.Metadata = setMetadata(discount, metadata)
	}
	if rule != nil {
		ruleUpdate := rule.(UpdateDiscountRuleInput)
		if rule.Value != nil {
			s.validateDiscountRule(*models.DiscountRule{
				Value: *rule.Value,
				discount.Rule.Type,
			})
		}
		// discount.Rule = s.discountRuleRepository.create(*models.DiscountRule{
		// 	...discount.Rule,
		// 	...ruleUpdate,
		// })
	}
	for key, value := range rest {
		if value != "" {
			discount[key] = value
		}
	}
	discount.Code = strings.ToUpper(discount.Code)
	result, err := s.repo.Save(discount)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *DiscountService) createDynamicCode(discountId string, data CreateDynamicDiscountInput) (*models.Discount, error) {
	discount, err := s.retrieve(discountId)
	if err != nil {
		return nil, err
	}
	if !discount.IsDynamic {
		return models.Discount{}, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"models.Discount must be set to dynamic",
			"500",
			nil,
		)
	}
	if data.Code == "" {
		return models.Discount{}, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"models.Discount must have a code",
			"500",
			nil,
		)
	}
	toCreate := CreateDynamicDiscountInput{
		// ...data,
		RuleId:           discount.RuleId,
		IsDynamic:        true,
		IsDisabled:       false,
		Code:             strings.ToUpper(data.Code),
		ParentDiscountId: discount.Id,
		UsageLimit:       discount.UsageLimit,
	}
	if discount.ValidDuration != nil {
		lastValidDate := time.Now()
		lastValidDate.Add(time.Second * time.Duration(toSeconds(parse(discount.ValidDuration))))
		toCreate.EndsAt = &lastValidDate
	}
	result, err := s.repo.Save(toCreate)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (s *DiscountService) deleteDynamicCode(discountId string, code string) error {
	s.repo := manager.withRepository(s.discountRepository_)
	discount, err := s.repo.findOne(map[string]interface{}{
		"parent_discount_id": discountId,
		"code":               code,
	})
	if err != nil {
		return err
	}
	if discount == nil {
		return nil
	}
	return s.repo.softRemove(discount)
}

func (s *DiscountService) addRegion(discountId string, regionId string) (*models.Discount, error) {
	discount, err := s.retrieve(discountId, repository.Options{Relations: []string{"regions", "rule"}})
	if err != nil {
		return nil, err
	}
	exists := false
	for _, r := range discount.Regions {
		if r.Id == regionId {
			exists = true
			break
		}
	}
	if exists {
		return discount, nil
	}
	if len(discount.Regions) == 1 && discount.Rule.Type == "fixed" {
		return models.Discount{}, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Fixed discounts can have one region",
			"500",
			nil,
		)
	}
	region := s.regionService_.retrieve(regionId)
	discount.Regions = append(discount.Regions, region)
	result, err := s.repo.Save(discount)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (s *DiscountService) removeRegion(discountId string, regionId string) (*models.Discount, error) {
	discount, err := s.retrieve(discountId, repository.Options{Relations: []string{"regions"}})
	if err != nil {
		return nil, err
	}
	exists := false
	for _, r := range discount.Regions {
		if r.Id == regionId {
			exists = true
			break
		}
	}
	if !exists {
		return discount, nil
	}
	discount.Regions = discount.Regions[:0]
	for _, r := range discount.Regions {
		if r.Id != regionId {
			discount.Regions = append(discount.Regions, r)
		}
	}
	result, err := s.repo.Save(discount)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *DiscountService) delete(discountId string) error {
	discount, err := s.repo.findOne(map[string]interface{}{"id": discountId})
	if err != nil {
		return err
	}
	if discount == nil {
		return nil
	}
	return s.repo.softRemove(discount)
}

func (s *DiscountService) validateDiscountForProduct(discountRuleId string, productId string) (bool, error) {
	if productId == "" {
		return false, nil
	}
	return discountConditionRepo.isValidForProduct(discountRuleId, productId)
}

func (s *DiscountService) calculateDiscountForLineItem(discountId string, lineItem LineItem, calculationContextData CalculationContextData) (int, error) {
	adjustment := 0
	if !lineItem.AllowDiscounts {
		return adjustment, nil
	}
	discount, err := s.retrieve(discountId, repository.Options{Relations: []string{"rule"}})
	if err != nil {
		return 0, err
	}
	discountType := discount.Rule.Type
	discountValue := discount.Rule.Value
	discountAllocation := discount.Rule.Allocation
	calculationContext, err := s.totalsService_.withTransaction(transactionManager).getCalculationContext(calculationContextData, CalculationContextConfig{ExcludeShipping: true})
	if err != nil {
		return 0, err
	}
	fullItemPrice := lineItem.UnitPrice * lineItem.Quantity
	includesTax := s.featureFlagRouter_.isFeatureEnabled(TaxInclusivePricingFeatureFlag.key) && lineItem.IncludesTax
	if includesTax {
		lineItemTotals, err := s.newTotalsService_.withTransaction(transactionManager).getLineItemTotals([]LineItem{lineItem}, LineItemTotalsConfig{IncludeTax: true, CalculationContext: calculationContext})
		if err != nil {
			return 0, err
		}
		fullItemPrice = lineItemTotals[lineItem.Id].Subtotal
	}
	if discountType == DiscountRuleType.PERCENTAGE {
		adjustment = int(math.Round(float64(fullItemPrice) / 100 * discountValue))
	} else if discountType == DiscountRuleType.FIXED && discountAllocation == DiscountAllocation.TOTAL {
		discountedItems := make([]LineItem, 0, len(calculationContextData.Items))
		for _, item := range calculationContextData.Items {
			if item.AllowDiscounts {
				discountedItems = append(discountedItems, item)
			}
		}
		totals, err := s.newTotalsService_.getLineItemTotals(discountedItems, LineItemTotalsConfig{IncludeTax: includesTax, CalculationContext: calculationContext})
		if err != nil {
			return 0, err
		}
		subtotal := 0
		for _, total := range totals {
			subtotal += total.Subtotal
		}
		nominator := math.Min(float64(discountValue), float64(subtotal))
		totalItemPercentage := float64(fullItemPrice) / float64(subtotal)
		adjustment = int(nominator * totalItemPercentage)
	} else {
		adjustment = discountValue * lineItem.Quantity
	}
	return int(math.Min(float64(adjustment), float64(fullItemPrice))), nil

}

func (s *DiscountService) validateDiscountForCartOrThrow(cart Cart, discount models.Discount) error {
	discounts := []models.Discount{discount}
	for _, disc := range discounts {
		if s.hasReachedLimit(disc) {
			return utils.NewApplictaionError(
				utils.NOT_ALLOWED,
				fmt.Sprintf("models.Discount %s has been used maximum allowed times", disc.Code),
				"500",
				nil,
			)
		}
		if s.hasNotStarted(disc) {
			return utils.NewApplictaionError(
				utils.NOT_ALLOWED,
				fmt.Sprintf("models.Discount %s is not valid yet", disc.Code),
				"500",
				nil,
			)
		}
		if s.hasExpired(disc) {
			return utils.NewApplictaionError(
				utils.NOT_ALLOWED,
				fmt.Sprintf("models.Discount %s is expired", disc.Code),
				"500",
				nil,
			)
		}
		if s.isDisabled(disc) {
			return utils.NewApplictaionError(
				utils.NOT_ALLOWED,
				fmt.Sprintf("The discount code %s is disabled", disc.Code),
				"500",
				nil,
			)
		}
		if cart.CustomerId == "" && s.hasCustomersGroupCondition(disc) {
			return utils.NewApplictaionError(
				utils.NOT_ALLOWED,
				fmt.Sprintf("models.Discount %s is only valid for specific customer", disc.Code),
				"500",
				nil,
			)
		}
		isValidForRegion, err := s.isValidForRegion(disc, cart.RegionId)
		if err != nil {
			return err
		}
		if !isValidForRegion {
			return utils.NewApplictaionError(
				utils.INVALID_DATA,
				"The discount is not available in current region",
				"500",
				nil,
			)
		}
		if cart.CustomerId != "" {
			canApplyForCustomer, err := s.canApplyForCustomer(disc.Rule.Id, cart.CustomerId)
			if err != nil {
				return err
			}
			if !canApplyForCustomer {
				return utils.NewApplictaionError(
					utils.NOT_ALLOWED,
					fmt.Sprintf("models.Discount %s is not valid for customer", disc.Code),
					"500",
					nil,
				)
			}
		}
	}
	return nil

}

func (s *DiscountService) hasCustomersGroupCondition(discount models.Discount) bool {
	for _, cond := range discount.Rule.Conditions {
		if cond.Type == DiscountConditionType.CUSTOMER_GROUPS {
			return true
		}
	}
	return false
}

func (s *DiscountService) hasReachedLimit(discount models.Discount) bool {
	count := discount.UsageCount
	limit := discount.UsageLimit
	return limit != nil && count >= *limit
}

func (s *DiscountService) hasNotStarted(discount models.Discount) bool {
	return discount.StartsAt.After(time.Now())
}

func (s *DiscountService) hasExpired(discount models.Discount) bool {
	if discount.EndsAt == nil {
		return false
	}
	return discount.EndsAt.Before(time.Now())
}

func (s *DiscountService) isDisabled(discount models.Discount) bool {
	return discount.IsDisabled
}

func (s *DiscountService) isValidForRegion(discount models.Discount, region_id string) bool {
	regions := discount.Regions
	if discount.ParentDiscountId != "" {
		parent, _ := retrieve(discount.ParentDiscountId, map[string]interface{}{
			"relations": []string{"rule", "regions"},
		})
		regions = parent.Regions
	}
	for _, r := range regions {
		if r.Id == region_id {
			return true
		}
	}
	return false
}

func (s *DiscountService) canApplyForCustomer(discountRuleId string, customerId string) bool {

	if customerId == "" {
		return false
	}
	customer, _ := customerService_.WithTransaction(manager).Retrieve(customerId, map[string]interface{}{
		"relations": []string{"groups"},
	})
	return discountConditionRepo.CanApplyForCustomer(discountRuleId, customer.Id)

}
