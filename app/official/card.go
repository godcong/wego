package official

import (
	"encoding/json"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"strings"
	"time"
)

// Card ...
type Card struct {
	*Account
	jssdk *JSSDK
}

func newCard(officialAccount *Account) *Card {
	return &Card{
		Account: officialAccount,
		jssdk:   NewJSSDK(officialAccount.Config),
	}
}

// NewCard ...
func NewCard(config *core.Config) *Card {
	return newCard(NewOfficialAccount(config))
}

//CreateLandingPage 创建货架接口
//	HTTP请求方式: POST
//	URL:https://api.weixin.qq.com/card/landingpage/create?access_token=$TOKEN
// input a CardLandingPage point or a util.map
func (c *Card) CreateLandingPage(able util.MapAble) core.Responder {
	resp := core.PostJSON(
		Link(cardLandingPageCreate),
		c.accessToken.GetToken().KeyMap(),
		able.ToMap(),
	)
	return resp
}

//Deposit 导入code接口
//	HTTP请求方式: POST
//	URL:http://api.weixin.qq.com/card/code/deposit?access_token=ACCESS_TOKEN
func (c *Card) Deposit(cardID string, code []string) core.Responder {
	resp := core.PostJSON(
		Link(cardCodeDeposit),
		c.accessToken.GetToken().KeyMap(),
		util.Map{
			"card_id": cardID,
			"code":    code,
		},
	)
	return resp
}

//GetDepositCount 查询导入code数目
//
//  HTTP请求方式: POST
//  URL:http://api.weixin.qq.com/card/code/getdepositcount?access_token=ACCESS_TOKEN
func (c *Card) GetDepositCount(cardID string) core.Responder {
	resp := core.PostJSON(
		Link(cardCodeGetDepositCount),
		c.accessToken.GetToken().KeyMap(),
		util.Map{
			"card_id": cardID,
		},
	)
	return resp
}

//CheckCode 核查code接口
//	HTTP请求方式: POST
//	HTTP调用:http://api.weixin.qq.com/card/code/checkcode?access_token=ACCESS_TOKEN
func (c *Card) CheckCode(cardID string, code []string) core.Responder {
	resp := core.PostJSON(
		Link(cardCodeCheckCode),
		c.accessToken.GetToken().KeyMap(),
		util.Map{
			"card_id": cardID,
			"code":    code,
		},
	)
	return resp
}

//GetCode 查询Code接口
//	HTTP请求方式: POST
//	HTTP调用:https://api.weixin.qq.com/card/code/get?access_token=TOKEN
//	参数说明:
//	参数名	必填	类型	示例值	描述
//	code	是	string(20)	110201201245	单张卡券的唯一标准。
//	card_id	否	string(32)	pFS7Fjg8kV1I dDz01r4SQwMkuCKc	卡券ID代表一类卡券。自定义code卡券必填。
//	check_consume	否	bool	true	是否校验code核销状态，填入true和false时的code异常状态返回数据不同。
func (c *Card) GetCode(maps util.Map) core.Responder {
	resp := core.PostJSON(
		Link(cardCodeGet),
		c.accessToken.GetToken().KeyMap(),
		maps,
	)
	return resp
}

//GetHTML 图文消息群发卡券
//	HTTP请求方式: POST
//	URL:https://api.weixin.qq.com/card/mpnews/gethtml?access_token=TOKEN
func (c *Card) GetHTML(cardID string) core.Responder {
	resp := core.PostJSON(
		Link(cardMPNewsGetHTML),
		c.accessToken.GetToken().KeyMap(),
		util.Map{
			"card_id": cardID,
		},
	)
	return resp
}

//SetTestWhiteListByID 设置测试白名单(by openid)
func (c *Card) SetTestWhiteListByID(list []string) core.Responder {
	return c.SetTestWhiteList("openid", list)
}

//SetTestWhiteListByName 设置测试白名单(by username)
func (c *Card) SetTestWhiteListByName(list []string) core.Responder {
	return c.SetTestWhiteList("username", list)
}

