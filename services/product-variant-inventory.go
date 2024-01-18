package services

import (
	"context"
	"fmt"
	"math"
	"slices"
	"sync"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
)

type ReserveQuantityContext struct {
	LocationId     uuid.UUID
	LineItemId     uuid.UUID
	SalesChannelId uuid.UUID
}

type AvailabilityContext struct {
	variantInventoryMap  map[uuid.UUID][]models.ProductVariantInventoryItem
	inventoryLocationMap map[uuid.UUID][]interfaces.InventoryLevelDTO
}

type ProductVariantInventoryService struct {
	ctx context.Context
	r   Registry
}

func NewProductVariantInventoryService(
	r Registry,
) *ProductVariantInventoryService {
	return &ProductVariantInventoryService{
		context.Background(),
		r,
	}
}

func (s *ProductVariantInventoryService) SetContext(context context.Context) *ProductVariantInventoryService {
	s.ctx = context
	return s
}

func (s *ProductVariantInventoryService) ConfirmInventory(variantId uuid.UUID, quantity int, context map[string]interface{}) (bool, *utils.ApplictaionError) {
	if variantId == uuid.Nil {
		return false, utils.NewApplictaionError(
			utils.INVALID_DATA,
			`"variantId" must be defined`,
			nil,
		)
	}

	productVariant, err := s.r.ProductVariantService().SetContext(s.ctx).Retrieve(variantId, &sql.Options{
		Selects: []string{
			"id",
			"allow_backorder",
			"manage_inventory",
			"inventory_quantity",
		},
	})
	if err != nil {
		return false, err
	}

	if productVariant.AllowBackorder || !productVariant.ManageInventory {
		return true, nil
	}

	if s.r.InventoryService() == nil {
		return productVariant.InventoryQuantity >= quantity, nil
	}

	variantInventory, err := s.ListByVariant(uuid.UUIDs{variantId})
	if err != nil {
		return false, err
	}

	if len(variantInventory) == 0 {
		return true, nil
	}

	var locationIds uuid.UUIDs
	if context["salesChannelId"] != nil {
		locationIds, err = s.r.SalesChannelLocationService().SetContext(s.ctx).ListLocationIds(uuid.UUIDs{context["salesChannelId"].(uuid.UUID)})
		if err != nil {
			return false, err
		}
	} else {
		stockLocations, err := s.r.StockLocationService().List(s.ctx, interfaces.FilterableStockLocation{}, &sql.Options{
			Selects: []string{"id"},
		})
		if err != nil {
			return false, err
		}
		for _, l := range stockLocations {
			locationIds = append(locationIds, l.Id)
		}
	}

	if len(locationIds) == 0 {
		return false, nil
	}

	hasInventory := make([]bool, len(variantInventory))
	for i, inventoryPart := range variantInventory {
		itemQuantity := inventoryPart.RequiredQuantity * quantity
		hasInventory[i], err = s.r.InventoryService().ConfirmInventory(s.ctx, inventoryPart.InventoryItemId.UUID, locationIds, itemQuantity)
		if err != nil {
			return false, err
		}
	}

	for _, inventory := range hasInventory {
		if !inventory {
			return false, nil
		}
	}

	return true, nil
}

