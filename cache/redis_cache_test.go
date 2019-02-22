package cache

import (
	"github.com/go-redis/redis"
	"github.com/godcong/wego/util"
	"strconv"
	"testing"
)

// TestRedisCache_Clear ...
func TestRedisCache_Clear(t *testing.T) {
	rds := NewRedisCache(&redis.Options{
		Addr:     "localhost:6379",
		Password: "2rXfzaNKqX1b",
		DB:       1,
	})
	RegisterCache(rds)
	for i := 0; i < 100; i++ {
		cache.Set(strconv.Itoa(i), util.GenerateRandomString(32))
	}

}
