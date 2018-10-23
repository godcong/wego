package mini

import (
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

/*Template Template */
type Template struct {
	Config
	*Program
	//client *core.Client
}

func newTemplate(program *Program) *Template {
	template := Template{
		Config:  defaultConfig,
		Program: program,
		//client:      program.GetClient(),
	}
	//template.client.SetDomain(core.NewDomain(""))
	return &template
}

/*NewTemplate NewTemplate */
func NewTemplate() *Template {
	return newTemplate(program)
}

/*List 获取小程序模板库标题列表
接口地址
https://api.weixin.qq.com/cgi-bin/wxopen/template/library/list?access_token=ACCESS_TOKEN
HTTP请求方式：
POST
POST参数说明：
参数	必填	说明
access_token	是	接口调用凭证
offset	是	offset和count用于分页，表示从offset开始，拉取count条记录，offset从0开始，count最大为20。
count	是	offset和count用于分页，表示从offset开始，拉取count条记录，offset从0开始，count最大为20。
返回码说明：
在调用模板消息接口后，会返回JSON数据包。
返回参数说明：
参数	说明
id	模板标题id（获取模板标题下的关键词库时需要）
title	模板标题内容
total_count	模板库标题总数
*/
func (t *Template) List(offset, count int) util.Map {
	return t.GetClient().HTTPPostJSON(
		t.client.Link(templateLibraryListURLSuffix), util.Map{"offset": offset, "count": count}, nil).ToMap()
}

/*Get 获取模板库某个模板标题下关键词库
接口地址
https://api.weixin.qq.com/cgi-bin/wxopen/template/library/get?access_token=ACCESS_TOKEN
HTTP请求方式：
POST
POST参数说明：
参数	必填	说明
access_token	是	接口调用凭证
id	是	模板标题id，可通过接口获取，也可登录小程序后台查看获取
返回码说明：
在调用模板消息接口后，会返回JSON数据包。
正常时的返回JSON数据包示例：
返回参数说明：
参数	说明
keyword_id	关键词id，添加模板时需要
name	关键词内容
example	关键词内容对应的示例
*/
func (t *Template) Get(id string) util.Map {
	return t.GetClient().HTTPPostJSON(
		t.client.Link(templateLibraryGetURLSuffix), util.Map{"id": id}, nil).ToMap()
}

/*Delete 删除帐号下的某个模板
接口地址
https://api.weixin.qq.com/cgi-bin/wxopen/template/del?access_token=ACCESS_TOKEN
HTTP请求方式：
POST
POST参数说明：
参数	必填	说明
access_token	是	接口调用凭证
template_id	是	要删除的模板id
返回码说明：
在调用模板消息接口后，会返回JSON数据包。
正常时的返回JSON数据包示例：
*/
func (t *Template) Delete(templateID string) util.Map {
	return t.GetClient().HTTPPostJSON(
		t.client.Link(templateDelURLSuffix), util.Map{"template_id": templateID}, nil).ToMap()
}

/*GetTemplates 获取帐号下已存在的模板列表
接口地址
https://api.weixin.qq.com/cgi-bin/wxopen/template/list?access_token=ACCESS_TOKEN
HTTP请求方式：
POST
POST参数说明：
参数	必填	说明
access_token	是	接口调用凭证
offset	是	offset和count用于分页，表示从offset开始，拉取count条记录，offset从0开始，count最大为20。最后一页的list长度可能小于请求的count
count	是	offset和count用于分页，表示从offset开始，拉取count条记录，offset从0开始，count最大为20。最后一页的list长度可能小于请求的count
返回码说明：
在调用模板消息接口后，会返回JSON数据包。
正常时的返回JSON数据包示例：
返回参数说明：
参数	说明
list	帐号下的模板列表
template_id	添加至帐号下的模板id，发送小程序模板消息时所需
title	模板标题
content	模板内容
example	模板内容示例
*/
func (t *Template) GetTemplates(offset, count int) util.Map {
	return t.GetClient().HTTPPostJSON(
		t.client.Link(templateListURLSuffix), util.Map{"offset": offset, "count": count}, nil).ToMap()
}

/*Add 组合模板并添加至帐号下的个人模板库
接口地址
https://api.weixin.qq.com/cgi-bin/wxopen/template/add?access_token=ACCESS_TOKEN
HTTP请求方式：
POST
POST参数说明：
参数	必填	说明
access_token	是	接口调用凭证
id	是	模板标题id，可通过接口获取，也可登录小程序后台查看获取
keyword_id_list	是	开发者自行组合好的模板关键词列表，关键词顺序可以自由搭配（例如[3,5,4]或[4,5,3]），最多支持10个关键词组合
返回码说明：
在调用模板消息接口后，会返回JSON数据包。
返回参数说明：
参数	说明
template_id	添加至帐号下的模板id，发送小程序模板消息时所需
*/
func (t *Template) Add(id string, keyword util.Map) util.Map {
	return t.GetClient().HTTPPostJSON(
		t.client.Link(templateAddURLSuffix), util.Map{"id": id, "keyword_id_list": keyword}, nil).ToMap()
}

/*Send 发送模板消息
接口地址：(ACCESS_TOKEN 需换成上文获取到的 access_token)
https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=ACCESS_TOKEN
HTTP请求方式：
POST
POST参数说明：
参数	必填	说明
touser	是	接收者（用户）的 openid
template_id	是	所需下发的模板消息的id
page	否	点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
form_id	是	表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
data	是	模板内容，不填则下发空模板
color	否	模板内容字体的颜色，不填默认黑色 【废弃】
emphasis_keyword	否	模板需要放大的关键词，不填则默认无放大
返回码说明：
在调用模板消息接口后，会返回JSON数据包。
错误时会返回错误码信息，说明如下：
返回码	说明
40037	template_id不正确
41028	form_id不正确，或者过期
41029	form_id已被使用
41030	page不正确
45009	接口调用超过限额（目前默认每个帐号日调用限额为100万）
*/
func (t *Template) Send(template *message.Template) core.Response {
	token := t.token.GetToken()
	resp := t.client.HTTPPostJSON(
		t.client.Link(templateSendURLSuffix),
		token.KeyMap(),
		template,
	)
	return resp
}
