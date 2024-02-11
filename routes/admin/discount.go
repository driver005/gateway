package admin

import (
	"fmt"

	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
)

type Discount struct {
	r Registry
}

func NewDiscount(r Registry) *Discount {
	m := Discount{r: r}
	return &m
}

func (m *Discount) SetRoutes(router fiber.Router) {
	route := router.Group("/discounts")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Get("/code/:code", m.GetDiscountByCode)

	// Dynamic codes
	route.Post("/:id/dynamic-codes", m.CreateDynamicCode)
	route.Delete("/:id/dynamic-codes/:code", m.DeleteConditon)

	// Discount region management
	route.Post("/:id/regions/:region_id", m.AddRegion)
	route.Delete("/:id/regions/:region_id", m.RemoveRegion)

	// Discount condition management
	route.Post("/:id/conditions", m.CreateConditon)
	route.Delete("/:id/conditions/:condition_id", m.DeleteConditon)

	route.Get("/:id/conditions/:condition_id", m.GetConditon)
	route.Post("/:id/conditions/:condition_id", m.UpdateConditon)
	route.Post("/:id/conditions/:condition_id/batch", m.AddResourcesToConditionBatch)
	route.Delete("/:id/conditions/:condition_id/batch", m.DeleteResourcesToConditionBatch)
}

func (m *Discount) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableDiscount](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.DiscountService().SetContext(context.Context()).ListAndCount(model, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   result,
		"count":  count,
		"offset": config.Skip,
		"limit":  config.Take,
	})
}

func (m *Discount) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateDiscountInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateDiscountInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.DiscountService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "disocunt",
		"deleted": true,
	})
}

func (m *Discount) GetConditon(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "condition_id")
	if err != nil {
		return err
	}

	result, err := m.r.DiscountConditionService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) GetDiscountByCode(context fiber.Ctx) error {
	config, err := api.BindConfig(context, m.r.Validator())
	if err != nil {
		return err
	}

	code := context.Params("code")

	result, err := m.r.DiscountService().SetContext(context.Context()).RetrieveByCode(code, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) AddRegion(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	regionId, err := utils.ParseUUID(context.Params("region_id"))
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).AddRegion(id, regionId)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) AddResourcesToConditionBatch(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.AddResourcesToConditionBatch](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	conditionId, err := utils.ParseUUID(context.Params("condition_id"))
	if err != nil {
		return err
	}

	condition, err := m.r.DiscountConditionService().SetContext(context.Context()).Retrieve(conditionId, &sql.Options{Selects: []string{"id", "type", "discount_rule_id"}})
	if err != nil {
		return err
	}

	input := &types.DiscountConditionInput{Id: conditionId, RuleId: condition.DiscountRuleId.UUID}
	if condition.Type == models.DiscountConditionTypeProducts {
		input.Products = model.Resources
	} else if condition.Type == models.DiscountConditionTypeProductTypes {
		input.ProductTypes = model.Resources
	} else if condition.Type == models.DiscountConditionTypeProductTags {
		input.ProductTags = model.Resources
	} else if condition.Type == models.DiscountConditionTypeProductCollections {
		input.ProductCollections = model.Resources
	} else if condition.Type == models.DiscountConditionTypeCustomerGroups {
		input.CustomerGroups = model.Resources
	}

	_, err = m.r.DiscountConditionService().SetContext(context.Context()).UpsertCondition(input, false)
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) CreateConditon(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.CreateConditon](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	discount, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	_, err = m.r.DiscountConditionService().SetContext(context.Context()).UpsertCondition(&types.DiscountConditionInput{Operator: model.Operator, RuleId: discount.RuleId.UUID}, false)
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) CreateDynamicCode(context fiber.Ctx) error {
	model, id, err := api.BindWithUUID[types.CreateDynamicDiscountInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	_, err = m.r.DiscountService().SetContext(context.Context()).CreateDynamicCode(id, model)
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) UpdateConditon(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.AdminUpsertConditionsReq](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	conditionId, err := utils.ParseUUID(context.Params("condition_id"))
	if err != nil {
		return err
	}

	condition, err := m.r.DiscountConditionService().SetContext(context.Context()).Retrieve(conditionId, &sql.Options{Selects: []string{"id"}})
	if err != nil {
		return err
	}

	discount, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, &sql.Options{Selects: []string{}})
	if err != nil {
		return err
	}

	_, err = m.r.DiscountConditionService().SetContext(context.Context()).UpsertCondition(&types.DiscountConditionInput{
		Id:                 condition.Id,
		RuleId:             discount.RuleId.UUID,
		Products:           model.Products,
		ProductCollections: model.ProductCollections,
		ProductTypes:       model.ProductTypes,
		ProductTags:        model.ProductTags,
		CustomerGroups:     model.CustomerGroups,
	}, false)
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) DeleteConditon(context fiber.Ctx) error {
	config, err := api.BindConfig(context, m.r.Validator())
	if err != nil {
		return err
	}

	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	conditionId, err := utils.ParseUUID(context.Params("condition_id"))
	if err != nil {
		return err
	}

	condition, err := m.r.DiscountConditionService().SetContext(context.Context()).Retrieve(conditionId, config)
	if err != nil {
		discount, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
		if err != nil {
			return err
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"id":       conditionId,
			"object":   "disocunt-condition",
			"deleted":  true,
			"discount": discount,
		})
	}

	discount, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, &sql.Options{Selects: []string{"id", "rule_id"}})
	if err != nil {
		return err
	}

	if condition.DiscountRuleId.UUID != discount.RuleId.UUID {
		return utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf(`Condition with id %s does not belong to Discount with id %s`, conditionId, id),
		)
	}

	if err := m.r.DiscountConditionService().SetContext(context.Context()).Delete(conditionId); err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":       conditionId,
		"object":   "disocunt-condition",
		"deleted":  true,
		"discount": result,
	})
}

func (m *Discount) DeleteDynamicCode(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	code := context.Params("code")

	if err := m.r.DiscountService().SetContext(context.Context()).DeleteDynamicCode(id, code); err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) DeleteResourcesToConditionBatch(context fiber.Ctx) error {
	model, id, config, err := api.BindAll[types.AddResourcesToConditionBatch](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	conditionId, err := utils.ParseUUID(context.Params("condition_id"))
	if err != nil {
		return err
	}

	condition, err := m.r.DiscountConditionService().SetContext(context.Context()).Retrieve(conditionId, &sql.Options{Selects: []string{"id", "type", "discount_rule_id"}})
	if err != nil {
		return err
	}

	input := &types.DiscountConditionInput{Id: conditionId, RuleId: condition.DiscountRuleId.UUID}
	if condition.Type == models.DiscountConditionTypeProducts {
		input.Products = model.Resources
	} else if condition.Type == models.DiscountConditionTypeProductTypes {
		input.ProductTypes = model.Resources
	} else if condition.Type == models.DiscountConditionTypeProductTags {
		input.ProductTags = model.Resources
	} else if condition.Type == models.DiscountConditionTypeProductCollections {
		input.ProductCollections = model.Resources
	} else if condition.Type == models.DiscountConditionTypeCustomerGroups {
		input.CustomerGroups = model.Resources
	}

	if err := m.r.DiscountConditionService().SetContext(context.Context()).RemoveResources(input); err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).Retrieve(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Discount) RemoveRegion(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	regionId, err := utils.ParseUUID(context.Params("region_id"))
	if err != nil {
		return err
	}

	result, err := m.r.DiscountService().SetContext(context.Context()).RemoveRegion(id, regionId)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}
