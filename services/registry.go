package services

import (
	"context"

	"github.com/driver005/gateway/config"
	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/repository"
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
	AnalyticsConfigService() *AnalyticsConfigService
	AuthService() *AuthService
	BatchJobService() *BatchJobService
	CartService() *CartService
	ClaimItemService() *ClaimItemService
	ClaimService() *ClaimService
	CsvParserService() *CsvParserService
	CurrencyService() *CurrencyService
	CustomShippingOptionService() *CustomShippingOptionService
	CustomerGroupService() *CustomerGroupService
	CustomerService() *CustomerService
	DiscountConditionService() *DiscountConditionService
	DiscountService() *DiscountService
	DraftOrderService() *DraftOrderService
	EventBus() *Bus
	DefaultFileService() *DefaultFileService
	FulfillmentProviderService() *FulfillmentProviderService
	FulfillmentService() *FulfillmentService
	GiftCardService() *GiftCardService
	IdempotencyKeyService() *IdempotencyKeyService
	InviteService() *InviteService
	LineItemAdjustmentService() *LineItemAdjustmentService
	LineItemService() *LineItemService
	NewTotalsService() *NewTotalsService
	NoteService() *NoteService
	NotificationService() *NotificationService
	OAuthService() *OAuthService
	OrderItemChangeService() *OrderItemChangeService
	OrderEditService() *OrderEditService
	OrderService() *OrderService
	PaymentCollectionService() *PaymentCollectionService
	PaymentProviderService() *PaymentProviderService
	PaymentService() *PaymentService
	PriceListService() *PriceListService
	PricingService() *PricingService
	ProductCategoryService() *ProductCategoryService
	ProductCollectionService() *ProductCollectionService
	ProductTagService() *ProductTagService
	ProductTaxRateService() *ProductTaxRateService
	ProductTypeService() *ProductTypeService
	ProductVariantInventoryService() *ProductVariantInventoryService
	ProductVariantService() *ProductVariantService
	ProductService() *ProductService
	PublishableApiKeyService() *PublishableApiKeyService
	RegionService() *RegionService
	ReturnReasonService() *ReturnReasonService
	ReturnService() *ReturnService
	SalesChannelInventoryService() *SalesChannelInventoryService
	SalesChannelLocationService() *SalesChannelLocationService
	SalesChannelService() *SalesChannelService
	DefaultSearchService() *DefaultSearchService
	ShippingOptionService() *ShippingOptionService
	ShippingProfileService() *ShippingProfileService
	ShippingTaxRateService() *ShippingTaxRateService
	StagedJobService() *StagedJobService
	StoreService() *StoreService
	StrategyResolverService() *StrategyResolverService
	SwapService() *SwapService
	SystemProviderService() *SystemProviderService
	SystemTaxService() *SystemTaxService
	TaxProviderService() *TaxProviderService
	TaxRateService() *TaxRateService
	TockenService() *TockenService
	TotalsService() *TotalsService
	UserService() *UserService

	//Interfaces
	PriceSelectionStrategy() interfaces.IPriceSelectionStrategy
	TaxCalculationStrategy() interfaces.ITaxCalculationStrategy
	InventoryService() interfaces.IInventoryService
	StockLocationService() interfaces.IStockLocationService
	CacheService() interfaces.ICacheService
	PricingModuleService() interfaces.IPricingModuleService
	BatchJobStrategy() interfaces.IBatchJobStrategy
}
