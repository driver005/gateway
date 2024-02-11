package types

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type FilterableDiscount struct {
	core.FilterModel

	Q          string                         `json:"q,omitempty" validate:"omitempty"`
	IsDynamic  bool                           `json:"is_dynamic,omitempty" validate:"omitempty"`
	IsDisabled bool                           `json:"is_disabled,omitempty" validate:"omitempty"`
	Rule       *AdminGetDiscountsDiscountRule `json:"rule,omitempty" validate:"omitempty"`
}

type AdminGetDiscountsDiscountRule struct {
	Type       *models.DiscountRule   `json:"type,omitempty" validate:"omitempty"`
	Allocation *models.AllocationType `json:"allocation,omitempty" validate:"omitempty"`
}

// AdminUpsertConditionsReq represents the fields to create or update a discount condition.
type AdminUpsertConditionsReq struct {
	Products           []string `json:"products,omitempty" validate:"omitempty"`
	ProductCollections []string `json:"product_collections,omitempty" validate:"omitempty"`
	ProductTypes       []string `json:"product_types,omitempty" validate:"omitempty"`
	ProductTags        []string `json:"product_tags,omitempty" validate:"omitempty"`
	CustomerGroups     []string `json:"customer_groups,omitempty" validate:"omitempty"`
}

// DiscountConditionMapTypeToProperty maps the discount condition type to its corresponding property.
var DiscountConditionMapTypeToProperty = map[models.DiscountConditionType]string{
	models.DiscountConditionTypeProducts:           "products",
	models.DiscountConditionTypeProductTypes:       "product_types",
	models.DiscountConditionTypeProductCollections: "product_collections",
	models.DiscountConditionTypeProductTags:        "product_tags",
	models.DiscountConditionTypeCustomerGroups:     "customer_groups",
}

// DiscountConditionInput represents the input for a discount condition.
type DiscountConditionInput struct {
	RuleId             uuid.UUID                        `json:"rule_id,omitempty" validate:"omitempty"`
	Id                 uuid.UUID                        `json:"id,omitempty" validate:"omitempty"`
	Operator           models.DiscountConditionOperator `json:"operator,omitempty" validate:"omitempty"`
	Products           []string                         `json:"products,omitempty" validate:"omitempty"`
	ProductCollections []string                         `json:"product_collections,omitempty" validate:"omitempty"`
	ProductTypes       []string                         `json:"product_types,omitempty" validate:"omitempty"`
	ProductTags        []string                         `json:"product_tags,omitempty" validate:"omitempty"`
	CustomerGroups     []string                         `json:"customer_groups,omitempty" validate:"omitempty"`
}

// CreateDiscountRuleInput represents the input for creating a discount rule.
type CreateDiscountRuleInput struct {
	Description string                   `json:"description,omitempty" validate:"omitempty"`
	Type        models.DiscountRuleType  `json:"type"`
	Value       float64                  `json:"value"`
	Allocation  models.AllocationType    `json:"allocation"`
	Conditions  []DiscountConditionInput `json:"conditions,omitempty" validate:"omitempty"`
}

// CreateDiscountInput represents the input for creating a discount.
type CreateDiscountInput struct {
	Code          string                  `json:"code"`
	Rule          CreateDiscountRuleInput `json:"rule"`
	IsDynamic     bool                    `json:"is_dynamic"`
	IsDisabled    bool                    `json:"is_disabled"`
	StartsAt      *time.Time              `json:"starts_at,omitempty" validate:"omitempty"`
	EndsAt        *time.Time              `json:"ends_at,omitempty" validate:"omitempty"`
	ValidDuration *time.Time              `json:"valid_duration,omitempty" validate:"omitempty"`
	UsageLimit    int                     `json:"usage_limit,omitempty" validate:"omitempty"`
	Regions       uuid.UUIDs              `json:"regions,omitempty" validate:"omitempty"`
	Metadata      core.JSONB              `json:"metadata,omitempty" validate:"omitempty"`
}

// UpdateDiscountRuleInput represents the input for updating a discount rule.
type UpdateDiscountRuleInput struct {
	Id          uuid.UUID                `json:"id"`
	Description string                   `json:"description,omitempty" validate:"omitempty"`
	Value       float64                  `json:"value,omitempty" validate:"omitempty"`
	Allocation  *models.AllocationType   `json:"allocation,omitempty" validate:"omitempty"`
	Conditions  []DiscountConditionInput `json:"conditions,omitempty" validate:"omitempty"`
}

// UpdateDiscountInput represents the input for updating a discount.
type UpdateDiscountInput struct {
	Code          string                   `json:"code,omitempty" validate:"omitempty"`
	Rule          *UpdateDiscountRuleInput `json:"rule,omitempty" validate:"omitempty"`
	IsDisabled    bool                     `json:"is_disabled,omitempty" validate:"omitempty"`
	StartsAt      *time.Time               `json:"starts_at,omitempty" validate:"omitempty"`
	EndsAt        *time.Time               `json:"ends_at,omitempty" validate:"omitempty"`
	ValidDuration *time.Time               `json:"valid_duration,omitempty" validate:"omitempty"`
	UsageLimit    int                      `json:"usage_limit,omitempty" validate:"omitempty"`
	Regions       uuid.UUIDs               `json:"regions,omitempty" validate:"omitempty"`
	Metadata      core.JSONB               `json:"metadata,omitempty" validate:"omitempty"`
}

// CreateDynamicDiscountInput represents the input for creating a dynamic discount.
type CreateDynamicDiscountInput struct {
	Code       string     `json:"code"`
	EndsAt     *time.Time `json:"ends_at,omitempty" validate:"omitempty"`
	UsageLimit int        `json:"usage_limit"`
	Metadata   core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

type AddResourcesToConditionBatch struct {
	Resources []string `json:"resources"`
}

type CreateConditon struct {
	Operator models.DiscountConditionOperator `json:"operator"`
}
