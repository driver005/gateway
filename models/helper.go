package models

type Filter struct {
	Size   int    `json:"size"`
	Offset int    `json:"offset"`
	Name   string `json:"client_name"`
	Owner  string `json:"owner"`
}