//SetTestWhiteList 设置测试白名单
//	HTTP请求方式: POST
//	URL:https://api.weixin.qq.com/card/testwhitelist/set?access_token=TOKEN
func (c *Card) SetTestWhiteList(typ string, list []string) core.Responder {
	resp := core.PostJSON(
		Link(cardTestWhiteListSet),
		c.accessToken.GetToken().KeyMap(),
		util.Map{
			typ: list,
		},
	)
	return resp
}

//CreateQrCode 创建二维码
//	HTTP请求方式: POST
//	URL:https://api.weixin.qq.com/card/qrcode/create?access_token=TOKEN
func (c *Card) CreateQrCode(able util.MapAble) core.Responder {
	resp := core.PostJSON(
		Link(cardQrcodeCreate),
		c.accessToken.GetToken().KeyMap(),
		able.ToMap(),
	)
	return resp
}

//Create 创建卡券
//	HTTP请求方式: POST
//	URL: https://api.weixin.qq.com/card/create?access_token=ACCESS_TOKEN
//	func (c *Card) Create(card *OneCard) core.Responder {
func (c *Card) Create(able util.MapAble) core.Responder {
	key := c.accessToken.GetToken().KeyMap()
	//_, d := maps.Get()
	resp := core.PostJSON(
		Link(cardCreate),
		key,
		util.Map{"card": able.ToMap()})

	return resp
}

//Get 查看卡券详情
//	开发者可以调用该接口查询某个card_id的创建信息、审核状态以及库存数量。
//	接口调用请求说明
//	HTTP请求方式: POSTURL:https://api.weixin.qq.com/card/get?access_token=TOKEN
func (c *Card) Get(cardID string) core.Responder {
	token := c.accessToken.GetToken()
	return core.PostJSON("card/get", token.KeyMap(), util.Map{"card_id": cardID})
}

//GetApplyProtocol 卡券开放类目查询接口
//	HTTP请求方式: GET
//	URL:https://api.weixin.qq.com/card/getapplyprotocol?access_token=TOKEN
func (c *Card) GetApplyProtocol() core.Responder {
	token := c.accessToken.GetToken()
	return core.Get(Link(cardGetApplyProtocol), token.KeyMap())
}

//GetColors 卡券开放类目查询接口
//	HTTP请求方式: GET
//	URL:https://api.weixin.qq.com/card/getcolors?access_token=TOKEN
func (c *Card) GetColors() core.Responder {
	token := c.accessToken.GetToken()
	return core.Get(Link(cardGetColors), token.KeyMap())
}

//Checkin 更新飞机票信息接口
//	接口调用请求说明
//	http请求方式: POST
//	URL:https://api.weixin.qq.com/card/boardingpass/checkin?access_token=TOKEN
func (c *Card) Checkin(p util.Map) core.Responder {
	token := c.accessToken.GetToken()
	return core.PostJSON(Link(cardBoardingpassCheckin), token.KeyMap(), p)
}

//Categories 卡券开放类目查询接口
//	接口说明
//	通过调用该接口查询卡券开放的类目ID，类目会随业务发展变更，请每次用接口去查询获取实时卡券类目。
//	注意：
//	1.本接口查询的返回值还有卡券资质ID,此处的卡券资质为：已微信认证的公众号通过微信公众平台申请卡券功能时，所需的资质。
//	2.对于第三方强授权模式，子商户无论选择什么类目，均提交营业执照即可，所以不用考虑此处返回的资质字段，返回值仅参考类目ID即可。
//	接口详情
//	接口调用请求说明
//	https请求方式: GET https://api.weixin.qq.com/card/getapplyprotocol?access_token=TOKEN
func (c *Card) Categories() core.Responder {
	token := c.accessToken.GetToken()
	return core.Get(Link(cardGetapplyprotocol), token.KeyMap())

}

//BatchGet 批量查询卡券列表
//	接口调用请求说明
//	HTTP请求方式: POST URL:https://api.weixin.qq.com/card/batchget?access_token=TOKEN
func (c *Card) BatchGet(offset, count int, statusList []CardStatus) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"offset":      offset,
		"count":       count,
		"status_list": statusList,
	}
	return core.PostJSON(Link(cardBatchget), token.KeyMap(), maps)
}

