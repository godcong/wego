package payment

const domain = "https://api.mch.weixin.qq.com"

const riskGetPublicKeyURLSuffix = "https://fraud.mch.weixin.qq.com/risk/getpublickey"
const downloadBillURLSuffix = "/pay/downloadbill"
const unifiedOrderURLSuffix = "/pay/unifiedorder"
const orderQueryURLSuffix = "/pay/orderquery"
const microPayURLSuffix = "/pay/micropay"
const reverseURLSuffix = "/secapi/pay/reverse"
const closeOrderURLSuffix = "/pay/closeorder"
const refundURLSuffix = "/secapi/pay/refund"
const refundQueryURLSuffix = "/pay/refundquery"
const sendRedPackURLSuffix = "/mmpaymkttransfers/sendredpack"
const getHBInfoURLSuffix = "/mmpaymkttransfers/gethbinfo"
const sendGroupRedPackURLSuffix = "/mmpaymkttransfers/sendgroupredpack"
const getTransferInfoURLSuffix = "/mmpaymkttransfers/gettransferinfo"
const promotionTransfersURLSuffix = "/mmpaymkttransfers/promotion/transfers"
const mmPaySpTransQueryBankURLSuffix = "/mmpaysptrans/query_bank"
const mmPaySpTransPayBankURLSuffix = "/mmpaysptrans/pay_bank"

const sendCouponURLSuffix = "/mmpaymkttransfers/send_coupon"
const queryCouponStockURLSuffix = "/mmpaymkttransfers/query_coupon_stock"
const queryCouponsInfoURLSuffix = "/mmpaymkttransfers/querycouponsinfo"

const authCodeToOpenidURLSuffix = "/tool/authcodetoopenid"

//bizPayURL biz pay url suffix
const bizPayURL = "weixin://wxpay/bizpayurl?"
const sandboxURLSuffix = "/sandboxnew"
const sandboxSignKeyURLSuffix = sandboxURLSuffix + "/pay/getsignkey"

/*FieldSign 定义:sign */
const FieldSign = "sign"

/*FieldSignType 定义:sign_type */
const FieldSignType = "sign_type"

/*HMACSHA256 定义:HMAC-SHA256 */
const HMACSHA256 = "HMAC-SHA256"

/*MD5 定义:MD5 */
const MD5 = "MD5"
