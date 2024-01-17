package utils

import (
	"github.com/google/uuid"
)

func ParseUUID(str string) (uuid.UUID, *ApplictaionError) {
	id, err := uuid.Parse(str)
	if err != nil {
		return uuid.Nil, NewApplictaionError(
			INVALID_DATA,
			err.Error(),
			nil,
		)
	}

	return id, nil
}

func ParseUUIDs(str []string) (uuid.UUIDs, *ApplictaionError) {
	var ids uuid.UUIDs
	for _, s := range str {
		id, err := uuid.Parse(s)
		if err != nil {
			return nil, NewApplictaionError(
				INVALID_DATA,
				err.Error(),
				nil,
			)
		}
		ids = append(ids, id)
	}

	return ids, nil
}
