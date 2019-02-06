package _bak_test

import (
	"testing"

	"github.com/godcong/wego/log"
)

// TestConfig ...
func TestConfig(t *testing.T) {
	t.Log(log.IsDebug())
	log.Debug("test")
	log.Error("test")
	log.Info("test")
	log.Warn("test")
	log.Fatal("test")
}
