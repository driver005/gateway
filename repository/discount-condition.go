package repository

import (
	"context"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DiscountConditionJoinTableForeignKey string

const (
	PRODUCT_ID            DiscountConditionJoinTableForeignKey = "product_id"
	PRODUCT_TYPE_ID       DiscountConditionJoinTableForeignKey = "product_type_id"
	PRODUCT_COLLECTION_ID DiscountConditionJoinTableForeignKey = "product_collection_id"
	PRODUCT_TAG_ID        DiscountConditionJoinTableForeignKey = "product_tag_id"
	CUSTOMER_GROUP_ID     DiscountConditionJoinTableForeignKey = "customer_group_id"
)

type DiscountConditionResourceType interface{}

type DiscountConditionRepo struct {
	sql.Repository[models.DiscountCondition]
}

func DiscountConditionRepository(db *gorm.DB) *DiscountConditionRepo {
	return &DiscountConditionRepo{*sql.NewRepository[models.DiscountCondition](db)}
}

func (r *DiscountConditionRepo) FindOneWithDiscount(conditionId uuid.UUID, discountId uuid.UUID) (*models.DiscountCondition, *utils.ApplictaionError) {
	condition := &models.DiscountCondition{}
	err := r.Db().Model(&models.DiscountCondition{}).
		Joins("LEFT JOIN condition on condition.discount_rule_id = discount.rule_id and discount.id = ? and condition.id = ?", discountId, conditionId).
		First(condition).Error
	if err != nil {
		return nil, utils.NewApplictaionError(
			utils.DB_ERROR,
			err.Error(),
			nil,
		)
	}
	return condition, nil
}

func (r *DiscountConditionRepo) GetJoinTableResourceIdentifiers(types models.DiscountConditionType) (*DiscountConditionResourceType, string, *DiscountConditionJoinTableForeignKey, string, string, string) {
	var conditionTable DiscountConditionResourceType
	joinTable := "product"
	joinTableForeignKey := PRODUCT_ID
	joinTableKey := "id"
	relatedTable := ""
	resourceKey := ""
	switch types {
	case "PRODUCTS":
		resourceKey = "id"
		joinTableForeignKey = PRODUCT_ID
		joinTable = "product"
		conditionTable = models.DiscountConditionProduct{}
	case "PRODUCT_TYPES":
		resourceKey = "type_id"
		joinTableForeignKey = PRODUCT_TYPE_ID
		joinTable = "product"
		relatedTable = "types"
		conditionTable = models.DiscountConditionProductType{}
	case "PRODUCT_COLLECTIONS":
		resourceKey = "collection_id"
		joinTableForeignKey = PRODUCT_COLLECTION_ID
		joinTable = "product"
		relatedTable = "collections"
		conditionTable = models.DiscountConditionProductCollection{}
	case "PRODUCT_TAGS":
		joinTableKey = "product_id"
		resourceKey = "product_tag_id"
		joinTableForeignKey = PRODUCT_TAG_ID
		joinTable = "product_tags"
		relatedTable = "tags"
		conditionTable = models.DiscountConditionProductTag{}
	case "CUSTOMER_GROUPS":
		joinTableKey = "customer_id"
		resourceKey = "customer_group_id"
		joinTable = "customer_group_customers"
		joinTableForeignKey = CUSTOMER_GROUP_ID
		conditionTable = models.DiscountConditionCustomerGroup{}
	}
	return &conditionTable, joinTable, &joinTableForeignKey, resourceKey, joinTableKey, relatedTable
}

// TODO: ADD
func (r *DiscountConditionRepo) RemoveConditionResources(ctx context.Context, id uuid.UUID, resource *models.DiscountCondition) *utils.ApplictaionError {
	conditionTable, _, joinTableForeignKey, _, _, _ := r.GetJoinTableResourceIdentifiers(resource.Type)
	if conditionTable == nil || joinTableForeignKey == nil {
		return nil
	}

	var res *models.DiscountCondition

	res.Id = id

	if err := r.Remove(ctx, res); err != nil {
		return err
	}

	return nil
}

func (r *DiscountConditionRepo) AddConditionResources(ctx context.Context, conditionId uuid.UUID, resource *models.DiscountCondition, overrideExisting bool) ([]models.DiscountCondition, *utils.ApplictaionError) {
	// toInsert := make([]map[string]string, len(resourceIds))
	conditionTable, _, joinTableForeignKey, _, _, _ := r.GetJoinTableResourceIdentifiers(resource.Type)
	if conditionTable == nil || joinTableForeignKey == nil {
		return nil, nil
	}

	// _, err := r.Db().Model(&models.DiscountCondition{}).
	// 	Insert().
	// 	OrIgnore(true).
	// 	Into(conditionTable).
	// 	Values(toInsert).
	// 	Execute()
	// if err != nil {
	// 	return nil, err
	// }
	// if overrideExisting {
	// 	_, err := r.Db().Model(&models.DiscountCondition{}).
	// 		Delete().
	// 		From(conditionTable).
	// 		Where(map[string]interface{}{
	// 			"condition_id":      conditionId,
	// 			joinTableForeignKey: typeorm.Not(typeorm.In(idsToInsert)),
	// 		}).
	// 		Execute()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	// return r.Manager.CreateQueryBuilder(conditionTable, "discon").
	// 	Select().
	// 	Where(map[string]interface{}{
	// 		"id": idsToInsert,
	// 	}).
	// 	GetMany()

	return nil, nil
}

func (r *DiscountConditionRepo) QueryConditionTable(types models.DiscountConditionType, conditionId uuid.UUID, resourceId uuid.UUID) (*int64, *utils.ApplictaionError) {
	conditionTable, joinTable, joinTableForeignKey, resourceKey, joinTableKey, _ := r.GetJoinTableResourceIdentifiers(types)
	medusaV2Flag := false

	//TODO: Add
	if types != "CUSTOMER_GROUPS" && medusaV2Flag {
		// module := medusajs.GetModuleInstance(modules.PRODUCT)[modules.PRODUCT]
		// prop := relatedTable
		// resource, err := module.Retrieve(resourceId, map[string]interface{}{
		// 	"select":    []string{prop + ".id"},
		// 	"relations": []string{prop},
		// })
		// if err != nil {
		// 	return 0, err
		// }
		// if resource == nil {
		// 	return 0, nil
		// }
		// relatedResourceIds := make([]string, len(resource[prop]))
		// for i, relatedResource := range resource[prop] {
		// 	relatedResourceIds[i] = relatedResource.id
		// }
		// if len(relatedResourceIds) == 0 {
		// 	return 0, nil
		// }
		// return r.Manager.CreateQueryBuilder(conditionTable, "dc").
		// 	Where("dc.condition_id = ? AND dc.?? IN (?)", conditionId, joinTableForeignKey, relatedResourceIds).
		// 	GetCount()
	}
	var count *int64
	err := r.Db().Model(conditionTable).
		Joins(joinTable, "resource", "dc.?? = resource.?? and resource.?? = ?", joinTableForeignKey, resourceKey, joinTableKey, resourceId).
		Where("dc.condition_id = ?", conditionId).
		Count(count).Error

	if err != nil {
		return nil, utils.NewApplictaionError(
			utils.DB_ERROR,
			err.Error(),
			nil,
		)
	}

	return count, nil
}

func (r *DiscountConditionRepo) IsValidForProduct(discountRuleId uuid.UUID, productId uuid.UUID) (bool, *utils.ApplictaionError) {
	var discountConditions []models.DiscountCondition
	err := r.Db().Model(&models.DiscountCondition{}).
		Select([]string{"discon.id", "discon.type", "discon.operator"}).
		Where("discon.discount_rule_id = ?", discountRuleId).
		Find(&discountConditions).Error
	if err != nil {
		return false, utils.NewApplictaionError(
			utils.DB_ERROR,
			err.Error(),
			nil,
		)
	}
	if len(discountConditions) == 0 {
		return true, nil
	}
	for _, condition := range discountConditions {
		if condition.Type == "CUSTOMER_GROUPS" {
			continue
		}
		numConditions, err := r.QueryConditionTable(condition.Type, condition.Id, productId)
		if err != nil {
			return false, err
		}
		if condition.Operator == "IN" && *numConditions == 0 {
			return false, nil
		}
		if condition.Operator == "NOT_IN" && *numConditions > 0 {
			return false, nil
		}
	}
	return true, nil
}

func (r *DiscountConditionRepo) CanApplyForCustomer(discountRuleId uuid.UUID, customerId uuid.UUID) (bool, *utils.ApplictaionError) {
	var discountConditions []models.DiscountCondition
	err := r.Db().Model(&models.DiscountCondition{}).
		Select([]string{"discon.id", "discon.type", "discon.operator"}).
		Where("discon.discount_rule_id = ?", discountRuleId).
		Where("discon.type = ?", "CUSTOMER_GROUPS").
		Find(&discountConditions).Error
	if err != nil {
		return false, utils.NewApplictaionError(
			utils.DB_ERROR,
			err.Error(),
			nil,
		)
	}
	if len(discountConditions) == 0 {
		return true, nil
	}
	for _, condition := range discountConditions {
		numConditions, err := r.QueryConditionTable("customer_groups", condition.Id, customerId)
		if err != nil {
			return false, err
		}
		if condition.Operator == "IN" && *numConditions == 0 {
			return false, nil
		}
		if condition.Operator == "NOT_IN" && *numConditions > 0 {
			return false, nil
		}
	}
	return true, nil
}
