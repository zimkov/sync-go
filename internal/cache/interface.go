package cache

type Cache interface {
	Get(key string) (string, bool)
	Put(key, value string)
}
