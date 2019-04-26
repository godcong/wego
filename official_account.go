package wego

import (
	"bytes"
	"context"
	"github.com/godcong/wego/util"
	log "github.com/sirupsen/logrus"
	"strings"
)

// OfficialAccount ...
type OfficialAccount struct {
	*OfficialAccountProperty
	BodyType    BodyType
	oauth       OAuthProperty
	client      *Client
	jssdk       *JSSDK
	accessToken *AccessToken
	remoteURL   string
	localHost   string
}

// NewOfficialAccount ...
func NewOfficialAccount(config *OfficialAccountProperty, options ...OfficialAccountOption) *OfficialAccount {
	officialAccount := &OfficialAccount{
		OfficialAccountProperty: config,
		BodyType:                BodyTypeJSON,
	}
	officialAccount.parse(options...)

	return officialAccount
}

func (obj *OfficialAccount) parse(options ...OfficialAccountOption) {
	if options == nil {
		return
	}
	for _, o := range options {
		o(obj)
	}
}

// Client ...
func (obj *OfficialAccount) Client() *Client {
	if obj.client == nil {
		obj.client = NewClient(ClientBodyType(obj.BodyType), ClientAccessToken(obj.accessToken))
	}
	return obj.client
}

// HandleAuthorizeNotify ...
func (obj *OfficialAccount) HandleAuthorizeNotify(hooks ...interface{}) ServeHTTPFunc {
	return obj.HandleAuthorize(hooks...).ServeHTTP
}

// HandleAuthorize ...
func (obj *OfficialAccount) HandleAuthorize(hooks ...interface{}) Notifier {
	notify := &authorizeNotify{
		OfficialAccount: obj,
	}
	for _, hook := range hooks {
		switch h := hook.(type) {
		case TokenHook:
			notify.TokenHook = h
		case UserHook:
			notify.UserHook = h
		case StateHook:
			notify.StateHook = h
		}
	}
	return notify
}

// GetUserInfo ...
func (obj *OfficialAccount) GetUserInfo(token *Token) (*WechatUser, error) {
	var info WechatUser
	var e error
	p := util.Map{
		"access_token": token.AccessToken,
		"openid":       token.OpenID,
		"lang":         "zh_CN",
	}
	responder := Get(
		snsUserinfo,
		p,
	)
	log.Debug("WechatUser|responder", string(responder.Bytes()), responder.Error())
	e = responder.Error()
	if e != nil {
		return &info, e
	}

	e = responder.Unmarshal(&info)
	if e != nil {
		return &info, e
	}
	return &info, nil
}

// Oauth2AuthorizeToken ...
func (obj *OfficialAccount) Oauth2AuthorizeToken(code string) (*Token, error) {
	var token Token
	var e error

	p := util.Map{
		"appid":      obj.AppID,
		"secret":     obj.AppSecret,
		"code":       code,
		"grant_type": "authorization_code",
	}

	uri := obj.RedirectURI()
	if uri != "" {
		p.Set("redirect_uri", uri)
	}

	responder := PostJSON(
		oauth2AccessToken,
		p,
		nil,
	)
	e = responder.Error()
	log.Debug("GetAuthorizeToken|response", string(responder.Bytes()), e)
	if e != nil {
		return &token, e
	}

	e = responder.Unmarshal(&token)
	if e != nil {
		return &token, e
	}
	return &token, nil
}

/*AuthCodeURL 生成授权地址URL*/
func (obj *OfficialAccount) AuthCodeURL(state string) string {
	log.Debug("AuthCodeURL", state)
	var buf bytes.Buffer
	buf.WriteString(oauth2Authorize)
	p := util.Map{
		"response_type": "code",
		"appid":         obj.AppID,
	}

	uri := obj.RedirectURI()
	if uri != "" {
		p.Set("redirect_uri", uri)
	}

	if obj.oauth.Scopes != nil {
		p.Set("scope", obj.oauth.Scopes)
	}
	if state != "" {
		// TODO(light): Docs say never to omit state; don't allow empty.
		p.Set("state", state)
	}
	buf.WriteByte('?')
	buf.WriteString(p.URLEncode())
	return buf.String()
}

