package wego_test

import (
	"log"
	"testing"

	"github.com/godcong/wego/core"
)

func TestConfig(t *testing.T) {
	t.Log(core.IsDebug())
	log.Debug("test")
	log.Error("test")
	core.Info("test")
	core.Warn("test")
	core.Fatal("test")
}
