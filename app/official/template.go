package official

import (
	//"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/message"

	"github.com/godcong/wego/util"
)

/*Template Template*/
type Template struct {
	*Account
}

func newTemplate(acc *Account) *Template {
	return &Template{
		Account: acc,
	}
}

/*NewTemplate NewTemplate */
func NewTemplate(config *core.Config) *Template {
	return newTemplate(NewOfficialAccount(config))
}

//SetIndustry 设置所属行业
// http请求方式: POST
// https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=ACCESS_TOKEN
//主行业	副行业	代码
//IT科技	互联网/电子商务	1
//IT科技	IT软件与服务	2
//IT科技	IT硬件与设备	3
//IT科技	电子技术	4
//IT科技	通信与运营商	5
//IT科技	网络游戏	6
//金融业	银行	7
//金融业	基金理财信托	8
//金融业	保险	9
//餐饮	餐饮	10
//酒店旅游	酒店	11
//酒店旅游	旅游	12
//运输与仓储	快递	13
//运输与仓储	物流	14
//运输与仓储	仓储	15
//教育	培训	16
//教育	院校	17
//政府与公共事业	学术科研	18
//政府与公共事业	交警	19
//政府与公共事业	博物馆	20
//政府与公共事业	公共事业非盈利机构	21
//医药护理	医药医疗	22
//医药护理	护理美容	23
//医药护理	保健与卫生	24
//交通工具	汽车相关	25
//交通工具	摩托车相关	26
//交通工具	火车相关	27
//交通工具	飞机相关	28
//房地产	建筑	29
//房地产	物业	30
//消费品	消费品	31
//商业服务	法律	32
//商业服务	会展	33
//商业服务	中介服务	34
//商业服务	认证	35
//商业服务	审计	36
//文体娱乐	传媒	37
//文体娱乐	体育	38
//文体娱乐	娱乐休闲	39
//印刷	印刷	40
//其它	其它	41
func (t *Template) SetIndustry(id1, id2 string) core.Response {
	resp := t.client.PostJSON(
		Link(templateAPISetIndustryURLSuffix),
		t.accessToken.GetToken().KeyMap(),
		util.Map{"industry_id1": id1, "industry_id2": id2},
	)
	return resp

}

//GetIndustry 获取设置的行业信息
// http请求方式:GET
// https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=ACCESS_TOKEN
func (t *Template) GetIndustry() core.Response {
	resp := t.client.Get(
		Link(templateGetIndustryURLSuffix),
		t.accessToken.GetToken().KeyMap())
	return resp
}

//Add 获得模板ID
// 获取模板:https://mp.weixin.qq.com/advanced/tmplmsg?action=list&t=tmplmsg/list&token=93895307&lang=zh_CN
// http请求方式: POST
// https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=ACCESS_TOKEN
func (t *Template) Add(shortID string) core.Response {
	resp := t.client.PostJSON(
		Link(templateAPIAddTemplateURLSuffix),
		t.accessToken.GetToken().KeyMap(),
		util.Map{"template_id_short": shortID})
	return resp
}

//Send 发送模板消息
//http请求方式: POST
//https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN
func (t *Template) Send(template *message.Template) core.Response {
	resp := t.client.PostJSON(
		Link(messageTemplateSendURLSuffix),
		t.accessToken.GetToken().KeyMap(),
		template,
	)
	return resp
}

//GetAllPrivate 获取模板列表
// url:https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=ACCESS_TOKEN
func (t *Template) GetAllPrivate() core.Response {
	resp := t.client.Get(
		Link(templateGetAllPrivateTemplateURLSuffix),
		t.accessToken.GetToken().KeyMap(),
	)
	return resp
}

//DelAllPrivate 删除模板
// url:https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=ACCESS_TOKEN
func (t *Template) DelAllPrivate(templateID string) core.Response {
	resp := t.client.PostJSON(
		Link(templateDelPrivateTemplateURLSuffix),
		t.accessToken.GetToken().KeyMap(),
		util.Map{"template_id": templateID},
	)
	return resp

}
