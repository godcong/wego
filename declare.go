package wego

const clearQuota = "/cgi-bin/clear_quota"
const getCallbackIP = "/cgi-bin/getcallbackip"
const sandboxNew = "sandboxnew"
const getSignKey = "pay/getsignkey"

// APIMCHUS ...
const APIMCHUS = "https://apius.mch.weixin.qq.com"

// APIMCHHK ...
const APIMCHHK = "https://apihk.mch.weixin.qq.com"

// APIMCHDefault ...
const APIMCHDefault = "https://api.mch.weixin.qq.com"

const apiWeixin = "https://api.weixin.qq.com"
const oauth2Authorize = "https://open.weixin.qq.com/connect/oauth2/authorize"
const oauth2AccessToken = "https://api.weixin.qq.com/sns/oauth2/access_token"
const snsUserinfo = "https://api.weixin.qq.com/sns/userinfo"

const riskGetPublicKey = "https://fraud.mch.weixin.qq.com/risk/getpublickey"

const mchSubMchManage = "/secapi/mch/submchmanage"
const mchModifymchinfo = "/secapi/mch/modifymchinfo"
const mktAddrecommendconf = "/secapi/mkt/addrecommendconf"
const mchAddSubDevConfig = "/secapi/mch/addsubdevconfig"

const mmpaymkttransfersSendRedPack = "/mmpaymkttransfers/sendredpack"
const mmpaymkttransfersGetHbInfo = "/mmpaymkttransfers/gethbinfo"
const mmpaymkttransfersSendGroupRedPack = "/mmpaymkttransfers/sendgroupredpack"
const mmpaymkttransfersGetTransferInfo = "/mmpaymkttransfers/gettransferinfo"
const mmpaymkttransfersPromotionTransfers = "/mmpaymkttransfers/promotion/transfers"

const mmpaymkttransfersSendCoupon = "/mmpaymkttransfers/send_coupon"
const mmpaymkttransfersQueryCouponStock = "/mmpaymkttransfers/query_coupon_stock"
const mmpaymkttransfersQueryCouponsInfo = "/mmpaymkttransfers/querycouponsinfo"

const mmpaysptransQueryBank = "/mmpaysptrans/query_bank"
const mmpaysptransPayBank = "/mmpaysptrans/pay_bank"

const sandboxURLSuffix = "/sandboxnew"
const sandboxSignKeyURLSuffix = sandboxURLSuffix + "/pay/getsignkey"

// BizPayURL ...
const bizPayURL = "weixin://wxpay/bizpayurl?"

const authCodeToOpenidURLSuffix = "/tools/authcodetoopenid"
const batchQueryComment = "/billcommentsp/batchquerycomment"
const payDownloadBill = "/pay/downloadbill"
const payDownloadFundFlow = "/pay/downloadfundflow"
const paySettlementquery = "/pay/settlementquery"
const payQueryexchagerate = "pay/queryexchagerate"
const payUnifiedOrder = "/pay/unifiedorder"
const payOrderQuery = "/pay/orderquery"
const payMicroPay = "/pay/micropay"
const payCloseOrder = "/pay/closeorder"
const payRefundQuery = "/pay/refundquery"

const payReverse = "/secapi/pay/reverse"
const payRefund = "/secapi/pay/refund"

//ticketGetTicket api address suffix
const ticketGetTicket = "/cgi-bin/ticket/getticket"

const wegoLocal = "http://localhost"
const notifyCB = "notify_cb"
const refundedCB = "refunded_cb"
const scannedCB = "scanned_cb"
const defaultKeepAlive = 30
const defaultTimeout = 30

/*accessTokenKey 键值 */
const accessTokenKey = "access_token"
const accessTokenURLSuffix = "/cgi-bin/token"

// POST ...
const POST = "POST"

// GET ...
const GET = "GET"
