package cache

import (
	"sync"
	"time"
)

/*MapCache MapCache */
type MapCache struct {
	sync.Map
}

type mapCacheData struct {
	value interface{}
	life  *time.Time
}

func init() {
	RegisterCache(NewMapCache())
}

// NewMapCache ...
func NewMapCache() *MapCache {
	return &MapCache{}
}

/*Get check exist */
func (m *MapCache) Get(key string) interface{} {
	return m.GetD(key, nil)
}

/*Set check exist */
func (m *MapCache) Set(key string, val interface{}) Cache {
	return m.SetWithTTL(key, val, 0)
}

/*GetD get interface with default */
func (m *MapCache) GetD(key string, v0 interface{}) interface{} {
	if v, b := m.Load(key); b {
		switch vv := v.(type) {
		case *mapCacheData:
			if vv.life != nil && vv.life.Before(time.Now()) {
				return nil
			}
			return vv.value
		}
	}
	return nil
}

/*SetWithTTL set interface with ttl */
func (m *MapCache) SetWithTTL(key string, val interface{}, ttl int64) Cache {
	t := time.Now().Add(time.Duration(ttl))
	m.Store(key, &mapCacheData{
		value: val,
		life:  &t,
	})
	return m
}

/*Has check exist */
func (m *MapCache) Has(key string) bool {
	if v, b := m.Load(key); b {
		switch vv := v.(type) {
		case *mapCacheData:
			if vv.life != nil && vv.life.Before(time.Now()) {
				return false
			}
			return true
		}
	}
	return false
}

/*Delete one value */
func (m *MapCache) Delete(key string) Cache {
	m.Delete(key)
	return m
}

/*Clear delete all values */
func (m *MapCache) Clear() {
	*m = MapCache{}
}

/*GetMultiple get multiple values */
func (m *MapCache) GetMultiple(keys ...string) map[string]interface{} {
	c := make(map[string]interface{})
	size := len(keys)
	for i := 0; i < size; i++ {
		if tmp := m.Get(keys[i]); tmp != nil {
			c[keys[i]] = tmp
		}
		//c[keys[i]] = nil
	}

	return c
}

/*SetMultiple set multiple values */
func (m *MapCache) SetMultiple(values map[string]interface{}) Cache {
	for k, v := range values {
		m.Set(k, v)
	}
	return m
}

/*DeleteMultiple delete multiple values */
func (m *MapCache) DeleteMultiple(keys ...string) Cache {
	size := len(keys)
	for i := 0; i < size; i++ {
		m.Delete(keys[i])
	}

	return m
}
