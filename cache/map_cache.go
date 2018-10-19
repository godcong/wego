package cache

import (
	"sync"
	"time"
)

/*MapCache MapCache */
type MapCache struct {
	sync.Map
}

/*NewMapCache NewMapCache */
func NewMapCache() Cache {
	c := &MapCache{}
	return c
}

/*Get check exist */
func (m *MapCache) Get(key string) interface{} {
	if v, b := m.Load(key); b {
		return v
	}
	return nil
}

/*Set check exist */
func (m *MapCache) Set(key string, val interface{}) Cache {
	m.Store(key, val)

	return m
}

/*GetD get interface with default */
func (m *MapCache) GetD(key string, v0 interface{}) interface{} {

	if v, b := m.Load(key); b {
		return v
	}
	return v0
}

/*SetWithTTL set interface with ttl */
func (m *MapCache) SetWithTTL(key string, val interface{}, ttl time.Time) Cache {
	//TODO: ttl not set
	m.Store(key, val)
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
func (m *MapCache) SetMultiple(values map[string]interface{}) {

	for k, v := range values {
		m.Set(k, v)
	}

}

/*DeleteMultiple delete multiple values */
func (m *MapCache) DeleteMultiple(keys []string) Cache {
	for _, k := range keys {
		m.Delete(k)
	}
	return m
}
