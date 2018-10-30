package payment_test

import (
	"github.com/godcong/wego/app/payment"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"testing"
)

func TestNewPayment(t *testing.T) {
	cfg := core.C(util.Map{
		"app_id": "wx1ad61aeexxxxxxx",
		"mch_id": "1498xxxxx32",                  //商户ID
		"key":    "O9aVVkxxxxxxxxxxxxxxxbZ2NQSJ", //支付key

		"notify_url": "https://host.address/uri", //支付回调地址

		//如需使用敏感接口（如退款、发送红包等）需要配置 API 证书路径(登录商户平台下载 API 证书)
		"cert_path": "cert/apiclient_cert.pem", //支付证书地址
		"key_path":  "cert/apiclient_key.pem",  //支付证书地址

		//银行转账功能
		"rootca_path": "cert/rootca.pem",     //(可不填)
		"pubkey_path": "cert/publickey.pem",  //(可不填)部分支付使用（如：银行转账）
		"prikey_path": "cert/privatekey.pem", //(可不填)部分支付使用（如：银行转账）
	})

	payment.NewPayment(cfg)

}
