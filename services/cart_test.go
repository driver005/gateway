package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

func TestNewCartService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *CartService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCartService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCartService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *CartService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_List(t *testing.T) {
	type args struct {
		selector types.FilterableCartProps
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  []models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_Retrieve(t *testing.T) {
	type args struct {
		id           uuid.UUID
		config       *sql.Options
		totalsConfig TotalsConfig
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.id, tt.args.config, tt.args.totalsConfig)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_RetrieveLegacy(t *testing.T) {
	type args struct {
		id           uuid.UUID
		config       *sql.Options
		totalsConfig TotalsConfig
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveLegacy(tt.args.id, tt.args.config, tt.args.totalsConfig)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.RetrieveLegacy() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.RetrieveLegacy() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_RetrieveWithTotals(t *testing.T) {
	type args struct {
		id           uuid.UUID
		config       *sql.Options
		totalsConfig TotalsConfig
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveWithTotals(tt.args.id, tt.args.config, tt.args.totalsConfig)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.RetrieveWithTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.RetrieveWithTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_Create(t *testing.T) {
	type args struct {
		data *types.CartCreateProps
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_getValidatedSalesChannel(t *testing.T) {
	type args struct {
		salesChannelId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.SalesChannel
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.getValidatedSalesChannel(tt.args.salesChannelId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.getValidatedSalesChannel() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.getValidatedSalesChannel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_RemoveLineItem(t *testing.T) {
	type args struct {
		id          uuid.UUID
		lineItemIds uuid.UUIDs
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RemoveLineItem(tt.args.id, tt.args.lineItemIds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.RemoveLineItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_validateLineItemShipping(t *testing.T) {
	type args struct {
		shippingMethods            []models.ShippingMethod
		lineItemShippingProfiledId uuid.UUID
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.validateLineItemShipping(tt.args.shippingMethods, tt.args.lineItemShippingProfiledId); got != tt.want {
				t.Errorf("CartService.validateLineItemShipping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_ValidateLineItem(t *testing.T) {
	type args struct {
		salesChannelId uuid.UUID
		lineItem       types.LineItemValidateData
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  bool
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ValidateLineItem(tt.args.salesChannelId, tt.args.lineItem)
			if got != tt.want {
				t.Errorf("CartService.ValidateLineItem() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.ValidateLineItem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_AddLineItem(t *testing.T) {
	type args struct {
		cartId                uuid.UUID
		lineItem              models.LineItem
		validateSalesChannels bool
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.AddLineItem(tt.args.cartId, tt.args.lineItem, tt.args.validateSalesChannels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.AddLineItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_AddOrUpdateLineItems(t *testing.T) {
	type args struct {
		cartId                uuid.UUID
		lineItems             []models.LineItem
		validateSalesChannels bool
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.AddOrUpdateLineItems(tt.args.cartId, tt.args.lineItems, tt.args.validateSalesChannels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.AddOrUpdateLineItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_UpdateLineItem(t *testing.T) {
	type args struct {
		cartId     uuid.UUID
		lineItemId uuid.UUID
		update     *types.LineItemUpdate
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdateLineItem(tt.args.cartId, tt.args.lineItemId, tt.args.update)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.UpdateLineItem() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.UpdateLineItem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_adjustFreeShipping(t *testing.T) {
	type args struct {
		cart      *models.Cart
		shouldAdd bool
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.adjustFreeShipping(tt.args.cart, tt.args.shouldAdd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.adjustFreeShipping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_Update(t *testing.T) {
	type args struct {
		id   uuid.UUID
		cart *models.Cart
		data *types.CartUpdateProps
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.cart, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_onSalesChannelChange(t *testing.T) {
	type args struct {
		cart              *models.Cart
		newSalesChannelId uuid.UUID
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.onSalesChannelChange(tt.args.cart, tt.args.newSalesChannelId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.onSalesChannelChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_updateCustomerId(t *testing.T) {
	type args struct {
		cart       *models.Cart
		customerId uuid.UUID
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.updateCustomerId(tt.args.cart, tt.args.customerId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.updateCustomerId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_createOrFetchGuestCustomerFromEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Customer
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.createOrFetchGuestCustomerFromEmail(tt.args.email)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.createOrFetchGuestCustomerFromEmail() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.createOrFetchGuestCustomerFromEmail() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_updateBillingAddress(t *testing.T) {
	type args struct {
		cart    *models.Cart
		id      uuid.UUID
		address *models.Address
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.updateBillingAddress(tt.args.cart, tt.args.id, tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.updateBillingAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_updateShippingAddress(t *testing.T) {
	type args struct {
		cart    *models.Cart
		id      uuid.UUID
		address *models.Address
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.updateShippingAddress(tt.args.cart, tt.args.id, tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.updateShippingAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_applyGiftCard(t *testing.T) {
	type args struct {
		cart *models.Cart
		code string
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.applyGiftCard(tt.args.cart, tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.applyGiftCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_ApplyDiscount(t *testing.T) {
	type args struct {
		cart         *models.Cart
		discountCode string
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ApplyDiscount(tt.args.cart, tt.args.discountCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.ApplyDiscount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_ApplyDiscounts(t *testing.T) {
	type args struct {
		cart          *models.Cart
		discountCodes []string
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ApplyDiscounts(tt.args.cart, tt.args.discountCodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.ApplyDiscounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_RemoveDiscount(t *testing.T) {
	type args struct {
		cartId       uuid.UUID
		discountCode string
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RemoveDiscount(tt.args.cartId, tt.args.discountCode)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.RemoveDiscount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.RemoveDiscount() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_UpdatePaymentSession(t *testing.T) {
	type args struct {
		cartId uuid.UUID
		update map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdatePaymentSession(tt.args.cartId, tt.args.update)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.UpdatePaymentSession() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.UpdatePaymentSession() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_AuthorizePayment(t *testing.T) {
	type args struct {
		id      uuid.UUID
		cart    *models.Cart
		context map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AuthorizePayment(tt.args.id, tt.args.cart, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.AuthorizePayment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.AuthorizePayment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_SetPaymentSession(t *testing.T) {
	type args struct {
		id         uuid.UUID
		providerId uuid.UUID
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetPaymentSession(tt.args.id, tt.args.providerId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.SetPaymentSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_SetPaymentSessions(t *testing.T) {
	type args struct {
		id   uuid.UUID
		cart *models.Cart
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetPaymentSessions(tt.args.id, tt.args.cart); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.SetPaymentSessions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_DeletePaymentSession(t *testing.T) {
	type args struct {
		id         uuid.UUID
		providerId uuid.UUID
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DeletePaymentSession(tt.args.id, tt.args.providerId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.DeletePaymentSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_RefreshPaymentSession(t *testing.T) {
	type args struct {
		id         uuid.UUID
		providerId uuid.UUID
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RefreshPaymentSession(tt.args.id, tt.args.providerId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.RefreshPaymentSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_AddShippingMethod(t *testing.T) {
	type args struct {
		id       uuid.UUID
		cart     *models.Cart
		optionId uuid.UUID
		data     map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddShippingMethod(tt.args.id, tt.args.cart, tt.args.optionId, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.AddShippingMethod() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.AddShippingMethod() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_findCustomShippingOption(t *testing.T) {
	type args struct {
		cartCustomShippingOptions []models.CustomShippingOption
		optionId                  uuid.UUID
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.CustomShippingOption
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.findCustomShippingOption(tt.args.cartCustomShippingOptions, tt.args.optionId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.findCustomShippingOption() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.findCustomShippingOption() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_updateUnitPrices(t *testing.T) {
	type args struct {
		cart       *models.Cart
		regionId   uuid.UUID
		customerId uuid.UUID
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.updateUnitPrices(tt.args.cart, tt.args.regionId, tt.args.customerId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.updateUnitPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_setRegion(t *testing.T) {
	type args struct {
		cart        *models.Cart
		regionId    uuid.UUID
		countryCode string
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.setRegion(tt.args.cart, tt.args.regionId, tt.args.countryCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.setRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_Delete(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Delete(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.Delete() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.Delete() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_SetMetadata(t *testing.T) {
	type args struct {
		id    uuid.UUID
		key   string
		value string
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetMetadata(tt.args.id, tt.args.key, tt.args.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.SetMetadata() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.SetMetadata() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_CreateTaxLines(t *testing.T) {
	type args struct {
		id   uuid.UUID
		cart *models.Cart
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.CreateTaxLines(tt.args.id, tt.args.cart); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.CreateTaxLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_DeleteTaxLines(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DeleteTaxLines(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.DeleteTaxLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_DecorateTotals(t *testing.T) {
	type args struct {
		cart         *models.Cart
		totalsConfig TotalsConfig
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.DecorateTotals(tt.args.cart, tt.args.totalsConfig)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.DecorateTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.DecorateTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_refreshAdjustments(t *testing.T) {
	type args struct {
		cart *models.Cart
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.refreshAdjustments(tt.args.cart); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.refreshAdjustments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartService_transformQueryForTotals(t *testing.T) {
	type args struct {
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  []string
		want1 []string
		want2 []types.TotalField
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.transformQueryForTotals(tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.transformQueryForTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.transformQueryForTotals() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("CartService.transformQueryForTotals() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestCartService_decorateTotals(t *testing.T) {
	type args struct {
		cart           *models.Cart
		totalsToSelect []types.TotalField
		config         TotalsConfig
	}
	tests := []struct {
		name  string
		s     *CartService
		args  args
		want  *models.Cart
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.decorateTotals(tt.args.cart, tt.args.totalsToSelect, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.decorateTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CartService.decorateTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCartService_getTotalsRelations(t *testing.T) {
	type args struct {
		config *sql.Options
	}
	tests := []struct {
		name string
		s    *CartService
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.getTotalsRelations(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartService.getTotalsRelations() = %v, want %v", got, tt.want)
			}
		})
	}
}
