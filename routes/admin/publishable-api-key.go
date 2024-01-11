package admin

type PublishableApiKey struct {
	r Registry
}

func NewPublishableApiKey(r Registry) *PublishableApiKey {
	m := PublishableApiKey{r: r}
	return &m
}
