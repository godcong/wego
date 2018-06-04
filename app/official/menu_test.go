package official_test

import (
	"testing"

	"github.com/godcong/wego/app/official"
	"github.com/godcong/wego/core/menu"
)

func TestNewMenu(t *testing.T) {
	menu := official.NewMenu()
	t.Log(menu)
	//testMenu_List(t, menu)
	//testMenu_AddButton(t, menu)
	//testMenu_Create(t, menu)
	testMenu_TryMatch(t, menu)
}

func testMenu_List(t *testing.T, m *official.Menu) {
	rlt := m.List()
	t.Log(rlt.ToString())
}

// func testMenu_AddButton(t *testing.T, m *official_account.Menu) {
// 	m.AddButton(menu.NewClickButton("click1", "fistkey"))
// }

func testMenu_Create(t *testing.T, m *official.Menu) {
	button := menu.NewBaseButton()
	// button.SetMatchRule&menu.MatchRule
	button.SetMatchRule(&menu.MatchRule{
		TagID:   "2",
		Sex:     "1",
		Country: "中国",
		//Province:           "广东",
		//City:               "广州",
		//ClientPlatformType: "2",
		//Language:           "zh_CN",
	})
	rlt := m.Create(button)
	t.Log(rlt.ToString())
}

func testMenu_TryMatch(t *testing.T, m *official.Menu) {
	rlt := m.TryMatch("ccdevil0910")
	t.Log(rlt.ToString())
}
