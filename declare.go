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

const sandbox = "/sandboxnew"
const sandboxSignKey = sandbox + "/pay/getsignkey"

// BizPayURL ...
const bizPayURL = "weixin://wxpay/bizpayurl?"

const authCodeToOpenid = "/tools/authcodetoopenid"
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
const accessToken = "/cgi-bin/token"

const getKFList = "/cgi-bin/customservice/getkflist"

const menuCreate = "/cgi-bin/menu/create"
const menuGet = "/cgi-bin/menu/get"
const menuDelete = "/cgi-bin/menu/delete"
const menuAddConditional = "/cgi-bin/menu/addconditional"
const menuDeleteConditional = "/cgi-bin/menu/delconditional"
const menuTryMatch = "/cgi-bin/menu/trymatch"

const templateAPISetIndustry = "/cgi-bin/template/api_set_industry"
const templateGetIndustry = "/cgi-bin/template/get_industry"
const templateAPIAddTemplate = "/cgi-bin/template/api_add_template"
const templateGetAllPrivateTemplate = "/cgi-bin/template/get_all_private_template"
const templateDelPrivateTemplate = "/cgi-bin/template/del_private_template"
const messageTemplateSend = "/cgi-bin/message/template/send"

const mediaUpload = "/cgi-bin/media/upload"
const mediaUploadImg = "/cgi-bin/media/uploadimg"
const mediaGet = "/cgi-bin/media/get"
const mediaGetJssdk = "/cgi-bin/media/get/jssdk"

const tagsCreate = "/cgi-bin/tags/create"
const tagsGet = "/cgi-bin/tags/get"
const tagsUpdate = "/cgi-bin/tags/update"
const tagsDelete = "/cgi-bin/tags/delete"

const tagsMembersBatchTagging = "/cgi-bin/tags/members/batchtagging"
const tagsMembersBatchUntagging = "/cgi-bin/tags/members/batchuntagging"
const tagsGetIDList = "/cgi-bin/tags/getidlist"
const tagsMembersGetBlackList = "/cgi-bin/tags/members/getblacklist"
const tagsMembersBatchBlackList = "/cgi-bin/tags/members/batchblacklist"
const tagsMembersBatchUnblackList = "/cgi-bin/tags/members/batchunblacklist"

const userTagGet = "/cgi-bin/user/tag/get"
const userInfoUpdateRemark = "/cgi-bin/user/info/updateremark"
const userInfo = "/cgi-bin/user/info"
const userInfoBatchGet = "/cgi-bin/user/info/batchget"
const userGet = "/cgi-bin/user/get"

const qrcodeCreate = "/cgi-bin/qrcode/create"
const showQrcode = "/cgi-bin/showqrcode"

const messageMassSend = "/cgi-bin/message/mass/send"
const messageMassSendall = "/cgi-bin/message/mass/sendall"
const messageMassPreview = "cgi-bin/message/mass/preview"
const messageMassDelete = "/cgi-bin/message/mass/delete"
const messageMassGet = "/cgi-bin/message/mass/get"

//DatacubeTimeLayout time format for datacube
const DatacubeTimeLayout = "2006-01-02"

// const tags_members_batchuntagging_URL_SUFFIX = "/cgi-bin/tags/members/batchuntagging"
// const tags_members_batchtagging_URL_SUFFIX = "/cgi-bin/tags/members/batchtagging"
// const tags_members_batchuntagging_URL_SUFFIX = "/cgi-bin/tags/members/batchuntagging"
const dataCubeGetUserSummary = "/datacube/getusersummary"
const dataCubeGetUserCumulate = "/datacube/getusercumulate"
const dataCubeGetArticleSummary = "/datacube/getarticlesummary"
const dataCubeGetArticleTotal = "/datacube/getarticletotal"
const dataCubeGetUserRead = "/datacube/getuserread"
const dataCubeGetUserReadHour = "/datacube/getuserreadhour"
const dataCubeGetUserShare = "/datacube/getusershare"
const dataCubeGetUserShareHour = "/datacube/getusersharehour"

