package core

import (
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"github.com/pelletier/go-toml"
)

//const FileLoadError = "cannot find config file"
//const ConfigReadError = "cannot read config file"

/*Tree Config Tree */
type Config toml.Tree

//Config Application config interface */
//type Config interface {
//	Get(s string) interface{}
//	GetD(s string, d interface{}) interface{}
//	Set(k string, v interface{}) *Tree
//	GetString(k string) string
//	GetStringD(k, d string) string
//	GetBool(s string) bool
//	GetBoolD(s string, d bool) bool
//	GetInt(s string) int64
//	GetIntD(s string, d int64) int64
//	GetSubConfig(s string) Config
//	GetTree(s string) interface{}
//	Unmarshal(v interface{}) error
//}

/*LoadConfig get config tree with file name*/
func LoadConfig(f string) (*Config, error) {
	t, e := toml.LoadFile(f)
	if e != nil {
		log.Println("filepath: " + f)
		log.Println(e.Error())
		return nil, e
	}
	return (*Config)(t), nil
}

/*GetSubConfig get sub config from current config */
func (t *Config) GetSubConfig(s string) *Config {
	if v, b := t.GetTree(s).(*toml.Tree); b {
		return (*Config)(v)
	}
	return (*Config)(nil)
}

/*GetTree get config tree */
func (t *Config) GetTree(s string) interface{} {
	if t == nil {
		return nil
	}

	return (*toml.Tree)(t).Get(s)
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
	(*toml.Tree)(t).Set(k, v)
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

func (t *Config) Unmarshal(v interface{}) error {
	return (*toml.Tree)(t).Unmarshal(v)
}

//Has check config elements
func (t *Config) Has(key string) bool {
	if t == nil {
		return false
	}

	return (*toml.Tree)(t).Has(key)
}

//NewConfig create a null config
func NewConfig() *Config {
	return &Config{}
}
