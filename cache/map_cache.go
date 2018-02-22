package cache

import (
	"sync"
)

type MapCache map[string]interface{}

var mutex sync.RWMutex

func (m *MapCache) GetD(key string, v0 interface{}) interface{} {
	v := m.Get(key)
	if v == nil {
		return v0
	}
	return v
}

//TODO: ttl not set
func (m *MapCache) SetWithTTL(key string, val interface{}, ttl int) Cache {
	mutex.Lock()
	(*m)[key] = val
	mutex.Unlock()
	return m
}

func NewMapCache() Cache {
	c := &MapCache{}
	return c
}

func (m *MapCache) Get(key string) interface{} {
	mutex.RLock()
	if v, b := (*m)[key]; b {
		return v
	}
	mutex.RUnlock()
	return nil
}

func (m *MapCache) Set(key string, val interface{}) Cache {
	mutex.Lock()
	(*m)[key] = val
	mutex.Unlock()
	return m
}

func (m *MapCache) Has(key string) bool {
	mutex.RLock()
	_, b := (*m)[key]
	mutex.RUnlock()
	return b
}

func (m *MapCache) Delete(key string) Cache {
	mutex.Lock()
	delete(*m, key)
	mutex.Unlock()
	return m
}

func (m *MapCache) Clear() {
	mutex.Lock()
	*m = make(MapCache)
	mutex.Unlock()
}

func (m *MapCache) GetMultiple(keys []string) map[string]interface{} {
	c := make(map[string]interface{})
	mutex.RLock()
	for _, k := range keys {
		if tmp, b := (*m)[k]; b {
			c[k] = tmp
		}
		c[k] = nil
	}
	mutex.RUnlock()
	return c
}

func (m *MapCache) SetMultiple(values map[string]interface{}) {
	mutex.Lock()
	for k, v := range values {
		(*m)[k] = v
	}
	mutex.Unlock()
}

func (m *MapCache) DeleteMultiple(keys []string) Cache {
	mutex.Lock()
	for _, k := range keys {
		delete(*m, k)
	}
	mutex.Unlock()
	return m
}
