package interfaces

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
)

type IFulfillmentService interface {
	GetIdentifier() string
	GetFulfillmentOptions() (map[string]interface{}, *utils.ApplictaionError)
	ValidateFulfillmentData(optionData *models.ShippingOption, data map[string]interface{}, cart *models.Cart) map[string]interface{}
	ValidateOption(data models.ShippingOption) bool
	CanCalculate(data map[string]interface{}) bool
	CalculatePrice(optionData *models.ShippingOption, data map[string]interface{}, cart *models.Cart) float64
	CreateFulfillment(data *models.ShippingMethod, items []models.LineItem, order *types.CreateFulfillmentOrder, fulfillment *models.Fulfillment) core.JSONB
	CancelFulfillment(fulfillment *models.Fulfillment) *models.Fulfillment
	GetFulfillmentDocuments(data interface{}) []interface{}
	CreateReturn(fromData *models.Return) core.JSONB
	GetReturnDocuments(data interface{}) []interface{}
	GetShipmentDocuments(data interface{}) []interface{}
	RetrieveDocuments(fulfillmentData map[string]interface{}, documentType string) map[string]interface{}
}
