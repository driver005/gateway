package core

import (
	"time"

	"github.com/google/uuid"
)

type NumberModel struct {
	N   int `json:"n,omitempty"`
	Lt  int `json:"lt,omitempty"`
	Gt  int `json:"gt,omitempty"`
	Gte int `json:"gte,omitempty"`
	Lte int `json:"lte,omitempty"`
}

type StringModel struct {
	N   string `json:"n,omitempty"`
	Lt  string `json:"lt,omitempty"`
	Gt  string `json:"gt,omitempty"`
	Gte string `json:"gte,omitempty"`
	Lte string `json:"lte,omitempty"`
}

type TimeModel struct {
	N   time.Time `json:"n,omitempty"`
	Lt  time.Time `json:"lt,omitempty"`
	Gt  time.Time `json:"gt,omitempty"`
	Gte time.Time `json:"gte,omitempty"`
	Lte time.Time `json:"lte,omitempty"`
}

type FilterModel struct {
	Id        []uuid.UUID `json:"id,omitempty"`
	CreatedAt *TimeModel  `json:"created_at,omitempty"`
	UpdatedAt *TimeModel  `json:"updated_at,omitempty"`
	DeletedAt *TimeModel  `json:"deleted_at,omitempty"`
}
