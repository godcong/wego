package wego

import (
	"github.com/godcong/wego/util"
	jsoniter "github.com/json-iterator/go"
)

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

// ToMap ...
func (t Template) ToMap() util.Map {
	var err error
	m := util.Map{}
	b, err := jsoniter.Marshal(t)
	if err != nil {
		return nil
	}
	err = jsoniter.Unmarshal(b, &m)
	if err != nil {
		return nil
	}
	return m
}
