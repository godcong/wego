package cache

import (
	"sync"
)

type MapCache struct {
	sync.Map
}

func (m *MapCache) GetD(key string, v0 interface{}) interface{} {

	if v, b := m.Load(key); b {
		return v
	}
	return v0
}

//TODO: ttl not set
func (m *MapCache) SetWithTTL(key string, val interface{}, ttl int) Cache {
	m.Store(key, val)
	return m
}

func NewMapCache() Cache {
	c := &MapCache{}

	return c
}

func (m *MapCache) Get(key string) interface{} {
	if v, b := m.Load(key); b {
		return v
	}
	return nil
}

func (m *MapCache) Set(key string, val interface{}) Cache {
	m.Store(key, val)

	return m
}

func (m *MapCache) Has(key string) bool {

	_, b := m.Load(key)
	return b
}

func (m *MapCache) Delete(key string) Cache {

	m.Delete(key)
	return m
}

func (m *MapCache) Clear() {
	*m = MapCache{}
}

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

func (m *MapCache) SetMultiple(values map[string]interface{}) {

	for k, v := range values {
		m.Set(k, v)
	}

}

func (m *MapCache) DeleteMultiple(keys []string) Cache {
	for _, k := range keys {
		m.Delete(k)
	}
	return m
}
