package registry

import (
	"github.com/driver005/gateway/services"
)

func (m *Base) AnalyticsConfigService() *services.AnalyticsConfigService {
	if m.analyticsConfigService == nil {
		m.analyticsConfigService = services.NewAnalyticsConfigService(m)
	}
	return m.analyticsConfigService
}
func (m *Base) AuthService() *services.AuthService {
	if m.authService == nil {
		m.authService = services.NewAuthService(m)
	}
	return m.authService
}
func (m *Base) BatchJobService() *services.BatchJobService {
	if m.batchJobService == nil {
		m.batchJobService = services.NewBatchJobService(
			m.Container(),
			m,
		)
	}
	return m.batchJobService
}
func (m *Base) CartService() *services.CartService {
	if m.cartService == nil {
		m.cartService = services.NewCartService(m)
	}
	return m.cartService
}
func (m *Base) ClaimItemService() *services.ClaimItemService {
	if m.claimItemService == nil {
		m.claimItemService = services.NewClaimItemService(m)
	}
	return m.claimItemService
}
func (m *Base) ClaimService() *services.ClaimService {
	if m.claimService == nil {
		m.claimService = services.NewClaimService(m)
	}
	return m.claimService
}
func (m *Base) CsvParserService() *services.CsvParserService {
	if m.csvParserService == nil {
		m.csvParserService = services.NewCsvParserService(
			services.CsvSchema{},
			"",
		)
	}
	return m.csvParserService
}
func (m *Base) CurrencyService() *services.CurrencyService {
	if m.currencyService == nil {
		m.currencyService = services.NewCurrencyService(m)
	}
	return m.currencyService
}
func (m *Base) CustomShippingOptionService() *services.CustomShippingOptionService {
	if m.customShippingOptionService == nil {
		m.customShippingOptionService = services.NewCustomShippingOptionService(m)
	}
	return m.customShippingOptionService
}
func (m *Base) CustomerGroupService() *services.CustomerGroupService {
	if m.customerGroupService == nil {
		m.customerGroupService = services.NewCustomerGroupService(m)
	}
	return m.customerGroupService
}
func (m *Base) CustomerService() *services.CustomerService {
	if m.customerService == nil {
		m.customerService = services.NewCustomerService(m)
	}
	return m.customerService
}
func (m *Base) DiscountConditionService() *services.DiscountConditionService {
	if m.discountConditionService == nil {
		m.discountConditionService = services.NewDiscountConditionService(m)
	}
	return m.discountConditionService
}
func (m *Base) DiscountService() *services.DiscountService {
	if m.discountService == nil {
		m.discountService = services.NewDiscountService(m)
	}
	return m.discountService
}
func (m *Base) DraftOrderService() *services.DraftOrderService {
	if m.draftOrderService == nil {
		m.draftOrderService = services.NewDraftOrderService(m)
	}
	return m.draftOrderService
}
func (m *Base) EventBus() *services.Bus {
	if m.eventBusService == nil {
		new := services.New()
		m.eventBusService = &new
	}
	return m.eventBusService
}
func (m *Base) DefaultFileService() *services.DefaultFileService {
	if m.defaultFileService == nil {
		m.defaultFileService = services.NewDefaultFileService()
	}
	return m.defaultFileService
}
func (m *Base) FulfillmentProviderService() *services.FulfillmentProviderService {
	if m.fulfillmentProviderService == nil {
		m.fulfillmentProviderService = services.NewFulfillmentProviderService(
			m.Container(),
			m,
		)
	}
	return m.fulfillmentProviderService
}
func (m *Base) FulfillmentService() *services.FulfillmentService {
	if m.fulfillmentService == nil {
		m.fulfillmentService = services.NewFulfillmentService(m)
	}
	return m.fulfillmentService
}
func (m *Base) GiftCardService() *services.GiftCardService {
	if m.giftCardService == nil {
		m.giftCardService = services.NewGiftCardService(m)
	}
	return m.giftCardService
}
func (m *Base) IdempotencyKeyService() *services.IdempotencyKeyService {
	if m.idempotencyKeyService == nil {
		m.idempotencyKeyService = services.NewIdempotencyKeyService(m)
	}
	return m.idempotencyKeyService
}
func (m *Base) InviteService() *services.InviteService {
	if m.inviteService == nil {
		m.inviteService = services.NewInviteService(m)
	}
	return m.inviteService
}
func (m *Base) LineItemAdjustmentService() *services.LineItemAdjustmentService {
	if m.lineItemAdjustmentService == nil {
		m.lineItemAdjustmentService = services.NewLineItemAdjustmentService(m)
	}
	return m.lineItemAdjustmentService
}
func (m *Base) LineItemService() *services.LineItemService {
	if m.lineItemService == nil {
		m.lineItemService = services.NewLineItemService(m)
	}
	return m.lineItemService
}
func (m *Base) NewTotalsService() *services.NewTotalsService {
	if m.newTotalsService == nil {
		m.newTotalsService = services.NewNewTotalsServices(m)
	}
	return m.newTotalsService
}
func (m *Base) NoteService() *services.NoteService {
	if m.noteService == nil {
		m.noteService = services.NewNoteService(m)
	}
	return m.noteService
}
func (m *Base) NotificationService() *services.NotificationService {
	if m.notificationService == nil {
		m.notificationService = services.NewNotificationService(
			m.Container(),
			nil,
			nil,
			m,
		)
	}
	return m.notificationService
}
func (m *Base) OAuthService() *services.OAuthService {
	if m.oAuthService == nil {
		m.oAuthService = services.NewOAuthService(
			m.Container(),
			m,
		)
	}
	return m.oAuthService
}
func (m *Base) OrderItemChangeService() *services.OrderItemChangeService {
	if m.orderItemChangeService == nil {
		m.orderItemChangeService = services.NewOrderItemChangeService(m)
	}
	return m.orderItemChangeService
}
func (m *Base) OrderEditService() *services.OrderEditService {
	if m.orderEditService == nil {
		m.orderEditService = services.NewOrderEditService(m)
	}
	return m.orderEditService
}
func (m *Base) OrderService() *services.OrderService {
	if m.orderService == nil {
		m.orderService = services.NewOrderService(m)
	}
	return m.orderService
}
func (m *Base) PaymentCollectionService() *services.PaymentCollectionService {
	if m.paymentCollectionService == nil {
		m.paymentCollectionService = services.NewPaymentCollectionService(m)
	}
	return m.paymentCollectionService
}
func (m *Base) PaymentProviderService() *services.PaymentProviderService {
	if m.paymentProviderService == nil {
		m.paymentProviderService = services.NewPaymentProviderService(
			m.Container(),
			m,
		)
	}
	return m.paymentProviderService
}
func (m *Base) PaymentService() *services.PaymentService {
	if m.paymentService == nil {
		m.paymentService = services.NewPaymentService(m)
	}
	return m.paymentService
}
func (m *Base) PriceListService() *services.PriceListService {
	if m.priceListService == nil {
		m.priceListService = services.NewPriceListService(m)
	}
	return m.priceListService
}
func (m *Base) PricingService() *services.PricingService {
	if m.pricingService == nil {
		m.pricingService = services.NewPricingService(m)
	}
	return m.pricingService
}
func (m *Base) ProductCategoryService() *services.ProductCategoryService {
	if m.productCategoryService == nil {
		m.productCategoryService = services.NewProductCategoryService(m)
	}
	return m.productCategoryService
}
func (m *Base) ProductCollectionService() *services.ProductCollectionService {
	if m.productCollectionService == nil {
		m.productCollectionService = services.NewProductCollectionService(m)
	}
	return m.productCollectionService
}
func (m *Base) ProductTagService() *services.ProductTagService {
	if m.productTagService == nil {
		m.productTagService = services.NewProductTagService(m)
	}
	return m.productTagService
}
func (m *Base) ProductTaxRateService() *services.ProductTaxRateService {
	if m.productTaxRateService == nil {
		m.productTaxRateService = services.NewProductTaxRateService(m)
	}
	return m.productTaxRateService
}
func (m *Base) ProductTypeService() *services.ProductTypeService {
	if m.productTypeService == nil {
		m.productTypeService = services.NewProductTypeService(m)
	}
	return m.productTypeService
}
func (m *Base) ProductVariantInventoryService() *services.ProductVariantInventoryService {
	if m.productVariantInventoryService == nil {
		m.productVariantInventoryService = services.NewProductVariantInventoryService(m)
	}
	return m.productVariantInventoryService
}
func (m *Base) ProductVariantService() *services.ProductVariantService {
	if m.productVariantService == nil {
		m.productVariantService = services.NewProductVariantService(m)
	}
	return m.productVariantService
}
func (m *Base) ProductService() *services.ProductService {
	if m.productService == nil {
		m.productService = services.NewProductService(m)
	}
	return m.productService
}
func (m *Base) PublishableApiKeyService() *services.PublishableApiKeyService {
	if m.publishableApiKeyService == nil {
		m.publishableApiKeyService = services.NewPublishableApiKeyService(m)
	}
	return m.publishableApiKeyService
}
func (m *Base) RegionService() *services.RegionService {
	if m.regionService == nil {
		m.regionService = services.NewRegionService(m)
	}
	return m.regionService
}
func (m *Base) ReturnReasonService() *services.ReturnReasonService {
	if m.returnReasonService == nil {
		m.returnReasonService = services.NewReturnReasonService(m)
	}
	return m.returnReasonService
}
func (m *Base) ReturnService() *services.ReturnService {
	if m.returnedService == nil {
		m.returnedService = services.NewReturnService(m)
	}
	return m.returnedService
}
func (m *Base) SalesChannelInventoryService() *services.SalesChannelInventoryService {
	if m.salesChannelInventoryService == nil {
		m.salesChannelInventoryService = services.NewSalesChannelInventoryService(m)
	}
	return m.salesChannelInventoryService
}
func (m *Base) SalesChannelLocationService() *services.SalesChannelLocationService {
	if m.salesChannelLocationService == nil {
		m.salesChannelLocationService = services.NewSalesChannelLocationService(m)
	}
	return m.salesChannelLocationService
}
func (m *Base) SalesChannelService() *services.SalesChannelService {
	if m.salesChannelService == nil {
		m.salesChannelService = services.NewSalesChannelService(m)
	}
	return m.salesChannelService
}
func (m *Base) DefaultSearchService() *services.DefaultSearchService {
	if m.defaultSearchService == nil {
		m.defaultSearchService = services.NewDefaultSearchService(m)
	}
	return m.defaultSearchService
}
func (m *Base) ShippingOptionService() *services.ShippingOptionService {
	if m.shippingOptionService == nil {
		m.shippingOptionService = services.NewShippingOptionService(m)
	}
	return m.shippingOptionService
}
func (m *Base) ShippingProfileService() *services.ShippingProfileService {
	if m.shippingProfileService == nil {
		m.shippingProfileService = services.NewShippingProfileService(m)
	}
	return m.shippingProfileService
}
func (m *Base) ShippingTaxRateService() *services.ShippingTaxRateService {
	if m.shippingTaxRateService == nil {
		m.shippingTaxRateService = services.NewShippingTaxRateService(m)
	}
	return m.shippingTaxRateService
}
func (m *Base) StagedJobService() *services.StagedJobService {
	if m.stagedJobService == nil {
		m.stagedJobService = services.NewStagedJobService(m)
	}
	return m.stagedJobService
}
func (m *Base) StoreService() *services.StoreService {
	if m.storeService == nil {
		m.storeService = services.NewStoreService(m)
	}
	return m.storeService
}
func (m *Base) StrategyResolverService() *services.StrategyResolverService {
	if m.strategyResolverService == nil {
		m.strategyResolverService = services.NewStrategyResolverService()
	}
	return m.strategyResolverService
}
func (m *Base) SwapService() *services.SwapService {
	if m.swapService == nil {
		m.swapService = services.NewSwapService(m)
	}
	return m.swapService
}
func (m *Base) SystemProviderService() *services.SystemProviderService {
	if m.systemProviderService == nil {
		m.systemProviderService = services.NewSystemProviderService()
	}
	return m.systemProviderService
}
func (m *Base) SystemTaxService() *services.SystemTaxService {
	if m.systemTaxService == nil {
		m.systemTaxService = services.NewSystemTaxService(m)
	}
	return m.systemTaxService
}
func (m *Base) TaxProviderService() *services.TaxProviderService {
	if m.taxProviderService == nil {
		m.taxProviderService = services.NewTaxProviderService(
			m.Container(),
			m,
		)
	}
	return m.taxProviderService
}
func (m *Base) TaxRateService() *services.TaxRateService {
	if m.taxRateService == nil {
		m.taxRateService = services.NewTaxRateService(m)
	}
	return m.taxRateService
}
func (m *Base) TockenService() *services.TockenService {
	if m.tockenService == nil {
		m.tockenService = services.NewTockenService(
			[]byte(m.Config().Secrets.JwtSecret),
		)
	}
	return m.tockenService
}
func (m *Base) TotalsService() *services.TotalsService {
	if m.totalsService == nil {
		m.totalsService = services.NewTotalService(m)
	}
	return m.totalsService
}
func (m *Base) UserService() *services.UserService {
	if m.userService == nil {
		m.userService = services.NewUserService(m)
	}
	return m.userService
}

func (m *Base) FlagRouter() *services.FlagRouter {
	if m.flagRouter == nil {
		m.flagRouter = services.NewFlagRouter(m)
	}
	return m.flagRouter
}