//Update 更改卡券信息接口
//	接口说明
//	支持更新所有卡券类型的部分通用字段及特殊卡券（会员卡、飞机票、电影票、会议门票）中特定字段的信息。
//	接口调用请求说明
//	HTTP请求方式: POST URL:https://api.weixin.qq.com/card/update?access_token=TOKEN
func (c *Card) Update(cardID string, p util.Map) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"card_id": cardID,
	}
	maps.Join(p)
	return core.PostJSON(Link(cardUpdate), token.KeyMap(), maps)
}

//Delete 删除卡券接口
//删除卡券接口允许商户删除任意一类卡券。删除卡券后，该卡券对应已生成的领取用二维码、添加到卡包JS API均会失效。 注意：如用户在商家删除卡券前已领取一张或多张该卡券依旧有效。即删除卡券不能删除已被用户领取，保存在微信客户端中的卡券。
//接口调用请求说明
//HTTP请求方式: POST URL:https://api.weixin.qq.com/card/delete?access_token=TOKEN
func (c *Card) Delete(cardID string) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"card_id": cardID,
	}
	return core.PostJSON(Link(cardDelete), token.KeyMap(), maps)
}

// GetUserCards ...
func (c *Card) GetUserCards(openID, cardID string) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"openid":  openID,
		"card_id": cardID,
	}
	return core.PostJSON(Link(cardUserGetcardlist), token.KeyMap(), maps)
}

// SetPayCell ...
func (c *Card) SetPayCell(cardID string, isOpen bool) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"is_open": isOpen,
		"card_id": cardID,
	}
	return core.PostJSON(Link(cardPaycellSet), token.KeyMap(), maps)
}

// ModifyStock ...
func (c *Card) ModifyStock(cardID string, option util.Map) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"card_id": cardID,
	}
	maps.Join(option)
	return core.PostJSON(Link(cardModifystock), token.KeyMap(), maps)
}

// PayActivate ...
func (c *Card) PayActivate() core.Responder {
	token := c.accessToken.GetToken()
	return core.Get(Link(cardPayActivate), token.KeyMap())
}

// PayGetPrice ...
func (c *Card) PayGetPrice(cardID string, quantity int) core.Responder {
	token := c.accessToken.GetToken()
	return core.PostJSON(Link(cardPayGetpayprice), token.KeyMap(), util.Map{
		"card_id":  cardID,
		"quantity": quantity,
	})
}

// PayGetCoinsInfo ...
func (c *Card) PayGetCoinsInfo() core.Responder {
	token := c.accessToken.GetToken()
	return core.Get(Link(cardPayGetcoinsinfo), token.KeyMap())
}

// PayRecharge ...
func (c *Card) PayRecharge(count int) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardPayGetpayprice), token, util.Map{
		"coin_count": count,
	})
}

// PayOrder ...
func (c *Card) PayOrder(orderID string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardPayGetorder), token, util.Map{
		"order_id": orderID,
	})

}

// PayGetOrderList ...
func (c *Card) PayGetOrderList(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardPayGetorderlist), token, util.MapNilMake(p))
}

// PayConfirm ...
func (c *Card) PayConfirm(cardID, orderID string, quantity int) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardPayConfirm), token, util.Map{
		"card_id":  cardID,
		"order_id": orderID,
		"quantity": quantity,
	})
}

// GeneralActivate ...
func (c *Card) GeneralActivate(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardGeneralcardActivate), token, util.MapNilMake(p))
}

// GeneralDeactivate ...
func (c *Card) GeneralDeactivate(cardID, code string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardGeneralcardUnactivate), token, util.Map{
		"card_id": cardID,
		"code":    code,
	})
}

//GeneralUpdateUser 更新用户礼品卡信息接口
//	接口说明
//	当礼品卡被使用后，开发者可以通过该接口变更某个礼品卡的余额信息。
//	接口调用请求说明
//	协议	HTTPS
//	http请求方式	POST
//	请求Url	https://api.weixin.qq.com/card/generalcard/updateuser?access_token=TOKEN
//	POST数据格式	JSON
func (c *Card) GeneralUpdateUser(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardGeneralcardUpdateuser), token, util.MapNilMake(p))
}

