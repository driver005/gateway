package services

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"slices"
	"strings"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

type TotalsConfig struct {
	ForceTaxes bool
}

type CartService struct {
	ctx context.Context
	r   Registry
}

func NewCartService(
	r Registry,
) *CartService {
	return &CartService{
		context.Background(),
		r,
	}
}

func (s *CartService) SetContext(context context.Context) *CartService {
	s.ctx = context
	return s
}

func (s *CartService) List(selector types.FilterableCartProps, config sql.Options) ([]models.Cart, *utils.ApplictaionError) {
	var res []models.Cart

	if reflect.DeepEqual(config, sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	query := sql.BuildQuery(selector, config)

	if err := s.r.CartRepository().Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CartService) Retrieve(id uuid.UUID, config sql.Options, totalsConfig TotalsConfig) (*models.Cart, *utils.ApplictaionError) {
	if id == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"id" must be defined`,
			"500",
			nil,
		)
	}
	var res *models.Cart

	featurev2 := true
	if featurev2 {
		if len(config.Relations) > 0 {
			for i := 0; i < len(config.Relations); i++ {
				if strings.HasPrefix(config.Relations[i], "items.variant") {
					config.Relations[i] = "items"
				}
			}
		}
		config.Relations = removeDuplicate(config.Relations)
	}
	_, _, totalsToSelect := s.transformQueryForTotals(config)
	if len(totalsToSelect) > 0 {
		return s.RetrieveLegacy(id, config, totalsConfig)
	}

	query := sql.BuildQuery(models.Cart{Model: core.Model{Id: id}}, config)
	if err := s.r.CartRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CartService) RetrieveLegacy(id uuid.UUID, config sql.Options, totalsConfig TotalsConfig) (*models.Cart, *utils.ApplictaionError) {
	var res *models.Cart
	query := sql.BuildQuery(models.Cart{Model: core.Model{Id: id}}, config)

	selects, relations, totalsToSelect := s.transformQueryForTotals(config)
	query.Selects = selects
	query.Relations = relations

	if err := s.r.CartRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return s.decorateTotals(res, totalsToSelect, totalsConfig)
}

func (s *CartService) RetrieveWithTotals(id uuid.UUID, config sql.Options, totalsConfig TotalsConfig) (*models.Cart, *utils.ApplictaionError) {
	relations := s.getTotalsRelations(config)
	config.Relations = relations

	featurev2 := true
	if featurev2 {
		if len(config.Relations) > 0 {
			for i := 0; i < len(config.Relations); i++ {
				if strings.HasPrefix(config.Relations[i], "items.variant") {
					config.Relations[i] = "items"
				}
			}
		}
		config.Relations = removeDuplicate(config.Relations)
	}
	cart, err := s.Retrieve(id, config, totalsConfig)
	if err != nil {
		return nil, err
	}
	return s.DecorateTotals(cart, totalsConfig)
}

func (s *CartService) Create(data *models.Cart) (*models.Cart, *utils.ApplictaionError) {
	cart := data

	feature := true
	featurev2 := true
	if feature && !featurev2 {
		salesChannel, err := s.getValidatedSalesChannel(data.SalesChannelId.UUID)
		if err != nil {
			return nil, err
		}
		cart.SalesChannelId = uuid.NullUUID{UUID: salesChannel.Id}
	}
	if data.CustomerId.UUID != uuid.Nil || data.Customer != nil {
		customer := data.Customer
		if data.CustomerId.UUID != uuid.Nil {
			c, err := s.r.CustomerService().SetContext(s.ctx).RetrieveById(data.CustomerId.UUID, sql.Options{})
			if err != nil {
				return nil, err
			}
			customer = c
		}
		cart.Customer = customer
		cart.CustomerId = uuid.NullUUID{UUID: customer.Id}
		cart.Email = customer.Email
	}
	if cart.Email == "" && data.Email != "" {
		customer, err := s.createOrFetchGuestCustomerFromEmail(data.Email)
		if err != nil {
			return nil, err
		}
		cart.Customer = customer
		cart.CustomerId = uuid.NullUUID{UUID: customer.Id}
		cart.Email = customer.Email
	}
	if data.RegionId.UUID == uuid.Nil && data.Region == nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`A region_id must be provided when creating a cart`,
			"500",
			nil,
		)
	}

	cart.RegionId = data.RegionId
	region := data.Region

	if region == nil {
		r, err := s.r.RegionService().SetContext(s.ctx).Retrieve(data.RegionId.UUID, sql.Options{
			Relations: []string{"countries"},
		})
		if err != nil {
			return nil, err
		}
		region = r
	}

	regCountries := []string{}
	for _, country := range region.Countries {
		regCountries = append(regCountries, country.Iso2)
	}

	if data.ShippingAddress == nil && data.ShippingAddressId.UUID == uuid.Nil {
		if len(region.Countries) == 1 {
			cart.ShippingAddress = &models.Address{
				CountryCode: regCountries[0],
			}
		}
	} else {
		if data.ShippingAddress != nil {
			if !slices.Contains(regCountries, data.ShippingAddress.CountryCode) {
				return nil, utils.NewApplictaionError(
					utils.NOT_ALLOWED,
					"Shipping country not in region",
					"500",
					nil,
				)
			}
			cart.ShippingAddress = data.ShippingAddress
		}
		if data.ShippingAddressId.UUID != uuid.Nil {
			var addr *models.Address

			query := sql.BuildQuery(models.Address{Model: core.Model{Id: data.ShippingAddressId.UUID}}, sql.Options{})

			if err := s.r.AddressRepository().FindOne(s.ctx, addr, query); err != nil {
				return nil, err
			}

			if addr != nil && !slices.Contains(regCountries, addr.CountryCode) {
				return nil, utils.NewApplictaionError(
					utils.NOT_ALLOWED,
					"Shipping country not in region",
					"500",
					nil,
				)
			}
			cart.ShippingAddressId = data.ShippingAddressId
		}
	}
	if data.BillingAddress != nil {
		if !slices.Contains(regCountries, data.BillingAddress.CountryCode) {
			return nil, utils.NewApplictaionError(
				utils.NOT_ALLOWED,
				"Billing country not in region",
				"500",
				nil,
			)
		}
		cart.BillingAddress = data.BillingAddress
	}
	if data.BillingAddressId.UUID != uuid.Nil {
		var addr *models.Address

		query := sql.BuildQuery(models.Address{Model: core.Model{Id: data.BillingAddressId.UUID}}, sql.Options{})

		if err := s.r.AddressRepository().FindOne(s.ctx, addr, query); err != nil {
			return nil, err
		}

		if addr != nil && !slices.Contains(regCountries, addr.CountryCode) {
			return nil, utils.NewApplictaionError(
				utils.NOT_ALLOWED,
				"Billing country not in region",
				"500",
				nil,
			)
		}
		cart.BillingAddressId = data.BillingAddressId
	}

	if err := s.r.CartRepository().Save(s.ctx, cart); err != nil {
		return nil, err
	}
	if feature && featurev2 {
		salesChannel, err := s.getValidatedSalesChannel(data.SalesChannelId.UUID)
		if err != nil {
			return nil, err
		}

		s.r.SalesChannelService().SetContext(s.ctx).Create(&models.SalesChannel{Model: core.Model{Id: salesChannel.Id}})
		// s.remoteLink.Create(map[string]interface{}{
		// 	"cartService": map[string]interface{}{
		// 		"cart_id": cart.Id,
		// 	},
		// 	"r.SalesChannelService()": map[string]interface{}{
		// 		"sales_channel_id": salesChannel.Id,
		// 	},
		// })
	}
	// s.eventBus_.Emit(CartService.Events.CREATED, map[string]interface{}{
	// 	"id": cart.Id,
	// })
	return cart, nil
}

func (s *CartService) getValidatedSalesChannel(salesChannelId uuid.UUID) (*models.SalesChannel, *utils.ApplictaionError) {
	var salesChannel *models.SalesChannel
	featurev2 := true
	if salesChannelId != uuid.Nil {
		if featurev2 {
			result, err := s.r.SalesChannelService().SetContext(s.ctx).RetrieveById(salesChannelId, sql.Options{})
			if err != nil {
				return nil, err
			}
			salesChannel = result
		} else {
			result, err := s.r.SalesChannelService().SetContext(s.ctx).RetrieveById(salesChannelId, sql.Options{})
			if err != nil {
				return nil, err
			}
			salesChannel = result
		}
	} else {
		result, err := s.r.StoreService().SetContext(s.ctx).Retrieve(sql.Options{
			Relations: []string{"default_sales_channel"},
		})
		if err != nil {
			return nil, err
		}
		salesChannel = result.DefaultSalesChannel
	}
	if salesChannel.IsDisabled {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			fmt.Sprintf("Unable to assign the cart to a disabled Sales Channel \"%s\"", salesChannel.Name),
			"500",
			nil,
		)
	}
	return salesChannel, nil
}

func (s *CartService) RemoveLineItem(id uuid.UUID, lineItemIds uuid.UUIDs) *utils.ApplictaionError {
	cart, err := s.Retrieve(id, sql.Options{
		Relations: []string{"items.variant.product.profiles", "shipping_methods"},
	}, TotalsConfig{})

	var lineItems []models.LineItem
	for _, item := range cart.Items {
		if slices.Contains(lineItemIds, item.Id) {
			item.HasShipping = false
			lineItems = append(lineItems, item)
		}
	}

	if len(lineItems) == 0 {
		return err
	}

	if len(cart.ShippingMethods) > 0 {
		err := s.r.ShippingOptionService().SetContext(s.ctx).DeleteShippingMethods(cart.ShippingMethods)
		if err != nil {
			return err
		}
	}

	if err := s.r.LineItemRepository().UpdateSlice(s.ctx, lineItems); err != nil {
		return err
	}

	for _, id := range lineItemIds {
		if err := s.r.LineItemService().SetContext(s.ctx).Delete(id); err != nil {
			return err
		}
	}

	result, err := s.Retrieve(id, sql.Options{
		Relations: []string{"items.variant.product.profiles", "discounts.rule", "region"},
	}, TotalsConfig{})
	if err != nil {
		return err
	}

	if err := s.refreshAdjustments(result); err != nil {
		return err
	}

	// s.eventBus_.Emit(CartService.Events.UPDATED, map[string]interface{}{
	// 	"id": cart.Id,
	// })
	return nil
}

func (s *CartService) validateLineItemShipping(
	shippingMethods []models.ShippingMethod,
	lineItemShippingProfiledId uuid.UUID,
) bool {
	if lineItemShippingProfiledId == uuid.Nil {
		return true
	}
	if len(shippingMethods) > 0 && lineItemShippingProfiledId != uuid.Nil {
		var selectedProfiles uuid.UUIDs
		for _, method := range shippingMethods {
			selectedProfiles = append(selectedProfiles, method.ShippingOption.ProfileId.UUID)
		}
		for _, profile := range selectedProfiles {
			if profile == lineItemShippingProfiledId {
				return true
			}
		}
	}
	return false
}

func (s *CartService) ValidateLineItem(
	salesChannelId uuid.UUID,
	lineItem models.LineItem,
) (bool, *utils.ApplictaionError) {
	if salesChannelId == uuid.Nil || lineItem.VariantId.UUID == uuid.Nil {
		return true, nil
	}

	if lineItem.Variant == nil || lineItem.Variant.ProductId.UUID == uuid.Nil {
		productVariant, err := s.r.ProductVariantService().SetContext(s.ctx).Retrieve(lineItem.VariantId.UUID, sql.Options{Selects: []string{"id", "product_id"}})
		if err != nil {
			return false, err
		}
		lineItem.Variant = productVariant
	}

	products, err := s.r.ProductService().SetContext(s.ctx).FilterProductsBySalesChannel(uuid.UUIDs{lineItem.Variant.ProductId.UUID}, salesChannelId, sql.Options{})
	if err != nil {
		return false, err
	}

	return len(products) > 0, nil
}

func (s *CartService) AddLineItem(
	cartId uuid.UUID,
	lineItem models.LineItem,
	validateSalesChannels bool,
) *utils.ApplictaionError {
	fields := []string{"id"}
	relations := []string{"shipping_methods"}

	feature := true
	featurev2 := true
	if feature {
		if featurev2 {
			relations = append(relations, "sales_channels")
		} else {
			fields = append(fields, "sales_channel_id")
		}
	}
	cart, err := s.Retrieve(cartId, sql.Options{
		Selects:   fields,
		Relations: relations,
	}, TotalsConfig{})
	if err != nil {
		return err
	}
	if feature {
		if validateSalesChannels {
			if lineItem.VariantId.UUID != uuid.Nil {
				lineItemIsValid, err := s.ValidateLineItem(cart.SalesChannelId.UUID, lineItem)
				if err != nil {
					return err
				}
				if !lineItemIsValid {
					return utils.NewApplictaionError(
						utils.INVALID_DATA,
						fmt.Sprintf("The product \"%s\" must belongs to the sales channel on which the cart has been created.", lineItem.Title),
						"500",
						nil,
					)
				}
			}
		}
	}

	var currentItem *models.LineItem

	if lineItem.ShouldMerge {
		existingItems, err := s.r.LineItemService().SetContext(s.ctx).List(models.LineItem{
			CartId:      uuid.NullUUID{UUID: cart.Id},
			VariantId:   lineItem.VariantId,
			ShouldMerge: true,
		}, sql.Options{
			Selects: []string{"id", "metadata", "quantity"},
			Take:    gox.NewInt(1),
		})
		if err != nil {
			return err
		}
		if len(existingItems) > 0 && reflect.DeepEqual(existingItems[0].Metadata, lineItem.Metadata) {
			currentItem = &existingItems[0]
		}
	}

	quantity := lineItem.Quantity

	if currentItem != nil {
		currentItem.Quantity += lineItem.Quantity
		quantity = currentItem.Quantity
	}

	if lineItem.VariantId.UUID != uuid.Nil {
		isCovered, err := s.r.ProductVariantInventoryService().SetContext(s.ctx).ConfirmInventory(lineItem.VariantId.UUID, quantity, map[string]interface{}{
			"salesChannelId": cart.SalesChannelId.UUID,
		})
		if err != nil {
			return err
		}
		if !isCovered {
			return utils.NewApplictaionError(
				utils.INSUFFICIENT_INVENTORY,
				fmt.Sprintf("Variant with id: %s does not have the required inventory", lineItem.VariantId.UUID),
				"500",
				nil,
			)
		}
	}
	if currentItem != nil {
		_, err = s.r.LineItemService().SetContext(s.ctx).Update(currentItem.Id, nil, &models.LineItem{
			Quantity: currentItem.Quantity,
		}, sql.Options{})
		if err != nil {
			return err
		}
	} else {
		_, err = s.r.LineItemService().SetContext(s.ctx).Create([]models.LineItem{{
			Quantity:    lineItem.Quantity,
			HasShipping: false,
			CartId:      uuid.NullUUID{UUID: cart.Id},
		}})
		if err != nil {
			return err
		}
	}

	_, err = s.r.LineItemService().SetContext(s.ctx).Update(
		uuid.Nil,
		&models.LineItem{
			CartId:      uuid.NullUUID{UUID: cartId},
			HasShipping: true,
		},
		&models.LineItem{
			HasShipping: false,
		},
		sql.Options{},
	)
	if err != nil {
		return err
	}

	if len(cart.ShippingMethods) > 0 {
		err = s.r.ShippingOptionService().SetContext(s.ctx).DeleteShippingMethods(cart.ShippingMethods)
		if err != nil {
			return err
		}
	}

	cart, err = s.Retrieve(cart.Id, sql.Options{
		Relations: []string{
			"items.variant.product.profiles",
			"discounts",
			"discounts.rule",
			"region",
		},
	}, TotalsConfig{})
	if err != nil {
		return err
	}

	err = s.refreshAdjustments(cart)
	if err != nil {
		return err
	}
	// err = s.eventBus_.emit(CartService.Events.UPDATED, map[string]interface{}{
	// 	"id": cart.Id,
	// })
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (s *CartService) AddOrUpdateLineItems(
	cartId uuid.UUID,
	lineItems []models.LineItem,
	validateSalesChannels bool,
) *utils.ApplictaionError {
	fields := []string{"id", "customerId", "region_id"}
	relations := []string{"shipping_methods"}

	feature := true
	featurev2 := true
	if feature {
		if featurev2 {
			relations = append(relations, "sales_channels")
		} else {
			fields = append(fields, "sales_channel_id")
		}
	}
	cart, err := s.Retrieve(cartId, sql.Options{
		Selects:   fields,
		Relations: relations,
	}, TotalsConfig{})
	if err != nil {
		return err
	}

	if feature {
		if validateSalesChannels {
			var areValid []bool
			for _, item := range lineItems {
				if item.VariantId.UUID != uuid.Nil {
					valid, err := s.ValidateLineItem(cart.SalesChannelId.UUID, item)
					if err != nil {
						return err
					}
					areValid = append(areValid, valid)
				} else {
					areValid = append(areValid, true)
				}
			}

			var invalidProducts []models.LineItem
			for i, valid := range areValid {
				if !valid {
					invalidProducts = append(invalidProducts, models.LineItem{
						Title: lineItems[i].Title,
					})
				}
			}
			if len(invalidProducts) > 0 {
				return utils.NewApplictaionError(
					utils.INVALID_DATA,
					fmt.Sprintf("The products [%v] must belongs to the sales channel on which the cart has been created.", invalidProducts),
					"500",
					nil,
				)
			}
		}
	}

	var variantIds uuid.UUIDs
	for _, item := range lineItems {
		variantIds = append(variantIds, item.VariantId.UUID)
	}

	existingItems, err := s.r.LineItemService().SetContext(s.ctx).List(models.LineItem{
		CartId:      uuid.NullUUID{UUID: cart.Id},
		ShouldMerge: true,
	}, sql.Options{
		Selects: []string{"id", "metadata", "quantity", "variant_id"},
		Specification: []sql.Specification{
			sql.In("variant_id", variantIds),
		},
	})
	if err != nil {
		return err
	}

	existingItemsVariantMap := make(map[uuid.UUID]models.LineItem)
	for _, item := range existingItems {
		existingItemsVariantMap[item.VariantId.UUID] = item
	}

	var lineItemsToCreate []models.LineItem
	lineItemsToUpdate := make(map[uuid.UUID]models.LineItem)
	for _, item := range lineItems {
		var currentItem *models.LineItem
		existingItem, ok := existingItemsVariantMap[item.VariantId.UUID]
		if item.ShouldMerge {
			if ok && reflect.DeepEqual(existingItem.Metadata, item.Metadata) {
				currentItem = &existingItem
			}
		}
		item.Quantity = currentItem.Quantity + item.Quantity
		if item.VariantId.UUID != uuid.Nil {
			isSufficient, err := s.r.ProductVariantInventoryService().SetContext(s.ctx).ConfirmInventory(item.VariantId.UUID, item.Quantity, map[string]interface{}{
				"salesChannelId": cart.SalesChannelId.UUID,
			})
			if err != nil {
				return err
			}

			if !isSufficient {
				return utils.NewApplictaionError(
					utils.INSUFFICIENT_INVENTORY,
					fmt.Sprintf("Variant with id: %s does not have the required inventory", item.VariantId.UUID),
					"500",
					nil,
				)
			}
		}
		if currentItem != nil {
			variantsPricing, err := s.r.PricingService().SetContext(s.ctx).GetProductVariantsPricing([]interfaces.Pricing{
				{
					VariantId: item.VariantId.UUID,
					Quantity:  item.Quantity,
				},
			},
				&interfaces.PricingContext{
					RegionId:              cart.RegionId.UUID,
					CustomerId:            cart.CustomerId.UUID,
					IncludeDiscountPrices: true,
				},
			)
			if err != nil {
				return err
			}
			calculatedPrice := variantsPricing[currentItem.VariantId.UUID].CalculatedPrice
			lineItemsToUpdate[currentItem.Id] = models.LineItem{
				Quantity:    item.Quantity,
				HasShipping: false,
			}
			if calculatedPrice != 0 {
				item := lineItemsToUpdate[currentItem.Id]
				item.UnitPrice = calculatedPrice
				lineItemsToUpdate[currentItem.Id] = item
			}
		} else {
			item.Variant = &models.ProductVariant{}
			item.HasShipping = false
			item.CartId = uuid.NullUUID{UUID: cart.Id}
			lineItemsToCreate = append(lineItemsToCreate, item)
		}
	}

	if len(lineItemsToUpdate) > 0 {
		for id, item := range lineItemsToUpdate {
			s.r.LineItemService().SetContext(s.ctx).Update(id, nil, &item, sql.Options{})
		}
		if err != nil {
			return err
		}
	}

	_, err = s.r.LineItemService().SetContext(s.ctx).Create(lineItemsToCreate)
	if err != nil {
		return err
	}
	_, err = s.r.LineItemService().SetContext(s.ctx).Update(uuid.Nil, &models.LineItem{
		CartId:      uuid.NullUUID{UUID: cartId},
		HasShipping: true,
	}, &models.LineItem{
		HasShipping: false,
	}, sql.Options{})
	if err != nil {
		return err
	}

	if len(cart.ShippingMethods) > 0 {
		err = s.r.ShippingOptionService().SetContext(s.ctx).DeleteShippingMethods(cart.ShippingMethods)
		if err != nil {
			return err
		}
	}
	cart, err = s.Retrieve(cart.Id, sql.Options{
		Relations: []string{
			"items.variant.product.profiles",
			"discounts",
			"discounts.rule",
			"region",
		},
	}, TotalsConfig{})
	if err != nil {
		return err
	}

	err = s.refreshAdjustments(cart)
	if err != nil {
		return err
	}
	// err = s.eventBus_.emit(CartService.Events.UPDATED, map[string]interface{}{
	// 	"id": cart.Id,
	// })
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (s *CartService) UpdateLineItem(
	cartId uuid.UUID,
	lineItemId uuid.UUID,
	update *models.LineItem,
	shouldCalculatePrices bool,
) (*models.Cart, *utils.ApplictaionError) {
	selectFields := []string{"id", "region_id", "customerId"}

	feature := true
	if feature {
		selectFields = append(selectFields, "sales_channel_id")
	}
	cart, err := s.Retrieve(cartId, sql.Options{
		Selects:   selectFields,
		Relations: []string{"shipping_methods"},
	}, TotalsConfig{})
	if err != nil {
		return nil, err
	}
	lineItem, err := s.r.LineItemService().SetContext(s.ctx).Retrieve(lineItemId, sql.Options{
		Selects: []string{"id", "quantity", "variant_id", "cart_id"},
	})
	if err != nil {
		return nil, err
	}
	if lineItem.CartId.UUID != cartId {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"A line item with the provided id doesn't exist in the cart",
			"500",
			nil,
		)
	}
	if update.Quantity != 0 {
		if lineItem.VariantId.UUID != uuid.Nil {
			hasInventory, err := s.r.ProductVariantInventoryService().SetContext(s.ctx).ConfirmInventory(lineItem.VariantId.UUID, update.Quantity, map[string]interface{}{
				"salesChannelId": cart.SalesChannelId.UUID,
			})
			if err != nil {
				return nil, err
			}
			if !hasInventory {
				return nil, utils.NewApplictaionError(
					utils.INSUFFICIENT_INVENTORY,
					"Inventory doesn't cover the desired quantity",
					"500",
					nil,
				)
			}
			if shouldCalculatePrices {
				variantsPricing, err := s.r.PricingService().SetContext(s.ctx).GetProductVariantsPricing([]interfaces.Pricing{
					{
						VariantId: lineItem.VariantId.UUID,
						Quantity:  update.Quantity,
					},
				}, &interfaces.PricingContext{
					RegionId:              cart.RegionId.UUID,
					CustomerId:            cart.CustomerId.UUID,
					IncludeDiscountPrices: true,
				})
				if err != nil {
					return nil, err
				}
				calculatedPrice := variantsPricing[lineItem.VariantId.UUID].CalculatedPrice
				update.UnitPrice = calculatedPrice
			}
		}
	}
	if len(cart.ShippingMethods) > 0 {
		err = s.r.ShippingOptionService().SetContext(s.ctx).DeleteShippingMethods(cart.ShippingMethods)
		if err != nil {
			return nil, err
		}
	}
	_, err = s.r.LineItemService().SetContext(s.ctx).Update(lineItemId, nil, update, sql.Options{})
	if err != nil {
		return nil, err
	}
	updatedCart, err := s.Retrieve(cartId, sql.Options{
		Relations: []string{
			"items.variant.product.profiles",
			"discounts",
			"discounts.rule",
			"region",
		},
	}, TotalsConfig{})
	if err != nil {
		return nil, err
	}
	err = s.refreshAdjustments(updatedCart)
	if err != nil {
		return nil, err
	}
	// err = s.eventBus_.emit(CartService.Events.UPDATED, map[string]interface{}{
	// 	"id": updatedCart.Id,
	// })
	// if err != nil {
	// 	return err
	// }
	return updatedCart, nil
}

func (s *CartService) adjustFreeShipping(
	cart *models.Cart,
	shouldAdd bool,
) *utils.ApplictaionError {
	if len(cart.ShippingMethods) > 0 {
		if shouldAdd {
			var shippingMethods []models.ShippingMethod
			for _, shippingMethod := range cart.ShippingMethods {
				shippingMethod.Price = 0
				shippingMethods = append(shippingMethods, shippingMethod)
			}
			if err := s.r.ShippingMethodRepository().UpdateSlice(s.ctx, shippingMethods); err != nil {
				return err
			}
		} else {
			var shippingMethods []models.ShippingMethod
			for _, shippingMethod := range cart.ShippingMethods {
				if shippingMethod.ShippingOption.Amount != 0 {
					price, err := s.r.ShippingOptionService().SetContext(s.ctx).GetPrice(shippingMethod.ShippingOption, shippingMethod.Data, cart)
					if err != nil {
						return err
					}
					shippingMethod.Price = price
				} else {
					shippingMethod.Price = shippingMethod.ShippingOption.Amount
				}
				shippingMethods = append(shippingMethods, shippingMethod)
			}
			if err := s.r.ShippingMethodRepository().UpdateSlice(s.ctx, shippingMethods); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *CartService) Update(id uuid.UUID, cart *models.Cart, data *models.Cart) (*models.Cart, *utils.ApplictaionError) {
	relations := []string{
		"items.variant.product.profiles",
		"shipping_methods",
		"shipping_methods.shipping_option",
		"shipping_address",
		"billing_address",
		"gift_cards",
		"customer",
		"region",
		"payment_sessions",
		"region.countries",
		"discounts",
		"discounts.rule",
	}
	if id == uuid.Nil {
		c, err := s.Retrieve(id, sql.Options{
			Relations: relations,
		}, TotalsConfig{})
		if err != nil {
			return nil, err
		}
		cart = c
	}

	// var originalCartCustomer *models.Customer
	// if cart.Customer != nil {
	// 	originalCartCustomer = cart.Customer
	// }

	if data.CustomerId.UUID != uuid.Nil {
		if err := s.updateCustomerId(cart, data.CustomerId.UUID); err != nil {
			return nil, err
		}
	} else if data.Email != "" {
		customer, err := s.createOrFetchGuestCustomerFromEmail(data.Email)
		if err != nil {
			return nil, err
		}
		cart.Customer = customer
		cart.CustomerId = uuid.NullUUID{UUID: customer.Id}
		cart.Email = customer.Email
	}
	if data.CustomerId.UUID != uuid.Nil || data.RegionId.UUID != uuid.Nil {
		if err := s.updateUnitPrices(cart, data.RegionId.UUID, data.CustomerId.UUID); err != nil {
			return nil, err
		}
	}
	if data.RegionId.UUID != uuid.Nil && cart.RegionId != data.RegionId {
		shippingAddress := data.ShippingAddress
		countryCode := shippingAddress.CountryCode
		if err := s.setRegion(cart, data.RegionId.UUID, countryCode); err != nil {
			return nil, err
		}
	}
	var billingAddress *models.Address
	if data.BillingAddressId.UUID != uuid.Nil {
		billingAddress = &models.Address{Model: core.Model{Id: data.BillingAddressId.UUID}}
	} else if data.BillingAddress != nil {
		billingAddress = data.BillingAddress
	}

	if billingAddress != nil {
		if err := s.updateBillingAddress(cart, uuid.Nil, billingAddress); err != nil {
			return nil, err
		}
	}

	var shippingAddress *models.Address
	if data.ShippingAddressId.UUID != uuid.Nil {
		shippingAddress = &models.Address{Model: core.Model{Id: data.ShippingAddressId.UUID}}
	} else if data.BillingAddress != nil {
		shippingAddress = data.ShippingAddress
	}

	if shippingAddress != nil {
		if err := s.updateShippingAddress(cart, uuid.Nil, shippingAddress); err != nil {
			return nil, err
		}
	}

	feature := false
	featurev2 := false
	if feature && data.SalesChannelId.UUID != uuid.Nil && data.SalesChannelId != cart.SalesChannelId {
		salesChannel, err := s.getValidatedSalesChannel(data.SalesChannelId.UUID)
		if err != nil {
			return nil, err
		}
		err = s.onSalesChannelChange(cart, data.SalesChannelId.UUID)
		if err != nil {
			return nil, err
		}
		// if featurev2 {
		// if cart.SalesChannelId.UUID != uuid.Nil {
		// err = s.remoteLink_.Dismiss(map[string]interface{}{
		// 	"cartService": map[string]interface{}{
		// 		"cart_id": cart.Id,
		// 	},
		// 	"r.SalesChannelService()": map[string]interface{}{
		// 		"sales_channel_id": cart.SalesChannelID,
		// 	},
		// })
		// if err != nil {
		// 	return nil, err
		// }
		// }
		// err = s.remoteLink_.Create(map[string]interface{}{
		// 	"cartService": map[string]interface{}{
		// 		"cart_id": cart.Id,
		// 	},
		// 	"r.SalesChannelService()": map[string]interface{}{
		// 		"sales_channel_id": salesChannel.Id,
		// 	},
		// })
		// if err != nil {
		// 	return nil, err
		// }
		// } else {
		if !featurev2 {
			cart.SalesChannelId = uuid.NullUUID{UUID: salesChannel.Id}
		}
	}
	if data.Discounts != nil && len(data.Discounts) > 0 {
		var previousDiscounts []models.Discount
		copy(previousDiscounts, cart.Discounts)
		cart.Discounts = []models.Discount{}

		var discountCodes []string
		for _, discount := range data.Discounts {
			discountCodes = append(discountCodes, discount.Code)
		}

		if err := s.ApplyDiscounts(cart, discountCodes); err != nil {
			return nil, err
		}
		hasFreeShipping := false
		for _, discount := range cart.Discounts {
			if discount.Rule != nil && discount.Rule.Type == models.DiscountRuleFreeShipping {
				hasFreeShipping = true
				break
			}
		}
		if len(previousDiscounts) > 0 && previousDiscounts[0].Rule.Type == models.DiscountRuleFreeShipping && !hasFreeShipping {
			if err := s.adjustFreeShipping(cart, false); err != nil {
				return nil, err
			}
		}
		if hasFreeShipping {
			if err := s.adjustFreeShipping(cart, true); err != nil {
				return nil, err
			}
		}
	} else if data.Discounts != nil && len(data.Discounts) == 0 {
		cart.Discounts = []models.Discount{}
		if err := s.refreshAdjustments(cart); err != nil {
			return nil, err
		}
	}
	if data.GiftCards != nil {
		cart.GiftCards = []models.GiftCard{}
		for _, gc := range data.GiftCards {
			if err := s.applyGiftCard(cart, gc.Code); err != nil {
				return nil, err
			}
		}
	}

	if data.Context != nil {
		var prevContext core.JSONB
		if cart.Context != nil {
			prevContext = cart.Context
		}

		cart.Context = core.JSONB{}
		for k, v := range prevContext {
			cart.Context[k] = v
		}
		for k, v := range data.Context {
			cart.Context[k] = v
		}
	}
	if data.CompletedAt != nil {
		cart.CompletedAt = data.CompletedAt
	}
	if data.PaymentAuthorizedAt != nil {
		cart.PaymentAuthorizedAt = data.PaymentAuthorizedAt
	}

	if err := s.r.CartRepository().Save(s.ctx, cart); err != nil {
		return nil, err
	}
	// if (data.Email != "" && data.Email != originalCartCustomer.Email) || (data.CustomerId.UUID != uuid.Nil && data.CustomerId.UUID != originalCartCustomer.Id) {
	// err = s.eventBus_.Emit(CartService.Events.CUSTOMER_UPDATED, updatedCart.Id)
	// if err != nil {
	// 	return nil, err
	// }
	// }
	// err = s.eventBus_.Emit(CartService.Events.UPDATED, updatedCart)
	// if err != nil {
	// 	return nil, err
	// }
	return cart, nil
}

func (s *CartService) onSalesChannelChange(cart *models.Cart, newSalesChannelId uuid.UUID) *utils.ApplictaionError {
	_, err := s.getValidatedSalesChannel(newSalesChannelId)
	if err != nil {
		return err
	}
	var productIDs uuid.UUIDs
	for _, item := range cart.Items {
		productIDs = append(productIDs, item.Variant.ProductId.UUID)
	}
	productsToKeep, err := s.r.ProductService().SetContext(s.ctx).FilterProductsBySalesChannel(productIDs, newSalesChannelId, sql.Options{
		Selects: []string{"id"},
		Take:    gox.NewInt(len(productIDs)),
	})
	if err != nil {
		return err
	}
	var productIDsToKeep uuid.UUIDs
	for _, product := range productsToKeep {
		productIDsToKeep = append(productIDsToKeep, product.Id)
	}

	var itemsToRemove []models.LineItem
	for _, item := range cart.Items {
		if !slices.Contains(productIDsToKeep, item.Variant.ProductId.UUID) {
			itemsToRemove = append(itemsToRemove, item)
		}
	}
	if len(itemsToRemove) == 0 {
		return nil
	}
	var itemIDsToRemove uuid.UUIDs
	for _, item := range itemsToRemove {
		itemIDsToRemove = append(itemIDsToRemove, item.Id)
	}
	if err := s.RemoveLineItem(cart.Id, itemIDsToRemove); err != nil {
		return err
	}
	cart.Items = cart.Items[:0]
	for _, item := range cart.Items {
		if !slices.Contains(productIDsToKeep, item.Variant.ProductId.UUID) {
			cart.Items = append(cart.Items, item)
		}
	}

	return nil
}

func (s *CartService) updateCustomerId(cart *models.Cart, customerId uuid.UUID) *utils.ApplictaionError {
	customer, err := s.r.CustomerService().SetContext(s.ctx).RetrieveById(customerId, sql.Options{})
	if err != nil {
		return err
	}
	cart.Customer = customer
	cart.CustomerId = uuid.NullUUID{UUID: customer.Id}
	cart.Email = customer.Email
	return nil
}

func (s *CartService) createOrFetchGuestCustomerFromEmail(email string) (*models.Customer, *utils.ApplictaionError) {
	if err := validator.New().Var(email, "required,email"); err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			err.Error(),
			"500",
			nil,
		)
	}
	customer, err := s.r.CustomerService().SetContext(s.ctx).RetrieveUnregisteredByEmail(email, sql.Options{})
	if err != nil {
		customer, err = s.r.CustomerService().SetContext(s.ctx).Create(&models.Customer{
			Email: email,
		})
		if err != nil {
			return nil, err
		}
	}
	return customer, nil
}

func (s *CartService) updateBillingAddress(cart *models.Cart, id uuid.UUID, address *models.Address) *utils.ApplictaionError {
	if id != uuid.Nil {
		var addr *models.Address

		query := sql.BuildQuery(models.Address{Model: core.Model{Id: id}}, sql.Options{})

		if err := s.r.AddressRepository().FindOne(s.ctx, addr, query); err != nil {
			return err
		}

		address = addr
	}
	if address.Id != uuid.Nil {
		cart.BillingAddress = address
	} else {
		if cart.BillingAddressId.UUID != uuid.Nil {
			var addr *models.Address

			query := sql.BuildQuery(models.Address{Model: core.Model{Id: cart.BillingAddressId.UUID}}, sql.Options{})

			if err := s.r.AddressRepository().FindOne(s.ctx, addr, query); err != nil {
				return err
			}

			copier.CopyWithOption(&address, addr, copier.Option{IgnoreEmpty: true, DeepCopy: true})

			if err := s.r.AddressRepository().Save(s.ctx, address); err != nil {
				return err
			}
		} else {
			cart.BillingAddress = address
		}
	}
	return nil
}

func (s *CartService) updateShippingAddress(cart *models.Cart, id uuid.UUID, address *models.Address) *utils.ApplictaionError {
	if id != uuid.Nil {
		var addr *models.Address

		query := sql.BuildQuery(models.Address{Model: core.Model{Id: id}}, sql.Options{})

		if err := s.r.AddressRepository().FindOne(s.ctx, addr, query); err != nil {
			return err
		}

		address = addr
	}
	if address.CountryCode != "" && !slices.ContainsFunc(cart.Region.Countries, func(c models.Country) bool {
		return address.CountryCode == c.Iso2
	}) {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Shipping country must be in the cart region",
			"500",
			nil,
		)
	}
	if address.Id != uuid.Nil {
		cart.ShippingAddress = address
	} else {
		if cart.ShippingAddressId.UUID != uuid.Nil {
			var addr *models.Address

			query := sql.BuildQuery(models.Address{Model: core.Model{Id: cart.ShippingAddressId.UUID}}, sql.Options{})

			if err := s.r.AddressRepository().FindOne(s.ctx, addr, query); err != nil {
				return err
			}

			copier.CopyWithOption(&address, addr, copier.Option{IgnoreEmpty: true, DeepCopy: true})

			if err := s.r.AddressRepository().Save(s.ctx, address); err != nil {
				return err
			}
		} else {
			cart.ShippingAddress = address
		}
	}
	return nil
}

func (s *CartService) applyGiftCard(cart *models.Cart, code string) *utils.ApplictaionError {
	giftCard, err := s.r.GiftCardService().SetContext(s.ctx).RetrieveByCode(code, sql.Options{})
	if err != nil {
		return err
	}
	if giftCard.IsDisabled {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"The gift card is disabled",
			"500",
			nil,
		)
	}
	if giftCard.RegionId != cart.RegionId {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"The gift card cannot be used in the current region",
			"500",
			nil,
		)
	}
	for _, gc := range cart.GiftCards {
		if gc.Id == giftCard.Id {
			return nil
		}
	}
	cart.GiftCards = append(cart.GiftCards, *giftCard)

	return nil
}

func (s *CartService) ApplyDiscount(cart *models.Cart, discountCode string) *utils.ApplictaionError {
	return s.ApplyDiscounts(cart, []string{discountCode})
}

func (s *CartService) ApplyDiscounts(cart *models.Cart, discountCodes []string) *utils.ApplictaionError {
	discounts, err := s.r.DiscountService().SetContext(s.ctx).ListByCodes(discountCodes, sql.Query{
		Relations: []string{"rule", "rule.conditions", "regions"},
	})
	if err != nil {
		return err
	}
	err = s.r.DiscountService().SetContext(s.ctx).validateDiscountForCartOrThrow(cart, discounts)
	if err != nil {
		return err
	}

	rules := make(map[uuid.UUID]models.DiscountRule)
	discountsMap := make(map[uuid.UUID]models.Discount)
	for _, d := range discounts {
		rules[d.Id] = *d.Rule
		discountsMap[d.Id] = d
	}

	var newDiscounts []models.Discount
	sawNotShipping := false

	var toParse []models.Discount
	for _, v := range discountsMap {
		toParse = append(toParse, v)
	}

	for _, discountToParse := range toParse {
		switch discountToParse.Rule.Type {
		case models.DiscountRuleFreeShipping:
			if discountToParse.Rule.Type == rules[discountToParse.Id].Type {
				newDiscounts = append(newDiscounts, discountsMap[discountToParse.Id])
			} else {
				newDiscounts = append(newDiscounts, discountToParse)
			}
		default:
			if !sawNotShipping {
				sawNotShipping = true
				if rules[discountToParse.Id].Type != models.DiscountRuleFreeShipping {
					newDiscounts = append(newDiscounts, discountsMap[discountToParse.Id])
				} else {
					newDiscounts = append(newDiscounts, discountToParse)
				}
			}
		}
	}

	cart.Discounts = []models.Discount{}
	cart.Discounts = append(cart.Discounts, newDiscounts...)
	hadNonFreeShippingDiscounts := false
	for _, rule := range rules {
		if rule.Type != models.DiscountRuleFreeShipping {
			hadNonFreeShippingDiscounts = true
			break
		}
	}
	if hadNonFreeShippingDiscounts && cart.Items != nil {
		err = s.refreshAdjustments(cart)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *CartService) RemoveDiscount(cartId uuid.UUID, discountCode string) (*models.Cart, *utils.ApplictaionError) {
	cart, err := s.Retrieve(cartId, sql.Options{
		Relations: []string{
			"items.variant.product.profiles",
			"region",
			"discounts",
			"discounts.rule",
			"payment_sessions",
			"shipping_methods",
			"shipping_methods.shipping_option",
		},
	}, TotalsConfig{})
	if err != nil {
		return nil, err
	}
	if cart.Discounts != nil {
		for i, discount := range cart.Discounts {
			if discount.Code == discountCode {
				cart.Discounts = append(cart.Discounts[:i], cart.Discounts[i+1:]...)
				break
			}
		}
	}
	if cart.Discounts != nil && len(cart.Discounts) > 0 {
		err = s.adjustFreeShipping(cart, false)
		if err != nil {
			return nil, err
		}
	}

	if err := s.r.CartRepository().Save(s.ctx, cart); err != nil {
		return nil, err
	}

	err = s.refreshAdjustments(cart)
	if err != nil {
		return nil, err
	}

	if cart.PaymentSessions != nil && len(cart.PaymentSessions) > 0 {
		err = s.setPaymentSessions(cartId, nil)
		if err != nil {
			return nil, err
		}
	}
	// err = s.eventBus_.Emit(CartService.Events.UPDATED, updatedCart)
	// if err != nil {
	// 	return nil, err
	// }
	return cart, nil
}

func (s *CartService) UpdatePaymentSession(cartId uuid.UUID, update map[string]interface{}) (*models.Cart, *utils.ApplictaionError) {
	cart, err := s.Retrieve(cartId, sql.Options{
		Relations: []string{"payment_sessions"},
	}, TotalsConfig{})
	if err != nil {
		return nil, err
	}
	if cart.PaymentSession != nil {
		_, err = s.r.PaymentProviderService().SetContext(s.ctx).UpdateSessionData(cart.PaymentSession, update)
		if err != nil {
			return nil, err
		}
	}
	updatedCart, err := s.Retrieve(cart.Id, sql.Options{}, TotalsConfig{})
	if err != nil {
		return nil, err
	}
	// err = s.eventBus_.Emit(CartService.Events.UPDATED, updatedCart)
	// if err != nil {
	// 	return nil, err
	// }
	return updatedCart, nil
}

func (s *CartService) AuthorizePayment(id uuid.UUID, cart *models.Cart, context map[string]interface{}) (*models.Cart, *utils.ApplictaionError) {
	if id != uuid.Nil {
		context["cart_id"] = id
	} else {
		context["cart_id"] = cart.Id
	}

	if id != uuid.Nil {
		c, err := s.RetrieveWithTotals(id, sql.Options{
			Relations: []string{"payment_sessions", "items.variant.product.profiles"},
		}, TotalsConfig{})
		if err != nil {
			return nil, err
		}
		cart = c
	}
	if cart.Total <= 0 {
		now := time.Now()
		cart.PaymentAuthorizedAt = &now
		if err := s.r.CartRepository().Save(s.ctx, &models.Cart{
			Model:               core.Model{Id: cart.Id},
			PaymentAuthorizedAt: cart.PaymentAuthorizedAt,
		}); err != nil {
			return nil, err
		}
		// err = this.eventBus_.emit(CartService.Events.UPDATED, cart)
		// if err != nil {
		// 	return nil, err
		// }
		return cart, nil
	}
	if cart.PaymentSession == nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"You cannot complete a cart without a payment session.",
			"500",
			nil,
		)
	}
	session, err := s.r.PaymentProviderService().SetContext(s.ctx).AuthorizePayment(cart.PaymentSession, context)
	if err != nil {
		return nil, err
	}
	freshCart, err := s.Retrieve(cart.Id, sql.Options{
		Relations: []string{"payment_sessions"},
	}, TotalsConfig{})
	if err != nil {
		return nil, err
	}
	if session.Status == models.PaymentSessionStatusAuthorized {
		payment, err := s.r.PaymentProviderService().SetContext(s.ctx).CreatePayment(&models.Payment{
			CartId:       uuid.NullUUID{UUID: cart.Id},
			CurrencyCode: cart.Region.CurrencyCode,
			Amount:       cart.Total,
			ProviderId:   freshCart.PaymentSession.ProviderId,
			Data:         freshCart.PaymentSession.Data,
		})
		if err != nil {
			return nil, err
		}
		freshCart.Payment = payment
		now := time.Now()
		freshCart.PaymentAuthorizedAt = &now
	}
	if err := s.r.CartRepository().Save(s.ctx, freshCart); err != nil {
		return nil, err
	}
	// err = this.eventBus_.emit(CartService.Events.UPDATED, updatedCart)
	// if err != nil {
	// 	return nil, err
	// }
	return freshCart, nil
}

func (s *CartService) SetPaymentSession(id uuid.UUID, providerId uuid.UUID) *utils.ApplictaionError {
	cart, err := s.RetrieveWithTotals(id, sql.Options{
		Relations: []string{
			"items.variant.product.profiles",
			"customer",
			"region",
			"region.payment_providers",
			"payment_sessions",
		},
	}, TotalsConfig{})
	if err != nil {
		return err
	}
	isProviderPresent := false
	for _, paymentProvider := range cart.Region.PaymentProviders {
		if paymentProvider.Id == providerId {
			isProviderPresent = true
			break
		}
	}
	if providerId != uuid.Nil && !isProviderPresent {
		return utils.NewApplictaionError(
			utils.NOT_ALLOWED,
			"The payment method is not available in this region",
			"500",
			nil,
		)
	}
	var currentlySelectedSession *models.PaymentSession
	for _, session := range cart.PaymentSessions {
		if session.IsSelected {
			currentlySelectedSession = &session
			break
		}
	}
	if currentlySelectedSession != nil && currentlySelectedSession.ProviderId.UUID != providerId {
		if currentlySelectedSession.IsInitiated {
			if err := s.r.PaymentProviderService().SetContext(s.ctx).DeleteSession(currentlySelectedSession); err != nil {
				return err
			}
		}
		currentlySelectedSession.IsInitiated = false
		currentlySelectedSession.IsSelected = false
		if err := s.r.PaymentSessionRepository().Save(s.ctx, currentlySelectedSession); err != nil {
			return err
		}
	}
	var cartPaymentSessions []models.PaymentSession
	for _, p := range cart.PaymentSessions {
		p.IsSelected = false
		p.IsInitiated = false
		cartPaymentSessions = append(cartPaymentSessions, p)
	}

	if err := s.r.PaymentSessionRepository().UpdateSlice(s.ctx, cartPaymentSessions); err != nil {
		return err
	}
	var paymentSession *models.PaymentSession
	for _, ps := range cart.PaymentSessions {
		if ps.ProviderId.UUID == providerId {
			paymentSession = &ps
			break
		}
	}
	if paymentSession == nil {
		return utils.NewApplictaionError(
			utils.UNEXPECTED_STATE,
			"Could not find payment session",
			"500",
			nil,
		)
	}
	sessionInput := &types.PaymentSessionInput{
		Cart:             *cart,
		Customer:         cart.Customer,
		Amount:           cart.Total,
		CurrencyCode:     cart.Region.CurrencyCode,
		ProviderId:       providerId,
		PaymentSessionId: paymentSession.Id,
	}
	if paymentSession.IsInitiated {
		_, err = s.r.PaymentProviderService().SetContext(s.ctx).UpdateSession(paymentSession, sessionInput)
		if err != nil {
			return err
		}
	} else {
		paymentSession, err = s.r.PaymentProviderService().SetContext(s.ctx).CreateSession(uuid.Nil, sessionInput)
		if err != nil {
			return err
		}
	}
	if err = s.r.PaymentSessionRepository().Update(s.ctx, &models.PaymentSession{
		Model:       core.Model{Id: paymentSession.Id},
		IsSelected:  true,
		IsInitiated: true,
	}); err != nil {
		return err
	}
	// err = this.eventBus_.emit(CartService.Events.UPDATED, map[string]interface{}{
	// 	"id": id,
	// })
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (s *CartService) setPaymentSessions(id uuid.UUID, cart *models.Cart) *utils.ApplictaionError {
	if id != uuid.Nil {
		c, err := s.RetrieveWithTotals(id, sql.Options{
			Relations: []string{
				"items.variant.product.profiles",
				"items.adjustments",
				"discounts",
				"discounts.rule",
				"gift_cards",
				"shipping_methods",
				"shipping_methods.shipping_option",
				"billing_address",
				"shipping_address",
				"region",
				"region.tax_rates",
				"region.payment_providers",
				"payment_sessions",
				"customer",
			},
		}, TotalsConfig{
			ForceTaxes: true,
		})
		if err != nil {
			return err
		}

		cart = c
	}
	total := cart.Total
	region := cart.Region
	deleteSessionAppropriately := func(session *models.PaymentSession) *utils.ApplictaionError {
		if session.IsInitiated {
			return s.r.PaymentProviderService().SetContext(s.ctx).DeleteSession(session)
		}
		if err := s.r.PaymentSessionRepository().Remove(s.ctx, session); err != nil {
			return err
		}

		return nil
	}
	if total <= 0 {
		for _, session := range cart.PaymentSessions {
			if err := deleteSessionAppropriately(&session); err != nil {
				return err
			}
		}
		return nil
	}
	var providerSet uuid.UUIDs
	for _, paymentProvider := range region.PaymentProviders {
		providerSet = append(providerSet, paymentProvider.Id)
	}
	var alreadyConsumedProviderIds uuid.UUIDs
	paymentSessionInput := &types.PaymentSessionInput{
		Cart:         *cart,
		Customer:     cart.Customer,
		Amount:       total,
		CurrencyCode: cart.Region.CurrencyCode,
	}
	paymentSessionData := &models.PaymentSession{
		CartId: uuid.NullUUID{UUID: cart.Id},
		Data:   models.JSONB{},
		Status: models.PaymentSessionStatusPending,
		Amount: total,
	}
	for _, session := range cart.PaymentSessions {
		if !slices.Contains(providerSet, session.ProviderId.UUID) {
			if err := deleteSessionAppropriately(&session); err != nil {
				return err
			}
		}
		if session.IsSelected && session.IsInitiated {
			paymentSessionInput.ProviderId = session.ProviderId.UUID
			_, err := s.r.PaymentProviderService().SetContext(s.ctx).UpdateSession(&session, paymentSessionInput)
			if err != nil {
				return err
			}
		} else {
			if session.IsInitiated {
				if err := s.r.PaymentProviderService().SetContext(s.ctx).DeleteSession(&session); err != nil {
					return err
				}
			} else {
				paymentSessionData = &session
				paymentSessionData.Amount = total
			}
			if err := s.r.PaymentSessionRepository().Save(s.ctx, paymentSessionData); err != nil {
				return err
			}
		}
		alreadyConsumedProviderIds = append(alreadyConsumedProviderIds, session.ProviderId.UUID)
	}
	if len(region.PaymentProviders) == 1 && cart.PaymentSession == nil {
		paymentProvider := region.PaymentProviders[0]
		paymentSessionInput.ProviderId = paymentProvider.Id
		paymentSession, err := s.r.PaymentProviderService().SetContext(s.ctx).CreateSession(uuid.Nil, paymentSessionInput)
		if err != nil {
			return err
		}
		if err := s.r.PaymentSessionRepository().Save(s.ctx, &models.PaymentSession{
			Model:       core.Model{Id: paymentSession.Id},
			IsSelected:  true,
			IsInitiated: true,
		}); err != nil {
			return err
		}

		return nil
	}
	for _, paymentProvider := range region.PaymentProviders {
		if slices.Contains(alreadyConsumedProviderIds, paymentProvider.Id) {
			continue
		}

		paymentSessionData.ProviderId = uuid.NullUUID{UUID: paymentProvider.Id}

		if err := s.r.PaymentSessionRepository().Save(s.ctx, paymentSessionData); err != nil {
			return err
		}
	}
	return nil
}

func (s *CartService) DeletePaymentSession(id uuid.UUID, providerId uuid.UUID) *utils.ApplictaionError {
	cart, err := s.Retrieve(id, sql.Options{
		Relations: []string{"payment_sessions"},
	}, TotalsConfig{})
	if err != nil {
		return err
	}
	if cart.PaymentSession != nil {
		var paymentSession *models.PaymentSession
		for _, session := range cart.PaymentSessions {
			if session.ProviderId.UUID == providerId {
				paymentSession = &session
				break
			}
		}
		var newPaymentSessions []models.PaymentSession
		for _, session := range cart.PaymentSessions {
			if session.ProviderId.UUID != providerId {
				newPaymentSessions = append(newPaymentSessions, session)
			}
		}
		cart.PaymentSessions = newPaymentSessions
		if paymentSession != nil {
			if paymentSession.IsSelected || paymentSession.IsInitiated {
				if err := s.r.PaymentProviderService().SetContext(s.ctx).DeleteSession(paymentSession); err != nil {
					return err
				}
			} else {
				if err := s.r.PaymentSessionRepository().Delete(s.ctx, &models.PaymentSession{
					Model: core.Model{Id: paymentSession.Id},
				}); err != nil {
					return err
				}
			}
		}
	}
	if err := s.r.CartRepository().Save(s.ctx, cart); err != nil {
		return err
	}
	// err = this.eventBus_.emit(CartService.Events.UPDATED, map[string]interface{}{
	// 	"id": cart.id,
	// })
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (s *CartService) RefreshPaymentSession(id uuid.UUID, providerId uuid.UUID) *utils.ApplictaionError {
	cart, err := s.RetrieveWithTotals(id, sql.Options{
		Relations: []string{"payment_sessions"},
	}, TotalsConfig{})
	if err != nil {
		return err
	}
	if cart.PaymentSessions != nil {
		var paymentSession *models.PaymentSession
		for _, session := range cart.PaymentSessions {
			if session.ProviderId.UUID == providerId {
				paymentSession = &session
				break
			}
		}
		if paymentSession != nil {
			if paymentSession.IsSelected {
				_, err = s.r.PaymentProviderService().SetContext(s.ctx).RefreshSession(paymentSession, &types.PaymentSessionInput{
					Cart:         *cart,
					Customer:     cart.Customer,
					Amount:       cart.Total,
					CurrencyCode: cart.Region.CurrencyCode,
					ProviderId:   providerId,
				})
				if err != nil {
					return err
				}
			} else {
				if err := s.r.PaymentSessionRepository().Save(s.ctx, &models.PaymentSession{
					Model:  core.Model{Id: paymentSession.Id},
					Amount: cart.Total,
				}); err != nil {
					return err
				}
			}
		}
	}
	// err = this.eventBus_.emit(CartService.Events.UPDATED, map[string]interface{}{
	// 	"id": id,
	// })
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (s *CartService) AddShippingMethod(id uuid.UUID, cart *models.Cart, optionId uuid.UUID, data map[string]interface{}) (*models.Cart, *utils.ApplictaionError) {
	if id != uuid.Nil {
		c, err := s.RetrieveWithTotals(id, sql.Options{
			Relations: []string{
				"shipping_methods",
				"shipping_methods.shipping_option",
				"items.variant.product.profiles",
				"payment_sessions",
			},
		}, TotalsConfig{})
		if err != nil {
			return nil, err
		}
		cart = c
	}
	cartCustomShippingOptions, err := s.r.CustomShippingOptionService().SetContext(s.ctx).List(models.CustomShippingOption{
		CartId: uuid.NullUUID{UUID: cart.Id},
	}, sql.Options{})
	if err != nil {
		return nil, err
	}

	customShippingOption, err := s.findCustomShippingOption(cartCustomShippingOptions, optionId)
	if err != nil {
		return nil, err
	}
	var shippingMethodConfig *models.ShippingMethod
	if customShippingOption != nil {
		shippingMethodConfig.CartId = uuid.NullUUID{UUID: cart.Id}
		shippingMethodConfig.Price = customShippingOption.Price
	} else {
		shippingMethodConfig.Cart = cart
	}
	newShippingMethod, err := s.r.ShippingOptionService().SetContext(s.ctx).CreateShippingMethod(optionId, data, shippingMethodConfig)
	if err != nil {
		return nil, err
	}
	methods := []models.ShippingMethod{*newShippingMethod}
	if len(cart.ShippingMethods) > 0 {
		for _, shippingMethod := range cart.ShippingMethods {
			if shippingMethod.ShippingOption.ProfileId == newShippingMethod.ShippingOption.ProfileId {
				err := s.r.ShippingOptionService().SetContext(s.ctx).DeleteShippingMethods([]models.ShippingMethod{shippingMethod})
				if err != nil {
					return nil, err
				}
			} else {
				methods = append(methods, shippingMethod)
			}
		}
	}
	if len(cart.Items) > 0 {
		productShippingProfileMap := make(map[uuid.UUID]uuid.UUID)
		featurev2 := false
		if featurev2 {
			var productIds uuid.UUIDs
			for _, item := range cart.Items {
				productIds = append(productIds, item.Variant.ProductId.UUID)
			}

			profileMap, err := s.r.ShippingProfileService().SetContext(s.ctx).GetMapProfileIdsByProductIds(productIds)
			if err != nil {
				return nil, err
			}

			productShippingProfileMap = *profileMap
		} else {
			for _, item := range cart.Items {
				productShippingProfileMap[item.Variant.Product.Id] = item.Variant.Product.ProfileId.UUID
			}
		}
		for _, item := range cart.Items {
			_, err := s.r.LineItemService().SetContext(s.ctx).Update(item.Id, nil, &models.LineItem{
				HasShipping: s.validateLineItemShipping(methods, productShippingProfileMap[item.Variant.ProductId.UUID]),
			}, sql.Options{})
			if err != nil {
				return nil, err
			}
		}
	}
	updatedCart, err := s.Retrieve(cart.Id, sql.Options{
		Relations: []string{
			"discounts",
			"discounts.rule",
			"shipping_methods",
			"shipping_methods.shipping_option",
		},
	}, TotalsConfig{})
	if err != nil {
		return nil, err
	}
	if slices.ContainsFunc(updatedCart.Discounts, func(discount models.Discount) bool {
		return discount.Rule.Type == models.DiscountRuleFreeShipping
	}) {
		if err := s.adjustFreeShipping(updatedCart, true); err != nil {
			return nil, err
		}
	}
	// err := s.eventBus_.emit(CartService.Events.UPDATED, updatedCart)
	// if err != nil {
	// 	return nil, err
	// }
	return updatedCart, nil
}

func (s *CartService) findCustomShippingOption(cartCustomShippingOptions []models.CustomShippingOption, optionId uuid.UUID) (*models.CustomShippingOption, *utils.ApplictaionError) {
	var customOption *models.CustomShippingOption
	for _, cso := range cartCustomShippingOptions {
		if cso.ShippingOptionId.UUID == optionId {
			customOption = &cso
			break
		}
	}
	hasCustomOptions := len(cartCustomShippingOptions) > 0
	if hasCustomOptions && customOption == nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Wrong shipping option",
			"500",
			nil,
		)
	}
	return customOption, nil
}

func (s *CartService) updateUnitPrices(cart *models.Cart, regionId uuid.UUID, customerId uuid.UUID) *utils.ApplictaionError {
	if len(cart.Items) == 0 {
		return nil
	}
	if regionId == uuid.Nil {
		regionId = cart.RegionId.UUID
	}
	if customerId == uuid.Nil {
		customerId = cart.CustomerId.UUID
	}
	region, err := s.r.RegionService().SetContext(s.ctx).Retrieve(regionId, sql.Options{
		Relations: []string{"countries"},
	})
	if err != nil {
		return err
	}

	var calculateVariantPriceData []interfaces.Pricing
	for _, item := range cart.Items {
		if item.VariantId.UUID != uuid.Nil {
			calculateVariantPriceData = append(calculateVariantPriceData, interfaces.Pricing{
				VariantId: item.VariantId.UUID,
				Quantity:  item.Quantity,
			})
		}
	}
	availablePriceMap, err := s.r.PriceSelectionStrategy().CalculateVariantPrice(calculateVariantPriceData, &interfaces.PricingContext{
		RegionId:              region.Id,
		CurrencyCode:          region.CurrencyCode,
		CustomerId:            customerId,
		IncludeDiscountPrices: true,
	})
	if err != nil {
		return err
	}
	for _, item := range cart.Items {
		if item.VariantId.UUID == uuid.Nil {
			continue
		}
		availablePrice, ok := availablePriceMap[item.VariantId.UUID]
		if ok {
			_, err := s.r.LineItemService().SetContext(s.ctx).Update(item.Id, nil, &models.LineItem{
				HasShipping: false,
				UnitPrice:   availablePrice.CalculatedPrice,
			}, sql.Options{})
			if err != nil {
				return err
			}
		} else {
			if err := s.r.LineItemService().SetContext(s.ctx).Delete(item.Id); err != nil {
				return err
			}
		}
	}
	var itemIds uuid.UUIDs
	for _, item := range cart.Items {
		itemIds = append(itemIds, item.Id)
	}
	items, err := s.r.LineItemService().SetContext(s.ctx).List(models.LineItem{}, sql.Options{
		Relations:     []string{"variant.product.profiles"},
		Specification: []sql.Specification{sql.In("id", itemIds)},
	})
	if err != nil {
		return err
	}

	cart.Items = items
	return nil
}

func (s *CartService) setRegion(cart *models.Cart, regionId uuid.UUID, countryCode string) *utils.ApplictaionError {
	if cart.CompletedAt != nil || cart.PaymentAuthorizedAt != nil {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Cannot change the region of a completed cart",
			"500",
			nil,
		)
	}
	region, err := s.r.RegionService().SetContext(s.ctx).Retrieve(regionId, sql.Options{
		Relations: []string{"countries"},
	})
	if err != nil {
		return err
	}
	cart.Region = region
	cart.RegionId = uuid.NullUUID{UUID: region.Id}

	var shippingAddress *models.Address
	if cart.SalesChannelId.UUID != uuid.Nil {
		query := sql.BuildQuery(models.Address{Model: core.Model{Id: cart.ShippingAddressId.UUID}}, sql.Options{})

		if err := s.r.AddressRepository().FindOne(s.ctx, shippingAddress, query); err != nil {
			return err
		}
	}
	if countryCode != "" {
		if !reflect.DeepEqual(shippingAddress, &models.Address{}) && shippingAddress.CountryCode != "" {
			shippingAddress.CountryCode = ""
		}
		if !slices.ContainsFunc(region.Countries, func(country models.Country) bool {
			return country.Iso2 == strings.ToLower(countryCode)
		}) {
			return utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Country not available in region",
				"500",
				nil,
			)
		}
		shippingAddress.CountryCode = strings.ToLower(countryCode)
		if err := s.r.AddressRepository().Save(s.ctx, shippingAddress); err != nil {
			return err
		}
		if err := s.updateShippingAddress(cart, uuid.Nil, shippingAddress); err != nil {
			return err
		}
	} else {
		if !reflect.DeepEqual(shippingAddress, &models.Address{}) && shippingAddress.CountryCode != "" {
			shippingAddress.CountryCode = ""
		}
		if len(region.Countries) == 1 {
			shippingAddress.CountryCode = region.Countries[0].Iso2
		}
		err = s.updateShippingAddress(cart, uuid.Nil, shippingAddress)
		if err != nil {
			return err
		}
	}
	if len(cart.ShippingMethods) > 0 {
		if err := s.r.ShippingOptionService().SetContext(s.ctx).DeleteShippingMethods(cart.ShippingMethods); err != nil {
			return err
		}
	}

	var discountIds uuid.UUIDs
	for _, discount := range cart.Discounts {
		discountIds = append(discountIds, discount.Id)
	}

	discounts, err := s.r.DiscountService().SetContext(s.ctx).List(types.FilterableDiscount{}, sql.Options{
		Relations:     []string{"rule", "regions"},
		Specification: []sql.Specification{sql.In("id", discountIds)},
	})
	if err != nil {
		return err
	}
	cart.Discounts = lo.Filter(discounts, func(discount models.Discount, index int) bool {
		for _, region := range discount.Regions {
			if region.Id == regionId {
				return true
			}
		}
		return false
	})
	if len(cart.Items) > 0 {
		if err := s.refreshAdjustments(cart); err != nil {
			return err
		}
	}
	cart.GiftCards = []models.GiftCard{}
	if len(cart.PaymentSessions) > 0 {
		if err := s.r.PaymentSessionRepository().DeleteSlice(s.ctx, cart.PaymentSessions); err != nil {
			return err
		}
		cart.PaymentSessions = []models.PaymentSession{}
		cart.PaymentSession = &models.PaymentSession{}
	}
	return nil
}

func (s *CartService) Delete(id uuid.UUID) (*models.Cart, *utils.ApplictaionError) {
	cart, err := s.Retrieve(id, sql.Options{
		Relations: []string{
			"items.variant.product.profiles",
			"discounts",
			"discounts.rule",
			"payment_sessions",
		},
	}, TotalsConfig{})
	if err != nil {
		return nil, err
	}
	if cart.CompletedAt != nil {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Completed carts cannot be deleted",
			"500",
			nil,
		)
	}
	if cart.PaymentAuthorizedAt != nil {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Can't delete a cart with an authorized payment",
			"500",
			nil,
		)
	}

	if err := s.r.CartRepository().Remove(s.ctx, cart); err != nil {
		return nil, err
	}
	return cart, nil
}

func (s *CartService) SetMetadata(id uuid.UUID, key string, value string) (*models.Cart, *utils.ApplictaionError) {
	if reflect.TypeOf(key).Kind() != reflect.String {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Key type is invalid. Metadata keys must be strings",
			"500",
			nil,
		)
	}
	cart, err := s.Retrieve(id, sql.Options{}, TotalsConfig{})
	if err != nil {
		return nil, err
	}

	cart.Metadata = cart.Metadata.Add(key, value)

	if err := s.r.CartRepository().Save(s.ctx, cart); err != nil {
		return nil, err
	}
	// err := s.eventBus_.emit(CartService.Events.UPDATED, updatedCart)
	// if err != nil {
	// 	return nil, err
	// }
	return cart, nil
}

func (s *CartService) CreateTaxLines(id uuid.UUID, cart *models.Cart) *utils.ApplictaionError {
	if id != uuid.Nil {
		c, err := s.Retrieve(id, sql.Options{
			Relations: []string{
				"customer",
				"discounts",
				"discounts.rule",
				"gift_cards",
				"items.variant.product.profiles",
				"items.adjustments",
				"region",
				"region.tax_rates",
				"shipping_address",
				"shipping_methods",
				"shipping_methods.shipping_option",
			},
		}, TotalsConfig{})
		if err != nil {
			return err
		}

		cart = c
	}

	calculationContext, err := s.r.TotalsService().GetCalculationContext(cart, nil, CalculationContextOptions{})
	if err != nil {
		return err
	}
	_, _, err = s.r.TaxProviderService().SetContext(s.ctx).CreateTaxLines(cart, nil, calculationContext)
	if err != nil {
		return err
	}
	return nil
}

func (s *CartService) DeleteTaxLines(id uuid.UUID) *utils.ApplictaionError {
	cart, err := s.Retrieve(id, sql.Options{
		Relations: []string{
			"items",
			"items.tax_lines",
			"shipping_methods",
			"shipping_methods.shipping_option",
			"shipping_methods.tax_lines",
		},
	}, TotalsConfig{})
	if err != nil {
		return err
	}

	if err := s.r.LineItemRepository().RemoveSlice(s.ctx, cart.Items); err != nil {
		return err
	}

	if err := s.r.ShippingMethodRepository().RemoveSlice(s.ctx, cart.ShippingMethods); err != nil {
		return err
	}

	return nil
}

func (s *CartService) DecorateTotals(cart *models.Cart, totalsConfig TotalsConfig) (*models.Cart, *utils.ApplictaionError) {
	calculationContext, err := s.r.TotalsService().GetCalculationContext(cart, nil, CalculationContextOptions{})
	if err != nil {
		return nil, err
	}
	includeTax := cart.Region.AutomaticTaxes
	if totalsConfig.ForceTaxes {
		includeTax = totalsConfig.ForceTaxes
	}
	cartItems := append([]models.LineItem{}, cart.Items...)
	cartShippingMethods := append([]models.ShippingMethod{}, cart.ShippingMethods...)
	if includeTax {
		taxLinesMaps, err := s.r.TaxProviderService().SetContext(s.ctx).GetTaxLinesMap(cartItems, calculationContext)
		if err != nil {
			return nil, err
		}
		for _, item := range cartItems {
			if item.IsReturn {
				continue
			}
			item.TaxLines = taxLinesMaps.LineItemsTaxLines[item.Id]
		}
		for _, method := range cartShippingMethods {
			method.TaxLines = taxLinesMaps.ShippingMethodsTaxLines[method.Id]
		}
	}
	itemsTotals, err := s.r.NewTotalsService().SetContext(s.ctx).GetLineItemTotals(cartItems, includeTax, calculationContext, nil)
	if err != nil {
		return nil, err
	}
	shippingTotals, err := s.r.NewTotalsService().SetContext(s.ctx).GetShippingMethodTotals(cartShippingMethods, includeTax, cart.Discounts, nil, calculationContext)
	if err != nil {
		return nil, err
	}
	cart.Subtotal = 0
	cart.DiscountTotal = 0
	cart.ItemTaxTotal = 0
	cart.ShippingTotal = 0
	cart.ShippingTaxTotal = 0
	for _, item := range cart.Items {
		itemWithTotals := item
		total, ok := itemsTotals[item.Id]
		if ok {
			itemWithTotals = total
		}
		cart.Subtotal += itemWithTotals.Subtotal
		cart.DiscountTotal += itemWithTotals.RawDiscountTotal
		cart.ItemTaxTotal += itemWithTotals.TaxTotal
		item = itemWithTotals
	}
	for _, method := range cart.ShippingMethods {
		methodWithTotals := method
		total, ok := shippingTotals[method.Id]
		if ok {
			methodWithTotals = total
		}
		cart.ShippingTotal += methodWithTotals.Subtotal
		cart.ShippingTaxTotal += methodWithTotals.TaxTotal
		method = methodWithTotals
	}
	cart.TaxTotal = cart.ItemTaxTotal + cart.ShippingTaxTotal
	cart.RawDiscountTotal = cart.DiscountTotal
	cart.DiscountTotal = math.Round(cart.DiscountTotal)
	giftCardableAmount := s.r.NewTotalsService().SetContext(s.ctx).GetGiftCardableAmount(cart.Region.GiftCardsTaxable, cart.Subtotal, cart.ShippingTotal, cart.DiscountTotal, cart.TaxTotal)
	giftCardTotal, err := s.r.NewTotalsService().SetContext(s.ctx).GetGiftCardTotals(giftCardableAmount, nil, cart.Region, cart.GiftCards)
	if err != nil {
		return nil, err
	}
	cart.GiftCardTotal = giftCardTotal.Total
	cart.GiftCardTaxTotal = giftCardTotal.TaxTotal
	cart.Total = cart.Subtotal + cart.ShippingTotal + cart.TaxTotal - (cart.GiftCardTaxTotal + cart.DiscountTotal + cart.GiftCardTaxTotal)
	return cart, nil
}

func (s *CartService) refreshAdjustments(cart *models.Cart) *utils.ApplictaionError {
	var nonReturnLineIds uuid.UUIDs
	for _, item := range cart.Items {
		if !item.IsReturn {
			nonReturnLineIds = append(nonReturnLineIds, item.Id)
		}
	}

	err := s.r.LineItemAdjustmentService().SetContext(s.ctx).Delete(uuid.Nil, &models.LineItemAdjustment{}, sql.Options{
		Specification: []sql.Specification{
			sql.In("item_id", nonReturnLineIds),
			sql.Not(sql.IsNull("discount_id")),
		},
	})
	if err != nil {
		return err
	}

	_, _, err = s.r.LineItemAdjustmentService().SetContext(s.ctx).CreateAdjustments(cart, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *CartService) transformQueryForTotals(config sql.Options) ([]string, []string, []types.TotalField) {
	selectFields := config.Selects
	relations := config.Relations

	if selectFields == nil {
		return selectFields, relations, []types.TotalField{}
	}

	totalFields := []types.TotalField{
		types.TotalFieldSubtotal,
		types.TotalFieldTaxTotal,
		types.TotalFieldShippingTotal,
		types.TotalFieldDiscountTotal,
		types.TotalFieldGiftCardTotal,
		types.TotalFieldTotal,
	}

	var totalsToSelect []types.TotalField
	for _, v := range selectFields {
		if slices.Contains(totalFields, types.TotalField(v)) {
			totalsToSelect = append(totalsToSelect, types.TotalField(v))
		}
	}

	if len(totalsToSelect) > 0 {
		relationSet := []string{
			"items",
			"items.tax_lines",
			"gift_cards",
			"discount,s",
			"discounts.rule",
			"shipping_methods",
			"shipping_address",
			"region",
			"region.tax_rates",
		}

		for _, relation := range config.Relations {
			if slices.Contains(relationSet, relation) {
				relationSet = append(relationSet, relation)
			}
		}

		selectFields = lo.Filter(selectFields, func(v string, index int) bool {
			return !slices.Contains(totalFields, types.TotalField(v))
		})

		return selectFields, relationSet, totalsToSelect
	}

	return selectFields, relations, totalsToSelect
}

func (s *CartService) decorateTotals(cart *models.Cart, totalsToSelect []types.TotalField, config TotalsConfig) (*models.Cart, *utils.ApplictaionError) {
	totals := make(map[types.TotalField]float64)

	for _, key := range totalsToSelect {
		switch key {
		case types.TotalFieldTotal:
			res, err := s.r.TotalsService().GetTotal(cart, nil, GetTotalsOptions{ForceTaxes: config.ForceTaxes})
			if err != nil {
				return nil, err
			}
			totals[types.TotalFieldTotal] = res
		case types.TotalFieldShippingTotal:
			res, err := s.r.TotalsService().GetShippingTotal(cart, nil)
			if err != nil {
				return nil, err
			}
			totals[types.TotalFieldShippingTotal] = res
		case types.TotalFieldDiscountTotal:
			res, err := s.r.TotalsService().GetDiscountTotal(cart, nil)
			if err != nil {
				return nil, err
			}
			totals[types.TotalFieldDiscountTotal] = res
		case types.TotalFieldTaxTotal:
			res, err := s.r.TotalsService().GetTaxTotal(cart, nil, config.ForceTaxes)
			if err != nil {
				return nil, err
			}
			totals[types.TotalFieldTaxTotal] = res
		case types.TotalFieldGiftCardTotal:
			giftCardBreakdown, err := s.r.TotalsService().GetGiftCardTotal(cart, nil, map[string]interface{}{})
			if err != nil {
				return nil, err
			}
			totals[types.TotalFieldGiftCardTotal] = giftCardBreakdown.Total
			totals[types.TotalFieldGiftCardTaxTotal] = giftCardBreakdown.TaxTotal
		case types.TotalFieldSubtotal:
			res, err := s.r.TotalsService().GetSubtotal(cart, nil, types.SubtotalOptions{})
			if err != nil {
				return nil, err
			}
			totals[types.TotalFieldSubtotal] = res
		}
	}

	for k, v := range totals {
		switch k {
		case types.TotalFieldShippingTotal:
			cart.ShippingTotal = v
		case types.TotalFieldDiscountTotal:
			cart.DiscountTotal = v
		case types.TotalFieldTaxTotal:
			cart.TaxTotal = v
		case types.TotalFieldRefundedTotal:
			cart.RefundedTotal = v
		case types.TotalFieldTotal:
			cart.Total = v
		case types.TotalFieldSubtotal:
			cart.Subtotal = v
		case types.TotalFieldRefundableAmount:
			cart.RefundableAmount = v
		case types.TotalFieldGiftCardTotal:
			cart.GiftCardTotal = v
		case types.TotalFieldGiftCardTaxTotal:
			cart.GiftCardTaxTotal = v
		}
	}

	return cart, nil
}

func (s *CartService) getTotalsRelations(config sql.Options) []string {
	relationSet := []string{
		"items.variant.product.profiles",
		"items.tax_lines",
		"items.adjustments",
		"gift_cards",
		"discounts",
		"discounts.rule",
		"shipping_methods",
		"shipping_methods.tax_lines",
		"shipping_address",
		"region",
		"region.tax_rates",
	}

	for _, relation := range config.Relations {
		if slices.Contains(relationSet, relation) {
			relationSet = append(relationSet, relation)
		}
	}

	return relationSet
}

func removeDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
