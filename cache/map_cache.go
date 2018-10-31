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

/*NewMapCache NewMapCache */
func NewMapCache() Cache {
	c := &MapCache{}
	return c
}

/*Get check exist */
func (m *MapCache) Get(key string) interface{} {
	return m.GetD(key, nil)
}

/*Set check exist */
func (m *MapCache) Set(key string, val interface{}) Cache {
	return m.SetWithTTL(key, val, nil)
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
func (m *MapCache) SetWithTTL(key string, val interface{}, ttl *time.Time) Cache {
	m.Store(key, &mapCacheData{
		value: val,
		life:  ttl,
	})
	return m
}

/*Has check exist */
func (m *MapCache) Has(key string) bool {
	_, b := m.Load(key)
	return b
}

/*Delete Delete one value */
func (m *MapCache) Delete(key string) Cache {
	m.Delete(key)
	return m
}

/*Clear delete all values */
func (m *MapCache) Clear() {
	*m = MapCache{}
}

/*GetMultiple get multiple values */
func (m *MapCache) GetMultiple(keys []string) map[string]interface{} {
	c := make(map[string]interface{})
	for _, k := range keys {
		if tmp := m.Get(k); tmp != nil {
			c[k] = tmp
		}
		c[k] = nil
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
func (m *MapCache) DeleteMultiple(keys []string) Cache {
	for _, k := range keys {
		m.Delete(k)
	}
	return m
}
