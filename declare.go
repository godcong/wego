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
const clearQuotaURLSuffix = "/cgi-bin/clear_quota"
const getCallbackIPURLSuffix = "/cgi-bin/getcallbackip"

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

const getKFListURLSuffix = "/cgi-bin/customservice/getkflist"

const menuCreateURLSuffix = "/cgi-bin/menu/create"
const menuGetURLSuffix = "/cgi-bin/menu/get"
const menuDeleteURLSuffix = "/cgi-bin/menu/delete"
const menuAddConditionalURLSuffix = "/cgi-bin/menu/addconditional"
const menuDeleteConditionalURLSuffix = "/cgi-bin/menu/delconditional"
const menuTryMatchURLSuffix = "/cgi-bin/menu/trymatch"
const getCurrentSelfMenuInfoURLSuffix = "/cgi-bin/get_current_selfmenu_info"

const templateAPISetIndustryURLSuffix = "/cgi-bin/template/api_set_industry"
const templateGetIndustryURLSuffix = "/cgi-bin/template/get_industry"
const templateAPIAddTemplateURLSuffix = "/cgi-bin/template/api_add_template"
const templateGetAllPrivateTemplateURLSuffix = "/cgi-bin/template/get_all_private_template"
const templateDelPrivateTemplateURLSuffix = "/cgi-bin/template/del_private_template"
const messageTemplateSendURLSuffix = "/cgi-bin/message/template/send"

const mediaUploadURLSuffix = "/cgi-bin/media/upload"
const mediaUploadImgURLSuffix = "/cgi-bin/media/uploadimg"
const mediaGetURLSuffix = "/cgi-bin/media/get"
const mediaGetJssdkURLSuffix = "/cgi-bin/media/get/jssdk"

const tagsCreateURLSuffix = "/cgi-bin/tags/create"
const tagsGetURLSuffix = "/cgi-bin/tags/get"
const tagsUpdateURLSuffix = "/cgi-bin/tags/update"
const tagsDeleteURLSuffix = "/cgi-bin/tags/delete"

const tagsMembersBatchTaggingURLSuffix = "/cgi-bin/tags/members/batchtagging"
const tagsMembersBatchUntaggingURLSuffix = "/cgi-bin/tags/members/batchuntagging"
const tagsGetIDListURLSuffix = "/cgi-bin/tags/getidlist"
const tagsMembersGetBlackListURLSuffix = "/cgi-bin/tags/members/getblacklist"
const tagsMembersBatchBlackListURLSuffix = "/cgi-bin/tags/members/batchblacklist"
const tagsMembersBatchUnblackListURLSuffix = "/cgi-bin/tags/members/batchunblacklist"

const userTagGetURLSuffix = "/cgi-bin/user/tag/get"
const userInfoUpdateRemarkURLSuffix = "/cgi-bin/user/info/updateremark"
const userInfoURLSuffix = "/cgi-bin/user/info"
const userInfoBatchGetURLSuffix = "/cgi-bin/user/info/batchget"
const userGetURLSuffix = "/cgi-bin/user/get"

const qrcodeCreateURLSuffix = "/cgi-bin/qrcode/create"
const showQrcodeURLSuffix = "/cgi-bin/showqrcode"

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
const dataCubeGetUserSummaryURLSuffix = "/datacube/getusersummary"
const dataCubeGetUserCumulateURLSuffix = "/datacube/getusercumulate"
const dataCubeGetArticleSummaryURLSuffix = "/datacube/getarticlesummary"
const dataCubeGetArticleTotalURLSuffix = "/datacube/getarticletotal"
const dataCubeGetUserReadURLSuffix = "/datacube/getuserread"
const dataCubeGetUserReadHourURLSuffix = "/datacube/getuserreadhour"
const dataCubeGetUserShareURLSuffix = "/datacube/getusershare"
const dataCubeGetUserShareHourURLSuffix = "/datacube/getusersharehour"