// RedirectURI ...
func (obj *OfficialAccount) RedirectURI() string {
	log.Debug("RedirectURI", obj.oauth.RedirectURI)
	if obj.oauth.RedirectURI != "" {
		if strings.Index(obj.oauth.RedirectURI, "http") == 0 {
			return obj.oauth.RedirectURI
		}
		return util.URL(obj.localHost, obj.oauth.RedirectURI)
	}
	return ""
}

// RemoteURL ...
func (obj *OfficialAccount) RemoteURL() string {
	return obj.remoteURL
}

/*ClearQuota 公众号的所有api调用（包括第三方帮其调用）次数进行清零
HTTP请求方式:POST
HTTP调用: https://api.weixin.qq.com/cgi-bin/clear_quota?access_token=ACCESS_TOKEN
*/
func (obj *OfficialAccount) ClearQuota() Responder {
	url := util.URL(obj.RemoteURL(), clearQuotaURLSuffix)
	token := obj.accessToken.GetToken()

	params := util.Map{
		"appid": obj.AppID,
	}
	return PostJSON(url, token.KeyMap(), params)

}

/*GetCallbackIP 请求微信的服务器IP列表
HTTP请求方式: GET
HTTP调用:https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN
*/
func (obj *OfficialAccount) GetCallbackIP() Responder {
	url := util.URL(obj.RemoteURL(), getCallbackIPURLSuffix)
	return obj.Client().Get(context.Background(), url, nil)
}

//MessageSend 根据OpenID列表群发【订阅号不可用，服务号认证后可用】
//接口调用请求说明
//http请求方式: POST
//https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN
func (obj *OfficialAccount) MessageSend(msg util.Map) Responder {
	url := util.URL(obj.RemoteURL(), messageMassSend)
	return obj.Client().Post(context.Background(), url, nil, msg)
}

//MessageSendAll 根据标签进行群发【订阅号与服务号认证后均可用】
//接口调用请求说明
//http请求方式: POST
//https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN
func (obj *OfficialAccount) MessageSendAll(msg util.Map) Responder {
	url := util.URL(obj.RemoteURL(), messageMassSendall)
	return obj.Client().Post(context.Background(), url, nil, msg)
}

//MessagePreview 预览接口【订阅号与服务号认证后均可用】
//开发者可通过该接口发送消息给指定用户，在手机端查看消息的样式和排版。为了满足第三方平台开发者的需求，在保留对openID预览能力的同时，增加了对指定微信号发送预览的能力，但该能力每日调用次数有限制（100次），请勿滥用。
//接口调用请求说明
//http请求方式: POST
//https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token=ACCESS_TOKEN
func (obj *OfficialAccount) MessagePreview(msg util.Map) Responder {
	url := util.URL(obj.RemoteURL(), messageMassPreview)
	return obj.Client().Post(context.Background(), url, nil, msg)

}

//MessageDelete 删除群发【订阅号与服务号认证后均可用】
//群发之后，随时可以通过该接口删除群发。
//接口调用请求说明
//http请求方式: POST
//https://api.weixin.qq.com/cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN
func (obj *OfficialAccount) MessageDelete(msgID string) Responder {
	url := util.URL(obj.RemoteURL(), messageMassDelete)
	return obj.Client().Post(context.Background(), url, nil, util.Map{"msg_id": msgID})

}

//MessageStatus 查询群发消息发送状态【订阅号与服务号认证后均可用】
//接口调用请求说明
//http请求方式: POST
//https://api.weixin.qq.com/cgi-bin/message/mass/get?access_token=ACCESS_TOKEN
func (obj *OfficialAccount) MessageStatus(msgID string) Responder {
	url := util.URL(obj.RemoteURL(), messageMassGet)
	return obj.Client().Post(context.Background(), url, nil, util.Map{"msg_id": msgID})

}

