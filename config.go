package wego

import (
	"github.com/pelletier/go-toml"
)

// Config 配置文件，用来生成Property各种属性
type Config struct {
	AppID       string
	AppSecret   string
	MchID       string
	MchKey      string
	PemCert     []byte
	PemKEY      []byte
	RootCA      []byte
	Token       string
	AesKey      string
	Scopes      []string
	RedirectURI string
}

// DefaultConfig ...
func DefaultConfig() *Config {
	return &Config{}
}

// LoadConfig ...
func LoadConfig(path string) *Config {
	cfg := DefaultConfig()
	t, e := toml.LoadFile(path)
	if e != nil {
		log.Info("filepath: " + path)
		log.Info(e.Error())
		return DefaultConfig()
	}

	e = t.Unmarshal(cfg)
	if e != nil {
		log.Info("filepath: " + path)
		log.Info(e.Error())
		return DefaultConfig()
	}

	return cfg
}
