# 支付 #

[微信支付开发文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)

## 配置
    cfg := C(util.Map{
        "app_id":"wx1ad61aeexxxxxxx",                //AppId
        "mch_id":"1498xxxxx32",                        //商户ID
        "key":"O9aVVkxxxxxxxxxxxxxxxbZ2NQSJ",    //支付key
        "notify_url":"https://host.address/uri", //支付回调地址
        
        //如需使用敏感接口（如退款、发送红包等）需要配置 API 证书路径(登录商户平台下载 API 证书)
        "cert_path":"cert/apiclient_cert.pem",   //支付证书地址
        "key_path":"cert/apiclient_key.pem",      //支付证书地址
        
        //银行转账功能
        "rootca_path":"cert/rootca.pem",     //(可不填)
        "pubkey_path":"cert/publickey.pem",  //(可不填)部分支付使用（如:银行转账）
        "prikey_path":"cert/privatekey.pem", //(可不填)部分支付使用（如:银行转账）
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
    obj:=wego.Payment() //使用config.toml配置文件
    或
    obj:=payment.NewPaymen(cfg) //使用自定义配置
    
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

    result:
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

### JSSDK
    官方文档:https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141115
    1.WeixinJSBridge:
    obj.JSSDK().BridgeConfig(#prepay_id#) 返回util.Map,使用Map.ToJson()转成json

    javascript:
    WeixinJSBridge.invoke(
           'getBrandWCPayRequest', <?= $json ?>,
           function(res){
               if(res.err_msg == "get_brand_wcpay_request:ok" ) {
                    // 使用以上方式判断前端返回,微信团队郑重提示：
                    // res.err_msg将在用户支付成功后返回
                    // ok，但并不保证它绝对可靠。
               }
           }
       );

    2.JSSDK: 
    obj.JSSDK().SdkConfig(#prepay_id#) 返回util.Map
    
    javascript:
    wx.chooseWXPay({
        timestamp: <?= $config['timestamp'] ?>,
        nonceStr: '<?= $config['nonceStr'] ?>',
        package: '<?= $config['package'] ?>',
        signType: '<?= $config['signType'] ?>',
        paySign: '<?= $config['paySign'] ?>', // 支付签名
        success: function (res) {
            // 支付成功后的回调函数
        }
    });
    
    3.小程序:
     obj.JSSDK().SdkConfig(#prepay_id#) // 返回util.Map
    
    javascript:
    wx.chooseWXPay({
        timeStamp: <?= $config['timestamp'] ?>, //注意 timeStamp 的格式
        nonceStr: '<?= $config['nonceStr'] ?>',
        package: '<?= $config['package'] ?>',
        signType: '<?= $config['signType'] ?>',
        paySign: '<?= $config['paySign'] ?>', // 支付签名
        success: function (res) {
            // 支付成功后的回调函数
        }
    });

### 下载对账单 
    官方文档:https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6
    obj.Bill().Download(util.Map{})
 
### 支付结果通知
    官方文档:https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_7
    
    支付结果通知
    在用户成功支付后，微信服务器会向该订单中设置的回调URL发起一个 POST 请求  
    而对于用户的退款操作，在退款成功之后也会有一个异步回调通知。
    
    本 SDK 内预置了相关方法，
    只需要实现 NotifyCallback 方法，在其中对自己的业务进行处理
    ScannedCallbackFunction := func(p util.Map) (maps util.Map, e error) {
        // 使用通知里的 "微信支付订单号" 或者 "商户订单号" 去自己的数据库找到订单
        order := dummyQuery(p.GetString("out_trade_no")) //通过out_trade_no查询订单,dummyQuery为查询订单函数
        if order != nil || order.GetString("status") != "paid" { // 如果订单不存在 或者 订单已经支付过了
            return nil, nil // 告诉微信，我已经处理完了，订单没找到，别再通知我了
        }

        if rc := p.GetString("return_code"); rc == "SUCCESS" { // return_code 表示通信状态，不代表支付状态
            // 用户是否支付成功
            if p.GetString("result_code") == "SUCCESS" {
                order.Set("paid", util.Time()) // 更新支付时间为当前时间
                order.Set("status", "paid")

                // 用户支付失败
            } else if rc == "FAIL" {
                order.Set("status", "pay_failed")
            }
        } else {
            return nil, errors.New("失败，请稍后再试")
        }
        dummySave(order)
        return nil, nil
    }
    第一参数仅扫码支付中用到,返回参数:util.Map{"prepay_id":"wxxxxxxxxxx"}
    第二返参异常结果为nil即可.
    
    获取监听函数
    方式1:
    serve1 := wego.Payment().HandlePaidNotify(PaidCallbackFunction).ServeHTTP
    http.HandleFunc("/scanned/callback/address", serve1)
    //方式2:
    http.HandleFunc("/scanned/callback/address2", wego.Payment().HandlePaid(PaidCallbackFunction))
    
    设置监听
    http.ListenAndServe(":8080", nil)

### 交易保障 
    官方文档:https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_8

### 退款结果通知 
    官方文档:https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=9
    //正在更新
    wego.Payment().HandleRefundNotify(PaidCallbackFunction).ServeHTTP

### 拉取订单评价数据 
    官方文档:https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_17&index=10
     obj.Bill().BatchQueryComment(util.Map{})
    
## 微信红包
    obj.RedPack()
    
## 企业付款
### 企业付款到零钱
    官方文档:https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
    obj.Transfer().ToBalance(util.Map{
        "partner_trade_no" : "1233455", // 商户订单号，需保持唯一性(只能是字母或者数字，不能包含有符号)
        "openid" : "oxTWIuGaIt6gTKsQRLau2M0yL16E",
        "check_name" : "FORCE_CHECK", // NO_CHECK：不校验真实姓名, FORCE_CHECK：强校验真实姓名
        "re_user_name" : "王小帅", // 如果 check_name 设置为FORCE_CHECK，则必填用户真实姓名
        "amount" : 10000, // 企业付款金额，单位为分
        "desc" : "理赔", // 企业付款操作说明信息。必填
    }) // 返回core.Response
    
### 查询企业付款到零钱
    官方文档:https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3
    obj.Transfer().QueryBalanceOrder(#partnerTradeNo#) // 商户订单号
    
###  企业付款到银行卡	
     官方文档:https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_2
     
      obj.Security().GetPublicKey()获取公钥保存,并转换成PKCS#8格式.
      转换方法参考官方文档:https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_7&index=4  
      使用前,必须已配置config的pubkey_path
   
     obj.Transfer().ToBankCard(util.Map{
        "partner_trade_no" : "1229222022",
        "enc_bank_no" : "6214830901234564", // 银行卡号
        "enc_true_name" : "聪神",   // 银行卡对应的用户真实姓名
        "bank_code" : "1001", // 银行编号
        "amount" : "100",  // 单位：分
        "desc" : "测试" // 商户订单号   
     })
   
###  查询企业付款到银行卡
    官方文档:https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_3
    obj.Transfer().QueryBankCardOrder(#partnerTradeNo#)
 