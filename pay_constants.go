package wego

type SignType int

const (
	SIGN_TYPE_MD5        SignType = iota
	SIGN_TYPE_HMACSHA256 SignType = iota
)

func (t SignType) String() string {
	if t == SIGN_TYPE_HMACSHA256 {
		return HMACSHA256
	}
	return MD5
}

const BASE_DOMAIN = "https://api.mch.weixin.qq.com"
const BACK_DOMAIN = "api2.mch.weixin.qq.com"
const HK_DOMAIN = "apihk.mch.weixin.qq.com"
const US_DOMAIN = "apius.mch.weixin.qq.com"
const BIZPAYURL = "weixin://wxpay/bizpayurl?"

const FAIL = "FAIL"
const SUCCESS = "SUCCESS"
const HMACSHA256 = "HMAC-SHA256"
const MD5 = "MD5"

const SYSTEMERROR = "SYSTEMERROR"
const BANKERROR = "BANKERROR"
const USERPAYING = "USERPAYING"

const FIELD_SIGN = "sign"
const FIELD_SIGN_TYPE = "sign_type"

//const SSLCERT_PATH = "./cert/apiclient_cert.pem"
//const SSLKEY_PATH = "./cert/apiclient_key.pem"

const MICROPAY_URL_SUFFIX = "/pay/micropay"
const UNIFIEDORDER_URL_SUFFIX = "/pay/unifiedorder"
const ORDERQUERY_URL_SUFFIX = "/pay/orderquery"
const REVERSE_URL_SUFFIX = "/secapi/pay/reverse"
const CLOSEORDER_URL_SUFFIX = "/pay/closeorder"
const REFUND_URL_SUFFIX = "/secapi/pay/refund"
const REFUNDQUERY_URL_SUFFIX = "/pay/refundquery"
const DOWNLOADBILL_URL_SUFFIX = "/pay/downloadbill"
const REPORT_URL_SUFFIX = "/payitil/report"
const SHORTURL_URL_SUFFIX = "/tools/shorturl"
const AUTHCODETOOPENID_URL_SUFFIX = "/tools/authcodetoopenid"

const SANDBOX_URL_SUFFIX = "/sandboxnew"
const SANDBOX_SIGNKEY_URL_SUFFIX = SANDBOX_URL_SUFFIX + "/pay/getsignkey"

const SENDREDPACK_URL_SUFFIX = "/mmpaymkttransfers/sendredpack"
const GETHBINFO_URL_SUFFIX = "/mmpaymkttransfers/gethbinfo"
const SENDGROUPREDPACK_URL_SUFFIX = "/mmpaymkttransfers/sendgroupredpack"
const GETTRANSFERINFO_URL_SUFFIX = "/mmpaymkttransfers/gettransferinfo"
const PROMOTION_TRANSFERS_URL_SUFFIX = "/mmpaymkttransfers/promotion/transfers"
const MMPAYSPTRANS_QUERY_BANK_URL_SUFFIX = "/mmpaymkttransfers/mmpaysptrans/query_bank"
const MMPAYSPTRANS_PAY_BANK_URL_SUFFIX = "/mmpaymkttransfers/mmpaysptrans/pay_bank"

const RISK_GETPUBLICKEY_URL_SUFFIX = "https://fraud.mch.weixin.qq.com/risk/getpublickey"
