package registry

import (
	"github.com/driver005/gateway/interfaces"
)

func (m *Base) PriceSelectionStrategy() interfaces.IPriceSelectionStrategy {
	if m.priceSelectionStrategy == nil {
		m.priceSelectionStrategy = nil
	}
	return m.priceSelectionStrategy
}
func (m *Base) TaxCalculationStrategy() interfaces.ITaxCalculationStrategy {
	if m.taxCalculationStrategy == nil {
		m.taxCalculationStrategy = nil
	}
	return m.taxCalculationStrategy
}
func (m *Base) InventoryService() interfaces.IInventoryService {
	if m.inventoryService == nil {
		m.inventoryService = nil
	}
	return m.inventoryService
}
func (m *Base) StockLocationService() interfaces.IStockLocationService {
	if m.stockLocationService == nil {
		m.stockLocationService = nil
	}
	return m.stockLocationService
}
func (m *Base) CacheService() interfaces.ICacheService {
	if m.cacheService == nil {
		m.cacheService = nil
	}
	return m.cacheService
}

func (m *Base) PricingModuleService() interfaces.IPricingModuleService {
	if m.pricingModuleService == nil {
		m.pricingModuleService = nil
	}
	return m.pricingModuleService
}

func (m *Base) FileService() interfaces.IFileService {
	if m.fileService == nil {
		m.fileService = nil
	}
	return m.fileService
}

func (m *Base) BatchJobStrategy() interfaces.IBatchJobStrategy {
	if m.batchJobStrategy == nil {
		m.batchJobStrategy = nil
	}
	return m.batchJobStrategy
}
