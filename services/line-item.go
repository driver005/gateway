package services

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
)

type LineItemService struct {
	ctx context.Context
	r   Registry
}

func NewLineItemService(
	r Registry,
) *LineItemService {
	return &LineItemService{
		context.Background(),
		r,
	}
}

func (s *LineItemService) SetContext(context context.Context) *LineItemService {
	s.ctx = context
	return s
}

func (s *LineItemService) Retrieve(lineItemId uuid.UUID, config *sql.Options) (*models.LineItem, *utils.ApplictaionError) {
	if lineItemId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"lineItemId" must be defined`,
			nil,
		)
	}
	var res *models.LineItem

	query := sql.BuildQuery(models.LineItem{Model: core.Model{Id: lineItemId}}, config)

	if err := s.r.LineItemRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LineItemService) List(selector models.LineItem, config *sql.Options) ([]models.LineItem, *utils.ApplictaionError) {
	var res []models.LineItem

	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	query := sql.BuildQuery[models.LineItem](selector, config)

	if err := s.r.LineItemRepository().Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *LineItemService) CreateReturnLines(returnId uuid.UUID, cartId uuid.UUID) (*models.LineItem, *utils.ApplictaionError) {
	lineItem, returnItem, err := s.r.LineItemRepository().FindByReturn(s.ctx, returnId)
	if err != nil {
		return nil, err
	}

	model := &models.LineItem{
		Model: core.Model{
			Metadata: lineItem.Metadata,
		},
		CartId: uuid.NullUUID{
			UUID:  cartId,
			Valid: true,
		},
		Thumbnail:      lineItem.Thumbnail,
		IsReturn:       true,
		Title:          lineItem.Title,
		VariantId:      lineItem.VariantId,
		UnitPrice:      -1 * lineItem.UnitPrice,
		Quantity:       returnItem.Quantity,
		AllowDiscounts: lineItem.AllowDiscounts,
		IncludesTax:    lineItem.IncludesTax,
		TaxLines:       lineItem.TaxLines,
		Adjustments:    lineItem.Adjustments,
	}

	if err := s.r.LineItemRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *LineItemService) Generate(
	variantId uuid.UUID,
	variant []types.GenerateInputData,
	regionId uuid.UUID,
	quantity int,
	context types.GenerateLineItemContext,
) ([]models.LineItem, *utils.ApplictaionError) {
	er := s.validateGenerateArguments(variantId, variant, regionId, quantity, context)
	if er != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			er.Error(),
			nil,
		)
	}

	if regionId == uuid.Nil {
		regionId = context.RegionId
	}
	resolvedData := []types.GenerateInputData{}
	if variantId != uuid.Nil {
		resolvedData = append(resolvedData, types.GenerateInputData{
			VariantId: variantId,
			Quantity:  quantity,
		})
	} else {
		resolvedData = variant
	}
	resolvedDataMap := make(map[uuid.UUID]types.GenerateInputData)
	var variantIds uuid.UUIDs

	for _, d := range resolvedData {
		resolvedDataMap[d.VariantId] = d
		variantIds = append(variantIds, d.VariantId)
	}
	variants, err := s.r.ProductVariantService().SetContext(s.ctx).List(types.FilterableProductVariant{}, &sql.Options{
		Relations:     []string{"product"},
		Specification: []sql.Specification{sql.In("id", variantIds)},
	})
	if err != nil {
		return nil, err
	}
	var inputDataVariantId uuid.UUIDs
	for _, d := range resolvedData {
		inputDataVariantId = append(inputDataVariantId, d.VariantId)
	}
	var foundVariants uuid.UUIDs
	for _, v := range variants {
		foundVariants = append(foundVariants, v.Id)
	}
	var notFoundVariants uuid.UUIDs
	for _, variantId := range inputDataVariantId {
		if !slices.Contains(foundVariants, variantId) {
			notFoundVariants = append(notFoundVariants, variantId)
		}
	}
	if len(notFoundVariants) > 0 {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			fmt.Sprintf("Unable to generate the line items, some variant has not been found: %s", strings.Join(notFoundVariants.Strings(), ", ")),
			nil,
		)
	}
	variantsMap := make(map[uuid.UUID]models.ProductVariant)
	variantsToCalculatePricingFor := []interfaces.Pricing{}
	for _, variant := range variants {
		variantsMap[variant.Id] = variant
		variantResolvedData := resolvedDataMap[variant.Id]
		if context.UnitPrice == 0.0 && variantResolvedData.UnitPrice == 0.0 {
			variantsToCalculatePricingFor = append(variantsToCalculatePricingFor, struct {
				VariantId uuid.UUID
				Quantity  int
			}{
				VariantId: variant.Id,
				Quantity:  variantResolvedData.Quantity,
			})
		}
	}
	variantsPricing := make(map[uuid.UUID]types.ProductVariantPricing)
	if len(variantsToCalculatePricingFor) > 0 {
		variantsPricing, err = s.r.PricingService().SetContext(s.ctx).GetProductVariantsPricing(variantsToCalculatePricingFor, &interfaces.PricingContext{
			RegionId:              regionId,
			CustomerId:            context.CustomerId,
			IncludeDiscountPrices: true,
		})
		if err != nil {
			return nil, err
		}
	}
	var generatedItems []models.LineItem
	for _, variantData := range resolvedData {
		variant := variantsMap[variantData.VariantId]
		variantPricing := variantsPricing[variantData.VariantId]
		lineItem, err := s.generateLineItem(variant, variantData.Quantity, GenerateLineItemContext{
			GenerateLineItemContext: types.GenerateLineItemContext{
				UnitPrice: variantData.UnitPrice,
				Metadata:  variantData.Metadata,
			},
			VariantPricing: &variantPricing,
		})
		if err != nil {
			return nil, err
		}
		if context.Cart != nil {
			adjustments, err := s.r.LineItemAdjustmentService().SetContext(s.ctx).GenerateAdjustments(types.CalculationContextData{
				Discounts:       context.Cart.Discounts,
				Items:           context.Cart.Items,
				Customer:        context.Cart.Customer,
				Region:          context.Cart.Region,
				ShippingAddress: context.Cart.ShippingAddress,
				ShippingMethods: context.Cart.ShippingMethods,
			}, lineItem, &variant)
			if err != nil {
				return nil, err
			}
			lineItem.Adjustments = adjustments
		}
		generatedItems = append(generatedItems, *lineItem)
	}
	return generatedItems, nil
}

type GenerateLineItemContext struct {
	types.GenerateLineItemContext
	VariantPricing *types.ProductVariantPricing
}

func (s *LineItemService) generateLineItem(variant models.ProductVariant, quantity int, context GenerateLineItemContext) (*models.LineItem, *utils.ApplictaionError) {
	unitPrice := 0.0
	unitPriceIncludesTax := false
	shouldMerge := false
	if context.UnitPrice != 0.0 {
		unitPrice = context.UnitPrice
	} else {
		if context.UnitPrice == 0.0 {
			shouldMerge = true
			unitPriceIncludesTax = context.VariantPricing.CalculatedPriceIncludesTax
			unitPrice = context.VariantPricing.CalculatedPrice
		}
		if unitPrice == 0.0 {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				fmt.Sprintf("Cannot generate line item for variant \"%s\" without a price", variant.Title),
				"500",
				nil,
			)
		}
	}
	lineItem := &models.LineItem{
		Model: core.Model{
			Metadata: context.Metadata,
		},
		UnitPrice:      unitPrice,
		Title:          variant.Product.Title,
		Description:    variant.Title,
		Thumbnail:      variant.Product.Thumbnail,
		VariantId:      uuid.NullUUID{UUID: variant.Id},
		Quantity:       quantity,
		AllowDiscounts: variant.Product.Discountable,
		IsGiftcard:     variant.Product.IsGiftcard,
		ShouldMerge:    shouldMerge,
	}

	feature := true
	tax := true

	if feature {
		lineItem.ProductId = variant.ProductId
	}
	if tax {
		lineItem.IncludesTax = unitPriceIncludesTax
	}
	lineItem.OrderEditId = uuid.NullUUID{UUID: context.OrderEditId}

	if err := s.r.LineItemRepository().Save(s.ctx, lineItem); err != nil {
		return nil, err
	}
	return lineItem, nil
}

func (s *LineItemService) Create(data []models.LineItem) ([]models.LineItem, *utils.ApplictaionError) {
	if err := s.r.LineItemRepository().SaveSlice(s.ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *LineItemService) Update(id uuid.UUID, selector *models.LineItem, Update *models.LineItem, config *sql.Options) (*models.LineItem, *utils.ApplictaionError) {
	if id != uuid.Nil {
		selector = &models.LineItem{Model: core.Model{Id: id}}
	}

	lineItems, err := s.List(*selector, config)
	if err != nil {
		return nil, err
	}

	Update.Id = lineItems[0].Id

	if err := s.r.LineItemRepository().Save(s.ctx, Update); err != nil {
		return nil, err
	}
	return Update, nil
}

func (s *LineItemService) Delete(id uuid.UUID) *utils.ApplictaionError {
	item, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}
	if err := s.r.LineItemRepository().Remove(s.ctx, item); err != nil {
		return err
	}
	return nil
}

func (s *LineItemService) DeleteWithTaxLines(id uuid.UUID) *utils.ApplictaionError {
	err := s.r.TaxProviderService().SetContext(s.ctx).ClearLineItemsTaxLines(uuid.UUIDs{id})
	if err != nil {
		return err
	}
	return s.Delete(id)
}

func (s *LineItemService) CreateTaxLine(data *models.LineItemTaxLine) (*models.LineItemTaxLine, *utils.ApplictaionError) {
	if err := s.r.LineItemTaxLineRepository().Create(s.ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *LineItemService) CloneTo(ids uuid.UUIDs, data *models.LineItem, options map[string]interface{}) ([]models.LineItem, *utils.ApplictaionError) {
	lineItems, err := s.List(models.LineItem{}, &sql.Options{
		Relations:     []string{"tax_lines", "adjustments"},
		Specification: []sql.Specification{sql.In("id", ids)},
	})
	if err != nil {
		return nil, err
	}

	var originalItemId uuid.NullUUID
	if data.OrderId.UUID == uuid.Nil && data.SwapId.UUID == uuid.Nil && data.ClaimOrderId.UUID == uuid.Nil && data.CartId.UUID == uuid.Nil && data.OrderEditId.UUID == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Unable to clone a line item that is not attached to at least one of: order_edit, order, swap, claim or cart.",
			nil,
		)
	}
	for i, item := range lineItems {
		if options["setOriginalLineItemId"].(bool) {
			originalItemId = uuid.NullUUID{UUID: item.Id}
		}
		item.Id = uuid.Nil

		if !reflect.ValueOf(data.OrderId).IsZero() {
			item.OrderId = data.OrderId
		}
		if !reflect.ValueOf(data.SwapId).IsZero() {
			item.SwapId = data.SwapId
		}
		if !reflect.ValueOf(data.ClaimOrderId).IsZero() {
			item.ClaimOrderId = data.ClaimOrderId
		}
		if !reflect.ValueOf(data.CartId).IsZero() {
			item.CartId = data.CartId
		}
		if !reflect.ValueOf(data.OrderEditId).IsZero() {
			item.OrderEditId = data.OrderEditId
		}
		if !reflect.ValueOf(originalItemId).IsZero() {
			item.OriginalItemId = originalItemId
		}
		item.TaxLines = []models.LineItemTaxLine{}
		for _, taxLine := range item.TaxLines {
			taxLine.Id = uuid.Nil
			taxLine.ItemId = uuid.NullUUID{}
			item.TaxLines = append(item.TaxLines, taxLine)
		}
		item.Adjustments = []models.LineItemAdjustment{}
		for _, adj := range item.Adjustments {
			adj.Id = uuid.Nil
			adj.ItemId = uuid.NullUUID{}
			item.Adjustments = append(item.Adjustments, adj)
		}
		lineItems[i] = item
	}

	if err := s.r.LineItemRepository().SaveSlice(s.ctx, lineItems); err != nil {
		return nil, err
	}

	return lineItems, nil
}

func (s *LineItemService) validateGenerateArguments(
	variantId uuid.UUID,
	variant []types.GenerateInputData,
	regionId uuid.UUID,
	quantity int,
	context types.GenerateLineItemContext,
) error {
	errorMessage := "Unable to generate the line item because one or more required argument(s) are missing"
	if variant == nil {
		if quantity == 0 || regionId == uuid.Nil {
			return fmt.Errorf("%s. Ensure quantity, regionId, and variantId are passed", errorMessage)
		}
		if variantId == uuid.Nil {
			return fmt.Errorf("%s. Ensure variant id is passed", errorMessage)
		}
		return nil
	}

	if context.RegionId == uuid.Nil && regionId == uuid.Nil {
		return fmt.Errorf("%s. Ensure region or region_id are passed", errorMessage)
	}

	hasMissingVariantId := false
	for _, d := range variant {
		if d.VariantId == uuid.Nil {
			hasMissingVariantId = true
			break
		}
	}
	if hasMissingVariantId {
		return fmt.Errorf("%s. Ensure a variant id is passed for each variant", errorMessage)
	}

	return nil
}
