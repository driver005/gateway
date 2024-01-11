package admin

type Reservation struct {
	r Registry
}

func NewReservation(r Registry) *Reservation {
	m := Reservation{r: r}
	return &m
}
