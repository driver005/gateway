package types

import "github.com/google/uuid"

type RegionDetails struct {
	Id      uuid.UUID
	TaxRate float64
}
