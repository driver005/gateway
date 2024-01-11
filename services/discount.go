package services

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type DiscountService struct {
	ctx context.Context
	r   Registry
}

func NewDiscountService(
	r Registry,
) *DiscountService {
	return &DiscountService{
		context.Background(),
		r,
	}
}

func (s *DiscountService) SetContext(context context.Context) *DiscountService {
	s.ctx = context
	return s
}

func (s *DiscountService) ValidateDiscountRule(discountRule *models.DiscountRule) (*models.DiscountRule, *utils.ApplictaionError) {
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

func (s *DiscountService) List(selector types.FilterableDiscount, config sql.Options) ([]models.Discount, *utils.ApplictaionError) {
	var discounts []models.Discount
	query := sql.BuildQuery(selector, config)
	if err := s.r.DiscountRepository().Find(s.ctx, discounts, query); err != nil {
		return nil, err
	}
	return discounts, nil
}

func (s *DiscountService) ListAndCount(selector types.FilterableDiscount, config sql.Options) ([]models.Discount, *int64, *utils.ApplictaionError) {
	var discounts []models.Discount
	query := sql.BuildQuery(selector, config)
	count, err := s.r.DiscountRepository().FindAndCount(s.ctx, discounts, query)
	if err != nil {
		return nil, nil, err
	}
	return discounts, count, nil
}

func (s *DiscountService) Create(discount *models.Discount) (*models.Discount, *utils.ApplictaionError) {
	conditions := discount.Rule.Conditions
	ruleToCreate := discount.Rule
	ruleToCreate.Conditions = nil
	validatedRule, err := s.ValidateDiscountRule(ruleToCreate)
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
		for i, region := range discount.Regions {
			res, err := s.r.RegionService().SetContext(s.ctx).Retrieve(region.Id, sql.Options{})
			if err != nil {
				return nil, err
			}
			discount.Regions[i] = *res
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
	if err := s.r.DiscountRuleRepository().Save(s.ctx, validatedRule); err != nil {
		return nil, err
	}
	discount.Rule = validatedRule
	if err := s.r.DiscountRepository().Save(s.ctx, discount); err != nil {
		return nil, err
	}
	if len(conditions) > 0 {
		for _, cond := range conditions {
			cond.DiscountRuleId = discount.RuleId
			_, err := s.r.DiscountConditionService().SetContext(s.ctx).UpsertCondition(&cond, false)

			if err != nil {
				return nil, err
			}
		}
	}
	// err = s.eventBus_.emit(DiscountService.Events.CREATED, map[string]interface{}{
	// 	"id": result.Id,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	return discount, nil

}

func (s *DiscountService) Retrieve(discountId uuid.UUID, config sql.Options) (*models.Discount, *utils.ApplictaionError) {
	if discountId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			`"discountId" must be defined`,
			"500",
			nil,
		)
	}
	var discount *models.Discount
	query := sql.BuildQuery(models.Discount{Model: core.Model{Id: discountId}}, config)
	if err := s.r.DiscountRepository().FindOne(s.ctx, discount, query); err != nil {
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

func (s *DiscountService) RetrieveByCode(discountCode string, config sql.Options) (*models.Discount, *utils.ApplictaionError) {
	var discount *models.Discount

	query := sql.BuildQuery(models.Discount{Code: strings.ToUpper(strings.TrimSpace(discountCode))}, config)
	if err := s.r.DiscountRepository().FindOne(s.ctx, discount, query); err != nil {
		return nil, err
	}
	if discount == nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("Discounts with code %s was not found", discountCode),
			"500",
			nil,
		)
	}
	return discount, nil
}

