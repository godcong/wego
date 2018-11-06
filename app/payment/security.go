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

/*NewSecurity NewSecurity */
func NewSecurity(config *core.Config) *Security {
	return newSecurity(NewPayment(config)).(*Security)
}

/*GetPublicKey 获取RSA加密公钥API
接口说明
请求Url	https://fraud.mch.weixin.qq.com/risk/getpublickey
是否需要证书	请求需要双向证书。 详见证书使用
请求方式	POST
请求参数
字段名	字段	必填	示例值	类型	说明
商户号	mch_id	是	1900000109	string(32)	微信支付分配的商户号
随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67Vs	string(32)	随机字符串，长度小于32位
签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	string(64)	商户自带签名
签名类型	sign_type	是	MD5	string(64)	签名类型，支持HMAC-SHA256和MD5。
返回参数
字段名	变量名	必填	类型	说明
返回状态码	return_code	是	String(16)
SUCCESS/FAIL
此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
返回信息	return_msg	是	String(128)
返回信息，如非空，为错误原因
签名失败
参数格式校验错误
以下字段在return_code为SUCCESS的时候有返回
业务结果	result_code	是	String(16)	SUCCESS/FAIL
错误代码	err_code	否	String(32)	错误码信息
错误代码描述	err_code_des	否	String(128)	结果信息描述
以下字段在return_code 和result_code都为SUCCESS的时候有返回
商户号	mch_id	是	string(32)	商户号
密钥	pub_key	是	String(2048)	RSA 公钥
错误码
错误代码	描述	解决方案
SIGNERROR	签名错误	签名前没有按照要求进行排序。没有使用商户平台设置的密钥进行签名，参数有空格或者进行了encode后进行签名
SYSTEMERROR	系统繁忙，请稍后重试	使用原请求参数重试
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
