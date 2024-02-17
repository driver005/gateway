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

func TestNewClaimService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *ClaimService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClaimService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClaimService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClaimService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *ClaimService
		args args
		want *ClaimService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClaimService_Retrieve(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *ClaimService
		args  args
		want  *models.ClaimOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ClaimService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClaimService_List(t *testing.T) {
	type args struct {
		selector *models.ClaimOrder
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *ClaimService
		args  args
		want  []models.ClaimOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ClaimService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClaimService_Update(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.UpdateClaimInput
	}
	tests := []struct {
		name  string
		s     *ClaimService
		args  args
		want  *models.ClaimOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ClaimService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClaimService_ValidateCreateClaimInput(t *testing.T) {
	type args struct {
		data *types.CreateClaimInput
	}
	tests := []struct {
		name string
		s    *ClaimService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ValidateCreateClaimInput(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimService.ValidateCreateClaimInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClaimService_GetRefundTotalForClaimLinesOnOrder(t *testing.T) {
	type args struct {
		order      *models.Order
		claimItems []types.CreateClaimItemInput
	}
	tests := []struct {
		name  string
		s     *ClaimService
		args  args
		want  *float64
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetRefundTotalForClaimLinesOnOrder(tt.args.order, tt.args.claimItems)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimService.GetRefundTotalForClaimLinesOnOrder() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ClaimService.GetRefundTotalForClaimLinesOnOrder() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClaimService_Create(t *testing.T) {
	type args struct {
		data *types.CreateClaimInput
	}
	tests := []struct {
		name  string
		s     *ClaimService
		args  args
		want  *models.ClaimOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ClaimService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClaimService_CreateFulfillment(t *testing.T) {
	type args struct {
		id             uuid.UUID
		noNotification bool
		locationId     uuid.UUID
		metadata       map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *ClaimService
		args  args
		want  *models.ClaimOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateFulfillment(tt.args.id, tt.args.noNotification, tt.args.locationId, tt.args.metadata)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimService.CreateFulfillment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ClaimService.CreateFulfillment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClaimService_CancelFulfillment(t *testing.T) {
	type args struct {
		fulfillmentId uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ClaimService
		args  args
		want  *models.ClaimOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CancelFulfillment(tt.args.fulfillmentId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimService.CancelFulfillment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ClaimService.CancelFulfillment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClaimService_ProcessRefund(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ClaimService
		args  args
		want  *models.ClaimOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.ProcessRefund(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimService.ProcessRefund() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ClaimService.ProcessRefund() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClaimService_CreateShipment(t *testing.T) {
	type args struct {
		id             uuid.UUID
		fulfillmentId  uuid.UUID
		trackingLinks  []models.TrackingLink
		noNotification bool
		metadata       map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *ClaimService
		args  args
		want  *models.ClaimOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateShipment(tt.args.id, tt.args.fulfillmentId, tt.args.trackingLinks, tt.args.noNotification, tt.args.metadata)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimService.CreateShipment() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ClaimService.CreateShipment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClaimService_Cancel(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name  string
		s     *ClaimService
		args  args
		want  *models.ClaimOrder
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Cancel(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClaimService.Cancel() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ClaimService.Cancel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
