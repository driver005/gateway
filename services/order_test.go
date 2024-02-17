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

func TestNewOrderService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *OrderService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *OrderService
		args args
		want *OrderService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableOrder
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  []models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableOrder
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  []models.Order
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("OrderService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestOrderService_Retrieve(t *testing.T) {
	type args struct {
		selector models.Order
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_RetrieveWithTotals(t *testing.T) {
	type args struct {
		selector models.Order
		config   *sql.Options
		context  types.TotalsContext
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveWithTotals(tt.args.selector, tt.args.config, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.RetrieveWithTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.RetrieveWithTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_RetrieveById(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveById(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.RetrieveById() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.RetrieveById() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_RetrieveLegacy(t *testing.T) {
	type args struct {
		id       uuid.UUID
		selector models.Order
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveLegacy(tt.args.id, tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.RetrieveLegacy() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.RetrieveLegacy() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_RetrieveByIdWithTotals(t *testing.T) {
	type args struct {
		id      uuid.UUID
		config  *sql.Options
		context types.TotalsContext
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByIdWithTotals(tt.args.id, tt.args.config, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.RetrieveByIdWithTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.RetrieveByIdWithTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_RetrieveByCartId(t *testing.T) {
	type args struct {
		cartId uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByCartId(tt.args.cartId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.RetrieveByCartId() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.RetrieveByCartId() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_RetrieveByCartIdWithTotals(t *testing.T) {
	type args struct {
		cartId uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByCartIdWithTotals(tt.args.cartId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.RetrieveByCartIdWithTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.RetrieveByCartIdWithTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_RetrieveByExternalId(t *testing.T) {
	type args struct {
		externalId uuid.UUID
		config     *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RetrieveByExternalId(tt.args.externalId, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.RetrieveByExternalId() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.RetrieveByExternalId() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_CompleteOrder(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CompleteOrder(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.CompleteOrder() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.CompleteOrder() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_CreateFromCart(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *models.Cart
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateFromCart(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.CreateFromCart() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.CreateFromCart() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_createGiftCardsFromLineItem(t *testing.T) {
	type args struct {
		order    *models.Order
		lineItem models.LineItem
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  []models.GiftCard
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.createGiftCardsFromLineItem(tt.args.order, tt.args.lineItem)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.createGiftCardsFromLineItem() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.createGiftCardsFromLineItem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_CreateShipment(t *testing.T) {
	type args struct {
		orderId       uuid.UUID
		fulfillmentId uuid.UUID
		trackingLinks []models.TrackingLink
		config        struct {
			NoNotification bool
			Metadata       map[string]interface{}
		}
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateShipment(tt.args.orderId, tt.args.fulfillmentId, tt.args.trackingLinks, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.CreateShipment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.CreateShipment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_UpdateBillingAddress(t *testing.T) {
	type args struct {
		order   *models.Order
		address *models.Address
	}
	tests := []struct {
		name string
		s    *OrderService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UpdateBillingAddress(tt.args.order, tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.UpdateBillingAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderService_UpdateShippingAddress(t *testing.T) {
	type args struct {
		order   *models.Order
		address *models.Address
	}
	tests := []struct {
		name string
		s    *OrderService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.UpdateShippingAddress(tt.args.order, tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.UpdateShippingAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderService_AddShippingMethod(t *testing.T) {
	type args struct {
		orderId  uuid.UUID
		optionId uuid.UUID
		data     map[string]interface{}
		config   *types.CreateShippingMethodDto
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddShippingMethod(tt.args.orderId, tt.args.optionId, tt.args.data, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.AddShippingMethod() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.AddShippingMethod() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_Update(t *testing.T) {
	type args struct {
		orderId uuid.UUID
		update  *types.UpdateOrderInput
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.orderId, tt.args.update)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_Cancel(t *testing.T) {
	type args struct {
		orderId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Cancel(tt.args.orderId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.Cancel() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.Cancel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_CapturePayment(t *testing.T) {
	type args struct {
		orderId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CapturePayment(tt.args.orderId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.CapturePayment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.CapturePayment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_ValidateFulfillmentLineItem(t *testing.T) {
	type args struct {
		item     *models.LineItem
		quantity int
	}
	tests := []struct {
		name string
		s    *OrderService
		args args
		want *models.LineItem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ValidateFulfillmentLineItem(tt.args.item, tt.args.quantity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.ValidateFulfillmentLineItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_paymentsAreCaptured(t *testing.T) {
	type args struct {
		payments []models.Payment
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paymentsAreCaptured(tt.args.payments); got != tt.want {
				t.Errorf("paymentsAreCaptured() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderService_CreateFulfillment(t *testing.T) {
	type args struct {
		id             uuid.UUID
		itemsToFulfill []types.FulFillmentItemType
		config         map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateFulfillment(tt.args.id, tt.args.itemsToFulfill, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.CreateFulfillment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.CreateFulfillment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_CancelFulfillment(t *testing.T) {
	type args struct {
		fulfillmentId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CancelFulfillment(tt.args.fulfillmentId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.CancelFulfillment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.CancelFulfillment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_GetFulfillmentItems(t *testing.T) {
	type args struct {
		order       models.Order
		items       []types.FulFillmentItemType
		transformer func(item models.LineItem, quantity int) interface{}
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  []models.LineItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetFulfillmentItems(tt.args.order, tt.args.items, tt.args.transformer)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.GetFulfillmentItems() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.GetFulfillmentItems() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_Archive(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Archive(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.Archive() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.Archive() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_CreateRefund(t *testing.T) {
	type args struct {
		id             uuid.UUID
		refundAmount   float64
		reason         models.RefundReason
		note           *string
		noNotification *bool
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateRefund(tt.args.id, tt.args.refundAmount, tt.args.reason, tt.args.note, tt.args.noNotification)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.CreateRefund() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.CreateRefund() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_decorateTotalsLegacy(t *testing.T) {
	type args struct {
		order        *models.Order
		totalsFields []string
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.decorateTotalsLegacy(tt.args.order, tt.args.totalsFields)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.decorateTotalsLegacy() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.decorateTotalsLegacy() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_decorateTotals(t *testing.T) {
	type args struct {
		order        *models.Order
		totalsFields []string
		context      types.TotalsContext
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.decorateTotals(tt.args.order, tt.args.totalsFields, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.decorateTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.decorateTotals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_RegisterReturnReceived(t *testing.T) {
	type args struct {
		id                 uuid.UUID
		receivedReturn     *models.Return
		customRefundAmount *float64
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  *models.Order
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RegisterReturnReceived(tt.args.id, tt.args.receivedReturn, tt.args.customRefundAmount)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.RegisterReturnReceived() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.RegisterReturnReceived() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrderService_transformQueryForTotals(t *testing.T) {
	type args struct {
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *OrderService
		args  args
		want  []string
		want1 []string
		want2 []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.transformQueryForTotals(tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.transformQueryForTotals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderService.transformQueryForTotals() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("OrderService.transformQueryForTotals() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestOrderService_GetTotalsRelations(t *testing.T) {
	type args struct {
		config *sql.Options
	}
	tests := []struct {
		name string
		s    *OrderService
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetTotalsRelations(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.GetTotalsRelations() = %v, want %v", got, tt.want)
			}
		})
	}
}