// MessageSendText ...
func (obj *OfficialAccount) MessageSendText() {

}

//CardCreateLandingPage 创建货架接口
//	HTTP请求方式: POST
//	URL:https://api.weixin.qq.com/card/landingpage/create?access_token=$TOKEN
//	func (c *OfficialAccount) CreateLandingPage(page *CardLandingPage) Responder {
func (obj *OfficialAccount) CardCreateLandingPage(p util.Map) Responder {
	url := util.URL(obj.RemoteURL(), cardLandingPageCreate)
	return obj.Client().Post(context.Background(), url, nil, p)

}

//CardDeposit 导入code接口
//	HTTP请求方式: POST
//	URL:http://api.weixin.qq.com/card/code/deposit?access_token=ACCESS_TOKEN
func (obj *OfficialAccount) CardDeposit(cardID string, code []string) Responder {
	url := util.URL(obj.RemoteURL(), cardCodeDeposit)
	return obj.Client().Post(context.Background(), url, nil, util.Map{
		"card_id": cardID,
		"code":    code,
	})

}

//CardGetDepositCount 查询导入code数目
//
//  HTTP请求方式: POST
//  URL:http://api.weixin.qq.com/card/code/getdepositcount?access_token=ACCESS_TOKEN
func (obj *OfficialAccount) CardGetDepositCount(cardID string) Responder {
	url := util.URL(obj.RemoteURL(), cardCodeGetDepositCount)
	return obj.Client().Post(context.Background(), url, nil, util.Map{
		"card_id": cardID,
	})
}

//CardCheckCode 核查code接口
//	HTTP请求方式: POST
//	HTTP调用:http://api.weixin.qq.com/card/code/checkcode?access_token=ACCESS_TOKEN
func (obj *OfficialAccount) CardCheckCode(cardID string, code []string) Responder {
	url := util.URL(obj.RemoteURL(), cardCodeCheckCode)
	return obj.Client().Post(context.Background(), url, nil, util.Map{
		"card_id": cardID,
		"code":    code,
	})
}

//CardGetCode 查询Code接口
//	HTTP请求方式: POST
//	HTTP调用:https://api.weixin.qq.com/card/code/get?access_token=TOKEN
//	参数说明:
//	参数名	必填	类型	示例值	描述
//	code	是	string(20)	110201201245	单张卡券的唯一标准。
//	card_id	否	string(32)	pFS7Fjg8kV1I dDz01r4SQwMkuCKc	卡券ID代表一类卡券。自定义code卡券必填。
//	check_consume	否	bool	true	是否校验code核销状态，填入true和false时的code异常状态返回数据不同。
func (obj *OfficialAccount) CardGetCode(p util.Map) Responder {
	url := util.URL(obj.RemoteURL(), cardCodeGet)
	return obj.Client().Post(context.Background(), url, nil, p)
}

//CardGetHTML 图文消息群发卡券
//	HTTP请求方式: POST
//	URL:https://api.weixin.qq.com/card/mpnews/gethtml?access_token=TOKEN
func (obj *OfficialAccount) CardGetHTML(cardID string) Responder {
	url := util.URL(obj.RemoteURL(), cardMPNewsGetHTML)
	return obj.Client().Post(context.Background(), url, nil, util.Map{
		"card_id": cardID,
	})
}

//CardSetTestWhiteListByID 设置测试白名单(by openid)
func (obj *OfficialAccount) CardSetTestWhiteListByID(list ...string) Responder {
	return obj.CardSetTestWhiteList("openid", list)
}

//CardSetTestWhiteListByName 设置测试白名单(by username)
func (obj *OfficialAccount) CardSetTestWhiteListByName(list ...string) Responder {
	return obj.CardSetTestWhiteList("username", list)
}

//CardSetTestWhiteList 设置测试白名单
//	HTTP请求方式: POST
//	URL:https://api.weixin.qq.com/card/testwhitelist/set?access_token=TOKEN
func (obj *OfficialAccount) CardSetTestWhiteList(typ string, list []string) Responder {
	url := util.URL(obj.RemoteURL(), cardTestWhiteListSet)
	return obj.Client().Post(context.Background(), url, nil, util.Map{
		typ: list,
	})
}

