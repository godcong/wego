package cache

import "time"

/*Cache define an cache interface */
type Cache interface {
	Get(key string) interface{}
	GetD(key string, v interface{}) interface{}
	Set(key string, val interface{}) Cache
	SetWithTTL(key string, val interface{}, ttl time.Time) Cache
	Has(key string) bool
	Delete(key string) Cache
	Clear()
	GetMultiple(keys []string) map[string]interface{}
	SetMultiple(values map[string]interface{})
	DeleteMultiple(keys []string) Cache
}

//var cache sync.Map
var cache Cache

/*DefaultCacheName defined the default cache */
const DefaultCacheName = "map_cache"

func init() {
	RegisterCache(DefaultCacheName, NewMapCache())
}

/*RegisterCache register cache to map */
func RegisterCache(name string, c Cache) {
	cache = c
}

/*GetCache get cache from map */
func GetCache(name ...string) Cache {
	return cache
}
