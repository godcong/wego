package mini

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Template Template */
type Template struct {
	*Program
}

func newTemplate(program *Program) interface{} {
	return &Template{
		Program: program,
	}
}

/*NewTemplate NewTemplate */
func NewTemplate(config *core.Config) *Template {
	return newTemplate(NewMiniProgram(config)).(*Template)
}

/*List 获取小程序模板库标题列表
详情请见: https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1500465446_j4CgR&token=&lang=zh_CN
接口地址
https://api.weixin.qq.com/cgi-bin/wxopen/template/library/list?access_token=ACCESS_TOKEN
HTTP请求方式:
POST
POST参数说明:
参数	必填	说明
access_token	是	接口调用凭证
offset	是	offset和count用于分页，表示从offset开始，拉取count条记录，offset从0开始，count最大为20。
count	是	offset和count用于分页，表示从offset开始，拉取count条记录，offset从0开始，count最大为20。
*/
func (t *Template) List(offset, count int) core.Responder {
	token := t.accessToken.GetToken()
	return core.PostJSON(Link(templateLibraryList), token.KeyMap(), util.Map{"offset": offset, "count": count})
}

/*Get 获取模板库某个模板标题下关键词库
详情请见: https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1500465446_j4CgR&token=&lang=zh_CN
接口地址
https://api.weixin.qq.com/cgi-bin/wxopen/template/library/get?access_token=ACCESS_TOKEN
HTTP请求方式:
POST
POST参数说明:
参数	必填	说明
access_token	是	接口调用凭证
id	是	模板标题id，可通过接口获取，也可登录小程序后台查看获取
*/
func (t *Template) Get(id string) core.Responder {
	token := t.accessToken.GetToken()
	return core.PostJSON(Link(templateLibraryGet), token.KeyMap(), util.Map{"id": id})
}

/*Delete 删除帐号下的某个模板
详情请见: https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1500465446_j4CgR&token=&lang=zh_CN
接口地址
https://api.weixin.qq.com/cgi-bin/wxopen/template/del?access_token=ACCESS_TOKEN
HTTP请求方式:
POST
POST参数说明:
参数	必填	说明
access_token	是	接口调用凭证
template_id	是	要删除的模板id
*/
func (t *Template) Delete(templateID string) core.Responder {
	token := t.accessToken.GetToken()
	return core.PostJSON(Link(templateDel), token.KeyMap(), util.Map{"template_id": templateID})
}

/*GetTemplates 获取帐号下已存在的模板列表
详情请见: https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1500465446_j4CgR&token=&lang=zh_CN
接口地址
https://api.weixin.qq.com/cgi-bin/wxopen/template/list?access_token=ACCESS_TOKEN
HTTP请求方式:
POST
POST参数说明:
参数	必填	说明
access_token	是	接口调用凭证
offset	是	offset和count用于分页，表示从offset开始，拉取count条记录，offset从0开始，count最大为20。最后一页的list长度可能小于请求的count
count	是	offset和count用于分页，表示从offset开始，拉取count条记录，offset从0开始，count最大为20。最后一页的list长度可能小于请求的count
*/
func (t *Template) GetTemplates(offset, count int) core.Responder {
	token := t.accessToken.GetToken()
	return core.PostJSON(Link(templateList), token.KeyMap(), util.Map{"offset": offset, "count": count})

}

/*Add 组合模板并添加至帐号下的个人模板库
详情请见: https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1500465446_j4CgR&token=&lang=zh_CN
接口地址
https://api.weixin.qq.com/cgi-bin/wxopen/template/add?access_token=ACCESS_TOKEN
HTTP请求方式:
POST
POST参数说明:
参数	必填	说明
access_token	是	接口调用凭证
id	是	模板标题id，可通过接口获取，也可登录小程序后台查看获取
keyword_id_list	是	开发者自行组合好的模板关键词列表，关键词顺序可以自由搭配（例如[3,5,4]或[4,5,3]），最多支持10个关键词组合
*/
func (t *Template) Add(id string, keywordIdList []int) core.Responder {
	token := t.accessToken.GetToken()
	return core.PostJSON(Link(templateAdd), token.KeyMap(), util.Map{"id": id, "keyword_id_list": keywordIdList})
}

/*Send 发送模板消息
详情请见: https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1500465446_j4CgR&token=&lang=zh_CN
接口地址:(ACCESS_TOKEN 需换成上文获取到的 access_token)
https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=ACCESS_TOKEN
HTTP请求方式:
POST
POST参数说明:
参数	必填	说明
touser	是	接收者（用户）的 openid
template_id	是	所需下发的模板消息的id
page	否	点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
form_id	是	表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
data	是	模板内容，不填则下发空模板
color	否	模板内容字体的颜色，不填默认黑色 【废弃】
emphasis_keyword	否	模板需要放大的关键词，不填则默认无放大
*/
func (t *Template) Send(maps util.Map) core.Responder {
	token := t.accessToken.GetToken()
	return core.PostJSON(Link(templateSend), token.KeyMap(), maps)
}
