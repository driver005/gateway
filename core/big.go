package core

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"
)

type BigJSON struct {
	*big.Int
}

func (b BigJSON) Value() (driver.Value, error) {
	if b.Int == nil {
		return nil, nil
	}
	return b.Int.String(), nil
}

func (b *BigJSON) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return json.Unmarshal(src, &b.Int)
	case string:
		return json.Unmarshal([]byte(src), &b.Int)
	default:
		return fmt.Errorf("unsupported type for BigJSON: %T", src)
	}
}
