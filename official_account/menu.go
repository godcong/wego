package official_account

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/menu"
)

type Menu struct {
	core.Config
	client  *core.Client
	token   *core.AccessToken
	buttons core.Map
}

func newMenu(config core.Config, client *core.Client, token *core.AccessToken) *Menu {
	//client.SetDomain(core.NewDomain("official_account"))
	//client.SetDataType(core.DATA_TYPE_JSON)
	return &Menu{
		client:  client,
		token:   token,
		buttons: core.Map{},
	}
}

func NewMenu(config core.Config, client *core.Client) *Menu {
	client.SetDomain(core.NewDomain("official_account"))
	client.SetDataType(core.DATA_TYPE_JSON)
	return newMenu(client)
	}
}

func DefaultMenu() *Menu {

}

func (m *Menu) SetButtons(b []*menu.Button) *Menu {
	m.buttons["button"] = b
	return m
}

func (m *Menu) GetButtons() []*menu.Button {
	if v, b := m.buttons["button"]; b {
		if v0, b := v.([]*menu.Button); b {
			return v0
		}
	}
	return nil
}

func (m *Menu) AddButton(b *menu.Button) *Menu {
	if v := m.GetButtons(); v != nil {
		m.buttons["button"] = append(v, b)
	} else {
		m.buttons["button"] = []*menu.Button{b}
	}
	return m
}

//个性化创建
//https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=ACCESS_TOKEN
//自定义菜单
//https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
func (m *Menu) Create(matchRule core.Map) *core.Response {
	token := m.token.GetToken()
	if matchRule == nil {
		resp := m.client.HttpPost(m.client.Link(MENU_CREATE_URL_SUFFIX), core.Map{
			core.REQUEST_TYPE_QUERY.String(): core.Map{"access_token": token.GetKey()},
			core.REQUEST_TYPE_JSON.String():  m.buttons,
		})
		return resp
	}
	return nil
}

func (m *Menu) List() *core.Response {
	token := m.token.GetToken()
	resp := m.client.HttpGet(m.client.Link(MENU_GET_URL_SUFFIX), core.Map{
		core.REQUEST_TYPE_QUERY.String(): core.Map{"access_token": token.GetKey()},
	})
	return resp

}

func (m *Menu) Current() *core.Response {
	token := m.token.GetToken()
	resp := m.client.HttpGet(m.client.Link(GET_CURRENT_SELFMENU_INFO_URL_SUFFIX), core.Map{
		core.REQUEST_TYPE_QUERY.String(): core.Map{"access_token": token.GetKey()},
	})
	return resp
}

func (m *Menu) Delete() *core.Response {
	token := m.token.GetToken()
	resp := m.client.HttpGet(m.client.Link(MENU_DELETE_URL_SUFFIX), core.Map{
		core.REQUEST_TYPE_QUERY.String(): core.Map{"access_token": token.GetKey()},
	})
	return resp
}
