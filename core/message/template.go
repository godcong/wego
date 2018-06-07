package message

/*ValueColor ValueColor */
type ValueColor struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

/*TemplateData TemplateData */
type TemplateData map[string]*ValueColor

/*TemplateMiniProgram TemplateMiniProgram */
type TemplateMiniProgram struct {
	AppID    string `json:"appid"`    //"appid":"xiaochengxuappid12345",
	PagePath string `json:"pagepath"` //"pagepath":"index?foo=bar"
}

//{
//	First            ValueColor `json:"first"`
//	KeyNote1         ValueColor `json:"keynote1"`
//	KeyNote2         ValueColor `json:"keynote2"`
//	KeyNote3         ValueColor `json:"keynote3"`
//	Remark           ValueColor `json:"remark"`
//	OrderProductName ValueColor `json:"orderProductName"`
//	//	"first": {
//	//	"value":"恭喜你购买成功！",
//	//	"color":"#173177"
//	//},
//	//	"keynote1":{
//	//	"value":"巧克力",
//	//	"color":"#173177"
//	//},
//	//	"keynote2": {
//	//	"value":"39.8元",
//	//	"color":"#173177"
//	//},
//	//	"keynote3": {
//	//	"value":"2014年9月22日",
//	//	"color":"#173177"
//	//},
//	//	"remark":{
//	//	"value":"欢迎再次购买！",
//	//	"color":"#173177"
//	//}
//}

/*Template Template */
type Template struct {
	ToUser          string              `json:"touser"`                     //"touser":"OPENID",
	TemplateID      string              `json:"template_id"`                // "template_id":"ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
	URL             string              `json:"url,omitempty"`              //"url":"http://weixin.qq.com/download",
	Data            TemplateData        `json:"data"`                       //"data":{...}
	MiniProgram     TemplateMiniProgram `json:"miniprogram,omitempty"`      //"miniprogram":
	Page            string              `json:"page,omitempty"`             //点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	FormID          string              `json:"form_id,omitempty"`          //小程序为必填 表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
	EmphasisKeyword string              `json:"emphasis_keyword,omitempty"` //模板需要放大的关键词，不填则默认无放大
}
