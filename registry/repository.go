package registry

import "github.com/driver005/gateway/repository"

func (m *Base) AddressRepository() *repository.AddressRepo {
	if m.addressRepo == nil {
		m.addressRepo = repository.AddressRepository(m.Context())
	}
	return m.addressRepo
}
func (m *Base) AnalyticsConfigRepository() *repository.AnalyticsConfigRepo {
	if m.analyticsConfigRepo == nil {
		m.analyticsConfigRepo = repository.AnalyticsConfigRepository(m.Context())
	}
	return m.analyticsConfigRepo
}
func (m *Base) BatchJobRepository() *repository.BatchJobRepo {
	if m.batchJobRepo == nil {
		m.batchJobRepo = repository.BatchJobRepository(m.Context())
	}
	return m.batchJobRepo
}
func (m *Base) CartRepository() *repository.CartRepo {
	if m.cartRepo == nil {
		m.cartRepo = repository.CartRepository(m.Context())
	}
	return m.cartRepo
}
func (m *Base) ClaimImageRepository() *repository.ClaimImageRepo {
	if m.claimImageRepo == nil {
		m.claimImageRepo = repository.ClaimImageRepository(m.Context())
	}
	return m.claimImageRepo
}
func (m *Base) ClaimItemRepository() *repository.ClaimItemRepo {
	if m.claimItemRepo == nil {
		m.claimItemRepo = repository.ClaimItemRepository(m.Context())
	}
	return m.claimItemRepo
}
func (m *Base) ClaimTagRepository() *repository.ClaimTagRepo {
	if m.claimTagRepo == nil {
		m.claimTagRepo = repository.ClaimTagRepository(m.Context())
	}
	return m.claimTagRepo
}
func (m *Base) ClaimRepository() *repository.ClaimRepo {
	if m.claimRepo == nil {
		m.claimRepo = repository.ClaimRepository(m.Context())
	}
	return m.claimRepo
}
func (m *Base) CountryRepository() *repository.CountryRepo {
	if m.countryRepo == nil {
		m.countryRepo = repository.CountryRepository(m.Context())
	}
	return m.countryRepo
}
func (m *Base) CurrencyRepository() *repository.CurrencyRepo {
	if m.currencyRepo == nil {
		m.currencyRepo = repository.CurrencyRepository(m.Context())
	}
	return m.currencyRepo
}
func (m *Base) CustomShippingOptionRepository() *repository.CustomShippingOptionRepo {
	if m.customShippingOptionRepo == nil {
		m.customShippingOptionRepo = repository.CustomShippingOptionRepository(m.Context())
	}
	return m.customShippingOptionRepo
}
func (m *Base) CustomerGroupRepository() *repository.CustomerGroupRepo {
	if m.customerGroupRepo == nil {
		m.customerGroupRepo = repository.CustomerGroupRepository(m.Context())
	}
	return m.customerGroupRepo
}
func (m *Base) CustomerRepository() *repository.CustomerRepo {
	if m.customerRepo == nil {
		m.customerRepo = repository.CustomerRepository(m.Context())
	}
	return m.customerRepo
}
func (m *Base) DiscountConditionRepository() *repository.DiscountConditionRepo {
	if m.discountConditionRepo == nil {
		m.discountConditionRepo = repository.DiscountConditionRepository(m.Context())
	}
	return m.discountConditionRepo
}
func (m *Base) DiscountRuleRepository() *repository.DiscountRuleRepo {
	if m.discountRuleRepo == nil {
		m.discountRuleRepo = repository.DiscountRuleRepository(m.Context())
	}
	return m.discountRuleRepo
}
func (m *Base) DiscountRepository() *repository.DiscountRepo {
	if m.discountRepo == nil {
		m.discountRepo = repository.DiscountRepository(m.Context())
	}
	return m.discountRepo
}
func (m *Base) DraftOrderRepository() *repository.DraftOrderRepo {
	if m.draftOrderRepo == nil {
		m.draftOrderRepo = repository.DraftOrderRepository(m.Context())
	}
	return m.draftOrderRepo
}
func (m *Base) FulfillmentProviderRepository() *repository.FulfillmentProviderRepo {
	if m.fulfillmentProviderRepo == nil {
		m.fulfillmentProviderRepo = repository.FulfillmentProviderRepository(m.Context())
	}
	return m.fulfillmentProviderRepo
}
func (m *Base) FulfillmentRepository() *repository.FulfillmentRepo {
	if m.fulfillmentRepo == nil {
		m.fulfillmentRepo = repository.FulfillmentRepository(m.Context())
	}
	return m.fulfillmentRepo
}
func (m *Base) GiftCardTransactionRepository() *repository.GiftCardTransactionRepo {
	if m.giftCardTransactionRepo == nil {
		m.giftCardTransactionRepo = repository.GiftCardTransactionRepository(m.Context())
	}
	return m.giftCardTransactionRepo
}
func (m *Base) GiftCardRepository() *repository.GiftCardRepo {
	if m.giftCardRepo == nil {
		m.giftCardRepo = repository.GiftCardRepository(m.Context())
	}
	return m.giftCardRepo
}
func (m *Base) IdempotencyKeyRepository() *repository.IdempotencyKeyRepo {
	if m.idempotencyKeyRepo == nil {
		m.idempotencyKeyRepo = repository.IdempotencyKeyRepository(m.Context())
	}
	return m.idempotencyKeyRepo
}
func (m *Base) ImageRepository() *repository.ImageRepo {
	if m.imageRepo == nil {
		m.imageRepo = repository.ImageRepository(m.Context())
	}
	return m.imageRepo
}
func (m *Base) InviteRepository() *repository.InviteRepo {
	if m.inviteRepo == nil {
		m.inviteRepo = repository.InviteRepository(m.Context())
	}
	return m.inviteRepo
}
func (m *Base) LineItemAdjustmentRepository() *repository.LineItemAdjustmentRepo {
	if m.ineItemAdjustmentRepo == nil {
		m.ineItemAdjustmentRepo = repository.LineItemAdjustmentRepository(m.Context())
	}
	return m.ineItemAdjustmentRepo
}
func (m *Base) LineItemTaxLineRepository() *repository.LineItemTaxLineRepo {
	if m.ineItemTaxLineRepo == nil {
		m.ineItemTaxLineRepo = repository.LineItemTaxLineRepository(m.Context())
	}
	return m.ineItemTaxLineRepo
}
func (m *Base) LineItemRepository() *repository.LineItemRepo {
	if m.ineItemRepo == nil {
		m.ineItemRepo = repository.LineItemRepository(m.Context())
	}
	return m.ineItemRepo
}
func (m *Base) MoneyAmountRepository() *repository.MoneyAmountRepo {
	if m.moneyAmountRepo == nil {
		m.moneyAmountRepo = repository.MoneyAmountRepository(m.Context())
	}
	return m.moneyAmountRepo
}
func (m *Base) NoteRepository() *repository.NoteRepo {
	if m.noteRepo == nil {
		m.noteRepo = repository.NoteRepository(m.Context())
	}
	return m.noteRepo
}
func (m *Base) NotificationProviderRepository() *repository.NotificationProviderRepo {
	if m.notificationProviderRepo == nil {
		m.notificationProviderRepo = repository.NotificationProviderRepository(m.Context())
	}
	return m.notificationProviderRepo
}
func (m *Base) NotificationRepository() *repository.NotificationRepo {
	if m.notificationRepo == nil {
		m.notificationRepo = repository.NotificationRepository(m.Context())
	}
	return m.notificationRepo
}
func (m *Base) OAuthRepository() *repository.OAuthRepo {
	if m.oAuthRepo == nil {
		m.oAuthRepo = repository.OAuthRepository(m.Context())
	}
	return m.oAuthRepo
}
func (m *Base) OrderEditRepository() *repository.OrderEditRepo {
	if m.orderEditRepo == nil {
		m.orderEditRepo = repository.OrderEditRepository(m.Context())
	}
	return m.orderEditRepo
}
func (m *Base) OrderItemChangeRepository() *repository.OrderItemChangeRepo {
	if m.orderItemChangeRepo == nil {
		m.orderItemChangeRepo = repository.OrderItemChangeRepository(m.Context())
	}
	return m.orderItemChangeRepo
}
func (m *Base) OrderRepository() *repository.OrderRepo {
	if m.orderRepo == nil {
		m.orderRepo = repository.OrderRepository(m.Context())
	}
	return m.orderRepo
}
func (m *Base) PaymentCollectionRepository() *repository.PaymentCollectionRepo {
	if m.paymentCollectionRepo == nil {
		m.paymentCollectionRepo = repository.PaymentCollectionRepository(m.Context())
	}
	return m.paymentCollectionRepo
}
func (m *Base) PaymentProviderRepository() *repository.PaymentProviderRepo {
	if m.paymentProviderRepo == nil {
		m.paymentProviderRepo = repository.PaymentProviderRepository(m.Context())
	}
	return m.paymentProviderRepo
}
func (m *Base) PaymentSessionRepository() *repository.PaymentSessionRepo {
	if m.paymentSessionRepo == nil {
		m.paymentSessionRepo = repository.PaymentSessionRepository(m.Context())
	}
	return m.paymentSessionRepo
}
func (m *Base) PaymentRepository() *repository.PaymentRepo {
	if m.paymentRepo == nil {
		m.paymentRepo = repository.PaymentRepository(m.Context())
	}
	return m.paymentRepo
}
func (m *Base) PriceListRepository() *repository.PriceListRepo {
	if m.priceListRepo == nil {
		m.priceListRepo = repository.PriceListRepository(m.Context())
	}
	return m.priceListRepo
}
func (m *Base) ProductCategoryRepository() *repository.ProductCategoryRepo {
	if m.productCategoryRepo == nil {
		m.productCategoryRepo = repository.ProductCategoryRepository(m.Context())
	}
	return m.productCategoryRepo
}
func (m *Base) ProductCollectionRepository() *repository.ProductCollectionRepo {
	if m.productCollectionRepo == nil {
		m.productCollectionRepo = repository.ProductCollectionRepository(m.Context())
	}
	return m.productCollectionRepo
}
func (m *Base) ProductOptionValueRepository() *repository.ProductOptionValueRepo {
	if m.productOptionValueRepo == nil {
		m.productOptionValueRepo = repository.ProductOptionValueRepository(m.Context())
	}
	return m.productOptionValueRepo
}
func (m *Base) ProductOptionRepository() *repository.ProductOptionRepo {
	if m.productOptionRepo == nil {
		m.productOptionRepo = repository.ProductOptionRepository(m.Context())
	}
	return m.productOptionRepo
}
func (m *Base) ProductTagRepository() *repository.ProductTagRepo {
	if m.productTagRepo == nil {
		m.productTagRepo = repository.ProductTagRepository(m.Context())
	}
	return m.productTagRepo
}
func (m *Base) ProductTaxRateRepository() *repository.ProductTaxRateRepo {
	if m.productTaxRateRepo == nil {
		m.productTaxRateRepo = repository.ProductTaxRateRepository(m.Context())
	}
	return m.productTaxRateRepo
}
func (m *Base) ProductTypeRepository() *repository.ProductTypeRepo {
	if m.productTypeRepo == nil {
		m.productTypeRepo = repository.ProductTypeRepository(m.Context())
	}
	return m.productTypeRepo
}
func (m *Base) ProductVariantInventoryItemRepository() *repository.ProductVariantInventoryItem {
	if m.productVariantInventoryItem == nil {
		m.productVariantInventoryItem = repository.ProductVariantInventoryItemRepository(m.Context())
	}
	return m.productVariantInventoryItem
}
func (m *Base) ProductVariantRepository() *repository.ProductVariantRepo {
	if m.productVariantRepo == nil {
		m.productVariantRepo = repository.ProductVariantRepository(m.Context())
	}
	return m.productVariantRepo
}
func (m *Base) ProductRepository() *repository.ProductRepo {
	if m.productRepo == nil {
		m.productRepo = repository.ProductRepository(m.Context())
	}
	return m.productRepo
}
func (m *Base) PublishableApiKeySalesChannelRepository() *repository.PublishableApiKeySalesChannelRepo {

	if m.publishableApiKeySalesChannelRepo == nil {
		m.publishableApiKeySalesChannelRepo = repository.PublishableApiKeySalesChannelRepository(m.Context())
	}
	return m.publishableApiKeySalesChannelRepo
}
func (m *Base) PublishableApiKeyRepository() *repository.PublishableApiKeyRepo {
	if m.publishableApiKeyRepo == nil {
		m.publishableApiKeyRepo = repository.PublishableApiKeyRepository(m.Context())
	}
	return m.publishableApiKeyRepo
}
func (m *Base) RefundRepository() *repository.RefundRepo {
	if m.refundRepo == nil {
		m.refundRepo = repository.RefundRepository(m.Context())
	}
	return m.refundRepo
}
func (m *Base) RegionRepository() *repository.RegionRepo {
	if m.regionRepo == nil {
		m.regionRepo = repository.RegionRepository(m.Context())
	}
	return m.regionRepo
}
func (m *Base) ReturnItemRepository() *repository.ReturnItemRepo {
	if m.returnItemRepo == nil {
		m.returnItemRepo = repository.ReturnItemRepository(m.Context())
	}
	return m.returnItemRepo
}
func (m *Base) ReturnReasonRepository() *repository.ReturnReasonRepo {
	if m.returnReasonRepo == nil {
		m.returnReasonRepo = repository.ReturnReasonRepository(m.Context())
	}
	return m.returnReasonRepo
}
func (m *Base) ReturnRepository() *repository.ReturnRepo {
	if m.returnRepo == nil {
		m.returnRepo = repository.ReturnRepository(m.Context())
	}
	return m.returnRepo
}
func (m *Base) SalesChannelLocationRepository() *repository.SalesChannelLocationRepo {
	if m.salesChannelLocationRepo == nil {
		m.salesChannelLocationRepo = repository.SalesChannelLocationRepository(m.Context())
	}
	return m.salesChannelLocationRepo
}
func (m *Base) SalesChannelRepository() *repository.SalesChannelRepo {
	if m.salesChannelRepo == nil {
		m.salesChannelRepo = repository.SalesChannelRepository(m.Context())
	}
	return m.salesChannelRepo
}
func (m *Base) ShippingMethodTaxLineRepository() *repository.ShippingMethodTaxLineRepo {
	if m.shippingMethodTaxLineRepo == nil {
		m.shippingMethodTaxLineRepo = repository.ShippingMethodTaxLineRepository(m.Context())
	}
	return m.shippingMethodTaxLineRepo
}
func (m *Base) ShippingMethodRepository() *repository.ShippingMethodRepo {
	if m.shippingMethodRepo == nil {
		m.shippingMethodRepo = repository.ShippingMethodRepository(m.Context())
	}
	return m.shippingMethodRepo
}
func (m *Base) ShippingOptionRequirementRepository() *repository.ShippingOptionRequirementRepo {
	if m.shippingOptionRequirementRepo == nil {
		m.shippingOptionRequirementRepo = repository.ShippingOptionRequirementRepository(m.Context())
	}
	return m.shippingOptionRequirementRepo
}
func (m *Base) ShippingOptionRepository() *repository.ShippingOptionRepo {
	if m.shippingOptionRepo == nil {
		m.shippingOptionRepo = repository.ShippingOptionRepository(m.Context())
	}
	return m.shippingOptionRepo
}
func (m *Base) ShippingProfileRepository() *repository.ShippingProfileRepo {
	if m.shippingProfileRepo == nil {
		m.shippingProfileRepo = repository.ShippingProfileRepository(m.Context())
	}
	return m.shippingProfileRepo
}
func (m *Base) ShippingTaxRateRepository() *repository.ShippingTaxRateRepo {
	if m.shippingTaxRateRepo == nil {
		m.shippingTaxRateRepo = repository.ShippingTaxRateRepository(m.Context())
	}
	return m.shippingTaxRateRepo
}
func (m *Base) StagedJobRepository() *repository.StagedJobRepo {
	if m.stagedJobRepo == nil {
		m.stagedJobRepo = repository.StagedJobRepository(m.Context())
	}
	return m.stagedJobRepo
}
func (m *Base) StoreRepository() *repository.StoreRepo {
	if m.storeRepo == nil {
		m.storeRepo = repository.StoreRepository(m.Context())
	}
	return m.storeRepo
}
func (m *Base) SwapRepository() *repository.SwapRepo {
	if m.swapRepo == nil {
		m.swapRepo = repository.SwapRepository(m.Context())
	}
	return m.swapRepo
}
func (m *Base) TaxProviderRepository() *repository.TaxProviderRepo {
	if m.taxProviderRepo == nil {
		m.taxProviderRepo = repository.TaxProviderRepository(m.Context())
	}
	return m.taxProviderRepo
}
func (m *Base) TaxRateRepository() *repository.TaxRateRepo {
	if m.taxRateRepo == nil {
		m.taxRateRepo = repository.TaxRateRepository(m.Context())
	}
	return m.taxRateRepo
}
func (m *Base) TrackingLinkRepository() *repository.TrackingLinkRepo {
	if m.trackingLinkRepo == nil {
		m.trackingLinkRepo = repository.TrackingLinkRepository(m.Context())
	}
	return m.trackingLinkRepo
}
func (m *Base) UserRepository() *repository.UserRepo {
	if m.userRepo == nil {
		m.userRepo = repository.UserRepository(m.Context())
	}
	return m.userRepo
}
