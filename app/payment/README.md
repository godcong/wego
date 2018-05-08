微信授权：
取得OpenId:
official_account.NewOauth()
token := oauth.AccessToken(#code#)
oauth.UserInfo(token)

也可绑定ServeHTTP到任何http Server:
注册回调监听
	oauth.RegisterAllCallback(func(w http.ResponseWriter, r *http.Request, val *official_account.CallbackValue) []byte {
		if val.Type == "info" {
			info := val.Value.(*core.UserInfo)
			log.Println("save info", *info)
		}
		return nil
	})
注册回调地址
	http.HandleFunc("/oauth_callback", oauth.ServeHTTP)


支付请求：

数据初始化：

data := make(util.Map)
data.Set("body", "腾讯充值中心-QQ会员充值")
data.Set("out_trade_no", out_trade_no)
data.Set("device_info", "")
data.Set("fee_type", "CNY")
data.Set("total_fee", "1")
data.Set("spbill_create_ip", "123.12.12.123")
data.Set("notify_url", "http://test.letiantian.me/wxpay/notify")
data.Set("trade_type", "NATIVE")
data.Set("product_id", "12")
//或者直接初始化

data := util.Map{
"body":"腾讯充值中心-QQ会员充值",
...,
}

a. 统一下单： https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1

调用接口： GetOrder().Unify(data)

b. 查询订单: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2

调用接口： GetOrder().Query(data)

c. 关闭订单: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3

调用接口： GetOrder().Close(data)

d. 申请退款: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4

调用接口： GetRefund().ByOutTradeNumber()
调用接口： GetRefund().ByTransactionId()

e. 查询退款: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5

调用接口： GetRefund().QueryByRefundId(id)
调用接口： GetRefund().QueryByOutRefundNumber(num)
调用接口： GetRefund().QueryByOutTradeNumber(num)
调用接口： GetRefund().QueryByTransactionId(id)

f. 下载对账单 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6

调用接口： GetBill().Get(bill_date, bill_type,op Map)

g. 支付结果通知 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_7

h. 交易保障 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_8

i. 退款结果通知 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=9

j. 拉取订单评价数据 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_17&index=10
