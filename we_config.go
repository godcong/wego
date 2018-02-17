package wego

import (
	"flag"
	"log"
	"strconv"

	"github.com/pelletier/go-toml"
)

const FileLoadError = "cannot find config file"
const ConfigReadError = "cannot read config file"

type Tree toml.Tree

var f = flag.String("f", "config.toml", "config file path")

//var useCache = false
var system System
var useCache = false
var configCache *Tree

func init() {
	flag.Parse()
	config := initLoader()
	useCache = system.UseCache
	if UseCache() {
		configCache = config
	}
	initLog(system)
	initDomain(GetConfig("domain"))
	Println(config)
}

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

type config struct {
	//System  `toml:"system"`
	Content *Tree
	//OfficialAccount map[string]ConfigMap `toml:"official_account"`
	//OpenPlatform    map[string]ConfigMap `toml:"open_platform"`
	//MiniProgram     map[string]ConfigMap `toml:"mini_program"`
	//Payment         map[string]ConfigMap `toml:"payment"`
	//Work            map[string]ConfigMap `toml:"work"`
}

//func (c *config) UseCache() bool {
//	if c == nil {
//		return false
//	}
//	return c.System.UseCache
//}

type Config interface {
	Get(s string) string
	GetBool(s string) bool
}

func ConfigTree() *Tree {
	t, e := toml.LoadFile(*f)
	if e != nil {
		log.Println("filepath: " + *f)
		log.Println(e.Error())
		panic(FileLoadError)
	}
	return (*Tree)(t)
}

func initLoader() *Tree {
	t := ConfigTree()
	t.GetTree("system").(*toml.Tree).Unmarshal(&system)
	return t
}

//
//func TransToMap(tree *toml.Tree, name string) map[string]Tree {
//	cm := make(map[string]Tree)
//	mp := tree.Get(name).(*toml.Tree).ToMap()
//	for key, value := range mp {
//		cm[key] = value.(map[string]Tree
//	}
//	return cm
//}

func ConfigLoader() *Tree {
	if system.UseCache {
		return configCache
	}
	return ConfigTree()
}

func GetConfig(s string) Config {
	c := ConfigLoader()
	if v, b := c.GetTree(s).(*toml.Tree); b {
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

	return strconv.Itoa(ParseInt(v))
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
