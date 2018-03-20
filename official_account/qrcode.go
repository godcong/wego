package official_account

import (
	"net/url"

	"github.com/godcong/wego/core"
)

type QrCodeScene struct {
	SceneId  int    `json:"scene_id,omitempty"`
	SceneStr string `json:"scene_str,omitempty"`
}

type QrCodeActionInfo struct {
	Scene QrCodeScene `json:"scene"`
}

type QrCodeAction struct {
	ExpireSeconds int              `json:"expire_seconds"`
	ActionName    string           `json:"action_name"`
	ActionInfo    QrCodeActionInfo `json:"action_info"`
}

type QrCode struct {
	core.Config
	*OfficialAccount
}

func newQrCode(officialAccount *OfficialAccount) *QrCode {
	return &QrCode{
		Config:          defaultConfig,
		OfficialAccount: officialAccount,
	}
}

func NewQrCode() *QrCode {
	return newQrCode(account)
}

// http请求方式: POST
// URL: https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN
// POST数据格式：json
// POST数据例子：{"expire_seconds": 604800, "action_name": "QR_SCENE", "action_info": {"scene": {"scene_id": 123}}}
//
// 或者也可以使用以下POST数据创建字符串形式的二维码参数：
// {"expire_seconds": 604800, "action_name": "QR_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}
// http请求方式: POST
// URL: https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN
// POST数据格式：json
// POST数据例子：{"action_name": "QR_LIMIT_SCENE", "action_info": {"scene": {"scene_id": 123}}}
//
// 或者也可以使用以下POST数据创建字符串形式的二维码参数：
// {"action_name": "QR_LIMIT_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}
// 成功:
// {"ticket":"gQFy7zwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFxLmNvbS9xLzAyOE1nSDFvTHdkeWkxeVNqTnhxMTcAAgR6E7FaAwQ8AAAA","expire_seconds":60,"url":"http:\/\/weixin.qq.com\/q\/028MgH1oLwdyi1ySjNxq17"}
func (q *QrCode) Create(action *QrCodeAction) *core.Response {
	core.Debug("QrCode|Create", action)
	p := q.token.GetToken().KeyMap()
	resp := q.client.HttpPost(
		q.client.Link(QRCODE_CREATE_URL_SUFFIX),
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
			core.REQUEST_TYPE_JSON.String():  action,
		})
	return resp
}

// HTTP GET请求（请使用https协议）https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=TICKET
// 提醒：TICKET记得进行UrlEncode
func (q *QrCode) ShowQrCode(ticket string) *core.Response {
	core.Debug("QrCode|ShowQrCode", ticket)
	q.client.SetDomain(core.NewDomain("mp"))
	// base64.URLEncoding.EncodeToString([]byte(ticket))
	resp := q.client.HttpGet(
		q.client.Link(SHOWQRCODE_URL_SUFFIX),
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): core.Map{
				"ticket": url.QueryEscape(ticket),
			},
		})
	return resp
}
