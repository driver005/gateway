package routes

import (
	"github.com/driver005/gateway/config"
	"github.com/driver005/gateway/middlewares"
	"github.com/driver005/gateway/routes/admin"
	"github.com/gofiber/fiber/v3"
)

type Registry interface {
	AdminRouter() fiber.Router
	StoreRouter() fiber.Router

	Config() *config.Config
	Middleware() *middlewares.Handler

	//Admin
	AdminAnalyticsConfig() *admin.AnalyticsConfig
	AdminApp() *admin.App
	AdminAuth() *admin.Auth
	AdminBatch() *admin.Batch
	AdminCollection() *admin.Collection
	AdminCurrencie() *admin.Currencie
	AdminCustomerGroup() *admin.CustomerGroup
	AdminCustomer() *admin.Customer
	AdminDiscount() *admin.Discount
	AdminDraftOrder() *admin.DraftOrder
	AdminGiftCard() *admin.GiftCard
	AdminInventoryItem() *admin.InventoryItem
	AdminInvite() *admin.Invite
	AdminNote() *admin.Note
	AdminNotification() *admin.Notification
	AdminOrderEdit() *admin.OrderEdit
	AdminOrder() *admin.Order
	AdminPaymentCollection() *admin.PaymentCollection
	AdminPayment() *admin.Payment
	AdminPriceList() *admin.PriceList
	AdminProductCategory() *admin.ProductCategory
	AdminProductTag() *admin.ProductTag
	AdminProductType() *admin.ProductType
	AdminProduct() *admin.Product
	AdminPublishableApiKey() *admin.PublishableApiKey
	AdminRegion() *admin.Region
	AdminReservation() *admin.Reservation
	AdminReturnReason() *admin.ReturnReason
	AdminReturn() *admin.Return
	AdminSalesChannel() *admin.SalesChannel
	AdminShippingOption() *admin.ShippingOption
	AdminShippingProfile() *admin.ShippingProfile
	AdminStockLocation() *admin.StockLocation
	AdminStore() *admin.Store
	AdminSwap() *admin.Swap
	AdminTaxRate() *admin.TaxRate
	AdminUpload() *admin.Upload
	AdminUser() *admin.User
	AdminVariant() *admin.Variant
}
