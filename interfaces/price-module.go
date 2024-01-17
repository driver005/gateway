package interfaces

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type RuleType struct {
	// The ID of the rule type.
	Id uuid.UUID `json:"id"`
	// The display name of the rule type.
	Name string `json:"name"`
	// The unique name used to later identify the rule_attribute. For example, it can be used in the `context` parameter of
	// the `calculatePrices` method to specify a rule for calculating the price.
	RuleAttribute string `json:"rule_attribute"`
	// The priority of the rule type. This is useful when calculating the price of a price set, and multiple rules satisfy
	// the provided context. The higher the value, the higher the priority of the rule type.
	DefaultPriority int `json:"default_priority"`
}

type PriceSet struct {
	// The ID of the price set.
	Id uuid.UUID `json:"id"`
	// The prices that belong to this price set.
	MoneyAmounts []models.MoneyAmount `json:"money_amounts,omitempty"`
	// The rule types applied on this price set.
	RuleTypes []RuleType `json:"rule_types,omitempty"`
}

type PriceListRuleValue struct {
	// The price list rule value's ID.
	Id uuid.UUID `json:"id"`
	// The rule's value.
	Value string `json:"value"`
	// The associated price list rule.
	PriceListRule PriceListRule `json:"price_list_rule"`
}

type PriceListRule struct {
	Id                  uuid.UUID
	Value               string
	RuleType            RuleType
	PriceList           models.PriceList
	PriceListRuleValues []PriceListRuleValue
}

type PriceRule struct {
	Id                    uuid.UUID
	PriceSetId            uuid.UUID
	PriceSet              PriceSet
	RuleTypeId            uuid.UUID
	RuleType              RuleType
	Value                 uuid.UUID
	Priority              int
	PriceSetMoneyAmountID string
	PriceListId           uuid.UUID
}

// PriceSetMoneyAmountRules represents the data transfer object for price set money amount rules
type PriceSetMoneyAmountRules struct {
	Id                  uuid.UUID
	PriceSetMoneyAmount PriceSetMoneyAmount
	RuleType            RuleType
	Value               string
}

// PriceSetMoneyAmount represents the data transfer object for price set money amount
type PriceSetMoneyAmount struct {
	Id          uuid.UUID
	Title       string
	PriceSet    PriceSet
	PriceList   models.PriceList
	PriceSetId  uuid.UUIDs
	PriceRules  []PriceRule
	MoneyAmount models.MoneyAmount
}

// AddPriceListPrices represents the data transfer object for adding prices to a price list
type AddPriceListPrices struct {
	// The ID of the price list to add prices to
	PriceListId uuid.UUID
	// The prices to add
	Prices []string
}

// CalculatedPrice represents the details of the calculated price
type CalculatedPrice struct {
	// The ID of the money amount selected as the calculated price
	MoneyAmountId uuid.UUID
	// The ID of the associated price list, if any
	PriceListId uuid.UUID
	// The type of the associated price list, if any
	PriceListType string
	// The `min_quantity` field defined on a money amount
	MinQuantity int
	// The `max_quantity` field defined on a money amount
	MaxQuantity int
}

// CalculatedPriceSet represents the calculated price set
type CalculatedPriceSet struct {
	// The ID of the price set
	Id uuid.UUID
	// Whether the calculated price is associated with a price list
	IsCalculatedPricePriceList bool `json:"is_calculated_price_price_list,omitempty"`
	// The amount of the calculated price, or `null` if there isn't a calculated price
	CalculatedAmount float64
	// Whether the original price is associated with a price list
	IsOriginalPricePriceList bool `json:"is_original_price_price_list,omitempty"`
	// The amount of the original price, or `null` if there isn't a calculated price
	OriginalAmount float64
	// The currency code of the calculated price, or null if there isn't a calculated price
	CurrencyCode string
	// The details of the calculated price
	CalculatedPrice *CalculatedPrice
	// The details of the original price
	OriginalPrice *CalculatedPrice
}

