package core_test

import (
	"github.com/godcong/wego/core"
	"testing"
)

func TestApp(t *testing.T) {
	app := core.NewApplication()

	var v core.Config

	app.Get("config", &v)

	t.Log(v)

}
