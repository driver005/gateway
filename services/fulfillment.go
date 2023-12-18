package services

import (
	"context"

	"github.com/driver005/gateway/repository"
)

type FulfillmentService struct {
	ctx                    context.Context
	repo                   *repository.FulfillmentRepo
	trackingLinkRepository *repository.TrackingLinkRepo
	lineItemRepository     *repository.LineItemRepo
}

func NewFulfillmentService(
	ctx context.Context,
	repo *repository.FulfillmentRepo,
	trackingLinkRepository *repository.TrackingLinkRepo,
	lineItemRepository *repository.LineItemRepo,
) *FulfillmentService {
	return &FulfillmentService{
		ctx,
		repo,
		trackingLinkRepository,
		lineItemRepository,
	}
}

func (fs *FulfillmentService) PartitionItems(shippingMethods []ShippingMethod, items []LineItem) []FulfillmentItemPartition {
	partitioned := []FulfillmentItemPartition{}
	if len(shippingMethods) == 1 {
		return []FulfillmentItemPartition{{Items: items, ShippingMethod: shippingMethods[0]}}
	}
	for _, method := range shippingMethods {
		temp := FulfillmentItemPartition{ShippingMethod: method}
		methodProfile := method.ShippingOption.ProfileID
		for _, item := range items {
			if item.Variant.Product.ProfileID == methodProfile {
				temp.Items = append(temp.Items, item)
			}
		}
		partitioned = append(partitioned, temp)
	}
	return partitioned
}

func (fs *FulfillmentService) GetFulfillmentItems(order CreateFulfillmentOrder, items []FulFillmentItemType) ([]*LineItem, error) {
	toReturn := []*LineItem{}
	for _, item := range items {
		found := false
		for _, i := range order.Items {
			if i.ID == item.ItemID {
				found = true
				if item.Quantity > i.Quantity-i.FulfilledQuantity {
					return nil, errors.New("Cannot fulfill more items than have been purchased")
				}
				toReturn = append(toReturn, &LineItem{
					ID:               i.ID,
					Quantity:         item.Quantity,
					FulfilledQuantity: i.FulfilledQuantity,
				})
				break
			}
		}
		if !found {
			return nil, nil
		}
	}
	return toReturn, nil
}

func (fs *FulfillmentService) ValidateFulfillmentLineItem(item *LineItem, quantity int) (*LineItem, error) {
	if item == nil {
		return nil, nil
	}
	if quantity > item.Quantity-item.FulfilledQuantity {
		return nil, errors.New("Cannot fulfill more items than have been purchased")
	}
	return &LineItem{
		ID:               item.ID,
		Quantity:         quantity,
		FulfilledQuantity: item.FulfilledQuantity,
	}, nil
}

func (fs *FulfillmentService) Retrieve(fulfillmentID string, config FindConfig) (*Fulfillment, error) {
	if fulfillmentID == "" {
		return nil, errors.New(`"fulfillmentID" must be defined`)
	}
	fulfillment := &Fulfillment{ID: fulfillmentID}
	return fulfillment, nil
}

func (fs *FulfillmentService) CreateFulfillment(order CreateFulfillmentOrder, itemsToFulfill []FulFillmentItemType, custom Partial<Fulfillment>) ([]*Fulfillment, error) {
	fulfillments := fs.PartitionItems(order.ShippingMethods, order.Items)
	created := []*Fulfillment{}
	for _, f := range fulfillments {
		ful := &Fulfillment{
			ProviderID: f.ShippingMethod.ShippingOption.ProviderID,
			Items:      []*FulfillmentItem{},
			Data:       map[string]interface{}{},
		}
		for _, item := range f.Items {
			ful.Items = append(ful.Items, &FulfillmentItem{
				ItemID:   item.ID,
				Quantity: item.Quantity,
			})
		}
		created = append(created, ful)
	}
	return created, nil
}

func (fs *FulfillmentService) CancelFulfillment(fulfillmentOrID interface{}) (*Fulfillment, error) {
	var fulfillment *Fulfillment
	switch fulfillmentOrID.(type) {
	case string:
		fulfillment = &Fulfillment{ID: fulfillmentOrID.(string)}
	case *Fulfillment:
		fulfillment = fulfillmentOrID.(*Fulfillment)
	default:
		return nil, errors.New("Invalid fulfillmentOrID type")
	}
	if fulfillment.ShippedAt != nil {
		return nil, errors.New("The fulfillment has already been shipped. Shipped fulfillments cannot be canceled")
	}
	fulfillment.CanceledAt = new(time.Time)
	return fulfillment, nil
}

func (fs *FulfillmentService) CreateShipment(fulfillmentID string, trackingLinks []TrackingLink, config CreateShipmentConfig) (*Fulfillment, error) {
	fulfillment := &Fulfillment{ID: fulfillmentID}
	if fulfillment.CanceledAt != nil {
		return nil, errors.New("Fulfillment has been canceled")
	}
	now := time.Now()
	fulfillment.ShippedAt = &now
	fulfillment.TrackingLinks = trackingLinks
	fulfillment.NoNotification = config.NoNotification
	fulfillment.Metadata = config.Metadata
	return fulfillment, nil
}
