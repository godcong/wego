package config

import (
	"strconv"

	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"github.com/pelletier/go-toml"
)

const FileLoadError = "cannot find config file"
const ConfigReadError = "cannot read config file"

type Tree toml.Tree

var useCache = false

//type config struct {
//	Content *Tree
//}

type Config interface {
	Get(s string) string
	GetD(s, d string) string
	Set(k, v string) *Tree
	GetBool(s string) bool
	GetConfig(s string) Config
	GetTree(s string) interface{}
}

func ConfigTree(f string) (*Tree, error) {
	t, e := toml.LoadFile(f)
	if e != nil {
		log.Println("filepath: " + f)
		log.Println(e.Error())
		return nil, e
	}
	return (*Tree)(t), nil
}

func treeLoader() *Tree {
	c := cache.GetCache()
	if UseCache() {
		return c.Get("cache").(*Tree)
	}

	t, err := ConfigTree(c.GetD("cache_path", "config.toml").(string))
	if err != nil {
		log.Error(err)
		return nil
	}
	return t
}

func GetConfig(path string) Config {
	log.Debug("GetConfig|path", path)
	c := treeLoader()
	if v, b := c.GetTree(path).(*toml.Tree); b {
		return (*Tree)(v)
	}
	return (*Tree)(nil)
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
	if t == nil {
		return nil
	}
	return (*toml.Tree)(t).Get(s)
}

func (t *Tree) Get(s string) string {
	v := t.GetTree(s)
	if v, b := v.(string); b {
		return v
	}
	if v0 := util.ParseInt(v); v0 == 0 {
		return ""
	} else {
		return strconv.FormatInt(v0, 10)
	}
}

func (t *Tree) GetD(s, d string) string {
	if v := t.Get(s); v != "" {
		return v
	}
	return d
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
