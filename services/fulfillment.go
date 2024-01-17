package services

import (
	"context"
	"time"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

type FulfillmentService struct {
	ctx context.Context
	r   Registry
}

func NewFulfillmentService(
	r Registry,
) *FulfillmentService {
	return &FulfillmentService{
		context.Background(),
		r,
	}
}

func (s *FulfillmentService) SetContext(context context.Context) *FulfillmentService {
	s.ctx = context
	return s
}

func (s *FulfillmentService) PartitionItems(shippingMethods []models.ShippingMethod, items []models.LineItem) []types.FulfillmentItemPartition {
	partitioned := []types.FulfillmentItemPartition{}
	if len(shippingMethods) == 1 {
		return []types.FulfillmentItemPartition{{Items: items, ShippingMethod: &shippingMethods[0]}}
	}
	for _, method := range shippingMethods {
		temp := types.FulfillmentItemPartition{ShippingMethod: &method}
		methodProfile := method.ShippingOption.ProfileId
		for _, item := range items {
			if item.Variant.Product.ProfileId == methodProfile {
				temp.Items = append(temp.Items, item)
			}
		}
		partitioned = append(partitioned, temp)
	}
	return partitioned
}

func (s *FulfillmentService) GetFulfillmentItems(order *types.CreateFulfillmentOrder, items []types.FulFillmentItemType) ([]models.LineItem, *utils.ApplictaionError) {
	var toReturn []models.LineItem
	for _, item := range items {
		found := false
		for _, i := range order.AdditionalItems {
			if i.Id == item.ItemId {
				found = true
				if item.Quantity > i.Quantity-i.FulfilledQuantity {
					return nil, utils.NewApplictaionError(
						utils.CONFLICT,
						"Cannot fulfill more items than have been purchased",
						"500",
						nil,
					)
				}
				toReturn = append(toReturn, models.LineItem{
					Model:             core.Model{Id: i.Id},
					Quantity:          item.Quantity,
					FulfilledQuantity: i.FulfilledQuantity,
				})
				break
			}
		}
		if !found {
			return nil, utils.NewApplictaionError(
				utils.NOT_FOUND,
				"Cannot find any items",
				"500",
				nil,
			)
		}
	}
	return toReturn, nil
}

func (s *FulfillmentService) ValidateFulfillmentLineItem(item *models.LineItem, quantity int) (*models.LineItem, *utils.ApplictaionError) {
	if item == nil {
		return nil, utils.NewApplictaionError(
			utils.NOT_FOUND,
			"Cannot find any items",
			nil,
		)
	}
	if quantity > item.Quantity-item.FulfilledQuantity {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Cannot fulfill more items than have been purchased",
			nil,
		)
	}
	return &models.LineItem{
		Model:             core.Model{Id: item.Id},
		Quantity:          quantity,
		FulfilledQuantity: item.FulfilledQuantity,
	}, nil
}

func (s *FulfillmentService) Retrieve(fulfillmentId uuid.UUID, config *sql.Options) (*models.Fulfillment, *utils.ApplictaionError) {
	if fulfillmentId == uuid.Nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"fulfillmentId" must be defined`,
			nil,
		)
	}
	var res *models.Fulfillment
	query := sql.BuildQuery(models.Fulfillment{Model: core.Model{Id: fulfillmentId}}, config)
	if err := s.r.FulfillmentRepository().FindOne(s.ctx, res, query); err == nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`Fulfillment with id `+fulfillmentId.String()+` was not found`,
			nil,
		)
	}
	return res, nil
}

func (s *FulfillmentService) CreateFulfillment(order *types.CreateFulfillmentOrder, itemsToFulfill []types.FulFillmentItemType, custom models.Fulfillment) ([]models.Fulfillment, *utils.ApplictaionError) {
	lineItems, err := s.GetFulfillmentItems(order, itemsToFulfill)
	if err != nil {
		return nil, err
	}
	fulfillments := s.PartitionItems(order.ShippingMethods, lineItems)
	var created []models.Fulfillment
	for _, f := range fulfillments {
		ful := &models.Fulfillment{
			ProviderId: f.ShippingMethod.ShippingOption.ProviderId,
			Items:      []models.FulfillmentItem{},
			Data:       map[string]interface{}{},
		}
		for _, item := range f.Items {
			ful.Items = append(ful.Items, models.FulfillmentItem{
				ItemId:   uuid.NullUUID{UUID: item.Id},
				Quantity: item.Quantity,
			})
		}

		if err := s.r.FulfillmentRepository().Save(s.ctx, ful); err != nil {
			return nil, err
		}

		var err *utils.ApplictaionError
		ful.Data, err = s.r.FulfillmentProviderService().SetContext(s.ctx).CreateFulfillment(f.ShippingMethod, f.Items, order, ful)
		if err != nil {
			return nil, err
		}

		if err := s.r.FulfillmentRepository().Save(s.ctx, ful); err != nil {
			return nil, err
		}

		created = append(created, *ful)
	}
	return created, nil
}

func (s *FulfillmentService) CancelFulfillment(fulfillmentId uuid.UUID, fulfillment *models.Fulfillment) (*models.Fulfillment, *utils.ApplictaionError) {
	var id uuid.UUID
	if fulfillmentId == uuid.Nil {
		id = fulfillmentId
	} else if fulfillment != nil {
		id = fulfillment.Id
	} else {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"Invalid fulfillmentOrID type",
			nil,
		)
	}
	fulfillment, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	if fulfillment.ShippedAt != nil {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"The fulfillment has already been shipped. Shipped fulfillments cannot be canceled",
			nil,
		)
	}

	if _, err := s.r.FulfillmentProviderService().SetContext(s.ctx).CancelFulfillment(fulfillment); err != nil {
		return nil, err
	}

	fulfillment.CanceledAt = new(time.Time)

	for _, item := range fulfillment.Items {
		litem, err := s.r.LineItemService().SetContext(s.ctx).Retrieve(item.ItemId.UUID, &sql.Options{})
		if err != nil {
			return nil, err
		}
		fulfilledQuantity := litem.FulfilledQuantity - item.Quantity

		_, err = s.r.LineItemService().SetContext(s.ctx).Update(litem.Id, nil, &models.LineItem{FulfilledQuantity: fulfilledQuantity}, &sql.Options{})
		if err != nil {
			return nil, err
		}
	}

	if err := s.r.FulfillmentRepository().Save(s.ctx, fulfillment); err != nil {
		return nil, err
	}

	return fulfillment, nil
}

func (s *FulfillmentService) CreateShipment(fulfillmentId uuid.UUID, trackingLinks []models.TrackingLink, config *types.CreateShipmentConfig) (*models.Fulfillment, *utils.ApplictaionError) {
	fulfillment, err := s.Retrieve(fulfillmentId, &sql.Options{})
	if err != nil {
		return nil, err
	}
	if fulfillment.CanceledAt != nil {
		return nil, utils.NewApplictaionError(
			utils.CONFLICT,
			"models.Fulfillment has been canceled",
			nil,
		)
	}

	for _, link := range trackingLinks {
		if err := s.r.TrackingLinkRepository().Create(s.ctx, &link); err != nil {
			return nil, err
		}
	}
	fulfillment.TrackingLinks = trackingLinks

	now := time.Now()
	fulfillment.ShippedAt = &now

	fulfillment.NoNotification = config.NoNotification
	fulfillment.Metadata = config.Metadata
	return fulfillment, nil
}
