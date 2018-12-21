package cache_test

import (
	"log"
	"testing"
	"time"

	"github.com/godcong/wego/cache"
)

// TestMapCache_SetWithTTL ...
func TestMapCache_SetWithTTL(t *testing.T) {
	tm := time.Now()
	c := cache.DefaultCache().SetWithTTL("hello", "nihao", &tm)
	c = cache.DefaultCache().SetWithTTL("hello1", "nihao1", &tm)
	log.Println(c.Get("hello"))
	time.Sleep(time.Duration(10) * time.Second)
	log.Println(c.Get("hello"))
	log.Println(c.Get("hello1"))
}
