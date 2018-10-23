package core

import (
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"github.com/pelletier/go-toml"
)

const configPath = "config.toml"

func init() {
	config, err := LoadConfig(configPath)
	if err != nil {
		log.Println("no config files loaded")
		log.Error(err)
		return
	}
	cache.Set("config", config)
}

/*Config Config Tree */
type Config struct {
	*toml.Tree
}

//DefaultConfig get the default config from cache
func DefaultConfig() *Config {
	if c := cache.Get("config"); c != nil {
		if v, b := c.(*Config); b {
			return v
		}
	}
	return nil
}

/*LoadConfig get config tree with file name*/
func LoadConfig(f string) (*Config, error) {
	t, e := toml.LoadFile(f)
	if e != nil {
		log.Println("filepath: " + f)
		log.Println(e.Error())
		return nil, e
	}
	return cfg(t), nil
}

/*GetSubConfig get sub config from current config */
func (t *Config) GetSubConfig(s string) *Config {
	if v, b := t.GetTree(s).(*toml.Tree); b {
		return cfg(v)
	}
	return cfg(nil)
}

/*GetTree get config tree */
func (t *Config) GetTree(s string) interface{} {
	if t == nil {
		return nil
	}

	return t.Tree.Get(s)
}

/*Get get an interface from config */
func (t *Config) Get(s string) interface{} {
	return t.GetTree(s)
}

/*GetD get interface with default value */
func (t *Config) GetD(s string, d interface{}) interface{} {
	v := t.GetTree(s)
	if v == nil {
		return d
	}
	return v
}

/*GetString get string with out default value */
func (t *Config) GetString(s string) string {
	v := t.GetTree(s)
	if v, b := v.(string); b {
		return v
	}
	return ""
}

/*GetStringD get string with default value */
func (t *Config) GetStringD(s, d string) string {
	v := t.GetTree(s)
	if v, b := v.(string); b {
		return v
	}
	return d
}

/*Set set value */
func (t *Config) Set(k string, v interface{}) *Config {
	t.Tree.Set(k, v)
	return t
}

/*GetBool get bool value */
func (t *Config) GetBool(s string) bool {
	v := t.GetTree(s)
	if v, b := v.(bool); b {
		return v
	}

	return false
}

//GetBoolD get bool with default value
func (t *Config) GetBoolD(s string, d bool) bool {
	v := t.GetTree(s)
	if v, b := v.(bool); b {
		return v
	}

	return d
}

//GetInt get int value
func (t *Config) GetInt(s string) int64 {
	v := t.GetTree(s)
	v0, b := util.ParseInt(v)
	if !b {
		return 0
	}
	return v0
}

//GetIntD get int with default value
func (t *Config) GetIntD(s string, d int64) int64 {
	v := t.GetTree(s)
	v0, b := util.ParseInt(v)
	if !b {
		return d
	}
	return v0
}

//Check check all input keys
//return 0 if all is exist
//return index when not found
func (t *Config) Check(arr ...string) int {
	for i, v := range arr {
		if t.Has(v) {
			continue
		}
		return i
	}
	return 0
}

//cfg create a null config
func cfg(tree *toml.Tree) *Config {
	return &Config{
		Tree: tree,
	}
}

//NewConfig create a new null config
func NewConfig() *Config {
	return &Config{
		Tree: &toml.Tree{},
	}
}
