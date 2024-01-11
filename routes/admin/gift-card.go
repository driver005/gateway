package admin

type GiftCard struct {
	r Registry
}

func NewGiftCard(r Registry) *GiftCard {
	m := GiftCard{r: r}
	return &m
}
