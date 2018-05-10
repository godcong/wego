package config_test

import (
	"testing"

	"github.com/godcong/wego/config"
)

func TestConfigLoader(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go config.GetRootConfig()

	}

}
