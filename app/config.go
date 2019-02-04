package app

import (
	"github.com/pelletier/go-toml"
	"log"
)

// Configure ...
type Configure struct {
	LocalHost   string
	PaidURL     string
	RefundedURL string
	ScannedURL  string
}

// Config ...
func Config() *Configure {
	return &Configure{
		LocalHost:   "http://localhost",
		PaidURL:     "paid_cb",
		RefundedURL: "refunded_cb",
		ScannedURL:  "scanned_cb",
	}
}

// LoadConfig ...
func LoadConfig(path string) *Configure {
	cfg := Config()
	t, e := toml.LoadFile(path)
	if e != nil {
		log.Println("filepath: " + path)
		log.Println(e.Error())
		return Config()
	}

	err := t.Unmarshal(&cfg)
	if err != nil {
		return Config()
	}

	return cfg

}
