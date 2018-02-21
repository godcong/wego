package cache

type MapCache map[string]interface{}

func NewMapCache() Cache {
	return &MapCache{}
}

func (m *MapCache) Get(key string) interface{} {
	if v, b := (*m)[key]; b {
		return v
	}
	return nil
}

func (m *MapCache) Set(key string, val interface{}) Cache {
	(*m)[key] = val
	return m
}

func (m *MapCache) Has(key string) bool {
	_, b := (*m)[key]
	return b
}

func (m *MapCache) Delete(key string) Cache {
	delete(*m, key)
	return m
}

func (m *MapCache) Clear() {
	*m = make(MapCache)
}

func (m *MapCache) GetMultiple(keys []string) map[string]interface{} {
	c := make(map[string]interface{})
	for _, k := range keys {
		if tmp, b := (*m)[k]; b {
			c[k] = tmp
		}
		c[k] = nil
	}
	return c
}

func (m *MapCache) SetMultiple(values map[string]interface{}) {
	for k, v := range values {
		(*m)[k] = v
	}
}

func (m *MapCache) DeleteMultiple(keys []string) Cache {
	for _, k := range keys {
		delete(*m, k)
	}
	return m
}
