package wego

import (
	"flag"
	"log"

	"github.com/pelletier/go-toml"
)

const FileLoadError = "cannot find config file"
const ConfigReadError = "cannot read config file"

type Tree = toml.Tree

var f = flag.String("f", "config.toml", "config file path")
var useCache = false
var configCache config

func init() {
	flag.Parse()
	config := initLoader()
	useCache = config.UseCache()
	if useCache {
		configCache = config
	}
	log.Println(config)
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
	System  `toml:"system"`
	Content *Tree
	//OfficialAccount map[string]ConfigMap `toml:"official_account"`
	//OpenPlatform    map[string]ConfigMap `toml:"open_platform"`
	//MiniProgram     map[string]ConfigMap `toml:"mini_program"`
	//Payment         map[string]ConfigMap `toml:"payment"`
	//Work            map[string]ConfigMap `toml:"work"`
}

func (c *config) UseCache() bool {
	if c == nil {
		return false
	}
	return c.System.UseCache
}

type Config interface {
	//UseCache() bool
	//Debug() bool

}

func ConfigTree() *Tree {
	t, e := toml.LoadFile(*f)
	if e != nil {
		log.Println("filepath: " + *f)
		log.Println(e.Error())
		panic(FileLoadError)
	}
	return t
}

func initLoader() config {
	c := config{}
	t := ConfigTree()
	t.Get("system").(*Tree).Unmarshal(&c.System)

	c.Content = t
	return c
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

func ConfigLoader(s string) *Tree {
	if useCache {
		return configCache.Content
	}
	return ConfigTree()
}
