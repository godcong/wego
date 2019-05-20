package wego

import (
	"context"
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/cipher"
	"github.com/godcong/wego/util"
	"golang.org/x/xerrors"
	"strconv"
)

//Payment ...
type Payment struct {
	*PaymentProperty
	BodyType    BodyType
	client      *Client
	safeClient  *Client
	sandbox     *Sandbox
	publicKey   string
	privateKey  string
	subMchID    string
	subAppID    string
	remoteURL   string
	localHost   string
	notifyURL   string
	refundedURL string
	scannedURL  string
}

// NewPayment ...
func NewPayment(config *PaymentProperty, options ...PaymentOption) *Payment {
	payment := &Payment{
		PaymentProperty: config,
		BodyType:        BodyTypeXML,
	}
	payment.parseOption(options...)

	return payment
}

func (obj *Payment) parseOption(options ...PaymentOption) {
	if options == nil {
		return
	}
	for _, o := range options {
		o(obj)
	}
}

// Client ...
func (obj *Payment) Client() *Client {
	if obj.client == nil {
		obj.client = NewClient(ClientBodyType(obj.BodyType))
	}
	return obj.client
}

// SafeClient ...
func (obj *Payment) SafeClient() *Client {
	if obj.safeClient == nil {
		obj.safeClient = NewClient(ClientBodyType(obj.BodyType), ClientSafeCert(obj.SafeCert))
	}
	return obj.safeClient
}

// SetKey ...
func (obj *Payment) SetKey(public, private string) {
	obj.privateKey = private
	obj.publicKey = public
}

// SetSubID ...
func (obj *Payment) SetSubID(mchid, appid string) {
	obj.subMchID = mchid
	obj.subAppID = appid
}

// HandleRefunded ...
func (obj *Payment) HandleRefunded(hook RequestHook) Notifier {
	return &paymentRefundedNotify{
		cipher:      cipher.New(cipher.AES256ECB, cipher.OptionKey(obj.Key)),
		RequestHook: hook,
	}
}

// HandleRefundedNotify ...
func (obj *Payment) HandleRefundedNotify(hook RequestHook) ServeHTTPFunc {
	return obj.HandleRefunded(hook).ServeHTTP
}

// HandleScannedNotify ...
func (obj *Payment) HandleScannedNotify(hook RequestHook) Notifier {
	return &paymentScannedNotify{
		Payment:     obj,
		RequestHook: hook,
	}
}

// HandleScanned ...
func (obj *Payment) HandleScanned(hook RequestHook) ServeHTTPFunc {
	return obj.HandleScannedNotify(hook).ServeHTTP
}

// HandlePaidNotify ...
func (obj *Payment) HandlePaidNotify(hook RequestHook) Notifier {
	return &paymentPaidNotify{
		Payment:     obj,
		RequestHook: hook,
	}
}

// HandlePaid ...
func (obj *Payment) HandlePaid(hook RequestHook) ServeHTTPFunc {
	return obj.HandlePaidNotify(hook).ServeHTTP
}

/*Pay 支付
接口地址
SDK下载:https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=11_1
https://api.mch.weixin.qq.com/pay/micropay
输入参数
名称	变量名	必填	类型	示例值	描述
公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID（企业号corpid即为此appId）
商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
设备号	device_info	否	String(32)	013467007045764	终端设备号(商户自定义，如门店编号)
随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
商品描述	body	是	String(128)	image形象店-深圳腾大- QQ公仔	商品简单描述，该字段须严格按照规范传递，具体请见参数规定
商品详情	detail	否	String(6000)
单品优惠功能字段，需要接入详见单品优惠详细说明
附加数据	attach	否	String(127)	说明	附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
商户订单号	out_trade_no	是	String(32)	1217752501201407033233368018	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。详见商户订单号
订单金额	total_fee	是	Int	888	订单总金额，单位为分，只能为整数，详见支付金额
货币类型	fee_type	否	String(16)	CNY	符合ISO4217标准的三位字母代码，默认人民币:CNY，详见货币类型
终端IP	spbill_create_ip	是	String(16)	8.8.8.8	调用微信支付API的机器IP
订单优惠标记	goods_tag	否	String(32)	1234	订单优惠标记，代金券或立减优惠功能的参数，详见代金券或立减优惠
指定支付方式	limit_pay	否	String(32)	no_credit	no_credit--指定不能使用信用卡支付
交易起始时间	time_start	否	String(14)	20091225091010	订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
交易结束时间	time_expire	否	String(14)	20091227091010 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。注意:最短失效时间间隔需大于1分钟
授权码	auth_code	是	String(128)	120061098828009406	扫码支付授权码，设备读取用户微信中的条码或者二维码信息（注:用户刷卡条形码规则:18位纯数字，以10、11、12、13、14、15开头）
+场景信息	scene_info	否	String(256)
该字段用于上报场景信息，目前支持上报实际门店信息。该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }} ，字段详细说明请点击行前的+展开
*/
func (obj *Payment) Pay(able util.MapAble) Responder {
	p := able.ToMap()
	p.Set("appid", obj.AppID)

	//set notify callback
	notify := obj.NotifyURL()
	if !p.Has("notify_url") {
		p.Set("notify_url", notify)
	}

	return obj.Request(payMicroPay, p)
}

