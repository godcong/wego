package wego_test

import (
	"testing"

	"github.com/godcong/wego/core"
)

func TestConfig(t *testing.T) {
	t.Log(core.IsDebug())
	core.Debug("test")
	core.Error("test")
	core.Info("test")
	core.Warn("test")
	core.Fatal("test")
}
