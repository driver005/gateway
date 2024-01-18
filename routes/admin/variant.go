package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/services"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type ResponseInventoryItem struct {
	interfaces.InventoryItemDTO
	LocationLevels   []interfaces.InventoryLevelDTO `json:"location_levels,omitempty"`
	StockedQuantity  int
	ReservedQuantity int
}

type SalesChannelAvailability struct {
	ChannelId         uuid.UUID `json:"channel_id"`
	ChannelName       string    `json:"channel_name"`
	AvailableQuantity int       `json:"available_quantity"`
}

type VariantInventory struct {
	Id                       uuid.UUID                  `json:"id"`
	Inventory                []ResponseInventoryItem    `json:"inventory"`
	SalesChannelAvailability []SalesChannelAvailability `json:"sales_channel_availability"`
}

type AdminGetVariantsParams struct {
	types.FilterableProductVariant
	CartId       uuid.UUID `json:"cart_id,omitempty" validate:"omitempty"`
	RegionId     uuid.UUID `json:"region_id,omitempty" validate:"omitempty"`
	CurrencyCode string    `json:"currency_code,omitempty" validate:"omitempty"`
	CustomerId   uuid.UUID `json:"customer_id,omitempty" validate:"omitempty"`
}

type Variant struct {
	r Registry
}

func NewVariant(r Registry) *Variant {
	m := Variant{r: r}
	return &m
}

func (m *Variant) SetRoutes(router fiber.Router) {
	route := router.Group("/variants")
	route.Get("/:id", m.Get)
	route.Get("/", m.List)

	route.Post("/:id/inventory", m.GetInventory)
}

func (m *Variant) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}

	rawVariant, err := m.r.ProductVariantService().Retrieve(id, config)
	if err != nil {
		return err
	}

	variant, err := m.r.PricingService().SetAdminVariantPricing([]models.ProductVariant{*rawVariant}, &interfaces.PricingContext{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(variant)
}

func (m *Variant) List(context fiber.Ctx) error {
	model, config, err := api.BindList[AdminGetVariantsParams](context)
	if err != nil {
		return err
	}

	rawVariants, count, err := m.r.ProductVariantService().ListAndCount(&types.FilterableProductVariant{
		Title:             model.Title,
		ProductId:         model.ProductId,
		Product:           model.Product,
		SKU:               model.SKU,
		Barcode:           model.Barcode,
		EAN:               model.EAN,
		UPC:               model.UPC,
		InventoryQuantity: model.InventoryQuantity,
		AllowBackorder:    model.AllowBackorder,
		ManageInventory:   model.ManageInventory,
		HSCode:            model.HSCode,
		OriginCountry:     model.OriginCountry,
		MidCode:           model.MidCode,
		Material:          model.Material,
		Weight:            model.Weight,
		Length:            model.Length,
		Height:            model.Height,
		Width:             model.Width,
	}, config)
	if err != nil {
		return err
	}

	regionId := model.RegionId
	currencyCode := model.CurrencyCode
	if model.CartId != uuid.Nil {
		cart, err := m.r.CartService().Retrieve(model.CartId, &sql.Options{
			Selects: []string{"id", "region_id"},
		}, services.TotalsConfig{})
		if err != nil {
			return err
		}
		region, err := m.r.RegionService().Retrieve(cart.RegionId.UUID, &sql.Options{
			Selects: []string{"id", "currency_code"},
		})
		if err != nil {
			return err
		}
		regionId = region.Id
		currencyCode = region.CurrencyCode
	}

	variants, err := m.r.PricingService().SetAdminVariantPricing(rawVariants, &interfaces.PricingContext{
		CartId:                model.CartId,
		RegionId:              regionId,
		CurrencyCode:          currencyCode,
		CustomerId:            model.CustomerId,
		IncludeDiscountPrices: true,
		IgnoreCache:           true,
	})
	if err != nil {
		return err
	}

	if m.r.InventoryService() != nil {
		salesChannelsIds, err := m.r.SalesChannelService().List(&types.FilterableSalesChannel{}, &sql.Options{
			Selects: []string{"id"},
		})
		if err != nil {
			return err
		}
		v, err := m.r.ProductVariantInventoryService().SetVariantAvailability(variants, func() uuid.UUIDs {
			var ids uuid.UUIDs
			for _, salesChannel := range salesChannelsIds {
				ids = append(ids, salesChannel.Id)
			}
			return ids
		}(), &services.AvailabilityContext{})
		if err != nil {
			return err
		}

		variants = v
	}

	return context.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"variants": variants,
		"count":    count,
		"offset":   config.Skip,
		"limit":    config.Take,
	})

}

func (m *Variant) GetInventory(context fiber.Ctx) error {
	id, err := utils.ParseUUID(context.Params("id"))
	if err != nil {
		return err
	}

	variant, _ := m.r.ProductVariantService().Retrieve(id, &sql.Options{Selects: []string{"id"}})
	responseVariant := VariantInventory{
		Id:                       variant.Id,
		Inventory:                []ResponseInventoryItem{},
		SalesChannelAvailability: []SalesChannelAvailability{},
	}

	rawChannels, _, err := m.r.SalesChannelService().ListAndCount(&types.FilterableSalesChannel{}, &sql.Options{})
	if err != nil {
		return err
	}
	var channels []models.SalesChannel
	for _, channel := range rawChannels {
		locationIds, _ := m.r.SalesChannelLocationService().ListLocationIds(uuid.UUIDs{channel.Id})
		channels = append(channels, models.SalesChannel{
			Model: core.Model{
				Id: channel.Id,
			},
			Name:        channel.Name,
			LocationIds: locationIds,
		})
	}

	variantInventoryItems, err := m.r.ProductVariantInventoryService().ListByVariant(uuid.UUIDs{variant.Id})
	if err != nil {
		return err
	}
	inventory, _ := m.r.ProductVariantInventoryService().ListInventoryItemsByVariant(variant.Id)
	in, err := joinLevels(inventory, uuid.UUIDs{}, m.r.InventoryService())
	if err != nil {
		return err
	}
	responseVariant.Inventory = in

	if len(inventory) > 0 {
		for _, channel := range channels {
			if len(channel.LocationIds) == 0 {
				responseVariant.SalesChannelAvailability = append(responseVariant.SalesChannelAvailability, SalesChannelAvailability{
					ChannelName:       channel.Name,
					ChannelId:         channel.Id,
					AvailableQuantity: 0,
				})
			} else {
				quantity, err := m.r.ProductVariantInventoryService().GetVariantQuantityFromVariantInventoryItems(variantInventoryItems, channel.Id)
				if err != nil {
					return err
				}
				responseVariant.SalesChannelAvailability = append(responseVariant.SalesChannelAvailability, SalesChannelAvailability{
					ChannelName:       channel.Name,
					ChannelId:         channel.Id,
					AvailableQuantity: *quantity,
				})
			}
		}
	}

	return context.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"variant": responseVariant,
	})
}
