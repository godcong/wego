# 支付 #

[微信支付开发文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)

## 配置
            cfg := C(util.Map{
                "app_id":"wx1ad61aeexxxxxxx",
                   app_id ='wx1ad61aeexxxxxxx'                //AppId
                            mch_id = '1498xxxxx32'                        //商户ID
                            key = 'O9aVVkxxxxxxxxxxxxxxxbZ2NQSJ'    //支付key
                            
                            notify_url ='https://host.address/uri' //支付回调地址
                            
                            //如需使用敏感接口（如退款、发送红包等）需要配置 API 证书路径(登录商户平台下载 API 证书)
                            cert_path = 'cert/apiclient_cert.pem'   //支付证书地址
                            key_path = 'cert/apiclient_key.pem'      //支付证书地址
                            
                            //银行转账功能
                            rootca_path = 'cert/rootca.pem'     //(可不填)
                            pubkey_path = "cert/publickey.pem"  //(可不填)部分支付使用（如:银行转账）
                            prikey_path = "cert/privatekey.pem" //(可不填)部分支付使用（如:银行转账）
            }


           通过配置config.toml文件
      
            //必要配置
            app_id ='wx1ad61aeexxxxxxx'                //AppId
            mch_id = '1498xxxxx32'                        //商户ID
            key = 'O9aVVkxxxxxxxxxxxxxxxbZ2NQSJ'    //支付key
            
            notify_url ='https://host.address/uri' //支付回调地址
            
            //如需使用敏感接口（如退款、发送红包等）需要配置 API 证书路径(登录商户平台下载 API 证书)
            cert_path = 'cert/apiclient_cert.pem'   //支付证书地址
            key_path = 'cert/apiclient_key.pem'      //支付证书地址
            
            //银行转账功能
            rootca_path = 'cert/rootca.pem'     //(可不填)
            pubkey_path = "cert/publickey.pem"  //(可不填)部分支付使用（如:银行转账）
            prikey_path = "cert/privatekey.pem" //(可不填)部分支付使用（如:银行转账）
    


## 创建支付对象
    obj:=wego.Payment()
    
## 通过授权码查询公众号Openid 
    obj.AuthCodeToOpenid(#authCode#)
    
## 订单    
### 统一下单
    官方文档: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
    m := make(util.Map)
	m.Set("body", "腾讯充值中心-QQ会员充值")
	m.Set("out_trade_no", 20181024123456)
	m.Set("total_fee", "666")
	m.Set("spbill_create_ip", "10.25.5.141") //可选，如不传该参数，SDK 将会自动获取IP 地址,可通过SetRequest方法帮助获取客户IP
	m.Set("notify_url", "https://test.letiantian.me/wxpay/notify") //支付结果通知，如果不设置则会使用config配置地址
	m.Set("trade_type", "NATIVE") //支付方式对应的值类型(JSAPI，NATIVE，APP)

    obj.Order().Unify(m)
    或
    obj.Order().Unify(util.Map{#请求参数#})

    resuilt:
    {"appid":"wx426b3015555a46be","code_url":"weixin://wxpay/bizpayurl?pr=D3sNT8y","mch_id":"1900009851","nonce_str":"FRFByNNdrzRuEGkp","prepay_id":"wx20180220113507842dff20340601189342","result_code":"SUCCESS","return_code":"SUCCESS","return_msg":"OK","sign":"D398DA0644A14D0BC00A8B82D8D4ECDC","trade_type":"NATIVE"}

### 订单查询
    官方文档: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
    通过商户系统内部的订单号(out_trade_no)查找退款订单
    obj.Order().QueryByOutTradeNumber(#out_trade_no#)
    
    通过微信订单号(transaction_id)查询订单
    obj.Order().QueryByTransactionID(#transaction_id#)
    
### 订单关闭
    官方文档:https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
    通过商户系统内部的订单号out_trade_no关闭订单
    obj.Order().Close(#out_trade_no#)

## 退款

### 发起退款
    官方文档:https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
    按照out_trade_no发起退款
    参数分别为微信订单号、商户退款单号、订单金额、退款金额、其他参数(options以util.Map形式传入)
    obj.Refund().ByOutTradeNumber(tradeNum, num, total, refund, options)

    按照transaction_id发起退款
    参数分别为商户订单号、商户退款单号、订单金额、退款金额、其他参数(options以util.Map形式传入)
    obj.Refund().ByTransactionId(tid, num, total, refund, options)
    
### 查询退款
    官方文档:https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5
    微信订单号
    obj.Refund().QueryByTransactionId(#transactionId#)
    商户订单号
    obj.Refund().QueryByOutTradeNumber(#outTradeNumber#)
    商户退款单号 
    obj.Refund().QueryByOutRefundNumber(#outRefundNumber#)
    微信退款单号 
    obj.Refund().QueryByRefundId(#refundId#)

f. 下载对账单 <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6>
 
g. 支付结果通知 <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_7>

h. 交易保障 <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_8>

i. 退款结果通知 <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=9>

j. 拉取订单评价数据 <https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_17&index=10>

none

