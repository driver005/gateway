package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
)

type SalesChannel struct {
	r Registry
}

func NewSalesChannel(r Registry) *SalesChannel {
	m := SalesChannel{r: r}
	return &m
}

func (m *SalesChannel) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.SalesChannelService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *SalesChannel) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableSalesChannel](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.SalesChannelService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *SalesChannel) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateSalesChannelInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.SalesChannelService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *SalesChannel) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateSalesChannelInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.SalesChannelService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *SalesChannel) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.SalesChannelService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "sales-channel",
		"deleted": true,
	})
}

func (m *SalesChannel) AddProductsBatch(context fiber.Ctx) error {
	return nil
}

func (m *SalesChannel) DeleteProductsBatch(context fiber.Ctx) error {
	return nil
}

func (m *SalesChannel) AddStockLocation(context fiber.Ctx) error {
	return nil
}

func (m *SalesChannel) RemoveStockLocation(context fiber.Ctx) error {
	return nil
}
