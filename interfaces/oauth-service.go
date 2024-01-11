package interfaces

// IOauthService is an interface for file connectors
type IOauthService interface {
	GenerateToken(code string) (map[string]interface{}, error)
	RefreshToken(code string) (map[string]interface{}, error)
	DestroyToken(code string) (map[string]interface{}, error)
}
