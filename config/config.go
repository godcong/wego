package config

import (
	"strconv"

	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"github.com/pelletier/go-toml"
)

//const FileLoadError = "cannot find config file"
//const ConfigReadError = "cannot read config file"

/*Tree Tree */
type Tree toml.Tree

var useCache = true

/*Config Config */
type Config interface {
	Get(s string) string
	GetD(s, d string) string
	Set(k, v string) *Tree
	GetBool(s string) bool
	GetConfig(s string) Config
	GetTree(s string) interface{}
}

/*GetConfigTree get config tree with file name*/
func GetConfigTree(f string) (*Tree, error) {
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
		t, b := c.Get("cache").(*Tree)
		if t != nil && b {
			return t
		}
	}

	t, err := GetConfigTree("config.toml")
	if err != nil {
		log.Error(err)
		return nil
	}
	if UseCache() {
		c.Set("cache", t)
	}
	return t
}

/*GetConfig get config with path */
func GetConfig(path string) Config {
	log.Debug("GetConfig|path", path)
	c := treeLoader()
	if v, b := c.GetTree(path).(*toml.Tree); b {
		return (*Tree)(v)
	}
	return (*Tree)(nil)
}

/*GetRootConfig get root config */
func GetRootConfig() Config {
	return treeLoader()
}

/*GetConfig get config or sub config */
func (t *Tree) GetConfig(s string) Config {
	if v, b := t.GetTree(s).(*toml.Tree); b {
		return (*Tree)(v)
	}
	return (*Tree)(nil)
}

/*GetTree get config tree */
func (t *Tree) GetTree(s string) interface{} {
	if t == nil {
		return nil
	}
	return (*toml.Tree)(t).Get(s)
}

/*Get get string with out default value */
func (t *Tree) Get(s string) string {
	v := t.GetTree(s)
	if v, b := v.(string); b {
		return v
	}
	v0 := util.ParseInt(v)
	if v0 == 0 {
		return ""
	}
	return strconv.FormatInt(v0, 10)

}

/*GetD get string with default value */
func (t *Tree) GetD(s, d string) string {
	if v := t.Get(s); v != "" {
		return v
	}
	return d
}

/*Set set string value */
func (t *Tree) Set(k, v string) *Tree {
	tt := (*toml.Tree)(t)
	tt.Set(k, v)
	return t
}

/*GetBool get bool value */
func (t *Tree) GetBool(s string) bool {
	v := t.GetTree(s)
	if v, b := v.(bool); b {
		return v
	}

	return false
}

/*CacheOn turn on cache */
func CacheOn() {
	useCache = true
}

/*CacheOff turn off cache */
func CacheOff() {
	useCache = false
}

/*UseCache return cache status */
func UseCache() bool {
	return useCache
}
