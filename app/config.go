package app

import (
	"github.com/pelletier/go-toml"
	"log"
)

// LocalHost ...
type LocalHost struct {
	Address     string `toml:"address"`
	PaidURL     string `toml:"paid_url"`
	RefundedURL string `toml:"refunded_url"`
	ScannedURL  string `toml:"scanned_url"`
}

// Configure ...
type Configure struct {
	LocalHost LocalHost `toml:"localhost"`
}

// Config ...
func Config() *Configure {
	return &Configure{
		LocalHost: LocalHost{
			Address:     "http://localhost",
			PaidURL:     "paid_cb",
			RefundedURL: "refunded_cb",
			ScannedURL:  "scanned_cb",
		},
	}
}

// LoadConfig ...
func LoadConfig(path string) *Configure {
	cfg := Config()
	t, err := toml.LoadFile(path)
	if err != nil {
		log.Println("filepath: " + path)
		log.Println(err.Error())
		return Config()
	}

	err = t.Unmarshal(cfg)
	if err != nil {
		log.Println("filepath: " + path)
		log.Println(err.Error())
		return Config()
	}

	return cfg

}
