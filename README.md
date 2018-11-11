# WEGO a wechat interface for go(golang)

[![GoDoc](https://godoc.org/github.com/godcong/wego?status.svg)](http://godoc.org/github.com/godcong/wego)
[![license](https://img.shields.io/github/license/godcong/wego.svg)](https://github.com/godcong/wego/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/godcong/wego)](https://goreportcard.com/report/github.com/godcong/wego)

## 世界上最好的Go微信开发SDK(时尚，简单，性能卓越)
### a fashion fast wechat api for golang

### WEGO借鉴了并参考了,当前许多流行的开源微信支付框架.并且为了使性能得到更大提高.全部使用golang进行重建架构开发

开发前,请务必仔细阅读并看懂

获取包:
> go get github.com/godcong/wego

建议使用golang/dep添加包:
> dep ensure --add github.com/godcong/wego

或者vgo添加:
> vgo get github.com/godcong/wego@v0.1.0 

接口说明可以参考godoc
> godoc -http:6060

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

       通过配置config.toml文件，具体参考config.toml.example

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




## Readme

> [公众号](https://github.com/godcong/wego/blob/master/app/official/README.md)

> [小程序](https://github.com/godcong/wego/blob/master/app/mini/README.md)

> [微信支付](https://github.com/godcong/wego/blob/master/app/payment/README.md)

> 开放平台 //TODO:

> 企业微信 //TODO:

> 企业微信开放平台 //TODO:

### 具体功能涵盖，微信模板，企业转账，微信红包，微信支付，微信客服，微信小程序等常用接口。
