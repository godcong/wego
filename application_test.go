package wego_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/godcong/wego/core"
)

func TestCoreButton(t *testing.T) {
	b1 := core.NewClickButton("hello", "run")
	b2 := core.NewViewButton("world", "y11e.com")
	b1.SetSub("sub", []*core.Button{b2})
	b3 := core.NewSubButton("sub", []*core.Button{b1})
	v, _ := json.Marshal(b3)
	log.Println(string(v))
}
