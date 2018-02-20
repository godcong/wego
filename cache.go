package wego

type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}) Cache
	Has(key string) bool
	Delete(key string) Cache
}
