package official_account

import (
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

type Template struct {
	config  config.Config
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
func (t *Template) SetIndustry(id1, id2 string) *net.Response {
	token := t.token.GetToken()
	resp := t.client.HttpPostJson(
		t.client.Link(API_SET_INDUSTRY_URL_SUFFIX),
		util.Map{"industry_id1": id1, "industry_id2": id2},
		util.Map{net.REQUEST_TYPE_QUERY.String(): token.KeyMap()})
	return resp

}

//成功：
//{"primary_industry":{"first_class":"IT科技","second_class":"互联网|电子商务"},"secondary_industry":{"first_class":"IT科技","second_class":"IT软件与服务"}}
func (t *Template) GetIndustry() *net.Response {
	token := t.token.GetToken()
	resp := t.client.HttpGet(
		t.client.Link(GET_INDUSTRY_URL_SUFFIX),
		util.Map{net.REQUEST_TYPE_QUERY.String(): token.KeyMap()})
	return resp
}

// 获取模板：https://mp.weixin.qq.com/advanced/tmplmsg?action=list&t=tmplmsg/list&token=93895307&lang=zh_CN
// http请求方式: POST
// https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=ACCESS_TOKEN
// 成功：
// {"errcode":0,"errmsg":"ok","template_id":"tAsZKUQO0zNkrfvsTi2JexHJ9ZPudXuZSdcurGzE7Yo"}
func (t *Template) Add(shortId string) *net.Response {
	token := t.token.GetToken()
	resp := t.client.HttpPostJson(
		t.client.Link(API_ADD_TEMPLATE_URL_SUFFIX),
		util.Map{"template_id_short": shortId},
		util.Map{net.REQUEST_TYPE_QUERY.String(): token.KeyMap()})
	return resp
}

//失败:
//{"errcode":44002,"errmsg":"empty post data hint: [s0462vr27]"}
//{"errcode":40003,"errmsg":"invalid openid hint: [7nhAqA0429ge31]"}
//成功:
//{"errcode":0,"errmsg":"ok","msgid":191569096301903872}
func (t *Template) Send(template *message.Template) *net.Response {

	resp := t.client.HttpPostJson(
		t.client.Link(MESSAGE_TEMPLATE_SEND_URL_SUFFIX),

		t.token.GetToken().KeyMap(),
		template,
	)
	return resp
}

//url:https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=ACCESS_TOKEN
//成功:
//{"template_list":[{"template_id":"tAsZKUQO0zNkrfvsTi2JexHJ9ZPudXuZSdcurGzE7Yo","title":"订单支付成功","primary_industry":"IT科技","deputy_industry":"互联网|电子商务","content":"{{first.DATA}}\n\n支付金额：{{orderMoneySum.DATA}}\n商品信息：{{orderProductName.DATA}}\n{{Remark.DATA}}","example":"我们已收到您的货款，开始为您打包商品，请耐心等待: )\n支付金额：30.00元\n商品信息：我是商品名字\n\n如有问题请致电400-828-1878或直接在微信留言，小易将第一时间为您服务！"},{"template_id":"sBMv7KrI5O66W-lqMQXKMVAs6sdtk0IKa7P1IoqC_mg","title":"订单支付成功","primary_industry":"IT科技","deputy_industry":"互联网|电子商务","content":"{{first.DATA}}\n\n支付金额：{{orderMoneySum.DATA}}\n商品信息：{{orderProductName.DATA}}\n{{Remark.DATA}}","example":"我们已收到您的货款，开始为您打包商品，请耐心等待: )\n支付金额：30.00元\n商品信息：我是商品名字\n\n如有问题请致电400-828-1878或直接在微信留言，小易将第一时间为您服务！"},{"template_id":"mO3VehTDPKVl-bJ1-58ZmfeFTuKwWP9yg6_tzkt1Ab0","title":"订阅模板消息","primary_industry":"","deputy_industry":"","content":"{{content.DATA}}","example":""},{"template_id":"awT3aSQJdtWqn7VRUNLzdEboGb1fONot3Z7SrsBtsjg","title":"订单支付成功","primary_industry":"IT科技","deputy_industry":"互联网|电子商务","content":"{{first.DATA}}\n\n支付金额：{{orderMoneySum.DATA}}\n商品信息：{{orderProductName.DATA}}\n{{Remark.DATA}}","example":"我们已收到您的货款，开始为您打包商品，请耐心等待: )\n支付金额：30.00元\n商品信息：我是商品名字\n\n如有问题请致电400-828-1878或直接在微信留言，小易将第一时间为您服务！"},{"template_id":"vc2ekfQmEP9qE9eBW9gGWaUrsLvztC9XOeB-cftLroo","title":"订单支付成功","primary_industry":"IT科技","deputy_industry":"互联网|电子商务","content":"{{first.DATA}}\n\n支付金额：{{orderMoneySum.DATA}}\n商品信息：{{orderProductName.DATA}}\n{{Remark.DATA}}","example":"我们已收到您的货款，开始为您打包商品，请耐心等待: )\n支付金额：30.00元\n商品信息：我是商品名字\n\n如有问题请致电400-828-1878或直接在微信留言，小易将第一时间为您服务！"}]}
func (t *Template) GetAllPrivate() *net.Response {
	token := t.token.GetToken()
	resp := t.client.HttpGet(
		t.client.Link(GET_ALL_PRIVATE_TEMPLATE_URL_SUFFIX),
		util.Map{net.REQUEST_TYPE_QUERY.String(): token.KeyMap()},
	)
	return resp
}

//url:https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=ACCESS_TOKEN
//成功:
//{"errcode":0,"errmsg":"ok"}
func (t *Template) DelAllPrivate(templateID string) *net.Response {
	token := t.token.GetToken()
	resp := t.client.HttpPostJson(
		t.client.Link(DEL_PRIVATE_TEMPLATE_URL_SUFFIX),
		util.Map{"template_id": templateID},
		util.Map{net.REQUEST_TYPE_QUERY.String(): token.KeyMap()},
	)
	return resp

}
