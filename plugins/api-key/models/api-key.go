package models

import "github.com/driver005/gateway/core"

type ApiKey struct {
	core.BaseModel
}

func (*ApiKey) TableName() string {
	return "api_key"
}
