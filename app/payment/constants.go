package payment

const domain = "https://api.mch.weixin.qq.com"

const riskGetPublicKey = "https://fraud.mch.weixin.qq.com/risk/getpublickey"

const batchQueryComment = "/billcommentsp/batchquerycomment"
const payDownloadBill = "/pay/downloadbill"
const payDownloadfundflow = "/pay/downloadfundflow"
const paySettlementquery = "/pay/settlementquery"
const payQueryexchagerate = "pay/queryexchagerate"
const payUnifiedOrder = "/pay/unifiedorder"
const payOrderQuery = "/pay/orderquery"
const payMicroPay = "/pay/micropay"
const payCloseOrder = "/pay/closeorder"
const payRefundQuery = "/pay/refundquery"

const payReverse = "/secapi/pay/reverse"
const payRefund = "/secapi/pay/refund"

const mchSubmchmanage = "/secapi/mch/submchmanage"
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

const authCodeToOpenidURLSuffix = "/tools/authcodetoopenid"

//bizPayURL biz pay url suffix
const bizPayURL = "weixin://wxpay/bizpayurl?"
const sandboxURLSuffix = "/sandboxnew"
const sandboxSignKeyURLSuffix = sandboxURLSuffix + "/pay/getsignkey"

// FieldSign ...
const FieldSign = "sign"

// FieldSignType ...
const FieldSignType = "sign_type"

// FieldLimit ...
const FieldLimit = "limit"

/*HMACSHA256 定义:HMAC-SHA256 */
const HMACSHA256 = "HMAC-SHA256"

/*MD5 定义:MD5 */
const MD5 = "MD5"
