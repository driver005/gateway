package registry

import "github.com/driver005/gateway/routes/store"

func (m *Base) StoreAuth() *store.Auth {
	if m.storeAuth == nil {
		m.storeAuth = store.NewAuth(m)
	}
	return m.storeAuth
}

func (m *Base) StoreCart() *store.Cart {
	if m.storeCart == nil {
		m.storeCart = store.NewCart(m)
	}
	return m.storeCart
}
func (m *Base) StoreCollection() *store.Collection {
	if m.storeCollection == nil {
		m.storeCollection = store.NewCollection(m)
	}
	return m.storeCollection
}
func (m *Base) StoreCustomer() *store.Customer {
	if m.storeCustomer == nil {
		m.storeCustomer = store.NewCustomer(m)
	}
	return m.storeCustomer
}
func (m *Base) StoreGiftCard() *store.GiftCard {
	if m.storeGiftCard == nil {
		m.storeGiftCard = store.NewGiftCard(m)
	}
	return m.storeGiftCard
}
func (m *Base) StoreOrder() *store.Order {
	if m.storeOrder == nil {
		m.storeOrder = store.NewOrder(m)
	}
	return m.storeOrder
}
func (m *Base) StoreOrderEdit() *store.OrderEdit {
	if m.storeOrderEdit == nil {
		m.storeOrderEdit = store.NewOrderEdit(m)
	}
	return m.storeOrderEdit
}
func (m *Base) StorePaymentCollection() *store.PaymentCollection {
	if m.storePaymentCollection == nil {
		m.storePaymentCollection = store.NewPaymentCollection(m)
	}
	return m.storePaymentCollection
}
func (m *Base) StoreProduct() *store.Product {
	if m.storeProduct == nil {
		m.storeProduct = store.NewProduct(m)
	}
	return m.storeProduct
}
func (m *Base) StoreProductCategory() *store.ProductCategory {
	if m.storeProductCategory == nil {
		m.storeProductCategory = store.NewProductCategory(m)
	}
	return m.storeProductCategory
}
func (m *Base) StoreProductTag() *store.ProductTag {
	if m.storeProductTag == nil {
		m.storeProductTag = store.NewProductTag(m)
	}
	return m.storeProductTag
}
func (m *Base) StoreProductType() *store.ProductType {
	if m.storeProductType == nil {
		m.storeProductType = store.NewProductType(m)
	}
	return m.storeProductType
}
func (m *Base) StoreRegion() *store.Region {
	if m.storeRegion == nil {
		m.storeRegion = store.NewRegion(m)
	}
	return m.storeRegion
}
func (m *Base) StoreReturn() *store.Return {
	if m.storeReturn == nil {
		m.storeReturn = store.NewReturn(m)
	}
	return m.storeReturn
}
func (m *Base) StoreReturnReason() *store.ReturnReason {
	if m.storeReturnReason == nil {
		m.storeReturnReason = store.NewReturnReason(m)
	}
	return m.storeReturnReason
}
func (m *Base) StoreShippingOption() *store.ShippingOption {
	if m.storeShippingOption == nil {
		m.storeShippingOption = store.NewShippingOption(m)
	}
	return m.storeShippingOption
}
func (m *Base) StoreSwap() *store.Swap {
	if m.storeSwap == nil {
		m.storeSwap = store.NewSwap(m)
	}
	return m.storeSwap
}
func (m *Base) StoreVariant() *store.Variant {
	if m.storeVariant == nil {
		m.storeVariant = store.NewVariant(m)
	}
	return m.storeVariant
}