//CardCreateQrCode 创建二维码
//	HTTP请求方式: POST
//	URL:https://api.weixin.qq.com/card/qrcode/create?access_token=TOKEN
func (obj *OfficialAccount) CardCreateQrCode(action *QrCodeAction) Responder {
	url := util.URL(obj.RemoteURL(), cardQrcodeCreate)
	return obj.Client().Post(context.Background(), url, nil, action)
}

//CardCreate 创建卡券
//	HTTP请求方式: POST
//	URL: https://api.weixin.qq.com/card/create?access_token=ACCESS_TOKEN
//	type *OneCard or Map
func (obj *OfficialAccount) CardCreate(maps util.MapAble) Responder {
	url := util.URL(obj.RemoteURL(), cardCreate)
	return obj.Client().Post(context.Background(), url, nil, util.Map{"card": maps})
}

//CardGet 查看卡券详情
//	开发者可以调用该接口查询某个card_id的创建信息、审核状态以及库存数量。
//	接口调用请求说明
//	HTTP请求方式: POSTURL:https://api.weixin.qq.com/card/get?access_token=TOKEN
func (obj *OfficialAccount) CardGet(cardID string) Responder {
	url := util.URL(obj.RemoteURL(), cardGet)
	return obj.Client().Post(context.Background(), url, nil, util.Map{"card_id": cardID})
}

//CardGetApplyProtocol 卡券开放类目查询接口
//	HTTP请求方式: GET
//	URL:https://api.weixin.qq.com/card/getapplyprotocol?access_token=TOKEN
func (obj *OfficialAccount) CardGetApplyProtocol() Responder {
	url := util.URL(obj.RemoteURL(), cardGetApplyProtocol)
	return obj.Client().Get(context.Background(), url, nil)
}

//CardGetColors 卡券开放类目查询接口
//	HTTP请求方式: GET
//	URL:https://api.weixin.qq.com/card/getcolors?access_token=TOKEN
func (obj *OfficialAccount) CardGetColors() Responder {
	url := util.URL(obj.RemoteURL(), cardGetColors)
	return obj.Client().Get(context.Background(), url, nil)
}

//CardCheckin 更新飞机票信息接口
//	接口调用请求说明
//	http请求方式: POST
//	URL:https://api.weixin.qq.com/card/boardingpass/checkin?access_token=TOKEN
func (obj *OfficialAccount) CardCheckin(p util.Map) Responder {
	url := util.URL(obj.RemoteURL(), cardBoardingpassCheckin)
	return obj.Client().Post(context.Background(), url, nil, p)
}

//CardCategories 卡券开放类目查询接口
//	接口说明
//	通过调用该接口查询卡券开放的类目ID，类目会随业务发展变更，请每次用接口去查询获取实时卡券类目。
//	注意：
//	1.本接口查询的返回值还有卡券资质ID,此处的卡券资质为：已微信认证的公众号通过微信公众平台申请卡券功能时，所需的资质。
//	2.对于第三方强授权模式，子商户无论选择什么类目，均提交营业执照即可，所以不用考虑此处返回的资质字段，返回值仅参考类目ID即可。
//	接口详情
//	接口调用请求说明
//	https请求方式: GET https://api.weixin.qq.com/card/getapplyprotocol?access_token=TOKEN
func (obj *OfficialAccount) CardCategories() Responder {
	url := util.URL(obj.RemoteURL(), cardGetapplyprotocol)
	return obj.Client().Get(context.Background(), url, nil)
}

//CardBatchGet 批量查询卡券列表
//	接口调用请求说明
//	HTTP请求方式: POST URL:https://api.weixin.qq.com/card/batchget?access_token=TOKEN
func (obj *OfficialAccount) CardBatchGet(offset, count int, statusList []CardStatus) Responder {
	p := util.Map{
		"offset":      offset,
		"count":       count,
		"status_list": statusList,
	}
	url := util.URL(obj.RemoteURL(), cardBatchget)
	return obj.Client().Post(context.Background(), url, nil, p)
}

