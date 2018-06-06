package wego_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/godcong/wego/core/menu"
)

func TestCoreButton(t *testing.T) {
	b1 := menu.NewClickButton("hello", "run")
	b2 := menu.NewViewButton("world", "y11e.com")
	//b1.SetSub("sub", []*core.Button{b2})
	b3 := menu.NewSubButton("sub", []*menu.Button{b1, b2})
	v, _ := json.Marshal(b3)
	log.Println(string(v))
}
