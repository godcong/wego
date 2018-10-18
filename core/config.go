package core

import (
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"github.com/pelletier/go-toml"
)

//const FileLoadError = "cannot find config file"
//const ConfigReadError = "cannot read config file"

/*Tree Config Tree */
type Tree toml.Tree

//Config Application config interface */
type Config interface {
	Get(s string) interface{}
	GetD(s string, d interface{}) interface{}
	Set(k string, v interface{}) *Tree
	GetString(k string) string
	GetStringD(k, d string) string
	GetBool(s string) bool
	GetBoolD(s string, d bool) bool
	GetInt(s string) int64
	GetIntD(s string, d int64) int64
	GetSubConfig(s string) Config
	GetTree(s string) interface{}
}

/*LoadConfig get config tree with file name*/
func LoadConfig(f string) (Config, error) {
	t, e := toml.LoadFile(f)
	if e != nil {
		log.Println("filepath: " + f)
		log.Println(e.Error())
		return nil, e
	}
	return (*Tree)(t), nil
}

/*GetSubConfig get sub config from current config */
func (t *Tree) GetSubConfig(s string) Config {
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

/*Get get an interface from config */
func (t *Tree) Get(s string) interface{} {
	return t.GetTree(s)
}

/*GetD get interface with default value */
func (t *Tree) GetD(s string, d interface{}) interface{} {
	v := t.GetTree(s)
	if v == nil {
		return d
	}
	return v
}

/*GetString get string with out default value */
func (t *Tree) GetString(s string) string {
	v := t.GetTree(s)
	if v, b := v.(string); b {
		return v
	}
	return ""
}

/*GetString get string with out default value */
func (t *Tree) GetStringD(s, d string) string {
	v := t.GetTree(s)
	if v, b := v.(string); b {
		return v
	}
	return d
}

/*Set set value */
func (t *Tree) Set(k string, v interface{}) *Tree {
	(*toml.Tree)(t).Set(k, v)
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

//GetBoolD get bool with default value
func (t *Tree) GetBoolD(s string, d bool) bool {
	v := t.GetTree(s)
	if v, b := v.(bool); b {
		return v
	}

	return d
}

//GetInt get int value
func (t *Tree) GetInt(s string) int64 {
	v := t.GetTree(s)
	v0, b := util.ParseInt(v)
	if !b {
		return 0
	}
	return v0
}

//GetIntD get int with default value
func (t *Tree) GetIntD(s string, d int64) int64 {
	v := t.GetTree(s)
	v0, b := util.ParseInt(v)
	if !b {
		return d
	}
	return v0
}

func (t *Tree) Has(key string) bool {
	if t == nil {
		return false
	}

	return (*toml.Tree)(t).Has(key)
}

//NewConfig create a null config
func NewConfig() Config {
	return &Tree{}
}
