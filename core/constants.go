package core

const MP_DOMAIN = "https://mp.weixin.qq.com"
const BASE_DOMAIN = "https://api.mch.weixin.qq.com"
const API_WEIXIN = "https://api.weixin.qq.com"
const BACK_DOMAIN = "https://api2.mch.weixin.qq.com"
const HK_DOMAIN = "https://apihk.mch.weixin.qq.com"
const US_DOMAIN = "https://apius.mch.weixin.qq.com"
const BIZPAYURL = "weixin://wxpay/bizpayurl?"
const FILE_API_WEIXIN = "http://file.api.weixin.qq.com"

const FAIL = "FAIL"
const SUCCESS = "SUCCESS"
const HMACSHA256 = "HMAC-SHA256"
const MD5 = "MD5"

const FIELD_SIGN = "sign"
const FIELD_SIGN_TYPE = "sign_type"

const SYSTEMERROR = "SYSTEMERROR"
const BANKERROR = "BANKERROR"
const USERPAYING = "USERPAYING"

//const SSLCERT_PATH = "./cert/apiclient_cert.pem"
//const SSLKEY_PATH = "./cert/apiclient_key.pem"

const REPORT_URL_SUFFIX = "/payitil/report"
const SHORTURL_URL_SUFFIX = "/tool/shorturl"

const SANDBOX_URL_SUFFIX = "/sandboxnew"
const SANDBOX_SIGNKEY_URL_SUFFIX = SANDBOX_URL_SUFFIX + "/pay/getsignkey"

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

const TEMPLATE_LIBRARY_LIST_URL_SUFFIX = "/cgi-bin/wxopen/template/library/list"
const TEMPLATE_LIBRARY_GET_URL_SUFFIX = "/cgi-bin/wxopen/template/library/get"

const CUSTOM_SEND_URL_SUFFIX = "/cgi-bin/message/custom/send"

const GETONLINEKFLIST_URL_SUFFIX = "/cgi-bin/customservice/getonlinekflist"
const KFACCOUNT_ADD_URL_SUFFIX = "/customservice/kfaccount/add"
const KFACCOUNT_UPDATE_URL_SUFFIX = "/customservice/kfaccount/update"
const KFACCOUNT_DEL_URL_SUFFIX = "/customservice/kfaccount/del"
const KFACCOUNT_INVITEWORKER_URL_SUFFIX = "/customservice/kfaccount/inviteworker"
const KFACCOUNT_UPLOADHEADIMG_URL_SUFFIX = "/customservice/kfaccount/uploadheadimg"
const MSGRECORD_GETMSGLIST_URL_SUFFIX = "/customservice/msgrecord/getmsglist"
