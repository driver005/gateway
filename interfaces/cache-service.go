package interfaces

type ICacheService interface {
	Get(key string) (interface{}, error)
	Set(key string, data interface{}, ttl *int) error
	Invalidate(key string) error
}