//CardUpdate 更改卡券信息接口
//	接口说明
//	支持更新所有卡券类型的部分通用字段及特殊卡券（会员卡、飞机票、电影票、会议门票）中特定字段的信息。
//	接口调用请求说明
//	HTTP请求方式: POST URL:https://api.weixin.qq.com/card/update?access_token=TOKEN
func (obj *OfficialAccount) CardUpdate(cardID string, p util.Map) Responder {
	p = util.CombineMaps(util.Map{
		"card_id": cardID,
	}, p)
	url := util.URL(obj.RemoteURL(), cardUpdate)
	return obj.Client().Post(context.Background(), url, nil, p)
}

//CardDelete 删除卡券接口
//删除卡券接口允许商户删除任意一类卡券。删除卡券后，该卡券对应已生成的领取用二维码、添加到卡包JS API均会失效。 注意：如用户在商家删除卡券前已领取一张或多张该卡券依旧有效。即删除卡券不能删除已被用户领取，保存在微信客户端中的卡券。
//接口调用请求说明
//HTTP请求方式: POST URL:https://api.weixin.qq.com/card/delete?access_token=TOKEN
func (obj *OfficialAccount) CardDelete(cardID string) Responder {
	url := util.URL(obj.RemoteURL(), cardDelete)
	return obj.Client().Post(context.Background(), url, nil, util.Map{
		"card_id": cardID,
	})
}

// CardGetUserCards ...
func (obj *OfficialAccount) CardGetUserCards(openID, cardID string) Responder {
	url := util.URL(obj.RemoteURL(), cardUserGetcardlist)
	return obj.Client().Post(context.Background(), url, nil, util.Map{
		"openid":  openID,
		"card_id": cardID,
	})
}

// CardSetPayCell ...
func (obj *OfficialAccount) CardSetPayCell(cardID string, isOpen bool) Responder {
	url := util.URL(obj.RemoteURL(), cardPaycellSet)
	return obj.Client().Post(context.Background(), url, nil, util.Map{
		"is_open": isOpen,
		"card_id": cardID,
	})
}

// CardModifyStock ...
func (obj *OfficialAccount) CardModifyStock(cardID string, option util.Map) Responder {
	url := util.URL(obj.RemoteURL(), cardModifystock)
	return obj.Client().Post(context.Background(), url, nil, util.Map{
		"card_id": cardID,
	})
}

//CardGetCardAPITicket get ticket
func (obj *OfficialAccount) CardGetCardAPITicket(refresh bool) {
	obj.jssdk.GetTicket("wx_card", refresh)
}

// JSSDK ...
func (obj *OfficialAccount) JSSDK() *JSSDK {
	//TODO:need fix
	return obj.jssdk
}

/*
CommentOpen 打开文章评论
 https 请求方式: POST
 https://api.weixin.qq.com/cgi-bin/comment/open?access_token=ACCESS_TOKEN
 失败:
  {"errcode":88000,"errmsg":"without comment privilege"}
*/
func (obj *OfficialAccount) CommentOpen(id, index int) Responder {
	url := util.URL(obj.RemoteURL(), commentOpenURLSuffix)
	return obj.Client().Post(context.Background(), url, nil, util.Map{
		"msg_data_id": id,
		"index":       index,
	})
}

/*
CommentClose 关闭评论
 https 请求方式: POST
 https://api.weixin.qq.com/cgi-bin/comment/close?access_token=ACCESS_TOKEN
 失败:
 {"errcode":88000,"errmsg":"without comment privilege"}
*/
func (obj *OfficialAccount) CommentClose(id, index int) Responder {
	url := util.URL(obj.RemoteURL(), commentCloseURLSuffix)
	return obj.Client().Post(context.Background(), url, nil, util.Map{
		"msg_data_id": id,
		"index":       index,
	})
}
