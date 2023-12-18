package models

import "database/sql/driver"

// The status of the Price List
type ClaimStatus string

// Defines values for ClaimStatus.
const (
	ClaimStatusReplace ClaimStatus = "replace"
	ClaimStatusRefund  ClaimStatus = "refund"
)

func (pl *ClaimStatus) Scan(value interface{}) error {
	*pl = ClaimStatus(value.([]byte))
	return nil
}

func (pl ClaimStatus) Value() (driver.Value, error) {
	return string(pl), nil
}
