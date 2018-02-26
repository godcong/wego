package core

import (
	"log"
	"strconv"

	"github.com/godcong/wego/cache"
	"github.com/pelletier/go-toml"
)

const FileLoadError = "cannot find config file"
const ConfigReadError = "cannot read config file"

type Tree toml.Tree

type System struct {
	//debug = true
	Debug bool `toml:"debug"`
	//response_type = 'array'
	ResponseType string `toml:"response_type"`
	//use_cache = true
	UseCache bool `toml:"use_cache"`
	Log      Log
}

type Log struct {
	//level = 'debug'
	Level string
	//file = 'logs/wechat.log'
	File string
}

//type config struct {
//	Content *Tree
//}

type Config interface {
	Get(s string) string
	Set(k, v string) *Tree
	GetBool(s string) bool
	GetConfig(s string) Config
	GetTree(s string) interface{}
}

var useCache = false

func ConfigTree(f string) *Tree {
	t, e := toml.LoadFile(f)
	if e != nil {
		log.Println("filepath: " + f)
		log.Println(e.Error())
		panic(FileLoadError)
	}
	return (*Tree)(t)
}

func treeLoader() *Tree {
	c := cache.GetCache()
	if UseCache() {
		return c.Get("cache").(*Tree)
	}
	return ConfigTree(c.GetD("cache_path", "config.toml").(string))
}

func GetConfig(s string) Config {
	c := treeLoader()
	if v, b := c.GetTree(s).(*toml.Tree); b {
		return (*Tree)(v)
	}
	return nil
}

func GetRootConfig() Config {
	return treeLoader()
}

func (t *Tree) GetConfig(s string) Config {
	if v, b := t.GetTree(s).(*toml.Tree); b {
		return (*Tree)(v)
	}
	return nil
}

func (t *Tree) GetTree(s string) interface{} {
	return (*toml.Tree)(t).Get(s)
}

func (t *Tree) Get(s string) string {
	v := t.GetTree(s)
	if v, b := v.(string); b {
		return v
	}
	if v0 := ParseInt(v); v0 == -1 {
		return ""
	} else {
		return strconv.FormatInt(v0, 10)
	}

}

func (t *Tree) Set(k, v string) *Tree {
	tt := (*toml.Tree)(t)
	tt.Set(k, v)
	return t
}

func (t *Tree) GetBool(s string) bool {
	v := t.GetTree(s)
	if v, b := v.(bool); b {
		return v
	}

	return false
}

func CacheOn() {
	useCache = true
}

func CacheOff() {
	useCache = false
}

func UseCache() bool {
	return useCache
}
