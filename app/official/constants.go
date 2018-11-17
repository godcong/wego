package official

const domain = "https://api.weixin.qq.com"

const clearQuotaURLSuffix = "/cgi-bin/clear_quota"
const getCallbackIPURLSuffix = "/cgi-bin/getcallbackip"

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

//oauth2AccessTokenURLSuffix api address suffix
const oauth2AccessTokenURLSuffix = "/sns/oauth2/access_token"

//oauth2RefreshTokenURLSuffix api address suffix
const oauth2RefreshTokenURLSuffix = "/sns/oauth2/refresh_token"

//oauth2UserinfoURLSuffix api address suffix
const oauth2UserinfoURLSuffix = "/sns/userinfo"

//oauth2AuthURLSuffix api address suffix
const oauth2AuthURLSuffix = "/sns/auth"

//oauth2AuthorizeURLSuffix api address suffix
const oauth2AuthorizeURLSuffix = "https://open.weixin.qq.com/connect/oauth2/authorize"

//defaultOauthRedirectURLSuffix api address suffix
const defaultOauthRedirectURLSuffix = "/oauth_redirect"

//snsapiBase api address suffix
const snsapiBase = "snsapi_base"

//snsapiUserinfo api address suffix
const snsapiUserinfo = "snsapi_userinfo"

//ticketGetTicket api address suffix
const ticketGetTicket = "/cgi-bin/ticket/getticket"

//cardLandingPageCreate api address suffix
const cardLandingPageCreate = "/card/landingpage/create"

//cardCodeDeposit api address suffix
const cardCodeDeposit = "/card/code/deposit"

//cardCodeGetDepositCount api address suffix
const cardCodeGetDepositCount = "/card/code/getdepositcount"

//cardQrcodeCreate api address suffix
const cardQrcodeCreate = "/card/qrcode/create"

//cardCodeCheckCode api address suffix
const cardCodeCheckCode = "/card/code/checkcode"

//cardCodeGet api address suffix
const cardCodeGet = "/card/code/get"

//cardMPNewsGetHTML api address suffix
const cardMPNewsGetHTML = "/card/mpnews/gethtml"

//cardTestWhiteListSet api address suffix
const cardTestWhiteListSet = "/card/testwhitelist/set"

//cardCreate api address suffix
const cardCreate = "/card/create"

//cardGetApplyProtocol api address suffix
const cardGetApplyProtocol = "/card/getapplyprotocol"

//cardGetColors api address suffix
const cardGetColors = "/card/getcolors"

//cardGetapplyprotocol api address suffix
const cardGetapplyprotocol = "/card/getapplyprotocol"

//cardBoardingpassCheckin api address suffix
const cardBoardingpassCheckin = "/card/boardingpass/checkin"

//poiAddPoi api address suffix
const poiAddPoi = "/cgi-bin/poi/addpoi"

//poiGetPoi api address suffix
const poiGetPoi = "/cgi-bin/poi/getpoi"

//poiUpdatePoi api address suffix
const poiUpdatePoi = "/cgi-bin/poi/updatepoi"

//poiGetListPoi api address suffix
const poiGetListPoi = "/cgi-bin/poi/getpoilist"

//poiDelPoi api address suffix
const poiDelPoi = "/cgi-bin/poi/delpoi"

//poiGetWXCategory api address suffix
const poiGetWXCategory = "/cgi-bin/poi/getwxcategory"

//getCurrentAutoReplyInfo api address suffix
const getCurrentAutoReplyInfo = "/cgi-bin/get_current_autoreply_info"

//getCurrentSelfMenuInfo api address suffix
const getCurrentSelfMenuInfo = "/cgi-bin/get_current_selfmenu_info"
