package admin

type Notification struct {
	r Registry
}

func NewNotification(r Registry) *Notification {
	m := Notification{r: r}
	return &m
}
