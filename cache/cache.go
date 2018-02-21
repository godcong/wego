package cache

type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}) Cache
	Has(key string) bool
	Delete(key string) Cache
	Clear()
	GetMultiple(keys []string) map[string]interface{}
	SetMultiple(values map[string]interface{})
	DeleteMultiple(keys []string) Cache
}

var cache map[string]Cache

const DefaultCacheName = "map_cache"

func init() {
	RegisterCache(DefaultCacheName, NewMapCache())
}

func RegisterCache(name string, c Cache) {
	if cache == nil {
		cache = make(map[string]Cache)
	}
	cache[name] = c
}

func GetCache(name ...string) Cache {
	if name != nil {
		if v, b := cache[name[0]]; b {
			return v
		}
	}
	return cache[DefaultCacheName]
}
