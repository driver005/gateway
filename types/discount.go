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

// @oas:schema:AdminPostDiscountsDiscountConditionsCondition
// type: object
// properties:
//
//	products:
//	   type: array
//	   description: list of product IDs if the condition's type is `products`.
//	   items:
//	     type: string
//	product_types:
//	   type: array
//	   description: list of product type IDs if the condition's type is `product_types`.
//	   items:
//	     type: string
//	product_collections:
//	   type: array
//	   description: list of product collection IDs if the condition's type is `product_collections`.
//	   items:
//	     type: string
//	product_tags:
//	   type: array
//	   description: list of product tag IDs if the condition's type is `product_tags`
//	   items:
//	     type: string
//	customer_groups:
//	   type: array
//	   description: list of customer group IDs if the condition's type is `customer_groups`.
//	   items:
//	     type: string
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

// @oas:schema:AdminPostDiscountsReq
// type: object
// description: "The details of the discount to create."
// required:
//   - code
//   - rule
//   - regions
//
// properties:
//
//	code:
//	  type: string
//	  description: A unique code that will be used to redeem the discount
//	is_dynamic:
//	  type: boolean
//	  description: Whether the discount should have multiple instances of itself, each with a different code. This can be useful for automatically generated discount codes that all have to follow a common set of rules.
//	  default: false
//	rule:
//	  description: The discount rule that defines how discounts are calculated
//	  type: object
//	  required:
//	     - type
//	     - value
//	     - allocation
//	  properties:
//	    description:
//	      type: string
//	      description: "A short description of the discount"
//	    type:
//	      type: string
//	      description: >-
//	        The type of the discount, can be `fixed` for discounts that reduce the price by a fixed amount, `percentage` for percentage reductions or `free_shipping` for shipping vouchers.
//	      enum: [fixed, percentage, free_shipping]
//	    value:
//	      type: number
//	      description: "The value that the discount represents. This will depend on the type of the discount."
//	    allocation:
//	      type: string
//	      description: >-
//	        The scope that the discount should apply to. `total` indicates that the discount should be applied on the cart total, and `item` indicates that the discount should be applied to each discountable item in the cart.
//	      enum: [total, item]
//	    conditions:
//	      type: array
//	      description: "A set of conditions that can be used to limit when the discount can be used. Only one of `products`, `product_types`, `product_collections`, `product_tags`, and `customer_groups` should be provided based on the discount condition's type."
//	      items:
//	        type: object
//	        required:
//	           - operator
//	        properties:
//	          operator:
//	            type: string
//	            description: >-
//	              Operator of the condition. `in` indicates that discountable resources are within the specified resources. `not_in` indicates that
//	              discountable resources are everything but the specified resources.
//	            enum: [in, not_in]
//	          products:
//	            type: array
//	            description: list of product IDs if the condition's type is `products`.
//	            items:
//	              type: string
//	          product_types:
//	            type: array
//	            description: list of product type IDs if the condition's type is `product_types`.
//	            items:
//	              type: string
//	          product_collections:
//	            type: array
//	            description: list of product collection IDs if the condition's type is `product_collections`.
//	            items:
//	              type: string
//	          product_tags:
//	            type: array
//	            description: list of product tag IDs if the condition's type is `product_tags`.
//	            items:
//	              type: string
//	          customer_groups:
//	            type: array
//	            description: list of customer group IDs if the condition's type is `customer_groups`.
//	            items:
//	              type: string
//	is_disabled:
//	  type: boolean
//	  description: >-
//	    Whether the discount code is disabled on creation. If set to `true`, it will not be available for customers.
//	  default: false
//	starts_at:
//	  type: string
//	  format: date-time
//	  description: The date and time at which the discount should be available.
//	ends_at:
//	  type: string
//	  format: date-time
//	  description: The date and time at which the discount should no longer be available.
//	valid_duration:
//	  type: string
//	  description: The duration the discount runs between
//	  example: P3Y6M4DT12H30M5S
//	regions:
//	  description: A list of region IDs representing the Regions in which the Discount can be used.
//	  type: array
//	  items:
//	    type: string
//	usage_limit:
//	  type: number
//	  description: Maximum number of times the discount can be used
//	metadata:
//	  description: An optional set of key-value pairs to hold additional information.
//	  type: object
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
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

