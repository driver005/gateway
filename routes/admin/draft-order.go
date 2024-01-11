package admin

type DraftOrder struct {
	r Registry
}

func NewDraftOrder(r Registry) *DraftOrder {
	m := DraftOrder{r: r}
	return &m
}