const dataCubeGetUpstreamMsg = "/datacube/getupstreammsg"
const dataCubeGetUpstreamMsgHour = "/datacube/getupstreammsghour"
const dataCubeGetUpstreamMsgWeek = "/datacube/getupstreammsgweek"
const dataCubeGetUpstreamMsgDist = "/datacube/getupstreammsgdist"
const dataCubeGetUpstreamMsgMonth = "/datacube/getupstreammsgmonth"
const dataCubeGetUpstreamMsgDistWeek = "/datacube/getupstreammsgdistweek"
const dataCubeGetUpstreamMsgDistMonth = "/datacube/getupstreammsgdistmonth"
const dataCubeGetInterfaceSummary = "/datacube/getinterfacesummary"
const dataCubeGetInterfaceSummaryHour = "/datacube/getinterfacesummaryhour"

const materialAddNews = "/cgi-bin/material/add_news"
const materialAddMaterial = "/cgi-bin/material/add_material"
const materialGetMaterial = "/cgi-bin/material/get_material"
const materialDelMaterial = "/cgi-bin/material/del_material"
const materialUpdateNews = "/cgi-bin/material/update_news"
const materialGetMaterialcount = "/cgi-bin/material/get_materialcount"
const materialBatchgetMaterial = "/cgi-bin/material/batchget_material"
const commentOpen = "/cgi-bin/comment/open"
const commentClose = "/cgi-bin/comment/close"
const commentList = "/cgi-bin/comment/list"
const commentMarkelect = "/cgi-bin/comment/markelect"
const commentUnmarkelect = "/cgi-bin/comment/unmarkelect"
const commentDelete = "/cgi-bin/comment/delete"
const commentReplyAdd = "/cgi-bin/comment/reply/add"
const commentReplyDelete = "/cgi-bin/comment/reply/delete"

//const oauth2AccessToken = "/sns/oauth2/access_token"
const oauth2RefreshToken = "/sns/oauth2/refresh_token"
const oauth2Userinfo = "/sns/userinfo"
const oauth2Auth = "/sns/auth"
const defaultOauthRedirect = "/oauth_redirect"
const snsapiBase = "snsapi_base"
const snsapiUserinfo = "snsapi_userinfo"
const cardLandingPageCreate = "/card/landingpage/create"
const cardCodeDeposit = "/card/code/deposit"
const cardCodeGetDepositCount = "/card/code/getdepositcount"
const cardQrcodeCreate = "/card/qrcode/create"
const cardCodeCheckCode = "/card/code/checkcode"
const cardCodeGet = "/card/code/get"
const cardMPNewsGetHTML = "/card/mpnews/gethtml"
const cardTestWhiteListSet = "/card/testwhitelist/set"
const cardCreate = "/card/create"
const cardGet = "/card/get"
const cardGetApplyProtocol = "/card/getapplyprotocol"
const cardGetColors = "/card/getcolors"
const cardGetapplyprotocol = "/card/getapplyprotocol"
const cardBatchget = "/card/batchget"
const cardUpdate = "/card/update"
const cardDelete = "/card/delete"
const cardUserGetcardlist = "/card/user/getcardlist"
const cardPaycellSet = "card/paycell/set"
const cardModifystock = "card/modifystock"
const cardBoardingpassCheckin = "/card/boardingpass/checkin"
const poiAddPoi = "/cgi-bin/poi/addpoi"
const poiGetPoi = "/cgi-bin/poi/getpoi"
const poiUpdatePoi = "/cgi-bin/poi/updatepoi"
const poiGetListPoi = "/cgi-bin/poi/getpoilist"
const poiDelPoi = "/cgi-bin/poi/delpoi"
const poiGetWXCategory = "/cgi-bin/poi/getwxcategory"
const getCurrentAutoReplyInfo = "/cgi-bin/get_current_autoreply_info"
const getCurrentSelfMenuInfo = "/cgi-bin/get_current_selfmenu_info"

// POST ...
const POST = "POST"

// GET ...
const GET = "GET"
