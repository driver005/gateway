package admin

import (
	"github.com/driver005/gateway/middlewares"
	"github.com/driver005/gateway/services"
	"github.com/gofiber/fiber/v3/middleware/session"
)

type Registry interface {
	Session() *session.Store

	Middleware() *middlewares.Handler

	AnalyticsConfigService() *services.AnalyticsConfigService
	AuthService() *services.AuthService
	BatchJobService() *services.BatchJobService
	CartService() *services.CartService
	ClaimItemService() *services.ClaimItemService
	ClaimService() *services.ClaimService
	CsvParserService() *services.CsvParserService
	CurrencyService() *services.CurrencyService
	CustomShippingOptionService() *services.CustomShippingOptionService
	CustomerGroupService() *services.CustomerGroupService
	CustomerService() *services.CustomerService
	DiscountConditionService() *services.DiscountConditionService
	DiscountService() *services.DiscountService
	DraftOrderService() *services.DraftOrderService
	EventBus() *services.Bus
	DefaultFileService() *services.DefaultFileService
	FulfillmentProviderService() *services.FulfillmentProviderService
	FulfillmentService() *services.FulfillmentService
	GiftCardService() *services.GiftCardService
	IdempotencyKeyService() *services.IdempotencyKeyService
	InviteService() *services.InviteService
	LineItemAdjustmentService() *services.LineItemAdjustmentService
	LineItemService() *services.LineItemService
	NewTotalsService() *services.NewTotalsService
	NoteService() *services.NoteService
	NotificationService() *services.NotificationService
	OAuthService() *services.OAuthService
	OrderItemChangeService() *services.OrderItemChangeService
	OrderEditService() *services.OrderEditService
	OrderService() *services.OrderService
	PaymentCollectionService() *services.PaymentCollectionService
	PaymentProviderService() *services.PaymentProviderService
	PaymentService() *services.PaymentService
	PriceListService() *services.PriceListService
	PricingService() *services.PricingService
	ProductCategoryService() *services.ProductCategoryService
	ProductCollectionService() *services.ProductCollectionService
	ProductTagService() *services.ProductTagService
	ProductTaxRateService() *services.ProductTaxRateService
	ProductTypeService() *services.ProductTypeService
	ProductVariantInventoryService() *services.ProductVariantInventoryService
	ProductVariantService() *services.ProductVariantService
	ProductService() *services.ProductService
	PublishableApiKeyService() *services.PublishableApiKeyService
	RegionService() *services.RegionService
	ReturnReasonService() *services.ReturnReasonService
	ReturnService() *services.ReturnService
	SalesChannelInventoryService() *services.SalesChannelInventoryService
	SalesChannelLocationService() *services.SalesChannelLocationService
	SalesChannelService() *services.SalesChannelService
	DefaultSearchService() *services.DefaultSearchService
	ShippingOptionService() *services.ShippingOptionService
	ShippingProfileService() *services.ShippingProfileService
	ShippingTaxRateService() *services.ShippingTaxRateService
	StagedJobService() *services.StagedJobService
	StoreService() *services.StoreService
	StrategyResolverService() *services.StrategyResolverService
	SwapService() *services.SwapService
	SystemProviderService() *services.SystemProviderService
	SystemTaxService() *services.SystemTaxService
	TaxProviderService() *services.TaxProviderService
	TaxRateService() *services.TaxRateService
	TockenService() *services.TockenService
	TotalsService() *services.TotalsService
	UserService() *services.UserService
}
