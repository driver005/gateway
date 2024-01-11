package admin

type StockLocation struct {
	r Registry
}

func NewStockLocation(r Registry) *StockLocation {
	m := StockLocation{r: r}
	return &m
}
