# 微信授权

## 第零步，配置config.toml中[payment.default]部分

            app_id ='wx1ad61aeef1903b93'                //AppId
            mch_id = '1498009232'                        //商户ID
            key = 'O9aVVkSTmgJK4qSibhSYpGQzRbZ2NQSJ'    //支付key
            notify_url ='https://mp.quick58.com/charge/callback' //支付回调地址
            cert_path = 'cert/apiclient_cert.pem'   //支付证书地址
            key_path = 'cert/apiclient_key.pem'      //支付证书地址
            rootca_path = 'cert/rootca.pem'     //(可不填)
            pubkey_path = "cert/publickey.pem"  //(可不填)部分支付使用（如：银行转账）
            prikey_path = "cert/privatekey.pem" //(可不填)部分支付使用（如：银行转账）

## 第一步,取得OpenId

>
    //创建微信授权
    oauth:=official_account.NewOauth()
    //生成一个跳转链接,state自定义
    oauth.AuthCodeURL(#state#)

    //监听回调地址取得code和state

    //输入code获取token
    token := oauth.AccessToken(#code#)

    //输入token获取用户信息
    oauth.UserInfo(token)

## 或绑定ServeHTTP到任何http Server

### 注册回调监听

>
    oauth.RegisterAllCallback(func(w http.ResponseWriter, r *http.Request, val *official_account.CallbackValue) []byte {
        if val.Type == "info" {
            info := val.Value.(*core.UserInfo)
            log.Println("save info", *info)
        }
        return nil
    })

### 注册回调地址

>
    http.HandleFunc("/oauth_callback", oauth.ServeHTTP)

## 第二部,支付请求：

数据初始化：
>
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

//也可以直接初始化

>
    data := util.Map{
    "body":"腾讯充值中心-QQ会员充值",
    ...,
    }

## 第三部,调用接口

### 所需参数可参考微信接口定义

before：
创建order对象：order:=payment.NewOrder()
创建refund对象: refund:=payment.NewRefund()
创建bill对象：bill:= payment.NewBill()

a. 统一下单： <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1>

调用接口： order.Unify(#data#)
PS：data参数及返回值请参考微信API

b. 查询订单: <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2>

调用接口： order.QueryByTransactionId(#transaction_id#)
          order.QueryByOutTradeNumber(#out_trade_no#)

c. 关闭订单: <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3>

调用接口： order.Close(#out_trade_no#)

d. 申请退款: <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4>

调用接口： refund.ByOutTradeNumber(#out_trade_no#, #out_refund_no#, #total_fee#, #refund_fee#,#options#)
调用接口： refund.ByTransactionId(#transaction_id#, #out_refund_no#, #total_fee#, #refund_fee#,#options#)

e. 查询退款: <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5>

调用接口： refund.QueryByRefundId(#refund_id#)
调用接口： refund.QueryByOutRefundNumber(#out_refund_no#)
调用接口： refund.QueryByOutTradeNumber(#out_trade_no#)
调用接口： refund.QueryByTransactionId(#transaction_id#)

f. 下载对账单 <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6>

调用接口： bill.Get(#bill_date#, #bill_type#,#options#)

g. 支付结果通知 <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_7>


h. 交易保障 <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_8>

i. 退款结果通知 <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=9>

j. 拉取订单评价数据 <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_17&index=10>
