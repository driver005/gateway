package core

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONB Interface for JSONB Field of yourTableName Table
type JSONB map[string]interface{}

// Value Marshal
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

func (a *JSONB) Add(key string, value interface{}) JSONB {
	data := *a
	data[key] = value
	a = &data
	return data
}
