package message

type ValueColor struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

type TemplateData map[string]*ValueColor

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

type Template struct {
	ToUser      string              `json:"touser"`                //"touser":"OPENID",
	TemplateId  string              `json:"template_id"`           // "template_id":"ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
	Url         string              `json:"url,omitempty"`         //"url":"http://weixin.qq.com/download",
	Data        TemplateData        `json:"data"`                  //"data":{...}
	MiniProgram TemplateMiniProgram `json:"miniprogram,omitempty"` //"miniprogram":
}
