package wego_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/godcong/wego/core/menu"
	"github.com/godcong/wego/official_account"
)

func TestCoreButton(t *testing.T) {
	b1 := menu.NewClickButton("hello", "run")
	b2 := menu.NewViewButton("world", "y11e.com")
	//b1.SetSub("sub", []*core.Button{b2})
	b3 := menu.NewSubButton("sub", []*menu.Button{b1, b2})
	v, _ := json.Marshal(b3)
	log.Println(string(v))
}

func TestCoreMenu(t *testing.T) {
	menus := official_account.NewMenu()

	b1 := menu.NewClickButton("hello1", "run1")
	b2 := menu.NewClickButton("hello2", "run2")
	b3 := menu.NewClickButton("hello3", "run3")
	b4 := menu.NewClickButton("hello4", "run4")
	b5 := menu.NewClickButton("h2", "gogogo")
	//m2 := menu.NewClickButton("main2", "gmmmmm")
	//m3 := menu.NewClickButton("main3", "gmm33333m")
	//b1.SetSub("sub", []*core.Button{b2})
	b6 := menu.NewSubButton("sub1", []*menu.Button{b1, b2, b3})
	b7 := menu.NewSubButton("sub2", []*menu.Button{b4, b5})
	b8 := menu.NewSubButton("sub3", []*menu.Button{b1, b2, b4})
	menus.AddButton(b6)
	menus.AddButton(b7)
	menus.AddButton(b8)
	//menus.AddButton(m2)
	//menus.AddButton(m3)

	log.Println(menus.Create(nil).ToString())
	log.Println(menus.List().ToString())
	log.Println(menus.Current().ToString())
}
