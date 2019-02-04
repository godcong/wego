package app

import (
	"crypto/md5"
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

//Payment ...
type Payment struct {
	*PaymentProperty
	property    *Property
	accessToken *AccessToken
	option      *PaymentOption
}

// PaymentOption ...
type PaymentOption struct {
	RemoteHost string
	Sandbox    SandboxProperty
	PaidURL    string
	RefundURL  string
}

// NewPayment ...
func NewPayment(property *Property, opts ...*PaymentOption) *Payment {
	var opt *PaymentOption
	if opts != nil {
		opt = opts[0]
	}
	return &Payment{
		PaymentProperty: property.Payment,
		accessToken:     nil,
		property:        property,
		option:          opt,
	}
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
func (obj *Payment) Pay(maps util.Map) core.Responder {
	maps.Set("appid", obj.AppID)

	//set notify callback
	notify := obj.option.PaidURL
	if !maps.Has("notify_url") {
		maps.Set("notify_url", notify)
	}

	return nil
	//TODO
	//return obj.Request(payMicroPay, maps)
}

// IsSandbox ...
func (obj *Payment) IsSandbox() bool {
	if obj.option != nil {
		return obj.option.Sandbox.UseSandbox
	}
	return false
}

/*GetKey 沙箱key(string类型) */
func (obj *Payment) GetKey() string {
	key := obj.Key
	if obj.IsSandbox() {
		cachedKey := cache.Get(obj.getCacheKey())
		if cachedKey != nil {
			key = cachedKey.(string)
		}

		response := obj.SandboxSignKey().ToMap()
		if response.GetString("return_code") == "SUCCESS" {
			key = response.GetString("sandbox_signkey")
			cache.SetWithTTL(obj.getCacheKey(), key, 3*24*3600)
		}
	}

	if 32 != len(key) {
		log.Error(fmt.Sprintf("%s should be 32 chars length.", key))
		return ""
	}

	return key

}

func (obj *Payment) getCacheKey() string {
	name := obj.option.Sandbox.AppID + "." + obj.option.Sandbox.MchID
	return "godcong.wego.payment.sandbox." + fmt.Sprintf("%x", md5.Sum([]byte(name)))
}

/*SandboxSignKey 沙箱key */
func (obj *Payment) SandboxSignKey() core.Responder {
	m := make(util.Map)
	m.Set("mch_id", obj.option.Sandbox.MchID)
	m.Set("nonce_str", util.GenerateNonceStr())
	sign := util.GenerateSignature(m, obj.option.Sandbox.Key, util.MakeSignMD5)
	m.Set("sign", sign)
	resp := core.PostXML(util.URL(obj.RemoteHost(), sandboxSignKeyURLSuffix), nil, m)

	return resp

}

// RemoteHost ...
func (obj *Payment) RemoteHost() string {
	if obj.option.RemoteHost != "" {
		return obj.option.RemoteHost
	}
	return BaseDomain
}

// LocalAddress ...
func (obj *Payment) LocalAddress() string {
	return obj.property.Local.Address
}

// PaidURL ...
func (obj *Payment) PaidURL() string {
	uri := obj.property.Local.PaidURL
	if obj.option.PaidURL != "" {
		uri = obj.option.PaidURL
	}
	return util.URL(obj.LocalAddress(), uri)
}

// RefundURL ...
func (obj *Payment) RefundURL() string {
	uri := obj.property.Local.RefundedURL
	if obj.option.RefundURL != "" {
		uri = obj.option.RefundURL
	}
	return util.URL(obj.LocalAddress(), uri)
}