// MeetingUpdateUser ...
func (c *Card) MeetingUpdateUser(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardMeetingticketUpdateuser), token, util.MapNilMake(p))
}

//GiftAdd 创建-礼品卡货架接口
//	接口说明
//	开发者可以通过该接口创建一个礼品卡货架并且用于公众号、门店的礼品卡售卖。
//	接口调用请求说明
//	协议	HTTPS
//	http请求方式	POST
//	请求Url	https://api.weixin.qq.com/card/giftcard/page/add?access_token=ACCESS_TOKEN
//	POST数据格式	JSON
func (c *Card) GiftAdd(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardGiftcardPageAdd), token, util.MapNilMake(p))
}

//GiftGet 查询-礼品卡货架信息接口
//	接口说明
//	开发者可以查询某个礼品卡货架信息。
//	接口调用请求说明
//	协议	HTTPS
//	http请求方式	POST
//	请求Url	https://api.weixin.qq.com/card/giftcard/page/get?access_token=ACCESS_TOKEN
//	POST数据格式	JSON
func (c *Card) GiftGet(pageID string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardGiftcardPageGet), token, util.Map{
		"page_id": pageID,
	})
}

//GiftUpdate 修改-礼品卡货架信息接口
//	接口说明
//	开发者可以通过该接口更新礼品卡货架信息。
//	接口调用请求说明
//	协议	HTTPS
//	http请求方式	POST
//	请求Url	https://api.weixin.qq.com/card/giftcard/page/update?access_token=ACCESS_TOKEN
//	POST数据格式	JSON
func (c *Card) GiftUpdate(pageID string, p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	p = util.MapNilMake(p)
	p.Set("page_id", pageID)
	return core.PostJSON(Link(cardGiftcardPageUpdate), token, p)
}

//GiftBatchGet 查询-礼品卡货架列表接口
//接口说明
//开发者可以通过该接口查询当前商户下所有的礼品卡货架id。
//接口调用请求说明
//协议	HTTPS
//http请求方式	POST
//请求Url	https://api.weixin.qq.com/card/giftcard/page/batchget?access_token=ACCESS_TOKEN
func (c *Card) GiftBatchGet() core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardGiftcardPageBatchget), token, util.Map{})
}

//GiftSet 下架-礼品卡货架接口
//接口说明
//开发者可以通过该接口查询当前商户下所有的礼品卡货架id。
//接口调用请求说明
//协议	HTTPS
//http请求方式	POST
//请求Url	https://api.weixin.qq.com/card/giftcard/maintain/set?access_token=ACCESS_TOKEN
func (c *Card) giftSet(pageID string, all, maintain bool) core.Responder {
	token := c.accessToken.KeyMap()
	maps := util.Map{
		"maintain": true,
	}
	if pageID == "" {
		maps.Set("all", all)
	} else {
		maps.Set("page_id", pageID)
	}
	return core.PostJSON(Link(cardGiftcardMaintainSet), token, maps)
}

// GiftSetByID ...
func (c *Card) GiftSetByID(pageID string, maintain bool) core.Responder {
	return c.giftSet(pageID, false, maintain)
}

// GiftSetAll ...
func (c *Card) GiftSetAll(all, maintain bool) core.Responder {
	return c.giftSet("", all, maintain)
}

//GiftRefund 退款接口
//	接口说明
//	开发者可以通过该接口对某一笔订单操作退款，注意该接口比较隐私，请开发者提高操作该功能的权限等级。
//	接口调用请求说明
//	协议	HTTPS
//	http请求方式	POST
//	请求Url	https://api.weixin.qq.com/card/giftcard/order/refund?access_token=ACCESS_TOKEN
//	POST数据格式	JSON
//	请求数据说明：
//	参数	说明	是否必填
//	order_id	须退款的订单id	是
func (c *Card) GiftRefund(orderID string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardGiftcardOrderRefund), token, util.Map{
		"order_id": orderID,
	})
}

