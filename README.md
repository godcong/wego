# WEGO a wechat interface for go(golang)

[![GoDoc](https://godoc.org/github.com/godcong/wego?status.svg)](http://godoc.org/github.com/godcong/wego)
[![license](https://img.shields.io/github/license/godcong/wego.svg)](https://github.com/godcong/wego/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/godcong/wego.svg?branch=master)](https://travis-ci.org/godcong/wego)
[![Go Report Card](https://goreportcard.com/badge/github.com/godcong/wego)](https://goreportcard.com/report/github.com/godcong/wego)

## a fashion fast wechat pay api for golang

## 时尚，简单，卓越，高性能的微信支付接口

获取包：
> go get github.com/godcong/wego

建议使用golang/dep添加包：
> dep ensure --add github.com/godcong/wego

或者vgo添加：
> vgo get github.com/godcong/wego@v0.1.0 

配置config.toml：
> 参考config.toml.example

接口说明可以参考godoc
> godoc -http:6060

使用哪个模块就New哪个模块，具体文档在各个模块下。

Wego搭建微信支付模块参考文档：
[支付模块README](https://github.com/godcong/wego/blob/master/app/payment/README.md)

Wego搭建公众号参考文档：
[公众号模块README](https://github.com/godcong/wego/blob/master/app/official/README.md)

Wego搭建小程序后台参考文档：
[小程序模块README](https://github.com/godcong/wego/blob/master/app/mini/README.md)

# PS:因微信接口超过,还有部分接口,因微信权限原因无法调试,故剩下部分进展比较缓慢.
# 如有需要的接口还未开发.请提交issue.


具体功能涵盖，微信模板，企业转账，微信红包，微信支付，微信客服，微信小程序等。