type IPricingModuleService interface {
	CalculatePrices(sharedContext context.Context, context map[string]interface{}, ids uuid.UUIDs) ([]CalculatedPriceSet, *utils.ApplictaionError)
	Retrieve(sharedContext context.Context, context map[string]interface{}, id uuid.UUID, config *sql.Options) (PriceSet, *utils.ApplictaionError)
	List(sharedContext context.Context, context map[string]interface{}, filters PriceSet, config *sql.Options) ([]PriceSet, *utils.ApplictaionError)
	ListAndCount(sharedContext context.Context, context map[string]interface{}, filters PriceSet, config *sql.Options) ([]PriceSet, *int64, *utils.ApplictaionError)
	Create(sharedContext context.Context, context map[string]interface{}, data PriceSet) (PriceSet, *utils.ApplictaionError)
	Creates(sharedContext context.Context, context map[string]interface{}, data []PriceSet) ([]PriceSet, *utils.ApplictaionError)
	Update(sharedContext context.Context, context map[string]interface{}, data []PriceSet) ([]PriceSet, *utils.ApplictaionError)
	RemoveRules(sharedContext context.Context, context map[string]interface{}, data []PriceSet) *utils.ApplictaionError
	Delete(sharedContext context.Context, context map[string]interface{}, ids []string) *utils.ApplictaionError
	AddPrice(sharedContext context.Context, context map[string]interface{}, data PriceSet) (PriceSet, *utils.ApplictaionError)
	AddPrices(sharedContext context.Context, context map[string]interface{}, data []PriceSet) ([]PriceSet, *utils.ApplictaionError)
	AddRule(sharedContext context.Context, context map[string]interface{}, data RuleType) (PriceSet, *utils.ApplictaionError)
	AddRules(sharedContext context.Context, context map[string]interface{}, data []RuleType) ([]PriceSet, *utils.ApplictaionError)
	RetrieveMoneyAmount(sharedContext context.Context, context map[string]interface{}, id uuid.UUID, config *sql.Options) (models.MoneyAmount, *utils.ApplictaionError)
	ListMoneyAmounts(sharedContext context.Context, context map[string]interface{}, filters models.MoneyAmount, config *sql.Options) ([]models.MoneyAmount, *utils.ApplictaionError)
	ListAndCountMoneyAmounts(sharedContext context.Context, context map[string]interface{}, filters models.MoneyAmount, config *sql.Options) ([]models.MoneyAmount, *int64, *utils.ApplictaionError)
	CreateMoneyAmounts(sharedContext context.Context, context map[string]interface{}, data []models.MoneyAmount) ([]models.MoneyAmount, *utils.ApplictaionError)
	UpdateMoneyAmounts(sharedContext context.Context, context map[string]interface{}, data []models.MoneyAmount) ([]models.MoneyAmount, *utils.ApplictaionError)
	DeleteMoneyAmounts(sharedContext context.Context, context map[string]interface{}, ids []string) *utils.ApplictaionError
	RetrieveCurrency(sharedContext context.Context, context map[string]interface{}, code string, config *sql.Options) (models.Currency, *utils.ApplictaionError)
	ListCurrencies(sharedContext context.Context, context map[string]interface{}, filters models.Currency, config *sql.Options) ([]models.Currency, *utils.ApplictaionError)
	ListAndCountCurrencies(sharedContext context.Context, context map[string]interface{}, filters models.Currency, config *sql.Options) ([]models.Currency, *int64, *utils.ApplictaionError)
	CreateCurrencies(sharedContext context.Context, context map[string]interface{}, data []models.Currency) ([]models.Currency, *utils.ApplictaionError)
	UpdateCurrencies(sharedContext context.Context, context map[string]interface{}, data []models.Currency) ([]models.Currency, *utils.ApplictaionError)
	DeleteCurrencies(sharedContext context.Context, context map[string]interface{}, currencyCodes []string) *utils.ApplictaionError
	RetrieveRuleType(sharedContext context.Context, context map[string]interface{}, id uuid.UUID, config *sql.Options) (RuleType, *utils.ApplictaionError)
	ListRuleTypes(sharedContext context.Context, context map[string]interface{}, filters RuleType, config *sql.Options) ([]RuleType, *utils.ApplictaionError)
	ListAndCountRuleTypes(sharedContext context.Context, context map[string]interface{}, filters RuleType, config *sql.Options) ([]RuleType, *int64, *utils.ApplictaionError)
	CreateRuleTypes(sharedContext context.Context, context map[string]interface{}, data []RuleType) ([]RuleType, *utils.ApplictaionError)
	UpdateRuleTypes(sharedContext context.Context, context map[string]interface{}, data []RuleType) ([]RuleType, *utils.ApplictaionError)
	DeleteRuleTypes(sharedContext context.Context, context map[string]interface{}, ruleTypeIds []string) *utils.ApplictaionError
	RetrievePriceSetMoneyAmountRules(sharedContext context.Context, context map[string]interface{}, id uuid.UUID, config *sql.Options) (PriceSetMoneyAmountRules, *utils.ApplictaionError)
	ListPriceSetMoneyAmountRules(sharedContext context.Context, context map[string]interface{}, filters PriceSetMoneyAmount, config *sql.Options) ([]PriceSetMoneyAmountRules, *utils.ApplictaionError)
	ListAndCountPriceSetMoneyAmountRules(sharedContext context.Context, context map[string]interface{}, filters PriceSetMoneyAmount, config *sql.Options) ([]PriceSetMoneyAmountRules, *int64, *utils.ApplictaionError)
	ListPriceSetMoneyAmounts(sharedContext context.Context, context map[string]interface{}, filters PriceSetMoneyAmount, config *sql.Options) ([]PriceSetMoneyAmount, *utils.ApplictaionError)
	ListAndCountPriceSetMoneyAmounts(sharedContext context.Context, context map[string]interface{}, filters PriceSetMoneyAmount, config *sql.Options) ([]PriceSetMoneyAmount, *int64, *utils.ApplictaionError)
	CreatePriceSetMoneyAmountRules(sharedContext context.Context, context map[string]interface{}, data []PriceSetMoneyAmount) ([]PriceSetMoneyAmountRules, *utils.ApplictaionError)
	UpdatePriceSetMoneyAmountRules(sharedContext context.Context, context map[string]interface{}, data []PriceSetMoneyAmount) ([]PriceSetMoneyAmountRules, *utils.ApplictaionError)
	DeletePriceSetMoneyAmountRules(sharedContext context.Context, context map[string]interface{}, ids []string) *utils.ApplictaionError
	RetrievePriceRule(sharedContext context.Context, context map[string]interface{}, id uuid.UUID, config *sql.Options) (PriceRule, *utils.ApplictaionError)
	ListPriceRules(sharedContext context.Context, context map[string]interface{}, filters PriceRule, config *sql.Options) ([]PriceRule, *utils.ApplictaionError)
	ListAndCountPriceRules(sharedContext context.Context, context map[string]interface{}, filters PriceRule, config *sql.Options) ([]PriceRule, *int64, *utils.ApplictaionError)
	CreatePriceRules(sharedContext context.Context, context map[string]interface{}, data []PriceRule) ([]PriceRule, *utils.ApplictaionError)
	UpdatePriceRules(sharedContext context.Context, context map[string]interface{}, data []PriceRule) ([]PriceRule, *utils.ApplictaionError)
	DeletePriceRules(sharedContext context.Context, context map[string]interface{}, priceRuleIds []string) *utils.ApplictaionError
	RetrievePriceList(sharedContext context.Context, context map[string]interface{}, id uuid.UUID, config *sql.Options) (models.PriceList, *utils.ApplictaionError)
	ListPriceLists(sharedContext context.Context, context map[string]interface{}, filters models.PriceList, config *sql.Options) ([]models.PriceList, *utils.ApplictaionError)
	ListAndCountPriceLists(sharedContext context.Context, context map[string]interface{}, filters models.PriceList, config *sql.Options) ([]models.PriceList, *int64, *utils.ApplictaionError)
	CreatePriceLists(sharedContext context.Context, context map[string]interface{}, data []models.PriceList) ([]models.PriceList, *utils.ApplictaionError)
	UpdatePriceLists(sharedContext context.Context, context map[string]interface{}, data []models.PriceList) ([]models.PriceList, *utils.ApplictaionError)
	DeletePriceLists(sharedContext context.Context, context map[string]interface{}, priceListIds []string) *utils.ApplictaionError
	RetrievePriceListRule(sharedContext context.Context, context map[string]interface{}, id uuid.UUID, config *sql.Options) (PriceListRule, *utils.ApplictaionError)
	ListPriceListRules(sharedContext context.Context, context map[string]interface{}, filters PriceListRule, config *sql.Options) ([]PriceListRule, *utils.ApplictaionError)
	ListAndCountPriceListRules(sharedContext context.Context, context map[string]interface{}, filters PriceListRule, config *sql.Options) ([]PriceListRule, *int64, *utils.ApplictaionError)
	CreatePriceListRules(sharedContext context.Context, context map[string]interface{}, data []PriceListRule) ([]PriceListRule, *utils.ApplictaionError)
	UpdatePriceListRules(sharedContext context.Context, context map[string]interface{}, data []PriceListRule) ([]PriceListRule, *utils.ApplictaionError)
	DeletePriceListRules(sharedContext context.Context, context map[string]interface{}, priceListRuleIds []string) *utils.ApplictaionError
	AddPriceListPrices(sharedContext context.Context, context map[string]interface{}, data []AddPriceListPrices) ([]models.PriceList, *utils.ApplictaionError)
	SetPriceListRules(sharedContext context.Context, context map[string]interface{}, data PriceListRule) (models.PriceList, *utils.ApplictaionError)
	RemovePriceListRules(sharedContext context.Context, context map[string]interface{}, data PriceListRule) (models.PriceList, *utils.ApplictaionError)
}
