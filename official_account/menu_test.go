package official_account_test

import (
	"testing"

	"github.com/godcong/wego/core/menu"
	"github.com/godcong/wego/official_account"
)

func TestNewMenu(t *testing.T) {
	menu := official_account.NewMenu()
	t.Log(menu)
	testMenu_List(t, menu)
	testMenu_AddButton(t, menu)
	testMenu_Create(t, menu)
}

func testMenu_List(t *testing.T, m *official_account.Menu) {
	rlt := m.List()
	t.Log(rlt.ToString())
}

func testMenu_AddButton(t *testing.T, m *official_account.Menu) {
	m.AddButton(menu.NewClickButton("click1", "fistkey"))
}

func testMenu_Create(t *testing.T, m *official_account.Menu) {
	m.SetMatchRule(&menu.MatchRule{
		TagId:   "2",
		Sex:     "1",
		Country: "中国",
		//Province:           "广东",
		//City:               "广州",
		//ClientPlatformType: "2",
		//Language:           "zh_CN",
	})
	rlt := m.Create()
	t.Log(rlt.ToString())
}