// @oas:schema:AdminPostDiscountsDiscountReq
// type: object
// description: "The details of the discount to update."
// properties:
//
//	code:
//	  type: string
//	  description: A unique code that will be used to redeem the discount
//	rule:
//	  description: The discount rule that defines how discounts are calculated
//	  type: object
//	  required:
//	    - id
//	  properties:
//	    id:
//	      type: string
//	      description: "The ID of the Rule"
//	    description:
//	      type: string
//	      description: "A short description of the discount"
//	    value:
//	      type: number
//	      description: "The value that the discount represents. This will depend on the type of the discount."
//	    allocation:
//	      type: string
//	      description: >-
//	        The scope that the discount should apply to. `total` indicates that the discount should be applied on the cart total, and `item` indicates that the discount should be applied to each discountable item in the cart.
//	      enum: [total, item]
//	    conditions:
//	      type: array
//	      description: "A set of conditions that can be used to limit when the discount can be used. Only one of `products`, `product_types`, `product_collections`, `product_tags`, and `customer_groups` should be provided based on the discount condition's type."
//	      items:
//	        type: object
//	        required:
//	          - operator
//	        properties:
//	          id:
//	            type: string
//	            description: "The ID of the condition"
//	          operator:
//	            type: string
//	            description: >-
//	              Operator of the condition. `in` indicates that discountable resources are within the specified resources. `not_in` indicates that
//	              discountable resources are everything but the specified resources.
//	            enum: [in, not_in]
//	          products:
//	            type: array
//	            description: list of product IDs if the condition's type is `products`.
//	            items:
//	              type: string
//	          product_types:
//	            type: array
//	            description: list of product type IDs if the condition's type is `product_types`.
//	            items:
//	              type: string
//	          product_collections:
//	            type: array
//	            description: list of product collection IDs if the condition's type is `product_collections`.
//	            items:
//	              type: string
//	          product_tags:
//	            type: array
//	            description: list of product tag IDs if the condition's type is `product_tags`.
//	            items:
//	              type: string
//	          customer_groups:
//	            type: array
//	            description: list of customer group IDs if the condition's type is `customer_groups`.
//	            items:
//	              type: string
//	is_disabled:
//	  type: boolean
//	  description: >-
//	    Whether the discount code is disabled on creation. If set to `true`, it will not be available for customers.
//	starts_at:
//	  type: string
//	  format: date-time
//	  description: The date and time at which the discount should be available.
//	ends_at:
//	  type: string
//	  format: date-time
//	  description: The date and time at which the discount should no longer be available.
//	valid_duration:
//	  type: string
//	  description: The duration the discount runs between
//	  example: P3Y6M4DT12H30M5S
//	usage_limit:
//	  type: number
//	  description: Maximum number of times the discount can be used
//	regions:
//	  description: A list of region IDs representing the Regions in which the Discount can be used.
//	  type: array
//	  items:
//	    type: string
//	metadata:
//	   description: An object containing metadata of the discount
//	   type: object
//	   externalDocs:
//	     description: "Learn about the metadata attribute, and how to delete and update it."
//	     url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
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

// @oas:schema:AdminPostDiscountsDiscountDynamicCodesReq
// type: object
// description: "The details of the dynamic discount to create."
// required:
//   - code
//
// properties:
//
//	code:
//	  type: string
//	  description: A unique code that will be used to redeem the Discount
//	usage_limit:
//	  type: number
//	  description: Maximum number of times the discount code can be used
//	  default: 1
//	metadata:
//	  type: object
//	  description: An optional set of key-value pairs to hold additional information.
//	  externalDocs:
//	    description: "Learn about the metadata attribute, and how to delete and update it."
//	    url: "https://docs.medusajs.com/development/entities/overview#metadata-attribute"
type CreateDynamicDiscountInput struct {
	Code       string     `json:"code"`
	EndsAt     *time.Time `json:"ends_at,omitempty" validate:"omitempty"`
	UsageLimit int        `json:"usage_limit"`
	Metadata   core.JSONB `json:"metadata,omitempty" validate:"omitempty"`
}

// @oas:schema:AdminDeleteDiscountsDiscountConditionsConditionBatchReq
// type: object
// description: "The resources to remove."
// required:
//   - resources
//
// properties:
//
//	resources:
//	  description: The resources to be removed from the discount condition
//	  type: array
//	  items:
//	    type: object
//	    required:
//	      - id
//	    properties:
//	      id:
//	        description: The id of the item
//	        type: string
type AddResourcesToConditionBatch struct {
	Resources []string `json:"resources"`
}

// @oas:schema:AdminPostDiscountsDiscountConditions
// type: object
// required:
//   - operator
//
// properties:
//
//	operator:
//	   description: >-
//	     Operator of the condition. `in` indicates that discountable resources are within the specified resources. `not_in` indicates that
//	     discountable resources are everything but the specified resources.
//	   type: string
//	   enum: [in, not_in]
//	products:
//	   type: array
//	   description: list of product IDs if the condition's type is `products`.
//	   items:
//	     type: string
//	product_types:
//	   type: array
//	   description: list of product type IDs if the condition's type is `product_types`.
//	   items:
//	     type: string
//	product_collections:
//	   type: array
//	   description: list of product collection IDs if the condition's type is `product_collections`.
//	   items:
//	     type: string
//	product_tags:
//	   type: array
//	   description: list of product tag IDs if the condition's type is `product_tags`.
//	   items:
//	     type: string
//	customer_groups:
//	   type: array
//	   description: list of customer group IDs if the condition's type is `customer_groups`.
//	   items:
//	     type: string
type CreateConditon struct {
	Operator models.DiscountConditionOperator `json:"operator"`
}
