package admin

import (
	"github.com/driver005/gateway/api"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Product struct {
	r Registry
}

func NewProduct(r Registry) *Product {
	m := Product{r: r}
	return &m
}

func (m *Product) SetRoutes(router fiber.Router) {
	route := router.Group("/products")
	route.Get("/:id", m.Get)
	route.Get("", m.List)
	route.Post("", m.Create)
	route.Post("/:id", m.Update)
	route.Delete("/:id", m.Delete)

	route.Get("/types", m.ListTypes)
	route.Get("/tag-usage", m.ListTagUsageCount)
	route.Get("/:id/variants", m.ListVariants)
	route.Post("/:id/variants", m.CreateVariant)
	route.Delete("/:id/variants/:variant_id", m.DeletOption)
	route.Post("/:id/variants/:variant_id", m.UpdateVariant)
	route.Post("/:id/options/:option_id", m.UpdateOption)
	route.Delete("/:id/options/:option_id", m.DeletOption)
	route.Post("/:id/options/", m.AddOption)
	route.Post("/:id/metadata", m.SetMetadata)
}

func (m *Product) Get(context fiber.Ctx) error {
	id, config, err := api.BindGet(context, "id")
	if err != nil {
		return err
	}
	result, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, config)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Product) List(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableProduct](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.ProductService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *Product) Create(context fiber.Ctx) error {
	model, err := api.BindCreate[types.CreateProductInput](context, m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductService().SetContext(context.Context()).Create(model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Product) Update(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateProductInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	result, err := m.r.ProductService().SetContext(context.Context()).Update(id, model)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Product) Delete(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	if err := m.r.ProductService().SetContext(context.Context()).Delete(id); err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      id,
		"object":  "product",
		"deleted": true,
	})
}

func (m *Product) AddOption(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.CreateProductProductOption](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if _, err := m.r.ProductService().SetContext(context.Context()).AddOption(id, model.Title); err != nil {
		return err
	}

	result, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Product) DeletOption(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	optionId, err := api.BindDelete(context, "option_id")
	if err != nil {
		return err
	}

	if _, err := m.r.ProductService().SetContext(context.Context()).DeleteOption(id, optionId); err != nil {
		return err
	}

	result, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"option_id": optionId,
		"object":    "option",
		"deleted":   true,
		"product":   result,
	})
}

func (m *Product) CreateVariant(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.CreateProductVariantInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	if m.r.InventoryService() != nil {
		variants, err := m.r.ProductVariantService().SetContext(context.Context()).Create(id, nil, []types.CreateProductVariantInput{*model})
		if err != nil {
			return err
		}

		for _, variant := range variants {
			if !variant.ManageInventory {
				continue
			}

			inventoryItem, err := m.r.InventoryService().CreateInventoryItem(context.Context(), interfaces.CreateInventoryItemInput{
				SKU:           variant.Sku,
				OriginCountry: variant.OriginCountry,
				HsCode:        variant.HsCode,
				MidCode:       variant.MIdCode,
				Material:      variant.Material,
				Weight:        variant.Weight,
				Length:        variant.Length,
				Height:        variant.Height,
				Width:         variant.Width,
			})
			if err != nil {
				return err
			}

			if _, err := m.r.ProductVariantInventoryService().AttachInventoryItem([]models.ProductVariantInventoryItem{
				{
					VariantId:       uuid.NullUUID{UUID: variant.Id},
					InventoryItemId: uuid.NullUUID{UUID: inventoryItem.Id},
				},
			}); err != nil {
				return err
			}
		}
	} else {
		_, err := m.r.ProductVariantService().SetContext(context.Context()).Create(id, nil, []types.CreateProductVariantInput{*model})
		if err != nil {
			return err
		}
	}

	raw, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.PricingService().SetContext(context.Context()).SetAdminProductPricing([]models.Product{*raw})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Product) DeletVariant(context fiber.Ctx) error {
	id, err := api.BindDelete(context, "id")
	if err != nil {
		return err
	}

	variantId, err := api.BindDelete(context, "variant_id")
	if err != nil {
		return err
	}

	if err := m.r.ProductVariantService().SetContext(context.Context()).Delete(uuid.UUIDs{variantId}); err != nil {
		return err
	}

	raw, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.PricingService().SetContext(context.Context()).SetAdminProductPricing([]models.Product{*raw})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"option_id": variantId,
		"object":    "product-variant",
		"deleted":   true,
		"product":   result,
	})
}

func (m *Product) ListTagUsageCount(context fiber.Ctx) error {
	result, err := m.r.ProductService().SetContext(context.Context()).ListTagsByUsage(-1)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Product) ListTypes(context fiber.Ctx) error {
	result, err := m.r.ProductService().SetContext(context.Context()).ListTypes()
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Product) ListVariants(context fiber.Ctx) error {
	model, config, err := api.BindList[types.FilterableProductVariant](context)
	if err != nil {
		return err
	}
	result, count, err := m.r.ProductVariantService().SetContext(context.Context()).ListAndCount(model, config)
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

func (m *Product) SetMetadata(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.Metadata](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	data := &types.UpdateProductInput{}
	data.Metadata = data.Metadata.Add(model.Key, model.Value)

	if _, err := m.r.ProductService().SetContext(context.Context()).Update(id, data); err != nil {
		return err
	}

	raw, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.PricingService().SetContext(context.Context()).SetAdminProductPricing([]models.Product{*raw})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Product) UpdateOption(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.CreateProductProductOption](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	optionId, err := api.BindDelete(context, "option_id")
	if err != nil {
		return err
	}

	if _, err := m.r.ProductService().SetContext(context.Context()).UpdateOption(id, optionId, &types.ProductOptionInput{
		Title: model.Title,
	}); err != nil {
		return err
	}

	raw, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.PricingService().SetContext(context.Context()).SetAdminProductPricing([]models.Product{*raw})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)
}

func (m *Product) UpdateVariant(context fiber.Ctx) error {
	model, id, err := api.BindUpdate[types.UpdateProductVariantInput](context, "id", m.r.Validator())
	if err != nil {
		return err
	}

	variantId, err := api.BindDelete(context, "variant_id")
	if err != nil {
		return err
	}

	model.ProductId = id

	if _, err := m.r.ProductVariantService().SetContext(context.Context()).Update(variantId, nil, model); err != nil {
		return err
	}

	raw, err := m.r.ProductService().SetContext(context.Context()).RetrieveById(id, &sql.Options{})
	if err != nil {
		return err
	}

	result, err := m.r.PricingService().SetContext(context.Context()).SetAdminProductPricing([]models.Product{*raw})
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(result)

}
