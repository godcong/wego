package cache

import (
	"github.com/go-redis/redis"
	"time"
)

// RedisCache ...
type RedisCache struct {
	client *redis.Client
}

// Get ...
func (r *RedisCache) Get(key string) interface{} {
	return r.GetD(key, nil)
}

// GetD ...
func (r *RedisCache) GetD(key string, v interface{}) interface{} {
	s, e := r.client.Get(key).Result()
	if e != nil || s == "" {
		return v
	}
	return s
}

// Set ...
func (r *RedisCache) Set(key string, val interface{}) Cache {
	r.client.Set(key, val, 0)
	return r
}

// SetWithTTL ...
func (r *RedisCache) SetWithTTL(key string, val interface{}, ttl int64) Cache {
	r.client.Set(key, val, time.Duration(ttl)*time.Second)
	return r
}

// Has ...
func (r *RedisCache) Has(key string) bool {
	_, err := r.client.Get(key).Result()
	if err == redis.Nil {
		return false
	}
	return true
}

// Delete ...
func (r *RedisCache) Delete(key string) Cache {
	r.client.Del(key)
	return r
}

// Clear ...
func (r *RedisCache) Clear() {
	r.client.FlushDB()
}

// GetMultiple ...
func (r *RedisCache) GetMultiple(keys ...string) map[string]interface{} {
	if keys == nil {
		return nil
	}
	val := r.client.MGet(keys...).Val()
	size := len(keys)
	result := make(map[string]interface{}, size)
	for i := 0; i < size; i++ {
		result[keys[i]] = val[i]
	}
	return result
}

// SetMultiple ...
func (r *RedisCache) SetMultiple(values map[string]interface{}) Cache {
	if values == nil {
		return cache
	}

	for key, value := range values {
		r.client.Set(key, value, 0)
	}
	return cache
}

// DeleteMultiple ...
func (r *RedisCache) DeleteMultiple(keys ...string) Cache {
	r.client.Del(keys...)
	return r
}

// Options ...
type Options struct {
	Addr string
}

// NewRedisCache ...
func NewRedisCache(op *redis.Options) *RedisCache {
	client := redis.NewClient(op)
	_, e := client.Ping().Result()
	if e != nil {
		panic(e)
	}
	return &RedisCache{client: client}
}
