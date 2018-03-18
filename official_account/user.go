package official_account

import (
	"encoding/json"

	"github.com/godcong/wego/core"
)

type User struct {
	config core.Config
	*OfficialAccount
}

func newUser(account *OfficialAccount) *User {
	return &User{
		config:          defaultConfig,
		OfficialAccount: account,
	}
}

func NewUser() *User {
	return newUser(account)
}

// http请求方式: POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=ACCESS_TOKEN
// POST数据格式：JSON
// POST数据例子：
// {
// "openid":"oDF3iY9ffA-hqb2vVvbr7qxf6A0Q",
// "remark":"pangzi"
// }
// 成功:
// {"errcode":0,"errmsg":"ok"}
// 失败:
// {"errcode":40013,"errmsg":"invalid appid"}
func (u *User) UpdateRemark(openid, remark string) *core.Response {
	core.Debug("User|UpdateRemark", openid, remark)
	p := u.token.GetToken().KeyMap()
	resp := u.client.HttpPostJson(
		u.client.Link(USER_INFO_UPDATEREMARK_URL_SUFFIX),
		core.Map{
			"openid": openid,
			"remark": remark,
		},
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// 接口调用请求说明
// http请求方式: GET
// https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
// 成功:
// {"subscribe":1,"openid":"o6_bmjrPTlm6_2sgVt7hMZOPfL2M","nickname":"Band","sex":1,"language":"zh_CN","city":"广州","province":"广东","country":"中国","headimgurl":"http://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0","subscribe_time":1382694957,"unionid":"o6_bmasdasdsad6_2sgVt7hMZOPfL""remark":"","groupid":0,"tagid_list":[128,2],"subscribe_scene":"ADD_SCENE_QR_CODE","qr_scene":98765,"qr_scene_str":""}
func (u *User) UserInfo(openid, lang string) *core.UserInfo {
	core.Debug("User|UpdateRemark", openid, lang)
	p := u.token.GetToken().KeyMap()
	p.Set("openid", openid)
	if lang != "" {
		p.Set("lang", lang)
	}

	resp := u.client.HttpGet(
		u.client.Link(USER_INFO_URL_SUFFIX),
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	var info core.UserInfo
	json.Unmarshal(resp.ToBytes(), &info)

	return &info
}

// http请求方式: POST
// https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=ACCESS_TOKEN
func (u *User) BatchGet(openids []string, lang string) *core.Response {

	core.Debug("User|BatchGet", openids, lang)
	p := u.token.GetToken().KeyMap()
	var list []struct {
		OpenId string
		Lang   string
	}
	// TODO:
	// for _, v := range openids {

	// list.Set("openid", v)
	// if lang != "" {
	// 	list.Set("lang", lang)
	// }
	// }
	resp := u.client.HttpPostJson(
		u.client.Link(USER_INFO_URL_SUFFIX),
		core.Map{
			"user_list": list,
		},
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}
