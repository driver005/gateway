package registry

import "github.com/driver005/gateway/routes/admin"

func (m *Base) AdminAuth() *admin.Auth {
	if m.adminAuth == nil {
		m.adminAuth = admin.NewAuth(m)
	}
	return m.adminAuth
}

func (m *Base) AdminBatch() *admin.Batch {
	if m.adminBatch == nil {
		m.adminBatch = admin.NewBatch(m)
	}
	return m.adminBatch
}

func (m *Base) AdminCollection() *admin.Collection {
	if m.adminCollection == nil {
		m.adminCollection = admin.NewCollection(m)
	}
	return m.adminCollection
}

func (m *Base) AdminCurrencie() *admin.Currencie {
	if m.adminCurrencie == nil {
		m.adminCurrencie = admin.NewCurrencie(m)
	}
	return m.adminCurrencie
}

func (m *Base) AdminCustomerGroup() *admin.CustomerGroup {
	if m.adminCustomerGroup == nil {
		m.adminCustomerGroup = admin.NewCustomerGroup(m)
	}
	return m.adminCustomerGroup
}

func (m *Base) AdminCustomer() *admin.Customer {
	if m.adminCustomer == nil {
		m.adminCustomer = admin.NewCustomer(m)
	}
	return m.adminCustomer
}

func (m *Base) AdminDiscount() *admin.Discount {
	if m.adminDiscount == nil {
		m.adminDiscount = admin.NewDiscount(m)
	}
	return m.adminDiscount
}

func (m *Base) AdminDraftOrder() *admin.DraftOrder {
	if m.adminDraftOrder == nil {
		m.adminDraftOrder = admin.NewDraftOrder(m)
	}
	return m.adminDraftOrder
}

func (m *Base) AdminGiftCard() *admin.GiftCard {
	if m.adminGiftCard == nil {
		m.adminGiftCard = admin.NewGiftCard(m)
	}
	return m.adminGiftCard
}

func (m *Base) AdminInventoryItem() *admin.InventoryItem {
	if m.adminInventoryItem == nil {
		m.adminInventoryItem = admin.NewInventoryItem(m)
	}
	return m.adminInventoryItem
}

func (m *Base) AdminInvite() *admin.Invite {
	if m.adminInvite == nil {
		m.adminInvite = admin.NewInvite(m)
	}
	return m.adminInvite
}

func (m *Base) AdminNote() *admin.Note {
	if m.adminNote == nil {
		m.adminNote = admin.NewNote(m)
	}
	return m.adminNote
}

func (m *Base) AdminNotification() *admin.Notification {
	if m.adminNotification == nil {
		m.adminNotification = admin.NewNotification(m)
	}
	return m.adminNotification
}

func (m *Base) AdminOrderEdit() *admin.OrderEdit {
	if m.adminOrderEdit == nil {
		m.adminOrderEdit = admin.NewOrderEdit(m)
	}
	return m.adminOrderEdit
}

func (m *Base) AdminOrder() *admin.Order {
	if m.adminOrder == nil {
		m.adminOrder = admin.NewOrder(m)
	}
	return m.adminOrder
}

func (m *Base) AdminPaymentCollection() *admin.PaymentCollection {
	if m.adminPaymentCollection == nil {
		m.adminPaymentCollection = admin.NewPaymentCollection(m)
	}
	return m.adminPaymentCollection
}

func (m *Base) AdminPayment() *admin.Payment {
	if m.adminPayment == nil {
		m.adminPayment = admin.NewPayment(m)
	}
	return m.adminPayment
}

func (m *Base) AdminPriceList() *admin.PriceList {
	if m.adminPriceList == nil {
		m.adminPriceList = admin.NewPriceList(m)
	}
	return m.adminPriceList
}

func (m *Base) AdminProductCategory() *admin.ProductCategory {
	if m.adminProductCategory == nil {
		m.adminProductCategory = admin.NewProductCategory(m)
	}
	return m.adminProductCategory
}

func (m *Base) AdminProductTag() *admin.ProductTag {
	if m.adminProductTag == nil {
		m.adminProductTag = admin.NewProductTag(m)
	}
	return m.adminProductTag
}

func (m *Base) AdminProductType() *admin.ProductType {
	if m.adminProductType == nil {
		m.adminProductType = admin.NewProductType(m)
	}
	return m.adminProductType
}

func (m *Base) AdminProduct() *admin.Product {
	if m.adminProduct == nil {
		m.adminProduct = admin.NewProduct(m)
	}
	return m.adminProduct
}

func (m *Base) AdminPublishableApiKey() *admin.PublishableApiKey {
	if m.adminPublishableApiKey == nil {
		m.adminPublishableApiKey = admin.NewPublishableApiKey(m)
	}
	return m.adminPublishableApiKey
}

func (m *Base) AdminRegion() *admin.Region {
	if m.adminRegion == nil {
		m.adminRegion = admin.NewRegion(m)
	}
	return m.adminRegion
}

func (m *Base) AdminReservation() *admin.Reservation {
	if m.adminReservation == nil {
		m.adminReservation = admin.NewReservation(m)
	}
	return m.adminReservation
}

func (m *Base) AdminReturnReason() *admin.ReturnReason {
	if m.adminReturnReason == nil {
		m.adminReturnReason = admin.NewReturnReason(m)
	}
	return m.adminReturnReason
}

func (m *Base) AdminReturn() *admin.Return {
	if m.adminReturn == nil {
		m.adminReturn = admin.NewReturn(m)
	}
	return m.adminReturn
}

func (m *Base) AdminSalesChannel() *admin.SalesChannel {
	if m.adminSalesChannel == nil {
		m.adminSalesChannel = admin.NewSalesChannel(m)
	}
	return m.adminSalesChannel
}

func (m *Base) AdminShippingOption() *admin.ShippingOption {
	if m.adminShippingOption == nil {
		m.adminShippingOption = admin.NewShippingOption(m)
	}
	return m.adminShippingOption
}

func (m *Base) AdminShippingProfile() *admin.ShippingProfile {
	if m.adminShippingProfile == nil {
		m.adminShippingProfile = admin.NewShippingProfile(m)
	}
	return m.adminShippingProfile
}

func (m *Base) AdminStockLocation() *admin.StockLocation {
	if m.adminStockLocation == nil {
		m.adminStockLocation = admin.NewStockLocation(m)
	}
	return m.adminStockLocation
}

func (m *Base) AdminStore() *admin.Store {
	if m.adminStore == nil {
		m.adminStore = admin.NewStore(m)
	}
	return m.adminStore
}

func (m *Base) AdminSwap() *admin.Swap {
	if m.adminSwap == nil {
		m.adminSwap = admin.NewSwap(m)
	}
	return m.adminSwap
}

func (m *Base) AdminTaxRate() *admin.TaxRate {
	if m.adminTaxRate == nil {
		m.adminTaxRate = admin.NewTaxRate(m)
	}
	return m.adminTaxRate
}

func (m *Base) AdminUpload() *admin.Upload {
	if m.adminUpload == nil {
		m.adminUpload = admin.NewUpload(m)
	}
	return m.adminUpload
}

func (m *Base) AdminUser() *admin.User {
	if m.adminUser == nil {
		m.adminUser = admin.NewUser(m)
	}
	return m.adminUser
}

func (m *Base) AdminVariant() *admin.Variant {
	if m.adminVariant == nil {
		m.adminVariant = admin.NewVariant(m)
	}
	return m.adminVariant
}
