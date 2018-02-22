package cache_test

import (
	"log"
	"testing"
	"time"

	"github.com/godcong/wego/cache"
)

func TestMapCache_SetWithTTL(t *testing.T) {
	c := cache.GetCache().SetWithTTL("hello", "nihao", 10)
	c = cache.GetCache().SetWithTTL("hello1", "nihao1", 100)
	log.Println(c.Get("hello"))
	time.Sleep(time.Duration(10) * time.Second)
	log.Println(c.Get("hello"))
	log.Println(c.Get("hello1"))
}
