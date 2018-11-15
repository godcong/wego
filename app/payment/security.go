package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Security Security */
type Security struct {
	*Payment
}

func newSecurity(pay *Payment) interface{} {
	return &Security{
		Payment: pay,
	}
}

// NewSecurity ...
func NewSecurity(config *core.Config) *Security {
	return newSecurity(NewPayment(config)).(*Security)
}

/*GetPublicKey 获取RSA加密公钥API
接口说明
请求Url	https://fraud.mch.weixin.qq.com/risk/getpublickey
是否需要证书	请求需要双向证书。 详见证书使用
请求方式	POST

PS: 可使用ToFile保存Key.需转换成PKCS#8使用.
RSA公钥格式PKCS#1,PKCS#8互转说明
PKCS#1 转 PKCS#8:
openssl rsa -RSAPublicKey_in -in <filename> -pubout
PKCS#8 转 PKCS#1:
openssl rsa -pubin -in <filename> -RSAPublicKey_out
*/
func (s *Security) GetPublicKey() core.Response {
	m := util.Map{"sign_type": "MD5"}
	maps := util.Map{
		core.DataTypeXML:      s.initRequest(m),
		core.DataTypeSecurity: s.Config,
	}
	return s.client.Request(riskGetPublicKey, "post", maps)
}
