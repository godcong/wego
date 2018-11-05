package payment

const domain = "https://api.mch.weixin.qq.com"

const riskGetPublicKey = "https://fraud.mch.weixin.qq.com/risk/getpublickey"

const payDownloadBill = "/pay/downloadbill"
const payUnifiedOrder = "/pay/unifiedorder"
const payOrderQuery = "/pay/orderquery"
const payMicroPay = "/pay/micropay"
const payCloseOrder = "/pay/closeorder"
const payRefundQuery = "/pay/refundquery"

const payReverse = "/secapi/pay/reverse"
const payRefund = "/secapi/pay/refund"
const mchSubmchmanage = "/secapi/mch/submchmanage"
const mchModifymchinfo = "secapi/mch/modifymchinfo"
const mktAddrecommendconf = "secapi/mkt/addrecommendconf"

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

/*FieldSign 定义:sign */
const FieldSign = "sign"

/*FieldSignType 定义:sign_type */
const FieldSignType = "sign_type"

/*HMACSHA256 定义:HMAC-SHA256 */
const HMACSHA256 = "HMAC-SHA256"

/*MD5 定义:MD5 */
const MD5 = "MD5"