//InvoiceSetPayMch 设置支付后开票信息
//	接口说明
//	商户可以通过该接口设置某个商户号发生收款后在支付消息上出现开票授权按钮。
//	请求url：
//	https://api.weixin.qq.com/card/invoice/setbizattr?action=set_pay_mch&access_token={access_token}
//	请求方法：POST
//	参数	类型	是否必填	描述
//	paymch_info	Object	是	授权页字段
//	paymch_info包含以下字段：
//	参数	类型	是否必填	描述
//	mchid	String	是	微信支付商户号
//	s_pappid	String	是	开票平台id，需要找开票平台提供
func (c *Card) InvoiceSetPayMch(mchID string, appID string) core.Responder {
	token := c.accessToken.KeyMap()
	token.Set("action", "set_pay_mch")
	maps := util.Map{}
	maps.Set("paymch_info.mchid", mchID)
	maps.Set("paymch_info.s_pappid", appID)
	return core.PostJSON(Link(cardInvoiceSetbizattr), token, maps)
}

//InvoiceGetPayMch 查询支付后开票信息接口
//	请求url：
//	https://api.weixin.qq.com/card/invoice/setbizattr?action=get_pay_mch&access_token={access_token}
//	请求方法：POST
func (c *Card) InvoiceGetPayMch() core.Responder {
	token := c.accessToken.KeyMap()
	token.Set("action", "get_pay_mch")
	return core.PostJSON(Link(cardInvoiceSetbizattr), token, util.Map{})
}

//InvoiceSetAuthField 设置授权页字段信息接口
//使用说明
//商户可以通过该接口设置用户授权时应该填写的内容
//请求说明url：
//https://api.weixin.qq.com/card/invoice/setbizattr?action=set_auth_field&access_token={access_token}
//请求方法：POST
//请求参数
//数据格式：JSON
//参数	类型	是否必填	描述
//auth_field	Object	是	授权页字段
//auth_field包含以下字段：
//参数	类型	是否必填	描述
//user_field	Object	是	授权页个人发票字段
//biz_field	Object	是	授权页单位发票字段
//user_field包含以下字段：
//参数	类型	是否必填	描述
//show_title	Int	否	是否填写抬头，0为否，1为是
//show_phone	Int	否	是否填写电话号码，0为否，1为是
//show_email	Int	否	是否填写邮箱，0为否，1为是
//custom_field	Object	否	自定义字段
//biz_field包含以下字段：
//参数	类型	是否必填	描述
//show_title	Int	否	是否填写抬头，0为否，1为是
//show_tax_no	Int	否	是否填写税号，0为否，1为是
//show_addr	Int	否	是否填写单位地址，0为否，1为是
//show_phone	Int	否	是否填写电话号码，0为否，1为是
//show_bank_type	Int	否	是否填写开户银行，0为否，1为是
//show_bank_no	Int	否	是否填写银行帐号，0为否，1为是
//custom_field	Object	否	自定义字段
//custom_field为list，每个对象包含以下字段：
//参数	类型	是否必填	描述
//key	String	是	自定义字段名称，最长5个字
func (c *Card) InvoiceSetAuthField(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	token.Set("action", "set_auth_field")
	return core.PostJSON(Link(cardInvoiceSetbizattr), token, p)
}

//InvoiceGetAuthField 查询授权页字段信息接口
//接口说明
//开发者可以通过该接口查看授权页抬头的填写项。
//请求说明url：
//https://api.weixin.qq.com/card/invoice/setbizattr?action=get_auth_field&access_token={access_token}
//请求方法：POST
func (c *Card) InvoiceGetAuthField() core.Responder {
	token := c.accessToken.KeyMap()
	token.Set("action", "get_auth_field")
	return core.PostJSON(Link(cardInvoiceSetbizattr), token, util.Map{})
}

//InvoiceGetAuthData 查询开票信息
//接口说明
//用户完成授权后，商户可以调用该接口查询某一个订单
//请求格式
//URL:
//https: //api.weixin.qq.com/card/invoice/getauthdata?access_token={access_token}
func (c *Card) InvoiceGetAuthData(orderID, appID string) core.Responder {
	token := c.accessToken.KeyMap()
	token.Set("action", "get_auth_field")
	return core.PostJSON(Link(cardInvoiceGetauthdata), token, util.Map{
		"order_id": orderID,
		"s_appid":  appID,
	})
}

