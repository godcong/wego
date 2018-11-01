package official

import (
	"net/url"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*QrCodeScene QrCodeScene*/
type QrCodeScene struct {
	SceneID  int    `json:"scene_id,omitempty"`
	SceneStr string `json:"scene_str,omitempty"`
}

/*QrCodeCard QrCodeCard*/
type QrCodeCard struct {
	CardID       string `json:"card_id,omitempty"`        // "card_id": "pFS7Fjg8kV1IdDz01r4SQwMkuCKc",
	Code         string `json:"code"`                     // "code": "198374613512",
	OpenID       string `json:"openid,omitempty"`         // "openid": "oFS7Fjl0WsZ9AMZqrI80nbIq8xrA",
	IsUniqueCode bool   `json:"is_unique_code,omitempty"` // "is_unique_code": false,
	OuterStr     string `json:"outer_str,omitempty"`      // "outer_str":"12b"
}

/*QrCodeCardList QrCodeCardList*/
type QrCodeCardList struct {
	CardID   string `json:"card_id,omitempty"`   // "card_id": "p1Pj9jgj3BcomSgtuW8B1wl-wo88",
	Code     string `json:"code"`                // "code": "198374613512",
	OuterStr string `json:"outer_str,omitempty"` // "outer_str":"12b"
}

/*QrCodeMultipleCard QrCodeMultipleCard*/
type QrCodeMultipleCard struct {
	CardList []QrCodeCardList `json:"card_list,omitempty"`
}

/*QrCodeActionInfo QrCodeActionInfo*/
type QrCodeActionInfo struct {
	Scene        *QrCodeScene        `json:"scene,omitempty"`
	Card         *QrCodeCard         `json:"card,omitempty"`
	MultipleCard *QrCodeMultipleCard `json:"multiple_card,omitempty"`
}

/*QrCodeAction QrCodeAction*/
type QrCodeAction struct {
	ExpireSeconds int              `json:"expire_seconds,omitempty"`
	ActionName    QrCodeActionName `json:"action_name"`
	ActionInfo    QrCodeActionInfo `json:"action_info"`
}

/*QrCodeActionName QrCodeActionName*/
type QrCodeActionName string

// QrMultipleCard ...
const (
	//QrMultipleCard QrMultipleCard
	QrMultipleCard QrCodeActionName = "QR_MULTIPLE_CARD"
	//QrCard QrCard
	QrCard QrCodeActionName = "QR_CARD"
	//QrScene QrScene
	QrScene QrCodeActionName = "QR_SCENE"
	//QrLimitStrScene QrLimitStrScene
	QrLimitStrScene QrCodeActionName = "QR_LIMIT_STR_SCENE"
)

/*QrCode QrCode*/
type QrCode struct {
	*Account
}

func newQrCode(acc *Account) *QrCode {
	return &QrCode{
		Account: acc,
	}
}

/*NewQrCode NewQrCode*/
func NewQrCode(config *core.Config) *QrCode {
	return newQrCode(NewAccount(config))
}

//Create 创建二维码ticket
// http请求方式: POST
// URL: https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN
// POST数据格式:json
// POST数据例子:{"expire_seconds": 604800, "action_name": "QR_SCENE", "action_info": {"scene": {"scene_id": 123}}}
//
// 或者也可以使用以下POST数据创建字符串形式的二维码参数:
// {"expire_seconds": 604800, "action_name": "QR_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}
// http请求方式: POST
// URL: https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN
// POST数据格式:json
// POST数据例子:{"action_name": "QR_LIMIT_SCENE", "action_info": {"scene": {"scene_id": 123}}}
//
// 或者也可以使用以下POST数据创建字符串形式的二维码参数:
// {"action_name": "QR_LIMIT_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}
// 成功:
// {"ticket":"gQFy7zwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFxLmNvbS9xLzAyOE1nSDFvTHdkeWkxeVNqTnhxMTcAAgR6E7FaAwQ8AAAA","expire_seconds":60,"url":"http:\/\/weixin.qq.com\/q\/028MgH1oLwdyi1ySjNxq17"}
func (q *QrCode) Create(action *QrCodeAction) core.Response {
	log.Debug("QrCode|Create", action)
	resp := q.client.PostJSON(
		Link(qrcodeCreateURLSuffix),
		q.accessToken.GetToken().KeyMap(),
		action,
	)
	return resp
}

//ShowQrCode 显示二维码
// HTTP GET请求（请使用https协议）https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=TICKET
// 提醒:TICKET记得进行UrlEncode
func (q *QrCode) ShowQrCode(ticket string) core.Response {
	log.Debug("QrCode|ShowQrCode", ticket)

	// base64.URLEncoding.EncodeToString([]byte(ticket))
	resp := q.client.Get(
		core.Link(showQrcodeURLSuffix, "mp"),
		util.Map{
			"ticket": url.QueryEscape(ticket),
		})
	return resp
}

/*String String*/
func (n QrCodeActionName) String() string {
	return string(n)
}
