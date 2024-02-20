package strategies

import (
	"context"

	"github.com/driver005/gateway/config"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/repository"
	"github.com/driver005/gateway/services"
	"gorm.io/gorm"
)

type Registry interface {
	Context() *gorm.DB
	Manager(ctx context.Context) *gorm.DB
	Config() *config.Config

	//Repository
	AddressRepository() *repository.AddressRepo
	AnalyticsConfigRepository() *repository.AnalyticsConfigRepo
	BatchJobRepository() *repository.BatchJobRepo
	CartRepository() *repository.CartRepo
	ClaimImageRepository() *repository.ClaimImageRepo
	ClaimItemRepository() *repository.ClaimItemRepo
	ClaimTagRepository() *repository.ClaimTagRepo
	ClaimRepository() *repository.ClaimRepo
	CountryRepository() *repository.CountryRepo
	CurrencyRepository() *repository.CurrencyRepo
	CustomShippingOptionRepository() *repository.CustomShippingOptionRepo
	CustomerGroupRepository() *repository.CustomerGroupRepo
	CustomerRepository() *repository.CustomerRepo
	DiscountConditionRepository() *repository.DiscountConditionRepo
	DiscountRuleRepository() *repository.DiscountRuleRepo
	DiscountRepository() *repository.DiscountRepo
	DraftOrderRepository() *repository.DraftOrderRepo
	FulfillmentProviderRepository() *repository.FulfillmentProviderRepo
	FulfillmentRepository() *repository.FulfillmentRepo
	GiftCardTransactionRepository() *repository.GiftCardTransactionRepo
	GiftCardRepository() *repository.GiftCardRepo
	IdempotencyKeyRepository() *repository.IdempotencyKeyRepo
	ImageRepository() *repository.ImageRepo
	InviteRepository() *repository.InviteRepo
	LineItemAdjustmentRepository() *repository.LineItemAdjustmentRepo
	LineItemTaxLineRepository() *repository.LineItemTaxLineRepo
	LineItemRepository() *repository.LineItemRepo
	MoneyAmountRepository() *repository.MoneyAmountRepo
	NoteRepository() *repository.NoteRepo
	NotificationProviderRepository() *repository.NotificationProviderRepo
	NotificationRepository() *repository.NotificationRepo
	OAuthRepository() *repository.OAuthRepo
	OrderEditRepository() *repository.OrderEditRepo
	OrderItemChangeRepository() *repository.OrderItemChangeRepo
	OrderRepository() *repository.OrderRepo
	PaymentCollectionRepository() *repository.PaymentCollectionRepo
	PaymentProviderRepository() *repository.PaymentProviderRepo
	PaymentSessionRepository() *repository.PaymentSessionRepo
	PaymentRepository() *repository.PaymentRepo
	PriceListRepository() *repository.PriceListRepo
	ProductCategoryRepository() *repository.ProductCategoryRepo
	ProductCollectionRepository() *repository.ProductCollectionRepo
	ProductOptionValueRepository() *repository.ProductOptionValueRepo
	ProductOptionRepository() *repository.ProductOptionRepo
	ProductTagRepository() *repository.ProductTagRepo
	ProductTaxRateRepository() *repository.ProductTaxRateRepo
	ProductTypeRepository() *repository.ProductTypeRepo
	ProductVariantInventoryItemRepository() *repository.ProductVariantInventoryItem
	ProductVariantRepository() *repository.ProductVariantRepo
	ProductRepository() *repository.ProductRepo
	PublishableApiKeySalesChannelRepository() *repository.PublishableApiKeySalesChannelRepo
	PublishableApiKeyRepository() *repository.PublishableApiKeyRepo
	RefundRepository() *repository.RefundRepo
	RegionRepository() *repository.RegionRepo
	ReturnItemRepository() *repository.ReturnItemRepo
	ReturnReasonRepository() *repository.ReturnReasonRepo
	ReturnRepository() *repository.ReturnRepo
	SalesChannelLocationRepository() *repository.SalesChannelLocationRepo
	SalesChannelRepository() *repository.SalesChannelRepo
	ShippingMethodTaxLineRepository() *repository.ShippingMethodTaxLineRepo
	ShippingMethodRepository() *repository.ShippingMethodRepo
	ShippingOptionRequirementRepository() *repository.ShippingOptionRequirementRepo
	ShippingOptionRepository() *repository.ShippingOptionRepo
	ShippingProfileRepository() *repository.ShippingProfileRepo
	ShippingTaxRateRepository() *repository.ShippingTaxRateRepo
	StagedJobRepository() *repository.StagedJobRepo
	StoreRepository() *repository.StoreRepo
	SwapRepository() *repository.SwapRepo
	TaxProviderRepository() *repository.TaxProviderRepo
	TaxRateRepository() *repository.TaxRateRepo
	TrackingLinkRepository() *repository.TrackingLinkRepo
	UserRepository() *repository.UserRepo

	//Services
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

	//Interfaces
	PriceSelectionStrategy() interfaces.IPriceSelectionStrategy
	TaxCalculationStrategy() interfaces.ITaxCalculationStrategy
	InventoryService() interfaces.IInventoryService
	StockLocationService() interfaces.IStockLocationService
	CacheService() interfaces.ICacheService
	PricingModuleService() interfaces.IPricingModuleService
	BatchJobStrategy() interfaces.IBatchJobStrategy
}
