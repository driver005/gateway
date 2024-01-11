package admin

type Invite struct {
	r Registry
}

func NewInvite(r Registry) *Invite {
	m := Invite{r: r}
	return &m
}
