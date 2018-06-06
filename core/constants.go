package core

/* domain defines */
const (
	MPDomain      = "https://mp.weixin.qq.com"
	BaseDomain    = "https://api.mch.weixin.qq.com"
	APIWeixin     = "https://api.weixin.qq.com"
	API2Domain    = "https://api2.mch.weixin.qq.com"
	HKDomain      = "https://apihk.mch.weixin.qq.com"
	USDomain      = "https://apius.mch.weixin.qq.com"
	BizPayURL     = "weixin://wxpay/bizpayurl?"
	FileAPIWeixin = "http://file.api.weixin.qq.com"
)

/*HMACSHA256 定义:HMAC-SHA256 */
const HMACSHA256 = "HMAC-SHA256"

/*MD5 定义:MD5 */
const MD5 = "MD5"

/*FieldSign 定义:sign */
const FieldSign = "sign"

/*FieldSignType 定义:sign_type */
const FieldSignType = "sign_type"

//const SYSTEMERROR = "SYSTEMERROR"
//const BANKERROR = "BANKERROR"
//const USERPAYING = "USERPAYING"

//const SSLCERT_PATH = "./cert/apiclient_cert.pem"
//const SSLKEY_PATH = "./cert/apiclient_key.pem"

const reportURLSuffix = "/payitil/report"
const shortURLSuffix = "/tool/shorturl"

const tokenURLSuffix = "/cgi-bin/token"

//const CUSTOM_SEND_URL_SUFFIX = "/cgi-bin/message/custom/send"

const getonlinekflistURLSuffix = "/cgi-bin/customservice/getonlinekflist"
const kfaccountAddURLSuffix = "/customservice/kfaccount/add"
const kfaccountUpdateURLSuffix = "/customservice/kfaccount/update"
const kfaccountDelURLSuffix = "/customservice/kfaccount/del"
const kfaccountInviteworkerURLSuffix = "/customservice/kfaccount/inviteworker"
const kfaccountUploadheadimgURLSuffix = "/customservice/kfaccount/uploadheadimg"
const msgrecordGetmsglistURLSuffix = "/customservice/msgrecord/getmsglist"

const clearQuotaURLSuffix = "/cgi-bin/clear_quota"
const getCallbackIPURLSuffix = "/cgi-bin/getcallbackip"

const sandboxURLSuffix = "/sandboxnew"
const sandboxSignKeyURLSuffix = sandboxURLSuffix + "/pay/getsignkey"
