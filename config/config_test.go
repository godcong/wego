package config

import (
	"log"
	"testing"

	"github.com/godcong/wego/core"
)

func TestConfigLoader(t *testing.T) {

	log.Println(core.GetSystemConfig())

}
