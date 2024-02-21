package models

type Filter struct {
	Size   int    `json:"size" gorm:"column:size"`
	Offset int    `json:"offset" gorm:"column:offset"`
	Name   string `json:"client_name" gorm:"column:client_name"`
	Owner  string `json:"owner" gorm:"column:owner"`
}
