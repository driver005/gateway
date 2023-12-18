package models

type Filter struct {
	// The maximum amount of clients to returned, upper bound is 500 clients.
	// in: query
	Size int `json:"size"`

	// The offset from where to start looking.
	// in: query
	Offset int `json:"offset"`

	// The name of the clients to filter by.
	// in: query
	Name string `json:"client_name"`

	// The owner of the clients to filter by.
	// in: query
	Owner string `json:"owner"`
}