//MemberActivate 接口激活
//	激活方式说明
//	接口激活通常需要开发者开发用户填写资料的网页。通常有两种激活流程：
//	用户必须在填写资料后才能领卡，领卡后开发者调用激活接口为用户激活会员卡；
//	是用户可以先领取会员卡，点击激活会员卡跳转至开发者设置的资料填写页面，填写完成后开发者调用激活接口为用户激活会员卡。
//	接口详情
//	接口调用请求说明
//	HTTP请求方式: POSTURL:https://api.weixin.qq.com/card/membercard/activate?access_token=TOKEN
func (c *Card) MemberActivate(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardMembercardActivate), token, p)
}

//SetMemberActivateUserForm 普通一键激活
//	一键激活是微信提供的快速便捷的激活方案，用户领取后点击“激活会员卡”会跳转至官方的资料填写页面，微信会自动拉取该用户之前填写过的开卡信息，用户无需重复填写， 同时避免了手机号验证的过程，从而实现一键激活的目的，提高了开卡率。具体流程如下图：
//	步骤一：在创建接口填入wx_activate字段
//	接口说明
//	设置微信一键开卡功能，现支持在创建会员卡时填入指定字段指定要一键激活，member_card中增加"wx_activate": true。 详情请见创建会员卡接口
//	若商户使用了自定义卡号，开发者可以设置用户填写信息后跳转至商户的网页，并由开发者进行激活。
//	参数说明
//	参数    是否必须    说明
//	member_card
//	wx_activate    否    填写true or false
//	开发者注意事项
//	1.填入了自动激活auto_activate字段，激活链接activate_url和一键开卡接口设置都会失效；
//	2.若同时传入了activate_url，则一键开卡接口设置会失效；
//	3.建议开发者activate_url、auto_activate和wx_activate只填写一项。
//	步骤二：设置开卡字段接口
//	开发者在创建时填入wx_activate字段后，需要调用该接口设置用户激活时需要填写的选项，否则一键开卡设置不生效。
//	接口调用请求说明
//	HTTP请求方式: POST
//	URL:https: //api.weixin.qq.com/card/membercard/activateuserform/set?access_token=TOKEN
func (c *Card) SetMemberActivateUserForm(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardMembercardActivateuserformSet), token, p)
}

// GetMemberUserInfo 查询会员信息
//	接口说明
//	支持开发者根据CardID和Code查询会员信息。
//	接口调用请求说明
//	HTTP请求方式: POST
//	URL:https://api.weixin.qq.com/card/membercard/userinfo/get?access_token=TOKEN
func (c *Card) GetMemberUserInfo(cardID, code string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardMembercardUserinfoGet), token, util.Map{
		"card_id": cardID,
		"code":    code,
	})
}

//GetMemberActivateTempInfo 设置开卡字段接口
//	步骤三：获取用户提交资料
//	用户填写并提交开卡资料后，会跳转到商户的网页，商户可以在网页内获取用户已填写的信息并进行开卡资质判断，信息确认等动作。
//	具体方式如下：
//	用户点击提交后，微信会在商户的url后面拼接获取用户填写信息的参数：activate_ticket、openid、card_id和加密code-encrypt_code,如商户填写的wx_activate_after_submit_url为www.qq.com,则拼接后的url为
//	www.qq.com&card_id=pbLatjvFdsLDUMoN8JqcsGeiMHKk&encrypt_code=Bupk8bb9xxxxxx3rdXV6fClBVtkHQplYohdzGvgDl4%3D&outer_str=&openid=obLatjjwDxxxxxxxoGIdwNqRXw&activate_ticket=fDZv9eMQAFfrNr3XBoqhb%2F%2BMSDM0yjDF6CdiUhC%2BOlEaxb0clsUxxxxxxxxxxxd6yQsjRMRu4kAcKTibYLN5tmHBdll1b6zQRsLF53MpKjGU%3D。
//	开发者可以根据activate_ticket获取到用户填写的信息，用于开发者页面的逻辑判断。
//	接口说明
//	支持开发者根据activate_ticket获取到用户填写的信息。
//	接口调用请求说明
//	HTTP请求方式: POSTURL:https://api.weixin.qq.com/card/membercard/activatetempinfo/get?access_token=TOKEN
func (c *Card) GetMemberActivateTempInfo(activateTicket string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardMembercardActivatetempinfoGet), token, util.Map{
		"activate_ticket": activateTicket,
	})
}