// UnifyResponse ...
type UnifyResponse struct {
	AppID      string `xml:"appid"`
	CodeURL    string `xml:"code_url"`
	DeviceInfo string `xml:"device_info"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	MchID      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	PrepayID   string `xml:"prepay_id"`
	ResultCode string `xml:"result_code"`
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Sign       string `xml:"sign"`
	TradeType  string `xml:"trade_type"`
}

//DownloadFundFlow 下载资金账单
//资金账单日期	bill_date	是	String(8)	20140603	下载对账单的日期，格式：20140603
//资金账户类型	account_type	是	String(8)	Basic
//账单的资金来源账户：
//Basic  基本账户
//Operation 运营账户
//Fees 手续费账户
//压缩账单	tar_type	否	String(8)	GZIP	非必传参数，固定值：GZIP，返回格式为.gzip的压缩包账单。不传则默认为数据流形式。
func (obj *Payment) DownloadFundFlow(bd string, at string, options ...util.Map) Responder {
	m := util.CombineMaps(util.Map{
		"appid":        obj.AppID,
		"bill_date":    bd,
		"sign_type":    util.HMACSHA256,
		"account_type": at,
	}, options...)

	return obj.SafeRequest(payDownloadFundFlow, m)
}

/*ReverseByOutTradeNumber 通过out_trade_no撤销订单
接口地址
https://apihk.mch.weixin.qq.com/secapi/pay/reverse    （建议接入点:东南亚）
https://apius.mch.weixin.qq.com/secapi/pay/reverse    （建议接入点:其它）
https://api.mch.weixin.qq.com/secapi/pay/reverse        （建议接入点:中国国内）
注:商户可根据实际请求情况选择最优域名进行访问，建议在接入时做好兼容，当访问其中一个域名出现异常时，可自动切换为其他域名。
是否需要证书
请求需要双向证书。 详见证书使用
请求参数
字段名	变量名	类型	必填	示例值	描述
商户订单号	out_trade_no	String(32)	是	1217752501201407033233368018	商户系统内部的订单号,transaction_id、out_trade_no二选一，如果同时存在优先级:transaction_id> out_trade_no
*/
func (obj *Payment) ReverseByOutTradeNumber(num string) Responder {
	return obj.reverse(util.Map{"out_trade_no": num})
}

/*ReverseByTransactionID 通过transaction_id撤销订单
接口地址
https://apihk.mch.weixin.qq.com/secapi/pay/reverse    （建议接入点:东南亚）
https://apius.mch.weixin.qq.com/secapi/pay/reverse    （建议接入点:其它）
https://api.mch.weixin.qq.com/secapi/pay/reverse        （建议接入点:中国国内）
注:商户可根据实际请求情况选择最优域名进行访问，建议在接入时做好兼容，当访问其中一个域名出现异常时，可自动切换为其他域名。
是否需要证书
请求需要双向证书。 详见证书使用
请求参数
字段名	变量名	类型	必填	示例值	描述
微信订单号	transaction_id	String(32)	否	1217752501201407033233368018	微信的订单号，优先使用
*/
func (obj *Payment) ReverseByTransactionID(id string) Responder {
	return obj.reverse(util.Map{"transaction_id": id})
}

func (obj *Payment) reverse(m util.Map) Responder {
	return obj.SafeRequest(payReverse, m)
}

/*Unify 统一下单
字段名	变量名	必填	类型	示例值	描述
商品描述	body	是	String(128)	Ipad mini  16G  白色	商品或支付单简要描述
商户订单号	out_trade_no	是	String(32)	20150806125346	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。 其他说明见商户订单号
标价金额	total_fee	是	Int	888	标价金额，单位为该币种最小计算单位，只能为整数，详见标价金额
交易类型	trade_type	是	String(16)	JSAPI	取值如下:JSAPI，NATIVE，APP，详细说明见参数规定
用户标识	openid	否	String(128)	oUpF8uMuAJO_M2pxb1Q9zNjWeS6o	trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识。openid如何获取，可参考【获取openid】。企业号请使用【企业号OAuth2.0接口】获取企业号内成员userid，再调用【企业号userid转openid接口】进行转换
*/
func (obj *Payment) Unify(m util.Map, options ...util.Map) Responder {
	m = util.CombineMaps(m, options...)
	if !m.Has("spbill_create_ip") {
		if m.Get("trade_type") == "NATIVE" {
			m.Set("spbill_create_ip", util.GetServerIP())
		}
		//TODO: getclientip with request
		//if obj.request != nil {
		//	m.Set("spbill_create_ip", core.GetClientIP(obj.request))
		//}
	}

	if !m.Has("notify_url") {
		m.Set("notify_url", obj.NotifyURL())
	}

	return obj.Request(payUnifiedOrder, m)
}

/*AuthCodeToOpenid 通过授权码查询公众号Openid
接口链接: https://api.mch.weixin.qq.com/tools/authcodetoopenid
通过授权码查询公众号Openid，调用查询后，该授权码只能由此商户号发起扣款，直至授权码更新。
参数:
authCode - 授权码
返回:
openid string
*/
func (obj *Payment) AuthCodeToOpenid(authCode string) Responder {
	m := make(util.Map)
	m.Set("appid", obj.AppID)
	m.Set("auth_code", authCode)
	return obj.Request(authCodeToOpenid, m)
}

//BillDownload 下载对账单
//接口链接
//https://api.mch.weixin.qq.com/pay/downloadbill
//是否需要证书
//不需要。
//请求参数
//字段名	变量名	必填	类型	示例值	描述
//对账单日期	bill_date	是	String(8)	20140603	下载对账单的日期，格式:20140603
func (obj *Payment) BillDownload(bd string, opts ...util.Map) Responder {
	m := util.CombineMaps(util.Map{
		"appid":     obj.AppID,
		"bill_date": bd,
	}, opts...)

	if !m.Has("bill_type") {
		m.Set("bill_type", "ALL")
	}

	return obj.Request(payDownloadBill, m)
}

//BillDownloadFundFlow 下载资金账单
//资金账单日期	bill_date	是	String(8)	20140603	下载对账单的日期，格式：20140603
//资金账户类型	account_type	是	String(8)	Basic
//账单的资金来源账户：
//Basic  基本账户
//Operation 运营账户
//Fees 手续费账户
//压缩账单	tar_type	否	String(8)	GZIP	非必传参数，固定值：GZIP，返回格式为.gzip的压缩包账单。不传则默认为数据流形式。
func (obj *Payment) BillDownloadFundFlow(bd string, at string, opts ...util.Map) Responder {
	m := util.CombineMaps(util.Map{
		"appid":        obj.AppID,
		"bill_date":    bd,
		"sign_type":    util.HMACSHA256,
		"account_type": at,
	}, opts...)

	return obj.SafeRequest(payDownloadFundFlow, m)
}

//BillBatchQueryComment 拉取订单评价数据
//接口链接
//https://api.mch.weixin.qq.com/billcommentsp/batchquerycomment
//开始时间 begin_time 是 String(19) 20170724000000 按用户评论时间批量拉取的起始时间，格式为yyyyMMddHHmmss
//结束时间 end_time 是 String(19) 20170725000000 按用户评论时间批量拉取的结束时间，格式为yyyyMMddHHmmss
//位移 offset 是 uint(64) 0 指定从某条记录的下一条开始返回记录。接口调用成功时，会返回本次查询最后一条数据的offset。商户需要翻页时，应该把本次调用返回的offset 作为下次调用的入参。注意offset是评论数据在微信支付后台保存的索引，未必是连续的
//条数 limit 否 uint(32) 100 一次拉取的条数, 最大值是200，默认是200
func (obj *Payment) BillBatchQueryComment(beginTime, endTime string, offset int, opts ...util.Map) Responder {
	m := util.CombineMaps(util.Map{
		"begin_time": beginTime,
		"end_time":   endTime,
		"offset":     strconv.Itoa(offset),
		"appid":      obj.AppID,
		"sign_type":  util.HMACSHA256,
	}, opts...)

	return obj.client.Post(context.Background(), obj.RequestURL(batchQueryComment), nil,
		obj.initPay(m, util.FieldSign, util.FieldSignType, util.FieldLimit))

}

// CouponSend ...
func (obj *Payment) CouponSend(opts ...util.Map) Responder {
	m := util.CombineMaps(util.Map{}, opts...)
	m.Set("appid", obj.AppID)
	m.Set("openid_count", 1)
	return obj.SafeRequest(mmpaymkttransfersSendCoupon, m)
}

// CouponQueryStock ...
func (obj *Payment) CouponQueryStock(opts ...util.Map) Responder {
	m := util.CombineMaps(util.Map{}, opts...)
	m.Set("appid", obj.AppID)
	return obj.SafeRequest(mmpaymkttransfersQueryCouponStock, m)
}

// CouponQueryInfo ...
func (obj *Payment) CouponQueryInfo(opts ...util.Map) Responder {
	m := util.CombineMaps(util.Map{}, opts...)
	m.Set("appid", obj.AppID)
	return obj.SafeRequest(mmpaymkttransfersQueryCouponsInfo, m)
}

// MerchantAddSubMerchant ...
func (obj *Payment) MerchantAddSubMerchant(maps util.Map) Responder {
	return obj.merchantManage("add", maps)
}

// MerchantQuerySubMerchantByMerchantID ...
func (obj *Payment) MerchantQuerySubMerchantByMerchantID(id string) Responder {
	return obj.merchantManage("dummyQuery", util.Map{"micro_mch_id": id})
}

// MerchantQuerySubMerchantByWeChatID ...
func (obj *Payment) MerchantQuerySubMerchantByWeChatID(id string) Responder {
	return obj.merchantManage("dummyQuery", util.Map{"recipient_wechatid": id})
}

// MerchantModifyInfo ...
func (obj *Payment) MerchantModifyInfo(maps util.Map) Responder {
	maps.Join(util.Map{
		"mch_id":     obj.MchID,
		"sub_mch_id": "",
	})
	return obj.SafeRequest(mchModifymchinfo, maps)
}

// MerchantAddRecommendConfBySubscribe ...
func (obj *Payment) MerchantAddRecommendConfBySubscribe(appID string) Responder {
	maps := util.Map{
		"subscribe_appid": appID,
		"mch_id":          obj.MchID,
		"sub_mch_id":      "",
		"sub_appid":       "",
	}
	return obj.SafeRequest(mktAddrecommendconf, maps)
}

// MerchantAddRecommendConfByReceipt ...
func (obj *Payment) MerchantAddRecommendConfByReceipt(appID string) Responder {
	maps := util.Map{
		"receipt_appid": appID,
		"mch_id":        obj.MchID,
		"sub_mch_id":    "",
		"sub_appid":     "",
	}
	return obj.SafeRequest(mktAddrecommendconf, maps)
}

func (obj *Payment) mchAddSubDevConfig() {
	//TODO
}

func (obj *Payment) merchantManage(action string, opts ...util.Map) Responder {
	m := util.CombineMaps(util.Map{
		"appid": obj.AppID,
	}, opts...)

	query := util.Map{"action": action}
	return obj.SafeClient().Post(context.Background(), obj.RequestURL(mchSubMchManage), query, obj.initPay(m))
}

/*OrderClose 关闭订单
字段名	变量名	必填	类型	示例值	描述
商户订单号	out_trade_no	是	String(32)	1217752501201407033233368018	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
*/
func (obj *Payment) OrderClose(no string) Responder {
	m := make(util.Map)
	m.Set("appid", obj.AppID)
	m.Set("out_trade_no", no)
	return obj.Request(payCloseOrder, m)
}

/** QueryOrder 查询订单
* 场景:刷卡支付、公共号支付、扫码支付、APP支付
接口地址
https://apihk.mch.weixin.qq.com/pay/orderquery    （建议接入点:东南亚）
https://apius.mch.weixin.qq.com/pay/orderquery    （建议接入点:其它）
https://api.mch.weixin.qq.com/pay/orderquery        （建议接入点:中国国内）
注:商户可根据实际请求情况选择最优域名进行访问，建议在接入时做好兼容，当访问其中一个域名出现异常时，可自动切换为其他域名。
是否需要证书
不需要
请求参数
字段名	变量名	必填	类型	示例值	描述
公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信分配的公众账号ID（企业号corpid即为此appId）
商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
微信订单号	transaction_id	二选一	String(32)	1009660380201506130728806387	微信的订单号，优先使用
商户订单号	out_trade_no	String(32)	20150806125346	商户系统内部的订单号，当没提供transaction_id时需要传这个。
随机字符串	nonce_str	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	随机字符串，不长于32位。推荐随机数生成算法
签名	sign	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	签名，详见签名生成算法
签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
*/
func (obj *Payment) orderQuery(m util.Map) Responder {
	return obj.Request(payOrderQuery, m)
}

/*OrderQueryByTransactionID 通过transaction_id查询订单
场景:刷卡支付、公共号支付、扫码支付、APP支付
字段名	变量名	必填	类型	示例值	描述
微信订单号	transaction_id	String(32)	1009660380201506130728806387	微信的订单号，优先使用
*/
func (obj *Payment) OrderQueryByTransactionID(id string) Responder {
	return obj.orderQuery(util.Map{"transaction_id": id})
}

/*OrderQueryByOutTradeNumber 通过out_trade_no查询订单
场景:刷卡支付、公共号支付、扫码支付、APP支付
字段名	变量名	必填	类型	示例值	描述
公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信分配的公众账号ID（企业号corpid即为此appId）
商户订单号	out_trade_no	String(32)	20150806125346	商户系统内部的订单号，当没提供transaction_id时需要传这个。
*/
func (obj *Payment) OrderQueryByOutTradeNumber(no string) Responder {
	return obj.orderQuery(util.Map{"out_trade_no": no})
}

/*RedPackInfo 查询红包记录
接口调用请求说明
请求Url	https://api.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo
是否需要证书	是（证书及使用说明详见商户证书）
请求方式	POST
请求参数
字段名	字段	必填	示例值	类型	说明
商户订单号	mch_billno	是	10000098201411111234567890	String(28)	商户发放红包的商户订单号
*/
func (obj *Payment) RedPackInfo(mchBillNo string) Responder {
	m := util.Map{
		"mch_billno": mchBillNo,
		"appid":      "app_id",
		"bill_type":  "MCHT",
	}
	return obj.SafeRequest(mmpaymkttransfersGetHbInfo, m)

}

/*RedPackSendNormal 发放普通红包
发放规则
1.发送频率限制------默认1800/min
2.发送个数上限------按照默认1800/min算
3.金额限制------默认红包金额为1-200元，如有需要，可前往商户平台进行设置和申请
4.其他其他限制吗？------单个用户可领取红包上线为10个/天，如有需要，可前往商户平台进行设置和申请
5.如果量上满足不了我们的需求，如何提高各个上限？------金额上限和用户当天领取次数上限可以在商户平台进行设置
注意-红包金额大于200或者小于1元时，请求参数scene_id必传，参数说明见下文。
注意2-根据监管要求，新申请商户号使用现金红包需要满足两个条件:1、入驻时间超过90天 2、连续正常交易30天。
注意3-移动应用的appid无法使用红包接口。
消息触达规则
现金红包发放后会以公众号消息的形式触达用户，不同情况下触达消息的形式会有差别，规则如下:
是否关注	关注时间	是否接收消息	触达消息
否	/	/	模版消息
是	<=50小时	是	模版消息
是	<=50小时	否	模版消息
是	>50小时	是	防伪消息
是	>50小时	否	模版消息
接口调用请求说明
请求Url	https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack
是否需要证书	是（证书及使用说明详见商户证书）
请求方式	POST
请求参数
字段名	字段	必填	示例值	类型	说明
随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	String(32)	随机字符串，不长于32位
签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	String(32)	详见签名生成算法
商户订单号	mch_billno	是	10000098201411111234567890	String(28)
商户订单号（每个订单号必须唯一。取值范围:0~9，a~z，A~Z）
接口根据商户订单号支持重入，如出现超时可再调用。
商户号	mch_id	是	10000098	String(32)	微信支付分配的商户号
子商户号	sub_mch_id	否	10000090	String(32)	微信支付分配的子商户号，服务商模式下必填
公众账号appid	wxappid	是	wx8888888888888888	String(32)	微信分配的公众账号ID（企业号corpid即为此appId）。接口传入的所有appid应该为公众号的appid（在mp.weixin.qq.com申请的），不能为APP的appid（在open.weixin.qq.com申请的）。
触达用户appid	msgappid	是	wx28b16568a629bb33	String(32)	服务商模式下触达用户时的appid(可填服务商自己的appid或子商户的appid)，服务商模式下必填，服务商模式下填入的子商户appid必须在微信支付商户平台中先录入，否则会校验不过。
商户名称	send_name	是	天虹百货	String(32)	红包发送者名称
用户openid	re_openid	是	oxTWIuGaIt6gTKsQRLau2M0yL16E	String(32)
接受红包的用户
用户在wxappid下的openid，服务商模式下可填入msgappid下的openid。
付款金额	total_amount	是	1000	int	付款金额，单位分
红包发放总人数	total_num	是	1	int
红包发放总人数
total_num=1
红包祝福语	wishing	是	感谢您参加猜灯谜活动，祝您元宵节快乐！	String(128)	红包祝福语
Ip地址	client_ip	是	192.168.0.1	String(15)	调用接口的机器Ip地址
活动名称	act_name	是	猜灯谜抢红包活动	String(32)	活动名称
备注	remark	是	猜越多得越多，快来抢！	String(256)	备注信息
场景id	scene_id	否	PRODUCT_8	String(32)
发放红包使用场景，红包金额大于200或者小于1元时必传
PRODUCT_1:商品促销
PRODUCT_2:抽奖
PRODUCT_3:虚拟物品兑奖
PRODUCT_4:企业内部福利
PRODUCT_5:渠道分润
PRODUCT_6:保险回馈
PRODUCT_7:彩票派奖
PRODUCT_8:税务刮奖
活动信息	risk_info	否	posttime%3d123123412%26clientversion%3d234134%26mobile%3d122344545%26deviceid%3dIOS	String(128)
posttime:用户操作的时间戳
mobile:业务系统账号的手机号，国家代码-手机号。不需要+号
deviceid :mac 地址或者设备唯一标识
clientversion :用户操作的客户端版本
把值为非空的信息用key=value进行拼接，再进行urlencode
urlencode(posttime=xx& mobile =xx&deviceid=xx)
扣钱方mchid	consume_mch_id	否	10000098	String(32)	常规模式下无效，服务商模式下选填，服务商模式下不填默认扣子商户的钱
*/
func (obj *Payment) RedPackSendNormal(m util.Map) Responder {
	m.Set("total_num", strconv.Itoa(1))
	m.Set("client_ip", util.GetServerIP())
	m.Set("wxappid", obj.AppID)
	return obj.SafeRequest(mmpaymkttransfersSendRedPack, m)
}

/*RedPackSendGroup 裂变红包
发放规则
裂变红包:一次可以发放一组红包。首先领取的用户为种子用户，种子用户领取一组红包当中的一个，并可以通过社交分享将剩下的红包给其他用户。裂变红包充分利用了人际传播的优势。
消息触达规则
现金红包发放后会以公众号消息的形式触达用户，不同情况下触达消息的形式会有差别，规则如下:
是否关注	关注时间	是否接收消息	触达消息
否	/	/	模版消息
是	<=50小时	是	模版消息
是	<=50小时	否	模版消息
是	>50小时	是	防伪消息
是	>50小时	否	模版消息
接口调用请求说明
请求Url	https://api.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack
是否需要证书	是（证书及使用说明详见商户证书）
请求方式	POST
请求参数
字段名	字段	必填	示例值	类型	说明
随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	String(32)	随机字符串，不长于32位
签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	String(32)	详见签名生成算法
商户订单号	mch_billno	是	10000098201411111234567890	String(28)
商户订单号（每个订单号必须唯一）
组成:mch_id+yyyymmdd+10位一天内不能重复的数字。
接口根据商户订单号支持重入，如出现超时可再调用。
商户号	mch_id	是	10000098	String(32)	微信支付分配的商户号
子商户号	sub_mch_id	否	10000090	String(32)	微信支付分配的子商户号，服务商模式下必填
公众账号appid	wxappid	是	wx8888888888888888	String(32)	微信分配的公众账号ID（企业号corpid即为此appId），接口传入的所有appid应该为公众号的appid（在mp.weixin.qq.com申请的），不能为APP的appid（在open.weixin.qq.com申请的）。
触达用户appid	msgappid	否	wx28b16568a629bb33	String(32)	服务商模式下触达用户时的appid(可填服务商自己的appid或子商户的appid)，服务商模式下必填，服务商模式下填入的子商户appid必须在微信支付商户平台中先录入，否则会校验不过。
商户名称	send_name	是	天虹百货	String(32)	红包发送者名称
用户openid	re_openid	是	oxTWIuGaIt6gTKsQRLau2M0yL16E	String(32)
接收红包的种子用户（首个用户）
用户在wxappid下的openid
付款金额	total_amount	是	1000	int	红包发放总金额，即一组红包金额总和，包括分享者的红包和裂变的红包，单位分
红包发放总人数	total_num	是	3	int	红包发放总人数，即总共有多少人可以领到该组红包（包括分享者）
红包金额设置方式	amt_type	是	ALL_RAND	String(32)
红包金额设置方式
ALL_RAND—全部随机,商户指定总金额和红包发放总人数，由微信支付随机计算出各红包金额
红包祝福语	wishing	是	感谢您参加猜灯谜活动，祝您元宵节快乐！	String(128)	红包祝福语
活动名称	act_name	是	猜灯谜抢红包活动	String(32)	活动名称
备注	remark	是	猜越多得越多，快来抢！	String(256)	备注信息
场景id	scene_id	否	PRODUCT_8	String(32)
PRODUCT_1:商品促销
PRODUCT_2:抽奖
PRODUCT_3:虚拟物品兑奖
PRODUCT_4:企业内部福利
PRODUCT_5:渠道分润
PRODUCT_6:保险回馈
PRODUCT_7:彩票派奖
PRODUCT_8:税务刮奖
活动信息	risk_info	否	posttime%3d123123412%26clientversion%3d234134%26mobile%3d122344545%26deviceid%3dIOS	String(128)
posttime:用户操作的时间戳
mobile:业务系统账号的手机号，国家代码-手机号。不需要+号
deviceid :mac 地址或者设备唯一标识
clientversion :用户操作的客户端版本
把值为非空的信息用key=value进行拼接，再进行urlencode
urlencode(posttime=xx& mobile =xx&deviceid=xx)
资金授权商户号	consume_mch_id	否	1222000096	String(32)
资金授权商户号
服务商替特约商户发放时使用
扣钱方mchid	consume_mch_id	否	10000098	String(32)	常规模式下无效，服务商模式下选填，服务商模式下不填默认扣子商户的钱
*/
func (obj *Payment) RedPackSendGroup(m util.Map) Responder {
	m.Set("amt_type", "ALL_RAND")
	m.Set("wxappid", obj.AppID)
	return obj.SafeRequest(mmpaymkttransfersSendGroupRedPack, m)
}

/*RefundByOutTradeNumber 按照out_trade_no发起退款
接口地址
接口链接:https://api.mch.weixin.qq.com/secapi/pay/refund
*/
func (obj *Payment) RefundByOutTradeNumber(tradeNum, num string, total, refund int, opts ...util.Map) Responder {
	m := util.CombineMaps(util.Map{"out_trade_no": tradeNum}, opts...)
	return obj.refund(num, total, refund, m)
}

/*RefundByTransactionID 按照transaction_id发起退款
接口地址
接口链接:https://api.mch.weixin.qq.com/secapi/pay/refund
*/
func (obj *Payment) RefundByTransactionID(tid, num string, total, refund int, opts ...util.Map) Responder {
	m := util.CombineMaps(util.Map{"transaction_id": tid}, opts...)
	return obj.refund(num, total, refund, m)
}

func (obj *Payment) refundQuery(m util.Map) Responder {
	return obj.Request(payRefundQuery, m)
}

/*RefundQueryByRefundID 按refund_id查找退款订单
接口地址
接口链接:https://api.mch.weixin.qq.com/pay/refundquery
*/
func (obj *Payment) RefundQueryByRefundID(id string) Responder {
	return obj.refundQuery(util.Map{"refund_id": id})
}

/*RefundQueryByOutRefundNumber 按out_refund_no查找退款订单
接口地址
接口链接:https://api.mch.weixin.qq.com/pay/refundquery
*/
func (obj *Payment) RefundQueryByOutRefundNumber(id string) Responder {
	return obj.refundQuery(util.Map{"out_refund_no": id})
}

/*RefundQueryByOutTradeNumber 按out_trade_no查找退款订单
接口地址
接口链接:https://api.mch.weixin.qq.com/pay/refundquery
*/
func (obj *Payment) RefundQueryByOutTradeNumber(id string) Responder {
	return obj.refundQuery(util.Map{"out_trade_no": id})
}

/*RefundQueryByTransactionID 按transaction_id查找退款订单
接口地址
接口链接:https://api.mch.weixin.qq.com/pay/refundquery
*/
func (obj *Payment) RefundQueryByTransactionID(id string) Responder {
	return obj.refundQuery(util.Map{"transaction_id": id})
}

func (obj *Payment) refund(num string, total, refund int, opts ...util.Map) Responder {
	m := util.CombineMaps(util.Map{
		"out_refund_no": num,
		"total_fee":     strconv.Itoa(total),
		"refund_fee":    strconv.Itoa(refund),
	})

	//set notify callback
	notify := obj.RefundURL()
	if !m.Has("notify_url") {
		m.Set("notify_url", notify)
	}
	return obj.SafeRequest(payRefund, m)
}

/*GetPublicKey 获取RSA加密公钥API
接口说明
请求Url	https://fraud.mch.weixin.qq.com/risk/getpublickey
是否需要证书	请求需要双向证书。 详见证书使用
请求方式	POST

PS: 可使用SaveTo保存Key.需转换成PKCS#8使用.
RSA公钥格式PKCS#1,PKCS#8互转说明
PKCS#1 转 PKCS#8:
openssl rsa -RSAPublicKey_in -in <filename> -pubout
PKCS#8 转 PKCS#1:
openssl rsa -pubin -in <filename> -RSAPublicKey_out
*/
func (obj *Payment) GetPublicKey() Responder {
	m := util.Map{"sign_type": "MD5"}
	return obj.SafeRequest(riskGetPublicKey, obj.initPay(m))
}

/*TransferQueryBalanceOrder 查询企业付款
接口说明
请求Url	https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo
是否需要证书	请求需要双向证书。 详见证书使用
请求方式	POST
请求参数
字段名	字段	必填	示例值	类型	说明
随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	String(32)	随机字符串，不长于32位
签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	String(32)	生成签名方式查看3.2.1节
商户订单号	partner_trade_no	是	10000098201411111234567890	String(28)	商户调用企业付款API时使用的商户订单号
商户号	mch_id	是	10000098	String(32)	微信支付分配的商户号
Appid	appid	是	wxe062425f740d30d8	String(32)	商户号的appid
*/
func (obj *Payment) TransferQueryBalanceOrder(s string) Responder {
	m := util.CombineMaps(util.Map{
		//"mchid":            obj.MchID,
		//"mch_appid":        obj.AppID,
		"partner_trade_no": s,
	})
	return obj.SafeRequest(mmpaymkttransfersGetTransferInfo, m)
}

/*TransferToBalance 企业付款
接口地址
接口链接:https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers
是否需要证书
请求需要双向证书。 详见证书使用
请求参数
字段名	变量名	必填	示例值	类型	描述
商户账号appid	mch_appid	是	wx8888888888888888	String	申请商户号的appid或商户号绑定的appid
商户号	mchid	是	1900000109	String(32)	微信支付分配的商户号
设备号	device_info	否	013467007045764	String(32)	微信支付分配的终端设备号
随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	String(32)	随机字符串，不长于32位
签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	String(32)	签名，详见签名算法
商户订单号	partner_trade_no	是	10000098201411111234567890	String	商户订单号，需保持唯一性
(只能是字母或者数字，不能包含有符号)
用户openid	openid	是	oxTWIuGaIt6gTKsQRLau2M0yL16E	String	商户appid下，某用户的openid
校验用户姓名选项	check_name	是	FORCE_CHECK	String	NO_CHECK:不校验真实姓名
FORCE_CHECK:强校验真实姓名
收款用户姓名	re_user_name	可选	王小王	String	收款用户真实姓名。
如果check_name设置为FORCE_CHECK，则必填用户真实姓名
金额	amount	是	10099	int	企业付款金额，单位为分
企业付款描述信息	desc	是	理赔	String	企业付款操作说明信息。必填。
Ip地址	spbill_create_ip	是	192.168.0.1	String(32)	该IP同在商户平台设置的IP白名单中的IP没有关联，该IP可传用户端或者服务端的IP。
*/
func (obj *Payment) TransferToBalance(opts ...util.Map) Responder {
	m := util.CombineMaps(util.Map{
		"mchid":     obj.MchID,
		"mch_appid": obj.AppID,
	}, opts...)

	if !m.Has("spbill_create_ip") {
		m.Set("spbill_create_ip", util.GetServerIP())
	}
	return obj.SafeRequest(mmpaymkttransfersPromotionTransfers, m)
}

/*TransferQueryBankCardOrder 查询企业付款银行卡API
接口说明
请求Url	https://api.mch.weixin.qq.com/mmpaysptrans/query_bank
是否需要证书	请求需要双向证书。 详见证书使用
请求方式	POST
请求参数
字段名	字段	必填	示例值	类型	说明
商户号	mch_id	是	1900000109	string(32)	微信支付分配的商户号
商户企业付款单号	partner_trade_no	是	1212121221227	string(32)	商户订单号，需保持唯一（只允许数字[0~9]或字母[A~Z]和[a~z]最短8位，最长32位）
随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67Vs	string(32)	随机字符串，长度小于32位
签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	string(32)	商户自带签名
返回参数
字段名	变量名	必填	类型	说明
返回状态码	return_code	是	String(16)
SUCCESS/FAIL
此字段是通信标识，非付款标识，付款是否成功需要查看result_code来判断
返回信息	return_msg	否	String(128)
返回信息，如非空，为错误原因
签名失败
参数格式校验错误
以下字段在return_code为SUCCESS的时候有返回
业务结果	result_code	是	String(16)	SUCCESS/FAIL，非付款标识，付款是否成功需要查看status字段来判断
错误代码	err_code	否	String(32)	错误码信息
错误代码描述	err_code_des	否	String(128)	结果信息描述
以下字段在return_code 和result_code都为SUCCESS的时候有返回
商户号	mch_id	是	string(32)	商户号
商户企业付款单号	partner_trade_no	是	string(32)	商户单号
微信企业付款单号	payment_no	是	string(64)	即为微信内部业务单号
银行卡号	bank_no_md5	是	string(32)	收款用户银行卡号(MD5加密)
用户真实姓名	true_name_md5	是	string(32)	收款人真实姓名（MD5加密）
代付金额	amount	是	int	代付订单金额RMB:分
代付单状态	status	是	string	代付订单状态:
PROCESSING（处理中，如有明确失败，则返回额外失败原因；否则没有错误原因）
SUCCESS（付款成功）
FAILED（付款失败,需要替换付款单号重新发起付款）
BANK_FAIL（银行退票，订单状态由付款成功流转至退票,退票时付款金额和手续费会自动退还）
手续费金额	cmms_amt	是	int	手续费订单金额 RMB:分
商户下单时间	create_time	是	String	微信侧订单创建时间
成功付款时间	pay_succ_time	否	String	微信侧付款成功时间（但无法保证银行不会退票）
失败原因	reason	否	String	订单失败原因（如:余额不足）
*/
func (obj *Payment) TransferQueryBankCardOrder(s string, opts ...util.Map) Responder {
	m := util.CombineMaps(util.Map{
		"partner_trade_no": s,
	}, opts...)
	return obj.SafeRequest(mmpaysptransQueryBank, m)
}

/*TransferToBankCard 转账至银行卡
接口介绍
业务流程	接口	简介
付款	企业付款到银行卡	用于企业向微信用户银行卡付款
目前支持接口API的方式向指定微信用户的银行卡付款。
接口调用规则:
◆ 单商户日限额——单日100w
◆ 单次限额——单次5w
◆ 单商户给同一银行卡单日限额——单日5w
接口地址
接口链接:https://api.mch.weixin.qq.com/mmpaysptrans/pay_bank
是否需要证书
请求需要双向证书。 详见证书使用
请求参数
字段名	变量名	必填	示例值	类型	描述
商户号	mch_id	是	1900000109	string(32)	微信支付分配的商户号
商户企业付款单号	partner_trade_no	是	1212121221227	string(32)	商户订单号，需保持唯一（只允许数字[0~9]或字母[A~Z]和[a~z]，最短8位，最长32位）
随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67Vs	string(32)	随机字符串，不长于32位
签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	string(32)	通过MD5签名算法计算得出的签名值，详见MD5签名生成算法
收款方银行卡号	enc_bank_no	是	8609cb22e1774a50a930e414cc71eca06121bcd266335cda230d24a7886a8d9f	string(64)	收款方银行卡号（采用标准RSA算法，公钥由微信侧提供）,详见获取RSA加密公钥API
收款方用户名	enc_true_name	是	ca775af5f841bdf424b2e6eb86a6e21e	string(64)	收款方用户名（采用标准RSA算法，公钥由微信侧提供）详见获取RSA加密公钥API
收款方开户行	bank_code	是	1001	string(64)	银行卡所在开户行编号,详见银行编号列表
付款金额	amount	是	100000	int	付款金额:RMB分（支付总额，不含手续费）
注:大于0的整数
付款说明	desc	否	理财	string	企业付款到银行卡付款说明,即订单备注（UTF8编码，允许100个字符以内）
返回参数
字段名	变量名	必填	类型	描述
返回状态码	return_code	是	String(16)	SUCCESS/FAIL
此字段是通信标识，非付款标识，付款是否成功需要查看result_code来判断
返回信息	return_msg	否	String(128)	返回信息，如非空，为错误原因
签名失败
参数格式校验错误
以下字段在return_code为SUCCESS的时候有返回
字段名	变量名	必填	类型	描述
业务结果	result_code	是	string(32)	SUCCESS/FAIL，注意:当状态为FAIL时，存在业务结果未明确的情况，所以如果状态为FAIL，请务必通过查询接口确认此次付款的结果（关注错误码err_code字段）。如果要继续进行这笔付款，请务必用原商户订单号和原参数来重入此接口。
错误代码	err_code
否
string(32)
错误码信息，注意:出现未明确的错误码时，如（SYSTEMERROR）等，请务必用原商户订单号重试，或通过查询接口确认此次付款的结果
错误代码描述	err_code_des	否	string(128)	错误信息描述
商户号	mch_id	是	string(32)	微信支付分配的商户号
商户企业付款单号	partner_trade_no	是	string(32)	商户订单号，需要保持唯一
代付金额	amount	是	int	代付金额RMB:分
随机字符串	nonce_str	是	string(32)	随机字符串，长度小于32位
签名	sign	是	string(32)	返回包携带签名给商户
以下字段在return_code 和result_code都为SUCCESS的时候有返回
字段名	变量名	必填	类型	描述
微信企业付款单号	payment_no	是	string(64)	代付成功后，返回的内部业务单号
手续费金额	cmms_amt	是	int	手续费金额 RMB:分
*/
func (obj *Payment) TransferToBankCard(p util.Map, opts ...util.Map) Responder {
	util.CombineMaps(p, opts...)
	if v := p.Check("bank_code", "partner_trade_no", "enc_bank_no", "enc_true_name", "amount"); v != -1 {
		return ErrResponder(xerrors.Errorf("the %d index of value is required", v))
	}

	//p.Set("nonce_str", util.GenerateUUID())

	c := cipher.New(cipher.RSA, cipher.OptionPublic(obj.PublicKey()), cipher.OptionPrivate(obj.PrivateKey()))

	ebn, e := c.Encrypt(p.Get("enc_bank_no"))
	if e != nil {
		return ErrResponder(e)
	}
	p.Set("enc_bank_no", string(ebn))

	etn, e := c.Encrypt(p.Get("enc_true_name"))
	if e != nil {
		return ErrResponder(e)
	}
	p.Set("enc_true_name", string(etn))
	//p.Set("sign", util.GenSign(p, obj.Key))
	return obj.SafeRequest(mmpaysptransPayBank, p)
}

// Request 默认请求
func (obj *Payment) Request(url string, p util.Map) Responder {
	return obj.SafeClient().Post(context.Background(), obj.RequestURL(url), nil, obj.initPay(p))
}

// SafeRequest 安全请求
func (obj *Payment) SafeRequest(url string, p util.Map) Responder {
	return obj.SafeClient().Post(context.Background(), obj.RequestURL(url), nil, obj.initPay(p))
}

func (obj *Payment) initPay(p util.Map, ignore ...string) util.Map {
	if !p.Has("mch_appid") || !p.Has("appid") {
		p.Set("appid", obj.AppID)
	}

	if !p.Has("mch_id") && !p.Has("mchid") {
		p.Set("mch_id", obj.MchID)
	}
	p.Set("nonce_str", util.GenerateUUID())
	if obj.subMchID != "" {
		p.Set("sub_mch_id", obj.subMchID)
	}
	if obj.subAppID != "" {
		p.Set("sub_appid", obj.subAppID)
	}

	if !p.Has("sign") {
		p.Set("sign", util.GenSign(p, obj.GetKey(), ignore...))
	}
	log.Debug("initPay end", p)
	return p
}

// UseSandbox ...
func (obj *Payment) UseSandbox() bool {
	return obj.sandbox != nil
}

/*GetKey 沙箱key(string类型) */
func (obj *Payment) GetKey() string {
	key := obj.Key
	if obj.UseSandbox() {
		keyName := obj.sandbox.getCacheKey()
		cachedKey := cache.Get(keyName)
		if cachedKey != nil {
			log.Debug("cached key:", keyName, cachedKey.(string))
			key = cachedKey.(string)
			return key
		}

		resp := obj.sandbox.SignKey().ToMap()
		if resp.GetString("return_code") == "SUCCESS" {
			key = resp.GetString("sandbox_signkey")
			log.Debug("key:", keyName, key)
			cache.SetWithTTL(keyName, key, 24*3600)
		}
	}

	if 32 != len(key) {
		log.Error(fmt.Sprintf("%s should be 32 chars length.", key))
		return ""
	}
	return key

}

// NotifyURL ...
func (obj *Payment) NotifyURL() string {
	return util.URL(obj.LocalHost(), paymentNotifyURL(obj))
}
func paymentNotifyURL(obj *Payment) string {
	if obj != nil && obj.notifyURL != "" {
		return obj.notifyURL
	}
	return notifyCB
}

// RefundURL ...
func (obj *Payment) RefundURL() string {
	return util.URL(obj.LocalHost(), paymentRefundURL(obj))
}

func paymentRefundURL(obj *Payment) string {
	if obj != nil && obj.refundedURL != "" {
		return obj.refundedURL
	}
	return refundedCB
}

// ScannedURL ...
func (obj *Payment) ScannedURL() string {
	return util.URL(obj.LocalHost(), paymentScannedURL(obj))
}

func paymentScannedURL(obj *Payment) string {
	if obj != nil && obj.scannedURL != "" {
		return obj.scannedURL
	}
	return scannedCB
}

// RequestURL ...
func (obj *Payment) RequestURL(uri string) string {
	if obj.UseSandbox() {
		return util.URL(obj.RemoteURL(), sandboxNew, uri)
	}
	return util.URL(obj.RemoteURL(), uri)
}

// RemoteURL ...
func (obj *Payment) RemoteURL() string {
	if obj == nil || obj.remoteURL == "" {
		return APIMCHDefault
	}
	return obj.remoteURL
}

// LocalHost ...
func (obj *Payment) LocalHost() string {
	return local(obj)
}

func local(obj *Payment) string {
	if obj != nil && obj.localHost != "" {
		return obj.localHost
	}
	return wegoLocal
}

// SubAppID ...
func (obj *Payment) SubAppID() string {
	return obj.subAppID
}

// SubMchID ...
func (obj *Payment) SubMchID() string {
	return obj.subMchID
}

// PrivateKey ...
func (obj *Payment) PrivateKey() string {
	return obj.privateKey
}

// PublicKey ...
func (obj *Payment) PublicKey() string {
	return obj.publicKey
}
