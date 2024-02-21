package core

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UUIDSlice uuid.UUIDs

func (UUIDSlice) GormDataType() string {
	return "string"
}

func (uuids *UUIDSlice) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("faild to scan UUIDSlice: %v", value)
	}

	if str == "" {
		return nil
	}

	strSplit := strings.Split(str[1:len(str)-1], ",")
	for _, uuidStr := range strSplit {
		if uuidStr == "" {
			continue
		}

		id, err := uuid.Parse(str)
		if err != nil {
			return err
		}

		*uuids = append(*uuids, id)
	}

	return nil
}

func (uuids *UUIDSlice) Value() (driver.Value, error) {
	var uuidStrs []string
	for _, u := range *uuids {
		uuidStrs = append(uuidStrs, u.String())
	}

	if len(uuidStrs) == 0 {
		return nil, nil
	}

	return fmt.Sprintf("{%s}", strings.Join(uuidStrs, ",")), nil
}