//MemberUpdateUser 更新会员信息
//	当会员持卡消费后，支持开发者调用该接口更新会员信息。会员卡交易后的每次信息变更需通过该接口通知微信，便于后续消息通知及其他扩展功能。
//	接口调用请求说明
//	HTTP请求方式: POSTURL:https: //api.weixin.qq.com/card/membercard/updateuser?access_token=TOKEN
func (c *Card) MemberUpdateUser(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardMembercardUpdateuser), token, p)
}

//AddPayGift 设置支付后投放卡券
//开通微信支付的商户可以设置在用户微信支付后自动为用户发送一条领卡消息，用户点击消息即可领取会员卡/优惠券。
//目前该功能仅支持微信支付商户号主体和制作会员卡公众号主体一致的情况下配置，否则报错。开发者可以登录
//“公众平台”-“公众号设置”、“微信支付商户平台首页”插卡企业主体信息是否一致。
//接口说明
//支持商户设置支付后投放卡券规则，可以区分时间段和金额区间发会员卡。
//接口调用请求说明
//HTTP请求方式: POSTURL:https://api.weixin.qq.com/card/paygiftcard/add?access_token=TOKEN
func (c *Card) AddPayGift(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardPaygiftcardAdd), token, p)
}

//MarkCode Mark(占用)Code接口
//朋友的券由于共享的特性，会出现多个消费者同时进入某一个卡券的自定义H5网页的情况，若该网页涉及线上下单、核销、支付等行为，会造成两个消费者同时使用同一张券，会有一个消费者使用失败的情况，为此我们设计了mark（占用）code接口。
//对于出示核销（消费者点击“出示使用”按钮）的场景，开发者直接调用核销接口，无需考虑mark逻辑，此时由客户端代为完成。
//对于消费者进入H5网页核销的情况，我们约定，开发者在帮助消费者核销卡券之前，必须帮助先将此code（卡券串码）与一个openid绑定（即mark住），才能进一步调用核销接口，否则报错。
//接口调用请求说明
//http请求方式: POST https://api.weixin.qq.com/card/code/mark?access_token=TOKEN
func (c *Card) MarkCode(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardCodeMark), token, p)
}

//GetBizUinInfo 拉取朋友的券数据接口
//	接口简介及开发注意事项
//	为支持开发者调用API查看卡券相关数据，微信卡券团队封装数据接口并面向具备卡券功能权限的开发者开放使用。开发者调用该接口可获取本商户下的所有卡券相关的总数据以及指定卡券的相关数据。
//	拉取卡券概况数据接口
//	接口说明
//	支持调用该接口拉取本商户的总体数据情况，包括时间区间内的各指标总量。
//	接口调用请求说明
//	http请求方式: POST https://api.weixin.qq.com/datacube/getcardbizuininfo?access_token=ACCESS_TOKEN
func (c *Card) GetBizUinInfo(beginDate, endDate time.Time, condSource int) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(datacubeGetcardbizuininfo), token, util.Map{
		"begin_date":  beginDate.Format(DatacubeTimeLayout), //请开发者按示例格式填写日期，否则会报错dateformaterror
		"end_date":    endDate.Format(DatacubeTimeLayout),
		"cond_source": condSource,
	})
}

//GetCardInfo 获取朋友的券数据接口
//	接口说明
//	支持开发者调用该接口拉取朋友的券在固定时间区间内的相关数据。
//	接口调用请求说明
//	http请求方式: POST https: //api.weixin.qq.com/datacube/getcardcardinfo?access_token=ACCESS_TOKEN
func (c *Card) GetCardInfo(cardID string, beginDate, endDate time.Time, condSource int) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(datacubeGetcardcardinfo), token, util.Map{
		"begin_date":  beginDate.Format(DatacubeTimeLayout), //请开发者按示例格式填写日期，否则会报错dateformaterror
		"end_date":    endDate.Format(DatacubeTimeLayout),
		"cond_source": condSource,
		"card_id":     cardID,
	})
}