func (s *ProductVariantInventoryService) Retrieve(inventoryItemId uuid.UUID, variantId uuid.UUID) (*models.ProductVariantInventoryItem, *utils.ApplictaionError) {
	var res *models.ProductVariantInventoryItem

	query := sql.BuildQuery(models.ProductVariantInventoryItem{InventoryItemId: uuid.NullUUID{UUID: inventoryItemId}, VariantId: uuid.NullUUID{UUID: variantId}}, &sql.Options{})

	if err := s.r.ProductVariantInventoryItemRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductVariantInventoryService) ListByItem(itemIds uuid.UUIDs) ([]models.ProductVariantInventoryItem, *utils.ApplictaionError) {
	var res []models.ProductVariantInventoryItem

	query := sql.BuildQuery(models.ProductVariantInventoryItem{}, &sql.Options{
		Specification: []sql.Specification{sql.In("inventory_item_id", itemIds)},
	})

	if err := s.r.ProductVariantInventoryItemRepository().Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductVariantInventoryService) ListByVariant(variantIds uuid.UUIDs) ([]models.ProductVariantInventoryItem, *utils.ApplictaionError) {
	var res []models.ProductVariantInventoryItem

	query := sql.BuildQuery(models.ProductVariantInventoryItem{}, &sql.Options{
		Specification: []sql.Specification{sql.In("variant_id", variantIds)},
	})

	if err := s.r.ProductVariantInventoryItemRepository().Find(s.ctx, res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductVariantInventoryService) ListVariantsByItem(itemId uuid.UUID) ([]models.ProductVariant, *utils.ApplictaionError) {
	if s.r.InventoryService() == nil {
		return nil, nil
	}

	variantInventorys, err := s.ListByItem(uuid.UUIDs{itemId})
	if err != nil {
		return nil, err
	}

	var ids uuid.UUIDs

	for _, variantInventory := range variantInventorys {
		ids = append(ids, variantInventory.Id)
	}

	items, err := s.r.ProductVariantService().SetContext(s.ctx).List(&types.FilterableProductVariant{}, &sql.Options{Specification: []sql.Specification{sql.In("id", ids)}})
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *ProductVariantInventoryService) ListInventoryItemsByVariant(variantId uuid.UUID) ([]interfaces.InventoryItemDTO, *utils.ApplictaionError) {
	if s.r.InventoryService() == nil {
		return nil, nil
	}

	variantInventorys, err := s.ListByVariant(uuid.UUIDs{variantId})
	if err != nil {
		return nil, err
	}

	var ids uuid.UUIDs

	for _, variantInventory := range variantInventorys {
		ids = append(ids, variantInventory.Id)
	}

	items, _, err := s.r.InventoryService().ListInventoryItems(s.ctx, interfaces.FilterableInventoryItemProps{}, &sql.Options{Specification: []sql.Specification{sql.In("id", ids)}})
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *ProductVariantInventoryService) AttachInventoryItem(data []models.ProductVariantInventoryItem) ([]models.ProductVariantInventoryItem, *utils.ApplictaionError) {

	var invalidDataEntries []models.ProductVariantInventoryItem
	for _, d := range data {
		if d.RequiredQuantity < 1 {
			invalidDataEntries = append(invalidDataEntries, d)
		}
	}

	if len(invalidDataEntries) > 0 {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("\"requiredQuantity\" must be greater than 0, the following entries are invalid: %v", invalidDataEntries),
			nil,
		)
	}

	var variantIds uuid.UUIDs

	for _, d := range data {
		variantIds = append(variantIds, d.VariantId.UUID)
	}

	variants, err := s.r.ProductVariantService().SetContext(s.ctx).List(&types.FilterableProductVariant{}, &sql.Options{Selects: []string{"id"}, Specification: []sql.Specification{sql.In("id", variantIds)}})
	if err != nil {
		return nil, err
	}

	var foundVariantIds uuid.UUIDs
	for _, v := range variants {
		foundVariantIds = append(foundVariantIds, v.Id)
	}

	var requestedVariantIds uuid.UUIDs
	for _, d := range data {
		requestedVariantIds = append(requestedVariantIds, d.VariantId.UUID)
	}

	if len(foundVariantIds) != len(requestedVariantIds) {
		difference := utils.GetSetDifference(requestedVariantIds, foundVariantIds)
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("Variants not found for the following ids: %v", difference),
			nil,
		)
	}

	var itemIds uuid.UUIDs

	for _, d := range data {
		itemIds = append(itemIds, d.InventoryItemId.UUID)
	}

	inventoryItems, _, err := s.r.InventoryService().ListInventoryItems(s.ctx, interfaces.FilterableInventoryItemProps{}, &sql.Options{Selects: []string{"id"}, Specification: []sql.Specification{sql.In("id", itemIds)}})
	if err != nil {
		return nil, err
	}

	var foundInventoryItemIds uuid.UUIDs
	for _, v := range inventoryItems {
		foundInventoryItemIds = append(foundInventoryItemIds, v.Id)
	}

	var requestedInventoryItemIds uuid.UUIDs
	for _, d := range data {
		requestedInventoryItemIds = append(requestedInventoryItemIds, d.InventoryItemId.UUID)
	}

	if len(foundInventoryItemIds) != len(requestedInventoryItemIds) {
		difference := utils.GetSetDifference(requestedInventoryItemIds, foundInventoryItemIds)
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			fmt.Sprintf("Inventory items not found for the following ids: %v", difference),
			nil,
		)
	}

	var existingAttachments []models.ProductVariantInventoryItem

	query := sql.BuildQuery(models.ProductVariantInventoryItem{}, &sql.Options{
		Specification: []sql.Specification{sql.In("inventory_item_id", itemIds), sql.In("variant_id", variantIds)},
	})

	if err := s.r.ProductVariantInventoryItemRepository().Find(s.ctx, existingAttachments, query); err != nil {
		return nil, err
	}

	existingMap := make(map[uuid.UUID]uuid.UUIDs)
	for _, curr := range existingAttachments {
		if set := existingMap[curr.VariantId.UUID]; set != nil {
			existingMap[curr.VariantId.UUID] = append(existingMap[curr.VariantId.UUID], curr.InventoryItemId.UUID)
		} else {
			existingMap[curr.VariantId.UUID] = uuid.UUIDs{curr.InventoryItemId.UUID}
		}
	}

	var toCreate []models.ProductVariantInventoryItem
	for _, d := range data {
		if slices.Contains(existingMap[d.VariantId.UUID], d.InventoryItemId.UUID) {
			continue
		}

		toCreate = append(toCreate, d)
	}

	if err = s.r.ProductVariantInventoryItemRepository().SaveSlice(s.ctx, toCreate); err != nil {
		return nil, err
	}

	return toCreate, nil
}

