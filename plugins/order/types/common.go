package types

import (
	"github.com/driver005/gateway/core"
)

type OrderSummary struct {
	Total                core.BigJSON `json:"total"`
	Subtotal             core.BigJSON `json:"subtotal"`
	TotalTax             core.BigJSON `json:"total_tax"`
	OrderedTotal         core.BigJSON `json:"ordered_total"`
	FulfilledTotal       core.BigJSON `json:"fulfilled_total"`
	ReturnedTotal        core.BigJSON `json:"returned_total"`
	ReturnRequestTotal   core.BigJSON `json:"return_request_total"`
	WriteOffTotal        core.BigJSON `json:"write_off_total"`
	ProjectedTotal       core.BigJSON `json:"projected_total"`
	NetTotal             core.BigJSON `json:"net_total"`
	NetSubtotal          core.BigJSON `json:"net_subtotal"`
	NetTotalTax          core.BigJSON `json:"net_total_tax"`
	FutureTotal          core.BigJSON `json:"future_total"`
	FutureSubtotal       core.BigJSON `json:"future_subtotal"`
	FutureTotalTax       core.BigJSON `json:"future_total_tax"`
	FutureProjectedTotal core.BigJSON `json:"future_projected_total"`
	Balance              core.BigJSON `json:"balance"`
	FutureBalance        core.BigJSON `json:"future_balance"`
}

type ItemSummary struct {
	ReturnableQuantity      core.BigJSON `json:"returnable_quantity"`
	OrderedQuantity         core.BigJSON `json:"ordered_quantity"`
	FulfilledQuantity       core.BigJSON `json:"fulfilled_quantity"`
	ReturnRequestedQuantity core.BigJSON `json:"return_requested_quantity"`
	ReturnReceivedQuantity  core.BigJSON `json:"return_received_quantity"`
	ReturnDismissedQuantity core.BigJSON `json:"return_dismissed_quantity"`
	WrittenOffQuantity      core.BigJSON `json:"written_off_quantity"`
}
