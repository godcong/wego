package app

import "testing"

// TestLoadConfig ...
func TestLoadConfig(t *testing.T) {
	config := LoadConfig("config.toml")
	t.Log(config)
}
