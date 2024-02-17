package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

func TestNewProductVariantInventoryService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *ProductVariantInventoryService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductVariantInventoryService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductVariantInventoryService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductVariantInventoryService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *ProductVariantInventoryService
		args args
		want *ProductVariantInventoryService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductVariantInventoryService_ConfirmInventory(t *testing.T) {
	type args struct {
		variantId uuid.UUID
		quantity  int
		context   map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *ProductVariantInventoryService
		args  args
		want  bool
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ConfirmInventory(tt.args.variantId, tt.args.quantity, tt.args.context)
			if got != tt.want {
				t.Errorf("ProductVariantInventoryService.ConfirmInventory() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantInventoryService.ConfirmInventory() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantInventoryService_Retrieve(t *testing.T) {
	type args struct {
		inventoryItemId uuid.UUID
		variantId       uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ProductVariantInventoryService
		args  args
		want  *models.ProductVariantInventoryItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.inventoryItemId, tt.args.variantId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantInventoryService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantInventoryService_ListByItem(t *testing.T) {
	type args struct {
		itemIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *ProductVariantInventoryService
		args  args
		want  []models.ProductVariantInventoryItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListByItem(tt.args.itemIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.ListByItem() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantInventoryService.ListByItem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantInventoryService_ListByVariant(t *testing.T) {
	type args struct {
		variantIds uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *ProductVariantInventoryService
		args  args
		want  []models.ProductVariantInventoryItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListByVariant(tt.args.variantIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.ListByVariant() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantInventoryService.ListByVariant() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantInventoryService_ListVariantsByItem(t *testing.T) {
	type args struct {
		itemId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ProductVariantInventoryService
		args  args
		want  []models.ProductVariant
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListVariantsByItem(tt.args.itemId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.ListVariantsByItem() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantInventoryService.ListVariantsByItem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantInventoryService_ListInventoryItemsByVariant(t *testing.T) {
	type args struct {
		variantId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ProductVariantInventoryService
		args  args
		want  []interfaces.InventoryItemDTO
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ListInventoryItemsByVariant(tt.args.variantId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.ListInventoryItemsByVariant() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantInventoryService.ListInventoryItemsByVariant() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantInventoryService_AttachInventoryItem(t *testing.T) {
	type args struct {
		data []models.ProductVariantInventoryItem
	}
	tests := []struct {
		name  string
		s     *ProductVariantInventoryService
		args  args
		want  []models.ProductVariantInventoryItem
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AttachInventoryItem(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.AttachInventoryItem() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantInventoryService.AttachInventoryItem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantInventoryService_DetachInventoryItem(t *testing.T) {
	type args struct {
		inventoryItemId uuid.UUID
		variantId       uuid.UUID
	}
	tests := []struct {
		name string
		s    *ProductVariantInventoryService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DetachInventoryItem(tt.args.inventoryItemId, tt.args.variantId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.DetachInventoryItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductVariantInventoryService_ReserveQuantity(t *testing.T) {
	type args struct {
		variantId uuid.UUID
		quantity  int
		context   ReserveQuantityContext
	}
	tests := []struct {
		name  string
		s     *ProductVariantInventoryService
		args  args
		want  []interfaces.ReservationItemDTO
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ReserveQuantity(tt.args.variantId, tt.args.quantity, tt.args.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.ReserveQuantity() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantInventoryService.ReserveQuantity() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantInventoryService_AdjustReservationsQuantityByLineItem(t *testing.T) {
	type args struct {
		lineItemId uuid.UUID
		variantId  uuid.UUID
		locationId uuid.UUID
		quantity   int
	}
	tests := []struct {
		name string
		s    *ProductVariantInventoryService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.AdjustReservationsQuantityByLineItem(tt.args.lineItemId, tt.args.variantId, tt.args.locationId, tt.args.quantity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.AdjustReservationsQuantityByLineItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductVariantInventoryService_ValidateInventoryAtLocation(t *testing.T) {
	type args struct {
		items      []models.LineItem
		locationId uuid.UUID
	}
	tests := []struct {
		name string
		s    *ProductVariantInventoryService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ValidateInventoryAtLocation(tt.args.items, tt.args.locationId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.ValidateInventoryAtLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductVariantInventoryService_DeleteReservationsByLineItem(t *testing.T) {
	type args struct {
		lineItemId uuid.UUID
		variantId  uuid.UUID
		quantity   int
	}
	tests := []struct {
		name string
		s    *ProductVariantInventoryService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DeleteReservationsByLineItem(tt.args.lineItemId, tt.args.variantId, tt.args.quantity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.DeleteReservationsByLineItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductVariantInventoryService_AdjustInventory(t *testing.T) {
	type args struct {
		variantId  uuid.UUID
		locationId uuid.UUID
		quantity   int
	}
	tests := []struct {
		name string
		s    *ProductVariantInventoryService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.AdjustInventory(tt.args.variantId, tt.args.locationId, tt.args.quantity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.AdjustInventory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductVariantInventoryService_SetVariantAvailability(t *testing.T) {
	type args struct {
		variants            []models.ProductVariant
		salesChannelId      uuid.UUIDs
		availabilityContext *AvailabilityContext
	}
	tests := []struct {
		name  string
		s     *ProductVariantInventoryService
		args  args
		want  []models.ProductVariant
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetVariantAvailability(tt.args.variants, tt.args.salesChannelId, tt.args.availabilityContext)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.SetVariantAvailability() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantInventoryService.SetVariantAvailability() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantInventoryService_GetAvailabilityContext(t *testing.T) {
	type args struct {
		variants        []models.ProductVariant
		salesChannelId  uuid.UUIDs
		existingContext *AvailabilityContext
	}
	tests := []struct {
		name  string
		s     *ProductVariantInventoryService
		args  args
		want  *AvailabilityContext
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetAvailabilityContext(tt.args.variants, tt.args.salesChannelId, tt.args.existingContext)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.GetAvailabilityContext() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantInventoryService.GetAvailabilityContext() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantInventoryService_SetProductAvailability(t *testing.T) {
	type args struct {
		products       []models.Product
		salesChannelId uuid.UUIDs
	}
	tests := []struct {
		name  string
		s     *ProductVariantInventoryService
		args  args
		want  []models.Product
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetProductAvailability(tt.args.products, tt.args.salesChannelId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.SetProductAvailability() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantInventoryService.SetProductAvailability() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductVariantInventoryService_GetVariantQuantityFromVariantInventoryItems(t *testing.T) {
	type args struct {
		variantInventoryItems []models.ProductVariantInventoryItem
		channelId             uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ProductVariantInventoryService
		args  args
		want  *int
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetVariantQuantityFromVariantInventoryItems(tt.args.variantInventoryItems, tt.args.channelId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductVariantInventoryService.GetVariantQuantityFromVariantInventoryItems() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductVariantInventoryService.GetVariantQuantityFromVariantInventoryItems() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
