package core

/* domain defines */
const (
	MPDomain   = "https://mp.weixin.qq.com"
	BaseDomain = "https://api.mch.weixin.qq.com"
	APIWeixin  = "https://api.weixin.qq.com"
	API2Domain = "https://api2.mch.weixin.qq.com"
	HKDomain   = "https://apihk.mch.weixin.qq.com"
	USDomain   = "https://apius.mch.weixin.qq.com"

	FileAPIWeixin = "http://file.api.weixin.qq.com"
)

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
