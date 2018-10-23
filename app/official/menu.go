package official

import (
	//"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/menu"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

/*Menu Menu*/
type Menu struct {
	//config  Config
	account *Account
	client  *core.Client
	token   *core.AccessToken
}

func newMenu(account *Account) *Menu {
	return &Menu{
		//config:  defaultConfig,
		account: account,
		client:  account.client,
		token:   account.token,
		//buttons: make(util.Map),
	}
}

/*NewMenu NewMenu*/
func NewMenu() *Menu {
	return newMenu(account)
}

//func (m *Menu) SetMatchRule(rule *menu.MatchRule) *Menu {
//	m.buttons["matchrule"] = rule
//	return m
//}
//
//func (m *Menu) SetMenuId(id int) *Menu {
//	m.menuid = id
//	return m
//}

//Create 创建菜单
//个性化创建
//https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=ACCESS_TOKEN
//成功:
//{"errcode":0,"errmsg":"ok"}
//自定义菜单
//https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
//成功:
// {"menuid":429680901}]
func (m *Menu) Create(buttons *menu.Button) core.Response {
	token := m.token.GetToken().KeyMap()
	if buttons.GetMatchRule() == nil {
		resp := m.client.HTTPPostJSON(
			m.client.Link(menuCreateURLSuffix),
			token,
			buttons)
		return resp
	}
	resp := m.client.HTTPPostJSON(
		m.client.Link(menuAddConditionalURLSuffix),
		token,
		buttons)
	return resp
}

/*List 自定义菜单查询接口
请求说明
http请求方式：GET
https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN
返回说明（无个性化菜单时）
参考URL:https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141014
*/
func (m *Menu) List() core.Response {
	resp := m.client.HTTPGet(m.client.Link(menuGetURLSuffix),
		m.token.GetToken().KeyMap(),
	)
	return resp
}

/*Current 获取自定义菜单配置接口
接口调用请求说明
http请求方式: GET（请使用https协议）https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=ACCESS_TOKEN
返回结果说明
参考URL:https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1434698695
*/
func (m *Menu) Current() core.Response {
	resp := m.client.HTTPGet(m.client.Link(getCurrentSelfMenuInfoURLSuffix),
		m.token.GetToken().KeyMap())
	return resp
}

/*TryMatch 测试个性化菜单匹配结果
http请求方式：POST（请使用https协议）
https://api.weixin.qq.com/cgi-bin/menu/trymatch?access_token=ACCESS_TOKEN
请求示例
{"user_id":"weixin"}
user_id可以是粉丝的OpenID，也可以是粉丝的微信号。
返回结果 该接口将返回菜单配置，示例如下：
{
    "button": [
        {
            "type": "view",
            "name": "tx",
            "url": "http://www.qq.com/",
            "sub_button": [ ]
        },
        {
            "type": "view",
            "name": "tx",
            "url": "http://www.qq.com/",
            "sub_button": [ ]
        },
        {
            "type": "view",
            "name": "tx",
            "url": "http://www.qq.com/",
            "sub_button": [ ]
        }
    ]
}
*/
func (m *Menu) TryMatch(userID string) core.Response {
	resp := m.client.HTTPPostJSON(m.client.Link(menuTryMatchURLSuffix),
		m.token.GetToken().KeyMap(),
		util.Map{"user_id": userID})
	return resp
}

/*Delete 自定义菜单删除接口
请求说明
http请求方式：GET
https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN
返回说明
对应创建接口，正确的Json返回结果:
{"errcode":0,"errmsg":"ok"}
*/
func (m *Menu) Delete(menuid int) core.Response {
	token := m.token.GetToken().KeyMap()
	if menuid == 0 {
		resp := m.client.HTTPGet(m.client.Link(menuDeleteURLSuffix),
			token)
		return resp
	}

	resp := m.client.HTTPPostJSON(m.client.Link(menuDeleteConditionalURLSuffix),
		util.Map{"menuid": menuid},
		token)
	return resp
}