func (s *ProductVariantInventoryService) DetachInventoryItem(inventoryItemId uuid.UUID, variantId uuid.UUID) *utils.ApplictaionError {
	var res []models.ProductVariantInventoryItem
	var selector models.ProductVariantInventoryItem

	selector.InventoryItemId = uuid.NullUUID{UUID: inventoryItemId}

	if variantId != uuid.Nil {
		selector.VariantId = uuid.NullUUID{UUID: variantId}
	}

	query := sql.BuildQuery(models.ProductVariantInventoryItem{}, &sql.Options{
		Specification: []sql.Specification{sql.In("variant_id", uuid.UUIDs{variantId})},
	})

	if err := s.r.ProductVariantInventoryItemRepository().Find(s.ctx, res, query); err != nil {
		return err
	}

	if len(res) > 0 {
		if err := s.r.ProductVariantInventoryItemRepository().RemoveSlice(s.ctx, res); err != nil {
			return err
		}
	}

	return nil
}

func (s *ProductVariantInventoryService) ReserveQuantity(variantId uuid.UUID, quantity int, context ReserveQuantityContext) ([]interfaces.ReservationItemDTO, *utils.ApplictaionError) {
	if s.r.InventoryService() == nil {
		variant, err := s.r.ProductVariantService().SetContext(s.ctx).Retrieve(variantId, &sql.Options{
			Selects: []string{"id", "inventory_quantity"},
		})
		if err != nil {
			return nil, err
		}

		_, err = s.r.ProductVariantService().SetContext(s.ctx).Update(variant.Id, nil, &types.UpdateProductVariantInput{
			InventoryQuantity: variant.InventoryQuantity - quantity,
		})
		if err != nil {
			return nil, err
		}
	}

	variantInventory, err := s.ListByVariant(uuid.UUIDs{variantId})
	if err != nil {
		return nil, err
	}

	if len(variantInventory) == 0 {
		return nil, nil
	}

	locationId := context.LocationId

	if locationId == uuid.Nil && context.SalesChannelId == uuid.Nil {
		locationIds, err := s.r.SalesChannelLocationService().SetContext(s.ctx).ListLocationIds(uuid.UUIDs{context.SalesChannelId})
		if err != nil {
			return nil, err
		}

		if len(locationIds) == 0 {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Must provide location_id or sales_channel_id to a Sales Channel that has associated Stock Locations",
				"500",
				nil,
			)
		}

		locations, count, err := s.r.InventoryService().ListInventoryLevels(s.ctx, interfaces.FilterableInventoryLevelProps{LocationId: locationIds, InventoryItemId: uuid.UUIDs{variantInventory[0].InventoryItemId.UUID}}, &sql.Options{})
		if err != nil {
			return nil, err
		}

		if count == gox.NewInt64(0) {
			return nil, utils.NewApplictaionError(
				utils.INVALID_DATA,
				"Must provide location_id or sales_channel_id to a Sales Channel that has associated locations with inventory levels",
				"500",
				nil,
			)
		}

		locationId = locations[0].LocationId
	}

	var reservationItems []interfaces.ReservationItemDTO
	for _, inventoryPart := range variantInventory {
		itemQuantity := inventoryPart.RequiredQuantity * quantity
		reservationItem, err := s.r.InventoryService().CreateReservationItems(s.ctx, []interfaces.CreateReservationItemInput{
			{
				LineItemId:      context.LineItemId,
				LocationId:      locationId,
				InventoryItemId: inventoryPart.InventoryItemId.UUID,
				Quantity:        itemQuantity,
			},
		})
		if err != nil {
			return nil, err
		}

		reservationItems = append(reservationItems, reservationItem...)
	}

	return reservationItems, nil
}

