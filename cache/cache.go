package cache

/*Cache define an cache interface */
type Cache interface {
	Get(key string) interface{}
	GetD(key string, v interface{}) interface{}
	Set(key string, val interface{}) Cache
	SetWithTTL(key string, val interface{}, ttl int64) Cache
	Has(key string) bool
	Delete(key string) Cache
	Clear()
	GetMultiple(keys ...string) map[string]interface{}
	SetMultiple(values map[string]interface{}) Cache
	DeleteMultiple(keys ...string) Cache
}

//var cache sync.Map
var cache Cache

/*RegisterCache register cache to map */
func RegisterCache(c Cache) {
	cache = c
}

/*DefaultCache get cache from map */
func DefaultCache() Cache {
	return cache
}

//Get get value
func Get(key string) interface{} {
	return cache.Get(key)
}

//GetD get value with default
func GetD(key string, v interface{}) interface{} {
	return cache.GetD(key, v)
}

//Set set value
func Set(key string, val interface{}) Cache {
	return cache.Set(key, val)
}

//SetWithTTL set value with time to life
func SetWithTTL(key string, val interface{}, ttl int64) Cache {
	return cache.SetWithTTL(key, val, ttl)
}

//Has check value
func Has(key string) bool {
	return cache.Has(key)
}

//Delete delete value
func Delete(key string) Cache {
	return cache.Delete(key)
}

//Clear clear all
func Clear() {
	cache.Clear()
}

//GetMultiple get multiple value
func GetMultiple(keys ...string) map[string]interface{} {
	return cache.GetMultiple(keys...)
}

//SetMultiple set multiple value
func SetMultiple(values map[string]interface{}) Cache {
	return cache.SetMultiple(values)
}

//DeleteMultiple delete multiple value
func DeleteMultiple(keys ...string) Cache {
	return cache.DeleteMultiple(keys...)
}
