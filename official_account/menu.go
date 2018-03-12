package official_account

import "github.com/godcong/wego/core"

type Menu struct {
	core.Config
	client  *core.Client
	token   *core.AccessToken
	buttons []*core.Button
}

func NewMenu(config core.Config, client *core.Client) *Menu {
	client.SetDomain(core.NewDomain("official_account"))
	client.SetDataType(core.DATA_TYPE_JSON)
	return &Menu{
		client:  client,
		token:   core.NewAccessToken(config, client),
		buttons: nil,
	}
}

func (m *Menu) SetButtons(b []*core.Button) {
	m.buttons = b
}

//个性化创建
//https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=ACCESS_TOKEN
//自定义菜单
//https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
func (m *Menu) Create(matchrule core.Map) {
	token := m.token.GetToken()
	if matchrule == nil {
		m.client.HttpPost(m.client.Link(MENU_CREATE_URL_SUFFIX), core.Map{
			core.REQUEST_TYPE_QUERY.String(): core.Map{"access_token": token.GetKey()},
			core.REQUEST_TYPE_JSON.String():  m.buttons,
		})
	}
}

func (m *Menu) List() {

}

func (m *Menu) Get() {

}

func (m *Menu) Current() {

}

//https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN
func (m *Menu) Delete() {

	//https: //api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN
}
