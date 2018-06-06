package core

const MpDomain = "https://mp.weixin.qq.com"
const BaseDomain = "https://api.mch.weixin.qq.com"
const ApiWeixin = "https://api.weixin.qq.com"
const BackDomain = "https://api2.mch.weixin.qq.com"
const HkDomain = "https://apihk.mch.weixin.qq.com"
const UsDomain = "https://apius.mch.weixin.qq.com"
const BizPayURL = "weixin://wxpay/bizpayurl?"
const FileApiWeixin = "http://file.api.weixin.qq.com"

const FAIL = "FAIL"
const SUCCESS = "SUCCESS"
const HMACSHA256 = "HMAC-SHA256"
const MD5 = "MD5"

const FieldSign = "sign"
const FieldSignType = "sign_type"

const SYSTEMERROR = "SYSTEMERROR"
const BANKERROR = "BANKERROR"
const USERPAYING = "USERPAYING"

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
const getcallbackipURLSuffix = "/cgi-bin/getcallbackip"

const sandboxUrlSuffix = "/sandboxnew"
const sandboxSignkeyUrlSuffix = sandboxUrlSuffix + "/pay/getsignkey"
