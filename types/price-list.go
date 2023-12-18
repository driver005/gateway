package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
)

type FilterablePriceList struct {
	core.FilterModel

	Q              string                   `json:"q,omitempty"`
	Status         []models.PriceListStatus `json:"status,omitempty"`
	Name           string                   `json:"name,omitempty"`
	CustomerGroups []string                 `json:"customer_groups,omitempty"`
	Description    string                   `json:"description,omitempty"`
	Type           []models.PriceListType   `json:"type,omitempty"`
}
