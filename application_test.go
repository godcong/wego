package wego_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/official_account"
)

func TestCoreButton(t *testing.T) {
	b1 := core.NewClickButton("hello", "run")
	b2 := core.NewViewButton("world", "y11e.com")
	//b1.SetSub("sub", []*core.Button{b2})
	b3 := core.NewSubButton("sub", []*core.Button{b1, b2})
	v, _ := json.Marshal(b3)
	log.Println(string(v))
}

func TestCoreMenu(t *testing.T) {
	menu := official_account.NewMenu(core.GetConfig("official_account.default"), core.NewClient(core.GetConfig("official_account.default")))
	menu.Create(nil)
	b1 := core.NewClickButton("hello", "run")
	b2 := core.NewViewButton("world", "y11e.com")
	//b1.SetSub("sub", []*core.Button{b2})
	b3 := core.NewSubButton("sub", []*core.Button{b1, b2})
	menu.SetButtons([]*core.Button{b3})
	log.Println()
}
