package official_account

import "github.com/godcong/wego/core"

type Template struct {
	config  core.Config
	account *OfficialAccount
	client  *core.Client
	token   *core.AccessToken
}

func newTemplate(account *OfficialAccount) *Template {
	return &Template{
		config:  defaultConfig,
		account: account,
		client:  account.client,
		token:   account.token,
	}
}

func NewTemplate() *Template {
	return newTemplate(account)
}

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
//成功：
//失败：
//{"errcode":43100,"errmsg":"change template too frequently hint: [ZJ3zDA0168vr23]"}
func (t *Template) SetIndustry(id1, id2 string) *core.Response {
	token := t.token.GetToken()
	resp := t.client.HttpPostJson(
		t.client.Link(API_SET_INDUSTRY_URL_SUFFIX),
		core.Map{"industry_id1": id1, "industry_id2": id2},
		core.Map{core.REQUEST_TYPE_QUERY.String(): token.KeyMap()})
	return resp

}

//成功：
//{"primary_industry":{"first_class":"IT科技","second_class":"互联网|电子商务"},"secondary_industry":{"first_class":"IT科技","second_class":"IT软件与服务"}}
func (t *Template) GetIndustry() *core.Response {
	token := t.token.GetToken()
	resp := t.client.HttpGet(
		t.client.Link(GET_INDUSTRY_URL_SUFFIX),
		core.Map{core.REQUEST_TYPE_QUERY.String(): token.KeyMap()})
	return resp
}

//成功：
//{"errcode":0,"errmsg":"ok","template_id":"tAsZKUQO0zNkrfvsTi2JexHJ9ZPudXuZSdcurGzE7Yo"}
func (t *Template) AddTemplate(shortId string) *core.Response {
	token := t.token.GetToken()
	resp := t.client.HttpPostJson(
		t.client.Link(API_ADD_TEMPLATE_URL_SUFFIX),
		core.Map{"template_id_short": shortId},
		core.Map{core.REQUEST_TYPE_QUERY.String(): token.KeyMap()})
	return resp
}
