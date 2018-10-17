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
	Get(s string) string
	GetD(s, d string) string
	Set(k, v string) *Tree
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

/*Get get string with out default value */
func (t *Tree) Get(s string) string {
	v := t.GetTree(s)
	if v, b := v.(string); b {
		return v
	}
	return ""
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
