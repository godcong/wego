package app

import (
	"github.com/pelletier/go-toml"
	"log"
)

// Configure ...
type Configure struct {
	Local LocalProperty `toml:"local"`
}

// Config ...
func Config() *Configure {
	return &Configure{
		Local: LocalProperty{
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
