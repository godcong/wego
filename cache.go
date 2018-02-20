package wego

type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}) Cache
	Has(key string) bool
	Delete(key string) Cache
	Clear()
	GetMultiple(keys []string) map[string]interface{}
	SetMultiple(values map[string]interface{})
	DeleteMultiple(keys []string)
}

type cache struct {
	Config
	app Application
}

func (c *cache) Get(key string) interface{} {
	panic("implement me")
}

func (c *cache) Set(key string, val interface{}) Cache {
	panic("implement me")
}

func (c *cache) Has(key string) bool {
	panic("implement me")
}

func (c *cache) Delete(key string) Cache {
	panic("implement me")
}

func (c *cache) Clear() {
	panic("implement me")
}

func (c *cache) GetMultiple(keys []string) map[string]interface{} {
	panic("implement me")
}

func (c *cache) SetMultiple(values map[string]interface{}) {
	panic("implement me")
}

func (c *cache) DeleteMultiple(keys []string) {
	panic("implement me")
}

func NewCache(application Application, config Config) Cache {
	return &cache{
		Config: config,
		app:    application,
		//client: application.Client(),
	}
}