func (s *DiscountService) ListByCodes(discountCodes []string, config sql.Query) ([]models.Discount, *utils.ApplictaionError) {
	normalizedCodes := make([]string, len(discountCodes))
	for i, code := range discountCodes {
		normalizedCodes[i] = strings.ToUpper(strings.TrimSpace(code))
	}
	var discounts []models.Discount

	if err := s.r.DiscountRepository().Specification(sql.In("code", normalizedCodes)).Find(s.ctx, discounts, config); err != nil {
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

func (s *DiscountService) Update(discountId uuid.UUID, Update *models.Discount) (*models.Discount, *utils.ApplictaionError) {
	if discountId == uuid.Nil {
		Update.Id = discountId
	}
	if err := s.r.DiscountRepository().FindOne(s.ctx, Update, sql.Query{Relations: []string{"rule"}}); err != nil {
		return nil, err
	}
	conditions := Update.Rule.Conditions
	Update.Rule.Conditions = nil
	rule := Update.Rule
	regions := Update.Regions
	if Update.StartsAt.Before(*Update.EndsAt) {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"ends_at" must be greater than "starts_at"`,
			"500",
			nil,
		)
	}

	if len(regions) > 1 && Update.Rule.Type == "fixed" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Fixed discounts can have one region",
			"500",
			nil,
		)
	}
	if len(conditions) > 0 {
		for _, cond := range conditions {
			cond.DiscountRuleId = Update.RuleId
			_, err := s.r.DiscountConditionService().SetContext(s.ctx).UpsertCondition(&cond, false)
			if err != nil {
				return nil, err
			}
		}
	}
	if regions != nil {
		Update.Regions = make([]models.Region, len(regions))
		for i, region := range regions {
			res, err := s.r.RegionService().SetContext(s.ctx).Retrieve(region.Id, sql.Options{})
			if err != nil {
				return nil, err
			}
			Update.Regions[i] = *res
		}
	}

	if rule != nil {
		if rule.Value != 0.0 {
			s.ValidateDiscountRule(&models.DiscountRule{
				Value: rule.Value,
				Type:  Update.Rule.Type,
			})
		}
		// discount.Rule = s.r.DiscountRuleRepository().Create(*models.DiscountRule{
		// 	...discount.Rule,
		// 	...ruleUpdate,
		// })
	}

	Update.Code = strings.ToUpper(Update.Code)
	if err := s.r.DiscountRepository().Upsert(s.ctx, Update); err != nil {
		return nil, err
	}
	return Update, nil
}

func (s *DiscountService) CreateDynamicCode(discountId uuid.UUID, data *models.Discount) (*models.Discount, *utils.ApplictaionError) {
	discount, err := s.Retrieve(discountId, sql.Options{})
	if err != nil {
		return nil, err
	}
	if !discount.IsDynamic {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"models.Discount must be set to dynamic",
			"500",
			nil,
		)
	}
	if data.Code == "" {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"models.Discount must have a code",
			"500",
			nil,
		)
	}
	toCreate := &models.Discount{
		// ...data,
		RuleId:           discount.RuleId,
		IsDynamic:        true,
		IsDisabled:       false,
		Code:             strings.ToUpper(data.Code),
		ParentDiscountId: uuid.NullUUID{UUID: discount.Id},
		UsageLimit:       discount.UsageLimit,
	}
	if discount.ValidDuration != nil {
		lastValidDate := time.Now()
		lastValidDate = lastValidDate.Add(time.Second * time.Duration(discount.ValidDuration.Second()))
		toCreate.EndsAt = &lastValidDate
	}
	if err := s.r.DiscountRepository().Save(s.ctx, toCreate); err != nil {
		return nil, err
	}
	return toCreate, nil

}

func (s *DiscountService) DeleteDynamicCode(discountId uuid.UUID, code string) *utils.ApplictaionError {
	var discount *models.Discount
	query := sql.BuildQuery(models.Discount{ParentDiscountId: uuid.NullUUID{UUID: discountId}, Code: code}, sql.Options{})
	if err := s.r.DiscountRepository().FindOne(s.ctx, discount, query); err != nil {
		return err
	}
	if discount == nil {
		return nil
	}
	return s.r.DiscountRepository().SoftRemove(s.ctx, discount)
}

func (s *DiscountService) AddRegion(discountId uuid.UUID, regionId uuid.UUID) (*models.Discount, *utils.ApplictaionError) {
	discount, err := s.Retrieve(discountId, sql.Options{
		Relations: []string{"regions", "rule"},
	})
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
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Fixed discounts can have one region",
			"500",
			nil,
		)
	}
	region, err := s.r.RegionService().SetContext(s.ctx).Retrieve(regionId, sql.Options{})
	if err != nil {
		return nil, err
	}
	discount.Regions = append(discount.Regions, *region)
	if err := s.r.DiscountRepository().Save(s.ctx, discount); err != nil {
		return nil, err
	}
	return discount, nil

}

func (s *DiscountService) RemoveRegion(discountId uuid.UUID, regionId uuid.UUID) (*models.Discount, *utils.ApplictaionError) {
	discount, err := s.Retrieve(discountId, sql.Options{
		Relations: []string{"regions"},
	})
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
	if err := s.r.DiscountRepository().Save(s.ctx, discount); err != nil {
		return nil, err
	}
	return discount, nil
}

func (s *DiscountService) Delete(discountId uuid.UUID) *utils.ApplictaionError {
	data, err := s.Retrieve(discountId, sql.Options{})
	if err != nil {
		return err
	}

	if err := s.r.DiscountRepository().SoftRemove(s.ctx, data); err != nil {
		return err
	}

	return nil
}

func (s *DiscountService) ValidateDiscountForProduct(discountRuleId uuid.UUID, productId uuid.UUID) (bool, *utils.ApplictaionError) {
	if productId == uuid.Nil {
		return false, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"`productId` can not be undifined",
			"500",
			nil,
		)
	}
	return s.r.DiscountConditionRepository().IsValidForProduct(discountRuleId, productId)
}

func (s *DiscountService) CalculateDiscountForLineItem(discountId uuid.UUID, lineItem *models.LineItem, cart *models.Cart) (float64, *utils.ApplictaionError) {
	adjustment := 0.0
	if !lineItem.AllowDiscounts {
		return adjustment, nil
	}
	discount, err := s.Retrieve(discountId, sql.Options{Relations: []string{"rule"}})
	if err != nil {
		return 0, err
	}
	discountType := discount.Rule.Type
	discountValue := discount.Rule.Value
	discountAllocation := discount.Rule.Allocation
	calculationContext, err := s.r.TotalsService().GetCalculationContext(lineItem.Cart, lineItem.Order, CalculationContextOptions{ExcludeShipping: true})
	if err != nil {
		return 0, err
	}
	fullItemPrice := lineItem.UnitPrice * float64(lineItem.Quantity)
	includesTax := true && lineItem.IncludesTax
	if includesTax {
		lineItemTotals, err := s.r.NewTotalsService().SetContext(s.ctx).GetLineItemTotals([]models.LineItem{*lineItem}, true, calculationContext, nil)
		if err != nil {
			return 0, err
		}
		fullItemPrice = lineItemTotals[lineItem.Id].Subtotal
	}
	if discountType == models.DiscountRulePersentage {
		adjustment = math.Round(fullItemPrice / 100 * discountValue)
	} else if discountType == models.DiscountRuleFixed && discountAllocation == models.AllocationTotal {
		discountedItems := make([]models.LineItem, 0, len(cart.Items))
		for _, item := range cart.Items {
			if item.AllowDiscounts {
				discountedItems = append(discountedItems, item)
			}
		}
		totals, err := s.r.NewTotalsService().SetContext(s.ctx).GetLineItemTotals(discountedItems, includesTax, calculationContext, nil)
		if err != nil {
			return 0, err
		}
		subtotal := 0.0
		for _, total := range totals {
			subtotal += total.Subtotal
		}
		nominator := math.Min(float64(discountValue), float64(subtotal))
		totalItemPercentage := float64(fullItemPrice) / float64(subtotal)
		adjustment = nominator * totalItemPercentage
	} else {
		adjustment = discountValue * float64(lineItem.Quantity)
	}
	return math.Min(float64(adjustment), float64(fullItemPrice)), nil

}

func (s *DiscountService) validateDiscountForCartOrThrow(cart *models.Cart, discounts []models.Discount) *utils.ApplictaionError {
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
		if cart.CustomerId.UUID == uuid.Nil && s.hasCustomersGroupCondition(disc) {
			return utils.NewApplictaionError(
				utils.NOT_ALLOWED,
				fmt.Sprintf("models.Discount %s is only valid for specific customer", disc.Code),
				"500",
				nil,
			)
		}
		isValidForRegion, err := s.isValidForRegion(disc, cart.RegionId.UUID)
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
		if cart.CustomerId.UUID != uuid.Nil {
			canApplyForCustomer, err := s.canApplyForCustomer(disc.Rule.Id, cart.CustomerId.UUID)
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
		if cond.Type == models.DiscountConditionTypeCustomerGroups {
			return true
		}
	}
	return false
}

func (s *DiscountService) hasReachedLimit(discount models.Discount) bool {
	count := discount.UsageCount
	limit := discount.UsageLimit
	return count >= limit
}

func (s *DiscountService) hasNotStarted(discount models.Discount) bool {
	return discount.StartsAt.After(time.Now())
}

func (s *DiscountService) hasExpired(discount models.Discount) bool {
	return discount.EndsAt.Before(time.Now())
}

func (s *DiscountService) isDisabled(discount models.Discount) bool {
	return discount.IsDisabled
}

func (s *DiscountService) isValidForRegion(discount models.Discount, region_id uuid.UUID) (bool, *utils.ApplictaionError) {
	regions := discount.Regions
	if discount.ParentDiscountId.UUID != uuid.Nil {
		parent, err := s.Retrieve(discount.ParentDiscountId.UUID, sql.Options{Relations: []string{"rule", "regions"}})
		if err != nil {
			return false, err
		}
		regions = parent.Regions
	}
	for _, r := range regions {
		if r.Id == region_id {
			return true, nil
		}
	}
	return false, nil
}

func (s *DiscountService) canApplyForCustomer(discountRuleId uuid.UUID, customerId uuid.UUID) (bool, *utils.ApplictaionError) {
	if customerId == uuid.Nil {
		return false, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"`customerId` can not be null",
			"500",
			nil,
		)
	}
	customer, _ := s.r.CustomerService().SetContext(s.ctx).RetrieveById(customerId, sql.Options{Relations: []string{"groups"}})
	return s.r.DiscountConditionRepository().CanApplyForCustomer(discountRuleId, customer.Id)

}
