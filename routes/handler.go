package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type Routes struct {
	r Registry
}

func New(r Registry) *Routes {
	return &Routes{
		r: r,
	}
}

func (r Routes) SetRoutes() {
	admin := r.r.AdminRouter()

	adminCors := r.r.Config().Server.AdminCors
	admin.Use(cors.New(cors.Config{
		AllowOrigins: adminCors,
	}))

	r.r.AdminAuth().SetRoutes(admin)
	r.r.AdminUser().UnauthenticatedUserRoutes(admin)
	r.r.AdminInvite().UnauthenticatedInviteRoutes(admin)

	admin.Use(convertMiddleware(r.r.Middleware().Authenticate())...)

	r.r.AdminAnalyticsConfig().SetRoutes(admin)
	r.r.AdminApp().SetRoutes(admin)
	r.r.AdminBatch().SetRoutes(admin)
	r.r.AdminCollection().SetRoutes(admin)
	r.r.AdminCurrencie().SetRoutes(admin)
	r.r.AdminCustomerGroup().SetRoutes(admin)
	r.r.AdminCustomer().SetRoutes(admin)
	r.r.AdminDiscount().SetRoutes(admin)
	r.r.AdminDraftOrder().SetRoutes(admin)
	r.r.AdminGiftCard().SetRoutes(admin)
	r.r.AdminInventoryItem().SetRoutes(admin)
	r.r.AdminInvite().SetRoutes(admin)
	r.r.AdminNote().SetRoutes(admin)
	r.r.AdminNotification().SetRoutes(admin)
	r.r.AdminOrderEdit().SetRoutes(admin)
	r.r.AdminOrder().SetRoutes(admin)
	r.r.AdminPaymentCollection().SetRoutes(admin)
	r.r.AdminPayment().SetRoutes(admin)
	r.r.AdminPriceList().SetRoutes(admin)
	r.r.AdminProductCategory().SetRoutes(admin)
	r.r.AdminProductTag().SetRoutes(admin)
	r.r.AdminProductType().SetRoutes(admin)
	r.r.AdminProduct().SetRoutes(admin)
	r.r.AdminPublishableApiKey().SetRoutes(admin)
	r.r.AdminRegion().SetRoutes(admin)
	r.r.AdminReservation().SetRoutes(admin)
	r.r.AdminReturnReason().SetRoutes(admin)
	r.r.AdminReturn().SetRoutes(admin)
	r.r.AdminSalesChannel().SetRoutes(admin)
	r.r.AdminShippingOption().SetRoutes(admin)
	r.r.AdminShippingProfile().SetRoutes(admin)
	r.r.AdminStockLocation().SetRoutes(admin)
	r.r.AdminStore().SetRoutes(admin)
	r.r.AdminSwap().SetRoutes(admin)
	r.r.AdminTaxRate().SetRoutes(admin)
	r.r.AdminUpload().SetRoutes(admin)
	r.r.AdminUser().SetRoutes(admin)
	r.r.AdminVariant().SetRoutes(admin)
}

func convertMiddleware(m []func(fiber.Ctx) error) []any {
	var result []any
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