func (s *ProductVariantInventoryService) AdjustReservationsQuantityByLineItem(lineItemId uuid.UUID, variantId uuid.UUID, locationId uuid.UUID, quantity int) *utils.ApplictaionError {
	if s.r.InventoryService() == nil {
		variant, err := s.r.ProductVariantService().SetContext(s.ctx).Retrieve(variantId, &sql.Options{
			Selects: []string{"id", "inventory_quantity", "manage_inventory"},
		})
		if err != nil {
			return err
		}

		if !variant.ManageInventory {
			return nil
		}

		_, err = s.r.ProductVariantService().SetContext(s.ctx).Update(variant.Id, nil, &types.UpdateProductVariantInput{
			InventoryQuantity: variant.InventoryQuantity - quantity,
		})
		if err != nil {
			return err
		}
	}

	if quantity > 0 {
		return utils.NewApplictaionError(
			utils.INVALID_DATA,
			"You can only reduce reservation quantities using AdjustReservationsQuantityByLineItem. If you wish to reserve more use Update or Create.",
			nil,
		)
	}

	reservations, reservationCount, err := s.r.InventoryService().ListReservationItems(s.ctx, interfaces.FilterableReservationItemProps{LineItemId: lineItemId}, &sql.Options{Order: gox.NewString("created_at DESC")})
	if err != nil {
		return err
	}

	slices.SortFunc(reservations, func(a interfaces.ReservationItemDTO, b interfaces.ReservationItemDTO) int {
		if a.LocationId == locationId {
			return -1
		}

		return 0
	})

	if *reservationCount > 0 {
		inventoryItems, err := s.ListByVariant(uuid.UUIDs{variantId})
		if err != nil {
			return err
		}

		productVariantInventory := inventoryItems[0]
		deltaUpdate := int(math.Abs(float64(quantity * productVariantInventory.RequiredQuantity)))
		exactReservation := slices.IndexFunc(reservations, func(r interfaces.ReservationItemDTO) bool {
			return r.Quantity == deltaUpdate && r.LocationId == locationId
		})
		if exactReservation != -1 {
			err = s.r.InventoryService().DeleteReservationItem(s.ctx, uuid.UUIDs{reservations[exactReservation].Id})
			if err != nil {
				return err
			}

			return nil
		}

		remainingQuantity := deltaUpdate
		var reservationsToDelete []interfaces.ReservationItemDTO
		var reservationToUpdate *interfaces.ReservationItemDTO
		for _, reservation := range reservations {
			if reservation.Quantity <= remainingQuantity {
				remainingQuantity -= reservation.Quantity
				reservationsToDelete = append(reservationsToDelete, reservation)
			} else {
				reservationToUpdate = &reservation
				break
			}
		}

		if len(reservationsToDelete) > 0 {
			var reservationIds uuid.UUIDs

			for _, reservation := range reservationsToDelete {
				reservationIds = append(reservationIds, reservation.Id)
			}

			if err := s.r.InventoryService().DeleteReservationItem(s.ctx, reservationIds); err != nil {
				return err
			}
		}

		if reservationToUpdate != nil {
			_, err = s.r.InventoryService().UpdateReservationItem(s.ctx, reservationToUpdate.Id, interfaces.UpdateReservationItemInput{
				Quantity: reservationToUpdate.Quantity - remainingQuantity,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *ProductVariantInventoryService) ValidateInventoryAtLocation(items []models.LineItem, locationId uuid.UUID) *utils.ApplictaionError {
	if s.r.InventoryService() == nil {
		return nil
	}

	var itemsToValidate []models.LineItem
	for _, item := range items {
		if item.VariantId.UUID != uuid.Nil {
			itemsToValidate = append(itemsToValidate, item)
		}
	}

	for _, item := range itemsToValidate {
		pvInventoryItems, err := s.ListByVariant(uuid.UUIDs{item.VariantId.UUID})
		if err != nil {
			return err
		}

		if len(pvInventoryItems) == 0 {
			continue
		}

		var ids uuid.UUIDs

		for _, variantInventory := range pvInventoryItems {
			ids = append(ids, variantInventory.InventoryItemId.UUID)
		}

		inventoryLevels, inventoryLevelCount, err := s.r.InventoryService().ListInventoryLevels(s.ctx, interfaces.FilterableInventoryLevelProps{LocationId: uuid.UUIDs{locationId}}, &sql.Options{Specification: []sql.Specification{sql.In("inventory_item_id", ids)}})
		if err != nil {
			return err
		}

		if inventoryLevelCount == gox.NewInt64(0) {
			return utils.NewApplictaionError(
				utils.INVALID_DATA,
				fmt.Sprintf("Inventory item for %s not found at location", item.Title),
				"500",
				nil,
			)
		}

		pviMap := make(map[uuid.UUID]*models.ProductVariantInventoryItem)
		for _, pvi := range pvInventoryItems {
			pviMap[pvi.InventoryItemId.UUID] = &pvi
		}

		for _, inventoryLevel := range inventoryLevels {
			pvInventoryItem := pviMap[inventoryLevel.InventoryItemId]
			if pvInventoryItem == nil || pvInventoryItem.RequiredQuantity*item.Quantity > inventoryLevel.StockedQuantity {
				return utils.NewApplictaionError(
					utils.INVALID_DATA,
					fmt.Sprintf("Insufficient stock for item: %s", item.Title),
					"500",
					nil,
				)
			}
		}
	}

	return nil
}

func (s *ProductVariantInventoryService) DeleteReservationsByLineItem(lineItemId uuid.UUID, variantId uuid.UUID, quantity int) *utils.ApplictaionError {
	if s.r.InventoryService() == nil {
		variant, err := s.r.ProductVariantService().SetContext(s.ctx).Retrieve(variantId, &sql.Options{
			Selects: []string{"id", "inventory_quantity", "manage_inventory"},
		})
		if err != nil {
			return err
		}

		if !variant.ManageInventory {
			return nil
		}

		_, err = s.r.ProductVariantService().SetContext(s.ctx).Update(variant.Id, nil, &types.UpdateProductVariantInput{
			InventoryQuantity: variant.InventoryQuantity - quantity,
		})
		if err != nil {
			return err
		}
	}

	err := s.r.InventoryService().DeleteReservationItemsByLineItem(s.ctx, uuid.UUIDs{lineItemId})
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductVariantInventoryService) AdjustInventory(variantId uuid.UUID, locationId string, quantity int) *utils.ApplictaionError {
	if s.r.InventoryService() == nil {
		variant, err := s.r.ProductVariantService().SetContext(s.ctx).Retrieve(variantId, &sql.Options{
			Selects: []string{"id", "inventory_quantity", "manage_inventory"},
		})
		if err != nil {
			return err
		}

		if !variant.ManageInventory {
			return nil
		}

		_, err = s.r.ProductVariantService().SetContext(s.ctx).Update(variant.Id, nil, &types.UpdateProductVariantInput{
			InventoryQuantity: variant.InventoryQuantity - quantity,
		})
		if err != nil {
			return err
		}
	}
	variantInventory, err := s.ListByVariant(uuid.UUIDs{variantId})
	if err != nil {
		return err
	}
	if len(variantInventory) == 0 {
		return nil
	}
	var wg sync.WaitGroup
	for _, inventoryPart := range variantInventory {
		itemQuantity := inventoryPart.RequiredQuantity * quantity
		wg.Add(1)
		err := func(inventoryItemId uuid.UUID, locationId string, itemQuantity int) *utils.ApplictaionError {
			defer wg.Done()
			_, err := s.r.InventoryService().AdjustInventory(s.ctx, inventoryItemId, locationId, itemQuantity)
			if err != nil {
				return err
			}
			return nil
		}(inventoryPart.InventoryItemId.UUID, locationId, itemQuantity)
		if err != nil {
			return err
		}
	}
	wg.Wait()
	return nil
}

func (s *ProductVariantInventoryService) SetVariantAvailability(variants []models.ProductVariant, salesChannelId uuid.UUIDs, availabilityContext *AvailabilityContext) ([]models.ProductVariant, *utils.ApplictaionError) {
	if s.r.InventoryService() == nil {
		return variants, nil
	}
	context, err := s.GetAvailabilityContext(variants, salesChannelId, availabilityContext)
	if err != nil {
		return nil, err
	}
	for i, variant := range variants {
		if variant.Id == uuid.Nil {
			continue
		}
		variant.Purchasable = variant.AllowBackorder
		if !variant.ManageInventory {
			variant.Purchasable = true
			variants[i] = variant
			continue
		}
		variantInventory, ok := context.variantInventoryMap[variant.Id]
		if !ok || len(variantInventory) == 0 {
			variant.InventoryQuantity = 0
			variant.Purchasable = true
			variants[i] = variant
			continue
		}
		if salesChannelId == nil {
			variant.InventoryQuantity = 0
			variant.Purchasable = false
			variants[i] = variant
			continue
		}
		locations, ok := context.inventoryLocationMap[variantInventory[0].InventoryItemId.UUID]
		if !ok {
			continue
		}
		variant.InventoryQuantity = 0
		for _, location := range locations {
			variant.InventoryQuantity += location.StockedQuantity - location.ReservedQuantity
		}
		variant.Purchasable = variant.InventoryQuantity > 0 || variant.AllowBackorder
		variants[i] = variant
	}
	return variants, nil
}

func (s *ProductVariantInventoryService) GetAvailabilityContext(variants []models.ProductVariant, salesChannelId uuid.UUIDs, existingContext *AvailabilityContext) (*AvailabilityContext, *utils.ApplictaionError) {
	variantInventoryMap := existingContext.variantInventoryMap
	inventoryLocationMap := existingContext.inventoryLocationMap
	if variantInventoryMap != nil {
		variantInventoryMap = make(map[uuid.UUID][]models.ProductVariantInventoryItem)
		var variantIds uuid.UUIDs
		for _, variant := range variants {
			variantIds = append(variantIds, variant.Id)
		}
		variantInventories, err := s.ListByVariant(variantIds)
		if err != nil {
			return nil, err
		}
		for _, inventory := range variantInventories {
			currentInventories, ok := variantInventoryMap[inventory.VariantId.UUID]
			if !ok {
				currentInventories = []models.ProductVariantInventoryItem{}
			}
			currentInventories = append(currentInventories, inventory)
			variantInventoryMap[inventory.VariantId.UUID] = currentInventories
		}
	}
	var locationIds uuid.UUIDs
	if salesChannelId != nil {
		if inventoryLocationMap == nil {
			inventoryLocationMap = make(map[uuid.UUID][]interfaces.InventoryLevelDTO)
		}
		locations, err := s.r.SalesChannelLocationService().SetContext(s.ctx).ListLocationIds(salesChannelId)
		if err != nil {
			return nil, err
		}
		locationIds = append(locationIds, locations...)
	}
	if inventoryLocationMap == nil {
		inventoryLocationMap = make(map[uuid.UUID][]interfaces.InventoryLevelDTO)
	}
	if len(locationIds) > 0 {
		var variantIds uuid.UUIDs
		for _, variant := range variantInventoryMap {
			for _, v := range variant {
				variantIds = append(variantIds, v.InventoryItemId.UUID)
			}
		}
		locationLevels, _, err := s.r.InventoryService().ListInventoryLevels(s.ctx, interfaces.FilterableInventoryLevelProps{}, &sql.Options{
			Specification: []sql.Specification{
				sql.In("location_id", locationIds),
				sql.In("inventory_item_id", variantIds),
			},
		})
		if err != nil {
			return nil, err
		}
		for _, locationLevel := range locationLevels {
			if _, ok := inventoryLocationMap[locationLevel.InventoryItemId]; !ok {
				inventoryLocationMap[locationLevel.InventoryItemId] = []interfaces.InventoryLevelDTO{}
			}
			inventoryLocationMap[locationLevel.InventoryItemId] = append(inventoryLocationMap[locationLevel.InventoryItemId], locationLevel)
		}
	}
	return &AvailabilityContext{
		variantInventoryMap:  variantInventoryMap,
		inventoryLocationMap: inventoryLocationMap,
	}, nil
}

func (s *ProductVariantInventoryService) SetProductAvailability(products []models.Product, salesChannelId uuid.UUIDs) ([]models.Product, *utils.ApplictaionError) {
	if s.r.InventoryService() == nil {
		return products, nil
	}
	var variants []models.ProductVariant
	for _, product := range products {
		for _, variant := range product.Variants {
			if variant.Id != uuid.Nil {
				variants = append(variants, variant)
			}
		}
	}
	availabilityContext, err := s.GetAvailabilityContext(variants, salesChannelId, nil)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	for _, product := range products {
		if len(product.Variants) == 0 {
			continue
		}
		wg.Add(1)
		err := func(product *models.Product) *utils.ApplictaionError {
			defer wg.Done()
			variants, err := s.SetVariantAvailability(product.Variants, salesChannelId, availabilityContext)
			if err != nil {
				return err
			}
			product.Variants = variants

			return nil
		}(&product)

		if err != nil {
			return nil, err
		}
	}
	wg.Wait()
	return products, nil
}

func (s *ProductVariantInventoryService) GetVariantQuantityFromVariantInventoryItems(variantInventoryItems []models.ProductVariantInventoryItem, channelId uuid.UUID) (*int, *utils.ApplictaionError) {
	variantItemsAreMixed := false
	for _, inventoryItem := range variantInventoryItems {
		if inventoryItem.VariantId != variantInventoryItems[0].VariantId {
			variantItemsAreMixed = true
			break
		}
	}
	if len(variantInventoryItems) > 0 && variantItemsAreMixed {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"All variant inventory items must belong to the same variant",
			nil,
		)
	}
	var quantities []int
	for _, variantInventory := range variantInventoryItems {
		quantity, err := s.r.SalesChannelInventoryService().SetContext(s.ctx).RetrieveAvailableItemQuantity(channelId, variantInventory.InventoryItemId.UUID)
		if err != nil {
			return nil, err
		}
		quantities = append(quantities, quantity/variantInventory.RequiredQuantity)
	}
	minQuantity := quantities[0]
	for _, quantity := range quantities {
		if quantity < minQuantity {
			minQuantity = quantity
		}
	}
	return &minQuantity, nil
}
