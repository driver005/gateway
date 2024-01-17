package admin

import (
	"context"
	"slices"

	"github.com/driver005/gateway/interfaces"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

func buildLevelsByInventoryItemId(inventoryLevels []interfaces.InventoryLevelDTO, locationIds uuid.UUIDs) map[uuid.UUID][]interfaces.InventoryLevelDTO {
	filteredLevels := make([]interfaces.InventoryLevelDTO, 0)
	for _, level := range inventoryLevels {
		if len(locationIds) == 0 || slices.Contains(locationIds, level.LocationId) {
			filteredLevels = append(filteredLevels, level)
		}
	}

	levelsByInventoryItemId := make(map[uuid.UUID][]interfaces.InventoryLevelDTO)
	for _, level := range filteredLevels {
		levelsByInventoryItemId[level.InventoryItemId] = append(levelsByInventoryItemId[level.InventoryItemId], level)
	}

	return levelsByInventoryItemId
}

func getLevelsByInventoryItemId(items []interfaces.InventoryItemDTO, locationIds uuid.UUIDs, inventoryService interfaces.IInventoryService) (map[uuid.UUID][]interfaces.InventoryLevelDTO, *utils.ApplictaionError) {
	selector := interfaces.FilterableInventoryLevelProps{
		InventoryItemId: lo.Map(items, func(inventoryItem interfaces.InventoryItemDTO, index int) uuid.UUID {
			return inventoryItem.Id
		}),
	}
	if len(locationIds) > 0 {
		selector.LocationId = locationIds
	}

	levels, _, err := inventoryService.ListInventoryLevels(context.Background(), selector, &sql.Options{})
	if err != nil {
		return nil, err
	}

	levelsWithAvailability := make([]interfaces.InventoryLevelDTO, len(levels))
	for i, level := range levels {
		availability, err := inventoryService.RetrieveAvailableQuantity(context.Background(), level.InventoryItemId, uuid.UUIDs{level.LocationId})
		if err != nil {
			return nil, err
		}
		level.AvailableQuantity = availability
		levelsWithAvailability[i] = level
	}

	return buildLevelsByInventoryItemId(levelsWithAvailability, locationIds), nil
}

func joinLevels(inventoryItems []interfaces.InventoryItemDTO, locationIds uuid.UUIDs, inventoryService interfaces.IInventoryService) ([]ResponseInventoryItem, *utils.ApplictaionError) {
	levelsByItemId, err := getLevelsByInventoryItemId(inventoryItems, locationIds, inventoryService)
	if err != nil {
		return nil, err
	}

	var responseInventoryItems []ResponseInventoryItem
	for _, inventoryItem := range inventoryItems {
		levels := levelsByItemId[inventoryItem.Id]
		if levels == nil {
			levels = []interfaces.InventoryLevelDTO{}
		}

		var reservedQuantity, stockedQuantity int
		for _, level := range levels {
			reservedQuantity += level.ReservedQuantity
			stockedQuantity += level.StockedQuantity
		}

		responseInventoryItem := ResponseInventoryItem{
			InventoryItemDTO: inventoryItem,
			ReservedQuantity: reservedQuantity,
			StockedQuantity:  stockedQuantity,
			LocationLevels:   levels,
		}

		responseInventoryItems = append(responseInventoryItems, responseInventoryItem)
	}

	return responseInventoryItems, nil
}
