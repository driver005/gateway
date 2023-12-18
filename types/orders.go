package types

import "github.com/driver005/gateway/models"

func IsOrder(object interface{}) bool {
	order, ok := object.(models.Order)
	return ok && order.Object == "order"
}
