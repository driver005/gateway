package admin

type InventoryItem struct {
	r Registry
}

func NewInventoryItem(r Registry) *InventoryItem {
	m := InventoryItem{r: r}
	return &m
}