// MovieUpdateUser ...
func (c *Card) MovieUpdateUser(p util.Map) core.Responder {
	token := c.accessToken.GetToken()
	return core.PostJSON(Link(cardMovieticketUpdateuser), token.KeyMap(), p)
}

// SubmitSubMerchant ...
func (c *Card) SubmitSubMerchant(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	p = p.Only([]string{"brand_name",
		"logo_url",
		"protocol",
		"end_time",
		"primary_category_id",
		"secondary_category_id",
		"agreement_media_id",
		"operator_media_id",
		"app_id"})
	return core.PostJSON(Link(cardSubmerchantSubmit), token, util.Map{"info": p})
}

// UpdateSubMerchant ...
func (c *Card) UpdateSubMerchant(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	p = p.Only([]string{
		"brand_name",
		"logo_url",
		"protocol",
		"end_time",
		"primary_category_id",
		"secondary_category_id",
		"agreement_media_id",
		"operator_media_id",
		"app_id",
	})
	return core.PostJSON(Link(cardSubmerchantUpdate), token, util.Map{"info": p})
}

// GetSubMerchant ...
func (c *Card) GetSubMerchant(mchID string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardSubmerchantget), token, util.Map{"merchant_id": mchID})

}

// BatchGetSubMerchant ...
func (c *Card) BatchGetSubMerchant(beginID, limit int, status string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardSubmerchantbatchget), token, util.Map{
		"begin_id": beginID,
		"limit":    limit,
		"status":   status,
	})
}

//GetCardAPITicket get ticket
func (c *Card) GetCardAPITicket(refresh bool) {
	c.jssdk.GetTicket("wx_card", refresh)
}

//NewOneCard 创建卡券信息
//	参数:
//	cardType 卡券类型
//	data	卡券信息 (可传nil)
func NewOneCard(cardType CardType, data util.Map) *OneCard {
	ct := strings.ToLower(cardType.String())
	return &OneCard{
		CardType: cardType,
		data: util.Map{
			"card_type": cardType,
			ct:          data,
		},
	}
}

//AddAdvancedInfo 添加卡券advanced_info
func (c *OneCard) AddAdvancedInfo(info *CardAdvancedInfo) *OneCard {
	return c.add("advanced_info", info)
}

//AddBaseInfo 添加卡券base_info
func (c *OneCard) AddBaseInfo(info *CardBaseInfo) *OneCard {
	return c.add("base_info", info)
}

//AddDealDetail 添加卡券deal_detail
func (c *OneCard) AddDealDetail(d string) *OneCard {
	return c.add("deal_detail", d)
}

func (c *OneCard) add(name string, info interface{}) *OneCard {
	ct := strings.ToLower(c.CardType.String())
	if c.data != nil {
		if v, b := c.data[ct].(util.Map); b {
			if v != nil {
				v[name] = info
			} else {
				v = util.Map{
					name: info,
				}
			}
			c.data[ct] = v
		}
	} else {
		c.data = util.Map{
			"card_type": c.CardType,
			ct: util.Map{
				name: info,
			},
		}
	}
	return c
}

//Set 设置卡券信息(包含base_info,advanced_info,deal_detail)
func (c *OneCard) Set(cardType CardType, data util.Map) {
	ct := strings.ToLower(cardType.String())
	c.CardType = cardType
	c.data = util.Map{
		"card_type": ct,
		ct:          data,
	}
}

//Get 获取卡券类型,卡券信息
func (c *OneCard) Get() (CardType, util.Map) {
	return c.CardType, c.data
}

// ToMap ...
func (c *OneCard) ToMap() util.Map {
	maps := util.Map{}
	v, err := json.Marshal(c)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(v, maps)
	if err != nil {
		return nil
	}
	return maps
}
