package core

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

const REPORT_URL_SUFFIX = "/payitil/report"
const SHORTURL_URL_SUFFIX = "/tools/shorturl"

const SANDBOX_URL_SUFFIX = "/sandboxnew"
const SANDBOX_SIGNKEY_URL_SUFFIX = SANDBOX_URL_SUFFIX + "/pay/getsignkey"

const API_WEIXIN_URL_SUFFIX = "https://api.weixin.qq.com"

const GETWXACODE_URL_SUFFIX = "/wxa/getwxacode"
const CREATEWXAQRCODE_URL_SUFFIX = "/cgi-bin/wxaapp/createwxaqrcode"
const GETWXACODEUNLIMIT_URL_SUFFIX = "/wxa/getwxacodeunlimit"
const SNS_JSCODE2SESSION_URL_SUFFIX = "/sns/jscode2session"
const CGI_BIN_TOKEN_URL_SUFFIX = "/cgi-bin/token"
const DATACUBE_VISITPAGE_URL_SUFFIX = "/datacube/getweanalysisappidvisitpage"
const DATACUBE_USERPORTRAIT_URL_SUFFIX = "/datacube/getweanalysisappiduserportrait"
const DATACUBE_MONTHLYRETAININFO_URL_SUFFIX = "/datacube/getweanalysisappidmonthlyretaininfo"
const DATACUBE_WEEKLYRETAININFO_URL_SUFFIX = "/datacube/getweanalysisappidweeklyretaininfo"
const DATACUBE_DAILYRETAININFO_URL_SUFFIX = "/datacube/getweanalysisappiddailyretaininfo"
const DATACUBE_VISITDISTRIBUTION_URL_SUFFIX = "/datacube/getweanalysisappidvisitdistribution"
const DATACUBE_MONTHLYVISITTREND_URL_SUFFIX = "/datacube/getweanalysisappidmonthlyvisittrend"
const DATACUBE_WEEKLYVISITTREND_URL_SUFFIX = "/datacube/getweanalysisappidweeklyvisittrend"
const DATACUBE_DAILYVISITTREND_URL_SUFFIX = "/datacube/getweanalysisappiddailyvisittrend"
const DATACUBE_DAILYSUMMARYTREND_URL_SUFFIX = "/datacube/getweanalysisappiddailysummarytrend"

const TEMPLATE_ADD_URL_SUFFIX = "/cgi-bin/wxopen/template/add"
const TEMPLATE_DEL_URL_SUFFIX = "/cgi-bin/wxopen/template/del"
const TEMPLATE_LIST_URL_SUFFIX = "/cgi-bin/wxopen/template/list"
const TEMPLATE_SEND_URL_SUFFIX = "/cgi-bin/message/wxopen/template/send"
const TEMPLATE_LIBRARY_LIST_URL_SUFFIX = "/cgi-bin/wxopen/template/library/list"
const TEMPLATE_LIBRARY_GET_URL_SUFFIX = "/cgi-bin/wxopen/template/library/get"

const CLEAR_QUOTA_URL_SUFFIX = "/cgi-bin/clear_quota"
const GETCALLBACKIP_URL_SUFFIX = "/cgi-bin/getcallbackip"

const CUSTOM_SEND_URL_SUFFIX = "/cgi-bin/message/custom/send"
const GETKFLIST_URL_SUFFIX = "/cgi-bin/customservice/getkflist"
const GETONLINEKFLIST_URL_SUFFIX = "/cgi-bin/customservice/getonlinekflist"
const KFACCOUNT_ADD_URL_SUFFIX = "/customservice/kfaccount/add"
const KFACCOUNT_UPDATE_URL_SUFFIX = "/customservice/kfaccount/update"
const KFACCOUNT_DEL_URL_SUFFIX = "/customservice/kfaccount/del"
const KFACCOUNT_INVITEWORKER_URL_SUFFIX = "/customservice/kfaccount/inviteworker"
const KFACCOUNT_UPLOADHEADIMG_URL_SUFFIX = "/customservice/kfaccount/uploadheadimg"
const MSGRECORD_GETMSGLIST_URL_SUFFIX = "/customservice/msgrecord/getmsglist"
