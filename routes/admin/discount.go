package admin

import (
	"github.com/driver005/gateway/api"
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
	route.Get("/", m.List)
	route.Post("/", m.Create)
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
	return nil
}

func (m *Discount) GetDiscountByCode(context fiber.Ctx) error {
	return nil
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
	model, id, err := api.BindWithString[types.UpdateDiscountInput](context, "id", m.r.Validator())
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

	&types.DiscountConditionInput{Id: conditionId, RuleId: condition.DiscountRuleId.UUID, }

	model, err := m.r.DiscountConditionService().SetContext(context.Context()).UpsertCondition(, false)
	if err != nil {
		return err
	}
}

func (m *Discount) CreateConditon(context fiber.Ctx) error {
	return nil
}

func (m *Discount) CreateDynamicCode(context fiber.Ctx) error {
	return nil
}

func (m *Discount) UpdateConditon(context fiber.Ctx) error {
	return nil
}

func (m *Discount) DeleteConditon(context fiber.Ctx) error {
	return nil
}

func (m *Discount) DeleteDynamicCode(context fiber.Ctx) error {
	return nil
}

func (m *Discount) DeleteResourcesToConditionBatch(context fiber.Ctx) error {
	return nil
}

func (m *Discount) RemoveRegion(context fiber.Ctx) error {
	return nil
}