const dataCubeGetUpstreamMsgURLSuffix = "/datacube/getupstreammsg"
const dataCubeGetUpstreamMsgHourURLSuffix = "/datacube/getupstreammsghour"
const dataCubeGetUpstreamMsgWeekURLSuffix = "/datacube/getupstreammsgweek"
const dataCubeGetUpstreamMsgDistURLSuffix = "/datacube/getupstreammsgdist"
const dataCubeGetUpstreamMsgMonthURLSuffix = "/datacube/getupstreammsgmonth"
const dataCubeGetUpstreamMsgDistWeekURLSuffix = "/datacube/getupstreammsgdistweek"
const dataCubeGetUpstreamMsgDistMonthURLSuffix = "/datacube/getupstreammsgdistmonth"
const dataCubeGetInterfaceSummaryURLSuffix = "/datacube/getinterfacesummary"
const dataCubeGetInterfaceSummaryHourURLSuffix = "/datacube/getinterfacesummaryhour"

const materialAddNewsURLSuffix = "/cgi-bin/material/add_news"
const materialAddMaterialURLSuffix = "/cgi-bin/material/add_material"
const materialGetMaterialURLSuffix = "/cgi-bin/material/get_material"
const materialDelMaterialURLSuffix = "/cgi-bin/material/del_material"
const materialUpdateNewsURLSuffix = "/cgi-bin/material/update_news"
const materialGetMaterialcountURLSuffix = "/cgi-bin/material/get_materialcount"
const materialBatchgetMaterialURLSuffix = "/cgi-bin/material/batchget_material"
const commentOpenURLSuffix = "/cgi-bin/comment/open"
const commentCloseURLSuffix = "/cgi-bin/comment/close"
const commentListURLSuffix = "/cgi-bin/comment/list"
const commentMarkelectURLSuffix = "/cgi-bin/comment/markelect"
const commentUnmarkelectURLSuffix = "/cgi-bin/comment/unmarkelect"
const commentDeleteURLSuffix = "/cgi-bin/comment/delete"
const commentReplyAddURLSuffix = "/cgi-bin/comment/reply/add"
const commentReplyDeleteURLSuffix = "/cgi-bin/comment/reply/delete"

const oauth2AccessTokenURLSuffix = "/sns/oauth2/access_token"
const oauth2RefreshTokenURLSuffix = "/sns/oauth2/refresh_token"
const oauth2UserinfoURLSuffix = "/sns/userinfo"
const oauth2AuthURLSuffix = "/sns/auth"
const oauth2AuthorizeURLSuffix = "https://open.weixin.qq.com/connect/oauth2/authorize"
const defaultOauthRedirectURLSuffix = "/oauth_redirect"
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

// CardScene ...
type CardScene string

// CardSceneNearBy ...
const (
	CardSceneNearBy         CardScene = "SCENE_NEAR_BY"          //CardSceneNearBy 附近
	CardSceneMenu           CardScene = "SCENE_MENU"             //CardSceneMenu 自定义菜单
	CardSceneQrcode         CardScene = "SCENE_QRCODE"           //CardSceneQrcode 二维码
	CardSceneArticle        CardScene = "SCENE_ARTICLE"          //CardSceneArticle 公众号文章
	CardSceneH5             CardScene = "SCENE_H5"               //CardSceneH5 H5页面
	CardSceneIvr            CardScene = "SCENE_IVR"              //CardSceneIvr 自动回复
	CardSceneCardCustomCell CardScene = "SCENE_CARD_CUSTOM_CELL" //CardSceneCardCustomCell 卡券自定义cell
)

//支持开发者拉出指定状态的卡券列表
type CardStatus string

// CARD_STATUS_NOT_VERIFY ...
const (
	CARD_STATUS_NOT_VERIFY  CardStatus = "CARD_STATUS_NOT_VERIFY"  //待审核
	CARD_STATUS_VERIFY_FAIL CardStatus = "CARD_STATUS_VERIFY_FAIL" //审核失败
	CARD_STATUS_VERIFY_OK   CardStatus = "CARD_STATUS_VERIFY_OK"   //通过审核
	CARD_STATUS_DELETE      CardStatus = "CARD_STATUS_DELETE"      //卡券被商户删除
	CARD_STATUS_DISPATCH    CardStatus = "CARD_STATUS_DISPATCH"    //在公众平台投放过的卡券；
)
